package service

import (
	"context"

	"github.com/leonzag/treport/internal/application/converter"
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/application/interfaces"
	"github.com/leonzag/treport/internal/domain/entity"
)

var _ interfaces.TokenService = new(tokenService)

type tokenService struct {
	tokenRepo     interfaces.TokenRepo
	cryptoService interfaces.CryptoService
}

func NewTokenService(
	tokenRepo interfaces.TokenRepo,
	cryptoService interfaces.CryptoService,
) *tokenService {
	return &tokenService{
		tokenRepo:     tokenRepo,
		cryptoService: cryptoService,
	}
}

func (s *tokenService) encryptAndHash(t dto.TokenDTO) (dto.TokenDTO, error) {
	if t.Password == "" {
		return t, nil
	}
	encToken, err := s.cryptoService.EncryptToken(t.Password, t.Token)
	if err != nil {
		return t, err
	}
	hashedPwd, err := s.cryptoService.HashPassword(t.Password)
	if err != nil {
		return t, err
	}
	t.Password = hashedPwd
	t.Token = encToken
	return t, nil
}

func (s *tokenService) decryptAndUnhash(t dto.TokenDTO, pwd string) (dto.TokenDTO, error) {
	if !s.cryptoService.CheckPassword(t.Password, pwd) {
		return t, entity.ErrTokenIncorrectPassword
	}
	if t.Password == "" {
		return t, nil
	}
	var err error
	t.Token, err = s.cryptoService.DecryptToken(pwd, t.Token)
	if err != nil {
		return t, err
	}
	t.Password = pwd
	return t, nil
}

func (s *tokenService) AddToken(ctx context.Context, tokenDTO dto.TokenDTO) error {
	tokenDTO, err := s.encryptAndHash(tokenDTO)
	if err != nil {
		return err
	}
	token := converter.ToTokenFromDTO(tokenDTO)
	return s.tokenRepo.Add(ctx, token)
}

func (s *tokenService) DeleteToken(ctx context.Context, tokenDTO dto.TokenDTO) error {
	token := converter.ToTokenFromDTO(tokenDTO)
	return s.tokenRepo.Delete(ctx, token)
}

func (s *tokenService) UpdateToken(ctx context.Context, tokenDTO dto.TokenDTO) error {
	tokenDTO, err := s.encryptAndHash(tokenDTO)
	if err != nil {
		return err
	}
	token := converter.ToTokenFromDTO(tokenDTO)
	return s.tokenRepo.Update(ctx, token)
}

func (s *tokenService) GetTokenByTitle(ctx context.Context, title string) (dto.TokenDTO, error) {
	token, error := s.tokenRepo.Get(ctx, title)
	if error != nil {
		return dto.TokenDTO{}, error
	}

	return converter.ToTokenDTOFromService(token), nil
}

func (s *tokenService) GetTokenByTitleDecrypted(ctx context.Context, title string, pwd string) (dto.TokenDTO, error) {
	token, err := s.GetTokenByTitle(ctx, title)
	if err != nil {
		return dto.TokenDTO{}, err
	}
	token, err = s.decryptAndUnhash(token, pwd)
	if err != nil {
		return dto.TokenDTO{}, err
	}
	return token, nil
}

func (s *tokenService) ListTokens(ctx context.Context) ([]dto.TokenDTO, error) {
	tokens, err := s.tokenRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	tokenDTOs := make([]dto.TokenDTO, 0, len(tokens))
	for _, token := range tokens {
		tokenDTOs = append(tokenDTOs, converter.ToTokenDTOFromService(token))
	}

	return tokenDTOs, nil
}

func (s *tokenService) ListTokensTitles(ctx context.Context) ([]string, error) {
	tokens, err := s.ListTokens(ctx)
	if err != nil {
		return nil, err
	}
	titles := make([]string, 0, len(tokens))
	for _, token := range tokens {
		titles = append(titles, token.Title)
	}

	return titles, nil
}

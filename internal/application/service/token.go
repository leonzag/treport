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

func (s *tokenService) AddToken(ctx context.Context, t dto.TokenRequestDTO) (*entity.Token, error) {
	_, err := s.tokenRepo.Get(ctx, t.Title)
	if err == nil {
		return nil, entity.ErrTokenExist
	}
	tokenEnc, err := s.encryptAndHash(t)
	if err != nil {
		return nil, err
	}
	token := converter.ToTokenFromDTO(tokenEnc)

	return s.tokenRepo.Add(ctx, token)
}

func (s *tokenService) DecryptToken(t *entity.Token, pwd string) (*entity.TokenDecrypted, error) {
	if !s.cryptoService.CheckPassword(t.Password, pwd) {
		return nil, entity.ErrTokenIncorrectPassword
	}

	var err error
	token := &entity.TokenDecrypted{Title: t.Title}

	if t.Password == "" {
		token.Token = t.Token
		return token, nil
	}
	token.Token, err = s.cryptoService.DecryptToken(pwd, t.Token)

	return token, err
}

func (s *tokenService) DeleteToken(ctx context.Context, tokenDTO dto.TokenRequestDTO) error {
	if _, err := s.GetTokenByTitle(ctx, tokenDTO.Token); err == nil {
		return entity.ErrTokenNotFound
	}
	token := converter.ToTokenFromRequestDTO(tokenDTO)

	return s.tokenRepo.Delete(ctx, token)
}

func (s *tokenService) UpdateToken(ctx context.Context, t dto.TokenRequestDTO) (*entity.Token, error) {
	if _, err := s.GetTokenByTitle(ctx, t.Token); err == nil {
		return nil, entity.ErrTokenNotFound
	}
	tokenEnc, err := s.encryptAndHash(t)
	if err != nil {
		return nil, err
	}
	token := converter.ToTokenFromDTO(tokenEnc)

	return s.tokenRepo.Update(ctx, token)
}

func (s *tokenService) GetTokenByTitle(ctx context.Context, title string) (*entity.Token, error) {
	token, err := s.tokenRepo.Get(ctx, title)
	if err != nil {
		return nil, entity.ErrTokenNotFound
	}

	return token, nil
}

func (s *tokenService) ListTokens(ctx context.Context) ([]*entity.Token, error) {
	return s.tokenRepo.List(ctx)
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

func (s *tokenService) encryptAndHash(t dto.TokenRequestDTO) (dto.TokenDTO, error) {
	token := dto.TokenDTO{
		Title:       t.Title,
		Description: t.Description,
		Token:       t.Token,
	}
	if t.Password == "" {
		return token, nil
	}

	var err error

	token.Token, err = s.cryptoService.EncryptToken(t.Password, t.Token)
	if err != nil {
		return token, err
	}
	token.Password, err = s.cryptoService.HashPassword(t.Password)
	if err != nil {
		return token, err
	}

	return token, nil
}

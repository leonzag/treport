package service

import (
	"github.com/leonzag/treport/internal/application/interfaces"
	"github.com/leonzag/treport/pkg/crypto"
)

var _ interfaces.CryptoService = new(cryptoService)

type cryptoService struct{}

func NewCryptoService() *cryptoService {
	return &cryptoService{}
}

func (s *cryptoService) HashPassword(pwd string) (string, error) {
	return crypto.HashPassword(pwd)
}

func (s *cryptoService) CheckPassword(hashed string, pwd string) bool {
	return crypto.CheckPassword(hashed, pwd)
}

func (s *cryptoService) EncryptToken(pwd string, token string) (string, error) {
	return crypto.EncryptToken(pwd, token)
}

func (s *cryptoService) DecryptToken(pwd string, encryptedToken string) (string, error) {
	return crypto.DecryptToken(pwd, encryptedToken)
}

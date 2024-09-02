package converter

import (
	"github.com/leonzag/treport/internal/application/dto"
	"github.com/leonzag/treport/internal/domain/entity"
)

func ToTokenFromDTO(token dto.TokenDTO) *entity.Token {
	return &entity.Token{
		Title:    token.Title,
		Password: token.Password,
		Token:    token.Token,
	}
}

func ToTokenDTOFromService(token *entity.Token) dto.TokenDTO {
	return dto.TokenDTO{
		Title:    token.Title,
		Password: token.Password,
		Token:    token.Token,
	}
}

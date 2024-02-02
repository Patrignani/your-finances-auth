package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IRefreshTokenSerice interface {
	CreateRefreshToken(userID string) (*entity.RefreshToken, error)
	FindById(refreshTokenId string) (*entity.RefreshToken, error)
}

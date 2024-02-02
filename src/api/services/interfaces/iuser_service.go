package interfaces

import (
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	"github.com/labstack/echo/v4"
)

type IUserService interface {
	Authenticate(ctx echo.Context, username string, password string) (*entity.User, error)
	FindById(userId string) (*entity.User, error)
}

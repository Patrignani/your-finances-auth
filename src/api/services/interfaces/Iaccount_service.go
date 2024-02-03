package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IAccountService interface {
	FindByUserId(userID string) (*entity.Account, error)
}

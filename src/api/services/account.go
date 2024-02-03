package services

import (
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	repository "github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/specifications"
	"github.com/Patrignani/your-finances-auth/src/api/services/interfaces"
)

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) interfaces.IAccountService {
	return &AccountService{repository: repository}
}

func (a *AccountService) FindByUserId(userID string) (*entity.Account, error) {
	project := map[string]int{
		"_id":        1,
		"user_id":    1,
		"account_id": 1,
	}

	specification := specifications.NewFindByOneAccountByUserId(userID, project)

	return a.repository.FindOneBySpecification(specification)
}

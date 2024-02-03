package repositories

import (
	"context"
	"time"

	"github.com/Patrignani/your-finances-auth/src/api/data"
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
)

const (
	accountCollection = "accounts"
)

type AccountRepository struct {
	context data.MongoDB
}

func NewAccountRepository(context data.MongoDB) interfaces.IAccountRepository {
	return &AccountRepository{context: context}
}

func (a *AccountRepository) FindOneBySpecification(specification interfaces.ISpecificationByOne) (*entity.Account, error) {
	filter, opts := specification.GetSpecification()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var account entity.Account
	if mgoErr := a.context.FindOne(ctx, accountCollection, filter, &account, opts); mgoErr != nil {
		return nil, mgoErr
	}

	return &account, nil
}

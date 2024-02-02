package repositories

import (
	"context"
	"time"

	"github.com/Patrignani/your-finances-auth/src/api/data"
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
)

const (
	userCollection = "users"
)

type UserRepository struct {
	context data.MongoDB
}

func NewUserRepository(context data.MongoDB) interfaces.IUserRepository {
	return &UserRepository{context: context}
}

func (u *UserRepository) FindOneBySpecification(specification interfaces.ISpecificationByOne) (*entity.User, error) {
	filter, opts := specification.GetSpecification()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user entity.User
	if mgoErr := u.context.FindOne(ctx, userCollection, filter, &user, opts); mgoErr != nil {
		return nil, mgoErr
	}

	return &user, nil
}

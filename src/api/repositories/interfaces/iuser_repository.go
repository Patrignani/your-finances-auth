package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IUserRepository interface {
	FindOneBySpecification(specification ISpecificationByOne) (*entity.User, error)
}

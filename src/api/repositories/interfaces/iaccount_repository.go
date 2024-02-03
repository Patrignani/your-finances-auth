package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IAccountRepository interface {
	FindOneBySpecification(specification ISpecificationByOne) (*entity.Account, error)
}

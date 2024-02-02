package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IClientRepository interface {
	Insert(client *entity.Client) error
	FindOneBySpecification(specification ISpecificationByOne) (*entity.Client, error)
}

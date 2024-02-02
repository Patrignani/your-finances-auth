package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IRefreshToken interface {
	Insert(refreshToken *entity.RefreshToken) error
	FindOneBySpecification(specification ISpecificationByOne) (*entity.RefreshToken, error)
	Update(filter map[string]interface{}, fields interface{}) error
}

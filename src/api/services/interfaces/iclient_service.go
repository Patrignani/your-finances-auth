package interfaces

import "github.com/Patrignani/your-finances-auth/src/api/entity"

type IClientService interface {
	Authenticate(clientId string, clientSecret string) (*entity.Client, error)
}

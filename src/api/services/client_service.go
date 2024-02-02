package services

import (
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	repository "github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/specifications"
	"github.com/Patrignani/your-finances-auth/src/api/services/interfaces"
)

type ClientService struct {
	repository repository.IClientRepository
}

func NewClientService(repository repository.IClientRepository) interfaces.IClientService {
	return &ClientService{repository: repository}
}

func (c *ClientService) Authenticate(clientId string, clientSecret string) (*entity.Client, error) {
	specification := specifications.NewFindClientByClientIdAndClientSecret(clientId, clientSecret, map[string]int{"_id": 1})
	return c.repository.FindOneBySpecification(specification)
}

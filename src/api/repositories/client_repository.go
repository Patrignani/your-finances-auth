package repositories

import (
	"context"
	"time"

	"github.com/Patrignani/your-finances-auth/src/api/data"
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
)

const (
	ClientCollection = "clients"
)

type ClientRepository struct {
	context data.MongoDB
}

func NewClientRepository(context data.MongoDB) interfaces.IClientRepository {
	return &ClientRepository{context: context}
}

func (c *ClientRepository) Insert(client *entity.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	id, err := c.context.Insert(ctx, ClientCollection, client)
	if err != nil {
		return err
	}

	client.ID = id

	return nil
}

func (u *ClientRepository) FindOneBySpecification(specification interfaces.ISpecificationByOne) (*entity.Client, error) {
	filter, opts := specification.GetSpecification()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var client entity.Client
	if mgoErr := u.context.FindOne(ctx, ClientCollection, filter, &client, opts); mgoErr != nil {
		return nil, mgoErr
	}

	return &client, nil
}

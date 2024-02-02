package specifications

import (
	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type FindOneByUsernameAndActive struct {
	username string
	active   bool
	project  map[string]int
}

func NewFindOneByUsernameAndActive(username string, active bool, project map[string]int) interfaces.ISpecificationByOne {
	return &FindOneByUsernameAndActive{username, active, project}
}

func (t *FindOneByUsernameAndActive) GetSpecification() (map[string]interface{}, *options.FindOneOptions) {
	opts := options.FindOne().
		SetProjection(t.project)

	filter := bson.M{"username": t.username, "active": t.active}

	return filter, opts
}

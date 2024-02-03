package specifications

import (
	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type FindByOneAccountByUserId struct {
	userId  string
	project map[string]int
}

func NewFindByOneAccountByUserId(userId string, project map[string]int) interfaces.ISpecificationByOne {
	return &FindByOneAccountByUserId{userId, project}
}

func (r *FindByOneAccountByUserId) GetSpecification() (map[string]interface{}, *options.FindOneOptions) {
	opts := options.FindOne().
		SetProjection(r.project)

	filter := bson.M{"user_id": r.userId}

	return filter, opts
}

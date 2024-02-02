package specifications

import (
	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type FindByOneUserId struct {
	userId  string
	active  bool
	project map[string]int
}

func NewFindByOneUserId(userId string, active bool, project map[string]int) interfaces.ISpecificationByOne {
	return &FindByOneUserId{userId, active, project}
}

func (u *FindByOneUserId) GetSpecification() (map[string]interface{}, *options.FindOneOptions) {
	opts := options.FindOne().
		SetProjection(u.project)

	ID, err := primitive.ObjectIDFromHex(u.userId)

	if err != nil {
		println(err.Error())
	}

	filter := bson.M{"_id": ID, "active": u.active}

	return filter, opts
}

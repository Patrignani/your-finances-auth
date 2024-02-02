package specifications

import (
	"time"

	"github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type FindByOneRefreshTokenId struct {
	refreshTokenId string
	expirationDate time.Time
	active         bool
	project        map[string]int
}

func NewFindByOneRefreshTokenId(refreshTokenId string, expirationDate time.Time, active bool, project map[string]int) interfaces.ISpecificationByOne {
	return &FindByOneRefreshTokenId{refreshTokenId, expirationDate, active, project}
}

func (r *FindByOneRefreshTokenId) GetSpecification() (map[string]interface{}, *options.FindOneOptions) {
	opts := options.FindOne().
		SetProjection(r.project)

	ID, err := primitive.ObjectIDFromHex(r.refreshTokenId)

	if err != nil {
		println(err.Error())
	}

	filter := bson.M{"_id": ID, "active": r.active, "expiration_date": bson.M{"$gte": r.expirationDate}}

	return filter, opts
}

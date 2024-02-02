package services

import (
	"sync"
	"time"

	"github.com/Patrignani/your-finances-auth/src/api/config"
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	repository "github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/specifications"
	"github.com/Patrignani/your-finances-auth/src/api/services/interfaces"
	"gopkg.in/mgo.v2/bson"
)

type RefreshTokenService struct {
	repository repository.IRefreshToken
}

func NewRefreshTokenService(repository repository.IRefreshToken) interfaces.IRefreshTokenSerice {
	return &RefreshTokenService{repository: repository}
}

func (r *RefreshTokenService) CreateRefreshToken(userID string) (*entity.RefreshToken, error) {
	var refresh *entity.RefreshToken = nil
	var erro error = nil

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		refreshToken := &entity.RefreshToken{
			UserID:         userID,
			Active:         true,
			CreateAt:       time.Now().UTC(),
			ExpirationDate: time.Now().Add(time.Duration(config.Env.RefreshTokenExpireTimeMinutes) * time.Minute),
		}

		r.DisableAllWithUserId(userID)
		erro = r.repository.Insert(refreshToken)
		refresh = refreshToken
	}()

	go func() {
		defer wg.Done()
		r.DisableAllWithExpiredDate()
	}()

	wg.Wait()

	return refresh, erro

}

func (r *RefreshTokenService) FindById(refreshTokenId string) (*entity.RefreshToken, error) {
	project := map[string]int{
		"_id":             1,
		"expiration_date": 1,
		"user_id":         1,
		"active":          1,
	}

	specification := specifications.NewFindByOneRefreshTokenId(refreshTokenId, time.Now().UTC(), true, project)

	return r.repository.FindOneBySpecification(specification)
}

func (r *RefreshTokenService) DisableAllWithExpiredDate() error {
	now := time.Now().UTC()

	filter := bson.M{
		"$and": []bson.M{
			{"expiration_date": bson.M{"$lt": now}},
			{"active": true},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"active":   false,
			"UpdateAt": now,
		},
	}

	return r.repository.Update(filter, update)
}

func (r *RefreshTokenService) DisableAllWithUserId(userID string) error {
	now := time.Now().UTC()

	filter := bson.M{
		"$and": []bson.M{
			{"user_id": userID},
			{"active": true},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"active":   false,
			"UpdateAt": now,
		},
	}

	return r.repository.Update(filter, update)
}

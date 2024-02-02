package facades

import (
	"context"
	"log"
	"time"

	"github.com/Patrignani/your-finances-auth/src/api/config"
	"github.com/Patrignani/your-finances-auth/src/api/data"
	"github.com/Patrignani/your-finances-auth/src/api/repositories"
	"github.com/Patrignani/your-finances-auth/src/api/services"
	"github.com/Patrignani/your-finances-auth/src/api/services/interfaces"
)

type AuthFacade struct {
	ClientService       interfaces.IClientService
	UserService         interfaces.IUserService
	RefreshTokenService interfaces.IRefreshTokenSerice
	AuthenticateService interfaces.IAuthenticateService
}

func CreateFacade() *AuthFacade {
	mongoContext := getMongoContext()

	//repo
	clientRepository := repositories.NewClientRepository(mongoContext)
	userRepository := repositories.NewUserRepository(mongoContext)
	refreshTokenRepository := repositories.NewRefreshTokenRepository(mongoContext)

	//services
	clientService := services.NewClientService(clientRepository)
	userService := services.NewUserService(userRepository)
	refreshTokenService := services.NewRefreshTokenService(refreshTokenRepository)
	authServices := services.NewAuthenticateService(clientService, userService, refreshTokenService)

	return &AuthFacade{clientService, userService, refreshTokenService, authServices}
}

func getMongoContext() data.MongoDB {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	mongo := data.GetInstance()

	if err := mongo.Initialize(ctx, config.Env.MongodbAddrs,
		config.Env.MongodbDatabase, config.Env.MongodbMaxPoolSize, time.Duration(config.Env.MongodbMaxConnIdleTine)*time.Minute); err != nil {
		log.Println("Could not resolve Data access layer", err)
	}

	return mongo
}

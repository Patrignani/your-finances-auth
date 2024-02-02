package services

import (
	"errors"
	"strings"

	"github.com/Patrignani/your-finances-auth/src/api/entity"
	repository "github.com/Patrignani/your-finances-auth/src/api/repositories/interfaces"
	"github.com/Patrignani/your-finances-auth/src/api/repositories/specifications"
	"github.com/Patrignani/your-finances-auth/src/api/services/interfaces"
	"github.com/labstack/echo/v4"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) interfaces.IUserService {
	return &UserService{repository: repository}
}

func (u *UserService) Authenticate(ctx echo.Context, username string, password string) (*entity.User, error) {
	msgError := "Not authorized"
	specification := specifications.NewFindOneByUsernameAndActive(username, true,
		map[string]int{"_id": 1, "username": 1, "password": 1, "seed": 1, "roles": 1, "permissions": 1, "two_factory_code": 1})

	user, err := u.repository.FindOneBySpecification(specification)

	if err != nil {
		return nil, errors.New(msgError)
	}

	if user == nil {
		return user, errors.New(msgError)
	}

	password += user.Seed
	passwordToCheck := []byte(password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), passwordToCheck)

	if err != nil {
		err = errors.New(msgError)
	}

	if len(strings.TrimSpace(user.TwoFactorCode)) > 0 {
		code := ctx.Request().Header.Get("2AF")

		if !u.verifyTwoFactorCode(user.TwoFactorCode, user.Username, code) {
			return nil, errors.New(msgError)
		}

	}

	return user, err
}

func (u *UserService) verifyTwoFactorCode(secret, accountName, userInputCode string) bool {
	key, err := totp.Generate(totp.GenerateOpts{
		Secret:      []byte(secret),
		Issuer:      "your-finances-auth",
		AccountName: accountName,
	})

	if err != nil {
		return false
	}

	return totp.Validate(userInputCode, key.Secret())
}

func (c *UserService) FindById(userId string) (*entity.User, error) {
	project := map[string]int{
		"_id":         1,
		"username":    1,
		"roles":       1,
		"permissions": 1,
		"active":      1,
	}

	specification := specifications.NewFindByOneUserId(userId, true, project)

	return c.repository.FindOneBySpecification(specification)
}

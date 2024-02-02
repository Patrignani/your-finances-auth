package services

import (
	oauth "github.com/Patrignani/simple-oauth"
	"github.com/Patrignani/your-finances-auth/src/api/entity"
	"github.com/Patrignani/your-finances-auth/src/api/services/interfaces"
	"github.com/labstack/echo/v4"
)

type AuthenticateService struct {
	clientService      interfaces.IClientService
	userService        interfaces.IUserService
	refreshTokenSerice interfaces.IRefreshTokenSerice
}

func NewAuthenticateService(client interfaces.IClientService, user interfaces.IUserService, refresh interfaces.IRefreshTokenSerice) interfaces.IAuthenticateService {
	return &AuthenticateService{clientService: client, userService: user, refreshTokenSerice: refresh}
}

func (a *AuthenticateService) ClientCredentialsAuthorization(c echo.Context, client *oauth.OAuthClient) oauth.AuthorizationRolesClient {
	authorization := oauth.AuthorizationRolesClient{}

	clientAuth, err := a.clientService.Authenticate(client.Client_id, client.Client_secret)

	if err != nil {
		authorization.Authorized = false
		return authorization
	}

	authorization.Authorized = clientAuth != nil && len(clientAuth.ID) > 0

	return authorization
}

func (a *AuthenticateService) PasswordAuthorization(c echo.Context, pass *oauth.OAuthPassword) oauth.AuthorizationRolesPassword {
	authorization := oauth.AuthorizationRolesPassword{}

	clientAuth, err := a.clientService.Authenticate(pass.Client_id, pass.Client_secret)

	if err != nil || clientAuth == nil {
		authorization.Authorized = false
		return authorization
	}

	user, err := a.userService.Authenticate(c, pass.Username, pass.Password)

	if err != nil || user == nil {
		authorization.Authorized = false
		return authorization
	}

	refresh, err := a.refreshTokenSerice.CreateRefreshToken(user.ID)

	if err != nil {
		authorization.Authorized = false
		return authorization
	}

	authorization.Authorized = true
	authorization.Roles = user.Roles
	authorization.RefreshToken = refresh.ID
	authorization.Claims = map[string]string{
		"uid": user.ID,
		"cid": clientAuth.ID,
	}

	return authorization
}

func (a *AuthenticateService) RefreshTokenCredentialsAuthorization(c echo.Context, refresh *oauth.OAuthRefreshToken) oauth.AuthorizationRolesRefresh {
	authorization := oauth.AuthorizationRolesRefresh{}

	var refreshT *entity.RefreshToken
	clientAuth, err := a.clientService.Authenticate(refresh.Client_id, refresh.Client_secret)

	if err != nil || clientAuth == nil {
		authorization.Authorized = false
		return authorization
	}

	refreshToken, err := a.refreshTokenSerice.FindById(refresh.Refresh_token)

	if refreshToken == nil || err != nil {
		authorization.Authorized = false
		return authorization
	}

	user, err := a.userService.FindById(refreshToken.UserID)

	if user == nil || err != nil {
		authorization.Authorized = false
		return authorization
	}

	refreshT, err = a.refreshTokenSerice.CreateRefreshToken(user.ID)

	if err != nil {
		authorization.Authorized = false
		return authorization
	}
	authorization.Authorized = true
	authorization.Roles = user.Roles
	authorization.RefreshToken = refreshT.ID
	authorization.Claims = map[string]string{
		"uid": user.ID,
		"cid": clientAuth.ID,
	}

	return authorization
}

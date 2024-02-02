package interfaces

import (
	oauth "github.com/Patrignani/simple-oauth"
	"github.com/labstack/echo/v4"
)

type IAuthenticateService interface {
	ClientCredentialsAuthorization(c echo.Context, client *oauth.OAuthClient) oauth.AuthorizationRolesClient
	PasswordAuthorization(c echo.Context, pass *oauth.OAuthPassword) oauth.AuthorizationRolesPassword
	RefreshTokenCredentialsAuthorization(c echo.Context, refresh *oauth.OAuthRefreshToken) oauth.AuthorizationRolesRefresh
}

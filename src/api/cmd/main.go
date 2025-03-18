package main

import (
	"net/http"
	"strings"

	oauth "github.com/Patrignani/simple-oauth"
	"github.com/Patrignani/your-finances-auth/src/api/config"
	"github.com/Patrignani/your-finances-auth/src/api/facades"
	t "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	authFacade := facades.CreateFacade()

	options := &oauth.OAuthSimpleOption{
		Key:                     config.Env.JwtKey,
		ExpireTimeMinutesClient: config.Env.JwtExpireTimeMinutesClient,
		ExpireTimeMinutes:       config.Env.JwtExpireTimeMinutes,
		AuthRouter:              "/auth",
	}

	authConfigure := &oauth.OAuthConfigure{
		ClientCredentialsAuthorization:       authFacade.AuthenticateService.ClientCredentialsAuthorization,
		PasswordAuthorization:                authFacade.AuthenticateService.PasswordAuthorization,
		RefreshTokenCredentialsAuthorization: authFacade.AuthenticateService.RefreshTokenCredentialsAuthorization,
		CustomActionRolesMiddleware:          customActionRolesMiddleware,
	}

	e := echo.New()

	e.Use(middleware.CORS())

	authRouter := oauth.NewAuthorization(authConfigure, options, e)

	authRouter.CreateAuthRouter()

	jwtValidate := authRouter.GetDefaultMiddleWareJwtValidate()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health check passed")
	})

	g := e.Group("check-jwt")
	g.Use(jwtValidate)
	g.GET("", func(c echo.Context) error {
		id := c.Get("user-id")
		cid := c.Get("cid")
		get := c.Get("user")

		println(cid, id)

		user := get.(*t.Token)
		claims := user.Claims.(t.MapClaims)
		roles := claims["roles"].([]interface{})

		rolesStr := []string{}
		permissionsStr := []string{}
		for _, role := range roles {
			rolesStr = append(rolesStr, role.(string))
		}

		ID := claims["user-id"].(string)
		return c.String(http.StatusOK, "Id:"+ID+" roles:"+strings.Join(rolesStr, ",")+" permissions:"+strings.Join(permissionsStr, ","))
	}, authRouter.RolesMiddleware("1", "5"))

	e.Logger.Fatal(e.Start(":8080"))
}

func customActionRolesMiddleware(c echo.Context, token *t.Token, claims t.MapClaims) error {
	if claims["uid"] != nil {
		c.Set("user-id", claims["user-id"].(string))
	}

	if claims["cid"] != nil {
		c.Set("cid", claims["cid"].(string))
	}

	return nil
}

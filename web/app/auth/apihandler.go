package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	iValidator "github.com/miiy/go-blog/pkg/gin/validator"
	"net/http"
)

type apiHandler struct {
}

func NewApiHandler() *apiHandler {
	return &apiHandler{}
}

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
)


//signUpParam := SignUpParam{
//	Username: ctx.PostForm("username"),
//	Password: ctx.PostForm("password"),
//	PasswordConfirmation: ctx.PostForm("password_confirmation"),
//	Email: ctx.PostForm("email"),
//}
//if err := module.validator.validateSignIn(&signInParam); err != nil {
//ctx.JSON(http.StatusBadRequest, BadRequestResponse(err))
//return
//}


// authorized
//authorized := router.Group("/")
//authorized.Use(middleware.JWTAuth())
//g.POST("/token/refresh", m.webHandler.refreshToken)
//
//


func (h *apiHandler) authUser(c *gin.Context) (*AuthUser, error) {
	cUser, exists := c.Get("auth")
	if !exists {
		return nil, ErrUnauthorized
	}
	if authUser, ok := cUser.(*AuthUser); ok {
		return authUser, nil
	}
	return nil, ErrInternalServerError
}

func (h *apiHandler) signUp(c *gin.Context) {
	var sup SignUpParam
	if err := c.ShouldBind(&sup); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse(iValidator.VErrorsTranslations(err)))
		return
	}

	user, err := module.service.signUp(&sup)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(user))
}

func (h *apiHandler) signIn(c *gin.Context) {
	var sip SignInParam
	if err := c.ShouldBind(&sip); err != nil {
		errs := iValidator.VErrorsTranslations(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(errs))
		return
	}

	user, err := module.service.signIn(&sip)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	token, err := module.jwtAuth.CreateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	apisUserSigIn := APIUserSigIn{
		AccessToken: token,
		TokenType: "Bearer",
		ExpiresIn: module.jwtAuth.Options.ExpiresIn,
		User: &APIUser{
			Username: user.Username,
		},
	}
	c.JSON(http.StatusOK, SuccessResponse(apisUserSigIn))
}

func (h *apiHandler) refreshToken(c *gin.Context) {
	var rtp RefreshTokenParam
	if err := c.ShouldBind(&rtp); err != nil {
		errs := iValidator.VErrorsTranslations(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(errs))
		return
	}

	token, err := module.jwtAuth.RefreshToken(rtp.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	apisUserSigIn := APIUserSigIn{
		AccessToken: token,
		TokenType: "Bearer",
		ExpiresIn: module.jwtAuth.Options.ExpiresIn,
	}
	c.JSON(http.StatusOK, SuccessResponse(apisUserSigIn))
}

func (h *apiHandler) logout(c *gin.Context) {
	fmt.Println(c.Get("auth"))
}

func (h *apiHandler) profile(c *gin.Context) {
	authUser, err := h.authUser(c)
	if err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	if user := module.repository.firstByUsername(authUser.Username); user != nil {
		c.JSON(http.StatusOK, SuccessResponse(APIUser{
			Username: user.Username,
		}))
		return
	}

	c.AbortWithStatus(http.StatusBadRequest)
}

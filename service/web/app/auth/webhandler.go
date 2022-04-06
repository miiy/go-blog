package auth

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	//"goblog.com/pkg/session"
	//iValidator "goblog.com/pkg/validator"
	"log"
	"net/http"
)

type webHandler struct {
	session *session.Session
}

func init()  {
	gob.Register(iValidator.ValidationErrorsTranslations{})
}

func NewHandler(session *session.Session) *webHandler {
	return &webHandler{
		session: session,
	}
}

func (h *webHandler) signUpForm(c *gin.Context) {
	rst := gin.H{
		"PageTitle": "Signup",
	}
	if flashes, exists := c.Get("flashes"); exists {
		rst["flashes"] = flashes
	}
	c.HTML(http.StatusOK, "auth/signup.html", rst)
}

func (h *webHandler) signUp(c *gin.Context) {
	var sup SignUpParam

	if err := c.ShouldBind(&sup); err != nil {
		errs := iValidator.VErrorsTranslations(err)
		if err = addFlash(c, errs); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.Redirect(http.StatusFound, "/signup")
		return
	}

	if _, err := module.service.signUp(&sup); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Redirect(http.StatusFound, "/signin")
}

func (h *webHandler) signInForm(c *gin.Context) {
	rst := gin.H{
		"PageTitle": "Signin",
	}

	if flashes, exists := c.Get("flashes"); exists {
		rst["flashes"] = flashes
	}
	c.HTML(http.StatusOK, "auth/signin.html", rst)
}

func (h *webHandler) signIn(c *gin.Context) {
	var sip SignInParam
	if err := c.ShouldBind(&sip); err != nil {
		errs := iValidator.VErrorsTranslations(err)

		if err = addFlash(c, errs); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Redirect(http.StatusFound, "/signin")
		return
	}

	user, err := module.service.signIn(&sip)
	if err != nil {
		errs := iValidator.ValidationErrorsTranslations{}
		errs["username"] = err.Error()
		if err = addFlash(c, errs); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Redirect(http.StatusFound, "/signin")
		return
	}

	sUser := AuthUser{
		Username: user.Username,
	}
	if err = addValue(c, "auth", sUser); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/auth/profile")
}

func (h *webHandler) logout(c *gin.Context) {
	if err := removeValue(c, "auth"); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (h *webHandler) profile(c *gin.Context) {
	sUser, exists := c.Get("auth")

	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, ok := sUser.(*AuthUser)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.HTML(http.StatusOK, "auth/profile.html", gin.H{
		"PageTitle": "Profile",
		"Username": user.Username,
	})
}

func removeValue(c *gin.Context, key string) error  {
	sess, err := module.session.Store.Get(c.Request, module.session.Options.SessionCookie)
	if err != nil {
		return err
	}
	delete(sess.Values, key)
	if err = sess.Save(c.Request, c.Writer); err != nil {
		return err
	}
	return nil
}

func addValue(c *gin.Context, key string, value interface{}) error  {
	sess, err := module.session.Store.Get(c.Request, module.session.Options.SessionCookie)
	if err != nil {
		return err
	}
	sess.Values[key] = value
	if err = sess.Save(c.Request, c.Writer); err != nil {
		return err
	}
	return nil
}

func addFlash(c *gin.Context, value interface{}) error  {
	sess, err := module.session.Store.Get(c.Request, module.session.Options.SessionCookie)
	if err != nil {
		return err
	}
	sess.AddFlash(value)
	if err = sess.Save(c.Request, c.Writer); err != nil {
		return err
	}
	return nil
}

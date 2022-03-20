package auth

import (
	"github.com/gin-gonic/gin"
	"goblog.com/pkg/jwtauth"
	//"goblog.com/pkg/session"
	//pkgValidator "goblog.com/pkg/validator"
	"log"
	"net/http"
)

func JWTAuthenticationMiddleware(jwtAuth *jwtauth.JWTAuth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := jwtauth.HeaderToken(ctx.Request)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := jwtAuth.ParseToken(token)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user := &AuthUser{
			Username: claims.Username,
		}
		ctx.Set("auth", user)
		ctx.Next()
	}
}

func SessionAuthenticationMiddleware(session *session.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := session.Store.Get(c.Request, session.Options.SessionCookie)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		user := sess.Values["auth"]
		if user == nil {
			c.Redirect(http.StatusMovedPermanently, "/signin")
			return
		}
		c.Set("auth", user)
		c.Next()

	}
}

func SessionFlashMiddleware(session *session.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := session.Store.Get(c.Request, session.Options.SessionCookie)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		flashes := sess.Flashes()
		if len(flashes) > 0 {
			for _, v := range flashes {
				if e, ok := v.(pkgValidator.ValidationErrorsTranslations); ok {
					c.Set("flashes", e)
				}
			}

			if err = sess.Save(c.Request, c.Writer); err != nil {
				log.Println(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		c.Next()

	}
}

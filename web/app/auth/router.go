package auth

import "github.com/gin-gonic/gin"

func (m *Module) RegisterRouter() {
	r := m.router
	session := m.session
	jwtAuth := m.jwtAuth

	r.GET("/signup", SessionFlashMiddleware(session), m.webHandler.signUpForm)
	r.POST("/signup", m.webHandler.signUp)
	r.GET("/signin", SessionFlashMiddleware(session), m.webHandler.signInForm)
	r.POST("/signin", m.webHandler.signIn)

	// auth
	g := m.router.Group("/auth").Use(SessionAuthenticationMiddleware(session))
	{
		g.GET("/logout", m.webHandler.logout)
		g.GET("/profile", m.webHandler.profile)
	}

	// api
	aG := r.Group("/api/v1/auth")
	{
		aG.POST("/signup", m.apiHandler.signUp)
		aG.POST("/signin", m.apiHandler.signIn)
		aG.POST("/token/refresh", m.apiHandler.refreshToken)

		authorized := aG.Use(JWTAuthenticationMiddleware(jwtAuth))
		{
			authorized.GET("/logout", m.apiHandler.logout)
			authorized.GET("/profile",  m.apiHandler.profile)
		}
	}
}

func registerWebRouter(r *gin.Engine)  {

}

func registerApiRouter(r *gin.Engine)  {


}

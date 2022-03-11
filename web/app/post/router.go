package post

func (m *Module) Router() {
	g := m.router.Group("/health")
	{
		g.GET("/", indexHandler)
		g.GET("/:id", showHandler)
		g.GET("/:id/create", createHandler)
		g.POST("/:id", storeHandler)
		g.GET("/:id/edit", editHandler)
		g.PUT("/:id", updateHandler)
		g.PATCH("/:id", updateHandler)
		g.DELETE("/:id", destroyHandler)
	}
}

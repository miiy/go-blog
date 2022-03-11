package about

func (m *Module) RegisterRouter() {
	g := m.router.Group("/about")
	{
		g.GET("/", indexHandler)
	}
}

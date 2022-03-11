package home

func (m *Module) RegisterRouter() {
	r := m.router
	r.GET("/", indexHandler)
}

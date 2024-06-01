package http

func (s *server) SetupRouter() {
	s.webSetUp()
	s.authRouter()
}

func (s *server) authRouter() {
	route := s.app.Group("/auth")
	route.Post("/login", s.Login, s.loginValidate)
	route.Get("/get-me", s.GetMe, s.IsAuthorized)
	route.Get("/get-user/:id", s.GetUser, s.IsAuthorized)
}

func (s *server) webSetUp() {
	route := s.app.Group("/web")
	route.Get("/", s.HomePageHandler)
}

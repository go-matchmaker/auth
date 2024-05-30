package http

func (s *server) SetupRouter() {
	s.webSetUp()
	s.authRouter()
}

func (s *server) authRouter() {
	route := s.app.Group("/auth")
	route.Post("/register", s.RegisterUser, s.registerValidate)
	route.Post("/login", s.Login, s.loginValidate)
	route.Get("/self-info", s.SelfInfo, s.IsAuthorized)
}

func (s *server) webSetUp() {
	route := s.app.Group("/web")
	route.Get("/", s.HomePageHandler)
}

package http

import (
	"auth/internal/converter"
	"auth/internal/core/domain/valueobject"
	"auth/internal/core/util"
	"auth/internal/dto"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func (s *server) RegisterUser(c fiber.Ctx) error {
	reqBody := new(dto.UserRegisterRequest)
	body := c.Body()
	if err := json.Unmarshal(body, reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	hashedPassword, err := util.HashPassword(reqBody.Password)
	if err != nil {
		return s.errorResponse(c, "error while trying to hash password", err, nil, fiber.StatusBadRequest)
	}

	userModel, err := converter.UserRegisterRequestToModel(reqBody, "user", hashedPassword)
	if err != nil {
		return s.errorResponse(c, "error while trying to convert user register to model", err, nil, fiber.StatusBadRequest)
	}
	userID, err := s.userService.Register(s.ctx, userModel)
	if err != nil {
		return s.errorResponse(c, "error while trying to register user", err, nil, fiber.StatusBadRequest)
	}

	zap.S().Info("User Registered Successfully! User:", userID)
	return s.successResponse(c, nil, "user registered successfully", fiber.StatusOK)
}

func (s *server) Login(c fiber.Ctx) error {
	reqBody := new(dto.UserLoginRequest)
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	ip := c.IP()
	userData, err := s.userService.Login(c.Context(), reqBody.Email, reqBody.Password, ip)
	if err != nil {
		return s.errorResponse(c, "error while trying to login", err, nil, fiber.StatusBadRequest)
	}

	userResponse := dto.NewUserLoginRequestResponse(userData)
	return s.successResponse(c, userResponse, "user logged in successfully", fiber.StatusOK)
}

func (s *server) SelfInfo(c fiber.Ctx) error {
	data, ok := c.Locals(AuthPayload).(*valueobject.Payload)
	if !ok {
		return s.errorResponse(c, "session not found", nil, nil, fiber.StatusBadRequest)
	}

	user, err := s.userService.UserSelfInfo(c.Context(), data.Email)
	if err != nil {
		return s.errorResponse(c, "error while trying to get user info", err, nil, fiber.StatusBadRequest)
	}

	return s.successResponse(c, user, "user info retrieved successfully", fiber.StatusOK)
}

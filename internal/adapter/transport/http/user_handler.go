package http

import (
	"auth/internal/dto"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
)

func (s *server) Login(c fiber.Ctx) error {
	reqBody := new(dto.UserLoginRequest)
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	userData, err := s.userService.Login(c.Context(), reqBody.Email, reqBody.Password)
	if err != nil {
		return s.errorResponse(c, "error while trying to login", err, nil, fiber.StatusBadRequest)
	}

	userResponse := dto.NewUserLoginRequestResponse(userData)
	return s.successResponse(c, userResponse, "user logged in successfully", fiber.StatusOK)
}

func (s *server) GetMe(c fiber.Ctx) error {
	return nil
	// userID := c.Locals("user").(*valueobject.UserClaims).UserID
	// userData, err := s.userService.GetMe(c.Context(), userID)
	// if err != nil {
	// 	return s.errorResponse(c, "error while trying to get user data", err, nil, fiber.StatusBadRequest)
	// }

	// userResponse := dto.NewUserResponse(userData)
	// return s.successResponse(c, userResponse, "user data fetched successfully", fiber.StatusOK)
}

func (s *server) GetUser(c fiber.Ctx) error {
	userID := c.Params("id")
	userData, err := s.userService.GetUser(c.Context(), userID)
	if err != nil {
		return s.errorResponse(c, "error while trying to get user data", err, nil, fiber.StatusBadRequest)
	}

	userResponse := dto.NewUserResponse(userData)
	return s.successResponse(c, userResponse, "user data fetched successfully", fiber.StatusOK)
}

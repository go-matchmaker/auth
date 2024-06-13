package http

import (
	"auth/internal/converter"
	"auth/internal/core/domain/valueobject"
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

	userResponse := converter.UserLoginModelToDto(userData)
	return s.successResponse(c, userResponse, "user logged in successfully", fiber.StatusOK)
}

func (s *server) GetMe(c fiber.Ctx) error {
	payload := c.Locals(AuthPayload).(*valueobject.Payload)

	userAggregate, err := s.userService.GetUserByID(c.Context(), payload.ID)
	if err != nil {
		return s.errorResponse(c, "error while trying to fetch user data", err, nil, fiber.StatusBadRequest)
	}

	userResponse := converter.GetUserModelToDto(userAggregate)
	return s.successResponse(c, userResponse, "user data fetched successfully", fiber.StatusOK)
}

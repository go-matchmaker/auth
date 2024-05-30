package http

import (
	"auth/internal/core/domain/valueobject"
	"errors"
	"github.com/gofiber/fiber/v3"
	"strings"
)

const (
	AuthHeader  = "Authorization"
	AuthType    = "Bearer"
	AuthPublic  = "AuthPublic"
	AuthPayload = "Payload"
)

func (s *server) IsAuthorized(c fiber.Ctx) error {
	token := c.Get(AuthHeader)
	if token == "" {
		return s.errorResponse(c, "authorization header is not provided", errors.New("authorization header is not provided"), nil, fiber.StatusUnauthorized)
	}

	fields := strings.Fields(token)
	if len(fields) != 2 {
		return s.errorResponse(c, "invalid authorization header format", errors.New("invalid authorization header format"), nil, fiber.StatusUnauthorized)
	}

	if fields[0] != AuthType {
		return s.errorResponse(c, "unsupported authorization type", errors.New("unsupported authorization type"), nil, fiber.StatusUnauthorized)
	}

	accessToken := fields[1]
	accessPublic := c.Get(AuthPublic)
	if accessPublic == "" {
		return s.errorResponse(c, "public key is not provided", errors.New("public key is not provided"), nil, fiber.StatusUnauthorized)
	}
	payload, err := s.tokenService.DecodeToken(accessToken, accessPublic)
	if err != nil {
		return s.errorResponse(c, "invalid access token", err, nil, fiber.StatusUnauthorized)
	}

	c.Locals(AuthPayload, payload)
	return c.Next()
}

func (s *server) IsAdmin(c fiber.Ctx) error {
	role, ok := c.Locals(AuthPayload).(*valueobject.Payload)
	if !ok {
		return s.errorResponse(c, "payload not found in context", errors.New("payload not found in context"), nil, fiber.StatusUnauthorized)
	}

	if role.Role != "admin" {
		return s.errorResponse(c, "unauthorized", errors.New("unauthorized"), nil, fiber.StatusUnauthorized)
	}

	return c.Next()
}

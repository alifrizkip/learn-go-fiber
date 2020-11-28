package auth

import (
	"learn-go-fiber/helper"
	"learn-go-fiber/modules/user"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service IService
}

func newHandler(service IService) *handler {
	return &handler{service}
}

func (h *handler) RegisterUser(c *fiber.Ctx) error {
	registerInput := new(registerUserInput)
	if err := c.BodyParser(registerInput); err != nil {
		if err != nil {
			return helper.SendAPIResponse(c)("Failed parse request body", fiber.StatusBadRequest, false, nil)
		}
	}

	errors := helper.ValidateRequest(*registerInput)
	if errors != nil {
		errorsData := helper.ErrorsData(errors)
		return helper.SendAPIResponse(c)("Failed to validate register data", fiber.StatusUnprocessableEntity, false, errorsData)
	}

	newUser, err := h.service.RegisterUser(registerInput)
	if err != nil {
		return helper.SendAPIResponse(c)("Failed to register new user", fiber.StatusBadRequest, false, nil)
	}

	registerResponse := FormatRegisterResponse(newUser)
	return helper.SendAPIResponse(c)("Success register new user", fiber.StatusOK, true, registerResponse)
}

func (h *handler) Login(c *fiber.Ctx) error {
	loginInput := new(loginInput)
	if err := c.BodyParser(loginInput); err != nil {
		if err != nil {
			return helper.SendAPIResponse(c)("Failed parse request body", fiber.StatusBadRequest, false, nil)
		}
	}

	errors := helper.ValidateRequest(*loginInput)
	if errors != nil {
		return helper.SendAPIResponse(c)("Failed to validate login data", fiber.StatusUnprocessableEntity, false, errors)
	}

	loggedinUser, err := h.service.Login(loginInput)
	if err != nil {
		errorsData := helper.ErrorsData(err.Error())
		return helper.SendAPIResponse(c)("Login failed", fiber.StatusUnprocessableEntity, false, errorsData)
	}

	token, err := h.service.GenerateToken(loggedinUser.ID)
	if errors != nil {
		return helper.SendAPIResponse(c)("Failed to validate login data", fiber.StatusUnprocessableEntity, false, errors)
	}

	loginResponse := FormatLoginResponse(loggedinUser, token)
	return helper.SendAPIResponse(c)("Success login", fiber.StatusOK, true, loginResponse)
}

func (h *handler) Profile(c *fiber.Ctx) error {
	currentUser := c.Locals("current_user").(*user.User)

	profileResponse := FormatRegisterResponse(currentUser)
	return helper.SendAPIResponse(c)("Profile data", fiber.StatusOK, true, profileResponse)
}

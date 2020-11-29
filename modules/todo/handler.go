package todo

import (
	"learn-go-fiber/helper"
	"learn-go-fiber/modules/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service IService
}

func newHandler(service IService) *handler {
	return &handler{service}
}

func (h *handler) GetAllTodos(c *fiber.Ctx) error {
	currentUser := c.Locals("current_user").(*user.User)

	todos, err := h.service.GetTodosByUserID(currentUser.ID)
	if err != nil {
		return helper.SendAPIResponse(c)("Failed to get todos data", fiber.StatusBadRequest, false, nil)
	}

	return helper.SendAPIResponse(c)("Success get todos data", fiber.StatusOK, true, todos)
}

func (h *handler) GetTodoDetail(c *fiber.Ctx) error {
	currentUser := c.Locals("current_user").(*user.User)
	id, _ := strconv.Atoi(c.Params("id"))
	todo, err := h.service.GetATodo(&Todo{ID: id, UserID: currentUser.ID})
	if err != nil {
		if err.Error() == "record not found" {
			return helper.SendAPIResponse(c)("Todo not found", fiber.StatusNotFound, false, nil)
		}

		return helper.SendAPIResponse(c)("Failed to get todo data", fiber.StatusBadRequest, false, nil)
	}

	return helper.SendAPIResponse(c)("Success get todo data", fiber.StatusOK, true, todo)
}

func (h *handler) CreateNewTodo(c *fiber.Ctx) error {
	currentUser := c.Locals("current_user").(*user.User)

	todoInput := new(createTodoInput)
	if err := c.BodyParser(todoInput); err != nil {
		if err != nil {
			return helper.SendAPIResponse(c)("Failed parse request body", fiber.StatusBadRequest, false, nil)
		}
	}

	errors := helper.ValidateRequest(*todoInput)
	if errors != nil {
		return helper.SendAPIResponse(c)("Failed to validate todo data", fiber.StatusUnprocessableEntity, false, errors)
	}

	todo, err := h.service.CreateTodo(todoInput, currentUser.ID)
	if err != nil {
		return helper.SendAPIResponse(c)("Failed to create new todo", fiber.StatusBadRequest, false, nil)
	}

	return helper.SendAPIResponse(c)("Success create todo", fiber.StatusOK, true, todo)
}

func (h *handler) CompleteTodo(c *fiber.Ctx) error {
	currentUser := c.Locals("current_user").(*user.User)
	id, _ := strconv.Atoi(c.Params("id"))
	todo, err := h.service.GetATodo(&Todo{ID: id, UserID: currentUser.ID})
	if err != nil {
		if err.Error() == "record not found" {
			return helper.SendAPIResponse(c)("Todo not found", fiber.StatusNotFound, false, nil)
		}

		return helper.SendAPIResponse(c)("Failed to get todo data", fiber.StatusBadRequest, false, nil)
	}

	todo, err = h.service.CompleteTodo(id)
	if err != nil {
		return helper.SendAPIResponse(c)("Failed to complete todo", fiber.StatusBadRequest, false, nil)
	}

	return helper.SendAPIResponse(c)("Success update todo to complete", fiber.StatusOK, true, todo)
}

func (h *handler) DeleteTodo(c *fiber.Ctx) error {
	currentUser := c.Locals("current_user").(*user.User)
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := h.service.GetATodo(&Todo{ID: id, UserID: currentUser.ID})
	if err != nil {
		if err.Error() == "record not found" {
			return helper.SendAPIResponse(c)("Todo not found", fiber.StatusNotFound, false, nil)
		}

		return helper.SendAPIResponse(c)("Failed to get todo data", fiber.StatusBadRequest, false, nil)
	}

	err = h.service.DeleteTodo(id)
	if err != nil {
		return helper.SendAPIResponse(c)("Failed to delete todo", fiber.StatusBadRequest, false, nil)
	}

	return helper.SendAPIResponse(c)("Success delete todo", fiber.StatusOK, true, nil)
}

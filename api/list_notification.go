package api

import "github.com/gofiber/fiber/v2"

func (srv *Service) ListNotifications(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	notifications, err := srv.repository.ListNotificationsByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Ocorreu um erro ao listar as notificações. Por favor, tente novamente mais tarde.",
		})
	}

	return c.JSON(notifications)
}

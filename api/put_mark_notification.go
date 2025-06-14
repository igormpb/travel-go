package api

import "github.com/gofiber/fiber/v2"

func (srv *Service) MarkNotificationRead(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := srv.repository.MarkNotificationAsRead(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to mark as read"})
	}

	return c.JSON(fiber.Map{"message": "Notification marked as read"})
}

package api

import "github.com/gofiber/fiber/v2"

func (srv *Service) ListTravels(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	status := c.Query("status")
	destination := c.Query("destination")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	travels, err := srv.repository.ListTravel(userID, status, destination, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ocorreu um problema ao buscar suas viagens. Por favor, tente novamente em instantes.",
		})
	}

	return c.JSON(travels)
}

package api

import "github.com/gofiber/fiber/v2"

func (srv *Service) GetTravelByID(c *fiber.Ctx) error {
	travelID := c.Params("id")
	userID := c.Locals("userId").(string)

	travel, err := srv.repository.GetTravelById(travelID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Não foi possível buscar as informações da viagem no momento. Tente novamente em instantes.",
		})
	}

	if travel == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Não encontramos a viagem solicitada. Verifique se o código está correto.",
		})
	}

	return c.JSON(travel)
}

package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/travel-go/models"
)

type CreateTravelRequest struct {
	RequesterName string `json:"requesterName"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departureDate"`
	ReturnDate    string `json:"returnDate"`
}

func (srv *Service) CreateTravel(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var body CreateTravelRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato de dados inválido. Verifique os campos e tente novamente.",
		})
	}

	if body.RequesterName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O nome do solicitante é obrigatório.",
		})
	}

	if body.Destination == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O destino da viagem é obrigatório.",
		})
	}

	// Converte as datas string para time.Time
	layout := "2006-01-02" // Ex: "2025-06-14"
	departureDate, err := time.Parse(layout, body.DepartureDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data de partida inválida. Use o formato YYYY-MM-DD.",
		})
	}

	returnDate, err := time.Parse(layout, body.ReturnDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data de retorno inválida. Use o formato YYYY-MM-DD.",
		})
	}

	if returnDate.Before(departureDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "A data de retorno não pode ser anterior à data de partida.",
		})
	}

	travel := &models.Travel{
		ID:            uuid.New(),
		RequesterName: body.RequesterName,
		Destination:   body.Destination,
		DepartureDate: departureDate,
		ReturnDate:    returnDate,
		Status:        models.StatusRequested,
		UserID:        uuid.MustParse(userID),
		CreatedAt:     time.Now(),
	}

	created, err := srv.repository.CreateTravel(travel)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro interno ao registrar a viagem. Tente novamente mais tarde.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(created)
}

package api

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/travel-go/models"
)

type UpdateStatusRequest struct {
	Status models.TravelStatus `json:"status"`
}

func (srv *Service) UpdateStatus(c *fiber.Ctx) error {
	travelID := c.Params("id")
	userID := c.Locals("userId").(string)

	travel, err := srv.repository.GetTravelById(travelID, userID)
	if err != nil || travel == nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pedido de viagem não encontrado."})
	}

	// Regra: dono do pedido não pode alterar status
	if travel.UserID.String() == userID {
		return c.Status(403).JSON(fiber.Map{"error": "Você não pode alterar o status do seu próprio pedido."})
	}

	var body UpdateStatusRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida de status. Verifique os dados e tente novamente."})
	}

	normalizedStatus := strings.ToUpper(string(body.Status))
	if !models.IsValidStatus(normalizedStatus) {
		return c.Status(400).JSON(fiber.Map{"error": "Status fornecido é inválido."})
	}
	body.Status = models.TravelStatus(normalizedStatus)

	// Regra: cancelar aprovado só se faltarem mais de 3 dias
	if travel.Status == models.StatusApproved && body.Status == models.StatusCanceled {
		if time.Until(travel.DepartureDate).Hours() < 72 {
			return c.Status(403).JSON(fiber.Map{"error": "Não é possível cancelar uma viagem aprovada com menos de 3 dias de antecedência."})
		}
	}

	// Atualizar status
	travel.Status = body.Status
	if err := srv.repository.UpdateTravel(travel); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao atualizar o status. Tente novamente em instantes."})
	}

	// Notificação simulada
	notif := &models.Notification{
		ID:      uuid.New(),
		UserID:  travel.UserID,
		Message: fmt.Sprintf("Seu pedido de viagem %s foi %s.", travel.ID, travel.Status),
	}
	_ = srv.repository.CreateNotification(notif)

	return c.JSON(travel)
}

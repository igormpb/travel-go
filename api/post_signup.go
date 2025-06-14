package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/igormpb/travel-go/models"
	"golang.org/x/crypto/bcrypt"
)

type PostSignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (srv *Service) PostSignup(c *fiber.Ctx) error {

	var body PostSignupRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Não foi possível entender os dados enviados. Verifique o formulário e tente novamente.",
		})
	}

	if body.Email == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Preencha todos os campos obrigatórios: e-mail, senha e nome.",
		})
	}

	existingUser, err := srv.repository.GetUserByEmail(body.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ocorreu um erro ao verificar o e-mail. Por favor, tente novamente em instantes.",
		})
	}
	if existingUser != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Este e-mail já está cadastrado. Tente fazer login ou use outro e-mail.",
		})
	}

	salt := uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ocorreu um problema ao proteger sua senha. Tente novamente mais tarde.",
		})
	}

	user := &models.User{
		ID:    uuid.New(),
		Name:  body.Name,
		Email: body.Email,
	}
	createdUser, err := srv.repository.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao salvar suas informações. Por favor, tente novamente.",
		})
	}

	creds := &models.Credentials{
		ID:           uuid.New(),
		UserID:       createdUser.ID,
		Salt:         salt,
		PasswordHash: string(hashedPassword),
	}
	if _, err := srv.repository.CreateCredentials(creds); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao finalizar seu cadastro. Tente novamente em instantes.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"mensagem": "Cadastro realizado com sucesso! Você já pode fazer login na plataforma.",
	})
}

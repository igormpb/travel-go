package api

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type PostSigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (srv *Service) PostSignin(c *fiber.Ctx) error {
	var body PostSigninRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Requisição inválida. Verifique os dados enviados e tente novamente.",
		})
	}

	user, err := srv.repository.GetUserByEmail(body.Email)
	if err != nil || user == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "E-mail ou senha inválidos.",
		})
	}

	creds, err := srv.repository.GetByUserId(user.ID.String())
	if err != nil || creds == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "E-mail ou senha inválidos.",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(body.Password+creds.Salt)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "E-mail ou senha inválidos.",
		})
	}

	// Gerar JWT
	claims := jwt.MapClaims{
		"userId": user.ID.String(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Erro ao gerar token de acesso. Por favor, tente novamente.",
		})
	}

	return c.JSON(fiber.Map{"token": tokenStr, "user": map[string]string{
		"name":  user.Name,
		"email": user.Email,
	}})
}

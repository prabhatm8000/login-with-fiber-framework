package authControllers

import (
	"os"
	"time"

	"example.com/login/models/authModel"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// matches the email and password in the database, generates a jwt token and sets it in a cookie
func Login(c *fiber.Ctx) error {
	var creds authModel.Credentials

	if err := c.BodyParser(&creds); err != nil {
		return err
	}

	user := authModel.User{Email: creds.Email}
	if err := user.FindByEmail(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	isMatch, err := user.MatchPassword(creds.Password)
	if err != nil || !isMatch {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	token, err := createJWTToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	cookie := fiber.Cookie{
		Name:     "auth-token",
		Value:    token,
		Expires:  time.Now().AddDate(0, 0, 3),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "Logged in successfully!", "name": user.Name, "email": user.Email})
}

// create a jwt token with the user id
func createJWTToken(user authModel.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().AddDate(0, 0, 3).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to load JWT secret")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Register(c *fiber.Ctx) error {
	var user authModel.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// if the user creation fails because the email already exists, return a 400 status code
	// if the user creation fails for any other reason, return a 500 status code
	_, err := user.Create()
	if err != nil {
		if err.Error() == "email already exists" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.JSON(fiber.Map{"message": "Successfully registered!"})
}

// remove cookie with auth-token
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "auth-token",
		Value:    "",
		Expires:  time.Now(),
		MaxAge:   -1,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "Successfully logged out!"})
}

func GetUser(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to find user"})
	}

	var user authModel.User

	if err := user.FindByID(userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to find user"})
	}

	return c.JSON(fiber.Map{"name": user.Name, "email": user.Email})
}

package authRoutes

import (
	"github.com/gofiber/fiber/v2"

	"example.com/login/controllers/authControllers"
	"example.com/login/middlewares/authMiddleware"
)

func AddLoginRoute(grp fiber.Router) {
	loginController := authControllers.Login
	registerController := authControllers.Register
	logoutController := authControllers.Logout
	getUserController := authControllers.GetUser

	grp.Post("/login", loginController)
	grp.Post("/register", registerController)
	grp.Get("/logout", authMiddleware.Auth, logoutController)
	grp.Get("/user", authMiddleware.Auth, getUserController)
}

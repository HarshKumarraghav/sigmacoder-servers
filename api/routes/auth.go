package routes

// The `import` statement is used to import packages that are required for the code to run. In this
// case, the code is importing the `http` package, the `auth` package from the `sigmacoder/pkg` directory,
// and the `fiber` package from the `github.com/gofiber/fiber/v2` repository. These packages are used
// in the code to handle HTTP requests and responses, and to interact with the authentication service.
import (
	"net/http"
	"os"
	"sigmacoder/pkg/auth"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// The function handles sign up requests by parsing the request body, calling the sign up service, and
// returning a JSON response with a refresh token.
func SignUpHandler(repo *auth.Repo, svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var in auth.InUser
		if err := c.BodyParser(&in); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "status": "failed1"})
		}
		refreshToken, err := svc.SignUp(in)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "status": "failed2"})
		}
		return c.Status(200).JSON(fiber.Map{"token": refreshToken, "status": "success"})
	}
}
func LoginHandler(repo *auth.Repo, svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var in auth.AuthBody
		if err := c.BodyParser(&in); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "status": "failed1"})
		}
		refreshToken, err := svc.Login(in.Email, in.Password)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "status": "failed2"})
		}
		return c.Status(200).JSON(fiber.Map{"token": refreshToken, "status": "success"})
	}
}

// The function handles sign up requests by parsing the request body, calling the sign up service, and
// returning a JSON response with a refresh token.
func CreateAuthRoutes(app *fiber.App, userRepo *auth.Repo, svc auth.Service) {
	app.Post("/api/auth/register", SignUpHandler(userRepo, svc))
	app.Post("/api/auth/login", LoginHandler(userRepo, svc))
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
}

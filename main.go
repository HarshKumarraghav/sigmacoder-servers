package main

// The `import` block is importing various packages and modules that are required for the application
// to run. These include:
import (
	"context"
	"log"
	"os"
	"sigmacoder/api/routes"
	"sigmacoder/pkg/allquestions"
	"sigmacoder/pkg/auth"
	"sigmacoder/pkg/configuration"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// `app := fiber.New()` is creating a new instance of the Fiber web framework, which will be used to
	// define and handle HTTP routes for the application.
	app := fiber.New()
	// `def` is a variable that holds a CORS (Cross-Origin Resource Sharing) configuration. It specifies
	// the allowed origins, methods, headers, and credentials for cross-origin requests. In this case, it
	// allows any origin, all HTTP methods, specific headers, and credentials to be included in the
	// request. This configuration is used by the `cors.New()` middleware to enable CORS for all routes in
	// the Fiber application.
	def := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Request-With",
		AllowCredentials: true,
	}
	// `app.Use(cors.New(def))` is adding a CORS middleware to the Fiber application, which allows
	// cross-origin requests from any origin. `godotenv.Load()` is loading environment variables from a
	// `.env` file into the application's environment.
	app.Use(cors.New(def))
	godotenv.Load()
	// `config := configuration.FromEnv()` is loading the application configuration from environment
	// variables using the `FromEnv()` method of the `configuration` package. This allows the application
	// to read configuration values such as the MongoDB URI and the application port from environment
	// variables, which can be set differently depending on the deployment environment.
	config := configuration.FromEnv()
	// This code is establishing a connection to a MongoDB database using the MongoDB Go driver. It creates
	// a new client instance using the `mongo.Connect()` method, passing in a context and options for the
	// client. The `config.MongoURI` value is used to specify the URI for the MongoDB database. If an error
	// occurs during the connection process, the program will log the error and exit using `log.Panic()`.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Panic(err)
	}
	// `db := client.Database("sigmacoder")` is creating a new database instance named "sigmacoder" using the
	// MongoDB client connection. This allows the application to interact with the "sigmacoder" database using
	// the methods provided by the MongoDB Go driver.
	db := client.Database("sigmacoder")

	// This code is creating a route for the root URL ("/") of the application using the HTTP GET method.
	// When a user makes a GET request to the root URL, the function passed as the second argument to
	// `app.Get()` is executed. This function returns a JSON response with a "ping" key and "pong" value,
	// indicating that the server is up and running.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"ping": "pong",
		})
	})
	// `userRepo := auth.NewRepo(db)` is creating a new instance of the `auth.Repo` struct, which is used
	// to interact with the MongoDB database and perform CRUD (Create, Read, Update, Delete) operations on
	// user data. The `db` variable is passed as an argument to the `NewRepo()` function to establish a
	// connection to the MongoDB database. The resulting `userRepo` variable is then used to pass the user
	// data to the authentication routes defined in the `routes` package.
	userRepo := auth.NewRepo(db)
	// The line `userSvc := auth.NewAuthService(userRepo.(*auth.Repo))` is creating a new instance of the
	// `auth.AuthService` struct, which is used to handle the logic and operations related to user
	// authentication.
	userSvc := auth.NewAuthService(userRepo.(*auth.Repo))
	// The line `allquestionRepo := allquestions.NewRepo(db)` is creating a new instance of the
	// `allquestions.Repo` struct, which is used to interact with the MongoDB database and perform CRUD
	// (Create, Read, Update, Delete) operations on all question data. The `db` variable, which represents
	// the MongoDB database connection, is passed as an argument to the `NewRepo()` function to establish a
	// connection to the database. The resulting `allquestionRepo` variable is then used to pass the all
	// question data to the routes defined in the `routes` package.
	allquestionRepo := allquestions.NewRepo(db)
	// `routes.CreatePhoneOtpRoutes(app, userSvc)` is creating and registering HTTP routes related to phone
	// OTP (One-Time Password) verification in the Fiber application. It is passing the `app` instance of
	// the Fiber application and a pointer to the `auth.AuthService` struct instance `userSvc` to the
	// `CreatePhoneOtpRoutes` function, which will define and register the necessary routes for phone OTP
	// verification. The `userSvc` instance is used to handle the logic and operations related to phone OTP
	// verification, such as sending OTPs and verifying OTPs.
	routes.CreatePhoneOtpRoutes(app, userSvc)
	// `routes.CreateAuthRoutes(app, userRepo.(*auth.Repo))` is creating and registering HTTP routes
	// related to user authentication in the Fiber application. It is passing the `app` instance of the
	// Fiber application and a pointer to the `auth.Repo` struct instance `userRepo` to the
	// `CreateAuthRoutes` function, which will define and register the necessary routes for user
	// authentication. The `userRepo.(*auth.Repo)` syntax is used to convert the `userRepo` variable to a
	// pointer to the `auth.Repo` struct type, which is required by the `CreateAuthRoutes` function.
	routes.CreateAuthRoutes(app, userRepo.(*auth.Repo), userSvc)
	// `routes.CreateAllQuestionRoutes(app, allquestionRepo.(*allquestions.Repo))` is creating and
	// registering HTTP routes related to all question data in the Fiber application. It is passing the
	// `app` instance of the Fiber application and a pointer to the `allquestions.Repo` struct instance
	// `allquestionRepo` to the `CreateAllQuestionRoutes` function, which will define and register the
	// necessary routes for all question data. The `allquestionRepo.(*allquestions.Repo)` syntax is used
	// to convert the `allquestionRepo` variable to a pointer to the `allquestions.Repo` struct type,
	// which is required by the `CreateAllQuestionRoutes` function.
	routes.CreateAllQuestionRoutes(app, allquestionRepo.(*allquestions.Repo))
	// `log.Panic(app.Listen(":" + os.Getenv("PORT")))` is a line of code that starts the Fiber
	// application and listens for incoming HTTP requests on the specified port.
	log.Panic(app.Listen(":" + os.Getenv("PORT")))
}

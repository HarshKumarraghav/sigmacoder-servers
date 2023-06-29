package routes

// The `import` statement is importing various packages that are needed for the implementation of the
// phone OTP routes in a Fiber app. These packages include:
import (
	"context"
	"log"
	"net/http"
	"os"
	"sigmacoder/pkg/auth"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

// `const appTimeout = time.Second * 10` is defining a constant variable `appTimeout` with a value of
// 10 seconds, which is the maximum amount of time allowed for a request to be processed before timing
// out. This is used in the `context.WithTimeout` function to set a timeout for the request context.
const appTimeout = time.Second * 10

// The OTPData type represents data for a phone number used in one-time password authentication, with
// the phone number being a required field.
// @property {string} PhoneNumber - PhoneNumber is a property of the OTPData struct that represents the
// phone number of the user for whom the OTP (One-Time Password) is being generated. It is of type
// string and has a JSON tag "phoneNumber" which is used for marshaling and unmarshaling JSON data. The
// "omitempty"
type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

// The VerifyData type contains a user's OTPData and a code to be validated, both of which are
// required.
// @property User - User is a pointer to an OTPData struct. It is marked as omitempty, which means that
// if the value is nil, it will not be included in the JSON output. The field is also marked as
// required in the validation tag, which means that it must be present and not nil when validating
// @property {string} Code - The "Code" property is a string that represents the OTP (One-Time
// Password) code that the user has entered for verification. It is a required field and must be
// provided in order to verify the user's identity.
type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}

// The type `jsonResponse` represents a JSON response with a status code, message, and data.
// @property {int} Status - Status is an integer property that represents the status of the response.
// It is typically used to indicate whether the request was successful or not. For example, a status
// code of 200 typically indicates success, while a status code of 404 indicates that the requested
// resource was not found.
// @property {string} Message - The "Message" property is a string that represents a message or
// description related to the response. It can be used to provide additional information about the
// response status or any errors that occurred during the request.
// @property {any} Data - The `Data` property is a field in the `jsonResponse` struct that represents
// the actual data being returned in the JSON response. The `any` type used for this field indicates
// that the data can be of any type, allowing for flexibility in the type of data that can be returned.
type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// `var validate = validator.New()` is creating a new instance of the `validator` struct from the
// `github.com/go-playground/validator/v10` package. This instance is used to validate the request body
// data in the `validateBody` function.
var validate = validator.New()

// The function validates the request body using a given struct and returns an error if validation
// fails.
func validateBody(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}

// The function writes a JSON response with a success message and data to a Fiber context.
func writeJSON(c *fiber.Ctx, status int, data interface{}) {
	c.JSON(jsonResponse{Status: status, Message: "success", Data: data})
}

// The function returns a JSON response with an error message and status code.
func errorJSON(c *fiber.Ctx, err error, status ...int) {
	statusCode := fiber.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.Status(statusCode).JSON(jsonResponse{Status: statusCode, Message: err.Error()})
}

// The function loads the .env file and returns the value of the TWILIO_ACCOUNT_SID environment
// variable.
func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

// This function loads the .env file and returns the value of the TWILIO_AUTHTOKEN environment
// variable.
func envAUTHTOKEN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}

// This function loads the environment variables from a .env file and returns the value of the
// TWILIO_SERVICES_ID variable.
func envSERVICESID() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}

// This line of code is creating a new instance of the `twilio.RestClient` struct and assigning it to
// the `client` variable. The `twilio.NewRestClientWithParams()` function is used to create the new
// instance, and it takes a `twilio.ClientParams` struct as an argument. The `twilio.ClientParams`
// struct contains the `Username` and `Password` fields, which are set to the values of the
// `envACCOUNTSID()` and `envAUTHTOKEN()` functions, respectively. These functions load the values of
// the `TWILIO_ACCOUNT_SID` and `TWILIO_AUTHTOKEN` environment variables from a `.env` file and return
// them. The `client` variable is then used to make requests to Twilio's API.
var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

// The function sends an OTP (one-time password) to a phone number using Twilio's API.
func twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

// The function verifies an OTP code sent to a phone number using Twilio API.
func twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}

// The function sends an OTP SMS message using Twilio API and returns a success message.
func sendSMS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var payload OTPData
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		newData := OTPData{
			PhoneNumber: payload.PhoneNumber,
		}
		_, err := twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			errorJSON(c, err)
			return err
		}
		writeJSON(c, http.StatusAccepted, "OTP sent successfully")
		return nil
	}
}

// The function verifies an SMS OTP code using Twilio API and returns a success message.
func verifySMS(svc auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, cancel := context.WithTimeout(c.Context(), appTimeout)
		defer cancel()
		var payload VerifyData

		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		newData := VerifyData{
			User: payload.User,
			Code: payload.Code,
		}
		token, err := svc.LoginPhoneOtp(newData.User.PhoneNumber)
		if err != nil {
			errorJSON(c, err)
			return err
		}

		err = twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			errorJSON(c, err)
			return err
		}
		return c.JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "OTP verified successfully",
			"token":   token,
		})

	}
}

// The function creates two routes for sending and verifying phone OTPs in a Fiber app.
func CreatePhoneOtpRoutes(app *fiber.App, svc auth.Service) {
	app.Post("/api/auth/sendotp", sendSMS())
	app.Post("/api/auth/verifyotp", verifySMS(svc))
}

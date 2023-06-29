package configuration

// `import "os"` is importing the `os` package, which provides a way to interact with the operating
// system. In this specific code, it is used to retrieve environment variables using the `os.Getenv()`
// function.
import "os"

// `var config Config` is declaring a variable named `config` of type `Config`. This variable will be
// used to store the configuration values retrieved from environment variables.
var config Config

// The Config type contains fields for a MongoDB URI, a port number, and a JWT secret.
// @property {string} MongoURI - MongoURI is a string that represents the connection string for MongoDB
// database. It typically includes the username, password, host, port, and database name.
// @property {string} Port - The `Port` property is a string that represents the port number on which
// the server will listen for incoming requests. This is typically a number between 0 and 65535 that is
// used to identify a specific process to which network traffic should be directed.
// @property {string} JwtSecret - JwtSecret is a property in the Config struct that represents the
// secret key used for JSON Web Token (JWT) authentication. JWT is a popular method for securely
// transmitting information between parties as a JSON object. The secret key is used to sign and verify
// the authenticity of the token.
type Config struct {
	MongoURI  string
	Port      string
	JwtSecret string
}

// The function retrieves configuration values from environment variables and returns them as a Config
// struct.
func FromEnv() Config {
	config := Config{
		MongoURI:  os.Getenv("MONGO_URI"),
		Port:      os.Getenv("PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
	return config
}

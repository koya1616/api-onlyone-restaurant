package configs

import (
	"os"
)

func EnvMongoURI() string {
	return os.Getenv("MONGOURI")
}

func EnvPort() string {
	return os.Getenv("PORT")
}

func EnvNeonDBString() string {
	return os.Getenv("NEONDB_STRING")
}

func EnvAdminName() string {
	return os.Getenv("ADMIN_NAME")
}

func EnvAdminPassword() string {
	return os.Getenv("ADMIN_PASSWORD")
}

func EnvJWTSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}

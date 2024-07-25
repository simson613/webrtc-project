package config

import (
	"os"
)

type SwaggerInterface interface {
	Path() string
}

type Swagger struct {
	path string
}

func initSwaggerConfig() *Swagger {
	path := os.Getenv("SWAGGER_PATH")
	if path == "" {
		path = "/auth"
	}

	return &Swagger{
		path: path,
	}
}

func (swagger *Swagger) Path() string {
	return swagger.path
}

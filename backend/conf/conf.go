package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

const EnvPrefix = "SUPER_BRUNFO"

var version string

type Main struct {
	AppHost     string
	AppPort     string
	DatabaseUrl string
	Version     string
}

func Get() (*Main, error) {
	secret := os.Getenv(fmt.Sprintf("%s_SECRET", EnvPrefix))

	c := new(Main)
	if err := json.Unmarshal([]byte(secret), c); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if c.AppHost == "" {
		c.AppHost = "0.0.0.0"
	}

	if c.AppPort == "" {
		c.AppPort = "8080"
	}

	c.Version = version

	return c, nil
}

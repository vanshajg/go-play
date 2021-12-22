package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type Config struct {
	Database struct {
		Dialect   string `default:"sqlite3"`
		Host      string `default:"book.db"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"false"`
	}
}

func Load() (*Config, string) {
	var env *string
	if value := os.Getenv("WEB_APP_ENV"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "dev", "switch environments")
	}

	config := &Config{}
	if err := configor.Load(config, "application."+*env+".yaml"); err != nil {
		fmt.Printf("Failed to read application.%s.yaml : %s", *env, err)
		os.Exit(2)
	}
	return config, *env
}

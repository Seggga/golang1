package conf

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

//ConfigType - is a structure with configuration elements user application required
type ConfigType struct {
	Port        uint32
	DBURL       string
	JaegerURL   string
	SentryURL   string
	KafkaBroker string
	SomeAppID   string
	SomeAppKey  string
}

//ReadFlags reads config data from command-line flags and writes it back in the ConfigType structure
func ReadFlags(config *ConfigType) error {

	// a set of flags used by the "flag"-package
	var (
		port        = flag.Uint("PORT", 8000, "Server port-number")
		dbURL       = flag.String("db_url", "", "Database URL (postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable).")
		jaegerURL   = flag.String("jaeger_url", "", "Jaeger URL with port number (http://example:9000).")
		sentryURL   = flag.String("sentry_url", "", "Sentry URL with port number(http://sentry:9000).")
		kafkaBroker = flag.String("kafka_broker", "", "Kafka broker URL with port number (kafka:9092).")
		someAppID   = flag.String("some_app_id", "", "Unique application identifier (testid).")
		someAppKey  = flag.String("some_app_key", "", "Test key for the application (testkey).")
	)

	flag.Parse()

	// obtain PORT data
	config.Port = uint32(*port)

	// data validation for "db_url" flag
	if matched, _ := regexp.MatchString(`\w+://\w+.?\w+\:`, *dbURL); !matched && *dbURL != "" {
		err := fmt.Errorf("Not valid data for the flag db_url, the field should be a URL-string, you entered %v", *dbURL)
		return err
	}
	config.DBURL = *dbURL

	// data validation for "jaeger_url" flag
	if matched, _ := regexp.MatchString(`\w+://\w+.?\w+\:`, *jaegerURL); !matched && *jaegerURL != "" {
		err := fmt.Errorf("Not valid data for the flag jaeger_url, the field should be a URL-string with port")
		return err
	}
	config.JaegerURL = *jaegerURL

	// data validation for "sentry_url" flag
	if matched, _ := regexp.MatchString(`\w+://\w+.?\w+\:`, *sentryURL); !matched && *sentryURL != "" {
		err := fmt.Errorf("Not valid data for the flag sentry_url, the field should be a URL-string with port")
		return err
	}
	config.SentryURL = *sentryURL

	// data validation for "kafka_broker" flag
	if matched, _ := regexp.MatchString(`\w+:\d+`, *kafkaBroker); !matched && *kafkaBroker != "" {
		err := fmt.Errorf("Not valid data for the flag kafka_broker, the field should be a URL-string with port")
		return err
	}
	config.KafkaBroker = *kafkaBroker

	config.SomeAppID = *someAppID
	config.SomeAppKey = *someAppKey

	return nil
}

//ReadEnv reads data from Environment variables and fills up the ConfigType structure
func ReadEnv(config *ConfigType) error {

	envData := os.Getenv("PORT")

	if envData == "" {
		envData = "8000"
	} else if matched, _ := regexp.MatchString(`^\d+`, envData); !matched {
		err := fmt.Errorf("Not valid data for the PORT variable. Entered value is %s", envData)
		return err
	}
	number, err := strconv.Atoi(envData)
	if err != nil {
		err1 := fmt.Errorf("Error parsing port number %v", err)
		return err1
	}
	config.Port = uint32(number)

	// check mandatory fields
	envData = os.Getenv("db_url")
	if envData == "" {
		err := fmt.Errorf("Database URL is not set")
		return err
	}
	if matched, _ := regexp.MatchString(`\w+://\w+.?\w+\:`, envData); !matched {
		err := fmt.Errorf("Variable db_url fail, expected URL-string, Entered value is %s", envData)
		return err
	}
	config.DBURL = envData

	// data validation for "jaeger_url" variable
	envData = os.Getenv("jaeger_url")
	if envData != "" {
		if matched, _ := regexp.MatchString(`\w+://\w+.?\w+\:`, envData); !matched {
			err := fmt.Errorf("Variable jaeger_url fail. Expected URL-string with port. Entered value is %s", envData)
			return err
		}
		config.JaegerURL = envData
	}

	// data validation for "sentry_url" variable
	envData = os.Getenv("sentry_url")
	if envData != "" {
		if matched, _ := regexp.MatchString(`\w+://\w+.?\w+\:`, envData); !matched {
			err := fmt.Errorf("Variable sentry_url fail. Expected URL-string with port. Entered value is %s", envData)
			return err
		}
		config.SentryURL = envData
	}

	// data validation for "kafka_broker" variable
	envData = os.Getenv("kafka_broker")
	if envData != "" {
		if matched, _ := regexp.MatchString(`\w+:\d+`, envData); !matched {
			err := fmt.Errorf("Variable kafka_broker fail. Expected URL-string with port. Entered value is %s", envData)
			return err
		}
		config.KafkaBroker = envData
	}

	// obtain "some_app_id" data
	envData = os.Getenv("some_app_id")
	if envData != "" {
		config.SomeAppID = envData
	}

	// obtain "some_app_key" data
	envData = os.Getenv("some_app_key")
	if envData != "" {
		config.SomeAppKey = envData
	}

	return nil
}

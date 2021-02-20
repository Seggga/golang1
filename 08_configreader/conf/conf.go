package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

//ConfigStruct - is a structure with configuration elements user application required
type ConfigStruct struct {
	Port        uint32
	DBURL       string
	JaegerURL   string
	SentryURL   string
	KafkaBroker string
	SomeAppID   string
	SomeAppKey  string
}

//ConfigJSON - is a structure with annotation for JSON unmarshaling
type ConfigJSON struct {
	Port        uint32 `json:"port"`
	DBURL       string `json:"db_url"`
	JaegerURL   string `json:"jaeger_url"`
	SentryURL   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppID   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

//ConfigYAML - is a structure with annotation for YAML unmarshalling
type ConfigYAML struct {
	Port        uint32 `yaml:"port"`
	DBURL       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

//ReadFlags reads config data from command-line flags and writes it back in the ConfigType structure
func ReadFlags(config *ConfigStruct) error {

	// a set of flags used by the "flag"-package
	var (
		port        = flag.Uint("PORT", 8000, "Server port-number")
		dbURL       = flag.String("db_url", "", "Database URL (postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable).")
		jaegerURL   = flag.String("jaeger_url", "", "Jaeger URL with port number (http://example:9000).")
		sentryURL   = flag.String("sentry_url", "", "Sentry URL with port number(http://sentry:9000).")
		kafkaBroker = flag.String("kafka_broker", "", "Kafka broker URL with port number (ssl://kafka:9092).")
		someAppID   = flag.String("some_app_id", "", "Unique application identifier (testid).")
		someAppKey  = flag.String("some_app_key", "", "Test key for the application (testkey).")
	)

	flag.Parse()

	// obtain PORT data
	config.Port = uint32(*port)

	// data validation 1 of 4
	// check, if db_url is valid, mandatory field
	if *dbURL == "" {
		return fmt.Errorf("error processing db_url, the field is mandatory")
	}
	if err := isURL(*dbURL); err != nil {
		return fmt.Errorf("error processing db_url, %v", err)
	}
	config.DBURL = *dbURL

	// data validation 2 of 4
	// check, if jaeger_url is valid
	if err := isURL(*jaegerURL); err != nil && *jaegerURL != "" {
		return fmt.Errorf("error processing jaeger_url, %v", err)
	}
	config.JaegerURL = *jaegerURL

	// data validation 3 of 4
	// data validation for "sentry_url" flag
	if err := isURL(*sentryURL); err != nil && *sentryURL != "" {
		return fmt.Errorf("error processing sentry_url, %v", err)
	}
	config.SentryURL = *sentryURL

	// data validation 4 of 4
	// data validation for "kafka_broker" flag
	if err := isURL(*kafkaBroker); err != nil && *kafkaBroker != "" {
		return fmt.Errorf("error processing kafka_broker, %v", err)
	}
	config.KafkaBroker = *kafkaBroker

	config.SomeAppID = *someAppID
	config.SomeAppKey = *someAppKey

	return nil
}

//ReadEnv reads data from Environment variables and fills up the ConfigType structure
func ReadEnv(config *ConfigStruct) error {

	// check, if PORT is valid, has default value 8000
	envData := os.Getenv("PORT")
	if envData == "" {
		envData = "8000"
	} else if matched, _ := regexp.MatchString(`^\d+`, envData); !matched {
		err := fmt.Errorf("not valid data for the PORT variable. Entered value is %s", envData)
		return err
	}
	number, err := strconv.Atoi(envData)
	if err != nil {
		return fmt.Errorf("error parsing port number %v", err)
	}
	config.Port = uint32(number)

	// data validation 1 of 4
	// check, if db_url is valid, mandatory field
	envData = os.Getenv("db_url")
	if envData == "" {
		return fmt.Errorf("error processing db_url, the field is mandatory")
	}
	if err := isURL(envData); err != nil {
		return fmt.Errorf("error processing db_url, %v", err)
	}
	config.DBURL = envData

	// data validation 2 of 4
	// check, if jaeger_url is valid
	envData = os.Getenv("jaeger_url")
	if err := isURL(envData); err != nil && envData != "" {
		return fmt.Errorf("error processing jaeger_url, %v", err)
	}
	config.JaegerURL = envData

	// data validation 3 of 4
	// check, if sentry_url is valid
	envData = os.Getenv("sentry_url")
	if err := isURL(envData); err != nil && envData != "" {
		return fmt.Errorf("error processing sentry_url, %v", err)
	}
	config.SentryURL = envData

	// data validation 4 of 4
	// check, if kafka_broker is valid
	envData = os.Getenv("kafka_broker")
	if err := isURL(envData); err != nil && envData != "" {
		return fmt.Errorf("error processing kafka_broker, %v", err)
	}
	config.KafkaBroker = envData

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

//ReadFile detects the filetype and decides, what function to call for config-parsing
func ReadFile(config *ConfigStruct, filePath string) error {
	//obtain "filepath" extension
	dotSlice := strings.Split(strings.ToLower(filePath), ".")
	fileExt := dotSlice[len(dotSlice)-1]
	switch fileExt {
	case "yml":
		return readYAML(config, filePath)
	case "yaml":
		return readYAML(config, filePath)
	case "json":
		return readJSON(config, filePath)
	default:
		return fmt.Errorf("unknown file type, *.json / *.yml / *.yaml expexted, *.%s detected", fileExt)
	}
}

//ReadYAML reads data config from YAML-file and fills up the ConfigStruct structure
func readYAML(config *ConfigStruct, filePath string) error {
	//file opening
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open file %s, %v", filePath, err)
	}
	//deffered file closing
	defer func() {
		err = file.Close()
		if err != nil {
			panic(fmt.Sprintf("error while closing file %s, %v", filePath, err))
		}
	}()

	configYAML := ConfigYAML{}
	err = yaml.NewDecoder(file).Decode(&configYAML)
	if err != nil {
		log.Printf("cannot decode YAML file %s in a structure: %v", filePath, err)
	}

	config.Port = configYAML.Port
	config.DBURL = configYAML.DBURL
	config.JaegerURL = configYAML.JaegerURL
	config.SentryURL = configYAML.SentryURL
	config.KafkaBroker = configYAML.KafkaBroker
	config.SomeAppID = configYAML.SomeAppID
	config.SomeAppKey = configYAML.SomeAppKey

	return nil
}

//ReadJSON reads data config from JSON-file and fills up the ConfigStruct structure
func readJSON(config *ConfigStruct, filePath string) error {
	//file opening
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open file %s, %v", filePath, err)
	}
	//deffered file closing
	defer func() {
		err = file.Close()
		if err != nil {
			panic(fmt.Sprintf("error while closing file %s, %v", filePath, err))
		}
	}()

	configJSON := ConfigJSON{}
	err = json.NewDecoder(file).Decode(&configJSON)
	if err != nil {
		log.Printf("cannot decode JSON file %s in a structure: %v", filePath, err)
	}

	config.Port = configJSON.Port
	config.DBURL = configJSON.DBURL
	config.JaegerURL = configJSON.JaegerURL
	config.SentryURL = configJSON.SentryURL
	config.KafkaBroker = configJSON.KafkaBroker
	config.SomeAppID = configJSON.SomeAppID
	config.SomeAppKey = configJSON.SomeAppKey

	return nil
}

//isURL implements URL validation
func isURL(data string) error {

	url, err := url.Parse(data)
	if err != nil {
		return err
	}
	// no protocol in db_url
	if url.Scheme == "" {
		return fmt.Errorf("No protocol has been set. Entered data is %v", data)
	}
	// no servername/address in db_url
	if url.Host == "" {
		return fmt.Errorf("No host has been set. Entered data is %v", data)
	}
	return nil
}

package main

import (
	"encoding/json"
	"flag"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// config struct to hold data for CLI flag arguments
type config struct {
	FILENAME    string
	LOGRESPONSE struct {
		ENABLED   bool
		DIRECTORY string
	}
}

// body struct holding data from JSON files
type body struct {
	METHOD      string `json:"method"`
	URL         string `json:"url"`
	HEADERS     []map[string]string
	FORMAT      string         `json:"format"`
	BODY        map[string]any `json:"body"`
	CREDENTIALS struct {
		USERNAME string
		PASSWORD string
	} `json:"credentials"`
}

type application struct {
	config  config
	payload body
	logger  *slog.Logger
}

func (app *application) runGet() (map[string]interface{}, error) {
	request, err := http.NewRequest(http.MethodGet, app.payload.URL, nil)
	if err != nil {
		app.logger.Error("Error connecting to endpoint", err)
		return nil, err
	}
	response, err := Client.Do(request)
	if err != nil {
		app.logger.Error("Error took place running GET request", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			app.logger.Error("Failed to close body closer", err)
		}
	}(response.Body)
	var m map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&m)

	if err != nil {
		app.logger.Error("Error encoding json object", err)
		return nil, err
	}

	return m, nil
}

func main() {
	var cfg config
	var payload body
	flag.StringVar(&cfg.FILENAME, "note", "", "Path to note to post")
	flag.BoolVar(&cfg.LOGRESPONSE.ENABLED, "log.enabled", true, "true or false to enable saved run logs")
	flag.StringVar(&cfg.LOGRESPONSE.DIRECTORY, "log.directory", "", "Directory where you want saved runs to be stored")
	flag.Parse()

	// init our new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	content, err := os.ReadFile(cfg.FILENAME)
	if err != nil {
		logger.Error("Error could not read file: ")
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		logger.Error("Error parsing JSON file", err)
	}
	logger.Info("METHOD", payload.URL)

	app := application{
		config:  cfg,
		payload: payload,
		logger:  logger,
	}

	switch {
	case strings.Contains(payload.METHOD, "GET"):
		response, err := app.runGet()
		if err != nil {
			app.logger.Error("Error in stack", err)
		}
		app.logger.Info("response data: ", response)

	}

}

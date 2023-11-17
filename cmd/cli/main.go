package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
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

func (app *application) runGet() ([]byte, error) {

	out, err := json.Marshal(app.payload.BODY)
	request, err := http.NewRequest(app.payload.METHOD, app.payload.URL, bytes.NewBuffer(out))

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
	content, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		app.logger.Error("Error Marshalling JSON", err)
		return nil, err
	}
	// parse the users input of -note to split out the path/filename.json
	filePathSplit := strings.Split(app.config.FILENAME, ".")

	// -2 should ALWAYS be the filename and not json extension
	fileNameData := filePathSplit[len(filePathSplit)-2]
	// split out the directory "/"
	fileNameSplit := strings.Split(fileNameData, "/")
	// get the last item in the slice
	requestFileName := fileNameSplit[len(fileNameSplit)-1]

	fileName := fmt.Sprintf("%s/Request-%s-%v", app.config.LOGRESPONSE.DIRECTORY, requestFileName, time.Now().Format("2006-01-02-15:04:05.json"))
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		app.logger.Error("Error writing data to file")
		return nil, err
	}
	app.logger.Info("Successfully wrote data to ", "file", fileName)

	if err != nil {
		app.logger.Error("Error encoding json object", err)
		return nil, err
	}
	return content, nil
}

func main() {
	var cfg config
	var payload body
	flag.StringVar(&cfg.FILENAME, "note", "", "Path to note to post")
	flag.BoolVar(&cfg.LOGRESPONSE.ENABLED, "log.enabled", true, "true or false to enable saved run logs")
	flag.StringVar(&cfg.LOGRESPONSE.DIRECTORY, "log.directory", "", "Directory where you want saved runs to be stored")
	flag.Parse()

	// init our new logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	content, err := os.ReadFile(cfg.FILENAME)
	if err != nil {
		logger.Error("Error could not read file", "error", err)
		os.Exit(1)
	}

	err = json.Unmarshal(content, &payload)
	if err != nil {
		logger.Error("Error parsing JSON file", "parsing error", err)
		os.Exit(1)
	}

	app := application{
		config:  cfg,
		payload: payload,
		logger:  logger,
	}

	response, err := app.runGet()
	if err != nil {
		app.logger.Error("Error in stack", err)
	}
	fmt.Println(string(response))

}

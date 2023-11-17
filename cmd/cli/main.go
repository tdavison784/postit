package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"os"
)

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
	config config
	logger *slog.Logger
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
	logger.Info("METHOD", payload.METHOD)

}

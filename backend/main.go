package main

import (
	"go.uber.org/zap"
)

// Structure for environment variables required
type Environment struct {
	emailAddress  string
	emailPassword string
	apiPort       string
}

// Variable for environment variables
var env Environment

// variable for logger
var logger *zap.Logger

func main() {
	// Set up logger
	logger = setUpLogger()
	// Set up env variables
	setUpEnvironment()
	logger.Info("Starting API service. Using port " + env.apiPort)
	// Set up API
	setUpGinEngine()
}

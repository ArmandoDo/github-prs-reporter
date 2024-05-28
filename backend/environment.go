package main

import (
	smptserver "backend/smtpserver"
	"os"

	"go.uber.org/zap"
)

// setUpEnvironment sets up the environment variables required for the app
func setUpEnvironment() {
	env.emailAddress = os.Getenv("EMAIL_ADDRESS")
	env.emailPassword = os.Getenv("EMAIL_PASSWORD")
	env.apiPort = os.Getenv("API_PORT")
	// Exit code if there is any env var missing
	if env.emailAddress == "" || env.emailPassword == "" || env.apiPort == "" {
		logger.Fatal(
			"MISSING ENV VARIABLES",
			zap.String("EMAIL_ADDRESS", env.emailAddress),
			zap.String("EMAIL_PASSWORD", env.emailPassword),
			zap.String("API_PORT", env.apiPort),
		)
		os.Exit(1)
	}
	// Set up smpt server using env variables
	smpt = &smptserver.Server{
		Email:    env.emailAddress,
		Password: env.emailPassword,
		Host:     "smtp.gmail.com",
		Port:     "587",
	}
}

package main

import (
	smptserver "backend/smtpserver"

	"go.uber.org/zap"
)

// Variable for smpt server
var smpt *smptserver.Server

// sendEmail sends an email with an attached file
func sendEmail(email, owner, repo, file string) error {
	// Set up smpt server
	sender := smptserver.New(smpt)
	// Set up email
	m := smptserver.NewMessage("Report of PRs for "+owner+"/"+repo,
		"This is the report of the latest PRs of"+owner+"/"+repo+"for the last 7 days.")
	m.To = []string{email}
	// Attach file
	err := m.AttachFile(file)
	if err != nil {
		return err
	}
	// Send email
	err = sender.Send(m, smpt)
	if err != nil {
		return err
	}
	// Generate log info
	logger.Info(
		"The report has been sent",
		zap.String("repo", repo),
		zap.String("owner", owner),
		zap.String("email", email),
		zap.String("body", m.Body),
		zap.String("subject", m.Subject),
	)

	return nil
}

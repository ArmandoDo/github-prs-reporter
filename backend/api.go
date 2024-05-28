package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ginEngine sets up the Gin Engine for the API
func setUpGinEngine() {
	eng := gin.New()
	eng.Use(cors.Default())
	eng.GET("/pullrequests", pullRequestReport)
	eng.Run(":" + env.apiPort)
}

// pullRequestReport generates a report if there are Pull requests during the last 7 days
func pullRequestReport(c *gin.Context) {
	email := c.Query("email")
	owner := c.Query("owner")
	repo := c.Query("repo")

	empty, err := getPullRequestsFromApi(email, owner, repo)
	// Check for errors
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else if empty {
		// Not PRs during the last 7 days
		c.JSON(http.StatusOK, gin.H{"Skipped": "Not PRs in the last 7 days for the repo " + owner + "/" + repo + "."})
	} else {
		// Report has been sent successfully
		c.JSON(http.StatusOK, gin.H{"Successful": "Pull request list from " + owner + "/" + repo + " sent to " + email + "."})
	}
}

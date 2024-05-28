package main

import (
	"backend/githubapi"
	"backend/reporter"
	"os"
	"time"

	"go.uber.org/zap"
)

func getPullRequestsFromApi(email, owner, repo string) (empty bool, err error) {
	// Path to store the tmp files
	path := "/tmp/"
	// Create and config github client
	client := githubapi.NewClient()
	config := &githubapi.Config{
		Direction: "desc",
		Page:      1,
		PerPage:   100,
		RepoName:  repo,
		RepoOwner: owner,
		State:     "all",
		Sort:      "created",
	}

	// Initial date (7 days before today)
	fd, err := time.Parse("2006-01-02", time.Now().AddDate(0, 0, -7).Format("2006-01-02"))
	// Check for errors
	if err != nil {
		return true, err
	}

	// Generate List of PRs
	var list githubapi.PullRequests
	iterate := true
	for iterate {
		// Get list of PRs from Github API
		prs, err := client.GetPullRequests(config)
		// Check for errors
		if err != nil {
			return true, err
		}

		// Append pull requests of current page
		list.Pulls = append(list.Pulls, prs.Pulls...)
		// Looking for a prs interface empty or for the last page
		if len(prs.Pulls) == 0 || len(prs.Pulls) < config.PerPage {
			iterate = false
		} else {
			// Verify last creation date in the interface
			ld := prs.Pulls[len(prs.Pulls)-1].CreatedAt.UTC()
			if ld.Before(fd) {
				// stop iteration
				iterate = false
			} else {
				// Next page
				config.Page = config.Page + 1
			}
		}
	}

	// Fill the table with the data
	table := []*reporter.Table{}
	for _, pull := range list.Pulls {
		// Verify if current pr is after the date fom
		if pull.CreatedAt.After(fd) {
			table = append(table, &reporter.Table{
				Id:        *pull.ID,
				Number:    *pull.Number,
				State:     *pull.State,
				Title:     *pull.Title,
				CreatedAt: pull.CreatedAt,
				ClosedAt:  pull.ClosedAt,
				MergedAt:  pull.MergedAt,
				Head:      *pull.Head.Ref,
				Base:      *pull.Base.Ref,
				UserId:    *pull.User.ID,
				UserName:  *pull.User.Login,
				Url:       *pull.HTMLURL,
			})
		} else {
			break
		}
	}
	// If table is empty, return empty as true
	if len(table) == 0 {
		logger.Info(
			"Not PRs in the last 7 days",
			zap.String("email", email),
			zap.String("owner", owner),
			zap.String("repo", repo),
			zap.Time("from", fd),
		)
		return true, nil
	} else {
		// Generate excel file
		name := owner + "_" + repo + "_" + time.Now().Format("2006_01_02T15_04_05") + ".xlsx"
		err = generateExcelReport(path, name, table)
		// Delete file when the function finish
		defer func() {
			os.Remove(path + name)
		}()

		if err != nil {
			return true, err
		}
		// Send excel file
		err = sendEmail(email, owner, repo, path+name)
		if err != nil {
			return true, err
		}
	}

	return false, nil
}

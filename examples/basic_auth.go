package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

// This example shows how to create a client with username and password.
//
// NOTE: This is not an official/supported way to use the API. So if possible
// please use `NewClient` or `NewOAuthClient` instead.
func basicAuthExample() {
	git, err := gitlab.NewBasicAuthClient(nil, "https://gitlab.company.com", "svanharmelen", "password")
	if err != nil {
		log.Fatal(err)
	}

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d projects", len(projects))
}

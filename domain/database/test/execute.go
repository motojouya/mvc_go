package test

import (
	//_ "internal/shelter/timezone"
	"database/sql"
	"fmt"
	"github.com/motojouya/ddd_go/domain/database/core"
	"github.com/motojouya/ddd_go/domain/database/behavior"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
)

func ExecuteDatabaseTest(pathToRoot string, run func(core.ORPer) int) {
	os.Setenv("TZ", "Asia/Tokyo") // `internal/shelter/timezone`だとできんかった？

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=ddd_go",
			"POSTGRES_USER=ddd_go",
			"POSTGRES_DB=ddd_go",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// resource.Expire(120) // Tell docker to hard kill the container in 120 seconds
	// pool.MaxWait = 120 * time.Second // exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	var database *sql.DB
	databaseUrl := fmt.Sprintf("postgres://ddd_go:ddd_go@%s/ddd_go?sslmode=disable", resource.GetHostPort("5432/tcp"))
	if err = pool.Retry(func() error {
		database, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return database.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	var migrateErr = core.Migrate(database, pathToRoot)
	if migrateErr != nil {
		log.Fatalf("Could not migrate database: %s", migrateErr)
	}

	var orp = behavior.CreateDatabase(database)

	code := run(orp)

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

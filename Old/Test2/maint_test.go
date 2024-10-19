package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestValidateDbConnection(t *testing.T) {
	ctx := context.Background()

	// Define the PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"}, // Use 5433 instead of 5432
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithStartupTimeout(60 * time.Second),
	}

	// Start the container
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}
	defer postgresC.Terminate(ctx)

	// Get the container's host and port
	host, err := postgresC.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %v", err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get mapped port: %v", err)
	}

	fmt.Printf("PostgreSQL container is running on %s:%s\n", host, port.Port())

	// Wait a bit to ensure the database is fully ready
	time.Sleep(5 * time.Second)

	// Test cases
	testCases := []struct {
		name     string
		host     string
		user     string
		password string
		port     int
		useSSL   bool
		wantOk   bool
		wantErr  bool
	}{
		{
			name:     "Valid connection",
			host:     host,
			user:     "testuser",
			password: "testpass",
			port:     port.Int(),
			useSSL:   false,
			wantOk:   true,
			wantErr:  false,
		},
		{
			name:     "Invalid password",
			host:     host,
			user:     "testuser",
			password: "wrongpass",
			port:     port.Int(),
			useSSL:   false,
			wantOk:   false,
			wantErr:  true,
		},
		{
			name:     "Missing host",
			host:     "",
			user:     "testuser",
			password: "testpass",
			port:     port.Int(),
			useSSL:   false,
			wantOk:   false,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Printf("Running test case: %s\n", tc.name)
			ok, err := ValidateDbConnection(tc.host, tc.user, tc.password, tc.port, tc.useSSL)

			if tc.wantErr {
				assert.Error(t, err)
				if err != nil {
					fmt.Printf("Expected error occurred: %v\n", err)
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.wantOk, ok)
			fmt.Printf("Test case %s completed. ok: %v, err: %v\n", tc.name, ok, err)
		})
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	api "people/internal/api/people/v1"
	"people/internal/repository"
	"people/internal/repository/migrator"
	peopleRepo "people/internal/repository/people"
	"people/internal/service"
	peopleService "people/internal/service/people"
	peopleV1 "people/shared/pkg/openapi/people/v1"
)

type Container struct {
	DB *sql.DB

	Migrator   migrator.Migrator
	PeopleRepo repository.PeopleRepository
	EmailRepo  repository.EmailRepository
	FriendRepo repository.FriendRepository

	PeopleService service.PeopleService

	PeopleHandler peopleV1.Handler

	Server *peopleV1.Server
}

func NewContainer() (*Container, error) {
	container := &Container{}

	if err := container.initDB(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	container.initRepositories()

	if err := container.initMigrator(); err != nil {
		return nil, fmt.Errorf("failed to initialize migrator: %w", err)
	}

	container.initServices()

	container.initHandlers()

	if err := container.initServer(); err != nil {
		return nil, fmt.Errorf("failed to initialize server: %w", err)
	}

	return container, nil
}

func (c *Container) initDB() error {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	log.Println(dbHost, dbPort, dbUser, dbPass, dbName)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	c.DB = db
	log.Println("Database connection established")
	return nil
}

func (c *Container) initRepositories() {
	repo := peopleRepo.NewRepository(c.DB)

	c.PeopleRepo = repo
	c.EmailRepo = repo
	c.FriendRepo = repo

	log.Println("Repositories initialized")
}

func (c *Container) initMigrator() error {
	migrationsDir := os.Getenv("MIGRATIONS_DIR")
	if migrationsDir == "" {
		migrationsDir = "migrations"
	}
	c.Migrator = migrator.NewMigrator(c.DB, migrationsDir)
	log.Printf("Running migrations from %s...", migrationsDir)
	if err := c.Migrator.Up(); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	log.Println("Migrations applied successfully")
	return nil
}

func (c *Container) initServices() {
	peopleService := peopleService.NewService(c.PeopleRepo, c.EmailRepo, c.FriendRepo)

	c.PeopleService = peopleService

	log.Println("Services initialized")
}

func (c *Container) initHandlers() {
	c.PeopleHandler = api.NewPeopleHandler(c.PeopleService)

	log.Println("Handlers initialized")
}

func (c *Container) initServer() error {
	server, err := peopleV1.NewServer(c.PeopleHandler)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	c.Server = server

	log.Println("Server initialized")
	return nil
}

func (c *Container) Close() error {
	if c.DB != nil {
		if err := c.DB.Close(); err != nil {
			return fmt.Errorf("failed to close database: %w", err)
		}
		log.Println("Database connection closed")
	}
	return nil
}

func (c *Container) GetServer() *peopleV1.Server {
	return c.Server
}

func (c *Container) GetDB() *sql.DB {
	return c.DB
}

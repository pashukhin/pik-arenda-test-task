package main

import (
	"flag"
	"fmt"
	"github.com/pashukhin/pik-arenda-test-task/handlers"
	"github.com/pashukhin/pik-arenda-test-task/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	httpAddr, dbHost, dbPort, dbUser, dbPassword, dbDatabase := Params()

	db := Db(dbHost, dbPort, dbUser, dbPassword, dbDatabase)

	e := echo.New()

	setRestHandlers := func(name string, h handlers.CrudHandler) {
		e.GET(fmt.Sprintf("/%s", name), h.List)
		e.POST(fmt.Sprintf("/%s", name), h.Create)
		e.GET(fmt.Sprintf("/%s/:id", name), h.Get)
		e.PUT(fmt.Sprintf("/%s/:id", name), h.Update)
		e.DELETE(fmt.Sprintf("/%s/:id", name), h.Delete)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// services
	freeScheduleService := service.NewFreeSchedule(db)
	scheduleService := service.NewSchedule(db, freeScheduleService)
	taskService := service.NewTask(db)
	workerService := service.NewWorker(db)
	freeSlotService := service.NewFreeSlot(freeScheduleService)

	// Handlers
	setRestHandlers("worker", handlers.NewWorker(workerService))
	setRestHandlers("schedule", handlers.NewSchedule(scheduleService))
	setRestHandlers("task", handlers.NewTask(taskService))
	setRestHandlers("free_schedule", handlers.NewFreeSchedule(freeScheduleService))
	setRestHandlers("free_slot", handlers.NewFreeSlot(freeSlotService))

	// Start server
	e.Logger.Fatal(e.Start(*httpAddr))
}

func Db(dbHost *string, dbPort *int, dbUser *string, dbPassword *string, dbDatabase *string) *sqlx.DB {
	//db, err := sqlx.Connect("postgres", "host=127.0.0.1 port=5432 user=user password=password dbname=db sslmode=disable")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *dbHost, *dbPort, *dbUser, *dbPassword, *dbDatabase)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}
	return db
}

func Params() (*string, *string, *int, *string, *string, *string) {
	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address")
	//httpAddr := flag.String("http.addr", ":8081", "HTTP listen address")
	dbHost := flag.String("db.host", "db", "db host")
	dbPort := flag.Int("db.port", 5432, "db port")
	dbUser := flag.String("db.user", "user", "db user")
	dbPassword := flag.String("db.password", "password", "db password")
	dbDatabase := flag.String("db.database", "db", "database name")
	flag.Parse()
	return httpAddr, dbHost, dbPort, dbUser, dbPassword, dbDatabase
}

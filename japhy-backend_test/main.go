package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	charmLog "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/japhy-tech/backend-test/database_actions"
	"github.com/japhy-tech/backend-test/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
)

const (
	MysqlDSN = "root:root@(mysql-test:3306)/core?parseTime=true"
	ApiPort  = "5000"
)

func main() {
	logger := charmLog.NewWithOptions(os.Stderr, charmLog.Options{
		Formatter:       charmLog.TextFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "🧑‍💻 backend-test",
		Level:           charmLog.DebugLevel,
	})

	err := database_actions.InitMigrator(MysqlDSN)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// GORM
	gormDB, err := gorm.Open(mysql.Open(MysqlDSN), &gorm.Config{
		Logger: gormLog.New(
			logger, // io writer
			gormLog.Config{
				SlowThreshold:             time.Second,  // Slow SQL threshold
				LogLevel:                  gormLog.Info, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,         // Disable color
			},
		),
	})
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	db, err := gormDB.DB()
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	// Actual MySQL DB bit
	defer db.Close()
	db.SetMaxIdleConns(0)
	err = db.Ping()
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}

	// Run migrations. This will trigger gorm.Automigrate() first
	msg, err := database_actions.RunMigrate("up", 0, gormDB, &internal.Breed{})
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info(msg)
	}

	logger.Info("Database connected")
	// App setup
	app := internal.NewApp(logger, gormDB)
	// Routing bit
	r := mux.NewRouter()
	app.RegisterRoutes(r.PathPrefix("/v1").Subrouter())
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	err = http.ListenAndServe(
		net.JoinHostPort("", ApiPort),
		r,
	)
	// =============================== Starting Msg ===============================
	logger.Info(fmt.Sprintf("Service started and listen on port %s", ApiPort))
}

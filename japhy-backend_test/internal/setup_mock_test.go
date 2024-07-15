package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	charmLog "github.com/charmbracelet/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func setupTest() (*gorm.DB, sqlmock.Sqlmock, *charmLog.Logger) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("failed to open sqlmock database: %v\n", err)
		os.Exit(1)
	}
	logger := charmLog.NewWithOptions(os.Stdout, charmLog.Options{
		Formatter:       charmLog.TextFormatter,
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "🧑‍💻 backend-test",
		Level:           charmLog.DebugLevel,
	})

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: gormLogger.New(
			logger, // io writer
			gormLogger.Config{
				SlowThreshold:             time.Second,     // Slow SQL threshold
				LogLevel:                  gormLogger.Info, // Log level
				IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,            // Disable color
			},
		),
	})
	if err != nil {
		fmt.Printf("failed to open gorm DB: %v", err)
		os.Exit(1)
	}
	return gormDB, sqlMock, logger
}

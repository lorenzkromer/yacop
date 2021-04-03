package test_data

import (
	"fmt"
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

// Initializes application config and SQLite database used for testing
func init() {
	// the test may be started from the home directory or a subdirectory
	err := config.LoadConfig("../../../config/test") // on host use absolute path
	if err != nil {
		panic(err)
	}
	common.TestDBInit()
}

// Resets testing database - deletes all tables, creates new ones using GORM migration and populates them using `db.sql` file
func ResetDB() *gorm.DB {
	if err := runSQLFile(config.Config.DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return config.Config.DB
}

func getSQLFile() string {
	return "../test_data/db.sql" // on host use absolute path
}

func GetTestCaseFolder() string {
	return "../test_data/test_case_data" // on host use absolute path
}

// Executes SQL file specified by file argument
func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if result := db.Exec(line); result.Error != nil {
			fmt.Println(line)
			return result.Error
		}
	}
	return nil
}

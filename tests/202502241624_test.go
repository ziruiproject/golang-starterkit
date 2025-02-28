package tests

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"template-go/databases"
	"testing"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.NoError(t, err)
}

func TestDatabasePostgre(t *testing.T) {
	os.Setenv("DB_DRIVER", "psql")

	db := databases.DatabaseInit()

	assert.NotNil(t, db)
	assert.IsType(t, &gorm.DB{}, db)

	con, _ := db.DB()
	err := con.Close()

	assert.NoError(t, err)
}

func TestDatabaseMysql(t *testing.T) {
	t.Skip()
	os.Setenv("DB_DRIVER", "mysql")

	db := databases.DatabaseInit()

	assert.NotNil(t, db)
	assert.IsType(t, &gorm.DB{}, db)

	con, _ := db.DB()
	err := con.Close()

	assert.NoError(t, err)
}

func TestDatabaseSqlite(t *testing.T) {
	os.Setenv("DB_DRIVER", "sqlite")
	os.Setenv("SQLITE_PATH", fmt.Sprintf("../%s", os.Getenv("SQLITE_PATH")))

	db := databases.DatabaseInit()

	assert.NotNil(t, db)
	assert.IsType(t, &gorm.DB{}, db)

	con, _ := db.DB()
	err := con.Close()

	assert.NoError(t, err)
}

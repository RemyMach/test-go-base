package db

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type MyTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *MyTestSuite) SetupTest() {
	db, err := sql.Open("postgres", "user=test password=test dbname=test sslmode=disable")
	suite.NoError(err)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS MyTable (MyColumn VARCHAR(255))")
	suite.NoError(err)

	_, err = db.Exec("INSERT INTO MyTable (MyColumn) VALUES ('Test data')")
	suite.NoError(err)

	suite.Db = db
}

func (suite *MyTestSuite) TearDownTest() {
	// Nettoyez votre base de données ici. Par exemple, supprimez toutes les lignes.
	_, err := suite.Db.Exec("DELETE FROM MyTable")
	suite.NoError(err)

	suite.Db.Close()
}

func (suite *MyTestSuite) TestDeleteRow() {
	// Supprimez une ligne ici et vérifiez ensuite qu'elle a bien été supprimée.
	_, err := suite.Db.Exec("DELETE FROM MyTable WHERE MyColumn = 'Test data'")
	suite.NoError(err)

	var count int
	err = suite.Db.QueryRow("SELECT COUNT(*) FROM MyTable WHERE MyColumn = 'Test data'").Scan(&count)
	suite.NoError(err)
	suite.Equal(0, count)
}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(MyTestSuite))
}

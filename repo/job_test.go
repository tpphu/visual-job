package repo

import (
	"errors"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type JobRepoTestSuite struct {
	suite.Suite
	jobRepo JobRepoImpl
	mock    sqlmock.Sqlmock
}

func (suite *JobRepoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	jobRepo := JobRepoImpl{}
	jobRepo.DB, _ = gorm.Open("mysql", db)
	suite.jobRepo = jobRepo
}

func (suite *JobRepoTestSuite) TearDownTest() {
	suite.jobRepo.DB.Close()
}

func TestJobRepoTestSuite(t *testing.T) {
	suite.Run(t, new(JobRepoTestSuite))
}

func (suite *JobRepoTestSuite) TestUserRepoFind() {
	suite.Run("find with having found id", func() {
		var ID uint = 5
		// Mock du lieu tra ve
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
			AddRow(ID, time.Now(), time.Now(), nil, "Phu")

		// Trong truong ho query
		suite.mock.ExpectQuery("SELECT \\* FROM `jobs`").
			WillReturnRows(rows)

		actual, err := suite.jobRepo.Find(ID)
		if err != nil {
			suite.Fail("Error should be nil")
		}
		if actual.ID != ID {
			suite.Fail("Id should be same")
		}
		if actual.DeletedAt != nil {
			suite.Fail("DeletedAt should be nil")
		}
	})

	suite.Run("find with not found id", func() {
		var ID uint = 6
		// Trong turong hop khong co cai id
		suite.mock.ExpectQuery("SELECT \\* FROM `jobs`").
			WillReturnError(errors.New("record not found"))
		_, err := suite.jobRepo.Find(ID)
		if err == nil {
			suite.Fail("Error should be not nil")
		}
	})
}

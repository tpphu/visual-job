package repo

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/tpphu/visual-job/model"
	"syreclabs.com/go/faker"
)

type ProductRepoTestSuite struct {
	suite.Suite
	productRepo ProductRepoImpl
	mock        sqlmock.Sqlmock
}

func (suite *ProductRepoTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	productRepo := ProductRepoImpl{}
	productRepo.DB, _ = gorm.Open("mysql", db)
	suite.productRepo = productRepo
}

func (suite *ProductRepoTestSuite) TearDownTest() {
	suite.productRepo.DB.Close()
}

func TestProductRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepoTestSuite))
}

func (suite *ProductRepoTestSuite) TestproductRepoCreate() {
	suite.Run("create with valid data", func() {
		var productID uint = 5
		product := model.Product{
			Name:     "name",
			Category: "category",
			Price:    1000,
		}
		// Mock SQL / DB
		suite.mock.ExpectExec("INSERT INTO `products`").
			WillReturnResult(sqlmock.NewResult(
				int64(productID),
				1,
			))
		actual, err := suite.productRepo.Create(product)
		if err != nil {
			suite.Fail("error should be nil")
		}
		if actual.ID != productID {
			suite.Fail("Id should be same")
		}
	})
	suite.Run("create with invalid data", func() {
		product := model.Product{
			Name:     faker.Commerce().ProductName(),
			Category: faker.Commerce().Department(),
			Price:    faker.Commerce().Price(),
		}
		//
		suite.mock.ExpectExec("INSERT INTO `products`").
			WillReturnError(errors.New("Title is exceed 32 characters"))
		_, err := suite.productRepo.Create(product)
		if err == nil {
			suite.Fail("Error should not nil")
		}
	})
}

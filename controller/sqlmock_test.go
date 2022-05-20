package controller_test

import (
	"regexp"
	"test-api/controller"
	"test-api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockDB struct {
	Mock sqlmock.Sqlmock
	Repo *controller.ProductController
}

func SetupMockDb() MockDB {
	var mockDB MockDB
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	openPostgres := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	testDb, err := gorm.Open(openPostgres, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	mockDB.Repo = controller.NewProductController(testDb)
	mockDB.Mock = mock
	return mockDB
}
func TestGetAllProducts_Success(t *testing.T) {
	mockDB := SetupMockDb()

	products := []model.Product{{ID: 1}}
	mockDB.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products"`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	res, err := mockDB.Repo.GetAllProducts()
	require.NoError(t, err)
	require.Equal(t, res, products)
}

func TestGetByID_Success(t *testing.T) {
	products := &model.Product{ID: 1}
	mockDB := SetupMockDb()

	mockDB.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products"`)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	res, err := mockDB.Repo.GetProductByID(1)

	require.NoError(t, err)
	require.Equal(t, res, products)
}
func TestCreateProduct(t *testing.T) {
	product := model.Product{Name: "test"}
	mockDB := SetupMockDb()
	mockDB.Mock.ExpectBegin()
	mockRows := sqlmock.NewRows([]string{"name"}).AddRow("test")
	mockDB.Mock.ExpectCommit()
	res := mockDB.Mock.ExpectQuery(
		regexp.QuoteMeta("INSERT INTO `products` (`name`) VALUES (?)")).WithArgs(product.Name).WillReturnRows(mockRows)

	err := mockDB.Repo.CreateProduct(&product)
	require.Equal(t, err, res)
}

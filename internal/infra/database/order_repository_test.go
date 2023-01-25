package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/virb30/goexpert-cleanarch/internal/entity"

	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db

}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Exec("DELETE FROM orders")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestShouldListOrders() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	anotherOrder, err := entity.NewOrder("456", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(anotherOrder.CalculateFinalPrice())
	err = repo.Save(anotherOrder)
	suite.NoError(err)

	orders, err := repo.GetAll()
	suite.NoError(err)

	suite.Len(orders, 2)
	suite.Equal(order.ID, orders[0].ID)
	suite.Equal(order.Price, orders[0].Price)
	suite.Equal(order.Tax, orders[0].Tax)
	suite.Equal(order.FinalPrice, orders[0].FinalPrice)
	suite.Equal(anotherOrder.ID, orders[1].ID)
	suite.Equal(anotherOrder.Price, orders[1].Price)
	suite.Equal(anotherOrder.Tax, orders[1].Tax)
	suite.Equal(anotherOrder.FinalPrice, orders[1].FinalPrice)
}

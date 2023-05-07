package database

import (
	"database/sql"

	"github.com/fabiofa8/pfa-go/internal/order/entity"
	_ "github.com/go-sql-driver/mysql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

// SQL create order table on mysql

// CREATE TABLE `orders` (
// 	`id` VARCHAR(36) NOT NULL,
// 	`price` FLOAT NOT NULL,
// 	`tax` FLOAT NOT NULL,
// 	`final_price` FLOAT NOT NULL,
// 	PRIMARY KEY (`id`)
// );

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	// BLANK IDENTIFIER - _ - is used to ignore the error return from Exec
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
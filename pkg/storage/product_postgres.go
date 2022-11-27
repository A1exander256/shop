package storage

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alexander256/shop/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProductStorage struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewProductStorage(db *sqlx.DB, log *logrus.Logger) *ProductStorage {
	return &ProductStorage{
		log: log,
		db:  db,
	}
}

func (s *ProductStorage) Create(product *models.Product) (int, error) {
	var id int
	tx, err := s.db.Begin()
	if err != nil {
		return id, err
	}
	createProductQuery := fmt.Sprintf("INSERT INTO %s (name, description, left_in_stock) VALUES ($1, $2, $3) RETURNING id", tableProducts)
	row := tx.QueryRow(createProductQuery, product.Name, product.Description, product.LeftInStock)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return id, err
	}

	for _, price := range product.Prices {
		createPricesQuery := fmt.Sprintf("INSERT INTO %s (product_id, currency, price) VALUES ($1, $2, $3)", tablePrices)
		if _, err := tx.Exec(createPricesQuery, id, price.Currency, price.Price); err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	return id, tx.Commit()
}

func (s *ProductStorage) Update(product *models.Product) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	setValuesProduct := make([]string, 0)
	argsProduct := make([]interface{}, 0)
	argIdProduct := 1

	if product.Name != "" {
		setValuesProduct = append(setValuesProduct, fmt.Sprintf("name=$%d", argIdProduct))
		argsProduct = append(argsProduct, product.Name)
		argIdProduct++
	}
	if product.Description != "" {
		setValuesProduct = append(setValuesProduct, fmt.Sprintf("description=$%d", argIdProduct))
		argsProduct = append(argsProduct, product.Description)
		argIdProduct++
	}
	if product.LeftInStock != 0 {
		setValuesProduct = append(setValuesProduct, fmt.Sprintf("left_in_stock=$%d", argIdProduct))
		argsProduct = append(argsProduct, product.LeftInStock)
		argIdProduct++
	}
	setProductQuery := strings.Join(setValuesProduct, ", ")
	updateProductQuery := fmt.Sprintf("UPDATE %s products SET %s WHERE products.id = $%d", tableProducts, setProductQuery, argIdProduct)
	argsProduct = append(argsProduct, product.Id)
	if _, err := tx.Exec(updateProductQuery, argsProduct...); err != nil {
		tx.Rollback()
		return err
	}

	for _, price := range product.Prices {
		var priceDB models.PriceProduct
		getPriceQuery := fmt.Sprintf("SELECT prices.id FROM %s prices WHERE prices.product_id = $1 AND prices.currency = $2", tablePrices)
		if err := s.db.Get(&priceDB, getPriceQuery, product.Id, price.Currency); err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}

		if priceDB.Id == 0 {
			createPriceQuery := fmt.Sprintf("INSERT INTO %s (product_id, currency, price) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM %s prices WHERE prices.product_id = $4 AND prices.currency = $5) RETURNING id", tablePrices, tablePrices)
			if _, err := tx.Exec(createPriceQuery, product.Id, price.Currency, price.Price, product.Id, price.Currency); err != nil {
				tx.Rollback()
				return err
			}
		} else {
			updatePriceQuery := fmt.Sprintf("UPDATE %s prices SET price = %f WHERE prices.product_id = $1 AND prices.currency = $2", tablePrices, price.Price)
			if _, err := tx.Exec(updatePriceQuery, product.Id, price.Currency); err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit()
}

func (s *ProductStorage) Delete(productId int) error {
	query := fmt.Sprintf("DELETE FROM %s products WHERE products.id=$1", tableProducts)
	_, err := s.db.Exec(query, productId)
	return err
}

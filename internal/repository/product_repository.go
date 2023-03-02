package repository

import (
	"database/sql"
	"merchant/internal/entity"
)

type ProductRepository struct {
	db *sql.DB
}

type IProductRepository interface {
	GetProductsByCustomerID(id int) (*sql.Rows, error)
	Update(id int, product entity.Product) error
	Delete(id int, product entity.Product) error
	Insert(id int, product entity.Product) error
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	var productRepo = ProductRepository{}
	productRepo.db = db
	return &productRepo
}

func (r *ProductRepository) GetProductsByCustomerID(id int) (*sql.Rows, error) {
	rows, err := r.db.Query(`select * from merchant_products where customer_id = $1`, id)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *ProductRepository) Update(id int, product entity.Product) error {
	_, err := r.db.Exec(`update merchant_products set name = $1, price = $2, quantity = $3 where customer_id = $4 and offer_id = $5`, product.Name, product.Price, product.Quantity, id, product.OfferID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Delete(id int, product entity.Product) error {
	_, err := r.db.Exec(`delete from merchant_products where customer_id = $1 and offer_id = $2`, id, product.OfferID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Insert(id int, product entity.Product) error {
	_, err := r.db.Exec(`insert into merchant_products values($1, $2, $3, $4, $5)`, id, product.OfferID, product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	}
	return nil
}

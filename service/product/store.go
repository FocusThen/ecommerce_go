package product

import (
	"database/sql"

	"github.com/focusthen/ecommerce_go/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Query("INSERT INTO products (name,description,image,price,quantity) VALUES (?,?,?,?,?)",
		product.Name,
		product.Description,
		product.Image,
		product.Price,
		product.Quantity)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)

	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&p.Quantity,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

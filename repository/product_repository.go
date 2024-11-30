package repository

import (
	"database/sql"
	"go-catalog/models"
	"strings"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) CreateProduct(product models.Product) (int64, error) {
	query := "INSERT INTO product (name, code, description, qty, active, deleted) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.DB.Exec(query, product.Name, product.Code, product.Description, product.Qty, product.Active, product.Deleted)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, name, code, description, qty, active, deleted FROM product WHERE deleted = false"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Code, &product.Description, &product.Qty, &product.Active, &product.Deleted); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(id int) (models.Product, error) {
	query := "SELECT id, name, code, description, qty, active, deleted FROM product WHERE id = ? AND deleted = false"
	row := r.DB.QueryRow(query, id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Code, &product.Description, &product.Qty, &product.Active, &product.Deleted)
	return product, err
}

func (r *ProductRepository) UpdateProduct(product models.Product) error {
	query := "UPDATE product SET name = ?, code = ?, description = ?, qty = ?, active = ? WHERE id = ? AND deleted = false"
	_, err := r.DB.Exec(query, product.Name, product.Code, product.Description, product.Qty, product.Active, product.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := "UPDATE product SET deleted = true WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *ProductRepository) GetProductByCodes(codes []string) ([]models.Product, error) {
	// Build placeholders for the query
	placeholders := make([]string, len(codes))
	args := make([]interface{}, len(codes))
	for i, code := range codes {
		placeholders[i] = "?"
		args[i] = code
	}

	// Create the SQL query
	query := "SELECT id, name, code, description, qty, active, deleted FROM product WHERE code IN (" + strings.Join(placeholders, ",") + ") AND deleted = false"

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Code, &product.Description, &product.Qty, &product.Active, &product.Deleted); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

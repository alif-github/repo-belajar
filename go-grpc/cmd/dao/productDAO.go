package dao

import (
	"database/sql"
	"fmt"
	productPb "github.com/alif-github/go-grpc/pb/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ProductDAO = productDAO{}.initiate()

type productDAO struct {
	FileName  string
	TableName string
}

func (input productDAO) initiate() productDAO {
	return productDAO{
		FileName:  "productDAO.go",
		TableName: "product",
	}
}

func (input productDAO) CountProducts(db *sql.DB) (total int, err error) {
	var (
		query string
	)

	query = fmt.Sprintf(`
		SELECT COUNT(p.id) 
		FROM %s p 
		INNER JOIN %s c on c.id = p.category_id `,
		input.TableName, CategoryDAO.TableName)
	err = db.QueryRow(query).Scan(&total)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return
	}

	return
}

func (input productDAO) GetListProducts(db *sql.DB, page int, limit int) (response []*productPb.Product, err error) {
	var (
		query    string
		rows     *sql.Rows
		products []*productPb.Product
	)

	query = fmt.Sprintf(`
		SELECT 
		p.id, p.name, p.price, 
		p.stock, c.id as category_id, c.name as category_name 
		FROM %s p 
		INNER JOIN %s c on c.id = p.category_id 
		LIMIT $1 OFFSET $2 `,
		input.TableName, CategoryDAO.TableName)

	offset := (page - 1) * limit
	params := []interface{}{limit, offset}
	rows, err = db.Query(query, params...)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return
	}

	defer rows.Close()
	for rows.Next() {
		var (
			product  productPb.Product
			category productPb.Category
		)

		err = rows.Scan(
			&product.Id, &product.Name, &product.Price,
			&product.Stock, &category.Id, &category.Name)

		if err != nil {
			err = status.Error(codes.Internal, err.Error())
			return
		}

		product.Category = &category
		products = append(products, &product)
	}

	response = products
	return
}

func (input productDAO) GetProductByID(db *sql.DB, id int) (result *productPb.Product, err error) {
	var (
		product  productPb.Product
		category productPb.Category
	)

	query := fmt.Sprintf(`
		SELECT 
		p.id, p.name, p.price, 
		p.stock, c.id as category_id, c.name as category_name 
		FROM %s p 
		INNER JOIN %s c on c.id = p.category_id 
		WHERE p.id = $1 `,
		input.TableName, CategoryDAO.TableName)

	param := []interface{}{id}
	row := db.QueryRow(query, param...)
	err = row.Scan(
		&product.Id, &product.Name, &product.Price,
		&product.Stock, &category.Id, &category.Name)
	if err != nil && err != sql.ErrNoRows {
		err = status.Error(codes.Internal, err.Error())
		return
	}

	product.Category = &category
	result = &product
	return
}

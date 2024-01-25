package dao

import (
	"database/sql"
	"fmt"
	productPb "github.com/alif-github/go-grpc/pb/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var CategoryDAO = categoryDAO{}.initiate()

type categoryDAO struct {
	FileName  string
	TableName string
}

func (input categoryDAO) initiate() categoryDAO {
	return categoryDAO{
		FileName:  "categoryDAO.go",
		TableName: "category",
	}
}

func (input categoryDAO) GetCategoryByName(db *sql.DB, name string) (category productPb.Category, err error) {
	query := fmt.Sprintf(`
		SELECT id, name 
		FROM %s WHERE LOWER(name) = $1`,
		input.TableName)

	param := []interface{}{name}
	row := db.QueryRow(query, param...)
	err = row.Scan(&category.Id, &category.Name)
	if err != nil && err != sql.ErrNoRows {
		err = status.Error(codes.Internal, err.Error())
		return
	}

	return
}

func (input categoryDAO) InsertCategory(tx *sql.Tx, newCategory *productPb.Category) (err error) {
	query := fmt.Sprintf(`
		INSERT INTO %s(id, name) 
		VALUES (nextval(), $1) 
		RETURNING id`,
		input.TableName)

	param := []interface{}{newCategory.GetName()}
	row := tx.QueryRow(query, param...)
	err = row.Scan(&newCategory.Id)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return
	}

	return
}

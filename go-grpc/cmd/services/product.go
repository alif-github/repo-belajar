package services

import (
	"context"
	"database/sql"
	"github.com/alif-github/go-grpc/cmd/config"
	"github.com/alif-github/go-grpc/cmd/dao"
	"github.com/alif-github/go-grpc/pb/pagination"
	productPb "github.com/alif-github/go-grpc/pb/product"
	"log"
	"math"
	"sync"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	mu sync.Mutex
}

func (d *ProductService) GetProducts(ctx context.Context, pageParam *productPb.Page) (products *productPb.Products, err error) {
	var (
		dataProducts []*productPb.Product
		total        int
		page         = 1
		limit        = 5
	)

	if pageParam.GetPage() != 0 {
		page = int(pageParam.GetPage())
	}

	log.Println("Incoming request Get Products...")
	defer func() {
		if err != nil {
			log.Fatalf("Error get data -> %v", err.Error())
		}
	}()

	//--- Get List Products
	dataProducts, err = dao.ProductDAO.GetListProducts(config.ServerAttribute.DBConnection, page, limit)
	if err != nil {
		return
	}

	//--- Count Products
	total, err = dao.ProductDAO.CountProducts(config.ServerAttribute.DBConnection)
	if err != nil {
		return
	}

	products = &productPb.Products{
		Pagination: &pagination.Pagination{
			Total:       uint64(total),
			PerPage:     uint32(limit),
			CurrentPage: uint32(page),
			LastPage:    uint32(math.Ceil(float64(total) / float64(limit))),
		},
		Data: dataProducts,
	}

	return
}

func (d *ProductService) GetProduct(ctx context.Context, productParam *productPb.Product) (result *productPb.Product, err error) {
	var id int
	log.Println("Incoming request Get Product...")
	if productParam.GetId() < 1 {
		return
	}

	id = int(productParam.GetId())
	result, err = dao.ProductDAO.GetProductByID(config.ServerAttribute.DBConnection, id)
	if err != nil {
		return
	}

	return
}

func (d *ProductService) CreateProduct(ctx context.Context, productParam *productPb.Product) (result *productPb.Product, err error) {
	var (
		product  productPb.Product
		category productPb.Category
		tx       *sql.Tx
		db       = config.ServerAttribute.DBConnection
	)

	log.Println("Incoming request Create Product...")
	product = productPb.Product{
		Name:     productParam.GetName(),
		Price:    productParam.GetPrice(),
		Stock:    productParam.GetStock(),
		Category: nil,
	}

	category = productPb.Category{
		Name: productParam.Category.GetName(),
	}

	tx, err = db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	//--- Get Category By Name
	category, err = dao.CategoryDAO.GetCategoryByName(config.ServerAttribute.DBConnection, category.GetName())
	if err != nil {
		return
	}

	//--- Insert to DB Category
	if category.GetId() < 1 {
		err = dao.CategoryDAO.InsertCategory(tx, &category)
		if err != nil {
			return
		}
	}

	//--- Add to Product
	product.Category = &category
	return
}

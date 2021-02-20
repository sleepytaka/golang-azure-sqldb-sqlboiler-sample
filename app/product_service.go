package app

import (
	"context"
	"database/sql"
	"errors"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang-azure-sqldb-sqlboiler/common/models"
	"gopkg.in/jeevatkm/go-model.v1"
)

var productBlacklist = []string {"ThumbNailPhoto", "rowguid", "ModifiedDate"}

type ProductService interface {
	create(req *ProductRequest)(*ProductResponse, error)
	update(id int, req *ProductRequest)(*ProductResponse, error)
	delete(id int) error
	findById(id int)(*ProductResponse, error)
	search(categoryName string) (*[]ProductResponse, error)
}

type productService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) ProductService {
	return &productService{db}
}

func (svc *productService) create(req *ProductRequest) (*ProductResponse, error) {
	prod, err := modelMapper(req)
	if err != nil {
		return nil, err
	}

	if err := prod.Insert(context.Background(), svc.db, boil.Blacklist(productBlacklist...));err != nil {
		return nil, err
	}

	return svc.findById(prod.ProductID)
}

func (svc *productService) update(id int, req *ProductRequest) (*ProductResponse, error) {
	prod, err := modelMapper(req)
	if err != nil {
		return nil, err
	}

	prod.ProductID = id
	if _, err := prod.Update(context.Background(), svc.db, boil.Blacklist(productBlacklist...)); err != nil {
		return nil, err
	}

	return svc.findById(id)
}

func (svc *productService) delete(id int) error {
	prod, err := models.FindProduct(context.Background(), svc.db, id)
	if err != nil {
		return err
	}

	if _, err := prod.Delete(context.Background(), svc.db); err != nil {
		return err
	}

	return nil
}
func (svc *productService) findById(id int) (*ProductResponse, error) {
	prod, err := models.FindProduct(context.Background(), svc.db, id)
	if err != nil {
		return nil, err
	}

	dst := &ProductResponse{}
	if errs := model.Copy(dst, prod); errs != nil {
		return nil, errors.New("copy model error")
	}

	uid := mssql.UniqueIdentifier{}
	if err := uid.Scan([]byte(prod.Rowguid)); err != nil {
		return nil, err
	}
	dst.RowGuid = uid.String()
	dst.Weight = prod.Weight
	return dst, nil
}

func (svc *productService) search(categoryName string) (*[]ProductResponse, error) {
	prods, err := models.Products(
		qm.Select("Product.*"),
		qm.InnerJoin("SalesLT.ProductCategory AS c on Product.ProductCategoryID = c.ProductCategoryID"),
		qm.Where("c.Name = ?", categoryName),
	).All(context.Background(), svc.db)
	if err != nil {
		return nil, err
	}

	size := len(prods)
	dst := make([]ProductResponse,size)
	for i := 0; i < size; i++ {
		prod := prods[i]
		if errs := model.Copy(&dst[i], prod); errs != nil {
			return nil, errors.New("copy model error")
		}

		uid := mssql.UniqueIdentifier{}
		if err := uid.Scan([]byte(prods[i].Rowguid)); err != nil {
			return nil, err
		}
		dst[i].RowGuid = uid.String()
		dst[i].Weight = prods[i].Weight
	}

	return &dst, nil
}

func modelMapper(src *ProductRequest) (*models.Product, error) {
	dst := &models.Product{}
	if errs := model.Copy(dst, src); errs != nil {
		return nil, errors.New("copy model error")
	}

	dst.Weight = src.Weight

	return dst, nil
}

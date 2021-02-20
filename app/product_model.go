package app

import (
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
	"time"
)

type ProductFindParam struct {
	Id						int            `uri:"id" binding:"required,min=1"`
}
type ProductSearchParam struct {
	CategoryName			string         `query:"categoryName" binding:"required"`
}

type ProductRequest struct {
	Name                   string            `json:"name" binding:"required"`
	ProductNumber          string            `json:"productNumber" binding:"required"`
	Color                  null.String       `json:"color"`
	StandardCost           string            `json:"standardCost" binding:"required,money"`
	ListPrice              string            `json:"listPrice" binding:"required"`
	Size                   null.String       `json:"size"`
	Weight                 types.NullDecimal `json:"weight"`
	ProductCategoryID      null.Int          `json:"productCategoryID"`
	ProductModelID         null.Int          `json:"productModelID"`
	SellStartDate          time.Time         `json:"sellStartDate" binding:"required"`
	SellEndDate            null.Time         `json:"sellEndDate"`
	DiscontinuedDate       null.Time         `json:"discontinuedDate"`
	ThumbnailPhotoFileName null.String       `json:"thumbnailPhotoFileName"`
}

type ProductResponse struct {
	ProductID              int               `json:"productID"`
	Name                   string            `json:"name"`
	ProductNumber          string            `json:"productNumber"`
	Color                  null.String       `json:"color"`
	StandardCost           string            `json:"standardCost"`
	ListPrice              string            `json:"listPrice"`
	Size                   null.String       `json:"size"`
	Weight                 types.NullDecimal `json:"weight"`
	ProductCategoryID      null.Int          `json:"productCategoryID"`
	ProductModelID         null.Int          `json:"productModelID"`
	SellStartDate          time.Time         `json:"sellStartDate"`
	SellEndDate            null.Time         `json:"sellEndDate"`
	DiscontinuedDate       null.Time         `json:"discontinuedDate"`
	ThumbnailPhotoFileName null.String       `json:"thumbnailPhotoFileName"`
	RowGuid                string            `json:"rowGuid"`
	ModifiedDate           time.Time         `json:"modifiedDate"`
}
package app

import (
	_ "encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindById_Success(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()


	rows := sqlmock.NewRows([]string{
		"ProductID",
		"Name",
		"ProductNumber",
		"Color",
		"StandardCost",
		"ListPrice",
		"Size",
		"Weight",
		"ProductCategoryID",
		"ProductModelID",
		"SellStartDate",
		"SellEndDate",
		"DiscontinuedDate",
		"ThumbNailPhoto",
		"ThumbnailPhotoFileName",
		"rowguid",
		"ModifiedDate",
	})

	tm, _ := time.Parse("2006-01-02", "2020-01-10")
	rows.AddRow(
		680,
		"HL Road Frame - Black, 58",
		"FR-R92B-58",
		"Black",
		1059.31,
		1431.50,
		58,
		1016.04,
		18,
		6,
		tm,
		nil,
		nil,
		nil,
		"no_image_available_small.gif",
		"43DD68D6-14A4-461F-9069-55309D90EA7E",
		tm)

	// 実行されるクエリ―を定義
	productID := 680
	q1 := "select * from [SalesLT].[Product] where [ProductID]=$1"
	mock.ExpectQuery(q1).
		WithArgs(productID).
		WillReturnRows(rows)


	svc := NewProductService(db)
	asserts := assert.New(t)

	prod, _ := svc.findById(680)
	asserts.Equal(680, prod.ProductID)
	asserts.Equal("HL Road Frame - Black, 58", prod.Name)
}


package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	create(ctx *gin.Context)
	update(ctx *gin.Context)
	delete(ctx *gin.Context)
	findById(ctx *gin.Context)
	search(ctx *gin.Context)
}

type productController struct {
	db *sql.DB
	svc ProductService
}

func NewProductController(db *sql.DB, svc ProductService) ProductController {
	return &productController{db, svc}
}

func (ctrl *productController) create(ctx *gin.Context) {
	var req ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.svc.create(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) update(ctx *gin.Context) {
	var req ProductRequest
	var prm ProductFindParam
	if err := ctx.ShouldBindUri(&prm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.svc.update(prm.Id, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) delete(ctx *gin.Context) {
	var prm ProductFindParam
	if err := ctx.ShouldBindUri(&prm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.svc.delete(prm.Id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (ctrl *productController) findById(ctx *gin.Context) {
	var prm ProductFindParam
	if err := ctx.ShouldBindUri(&prm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.svc.findById(prm.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) search(ctx *gin.Context) {
	var prm ProductSearchParam
	if err := ctx.ShouldBindQuery(&prm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ctrl.svc.search(prm.CategoryName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
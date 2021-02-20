package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang-azure-sqldb-sqlboiler/common/validations"
	"log"
)

type Handler interface {
	Run() error
}

type (
	handler struct {
		router   *gin.Engine
	}
)

// New 初期化
func NewHandler(engine *gin.Engine, pc ProductController) Handler {
	h := &handler{
		router : engine,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("money", validations.MoneyValidation); err != nil {
			log.Fatal(err)
		}
	}

	h.router.GET("/products/:id", pc.findById)
	h.router.GET("/products/", pc.search)
	h.router.POST("/products/", pc.create)
	h.router.PUT("/products/:id", pc.update)
	h.router.DELETE("/products/:id", pc.delete)
	return h
}

// サーバー起動
func (h *handler) Run() error {
	return h.router.Run(":3000")
}

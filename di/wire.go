// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"golang-azure-sqldb-sqlboiler/app"
	"golang-azure-sqldb-sqlboiler/common/database"
)

func InitializeHandler(conn string) (app.Handler, func(), error) {
	wire.Build(
		app.NewHandler,
		database.NewDB,
		gin.Default,
		app.NewProductService,
		app.NewProductController,
	)
	return nil, nil, nil
}

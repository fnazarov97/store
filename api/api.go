package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, cache storage.CacheStorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, cache, logger)

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customMiddleware())
	//AUTH
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	//USER
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	//CATEGORY
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.AuthMiddleware(), handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.PATCH("/category/:id", handler.UpdatePatchCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	//BRAND
	r.POST("/brand", handler.CreateBrand)
	r.GET("/brand/:id", handler.AuthMiddleware(), handler.GetByIdBrand)
	r.GET("/brand", handler.AuthMiddleware(), handler.GetListBrand)
	r.PUT("/brand/:id", handler.UpdateBrand)
	r.PATCH("/brand/:id", handler.UpdatePatchBrand)
	r.DELETE("/brand/:id", handler.DeleteBrand)

	//PRODUCT
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.PATCH("/product/:id", handler.UpdatePatchProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	//CUSTOMER
	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.PATCH("/customer/:id", handler.UpdatePatchCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	//STORE
	r.POST("/store", handler.CreateStore)
	r.GET("/store/:id", handler.GetByIdStore)
	r.GET("/store", handler.GetListStore)
	r.PUT("/store/:id", handler.UpdateStore)
	r.PATCH("/store/:id", handler.UpdatePatchStore)
	r.DELETE("/store/:id", handler.DeleteStore)

	//STAFF
	r.POST("/staff", handler.CreateStaff)
	r.GET("/staff/:id", handler.GetByIdStaff)
	r.GET("/staff", handler.GetListStaff)
	r.PUT("/staff/:id", handler.UpdateStaff)
	r.PATCH("/staff/:id", handler.UpdatePatchStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)

	//ORDER
	r.POST("/order", handler.CreateOrder)
	r.GET("/order/:id", handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrder)
	r.PUT("/order/:id", handler.UpdateOrder)
	r.PATCH("/order/:id", handler.UpdatePatchOrder)
	r.DELETE("/order/:id", handler.DeleteOrder)
	r.POST("/order_item", handler.CreateOrderItem)
	r.DELETE("/order_item/:id", handler.DeleteOrderItem)

	//STOCK
	r.POST("/stock", handler.CreateStock)
	r.GET("/stock/:id", handler.GetByIdStock)
	r.GET("/stock", handler.GetListStock)
	r.PUT("/stock/:id", handler.UpdateStock)
	r.PATCH("/stock/:id", handler.UpdatePatchStock)
	r.DELETE("/stock/:id", handler.DeleteStock)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

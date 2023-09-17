package main

import (
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	//"GoGoGo/cmd/server/config"
	"GoGoGo/cmd/server/external/database"
	"GoGoGo/cmd/server/handler"

	//"GoGoGo/cmd/server/middlewares"
	"GoGoGo/internal/dentists"
	//"net/http"
	//"os"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/swag/example/basic/docs"
)

// @title Certified Tech Developer
// @version 1.0
// @description This API Handle Products.
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	fmt.Println("hola database")

	db, err := database.Init()
	fmt.Println("sql Open")
	if err != nil {
		panic(err)
	}
	fmt.Println("hola database")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("hola database")

	router := gin.Default()

	router.GET("/belu", func(c *gin.Context) {
		c.JSON(200, gin.H{"massage": "Hola Belu"})
	})

	myDatabase := database.NewDatabase(db)

	dentistsService := dentists.NewService(myDatabase)

	dentistsHandler := handler.NewDentistsHandler(dentistsService, dentistsService, dentistsService, dentistsService)

	dentistGroup := router.Group("/dentists")

	dentistGroup.POST("/", dentistsHandler.PostDentist)
	dentistGroup.GET("/:id", dentistsHandler.GetDentistByID)
	dentistGroup.PUT("/:id", dentistsHandler.PutDentist)
	dentistGroup.PATCH("/:id", dentistsHandler.PatchDentist)
	dentistGroup.DELETE("/:id", dentistsHandler.DeleteDentist)

	err = router.Run()

	if err != nil {
		panic(err)
	}

	/* 	env := os.Getenv("ENV")
	   	if env == "" {
	   		env = "local"
	   	}

	   	if env == "local" {
	   		err := godotenv.Load()
	   		if err != nil {
	   			panic(err)
	   		}
	   	}

	   	cfg, err := config.NewConfig(env)

	   	if err != nil {
	   		panic(err)
	   	}

	   	authMidd := middlewares.NewAuth(cfg.PublicConfig.PublicKey, cfg.PrivateConfig.SecretKey)

	   	router := gin.New()

	   	customRecovery := gin.CustomRecovery(middlewares.RecoveryWithLog)

	   	router.Use(customRecovery)

	   	// docs endpoint
	   	docs.SwaggerInfo.Host = os.Getenv("HOST")
	   	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	   	router.GET("/ping", func(context *gin.Context) {
	   		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	   	})

	   	postgresDatabase, err := database.NewPostgresSQLDatabase(cfg.PublicConfig.PostgresHost,
	   		cfg.PublicConfig.PostgresPort, cfg.PublicConfig.PostgresUser, cfg.PrivateConfig.PostgresPassword,
	   		cfg.PublicConfig.PostgresDatabase)

	   	if err != nil {
	   		panic(err)
	   	}


	   	myDatabase := database.NewDatabase(postgresDatabase)

	   	productsService := products.NewService(myDatabase)

	   	productsHandler := handler.NewProductsHandler(productsService, productsService)

	   	productsGroup := router.Group("/products")

	   	productsGroup.GET("/:id", productsHandler.GetProductByID)

	   	productsGroup.PUT("/:id", authMidd.AuthHeader, productsHandler.PutProduct)

	   	err = router.Run()

	   	if err != nil {
	   		panic(err)
	   	}
	*/
}

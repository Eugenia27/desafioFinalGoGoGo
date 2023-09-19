package main

import (
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"GoGoGo/cmd/server/config"
	"GoGoGo/cmd/server/external/database"
	"GoGoGo/cmd/server/handler"
	"GoGoGo/cmd/server/middlewares"
	"GoGoGo/internal/appointments"
	"GoGoGo/internal/dentists"
	"GoGoGo/internal/patients"
	"fmt"

	//"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/swag/example/basic/docs"
)

func main() {

	godotenv.Load()
	env := os.Getenv("ENV")
	fmt.Println("env load ----->" + env)
	if env == "" {
		env = "local"
	}

	db, err := database.Init()
	fmt.Println("sql Open")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	cfg, err := config.NewConfig(env)
	if err != nil {
		panic(err)
	}

	authMidd := middlewares.NewAuth(cfg.PublicConfig.PublicKey) //, cfg.PrivateConfig.SecretKey)

	router := gin.Default()

	dentistRepository := database.NewDentistRepository(db)
	patientRepository := database.NewPatientRepository(db)
	appointmentRepository := database.NewAppointmentRepository(db)

	dentistsService := dentists.NewService(dentistRepository)
	dentistsHandler := handler.NewDentistsHandler(dentistsService, dentistsService, dentistsService, dentistsService)
	dentistGroup := router.Group("/dentists")
	dentistGroup.POST("/", authMidd.AuthHeader, dentistsHandler.PostDentist)
	dentistGroup.GET("/:id", dentistsHandler.GetDentistByID)
	dentistGroup.PUT("/:id", authMidd.AuthHeader, dentistsHandler.PutDentist)
	dentistGroup.PATCH("/:id", authMidd.AuthHeader, dentistsHandler.PatchDentist)
	dentistGroup.DELETE("/:id", authMidd.AuthHeader, dentistsHandler.DeleteDentist)

	patientsService := patients.NewService(patientRepository)
	patientsHandler := handler.NewPatientsHandler(patientsService, patientsService, patientsService, patientsService)
	patientGroup := router.Group("/patients")
	patientGroup.POST("/", authMidd.AuthHeader, patientsHandler.PostPatient)
	patientGroup.GET("/:id", patientsHandler.GetPatientByID)
	patientGroup.PUT("/:id", authMidd.AuthHeader, patientsHandler.PutPatient)
	patientGroup.PATCH("/:id", authMidd.AuthHeader, patientsHandler.PatchPatient)
	patientGroup.DELETE("/:id", authMidd.AuthHeader, patientsHandler.DeletePatient)

	appointmentsService := appointments.NewService(appointmentRepository)
	appointmentHandler := handler.NewAppointmentsHandler(appointmentsService, appointmentsService, appointmentsService, appointmentsService, appointmentsService)
	appointmentGroup := router.Group("/appointments")
	appointmentGroup.POST("/", authMidd.AuthHeader, appointmentHandler.PostAppointment)
	appointmentGroup.GET("/:id", appointmentHandler.GetAppointmentByID)
	appointmentGroup.PUT("/:id", authMidd.AuthHeader, appointmentHandler.PutAppointment)
	appointmentGroup.PATCH("/:id", authMidd.AuthHeader, appointmentHandler.PatchAppointment)
	appointmentGroup.DELETE("/:id", authMidd.AuthHeader, appointmentHandler.DeleteAppointment)
	appointmentGroup.GET("/by_patients/:credential_id", appointmentHandler.GetAppointmentByPatient)

	err = router.Run()
	if err != nil {
		panic(err)
	}

	/*
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

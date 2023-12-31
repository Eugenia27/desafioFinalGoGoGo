package main

import (
	"GoGoGo/cmd/server/config"
	"GoGoGo/cmd/server/external/database"
	"GoGoGo/cmd/server/handler"
	"GoGoGo/cmd/server/middlewares"
	"GoGoGo/docs"
	"GoGoGo/internal/appointments"
	"GoGoGo/internal/dentists"
	"GoGoGo/internal/patients"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Clinica Odotntologica Dientes GoGoGo
// @version 1.0
// @description This API Handle Clinica.
// @termsOfService https://developers.gogogo.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.gogogo.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	godotenv.Load()
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	db, err := database.Init()
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

	authMidd := middlewares.NewAuth(cfg.PublicConfig.PublicKey)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	if docs.SwaggerInfo.Host == "" {
		docs.SwaggerInfo.Host = "localhost:8080"
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

}

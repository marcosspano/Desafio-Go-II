package main

import (
	"desafioFinal/cmd/server/handler"
	"desafioFinal/internal/appointment"
	"desafioFinal/internal/dentist"
	"desafioFinal/internal/patient"
	"desafioFinal/pkg/store"
	"log"

	"time"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file", err.Error())
	}

}

func main() {

	sqlStore := store.NewSQLStore()
	apStore := store.NewSQLAp()

	appRepo := appointment.NewRepository(apStore)
	appService := appointment.NewService(appRepo)
	appHandler := handler.NewAppointmentHandler(appService)

	dentistRepo := dentist.NewRepository(sqlStore)
	dentistService := dentist.NewService(dentistRepo)
	dentistHandler := handler.NewDentistHandler(dentistService)

	patientRepo := patient.NewRepository(sqlStore)
	patientService := patient.NewService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientService)

	r := gin.Default()

	api := r.Group("/api")
	{
		appointments := api.Group("/appointments")
		{
			appointments.GET("", appHandler.GetAll())
			appointments.GET(":id", appHandler.GetByID())
			appointments.GET("/patient/:identity_number", appHandler.GetAllByIdentityNumber())
			appointments.GET("/dentist/:license_number", appHandler.GetAllByLicenseNumber())
			appointments.POST("", appHandler.Post())
			appointments.PUT(":id", appHandler.Put())
			appointments.PATCH(":id", appHandler.Patch())
			appointments.DELETE(":id", appHandler.Delete())
		}
		dentists := api.Group("/dentists")
		{
			dentists.GET("", dentistHandler.GetAll())
			dentists.GET(":id", dentistHandler.GetByID())
			dentists.POST("", dentistHandler.Post())
			dentists.PUT(":id", dentistHandler.Put())
			dentists.PATCH(":id", dentistHandler.Patch())
			dentists.DELETE(":id", dentistHandler.Delete())
		}
		patients := api.Group("/patients")
		{
			patients.GET("", patientHandler.GetAll())
			patients.GET(":id", patientHandler.GetByID())
			patients.POST("", patientHandler.Post())
			patients.PUT(":id", patientHandler.Put())
			patients.PATCH(":id", patientHandler.Patch())
			patients.DELETE(":id", patientHandler.Delete())
		}
	}

	r.Run(":8080")

}

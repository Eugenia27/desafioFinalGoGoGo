package handler

import (
	"GoGoGo/internal/patients"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientsGetter interface {
	GetByID(id int) (patients.Patient, error)
}

// type PatientCreator interface {
// 	ModifyByID(id int, dentist patients.Patient) (patients.Patient, error)
// }

type PatientsHandler struct {
	patientsGetter PatientsGetter
	//patientsCreator PatientCreator
}

func NewPatientsHandler(getter PatientsGetter) *PatientsHandler {
	return &PatientsHandler{
		patientsGetter: getter,
		//patientsCreator: creator,
	}
}

// func NewPatientsHandler(getter PatientGetter, creator PatientCreator) *PatientsHandler {
// 	return &PatientsHandler{
// 		patientsGetter:  getter,
// 		//patientsCreator: creator,
// 	}
// }

// GetPatientByID godoc
// @Summary      Gets a dentist by id
// @Description  Gets a dentist by id from the repository
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} patients.Patient
// @Router       /patients/{id} [get]
func (ph *PatientsHandler) GetPatientByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	patient, err := ph.patientsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

// func (ph *PatientsHandler) PutPatient(ctx *gin.Context) {
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.JSON(400, gin.H{"error": "invalid id"})
// 		return
// 	}
// 	dentistRequest := patients.Patient{}
// 	err = ctx.BindJSON(&dentistRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	dentist, err := ph.patientsCreator.ModifyByID(id, dentistRequest)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": "internal error"})
// 		return
// 	}
// 	ctx.JSON(200, dentist)
// }

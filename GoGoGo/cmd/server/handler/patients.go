package handler

import (
	"GoGoGo/internal/patients"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientCreator interface {
	Save(patient patients.Patient) (patients.Patient, error)
}

type PatientsGetter interface {
	GetByID(id int) (patients.Patient, error)
}

type PatientUpdate interface {
	ModifyByID(id int, patient patients.Patient) (patients.Patient, error)
}

type PatientDelete interface {
	Delete(id int) error
}

type PatientsHandler struct {
	patientsCreator PatientCreator
	patientsGetter  PatientsGetter
	patientsUpdate  PatientUpdate
	patientsDelete  PatientDelete
}

func NewPatientsHandler(creator PatientCreator, getter PatientsGetter, update PatientUpdate, delete PatientDelete) *PatientsHandler {
	return &PatientsHandler{
		patientsCreator: creator,
		patientsGetter:  getter,
		patientsUpdate:  update,
		patientsDelete:  delete,
	}
}

// GetPatientByID godoc
// @Summary      Gets a patient by id
// @Description  Gets a patient by id from the repository
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} patients.Patient
// @Router       /patients/{id} [get]

func (ph *PatientsHandler) PostPatient(ctx *gin.Context) {
	patientRequest := patients.Patient{}
	err := ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(patientRequest)

	patient, err := ph.patientsCreator.Save(patientRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusCreated, patient)
}

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

func (ph *PatientsHandler) PutPatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	patientRequest := patients.Patient{}
	err = ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := ph.patientsUpdate.ModifyByID(id, patientRequest)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

func (ph *PatientsHandler) PatchPatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	patientRequest := patients.Patient{}
	err = ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patientDB, err := ph.patientsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	if patientRequest.FirstName != "" {
		patientDB.FirstName = patientRequest.FirstName
	}
	if patientRequest.LastName != "" {
		patientDB.LastName = patientRequest.LastName
	}
	if patientRequest.Address != "" {
		patientDB.Address = patientRequest.Address
	}
	if patientRequest.CredentialID != "" {
		patientDB.CredentialID = patientRequest.CredentialID
	}
	if patientRequest.DischargeDate != "" {
		patientDB.DischargeDate = patientRequest.DischargeDate
	}

	patientUpdate, err := ph.patientsUpdate.ModifyByID(id, patientDB)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, patientUpdate)
}

func (ph *PatientsHandler) DeletePatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = ph.patientsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	deleted := ph.patientsDelete.Delete(id)
	if deleted != nil {
		ctx.JSON(404, gin.H{"error": "delete failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "patient deleted"})
}

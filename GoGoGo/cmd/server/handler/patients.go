package handler

import (
	"GoGoGo/internal/patients"
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

// PostPatient godoc
// @Summary      Guarda nuevo paciente
// @Description  Guarda nuevo paciente en la base de datos
// @Tags         patients
// @Produce      json
// @Body         {first_name:"first_name", last_name:"last_name", address:"calle altura", credential_id:"1234", discharge_date:"AAAA-MM-DD"}
// @Success      201 {object} patients.Patient
// @Failure      500 Error interno
// @Router       /patients [post]
func (ph *PatientsHandler) PostPatient(ctx *gin.Context) {
	patientRequest := patients.Patient{}
	err := ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient, err := ph.patientsCreator.Save(patientRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusCreated, patient)
}

// GetPatientByID godoc
// @Summary      Obtiene un paciente por id
// @Description  Obtiene un paciente por id desde la base de datos
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} patients.Patient
// @Failure      404 Paciente no encontrado, ID incorrecto
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

// PutPatient godoc
// @Summary      Modifica un paciente existente
// @Description  Modifica un paciente existente por id y lo almacena en la base de datos
// @Tags         patients
// @Produce      json
// @Body         {first_name:"first_name", last_name:"last_name", address:"calle altura", credential_id:"1234", discharge_date:"AAAA-MM-DD"}
// @Success      201 {object} patients.Patient
// @Failure      400 Error caracter inválido
// @Failure      500 Error interno
// @Router       /patients/{id} [put]
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
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

// PatchPatient godoc
// @Summary      Modifica uno o más atributos de un paciente existente
// @Description  Modifica uno o más atributos de un paciente existente y almacena los cambios en la base de datos
// @Tags         patients
// @Produce      json
// @Body         {[first_name:"first_name", last_name:"last_name"], [address:"calle altura"], [credential_id:"1234"], [discharge_date:"AAAA-MM-DD"]}
// @Success      201 {object} patients.Patient
// @Failure      400 Paciente no encontrado, ID incorrecto
// @Failure      500 Error interno
// @Router       /patients/{id} [patch]
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

// DeletePatient godoc
// @Summary      Elimina un paciente por id
// @Description  Elimina un paciente por id en la base de datos
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 Paciente eliminado
// @Failure      404 Paciente no encontrado, ID incorrecto
// @Router       /patients/{id} [delete]
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

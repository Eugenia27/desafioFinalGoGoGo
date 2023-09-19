package handler

import (
	"GoGoGo/internal/appointments"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentCreator interface {
	Save(appointment appointments.Appointment) (appointments.Appointment, error)
}
type AppointmentGetter interface {
	GetByID(id int) (appointments.Appointment, error)
}

type AppointmentGetterByPatient interface {
	GetByCredentialID(credentialId string) ([]appointments.AppointmentDTO, error)
}

type AppointmentUpdate interface {
	ModifyByID(id int, appointment appointments.Appointment) (appointments.Appointment, error)
}

type AppointmentDelete interface {
	Delete(id int) error
}

type AppointmentsHandler struct {
	appointmentsCreator         AppointmentCreator
	appointmentsGetter          AppointmentGetter
	appointmentsGetterByPatient AppointmentGetterByPatient
	appointmentsUpdate          AppointmentUpdate
	appointmentsDelete          AppointmentDelete
}

func NewAppointmentsHandler(creator AppointmentCreator, getter AppointmentGetter, getterByPatient AppointmentGetterByPatient, update AppointmentUpdate, delete AppointmentDelete) *AppointmentsHandler {
	return &AppointmentsHandler{
		appointmentsGetter:          getter,
		appointmentsGetterByPatient: getterByPatient,
		appointmentsCreator:         creator,
		appointmentsDelete:          delete,
		appointmentsUpdate:          update,
	}
}

// PostAppointment
// @Summary      Crea un nuevo turno
// @Description  Crea un nuevo turno y lo almacena en la base de datos.
// @Tags         appointments
// @Produce      json
// @Param 		 body body appointments.Appointment true "Appointment"
// @Success      201 {object} appointments.Appointment
// @Failure		 400 {object} error
// @Failure		 500 {object} error: internal error
// @Router       /appointments [post]
func (ph *AppointmentsHandler) PostAppointment(ctx *gin.Context) {

	appointmentRequest := appointments.Appointment{}
	err := ctx.BindJSON(&appointmentRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(appointmentRequest)

	appointment, err := ph.appointmentsCreator.Save(appointmentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, appointment)
}

// GetAppointmentByID
// @Summary      Obtener un turno por su ID
// @Description  Obtener un turno por su ID existente en la base de datos.
// @Tags         appointments
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} appointments.Appointment
// @Failure		 400 {object} error: invalid id
// @Failure	     404 {object} error: appointment not found
// @Router       /appointments/{id} [get]
func (ph *AppointmentsHandler) GetAppointmentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	appointment, err := ph.appointmentsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}
	ctx.JSON(http.StatusOK, appointment)
}

// PutAppointment
// @Summary      Modificar un turno
// @Description  Modifica la información de un turno existente en la base de datos. Este endpoint permite actualizar los detalles de un odontólogo identificado por su ID.
// @Tags         appointments
// @Produce      json
// @Param        id path string true "ID"
// @Param 		 body body appointments.Appointment true "Appointment"
// @Success      200 {object} appointments.Appointment
// @Failure		 400 {object} error: invalid id
// @Failure		 500 {object} error: internal error
// @Router       /appointments/{id} [put]
func (ph *AppointmentsHandler) PutAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	appointmentRequest := appointments.Appointment{}
	err = ctx.BindJSON(&appointmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment, err := ph.appointmentsUpdate.ModifyByID(id, appointmentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, appointment)
}

// PatchAppointment
// @Summary      Modificar turno
// @Description  Modifica la información de un turno existente en la base de datos por alguno de sus campos. Los campos no proporcionados en la solicitud permanecerán sin cambios en el odontólogo.
// @Tags         appointments
// @Produce      json
// @Param        id path string true "ID"
// @Param 		 body body appointments.Appointment true "Appointment"
// @Success      200 {object} appointments.Appointment
// @Failure		 400 {object} error: invalid id
// @Failure		 404 {object} error: dentist not found
// @Failure		 500 {object} error: internal error
// @Router       /appointments/{id} [patch]
func (ph *AppointmentsHandler) PatchAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	appointmentRequest := appointments.Appointment{}
	err = ctx.BindJSON(&appointmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointmentDB, err := ph.appointmentsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	if appointmentRequest.Date != "" {
		appointmentDB.Date = appointmentRequest.Date
	}
	if appointmentRequest.Notes != "" {
		appointmentDB.Notes = appointmentRequest.Notes
	}
	if appointmentRequest.DentistIdDentist != 0 {
		appointmentDB.DentistIdDentist = appointmentRequest.DentistIdDentist
	}
	if appointmentRequest.PatientIdPatient != 0 {
		appointmentDB.PatientIdPatient = appointmentRequest.PatientIdPatient
	}

	appointmentUpdate, err := ph.appointmentsUpdate.ModifyByID(id, appointmentDB)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, appointmentUpdate)
}

// DeleteAppointment
// @Summary      Eliminar a un turno por su ID.
// @Description  Elimina un odontólogo existente en la base de datos por su ID.
// @Tags         appointments
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 Appointment deleted
// @Failure		 400 {object} error: invalid id
// @Failure		 404 {object} error: appointment not found
// @Failure		 400 {object} error: delete failed
// @Router       /appointments/{id} [delete]
func (ph *AppointmentsHandler) DeleteAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = ph.appointmentsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	deleted := ph.appointmentsDelete.Delete(id)
	if deleted != nil {
		ctx.JSON(404, gin.H{"error": "delete failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "appointment deleted"})
}

// GetAppointmentByPatient
// @Summary      Obtener turnos por DNI de paciente
// @Description  Obtener turnos por DNI de paciente en la base de datos.
// @Tags         appointments
// @Produce      json
// @Param        credential_id path string true "Credential_ID"
// @Success      200 {object} appointments.Appointment
// @Failure	     404 {object} error: dentist not found
// @Router       /appointments/{credential_id} [get]
func (ph *AppointmentsHandler) GetAppointmentByPatient(ctx *gin.Context) {
	credentialId := ctx.Param("credential_id")
	appointment, err := ph.appointmentsGetterByPatient.GetByCredentialID(credentialId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}
	ctx.JSON(http.StatusOK, appointment)
}

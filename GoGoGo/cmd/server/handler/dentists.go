package handler

import (
	"GoGoGo/internal/dentists"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistCreator interface {
	Save(dentist dentists.Dentist) (dentists.Dentist, error)
}
type DentistGetter interface {
	GetByID(id int) (dentists.Dentist, error)
}

type DentistUpdate interface {
	ModifyByID(id int, dentist dentists.Dentist) (dentists.Dentist, error)
}

type DentistDelete interface {
	Delete(id int) error
}

type DentistsHandler struct {
	dentistsCreator DentistCreator
	dentistsGetter  DentistGetter
	dentistsUpdate  DentistUpdate
	dentistsDelete  DentistDelete
}

func NewDentistsHandler(creator DentistCreator, getter DentistGetter, update DentistUpdate, delete DentistDelete) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		dentistsCreator: creator,
		dentistsDelete:  delete,
		dentistsUpdate:  update,
	}
}

// PostDentist
// @Summary      Crea un nuevo odontologo
// @Description  Crea un nuevo odontólogo y lo almacena en la base de datos.
// @Tags         dentists
// @Produce      json
// @Param 		 body body dentists.Dentist true "Dentist"
// @Success      201 {object} dentists.Dentist
// @Failure		 400 {object} error
// @Failure		 500 {object} error: internal error
// @Router       /dentists [post]
func (ph *DentistsHandler) PostDentist(ctx *gin.Context) {

	dentistRequest := dentists.Dentist{}

	err := ctx.BindJSON(&dentistRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dentist, err := ph.dentistsCreator.Save(dentistRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, dentist)
}

// GetDentistByID
// @Summary      Obtener un odontologo por su ID
// @Description  Obtener un odontologo por su ID existente en la base de datos.
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} dentists.Dentist
// @Failure		 400 {object} error: invalid id
// @Failure	     404 {object} error: dentist not found
// @Router       /dentists/{id} [get]
func (ph *DentistsHandler) GetDentistByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	dentist, err := ph.dentistsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// PutDentist
// @Summary      Modificar un odontologo
// @Description  Modifica la información de un odontologo existente en la base de datos. Este endpoint permite actualizar los detalles de un odontólogo identificado por su ID.
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Param 		 body body dentists.Dentist true "Dentist"
// @Success      200 {object} dentists.Dentist
// @Failure		 400 {object} error: invalid id
// @Failure		 500 {object} error: internal error
// @Router       /dentists/{id} [put]
func (ph *DentistsHandler) PutDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	dentistRequest := dentists.Dentist{}
	err = ctx.BindJSON(&dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dentist, err := ph.dentistsUpdate.ModifyByID(id, dentistRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, dentist)
}

// PatchDentist
// @Summary      Modificar odontologo
// @Description  Modifica la información de un odontologo existente en la base de datos por alguno de sus campos. Los campos no proporcionados en la solicitud permanecerán sin cambios en el odontólogo.
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Param 		 body body dentists.Dentist true "Dentist"
// @Success      200 {object} dentists.Dentist
// @Failure		 400 {object} error: invalid id
// @Failure		 400 {object} error: dentist not found
// @Failure		 500 {object} error: internal error
// @Router       /dentists/{id} [patch]
func (ph *DentistsHandler) PatchDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	dentistRequest := dentists.Dentist{}
	err = ctx.BindJSON(&dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dentistDB, err := ph.dentistsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}

	if dentistRequest.FirstName != "" {
		dentistDB.FirstName = dentistRequest.FirstName
	}
	if dentistRequest.LastName != "" {
		dentistDB.LastName = dentistRequest.LastName
	}
	if dentistRequest.RegistrationNumber != 0 {
		dentistDB.RegistrationNumber = dentistRequest.RegistrationNumber
	}

	dentistUpdate, err := ph.dentistsUpdate.ModifyByID(id, dentistDB)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, dentistUpdate)
}

// DeleteDentist
// @Summary      Eliminar a un odontologo por su ID.
// @Description  Elimina un odontólogo existente en la base de datos por su ID.
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Success      200  Dentist deleted
// @Failure		 400 {object} error: invalid id
// @Failure		 400 {object} error: dentist not found
// @Failure		 400 {object} error: delete failed
// @Router       /dentists/{id} [delete]
func (ph *DentistsHandler) DeleteDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = ph.dentistsGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}

	deleted := ph.dentistsDelete.Delete(id)
	if deleted != nil {
		ctx.JSON(404, gin.H{"error": "delete failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dentist deleted"})
}

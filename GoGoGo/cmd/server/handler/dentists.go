package handler

import (
	"GoGoGo/internal/dentists"
	"fmt"
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
	dentistUpdate   DentistUpdate
	dentistDelete   DentistDelete
}

func NewDentistsHandler(creator DentistCreator, getter DentistGetter, update DentistUpdate, delete DentistDelete) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		dentistsCreator: creator,
		dentistDelete:   delete,
		dentistUpdate:   update,
	}
}

func (ph *DentistsHandler) PostDentist(ctx *gin.Context) {

	dentistRequest := dentists.Dentist{}

	err := ctx.BindJSON(&dentistRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(dentistRequest)

	dentist, err := ph.dentistsCreator.Save(dentistRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, dentist)
}

// GetDentistByID godoc
// @Summary      Gets a dentist by id
// @Description  Gets a dentist by id from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} dentists.Dentist
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

	dentist, err := ph.dentistUpdate.ModifyByID(id, dentistRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, dentist)
}

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

	dentistUpdate, err := ph.dentistUpdate.ModifyByID(id, dentistDB)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, dentistUpdate)
}

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

	deleted := ph.dentistDelete.Delete(id)
	if deleted != nil {
		ctx.JSON(404, gin.H{"error": "delete failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "dentist deleted"})
}

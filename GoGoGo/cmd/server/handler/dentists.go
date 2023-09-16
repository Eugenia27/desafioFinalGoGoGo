package handler

import (
	"GoGoGo/internal/dentists"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistGetter interface {
	GetByID(id int) (dentists.Dentist, error)
}

// type DentistCreator interface {
// 	ModifyByID(id int, dentist dentists.Dentist) (dentists.Dentist, error)
// }

type DentistsHandler struct {
	dentistsGetter  DentistGetter
	//dentistsCreator DentistCreator
}

func NewDentistsHandler(getter DentistGetter) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		//dentistsCreator: creator,
	}
}

// func NewDentistsHandler(getter DentistGetter, creator DentistCreator) *DentistsHandler {
// 	return &DentistsHandler{
// 		dentistsGetter:  getter,
// 		//dentistsCreator: creator,
// 	}
// }

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

// func (ph *DentistsHandler) PutDentist(ctx *gin.Context) {
// 	idParam := ctx.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		ctx.JSON(400, gin.H{"error": "invalid id"})
// 		return
// 	}
// 	dentistRequest := dentists.Dentist{}
// 	err = ctx.BindJSON(&dentistRequest)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	dentist, err := ph.dentistsCreator.ModifyByID(id, dentistRequest)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": "internal error"})
// 		return
// 	}
// 	ctx.JSON(200, dentist)
// }
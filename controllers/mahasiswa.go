package controllers

import (
	"golang3/forms"
	"golang3/models"

	"github.com/gin-gonic/gin"
)

var MahasiswaModel = new(models.MahasiswaModel)

type MahasiswaController struct{}

//FUNGSI CREATE
func (mahasiswa *MahasiswaController) Create(c *gin.Context) {
	var data forms.MahasiswaCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	err := MahasiswaModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"message": "Mahasiswa could not be Create", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "Mahasiswa Created"})
}

//FUNGSI MENAMPILKAN DATA
func (mahasiswa *MahasiswaController) Find(c *gin.Context) {
	list, err := MahasiswaModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

//FUNGSI MENAMPILKAN DATA BY ID
func (mahasiswa *MahasiswaController) FindId(c *gin.Context) {
	id := c.Param("id")
	list, err := MahasiswaModel.FindId(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find ID Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

//FUNGSI MENAMPILKAN DATA BY ID
func (mahasiswa *MahasiswaController) FindNim(c *gin.Context) {
	nim := c.Param("nim")
	list, err := MahasiswaModel.FindNim(nim)
	if err != nil {
		c.JSON(404, gin.H{"message": "Find NIM Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

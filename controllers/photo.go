package controllers

import (
	"golang3/forms"
	"golang3/models"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var PhotoModel = new(models.PhotoModel)

type PhotoController struct{}

//UPLOAD FOTO

func (photo *PhotoController) Upload(c *gin.Context) {
	//dir := config.GetDir()
	file, _, err := c.Request.FormFile("gambar") // Here is the bug
	filename := "null"
	if err == nil {
		if err != nil {
			c.JSON(404, gin.H{"message": "terjadi kesalahan", "error": err.Error()})
			c.Abort()
			return
		} // Err Handling

		path := "./assets/foto/mahasiswa/"
		time := time.Now().UnixNano() / 1000000
		filename = strconv.Itoa(int(time)) + ".png"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
		out, err := os.Create(path + "/" + filename)
		if err != nil {
			c.JSON(404, gin.H{"message": "terjadi kesalahan", "error": err.Error()})
			c.Abort()
		} // Err Handling
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(404, gin.H{"message": "terjadi kesalahan", "error": err.Error()})
			c.Abort()
		} // Err Handling
	}

	data := forms.PhotoCommand{
		Gambar: "foto/mahasiswa/" + filename,
		Nama:   c.PostForm("nama"),
	}
	er := PhotoModel.Create(data)
	if er != nil {
		if strings.Contains(err.Error(), "duplicate") == true {
			c.JSON(400, gin.H{
				"message": "Photo Mahasiswa sudah ada",
				"error":   er.Error(),
			})
			c.Abort()
		} else {
			c.JSON(400, gin.H{
				"message": "Photo Mahasiswa gagal dibuat",
				"error":   er.Error(),
			})
			c.Abort()
		}
		return
	} else {
		c.JSON(200, gin.H{"message": "Photo Mahasiswa Created", "status": "ok"})
	}

}

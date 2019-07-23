package main

import (
	"golang3/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/mahasiswa")
	{
		Mahasiswa := new(controllers.MahasiswaController)
		Photo := new(controllers.PhotoController)
		//MEMBUAT DATA
		v1.POST("/create_mhs", Mahasiswa.Create)
		//DATA ALL
		v1.GET("/view_mhs", Mahasiswa.Find)
		//DATA BY ID
		v1.GET("/view_mhs/:id", Mahasiswa.FindId)
		//DATA BY NAME
		v1.GET("/view_mhs_nim/:nim", Mahasiswa.FindNim)
		//UPLOAD GAMBAR + FORM INPUTAN
		v1.POST("/upload_gmbr", Photo.Upload)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "Not Found"})
	})

	router.Run()
}

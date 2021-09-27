package main

import (
	"crud.com/config"
	"crud.com/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Setup config tempat database
	db := config.SetupDatabaseConnection()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/kategori", service.FindAllKategori)
	r.POST("/kategori", service.InputKategori)
	r.PUT("/kategori/:id", service.UpdateKategori)
	r.DELETE("/kategori/:id", service.DeleteKategori)

	r.GET("/barang", service.FindAllBarang)
	r.POST("/barang", service.InputBarang)
	r.PUT("/barang/:id", service.UpdateBarang)
	r.DELETE("/barang/:id", service.DeleteBarang)

	r.Run()
}

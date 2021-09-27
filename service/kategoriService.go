package service

import (
	"net/http"

	"crud.com/models/dto"
	"crud.com/models/entity"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindAllKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var kate []entity.Kategori
	db.Find(&kate)
	c.JSON(http.StatusOK, gin.H{"data": kate})
}

// Tambah data kategori
func InputKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput dto.KategoriCreateDTO
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	kate := entity.Kategori{
		NamaKategori: dataInput.NamaKategori,
	}

	db.Create(&kate)

	c.JSON(http.StatusOK, gin.H{"data": kate})
}

// Ubah data kategori
func UpdateKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var kate entity.Kategori
	if err := db.Where("id = ?", c.Param("id")).First(&kate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data kategori tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput dto.KategoriCreateDTO
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&kate).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": kate})
}

// Hapus data kategori
func DeleteKategori(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var kate entity.Kategori
	if err := db.Where("id = ?", c.Param("id")).First(&kate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data kategori tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&kate)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

package service

import (
	"net/http"

	"crud.com/models/dto"
	"crud.com/models/entity"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func FindAllBarang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var barang []entity.Barang
	db.Find(&barang)
	c.JSON(http.StatusOK, gin.H{"data": barang})
}

// Tambah data barang
func InputBarang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput dto.BarangCreateDTO
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	barang := entity.Barang{
		NamaBarang: dataInput.NamaBarang,
		Deskripsi:  dataInput.Deskripsi,
		Harga:      dataInput.Harga,
		IdKategori: dataInput.IdKategori,
	}

	db.Create(&barang)

	c.JSON(http.StatusOK, gin.H{"data": barang})
}

// Ubah data barang
func UpdateBarang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var barang entity.Barang
	if err := db.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data barang tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput dto.BarangCreateDTO
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&barang).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{"data": barang})
}

// Hapus data barang
func DeleteBarang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek data
	var barang entity.Barang
	if err := db.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data barang tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&barang)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

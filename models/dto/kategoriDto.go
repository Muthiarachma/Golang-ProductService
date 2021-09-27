package dto

//KategoriCreateDTO untuk create
type KategoriCreateDTO struct {
	NamaKategori string `json:"namaKategori" form:"namaKategori" binding:"required"`
}

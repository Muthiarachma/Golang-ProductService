package dto

//BarangCreateDTO untuk create
type BarangCreateDTO struct {
	NamaBarang string  `json:"namaBarang" form:"namaBarang" binding:"required"`
	Deskripsi  string  `json:"deskripsi" form:"deskripsi" binding:"required"`
	Harga      float32 `json:"harga" form:"harga" binding:"required"`
	IdKategori uint32  `json:"IdKategori,omitempty" from:"IdKategori,omitempty"`
}

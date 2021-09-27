package entity

type Barang struct {
	ID         uint32  `gorm:"primary_key:auto_increment" json:"id"`
	NamaBarang string  `gorm:"type:varchar(255)" json:"namaBarang"`
	Deskripsi  string  `gorm:"type:varchar(255)" json:"deskripsi"`
	Harga      float32 `gorm:"type:float" json:"harga"`
	// Kategori   Kategori `gorm:"foreignkey:IdKategori;constraint:onUpadte:CASCADE,onDelete:CASCADE"`
	IdKategori uint32 `gorm:"foreignkey:IdKategori"`
}

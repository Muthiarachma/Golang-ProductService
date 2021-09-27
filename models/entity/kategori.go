package entity

type Kategori struct {
	ID           uint32 `gorm:"primary_key:auto_increment" json:"id"`
	NamaKategori string `gorm:"type:varchar(255)" json:"namaKategori"`
}

package models

type Siswa struct {
	ID          int    `json:"id"`
	Nis        string `json:"nis" gorm:"type: varchar(20)"`
	Name       string `json:"name" gorm:"type: varchar(255)"`
	Alamat    string `json:"alamat" gorm:"type: varchar(255)"`

}
package userdto

type CreatSiswa struct {
	Nis string `json:"nis" form:"nis" validate:"required"`
	Name string `json:"name" form:"name" validate:"required"`
	Alamat string `json:"alamat" form:"alamat" validate:"required"`
}

type SiswaResponse struct {
	Id int `json:"id"`
	Nis string `json:"nis"`
	Name string `json:"name"`
	Alamat string `json:"alamat"`
}
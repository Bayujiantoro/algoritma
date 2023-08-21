package repositories

import (
	"test/models"

	"gorm.io/gorm"
)

type SiswaRepo interface {
	FindSiswa() ([]models.Siswa, error)
	UpdateSiswa(siswa models.Siswa ) (models.Siswa, error)
	CreateSiswa(siswa models.Siswa)  (models.Siswa, error)
	DeleteSiswa(siswa models.Siswa , Id int) (models.Siswa, error)
	GetSiswa(ID int) (models.Siswa,error)
}

func RepoSiswa(db *gorm.DB) *repositories {
	return &repositories{db}
}


func (r *repositories) GetSiswa(ID int)(models.Siswa, error) {
	var Siswa models.Siswa

	err := r.db.First(&Siswa, ID).Error
	return Siswa, err
}

func (r *repositories) FindSiswa() ([]models.Siswa, error) {
	var User []models.Siswa
	err := r.db.Find(&User).Error

	return User, err
}

func (r *repositories) CreateSiswa(user models.Siswa) (models.Siswa, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *repositories) UpdateSiswa(user models.Siswa) (models.Siswa, error) {

	err := r.db.Save(&user).Error

	return user , err
}

func (r *repositories) DeleteSiswa(user models.Siswa, ID int) (models.Siswa, error) {
	err := r.db.Delete(&user,  ID).Scan(&user).Error

	return user, err
}
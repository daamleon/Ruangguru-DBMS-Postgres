package repository

import (
	"a21hc3NpZ25tZW50/model"
	// "fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	err := s.db.Create(student).Error
	return err
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error
	return err
}

func (s *studentRepoImpl) Delete(id int) error {
	err := s.db.Where("id = ?", id).Delete(&model.Student{}).Error
	return err
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	err := s.db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
    studentClasses := []model.StudentClass{} // Inisialisasi slice kosong
    err := s.db.Table("students").
        Select(`students.name, students.address, 
                classes.name as class_name, 
                classes.professor, 
                classes.room_number`).
        Joins("left join classes on students.class_id = classes.id").
        Scan(&studentClasses).Error
    if err != nil {
        return nil, err
    }
    return &studentClasses, nil
}



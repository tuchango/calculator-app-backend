package calculationService

import "gorm.io/gorm"

type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculations() ([]Calculation, error)
	GetCalculationById(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}
}

func (r *calcRepository) CreateCalculation(calc Calculation) error {
	err := r.db.Create(&calc).Error
	return err
}

func (r *calcRepository) GetAllCalculations() ([]Calculation, error) {
	var calcs []Calculation
	err := r.db.Find(&calcs).Error
	return calcs, err
}

func (r *calcRepository) GetCalculationById(id string) (Calculation, error) {
	var calc Calculation
	err := r.db.First(&calc, "id = ?", id).Error
	return calc, err
}

func (r *calcRepository) UpdateCalculation(calc Calculation) error {
	err := r.db.Save(&calc).Error
	return err
}

func (r *calcRepository) DeleteCalculation(id string) error {
	err := r.db.Delete(&Calculation{}, "id = ?", id).Error
	return err
}

package calculationService

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculations() ([]Calculation, error)
	GetCalculationById(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r}
}

func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) GetAllCalculations() ([]Calculation, error) {
	calcs, err := s.repo.GetAllCalculations()
	if err != nil {
		return nil, err
	}

	return calcs, nil
}

func (s *calcService) GetCalculationById(id string) (Calculation, error) {
	calc, err := s.repo.GetCalculationById(id)
	if err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) UpdateCalculation(id, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationById(id)
	if err != nil {
		return Calculation{}, err
	}

	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc.Expression = expression
	calc.Result = result

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	err := s.repo.DeleteCalculation(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	res, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", res), nil
}

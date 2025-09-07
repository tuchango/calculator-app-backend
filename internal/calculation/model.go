package calculation

type Calculation struct {
	ID         string `json:"id" gorm:"primary key"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}

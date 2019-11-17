package model

// Indicates one of the constraints of a range of business values.
type Term1 struct {

	// Provides the relationship between a variable and a fixed value.
	Operator *Operator1Code `xml:"Oprtr"`

	// Indicates the value.
	Value *RateOrAbsoluteValue1Choice `xml:"Val"`
}

func (t *Term1) SetOperator(value string) {
	t.Operator = (*Operator1Code)(&value)
}

func (t *Term1) AddValue() *RateOrAbsoluteValue1Choice {
	t.Value = new(RateOrAbsoluteValue1Choice)
	return t.Value
}

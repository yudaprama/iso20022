package model

// Provides constrains on a range of business values.
type AmountOrPercentageRange struct {

	// Indication of the relationship between two variables.
	Operation *Operation1Code `xml:"Opr,omitempty"`

	// Indicates one of the constraints of a range of business values.
	Term []*Term1 `xml:"Term,omitempty"`
}

func (a *AmountOrPercentageRange) SetOperation(value string) {
	a.Operation = (*Operation1Code)(&value)
}

func (a *AmountOrPercentageRange) AddTerm() *Term1 {
	newValue := new(Term1)
	a.Term = append(a.Term, newValue)
	return newValue
}

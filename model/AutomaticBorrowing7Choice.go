package model

// Choice of format for the automatic borrowing information.
type AutomaticBorrowing7Choice struct {

	// Condition for automatic borrowing expressed as an ISO 20022 code.
	Code *AutoBorrowing2Code `xml:"Cd"`

	// Condition for automatic borrowing expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AutomaticBorrowing7Choice) SetCode(value string) {
	a.Code = (*AutoBorrowing2Code)(&value)
}

func (a *AutomaticBorrowing7Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}

package model

// Choice of format for the automatic borrowing information.
type AutomaticBorrowing6Choice struct {

	// Condition for automatic borrowing expressed as an ISO 20022 code.
	Code *AutoBorrowing1Code `xml:"Cd"`

	// Condition for automatic borrowing expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AutomaticBorrowing6Choice) SetCode(value string) {
	a.Code = (*AutoBorrowing1Code)(&value)
}

func (a *AutomaticBorrowing6Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}

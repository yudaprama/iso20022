package model

// Choice of format for the automatic borrowing information.
type AutomaticBorrowing5Choice struct {

	// Condition for automatic borrowing expressed as an ISO 20022 code.
	Code *AutoBorrowing1Code `xml:"Cd"`

	// Condition for automatic borrowing expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (a *AutomaticBorrowing5Choice) SetCode(value string) {
	a.Code = (*AutoBorrowing1Code)(&value)
}

func (a *AutomaticBorrowing5Choice) AddProprietary() *GenericIdentification38 {
	a.Proprietary = new(GenericIdentification38)
	return a.Proprietary
}

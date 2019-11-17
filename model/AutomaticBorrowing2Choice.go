package model

// Choice of format for the automatic borrowing information.
type AutomaticBorrowing2Choice struct {

	// Condition for automatic borrowing expressed as an ISO 20022 code.
	Code *AutoBorrowing2Code `xml:"Cd"`

	// Condition for automatic borrowing expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AutomaticBorrowing2Choice) SetCode(value string) {
	a.Code = (*AutoBorrowing2Code)(&value)
}

func (a *AutomaticBorrowing2Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}

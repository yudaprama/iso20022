package model

// Choice of format for the automatic borrowing information.
type AutomaticBorrowing1Choice struct {

	// Condition for automatic borrowing expressed as an ISO 20022 code.
	Code *AutoBorrowing1Code `xml:"Cd"`

	// Condition for automatic borrowing expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AutomaticBorrowing1Choice) SetCode(value string) {
	a.Code = (*AutoBorrowing1Code)(&value)
}

func (a *AutomaticBorrowing1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}

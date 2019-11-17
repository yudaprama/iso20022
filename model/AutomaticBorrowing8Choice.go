package model

// Choice of format for the automatic borrowing information.
type AutomaticBorrowing8Choice struct {

	// Condition for automatic borrowing expressed as an ISO 20022 code.
	Code *AutoBorrowing1Code `xml:"Cd"`

	// Condition for automatic borrowing expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AutomaticBorrowing8Choice) SetCode(value string) {
	a.Code = (*AutoBorrowing1Code)(&value)
}

func (a *AutomaticBorrowing8Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}

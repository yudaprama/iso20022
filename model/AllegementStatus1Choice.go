package model

// Choice of format for the allegement status.
type AllegementStatus1Choice struct {

	// Status of the allegement reported.
	Code *AllegementStatus1Code `xml:"Cd"`

	// Status of the allegement reported.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (a *AllegementStatus1Choice) SetCode(value string) {
	a.Code = (*AllegementStatus1Code)(&value)
}

func (a *AllegementStatus1Choice) AddProprietary() *GenericIdentification20 {
	a.Proprietary = new(GenericIdentification20)
	return a.Proprietary
}

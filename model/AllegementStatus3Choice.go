package model

// Choice of format for the allegement status.
type AllegementStatus3Choice struct {

	// Status of the allegement reported.
	Code *AllegementStatus1Code `xml:"Cd"`

	// Status of the allegement reported.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (a *AllegementStatus3Choice) SetCode(value string) {
	a.Code = (*AllegementStatus1Code)(&value)
}

func (a *AllegementStatus3Choice) AddProprietary() *GenericIdentification30 {
	a.Proprietary = new(GenericIdentification30)
	return a.Proprietary
}

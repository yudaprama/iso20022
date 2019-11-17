package model

// Choice of format for the allegement status.
type AllegementStatus4Choice struct {

	// Status of the allegement reported.
	Code *AllegementStatus1Code `xml:"Cd"`

	// Status of the allegement reported.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AllegementStatus4Choice) SetCode(value string) {
	a.Code = (*AllegementStatus1Code)(&value)
}

func (a *AllegementStatus4Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}

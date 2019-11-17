package model

// Indicates the type of obligation.
type ObligationType1Choice struct {

	// Indicates the type of the obligation using a code.
	Code *ObligationType1Code `xml:"Cd"`

	// Indicates the type of the obligation using a proprietary format.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *ObligationType1Choice) SetCode(value string) {
	o.Code = (*ObligationType1Code)(&value)
}

func (o *ObligationType1Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}

package model

// Choice of format for the option style.
type OptionStyle8Choice struct {

	// Option style expressed as an ISO 20022 code.
	Code *OptionStyle2Code `xml:"Cd"`

	// Option style expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionStyle8Choice) SetCode(value string) {
	o.Code = (*OptionStyle2Code)(&value)
}

func (o *OptionStyle8Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}

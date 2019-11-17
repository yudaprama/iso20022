package model

// Choice of format for the option style.
type OptionStyle6Choice struct {

	// Option style expressed as an ISO 20022 code.
	Code *OptionStyle4Code `xml:"Cd"`

	// Option style expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (o *OptionStyle6Choice) SetCode(value string) {
	o.Code = (*OptionStyle4Code)(&value)
}

func (o *OptionStyle6Choice) AddProprietary() *GenericIdentification38 {
	o.Proprietary = new(GenericIdentification38)
	return o.Proprietary
}

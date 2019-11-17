package model

// Choice of format for the option style.
type OptionStyle9Choice struct {

	// Option style expressed as an ISO 20022 code.
	Code *OptionStyle2Code `xml:"Cd"`

	// Option style expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OptionStyle9Choice) SetCode(value string) {
	o.Code = (*OptionStyle2Code)(&value)
}

func (o *OptionStyle9Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}

package model

// Choice of format for the option style.
type OptionStyle4Choice struct {

	// Option style expressed as an ISO 20022 code.
	Code *OptionStyle2Code `xml:"Cd"`

	// Option style expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionStyle4Choice) SetCode(value string) {
	o.Code = (*OptionStyle2Code)(&value)
}

func (o *OptionStyle4Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}

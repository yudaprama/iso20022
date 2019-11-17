package model

// Choice of format for the option type.
type OptionType2Choice struct {

	// Option type expressed as an ISO 20022 code.
	Code *OptionType1Code `xml:"Cd"`

	// Option type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (o *OptionType2Choice) SetCode(value string) {
	o.Code = (*OptionType1Code)(&value)
}

func (o *OptionType2Choice) AddProprietary() *GenericIdentification20 {
	o.Proprietary = new(GenericIdentification20)
	return o.Proprietary
}

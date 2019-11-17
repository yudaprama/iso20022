package model

// Choice of format for the option type.
type OptionType6Choice struct {

	// Option type expressed as an ISO 20022 code.
	Code *OptionType1Code `xml:"Cd"`

	// Option type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (o *OptionType6Choice) SetCode(value string) {
	o.Code = (*OptionType1Code)(&value)
}

func (o *OptionType6Choice) AddProprietary() *GenericIdentification30 {
	o.Proprietary = new(GenericIdentification30)
	return o.Proprietary
}

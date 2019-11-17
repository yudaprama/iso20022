package model

// Choice of format for the option type.
type OptionType4Choice struct {

	// Option type expressed as an ISO 20022 code.
	Code *OptionType1Code `xml:"Cd"`

	// Option type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (o *OptionType4Choice) SetCode(value string) {
	o.Code = (*OptionType1Code)(&value)
}

func (o *OptionType4Choice) AddProprietary() *GenericIdentification38 {
	o.Proprietary = new(GenericIdentification38)
	return o.Proprietary
}

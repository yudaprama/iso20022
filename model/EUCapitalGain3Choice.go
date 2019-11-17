package model

// Choice of formats for the EU capital gain.
type EUCapitalGain3Choice struct {

	// EU capital gain expressed as a code.
	Code *EUCapitalGain2Code `xml:"Cd"`

	// EU capital gain expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (e *EUCapitalGain3Choice) SetCode(value string) {
	e.Code = (*EUCapitalGain2Code)(&value)
}

func (e *EUCapitalGain3Choice) AddProprietary() *GenericIdentification47 {
	e.Proprietary = new(GenericIdentification47)
	return e.Proprietary
}

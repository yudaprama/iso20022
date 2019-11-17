package model

// Choice of formats for a settlement purpose.
type Purpose3Choice struct {

	// Underlying reason for the SSI instruction, expressed as a code.
	SecuritiesPurposeCode *ExternalSecuritiesPurpose1Code `xml:"SctiesPurpCd"`

	// Underlying reason for the SSI instruction, expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (p *Purpose3Choice) SetSecuritiesPurposeCode(value string) {
	p.SecuritiesPurposeCode = (*ExternalSecuritiesPurpose1Code)(&value)
}

func (p *Purpose3Choice) AddProprietary() *GenericIdentification1 {
	p.Proprietary = new(GenericIdentification1)
	return p.Proprietary
}

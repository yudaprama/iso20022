package model

// Choice between a code and or a data source scheme to determine the account type.
type PurposeCode7Choice struct {

	// Securities account purpose as an ISO 20022 code.
	Code *SecuritiesAccountPurposeType1Code `xml:"Cd"`

	// Securities account purpose as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PurposeCode7Choice) SetCode(value string) {
	p.Code = (*SecuritiesAccountPurposeType1Code)(&value)
}

func (p *PurposeCode7Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}

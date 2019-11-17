package model

// Choice between a code and or a data source scheme to determine the account type.
type PurposeCode8Choice struct {

	// Securities account purpose as an ISO 20022 code.
	Code *SecuritiesAccountPurposeType1Code `xml:"Cd"`

	// Securities account purpose as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PurposeCode8Choice) SetCode(value string) {
	p.Code = (*SecuritiesAccountPurposeType1Code)(&value)
}

func (p *PurposeCode8Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}

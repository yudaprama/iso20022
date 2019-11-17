package model

// Choice between a code and or a data source scheme to determine the account type.
type PurposeCode5Choice struct {

	// Specifies the type of securities account.
	Code *SecuritiesAccountPurposeType1Code `xml:"Cd"`

	// Specifies the execution priority of the instruction with a proprietary scheme.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (p *PurposeCode5Choice) SetCode(value string) {
	p.Code = (*SecuritiesAccountPurposeType1Code)(&value)
}

func (p *PurposeCode5Choice) AddProprietary() *GenericIdentification38 {
	p.Proprietary = new(GenericIdentification38)
	return p.Proprietary
}

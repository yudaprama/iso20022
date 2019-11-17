package model

// Choice between a code and or a data source scheme to determine the account type.
type PurposeCode1Choice struct {

	// Specifies the type of securities account.
	Code *SecuritiesAccountPurposeType1Code `xml:"Cd"`

	// Account type is defined using a data source scheme.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *PurposeCode1Choice) SetCode(value string) {
	p.Code = (*SecuritiesAccountPurposeType1Code)(&value)
}

func (p *PurposeCode1Choice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}

package model

// Choice of format for a party role.
type Role5Choice struct {

	// Role expressed in as a code.
	Code *InvestmentFundRole2Code `xml:"Cd"`

	// Role expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (r *Role5Choice) SetCode(value string) {
	r.Code = (*InvestmentFundRole2Code)(&value)
}

func (r *Role5Choice) AddProprietary() *GenericIdentification36 {
	r.Proprietary = new(GenericIdentification36)
	return r.Proprietary
}

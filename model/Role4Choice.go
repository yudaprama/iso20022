package model

// Choice of format for a party role.
type Role4Choice struct {

	// Role of the party in the activity expressed as a code.
	Code *InvestmentFundRole2Code `xml:"Cd"`

	// Role of the party in the activity expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *Role4Choice) SetCode(value string) {
	r.Code = (*InvestmentFundRole2Code)(&value)
}

func (r *Role4Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}

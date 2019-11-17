package model

// Choice of format for a party role.
type Role7Choice struct {

	// Role of the party in the activity expressed as an ISO 20022 code.
	Code *InvestmentFundRole2Code `xml:"Cd"`

	// Role of the party in the activity expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`

	// Role of the party in the activity.
	Text *RestrictedFINXMax350Text `xml:"Txt"`
}

func (r *Role7Choice) SetCode(value string) {
	r.Code = (*InvestmentFundRole2Code)(&value)
}

func (r *Role7Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}

func (r *Role7Choice) SetText(value string) {
	r.Text = (*RestrictedFINXMax350Text)(&value)
}

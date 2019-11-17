package model

// Choice of format for a party role.
type Role2Choice struct {

	// Role of the party in the activity expressed as an ISO 20022 code.
	Code *InvestmentFundRole2Code `xml:"Cd"`

	// Role of the party in the activity expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`

	// Role of the party in the activity.
	Text *Max350Text `xml:"Txt"`
}

func (r *Role2Choice) SetCode(value string) {
	r.Code = (*InvestmentFundRole2Code)(&value)
}

func (r *Role2Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}

func (r *Role2Choice) SetText(value string) {
	r.Text = (*Max350Text)(&value)
}

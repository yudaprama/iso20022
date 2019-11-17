package model

// Choice of format for a party role.
type Role6Choice struct {

	// Role of the party in the activity expressed as an ISO 20022 code.
	Code *InvestmentFundRole2Code `xml:"Cd"`

	// Role of the party in the activity expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`

	// Role of the party in the activity.
	Text *Max350Text `xml:"Txt"`
}

func (r *Role6Choice) SetCode(value string) {
	r.Code = (*InvestmentFundRole2Code)(&value)
}

func (r *Role6Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}

func (r *Role6Choice) SetText(value string) {
	r.Text = (*Max350Text)(&value)
}

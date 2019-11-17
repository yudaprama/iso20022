package model

// Restrictions that may be applied to an account or investment plan.
type AccountRestrictions1 struct {

	// Restrictions and/or limitations on the account or account party.
	Limitation *Max350Text `xml:"Lmttn,omitempty"`

	// Additional information.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Information or instructions for the by-passing of validations in the account registration process.
	AccountValidation *Max350Text `xml:"AcctVldtn,omitempty"`

	// Type or identification of the restriction
	Type *Max35Text `xml:"Tp,omitempty"`

	// Regulator that may have to be informed about restrictions or limitations on the account or account party.
	Regulator *PartyIdentification70Choice `xml:"Rgltr,omitempty"`

	// Status of the restriction.
	Status *RestrictionStatus1Choice `xml:"Sts,omitempty"`

	// Period of the restriction.
	Period *DateTimePeriodDetails1 `xml:"Prd,omitempty"`
}

func (a *AccountRestrictions1) SetLimitation(value string) {
	a.Limitation = (*Max350Text)(&value)
}

func (a *AccountRestrictions1) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max350Text)(&value)
}

func (a *AccountRestrictions1) SetAccountValidation(value string) {
	a.AccountValidation = (*Max350Text)(&value)
}

func (a *AccountRestrictions1) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AccountRestrictions1) AddRegulator() *PartyIdentification70Choice {
	a.Regulator = new(PartyIdentification70Choice)
	return a.Regulator
}

func (a *AccountRestrictions1) AddStatus() *RestrictionStatus1Choice {
	a.Status = new(RestrictionStatus1Choice)
	return a.Status
}

func (a *AccountRestrictions1) AddPeriod() *DateTimePeriodDetails1 {
	a.Period = new(DateTimePeriodDetails1)
	return a.Period
}

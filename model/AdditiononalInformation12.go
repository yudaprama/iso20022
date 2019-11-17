package model

// Restrictions, remarks or notes that may be applied to an account or investment plan.
type AdditiononalInformation12 struct {

	// Restrictions and/or limitations on the account or party.
	Limitation *Max350Text `xml:"Lmttn,omitempty"`

	// Additional information such as remarks or notes that must be conveyed about the account management activity or party.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Information or instructions for the by-passing of validations in the account registration process.
	AccountValidation *Max350Text `xml:"AcctVldtn,omitempty"`

	// Type or identification of the remark, note, limitation or restriction
	Type *Max35Text `xml:"Tp,omitempty"`

	// Regulator that may have to be informed about the remark, note, limitation or restriction.
	Regulator *PartyIdentification70Choice `xml:"Rgltr,omitempty"`

	// Status of the remark , note, limitation or restriction.
	Status *RestrictionStatus1Choice `xml:"Sts,omitempty"`

	// Period of the restriction.
	Period *DateTimePeriodDetails1 `xml:"Prd,omitempty"`
}

func (a *AdditiononalInformation12) SetLimitation(value string) {
	a.Limitation = (*Max350Text)(&value)
}

func (a *AdditiononalInformation12) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max350Text)(&value)
}

func (a *AdditiononalInformation12) SetAccountValidation(value string) {
	a.AccountValidation = (*Max350Text)(&value)
}

func (a *AdditiononalInformation12) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AdditiononalInformation12) AddRegulator() *PartyIdentification70Choice {
	a.Regulator = new(PartyIdentification70Choice)
	return a.Regulator
}

func (a *AdditiononalInformation12) AddStatus() *RestrictionStatus1Choice {
	a.Status = new(RestrictionStatus1Choice)
	return a.Status
}

func (a *AdditiononalInformation12) AddPeriod() *DateTimePeriodDetails1 {
	a.Period = new(DateTimePeriodDetails1)
	return a.Period
}

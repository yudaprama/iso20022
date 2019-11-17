package model

// Party and account information.
type ExtendedParty9 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation12 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty9) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty9) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation12 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation12)
	return e.OtherPartyDetails
}

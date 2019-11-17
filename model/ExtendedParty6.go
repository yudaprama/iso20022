package model

// Other type of party.
type ExtendedParty6 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation9 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty6) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty6) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation9 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation9)
	return e.OtherPartyDetails
}

package model

// Other type of party.
type ExtendedParty4 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation7 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty4) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty4) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation7 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation7)
	return e.OtherPartyDetails
}

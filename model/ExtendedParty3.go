package model

// Other type of party.
type ExtendedParty3 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation6 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty3) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty3) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation6 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation6)
	return e.OtherPartyDetails
}

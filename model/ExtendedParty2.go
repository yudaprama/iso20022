package model

// Other type of party.
type ExtendedParty2 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation5 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty2) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty2) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation5 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation5)
	return e.OtherPartyDetails
}

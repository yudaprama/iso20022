package model

// Other type of party.
type ExtendedParty12 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation15 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty12) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty12) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation15 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation15)
	return e.OtherPartyDetails
}

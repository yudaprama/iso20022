package model

// Other type of party.
type ExtendedParty1 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation4 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty1) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty1) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation4 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation4)
	return e.OtherPartyDetails
}

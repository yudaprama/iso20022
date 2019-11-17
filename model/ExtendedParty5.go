package model

// Other type of party.
type ExtendedParty5 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation8 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty5) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty5) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation8 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation8)
	return e.OtherPartyDetails
}

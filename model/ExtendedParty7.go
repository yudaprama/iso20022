package model

// Other type of party.
type ExtendedParty7 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation10 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty7) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty7) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation10 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation10)
	return e.OtherPartyDetails
}

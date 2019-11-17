package model

// Other type of party.
type ExtendedParty10 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation13 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty10) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty10) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation13 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation13)
	return e.OtherPartyDetails
}

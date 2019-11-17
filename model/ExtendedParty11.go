package model

// Party and account information.
type ExtendedParty11 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation14 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty11) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty11) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation14 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation14)
	return e.OtherPartyDetails
}

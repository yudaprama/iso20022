package model

// Other type of party.
type ExtendedParty8 struct {

	// Other type of party's role.
	ExtendedPartyRole *Extended350Code `xml:"XtndedPtyRole"`

	// Detailed ownership information about a party.
	OtherPartyDetails *InvestmentAccountOwnershipInformation11 `xml:"OthrPtyDtls"`
}

func (e *ExtendedParty8) SetExtendedPartyRole(value string) {
	e.ExtendedPartyRole = (*Extended350Code)(&value)
}

func (e *ExtendedParty8) AddOtherPartyDetails() *InvestmentAccountOwnershipInformation11 {
	e.OtherPartyDetails = new(InvestmentAccountOwnershipInformation11)
	return e.OtherPartyDetails
}

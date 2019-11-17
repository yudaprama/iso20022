package model

// Specifies the party and owner type.
type AccountRole1 struct {

	// Account owner/user identification and contact information.
	Party *PartyIdentification41 `xml:"Pty"`

	// Defines account owners/users relation to the account.
	OwnerType *OwnerType1 `xml:"OwnrTp"`

	// Start date related to the role.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// End date related to the role.
	EndDate *ISODate `xml:"EndDt,omitempty"`
}

func (a *AccountRole1) AddParty() *PartyIdentification41 {
	a.Party = new(PartyIdentification41)
	return a.Party
}

func (a *AccountRole1) AddOwnerType() *OwnerType1 {
	a.OwnerType = new(OwnerType1)
	return a.OwnerType
}

func (a *AccountRole1) SetStartDate(value string) {
	a.StartDate = (*ISODate)(&value)
}

func (a *AccountRole1) SetEndDate(value string) {
	a.EndDate = (*ISODate)(&value)
}

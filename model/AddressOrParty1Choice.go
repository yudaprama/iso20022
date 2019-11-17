package model

// Choice of either address, or name and address.
type AddressOrParty1Choice struct {

	// New beneficiary address.
	NewAddress *PostalAddress6 `xml:"NewAdr"`

	// New beneficiary.
	NewBeneficiary *NameAndAddress10 `xml:"NewBnfcry"`
}

func (a *AddressOrParty1Choice) AddNewAddress() *PostalAddress6 {
	a.NewAddress = new(PostalAddress6)
	return a.NewAddress
}

func (a *AddressOrParty1Choice) AddNewBeneficiary() *NameAndAddress10 {
	a.NewBeneficiary = new(NameAndAddress10)
	return a.NewBeneficiary
}

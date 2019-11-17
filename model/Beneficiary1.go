package model

// Details related to the beneficiary.
type Beneficiary1 struct {

	// New beneficiary address, or new beneficiary name and address.
	NewAddressOrNewBeneficiary *AddressOrParty1Choice `xml:"NewAdrOrNewBnfcry"`

	// Additional information concerning the amended beneficiary details.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (b *Beneficiary1) AddNewAddressOrNewBeneficiary() *AddressOrParty1Choice {
	b.NewAddressOrNewBeneficiary = new(AddressOrParty1Choice)
	return b.NewAddressOrNewBeneficiary
}

func (b *Beneficiary1) AddAdditionalInformation(value string) {
	b.AdditionalInformation = append(b.AdditionalInformation, (*Max2000Text)(&value))
}

package model

// Details related to the local undertaking.
type Undertaking11 struct {

	// Details related to the requested new amount for the local undertaking.
	NewUndertakingAmount *UndertakingAmount2 `xml:"NewUdrtkgAmt,omitempty"`

	// Details related to the requested new expiry terms for the local undertaking.
	NewExpiryDetails *ExpiryDetails1 `xml:"NewXpryDtls,omitempty"`

	// Details related to the requested new beneficiary for the local undertaking.
	NewBeneficiary *PartyIdentification43 `xml:"NewBnfcry,omitempty"`

	// Details related to the requested new terms and conditions for the local undertaking.
	NewUndertakingTermsAndConditions *Narrative1 `xml:"NewUdrtkgTermsAndConds,omitempty"`

	// Details related to the delivery channel for the amended local undertaking.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`
}

func (u *Undertaking11) AddNewUndertakingAmount() *UndertakingAmount2 {
	u.NewUndertakingAmount = new(UndertakingAmount2)
	return u.NewUndertakingAmount
}

func (u *Undertaking11) AddNewExpiryDetails() *ExpiryDetails1 {
	u.NewExpiryDetails = new(ExpiryDetails1)
	return u.NewExpiryDetails
}

func (u *Undertaking11) AddNewBeneficiary() *PartyIdentification43 {
	u.NewBeneficiary = new(PartyIdentification43)
	return u.NewBeneficiary
}

func (u *Undertaking11) AddNewUndertakingTermsAndConditions() *Narrative1 {
	u.NewUndertakingTermsAndConditions = new(Narrative1)
	return u.NewUndertakingTermsAndConditions
}

func (u *Undertaking11) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

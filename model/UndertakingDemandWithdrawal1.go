package model

// Details of the demand withdrawal notification.
type UndertakingDemandWithdrawal1 struct {

	// Details related to the identification of the undertaking.
	UndertakingIdentification *Undertaking6 `xml:"UdrtkgId"`

	// Unique and unambiguous identifier assigned by the advising party to the undertaking.
	AdvisingPartyReferenceNumber *Max35Text `xml:"AdvsgPtyRefNb,omitempty"`

	// Details related to the demand.
	DemandDetails *Demand3 `xml:"DmndDtls"`

	// Unique and unambiguous identifier assigned by the confirmer to the undertaking.
	ConfirmerReferenceNumber *Max35Text `xml:"CnfrmrRefNb,omitempty"`
}

func (u *UndertakingDemandWithdrawal1) AddUndertakingIdentification() *Undertaking6 {
	u.UndertakingIdentification = new(Undertaking6)
	return u.UndertakingIdentification
}

func (u *UndertakingDemandWithdrawal1) SetAdvisingPartyReferenceNumber(value string) {
	u.AdvisingPartyReferenceNumber = (*Max35Text)(&value)
}

func (u *UndertakingDemandWithdrawal1) AddDemandDetails() *Demand3 {
	u.DemandDetails = new(Demand3)
	return u.DemandDetails
}

func (u *UndertakingDemandWithdrawal1) SetConfirmerReferenceNumber(value string) {
	u.ConfirmerReferenceNumber = (*Max35Text)(&value)
}

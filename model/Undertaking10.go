package model

// Details related to the undertaking.
type Undertaking10 struct {

	// Details related to the requested new amount for the counter-undertaking.
	NewUndertakingAmount *UndertakingAmount2 `xml:"NewUdrtkgAmt,omitempty"`

	// Details related to the requested new expiry terms for the counter-undertaking.
	NewExpiryDetails *ExpiryDetails1 `xml:"NewXpryDtls,omitempty"`
}

func (u *Undertaking10) AddNewUndertakingAmount() *UndertakingAmount2 {
	u.NewUndertakingAmount = new(UndertakingAmount2)
	return u.NewUndertakingAmount
}

func (u *Undertaking10) AddNewExpiryDetails() *ExpiryDetails1 {
	u.NewExpiryDetails = new(ExpiryDetails1)
	return u.NewExpiryDetails
}

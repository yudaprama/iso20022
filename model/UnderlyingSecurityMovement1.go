package model

// Provides information about the underlying securities movement.
type UnderlyingSecurityMovement1 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Quantity of financial instrument.
	SecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"SctiesQty"`

	// Provides information about the debited/credited securities account.
	AccountDetails []*SecuritiesAccount8 `xml:"AcctDtls"`
}

func (u *UnderlyingSecurityMovement1) AddSecurityIdentification() *SecurityIdentification7 {
	u.SecurityIdentification = new(SecurityIdentification7)
	return u.SecurityIdentification
}

func (u *UnderlyingSecurityMovement1) AddSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	u.SecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return u.SecuritiesQuantity
}

func (u *UnderlyingSecurityMovement1) AddAccountDetails() *SecuritiesAccount8 {
	newValue := new(SecuritiesAccount8)
	u.AccountDetails = append(u.AccountDetails, newValue)
	return newValue
}

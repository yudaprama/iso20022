package model

// Provides information about the securities movement.
type SecurityMovement1 struct {

	// Identification of the movement.
	MovementIdentification *Max35Text `xml:"MvmntId,omitempty"`

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Quantitty of financial instrument.
	SecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"SctiesQty"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*SecuritiesAccount12 `xml:"AcctDtls"`
}

func (s *SecurityMovement1) SetMovementIdentification(value string) {
	s.MovementIdentification = (*Max35Text)(&value)
}

func (s *SecurityMovement1) AddSecurityIdentification() *SecurityIdentification7 {
	s.SecurityIdentification = new(SecurityIdentification7)
	return s.SecurityIdentification
}

func (s *SecurityMovement1) AddSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.SecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.SecuritiesQuantity
}

func (s *SecurityMovement1) AddAccountDetails() *SecuritiesAccount12 {
	newValue := new(SecuritiesAccount12)
	s.AccountDetails = append(s.AccountDetails, newValue)
	return newValue
}

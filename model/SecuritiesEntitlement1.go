package model

// Provides entitlement information.
type SecuritiesEntitlement1 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"EntitldSctiesQty"`
}

func (s *SecuritiesEntitlement1) AddSecurityIdentification() *SecurityIdentification7 {
	s.SecurityIdentification = new(SecurityIdentification7)
	return s.SecurityIdentification
}

func (s *SecuritiesEntitlement1) AddEntitledSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.EntitledSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.EntitledSecuritiesQuantity
}

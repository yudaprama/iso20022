package model

// Sale context in which the transaction is performed.
type SaleContext1 struct {

	// Identification of the sale terminal (electronic cash register) or the sale system.
	SaleIdentification *Max35Text `xml:"SaleId,omitempty"`

	// Identify a sale transaction assigned by the sale system.
	SaleReferenceNumber *Max35Text `xml:"SaleRefNb,omitempty"`

	// Identifier of the reconciliation between the Sale system and the POI system.
	SaleReconciliationIdentification *Max35Text `xml:"SaleRcncltnId,omitempty"`

	// Identification of the cashier who carried out the transaction.
	CashierIdentification *Max35Text `xml:"CshrId,omitempty"`

	// Identifies the shift of the cashier.
	ShiftNumber *Max2NumericText `xml:"ShftNb,omitempty"`

	// Additional information associated with the sale transaction.
	AdditionalSaleData *Max70Text `xml:"AddtlSaleData,omitempty"`
}

func (s *SaleContext1) SetSaleIdentification(value string) {
	s.SaleIdentification = (*Max35Text)(&value)
}

func (s *SaleContext1) SetSaleReferenceNumber(value string) {
	s.SaleReferenceNumber = (*Max35Text)(&value)
}

func (s *SaleContext1) SetSaleReconciliationIdentification(value string) {
	s.SaleReconciliationIdentification = (*Max35Text)(&value)
}

func (s *SaleContext1) SetCashierIdentification(value string) {
	s.CashierIdentification = (*Max35Text)(&value)
}

func (s *SaleContext1) SetShiftNumber(value string) {
	s.ShiftNumber = (*Max2NumericText)(&value)
}

func (s *SaleContext1) SetAdditionalSaleData(value string) {
	s.AdditionalSaleData = (*Max70Text)(&value)
}

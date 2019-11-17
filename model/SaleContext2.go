package model

// Sale context in which the transaction is performed.
type SaleContext2 struct {

	// Identification of the sale terminal (electronic cash register or point of sale terminal) or the sale system.
	SaleIdentification *Max35Text `xml:"SaleId,omitempty"`

	// Identify a sale transaction assigned by the sale system.
	SaleReferenceNumber *Max35Text `xml:"SaleRefNb,omitempty"`

	// Identifier of the reconciliation between the Sale system and the POI system.
	SaleReconciliationIdentification *Max35Text `xml:"SaleRcncltnId,omitempty"`

	// Identification of the cashier who carried out the transaction.
	CashierIdentification *Max35Text `xml:"CshrId,omitempty"`

	// Identifies the shift of the cashier.
	ShiftNumber *Max2NumericText `xml:"ShftNb,omitempty"`

	// Identification of the purchase order.
	PurchaseOrderNumber *Max35Text `xml:"PurchsOrdrNb,omitempty"`

	// Identification of the invoice.
	InvoiceNumber *Max35Text `xml:"InvcNb,omitempty"`

	// Identification allocated by the sale system and given to the customer.
	DeliveryNoteNumber *Max35Text `xml:"DlvryNoteNb,omitempty"`

	// Merchant using the payment services of a payment facilitator, acting as a card acceptor.
	SponsoredMerchant []*Organisation26 `xml:"SpnsrdMrchnt,omitempty"`

	// True if the payment transaction is a partial payment of the sale transaction.
	SplitPayment *TrueFalseIndicator `xml:"SpltPmt,omitempty"`

	// Remaining amount to complete the sale transaction, if a partial payment has been completed for the sale transaction.
	RemainingAmount *ImpliedCurrencyAndAmount `xml:"RmngAmt,omitempty"`

	// Additional information associated with the sale transaction.
	AdditionalSaleData *Max70Text `xml:"AddtlSaleData,omitempty"`
}

func (s *SaleContext2) SetSaleIdentification(value string) {
	s.SaleIdentification = (*Max35Text)(&value)
}

func (s *SaleContext2) SetSaleReferenceNumber(value string) {
	s.SaleReferenceNumber = (*Max35Text)(&value)
}

func (s *SaleContext2) SetSaleReconciliationIdentification(value string) {
	s.SaleReconciliationIdentification = (*Max35Text)(&value)
}

func (s *SaleContext2) SetCashierIdentification(value string) {
	s.CashierIdentification = (*Max35Text)(&value)
}

func (s *SaleContext2) SetShiftNumber(value string) {
	s.ShiftNumber = (*Max2NumericText)(&value)
}

func (s *SaleContext2) SetPurchaseOrderNumber(value string) {
	s.PurchaseOrderNumber = (*Max35Text)(&value)
}

func (s *SaleContext2) SetInvoiceNumber(value string) {
	s.InvoiceNumber = (*Max35Text)(&value)
}

func (s *SaleContext2) SetDeliveryNoteNumber(value string) {
	s.DeliveryNoteNumber = (*Max35Text)(&value)
}

func (s *SaleContext2) AddSponsoredMerchant() *Organisation26 {
	newValue := new(Organisation26)
	s.SponsoredMerchant = append(s.SponsoredMerchant, newValue)
	return newValue
}

func (s *SaleContext2) SetSplitPayment(value string) {
	s.SplitPayment = (*TrueFalseIndicator)(&value)
}

func (s *SaleContext2) SetRemainingAmount(value, currency string) {
	s.RemainingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (s *SaleContext2) SetAdditionalSaleData(value string) {
	s.AdditionalSaleData = (*Max70Text)(&value)
}

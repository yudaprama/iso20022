package model

// Provides information about  a movement that failed the settlement.
type FailedMovement1 struct {

	// Amount of cash.
	CashAmount *ActiveCurrencyAndAmount `xml:"CshAmt"`

	// Quantity of the financial instrument.
	SecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"SctiesQty"`

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification7 `xml:"SctyId,omitempty"`

	// The reason for the settlement failure.
	Reason *FailedSettlementReason1FormatChoice `xml:"Rsn"`
}

func (f *FailedMovement1) SetCashAmount(value, currency string) {
	f.CashAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FailedMovement1) AddSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	f.SecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return f.SecuritiesQuantity
}

func (f *FailedMovement1) AddSecurityIdentification() *SecurityIdentification7 {
	f.SecurityIdentification = new(SecurityIdentification7)
	return f.SecurityIdentification
}

func (f *FailedMovement1) AddReason() *FailedSettlementReason1FormatChoice {
	f.Reason = new(FailedSettlementReason1FormatChoice)
	return f.Reason
}

package model

// Reporting per financial instrument.
type FinancialInstrumentDetails23 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation16 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance4 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance4 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction50 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails23) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails23) AddPriceDetails() *PriceInformation16 {
	f.PriceDetails = new(PriceInformation16)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails23) AddSafekeepingPlace() *SafeKeepingPlace2 {
	f.SafekeepingPlace = new(SafeKeepingPlace2)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails23) AddOpeningBalance() *OpeningBalance4 {
	f.OpeningBalance = new(OpeningBalance4)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails23) AddClosingBalance() *ClosingBalance4 {
	f.ClosingBalance = new(ClosingBalance4)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails23) AddTransaction() *Transaction50 {
	newValue := new(Transaction50)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

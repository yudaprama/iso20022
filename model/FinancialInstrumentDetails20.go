package model

// Reporting per financial instrument.
type FinancialInstrumentDetails20 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Information regarding the price of the instrument.
	PriceDetails *PriceInformation13 `xml:"PricDtls,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Opening balance for the statement period (first opening balance) or of this page (intermediary opening balance).
	OpeningBalance *OpeningBalance3 `xml:"OpngBal,omitempty"`

	// Closing balance for the statement period (final closing balance) or of this page (intermediary closing balance).
	ClosingBalance *ClosingBalance3 `xml:"ClsgBal,omitempty"`

	// Transaction details.
	Transaction []*Transaction46 `xml:"Tx"`
}

func (f *FinancialInstrumentDetails20) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails20) AddPriceDetails() *PriceInformation13 {
	f.PriceDetails = new(PriceInformation13)
	return f.PriceDetails
}

func (f *FinancialInstrumentDetails20) AddSafekeepingPlace() *SafeKeepingPlace1 {
	f.SafekeepingPlace = new(SafeKeepingPlace1)
	return f.SafekeepingPlace
}

func (f *FinancialInstrumentDetails20) AddOpeningBalance() *OpeningBalance3 {
	f.OpeningBalance = new(OpeningBalance3)
	return f.OpeningBalance
}

func (f *FinancialInstrumentDetails20) AddClosingBalance() *ClosingBalance3 {
	f.ClosingBalance = new(ClosingBalance3)
	return f.ClosingBalance
}

func (f *FinancialInstrumentDetails20) AddTransaction() *Transaction46 {
	newValue := new(Transaction46)
	f.Transaction = append(f.Transaction, newValue)
	return newValue
}

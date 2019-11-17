package model

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation4 struct {

	// Party on the credit side of the transaction to which the tax applies.
	Creditor *TaxParty1 `xml:"Cdtr,omitempty"`

	// Identifies the party on the debit side of the transaction to which the tax applies.
	Debtor *TaxParty2 `xml:"Dbtr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor, in this case, to the taxing authority.
	UltimateDebtor *TaxParty2 `xml:"UltmtDbtr,omitempty"`

	// Territorial part of a country to which the tax payment is related.
	AdministrationZone *Max35Text `xml:"AdmstnZone,omitempty"`

	// Tax reference information that is specific to a taxing agency.
	ReferenceNumber *Max140Text `xml:"RefNb,omitempty"`

	// Method used to indicate the underlying business or how the tax is paid.
	Method *Max35Text `xml:"Mtd,omitempty"`

	// Total amount of money on which the tax is based.
	TotalTaxableBaseAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`

	// Total amount of money as result of the calculation of the tax.
	TotalTaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`

	// Date by which tax is due.
	Date *ISODate `xml:"Dt,omitempty"`

	// Sequential number of the tax report.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Record of tax details.
	Record []*TaxRecord1 `xml:"Rcrd,omitempty"`
}

func (t *TaxInformation4) AddCreditor() *TaxParty1 {
	t.Creditor = new(TaxParty1)
	return t.Creditor
}

func (t *TaxInformation4) AddDebtor() *TaxParty2 {
	t.Debtor = new(TaxParty2)
	return t.Debtor
}

func (t *TaxInformation4) AddUltimateDebtor() *TaxParty2 {
	t.UltimateDebtor = new(TaxParty2)
	return t.UltimateDebtor
}

func (t *TaxInformation4) SetAdministrationZone(value string) {
	t.AdministrationZone = (*Max35Text)(&value)
}

func (t *TaxInformation4) SetReferenceNumber(value string) {
	t.ReferenceNumber = (*Max140Text)(&value)
}

func (t *TaxInformation4) SetMethod(value string) {
	t.Method = (*Max35Text)(&value)
}

func (t *TaxInformation4) SetTotalTaxableBaseAmount(value, currency string) {
	t.TotalTaxableBaseAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation4) SetTotalTaxAmount(value, currency string) {
	t.TotalTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation4) SetDate(value string) {
	t.Date = (*ISODate)(&value)
}

func (t *TaxInformation4) SetSequenceNumber(value string) {
	t.SequenceNumber = (*Number)(&value)
}

func (t *TaxInformation4) AddRecord() *TaxRecord1 {
	newValue := new(TaxRecord1)
	t.Record = append(t.Record, newValue)
	return newValue
}

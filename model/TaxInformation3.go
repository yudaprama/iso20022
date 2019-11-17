package model

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation3 struct {

	// Party on the credit side of the transaction to which the tax applies.
	Creditor *TaxParty1 `xml:"Cdtr,omitempty"`

	// Set of elements used to identify the party on the debit side of the transaction to which the tax applies.
	Debtor *TaxParty2 `xml:"Dbtr,omitempty"`

	// Territorial part of a country to which the tax payment is related.
	AdministrationZone *Max35Text `xml:"AdmstnZn,omitempty"`

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

func (t *TaxInformation3) AddCreditor() *TaxParty1 {
	t.Creditor = new(TaxParty1)
	return t.Creditor
}

func (t *TaxInformation3) AddDebtor() *TaxParty2 {
	t.Debtor = new(TaxParty2)
	return t.Debtor
}

func (t *TaxInformation3) SetAdministrationZone(value string) {
	t.AdministrationZone = (*Max35Text)(&value)
}

func (t *TaxInformation3) SetReferenceNumber(value string) {
	t.ReferenceNumber = (*Max140Text)(&value)
}

func (t *TaxInformation3) SetMethod(value string) {
	t.Method = (*Max35Text)(&value)
}

func (t *TaxInformation3) SetTotalTaxableBaseAmount(value, currency string) {
	t.TotalTaxableBaseAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation3) SetTotalTaxAmount(value, currency string) {
	t.TotalTaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation3) SetDate(value string) {
	t.Date = (*ISODate)(&value)
}

func (t *TaxInformation3) SetSequenceNumber(value string) {
	t.SequenceNumber = (*Number)(&value)
}

func (t *TaxInformation3) AddRecord() *TaxRecord1 {
	newValue := new(TaxRecord1)
	t.Record = append(t.Record, newValue)
	return newValue
}

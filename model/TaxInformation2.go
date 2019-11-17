package model

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation2 struct {

	// Tax Identification Number of the creditor.
	CreditorTaxIdentification *Max35Text `xml:"CdtrTaxId,omitempty"`

	// Type of tax payer (creditor).
	CreditorTaxType *Max35Text `xml:"CdtrTaxTp,omitempty"`

	// Tax Identification Number of the debtor.
	DebtorTaxIdentification *Max35Text `xml:"DbtrTaxId,omitempty"`

	// Tax reference information that is specific to a taxing agency.
	TaxReferenceNumber *Max140Text `xml:"TaxRefNb,omitempty"`

	// Total amount of money on which the tax is based.
	TotalTaxableBaseAmount *CurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`

	// Amount of money resulting from the calculation of the tax.
	TotalTaxAmount *CurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`

	// Date by which tax is due.
	TaxDate *ISODate `xml:"TaxDt,omitempty"`

	// Set of characteristics defining the type of tax.
	TaxTypeInformation []*TaxDetails `xml:"TaxTpInf,omitempty"`
}

func (t *TaxInformation2) SetCreditorTaxIdentification(value string) {
	t.CreditorTaxIdentification = (*Max35Text)(&value)
}

func (t *TaxInformation2) SetCreditorTaxType(value string) {
	t.CreditorTaxType = (*Max35Text)(&value)
}

func (t *TaxInformation2) SetDebtorTaxIdentification(value string) {
	t.DebtorTaxIdentification = (*Max35Text)(&value)
}

func (t *TaxInformation2) SetTaxReferenceNumber(value string) {
	t.TaxReferenceNumber = (*Max140Text)(&value)
}

func (t *TaxInformation2) SetTotalTaxableBaseAmount(value, currency string) {
	t.TotalTaxableBaseAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TaxInformation2) SetTotalTaxAmount(value, currency string) {
	t.TotalTaxAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TaxInformation2) SetTaxDate(value string) {
	t.TaxDate = (*ISODate)(&value)
}

func (t *TaxInformation2) AddTaxTypeInformation() *TaxDetails {
	newValue := new(TaxDetails)
	t.TaxTypeInformation = append(t.TaxTypeInformation, newValue)
	return newValue
}

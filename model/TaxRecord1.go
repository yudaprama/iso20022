package model

// Set of elements used to define the tax record.
type TaxRecord1 struct {

	// High level code to identify the type of tax details.
	Type *Max35Text `xml:"Tp,omitempty"`

	// Specifies the tax code as published by the tax authority.
	Category *Max35Text `xml:"Ctgy,omitempty"`

	// Provides further details of the category tax code.
	CategoryDetails *Max35Text `xml:"CtgyDtls,omitempty"`

	// Code provided by local authority to identify the status of the party that has drawn up the settlement document.
	DebtorStatus *Max35Text `xml:"DbtrSts,omitempty"`

	// Identification number of the tax report as assigned by the taxing authority.
	CertificateIdentification *Max35Text `xml:"CertId,omitempty"`

	// Identifies, in a coded form, on which template the tax report is to be provided.
	FormsCode *Max35Text `xml:"FrmsCd,omitempty"`

	// Set of elements used to provide details on the period of time related to the tax payment.
	Period *TaxPeriod1 `xml:"Prd,omitempty"`

	// Set of elements used to provide information on the amount of the tax record.
	TaxAmount *TaxAmount1 `xml:"TaxAmt,omitempty"`

	// Further details of the tax record.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxRecord1) SetType(value string) {
	t.Type = (*Max35Text)(&value)
}

func (t *TaxRecord1) SetCategory(value string) {
	t.Category = (*Max35Text)(&value)
}

func (t *TaxRecord1) SetCategoryDetails(value string) {
	t.CategoryDetails = (*Max35Text)(&value)
}

func (t *TaxRecord1) SetDebtorStatus(value string) {
	t.DebtorStatus = (*Max35Text)(&value)
}

func (t *TaxRecord1) SetCertificateIdentification(value string) {
	t.CertificateIdentification = (*Max35Text)(&value)
}

func (t *TaxRecord1) SetFormsCode(value string) {
	t.FormsCode = (*Max35Text)(&value)
}

func (t *TaxRecord1) AddPeriod() *TaxPeriod1 {
	t.Period = new(TaxPeriod1)
	return t.Period
}

func (t *TaxRecord1) AddTaxAmount() *TaxAmount1 {
	t.TaxAmount = new(TaxAmount1)
	return t.TaxAmount
}

func (t *TaxRecord1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max140Text)(&value)
}

package model

// Specifies information about a financing relation between two parties represented by a document, for example invoice, credit.
// The component may include an external document describing details of the underlying trade object using an external schema.
type FinancialItem1 struct {

	// Parameters identifying the context of the item.
	ItemContext *FinancialItemParameters1 `xml:"ItmCntxt"`

	// Identifier of financial document that is the base document for this item, for example an invoice number.
	FinancialDocumentReference []*QualifiedDocumentInformation1 `xml:"FinDocRef,omitempty"`

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the total amount related to the financial document.
	TotalAmount *InvoiceTotals1 `xml:"TtlAmt"`

	// Specifies the remaining monetary amount.
	DueAmount *ActiveCurrencyAndAmount `xml:"DueAmt,omitempty"`

	// Instalment information for payment.
	InstalmentInformation []*Instalment2 `xml:"InstlmtInf,omitempty"`

	// Additional proprietary textual information concerning the item.
	AdditionalInformation *Max2000Text `xml:"AddtlInf,omitempty"`

	// Associated free form document, for example a delivery confirmation.
	AssociatedDocument []*QualifiedDocumentInformation1 `xml:"AssoctdDoc,omitempty"`

	// Validation status of the item.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf,omitempty"`

	// Financing status if applicable for the item.
	FinancingStatus *FinancingInformationAndStatus1 `xml:"FincgSts,omitempty"`

	// Structured proprietary information concerning details of the financial item.
	ProprietaryDetails *SupplementaryData1 `xml:"PrtryDtls,omitempty"`
}

func (f *FinancialItem1) AddItemContext() *FinancialItemParameters1 {
	f.ItemContext = new(FinancialItemParameters1)
	return f.ItemContext
}

func (f *FinancialItem1) AddFinancialDocumentReference() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.FinancialDocumentReference = append(f.FinancialDocumentReference, newValue)
	return newValue
}

func (f *FinancialItem1) SetCreditDebitIndicator(value string) {
	f.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (f *FinancialItem1) AddTotalAmount() *InvoiceTotals1 {
	f.TotalAmount = new(InvoiceTotals1)
	return f.TotalAmount
}

func (f *FinancialItem1) SetDueAmount(value, currency string) {
	f.DueAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialItem1) AddInstalmentInformation() *Instalment2 {
	newValue := new(Instalment2)
	f.InstalmentInformation = append(f.InstalmentInformation, newValue)
	return newValue
}

func (f *FinancialItem1) SetAdditionalInformation(value string) {
	f.AdditionalInformation = (*Max2000Text)(&value)
}

func (f *FinancialItem1) AddAssociatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.AssociatedDocument = append(f.AssociatedDocument, newValue)
	return newValue
}

func (f *FinancialItem1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	f.ValidationStatusInformation = new(ValidationStatusInformation1)
	return f.ValidationStatusInformation
}

func (f *FinancialItem1) AddFinancingStatus() *FinancingInformationAndStatus1 {
	f.FinancingStatus = new(FinancingInformationAndStatus1)
	return f.FinancingStatus
}

func (f *FinancialItem1) AddProprietaryDetails() *SupplementaryData1 {
	f.ProprietaryDetails = new(SupplementaryData1)
	return f.ProprietaryDetails
}

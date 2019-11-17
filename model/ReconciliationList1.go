package model

// Specifies a list of reconciliation information concerning financial items.
type ReconciliationList1 struct {

	// Date of creation of this document.
	Date *ISODate `xml:"Dt"`

	// Reference to related documents for example to original assignment in a status response or retry.
	RelatedDocument []*QualifiedDocumentInformation1 `xml:"RltdDoc,omitempty"`

	// Party to be advised.
	Recipient *QualifiedPartyIdentification1 `xml:"Rcpt"`

	// Informing party.
	Advisor *QualifiedPartyIdentification1 `xml:"Advsr"`

	// Identification parameters.
	Parameters *FinancialItemParameters1 `xml:"Params"`

	// Reference to a payment instruction.
	PaymentReference *PaymentIdentification1 `xml:"PmtRef"`

	// Set of elements used to further specify the type of transaction.
	PaymentMeans *PaymentMeans1 `xml:"PmtMeans"`

	// Effective date of payment.
	PaymentDate *ISODate `xml:"PmtDt"`

	// Terms of the payment.
	PaymentTerms *PaymentTerms6 `xml:"PmtTerms"`

	// Amount of the referenced payment.
	PaymentAmount *CurrencyAndAmount `xml:"PmtAmt"`

	// Financial item impacted by the payment.
	Item []*FinancialItem1 `xml:"Itm"`

	// Number of individual items contained in the list.
	ItemCount *Max15NumericText `xml:"ItmCnt"`

	// Total of all individual amounts included in the list, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Associated free form document.
	AssociatedDocument []*QualifiedDocumentInformation1 `xml:"AssoctdDoc,omitempty"`

	// Additional proprietary formal information concerning the list.
	AdditionalInformation *Max2000Text `xml:"AddtlInf,omitempty"`

	// Validation status of the list.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf,omitempty"`
}

func (r *ReconciliationList1) SetDate(value string) {
	r.Date = (*ISODate)(&value)
}

func (r *ReconciliationList1) AddRelatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	r.RelatedDocument = append(r.RelatedDocument, newValue)
	return newValue
}

func (r *ReconciliationList1) AddRecipient() *QualifiedPartyIdentification1 {
	r.Recipient = new(QualifiedPartyIdentification1)
	return r.Recipient
}

func (r *ReconciliationList1) AddAdvisor() *QualifiedPartyIdentification1 {
	r.Advisor = new(QualifiedPartyIdentification1)
	return r.Advisor
}

func (r *ReconciliationList1) AddParameters() *FinancialItemParameters1 {
	r.Parameters = new(FinancialItemParameters1)
	return r.Parameters
}

func (r *ReconciliationList1) AddPaymentReference() *PaymentIdentification1 {
	r.PaymentReference = new(PaymentIdentification1)
	return r.PaymentReference
}

func (r *ReconciliationList1) AddPaymentMeans() *PaymentMeans1 {
	r.PaymentMeans = new(PaymentMeans1)
	return r.PaymentMeans
}

func (r *ReconciliationList1) SetPaymentDate(value string) {
	r.PaymentDate = (*ISODate)(&value)
}

func (r *ReconciliationList1) AddPaymentTerms() *PaymentTerms6 {
	r.PaymentTerms = new(PaymentTerms6)
	return r.PaymentTerms
}

func (r *ReconciliationList1) SetPaymentAmount(value, currency string) {
	r.PaymentAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReconciliationList1) AddItem() *FinancialItem1 {
	newValue := new(FinancialItem1)
	r.Item = append(r.Item, newValue)
	return newValue
}

func (r *ReconciliationList1) SetItemCount(value string) {
	r.ItemCount = (*Max15NumericText)(&value)
}

func (r *ReconciliationList1) SetControlSum(value string) {
	r.ControlSum = (*DecimalNumber)(&value)
}

func (r *ReconciliationList1) AddAssociatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	r.AssociatedDocument = append(r.AssociatedDocument, newValue)
	return newValue
}

func (r *ReconciliationList1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max2000Text)(&value)
}

func (r *ReconciliationList1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	r.ValidationStatusInformation = new(ValidationStatusInformation1)
	return r.ValidationStatusInformation
}

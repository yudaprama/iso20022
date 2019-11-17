package model

// Specifies a list of financing items exchanged between two parties, for example invoice, credit, financing request.
type FinancingItemList1 struct {

	// Identification assigned to unambiguously identify the financing item list.
	Identifier *Max35Text `xml:"Idr"`

	// Date of creation of this document.
	IssueDate *ISODate `xml:"IsseDt"`

	// Reference to related documents for example to original assignment in a status response or retry.
	RelatedDocument []*QualifiedDocumentInformation1 `xml:"RltdDoc,omitempty"`

	// Cut off date for items used to establish the total request amount.
	AmountCutOffDate *ISODate `xml:"AmtCutOffDt,omitempty"`

	// Party to which the list is assigned.
	Assignee *QualifiedPartyIdentification1 `xml:"Assgne"`

	// Party assigning the list.
	Assigner *QualifiedPartyIdentification1 `xml:"Assgnr"`

	// Identifies parties that notify the assignment(s) and the notified parties.
	NotificationInformation []*FinancingNotificationParties1 `xml:"NtfctnInf,omitempty"`

	// List of items/transactions.
	FinancialItem []*FinancialItem1 `xml:"FinItm,omitempty"`

	// Number of individual items contained in the list.
	ItemCount *Max15NumericText `xml:"ItmCnt"`

	// Total of all individual amounts included in the list, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount in all items. Requires same currency, necessary when financing request is in percentage.
	TotalRequestAmount *ActiveCurrencyAndAmount `xml:"TtlReqAmt,omitempty"`

	// Total amount requested.
	TotalRequestFinancing *FinancingRateOrAmountChoice `xml:"TtlReqFincg,omitempty"`

	// Acceptable exchange rate for financing by different currency.
	AgreedRate *AgreedRate1 `xml:"AgrdRate,omitempty"`

	// Instalment for the financing.
	FinancingInstalment []*Instalment2 `xml:"FincgInstlmt,omitempty"`

	// Additional free form information concerning the list.
	AdditionalInformation *Max2000Text `xml:"AddtlInf,omitempty"`

	// Validation status of the list.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf,omitempty"`

	// Financing status if applicable to the nature of the items.
	FinancingStatus *FinancingInformationAndStatus1 `xml:"FincgSts,omitempty"`
}

func (f *FinancingItemList1) SetIdentifier(value string) {
	f.Identifier = (*Max35Text)(&value)
}

func (f *FinancingItemList1) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancingItemList1) AddRelatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.RelatedDocument = append(f.RelatedDocument, newValue)
	return newValue
}

func (f *FinancingItemList1) SetAmountCutOffDate(value string) {
	f.AmountCutOffDate = (*ISODate)(&value)
}

func (f *FinancingItemList1) AddAssignee() *QualifiedPartyIdentification1 {
	f.Assignee = new(QualifiedPartyIdentification1)
	return f.Assignee
}

func (f *FinancingItemList1) AddAssigner() *QualifiedPartyIdentification1 {
	f.Assigner = new(QualifiedPartyIdentification1)
	return f.Assigner
}

func (f *FinancingItemList1) AddNotificationInformation() *FinancingNotificationParties1 {
	newValue := new(FinancingNotificationParties1)
	f.NotificationInformation = append(f.NotificationInformation, newValue)
	return newValue
}

func (f *FinancingItemList1) AddFinancialItem() *FinancialItem1 {
	newValue := new(FinancialItem1)
	f.FinancialItem = append(f.FinancialItem, newValue)
	return newValue
}

func (f *FinancingItemList1) SetItemCount(value string) {
	f.ItemCount = (*Max15NumericText)(&value)
}

func (f *FinancingItemList1) SetControlSum(value string) {
	f.ControlSum = (*DecimalNumber)(&value)
}

func (f *FinancingItemList1) SetTotalRequestAmount(value, currency string) {
	f.TotalRequestAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancingItemList1) AddTotalRequestFinancing() *FinancingRateOrAmountChoice {
	f.TotalRequestFinancing = new(FinancingRateOrAmountChoice)
	return f.TotalRequestFinancing
}

func (f *FinancingItemList1) AddAgreedRate() *AgreedRate1 {
	f.AgreedRate = new(AgreedRate1)
	return f.AgreedRate
}

func (f *FinancingItemList1) AddFinancingInstalment() *Instalment2 {
	newValue := new(Instalment2)
	f.FinancingInstalment = append(f.FinancingInstalment, newValue)
	return newValue
}

func (f *FinancingItemList1) SetAdditionalInformation(value string) {
	f.AdditionalInformation = (*Max2000Text)(&value)
}

func (f *FinancingItemList1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	f.ValidationStatusInformation = new(ValidationStatusInformation1)
	return f.ValidationStatusInformation
}

func (f *FinancingItemList1) AddFinancingStatus() *FinancingInformationAndStatus1 {
	f.FinancingStatus = new(FinancingInformationAndStatus1)
	return f.FinancingStatus
}

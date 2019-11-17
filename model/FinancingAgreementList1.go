package model

// Defines a list of party management registration and guarantee requests.
type FinancingAgreementList1 struct {

	// Identification assigned to unambiguously identify the agreement list.
	Identifier *Max35Text `xml:"Idr"`

	// Creation date of this list.
	Date *ISODate `xml:"Dt"`

	// Reference to related document.
	RelatedDocument []*QualifiedDocumentInformation1 `xml:"RltdDoc,omitempty"`

	// Requestor of the agreement(s).
	AgreementRequestor *QualifiedPartyIdentification1 `xml:"AgrmtRqstr"`

	// Party the agreement(s) are (to be) made with.
	AgreementResponder *QualifiedPartyIdentification1 `xml:"AgrmtRspndr"`

	// Applicant of the guarantee.
	GuaranteeApplicant *QualifiedPartyIdentification1 `xml:"GrntApplcnt"`

	// Beneficiary of the guarantee.
	GuaranteeBeneficiary *QualifiedPartyIdentification1 `xml:"GrntNbfcry"`

	// Party that issues the guarantee.
	GuaranteeIssuer *QualifiedPartyIdentification1 `xml:"GrntIssr"`

	// Party or parties to notify and to acknowledge the agreement.
	NotificationInformation []*FinancingNotificationParties1 `xml:"NtfctnInf,omitempty"`

	// List of agreement items.
	Item []*FinancingAgreementItem1 `xml:"Itm"`

	// Number of individual items contained in the list.
	ItemCount *Max15NumericText `xml:"ItmCnt"`

	// Total of all individual amounts included in the list, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Additional proprietary formal information concerning the list.
	AdditionalInformation *Max2000Text `xml:"AddtlInf,omitempty"`

	// Validation status of the list.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf,omitempty"`
}

func (f *FinancingAgreementList1) SetIdentifier(value string) {
	f.Identifier = (*Max35Text)(&value)
}

func (f *FinancingAgreementList1) SetDate(value string) {
	f.Date = (*ISODate)(&value)
}

func (f *FinancingAgreementList1) AddRelatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.RelatedDocument = append(f.RelatedDocument, newValue)
	return newValue
}

func (f *FinancingAgreementList1) AddAgreementRequestor() *QualifiedPartyIdentification1 {
	f.AgreementRequestor = new(QualifiedPartyIdentification1)
	return f.AgreementRequestor
}

func (f *FinancingAgreementList1) AddAgreementResponder() *QualifiedPartyIdentification1 {
	f.AgreementResponder = new(QualifiedPartyIdentification1)
	return f.AgreementResponder
}

func (f *FinancingAgreementList1) AddGuaranteeApplicant() *QualifiedPartyIdentification1 {
	f.GuaranteeApplicant = new(QualifiedPartyIdentification1)
	return f.GuaranteeApplicant
}

func (f *FinancingAgreementList1) AddGuaranteeBeneficiary() *QualifiedPartyIdentification1 {
	f.GuaranteeBeneficiary = new(QualifiedPartyIdentification1)
	return f.GuaranteeBeneficiary
}

func (f *FinancingAgreementList1) AddGuaranteeIssuer() *QualifiedPartyIdentification1 {
	f.GuaranteeIssuer = new(QualifiedPartyIdentification1)
	return f.GuaranteeIssuer
}

func (f *FinancingAgreementList1) AddNotificationInformation() *FinancingNotificationParties1 {
	newValue := new(FinancingNotificationParties1)
	f.NotificationInformation = append(f.NotificationInformation, newValue)
	return newValue
}

func (f *FinancingAgreementList1) AddItem() *FinancingAgreementItem1 {
	newValue := new(FinancingAgreementItem1)
	f.Item = append(f.Item, newValue)
	return newValue
}

func (f *FinancingAgreementList1) SetItemCount(value string) {
	f.ItemCount = (*Max15NumericText)(&value)
}

func (f *FinancingAgreementList1) SetControlSum(value string) {
	f.ControlSum = (*DecimalNumber)(&value)
}

func (f *FinancingAgreementList1) SetAdditionalInformation(value string) {
	f.AdditionalInformation = (*Max2000Text)(&value)
}

func (f *FinancingAgreementList1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	f.ValidationStatusInformation = new(ValidationStatusInformation1)
	return f.ValidationStatusInformation
}

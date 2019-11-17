package model

// Describes a financing relation between two parties, for example invoice, credit, financing request, cash accounts.
type FinancingAgreementItem1 struct {

	// Parameters related to the context of the item.
	ItemContext *FinancialItemParameters1 `xml:"ItmCntxt"`

	// Code to indicate the action concerning the item.
	ItemAction *AgreementItemAction1Code `xml:"ItmActn,omitempty"`

	// Desired payment instruction to be used by buyer.
	PaymentInstrument *PaymentInstrumentCode `xml:"PmtInstrm,omitempty"`

	// Validation status of the item.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf,omitempty"`

	// Guarantee is (to be) provided according current rating.
	Rating *YesNoIndicator `xml:"Ratg"`

	// Set to yes if the agreement was rejected and needs to be re-opened for arbitrage.
	ReopenIndication *YesNoIndicator `xml:"ReopIndctn"`

	// Issuers, amounts and periods to be guaranteed. At a given date, the sum of all issuers is guaranteed, covered as specified by rank/position and excess. For each period, the maximum value at a given date is used.
	Guarantee []*GuaranteeDetails1 `xml:"Grnt,omitempty"`

	// Status of guarantee if applicable.
	GuaranteeStatus *ValidationStatusInformation1 `xml:"GrntSts,omitempty"`

	// Reference to the guarantee letter issued by a guarantee provider.
	RelatedGuaranteeLetter *QualifiedDocumentInformation1 `xml:"RltdGrntLttr,omitempty"`

	// Associated free form document.
	AssociatedDocument []*QualifiedDocumentInformation1 `xml:"AssoctdDoc,omitempty"`

	// Free form textual information related to the agreement.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (f *FinancingAgreementItem1) AddItemContext() *FinancialItemParameters1 {
	f.ItemContext = new(FinancialItemParameters1)
	return f.ItemContext
}

func (f *FinancingAgreementItem1) SetItemAction(value string) {
	f.ItemAction = (*AgreementItemAction1Code)(&value)
}

func (f *FinancingAgreementItem1) SetPaymentInstrument(value string) {
	f.PaymentInstrument = (*PaymentInstrumentCode)(&value)
}

func (f *FinancingAgreementItem1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	f.ValidationStatusInformation = new(ValidationStatusInformation1)
	return f.ValidationStatusInformation
}

func (f *FinancingAgreementItem1) SetRating(value string) {
	f.Rating = (*YesNoIndicator)(&value)
}

func (f *FinancingAgreementItem1) SetReopenIndication(value string) {
	f.ReopenIndication = (*YesNoIndicator)(&value)
}

func (f *FinancingAgreementItem1) AddGuarantee() *GuaranteeDetails1 {
	newValue := new(GuaranteeDetails1)
	f.Guarantee = append(f.Guarantee, newValue)
	return newValue
}

func (f *FinancingAgreementItem1) AddGuaranteeStatus() *ValidationStatusInformation1 {
	f.GuaranteeStatus = new(ValidationStatusInformation1)
	return f.GuaranteeStatus
}

func (f *FinancingAgreementItem1) AddRelatedGuaranteeLetter() *QualifiedDocumentInformation1 {
	f.RelatedGuaranteeLetter = new(QualifiedDocumentInformation1)
	return f.RelatedGuaranteeLetter
}

func (f *FinancingAgreementItem1) AddAssociatedDocument() *QualifiedDocumentInformation1 {
	newValue := new(QualifiedDocumentInformation1)
	f.AssociatedDocument = append(f.AssociatedDocument, newValue)
	return newValue
}

func (f *FinancingAgreementItem1) AddAdditionalInformation(value string) {
	f.AdditionalInformation = append(f.AdditionalInformation, (*Max2000Text)(&value))
}

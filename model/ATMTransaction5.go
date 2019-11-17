package model

// Withdrawal transaction for which the completion is sent.
type ATMTransaction5 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Outcome of the financial transaction for the customer.
	TransactionStatus *ATMTransactionStatus1Code `xml:"TxSts"`

	// Incident occurring during the processing of the transaction.
	Incident []*FailureReason4Code `xml:"Incdnt,omitempty"`

	// Explanation of the incident.
	IncidentDetail []*Max70Text `xml:"IncdntDtl,omitempty"`

	// Identification of the reconciliation period assigned by the ATM.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// True if the customer has requested a receipt.
	RequestedReceipt *TrueFalseIndicator `xml:"ReqdRct,omitempty"`

	// True if a receipt has been printed and presented to the customer.
	ReceiptPrinted *TrueFalseIndicator `xml:"RctPrtd,omitempty"`

	// True when the card was captured by the ATM.
	CapturedCard *TrueFalseIndicator `xml:"CaptrdCard,omitempty"`

	// Outcome of the withdrawal authorisation.
	AuthorisationResult *AuthorisationResult9 `xml:"AuthstnRslt,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction5) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction5) SetTransactionStatus(value string) {
	a.TransactionStatus = (*ATMTransactionStatus1Code)(&value)
}

func (a *ATMTransaction5) AddIncident(value string) {
	a.Incident = append(a.Incident, (*FailureReason4Code)(&value))
}

func (a *ATMTransaction5) AddIncidentDetail(value string) {
	a.IncidentDetail = append(a.IncidentDetail, (*Max70Text)(&value))
}

func (a *ATMTransaction5) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction5) SetRequestedReceipt(value string) {
	a.RequestedReceipt = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction5) SetReceiptPrinted(value string) {
	a.ReceiptPrinted = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction5) SetCapturedCard(value string) {
	a.CapturedCard = (*TrueFalseIndicator)(&value)
}

func (a *ATMTransaction5) AddAuthorisationResult() *AuthorisationResult9 {
	a.AuthorisationResult = new(AuthorisationResult9)
	return a.AuthorisationResult
}

func (a *ATMTransaction5) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

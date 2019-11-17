package model

// Transaction for which the service is requested.
type ATMTransaction9 struct {

	// Identification of the transaction assigned by the ATM.
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the reconciliation period assigned by the ATM manager.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderNewPIN *OnLinePIN5 `xml:"CrdhldrNewPIN,omitempty"`

	// Sequence of one or more TLV data elements from the ATM application, in accordance with ISO 7816-6, not in a specific order. Present if the transaction is performed with an EMV chip card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (a *ATMTransaction9) AddTransactionIdentification() *TransactionIdentifier1 {
	a.TransactionIdentification = new(TransactionIdentifier1)
	return a.TransactionIdentification
}

func (a *ATMTransaction9) SetReconciliationIdentification(value string) {
	a.ReconciliationIdentification = (*Max35Text)(&value)
}

func (a *ATMTransaction9) AddCardholderNewPIN() *OnLinePIN5 {
	a.CardholderNewPIN = new(OnLinePIN5)
	return a.CardholderNewPIN
}

func (a *ATMTransaction9) SetICCRelatedData(value string) {
	a.ICCRelatedData = (*Max10000Binary)(&value)
}

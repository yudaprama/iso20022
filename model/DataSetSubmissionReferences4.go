package model

// Provides references to the transactions both for the matching application and for the user.
type DataSetSubmissionReferences4 struct {

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Own reference to the transaction for the financial institution.
	UserTransactionReference []*DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Specifies that this message should force the matching application to match all data sets it has received so far for the transaction identified by the transaction identification.
	ForcedMatch *YesNoIndicator `xml:"ForcdMtch"`

	// Unique identification assigned by the matching engine to the baseline when it is established.
	EstablishedBaselineIdentification *DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *BaselineStatus3Code `xml:"TxSts"`
}

func (d *DataSetSubmissionReferences4) SetTransactionIdentification(value string) {
	d.TransactionIdentification = (*Max35Text)(&value)
}

func (d *DataSetSubmissionReferences4) AddPurchaseOrderReference() *DocumentIdentification7 {
	d.PurchaseOrderReference = new(DocumentIdentification7)
	return d.PurchaseOrderReference
}

func (d *DataSetSubmissionReferences4) AddUserTransactionReference() *DocumentIdentification5 {
	newValue := new(DocumentIdentification5)
	d.UserTransactionReference = append(d.UserTransactionReference, newValue)
	return newValue
}

func (d *DataSetSubmissionReferences4) SetForcedMatch(value string) {
	d.ForcedMatch = (*YesNoIndicator)(&value)
}

func (d *DataSetSubmissionReferences4) AddEstablishedBaselineIdentification() *DocumentIdentification3 {
	d.EstablishedBaselineIdentification = new(DocumentIdentification3)
	return d.EstablishedBaselineIdentification
}

func (d *DataSetSubmissionReferences4) SetTransactionStatus(value string) {
	d.TransactionStatus = (*BaselineStatus3Code)(&value)
}

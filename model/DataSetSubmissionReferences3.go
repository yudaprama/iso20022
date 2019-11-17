package model

// Provides references to the submitted data set both for the matching application and for the user.
type DataSetSubmissionReferences3 struct {

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Provides reference to the transaction for the financial institution that submits the data set.
	SubmitterTransactionReference *Max35Text `xml:"SubmitrTxRef,omitempty"`

	// Specifies that this message should force the matching application to match all data sets it has received so far for the transaction identified by the transaction identification.
	ForcedMatch *YesNoIndicator `xml:"ForcdMtch"`
}

func (d *DataSetSubmissionReferences3) SetTransactionIdentification(value string) {
	d.TransactionIdentification = (*Max35Text)(&value)
}

func (d *DataSetSubmissionReferences3) AddPurchaseOrderReference() *DocumentIdentification7 {
	d.PurchaseOrderReference = new(DocumentIdentification7)
	return d.PurchaseOrderReference
}

func (d *DataSetSubmissionReferences3) SetSubmitterTransactionReference(value string) {
	d.SubmitterTransactionReference = (*Max35Text)(&value)
}

func (d *DataSetSubmissionReferences3) SetForcedMatch(value string) {
	d.ForcedMatch = (*YesNoIndicator)(&value)
}

package model

// Describes the events that occurred for one transaction.
type ActivityReportItems2 struct {

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Entity for which the activity is reported.
	ReportedEntity *BICIdentification1 `xml:"RptdNtty"`

	// Describes an activity that took place during a period.
	ReportedItem []*ActivityDetails1 `xml:"RptdItm"`

	// Next processing step required.
	PendingRequestForAction []*PendingActivity2 `xml:"PdgReqForActn,omitempty"`
}

func (a *ActivityReportItems2) SetTransactionIdentification(value string) {
	a.TransactionIdentification = (*Max35Text)(&value)
}

func (a *ActivityReportItems2) AddUserTransactionReference() *DocumentIdentification5 {
	newValue := new(DocumentIdentification5)
	a.UserTransactionReference = append(a.UserTransactionReference, newValue)
	return newValue
}

func (a *ActivityReportItems2) AddReportedEntity() *BICIdentification1 {
	a.ReportedEntity = new(BICIdentification1)
	return a.ReportedEntity
}

func (a *ActivityReportItems2) AddReportedItem() *ActivityDetails1 {
	newValue := new(ActivityDetails1)
	a.ReportedItem = append(a.ReportedItem, newValue)
	return newValue
}

func (a *ActivityReportItems2) AddPendingRequestForAction() *PendingActivity2 {
	newValue := new(PendingActivity2)
	a.PendingRequestForAction = append(a.PendingRequestForAction, newValue)
	return newValue
}

package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04700101 struct {
	XMLName xml.Name           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.047.001.01 Document"`
	Message *SpecialRequestV01 `xml:"SpclReq"`
}

func (d *Document04700101) AddMessage() *SpecialRequestV01 {
	d.Message = new(SpecialRequestV01)
	return d.Message
}

// Scope
// The SpecialRequest message is sent by a party to the transaction to the matching application if special circumstances are such that it cannot take part any longer to a specific transaction or that it cannot fulfill its role in the transaction.
// Usage
// The SpecialRequest message can be sent at any time during the life time of a transaction as long as the transaction is established and not yet closed.
type SpecialRequestV01 struct {

	// Identifies the acceptance message.
	RequestIdentification *model.MessageIdentification1 `xml:"ReqId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the identification of the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Type and details of the notification.
	Notification *model.Notification1 `xml:"Ntfctn"`
}

func (s *SpecialRequestV01) AddRequestIdentification() *model.MessageIdentification1 {
	s.RequestIdentification = new(model.MessageIdentification1)
	return s.RequestIdentification
}

func (s *SpecialRequestV01) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	s.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return s.TransactionIdentification
}

func (s *SpecialRequestV01) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	s.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return s.SubmitterTransactionReference
}

func (s *SpecialRequestV01) AddNotification() *model.Notification1 {
	s.Notification = new(model.Notification1)
	return s.Notification
}

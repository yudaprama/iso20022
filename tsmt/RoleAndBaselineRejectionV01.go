package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05000101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.050.001.01 Document"`
	Message *RoleAndBaselineRejectionV01 `xml:"RoleAndBaselnRjctn"`
}

func (d *Document05000101) AddMessage() *RoleAndBaselineRejectionV01 {
	d.Message = new(RoleAndBaselineRejectionV01)
	return d.Message
}

// Scope
// The RoleAndBaselineRejection message is sent by a secondary bank to the matching application if it rejects to join the transaction based on the baseline and the role that it is expected to play.
// Usage
// The RoleAndBaselineRejection message is sent in response to a message that is a direct request to join a transaction.
type RoleAndBaselineRejectionV01 struct {

	// Identifies the rejection message.
	RejectionIdentification *model.MessageIdentification1 `xml:"RjctnId"`

	// Reference to the message that contained the baseline and is rejected.
	RelatedMessageReference *model.MessageIdentification1 `xml:"RltdMsgRef"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reason why the user cannot accept the request.
	RejectionReason *model.Reason2 `xml:"RjctnRsn,omitempty"`
}

func (r *RoleAndBaselineRejectionV01) AddRejectionIdentification() *model.MessageIdentification1 {
	r.RejectionIdentification = new(model.MessageIdentification1)
	return r.RejectionIdentification
}

func (r *RoleAndBaselineRejectionV01) AddRelatedMessageReference() *model.MessageIdentification1 {
	r.RelatedMessageReference = new(model.MessageIdentification1)
	return r.RelatedMessageReference
}

func (r *RoleAndBaselineRejectionV01) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	r.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return r.TransactionIdentification
}

func (r *RoleAndBaselineRejectionV01) AddRejectionReason() *model.Reason2 {
	r.RejectionReason = new(model.Reason2)
	return r.RejectionReason
}

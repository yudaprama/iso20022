package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04900101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.049.001.01 Document"`
	Message *RoleAndBaselineAcceptanceV01 `xml:"RoleAndBaselnAccptnc"`
}

func (d *Document04900101) AddMessage() *RoleAndBaselineAcceptanceV01 {
	d.Message = new(RoleAndBaselineAcceptanceV01)
	return d.Message
}

// Scope
// The RoleAndBaselineAcceptance message is sent by a secondary bank to the matching application if it accepts to join the transaction based on the baseline and the role that it is expected to play.
// Usage
// The RoleAndBaselineAcceptance message is sent in response to a message that is a direct request to join a transaction.
type RoleAndBaselineAcceptanceV01 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *model.MessageIdentification1 `xml:"AccptncId"`

	// Reference to the message that contained the baseline and is accepted.
	RelatedMessageReference *model.MessageIdentification1 `xml:"RltdMsgRef"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`
}

func (r *RoleAndBaselineAcceptanceV01) AddAcceptanceIdentification() *model.MessageIdentification1 {
	r.AcceptanceIdentification = new(model.MessageIdentification1)
	return r.AcceptanceIdentification
}

func (r *RoleAndBaselineAcceptanceV01) AddRelatedMessageReference() *model.MessageIdentification1 {
	r.RelatedMessageReference = new(model.MessageIdentification1)
	return r.RelatedMessageReference
}

func (r *RoleAndBaselineAcceptanceV01) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	r.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return r.TransactionIdentification
}

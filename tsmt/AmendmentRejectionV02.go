package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.007.001.02 Document"`
	Message *AmendmentRejectionV02 `xml:"AmdmntRjctn"`
}

func (d *Document00700102) AddMessage() *AmendmentRejectionV02 {
	d.Message = new(AmendmentRejectionV02)
	return d.Message
}

// Scope
// The AmendmentRejection message is sent by the party requested to accept or reject an amendment to the matching application.
// This message is used to reject an amendment request.
// Usage
// The AmendmentRejection message can be sent by the party requested to accept or reject an amendment to inform that it rejects the requested amendment.
// The message can be sent in response to a FullPushThroughReport and DeltaReport message conveying the details of a BaselineAmendmentRequest message.
// The acceptance of an amendment request can be achieved by sending an AmendmentAcceptance message.
type AmendmentRejectionV02 struct {

	// Identifies the rejection message.
	RejectionIdentification *model.MessageIdentification1 `xml:"RjctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the delta report that contained the amendment.
	DeltaReportReference *model.MessageIdentification1 `xml:"DltaRptRef"`

	// Sequence number of the rejected baseline amendment.
	RejectedAmendmentNumber *model.Count1 `xml:"RjctdAmdmntNb"`

	// Specifies the reaons for rejecting the amendment.
	RejectionReason *model.RejectionReason1Choice `xml:"RjctnRsn"`
}

func (a *AmendmentRejectionV02) AddRejectionIdentification() *model.MessageIdentification1 {
	a.RejectionIdentification = new(model.MessageIdentification1)
	return a.RejectionIdentification
}

func (a *AmendmentRejectionV02) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	a.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AmendmentRejectionV02) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	a.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return a.SubmitterTransactionReference
}

func (a *AmendmentRejectionV02) AddDeltaReportReference() *model.MessageIdentification1 {
	a.DeltaReportReference = new(model.MessageIdentification1)
	return a.DeltaReportReference
}

func (a *AmendmentRejectionV02) AddRejectedAmendmentNumber() *model.Count1 {
	a.RejectedAmendmentNumber = new(model.Count1)
	return a.RejectedAmendmentNumber
}

func (a *AmendmentRejectionV02) AddRejectionReason() *model.RejectionReason1Choice {
	a.RejectionReason = new(model.RejectionReason1Choice)
	return a.RejectionReason
}

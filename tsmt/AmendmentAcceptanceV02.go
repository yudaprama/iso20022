package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.005.001.02 Document"`
	Message *AmendmentAcceptanceV02 `xml:"AmdmntAccptnc"`
}

func (d *Document00500102) AddMessage() *AmendmentAcceptanceV02 {
	d.Message = new(AmendmentAcceptanceV02)
	return d.Message
}

// Scope
// The AmendmentAcceptance message is sent by the party requested to accept or reject an amendment to the matching application.
// The message is used to accept an amendment request.
// Usage
// The AmendmentAcceptance message can be sent by the party requested to accept or reject an amendment to inform that it accepts the requested amendment.
// The message can be sent in response to a FullPushThroughReport and DeltaReport message conveying the details of a BaselineAmendmentRequest message.
// The rejection of an amendment request can be achieved by sending an AmendmentRejection message.
type AmendmentAcceptanceV02 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *model.MessageIdentification1 `xml:"AccptncId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the identification of the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the delta report that contained the amendment.
	DeltaReportReference *model.MessageIdentification1 `xml:"DltaRptRef"`

	// Sequence number of the accepted baseline amendment.
	AcceptedAmendmentNumber *model.Count1 `xml:"AccptdAmdmntNb"`
}

func (a *AmendmentAcceptanceV02) AddAcceptanceIdentification() *model.MessageIdentification1 {
	a.AcceptanceIdentification = new(model.MessageIdentification1)
	return a.AcceptanceIdentification
}

func (a *AmendmentAcceptanceV02) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	a.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AmendmentAcceptanceV02) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	a.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return a.SubmitterTransactionReference
}

func (a *AmendmentAcceptanceV02) AddDeltaReportReference() *model.MessageIdentification1 {
	a.DeltaReportReference = new(model.MessageIdentification1)
	return a.DeltaReportReference
}

func (a *AmendmentAcceptanceV02) AddAcceptedAmendmentNumber() *model.Count1 {
	a.AcceptedAmendmentNumber = new(model.Count1)
	return a.AcceptedAmendmentNumber
}

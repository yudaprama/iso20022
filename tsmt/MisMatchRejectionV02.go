package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200102 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.022.001.02 Document"`
	Message *MisMatchRejectionV02 `xml:"MisMtchRjctn"`
}

func (d *Document02200102) AddMessage() *MisMatchRejectionV02 {
	d.Message = new(MisMatchRejectionV02)
	return d.Message
}

// Scope
// The MisMatchRejection message is sent by the party requested to accept or reject data set mis-matches to the matching application.
// This message is used to reject mis-matches between data sets and the related baseline.
// Usage
// The MisMatchRejection message can be sent by the party requested to accept or reject data set mis-match to inform that it rejects the data set(s).
// The message can be sent in response to a DataSetMatchReport message conveying mis-matches.
// The information about the rejection of the mis-matched data sets will be forwarded by the matching application to the submitter of the data sets by a MisMatchRejectionNotification message.
// The acceptance of mis-matched data sets can be achieved by sending a MisMatchAcceptance message.
type MisMatchRejectionV02 struct {

	// Identifies the rejection message.
	RejectionIdentification *model.MessageIdentification1 `xml:"RjctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *model.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *model.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the report that contained the difference.
	DataSetMatchReportReference *model.MessageIdentification1 `xml:"DataSetMtchRptRef"`

	// Reason why the user cannot accept the request.
	RejectionReason *model.RejectionReason1Choice `xml:"RjctnRsn"`
}

func (m *MisMatchRejectionV02) AddRejectionIdentification() *model.MessageIdentification1 {
	m.RejectionIdentification = new(model.MessageIdentification1)
	return m.RejectionIdentification
}

func (m *MisMatchRejectionV02) AddTransactionIdentification() *model.SimpleIdentificationInformation {
	m.TransactionIdentification = new(model.SimpleIdentificationInformation)
	return m.TransactionIdentification
}

func (m *MisMatchRejectionV02) AddSubmitterTransactionReference() *model.SimpleIdentificationInformation {
	m.SubmitterTransactionReference = new(model.SimpleIdentificationInformation)
	return m.SubmitterTransactionReference
}

func (m *MisMatchRejectionV02) AddDataSetMatchReportReference() *model.MessageIdentification1 {
	m.DataSetMatchReportReference = new(model.MessageIdentification1)
	return m.DataSetMatchReportReference
}

func (m *MisMatchRejectionV02) AddRejectionReason() *model.RejectionReason1Choice {
	m.RejectionReason = new(model.RejectionReason1Choice)
	return m.RejectionReason
}

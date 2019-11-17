package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100105 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.05 Document"`
	Message *TransferInstructionStatusReportV05 `xml:"TrfInstrStsRpt"`
}

func (d *Document01100105) AddMessage() *TransferInstructionStatusReportV05 {
	d.Message = new(TransferInstructionStatusReportV05)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the TransferInstructionStatusReport message to the instructing party, for example, an investment manager or one of its authorised representatives to provide the status of a previously received transfer instruction.
// Usage
// The TransferInstructionStatusReport message is used to report on the status of a transfer in or transfer out instruction. The reference of the transfer instruction for which the status is reported is identified in TransferReference.
// The message identification of the transfer instruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// One of the following statuses can be reported:
// - an accepted status, or,
// - an already executed status, or,
// - a sent to next party status, or,
// - a matched status, or,
// - a settled status, or,
// - a pending settlement status and the reason for the status, or,
// - an unmatched status and the reason for the status, or,
// - an in-repair status and the reason for the status, or,
// - a rejected status and the reason for the status, or,
// - a failed settlement status and the reason for the status, or,
// - a cancelled status and the reason for the status, or,
// - a cancelled status and the reason for the status, or,
// - a cancellation pending status and the reason for the status.
type TransferInstructionStatusReportV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *model.AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Reference to the message or communication that was previously received.
	Reference *model.References49Choice `xml:"Ref,omitempty"`

	// Status of the transfer instruction.
	StatusReport *model.TransferStatusAndReason4 `xml:"StsRpt"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInstructionStatusReportV05) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInstructionStatusReportV05) AddCounterpartyReference() *model.AdditionalReference7 {
	t.CounterpartyReference = new(model.AdditionalReference7)
	return t.CounterpartyReference
}

func (t *TransferInstructionStatusReportV05) AddReference() *model.References49Choice {
	t.Reference = new(model.References49Choice)
	return t.Reference
}

func (t *TransferInstructionStatusReportV05) AddStatusReport() *model.TransferStatusAndReason4 {
	t.StatusReport = new(model.TransferStatusAndReason4)
	return t.StatusReport
}

func (t *TransferInstructionStatusReportV05) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInstructionStatusReportV05) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100102 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.02 Document"`
	Message *TransferInstructionStatusReportV02 `xml:"TrfInstrStsRptV02"`
}

func (d *Document01100102) AddMessage() *TransferInstructionStatusReportV02 {
	d.Message = new(TransferInstructionStatusReportV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the TransferInstructionStatusReport message to the instructing party, eg, an investment manager or one of its authorised representatives to provide the status of a previously received transfer instruction.
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
type TransferInstructionStatusReportV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference []*model.AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Status of the transfer instruction.
	StatusReport *model.TransferStatusAndReason2 `xml:"StsRpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferInstructionStatusReportV02) AddMessageIdentification() *model.MessageIdentification1 {
	t.MessageIdentification = new(model.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInstructionStatusReportV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	t.RelatedReference = append(t.RelatedReference, newValue)
	return newValue
}

func (t *TransferInstructionStatusReportV02) AddOtherReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	t.OtherReference = append(t.OtherReference, newValue)
	return newValue
}

func (t *TransferInstructionStatusReportV02) AddStatusReport() *model.TransferStatusAndReason2 {
	t.StatusReport = new(model.TransferStatusAndReason2)
	return t.StatusReport
}

func (t *TransferInstructionStatusReportV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

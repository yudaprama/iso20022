package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.011.001.01 Document"`
	Message *TransferInstructionStatusReport `xml:"sese.011.001.01"`
}

func (d *Document01100101) AddMessage() *TransferInstructionStatusReport {
	d.Message = new(TransferInstructionStatusReport)
	return d.Message
}

// Scope
// The TransferInstructionStatusReport message is sent by an instructing party to the executing party. The instructing party may be an investor, a transfer agent, or an intermediary, etc. The executing party may be a transfer agent, or an intermediary, etc.
// This message gives the status of a transfer instruction, and can be used from the time the executing party receives the transfer instruction until its execution.
// Usage
// The TransferInstructionStatusReport message is sent by an executing party to the instructing party. The message can be used to report one of the following
// - the status of the transfer instruction (using a code)or
// - the repair status or
// - the unmatched status or
// - the rejection status or
// - the pending settlement status.
// Further information about repair, unmatched, rejected or pending settlement statuses must be specified using either codes or unstructured information.
type TransferInstructionStatusReport struct {

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *model.AdditionalReference2 `xml:"OthrRef"`

	// Status of the transfer instruction.
	StatusReport *model.TransferStatusAndReason `xml:"StsRpt"`
}

func (t *TransferInstructionStatusReport) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInstructionStatusReport) AddOtherReference() *model.AdditionalReference2 {
	t.OtherReference = new(model.AdditionalReference2)
	return t.OtherReference
}

func (t *TransferInstructionStatusReport) AddStatusReport() *model.TransferStatusAndReason {
	t.StatusReport = new(model.TransferStatusAndReason)
	return t.StatusReport
}

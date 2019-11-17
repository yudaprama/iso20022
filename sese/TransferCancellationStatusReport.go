package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.01 Document"`
	Message *TransferCancellationStatusReport `xml:"sese.010.001.01"`
}

func (d *Document01000101) AddMessage() *TransferCancellationStatusReport {
	d.Message = new(TransferCancellationStatusReport)
	return d.Message
}

// Scope
// The TransferCancellationStatusReport message is sent by an executing party to the instructing party.
// The message gives the status of a transfer cancellation instruction that was previously sent by the instructing party.
// Usage
// The TransferCancellationStatusReport message is sent by an executing party to the instructing party. The message can be used to report that either
// - the cancellation has been acted upon or
// - the cancellation is rejected.
// In both cases, the reason must be specified using either a code or unstructured information.
type TransferCancellationStatusReport struct {

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef"`

	// Reference to the linked message sent in a proprietary way or the reference of a system.
	OtherReference *model.AdditionalReference2 `xml:"OthrRef"`

	// Status of the transfer cancellation instruction.
	StatusReport *model.CancellationStatusAndReason `xml:"StsRpt"`
}

func (t *TransferCancellationStatusReport) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	t.RelatedReference = append(t.RelatedReference, newValue)
	return newValue
}

func (t *TransferCancellationStatusReport) AddOtherReference() *model.AdditionalReference2 {
	t.OtherReference = new(model.AdditionalReference2)
	return t.OtherReference
}

func (t *TransferCancellationStatusReport) AddStatusReport() *model.CancellationStatusAndReason {
	t.StatusReport = new(model.CancellationStatusAndReason)
	return t.StatusReport
}

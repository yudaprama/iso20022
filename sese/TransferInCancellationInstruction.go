package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.006.001.01 Document"`
	Message *TransferInCancellationInstruction `xml:"sese.006.001.01"`
}

func (d *Document00600101) AddMessage() *TransferInCancellationInstruction {
	d.Message = new(TransferInCancellationInstruction)
	return d.Message
}

// Scope
// TheTransferInCancellationInstruction message is sent by an instructing party, or an instructing party's designated agent, to the executing party.
// This message is used to request the cancellation of a TransferInInstruction that was previously sent by the instructing party.
// Usage
// TheTransferInCancellationInstruction message is sent by an instructing party to request cancellation of a previously sent TransferInInstruction.
// This message must contain the reference of the message to be cancelled. The message may also contain all the details of the message to be cancelled, but this is not recommended.
type TransferInCancellationInstruction struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// The transfer in request message to cancel.
	TransferInToBeCancelled *model.TransferIn2 `xml:"TrfInToBeCanc,omitempty"`
}

func (t *TransferInCancellationInstruction) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferInCancellationInstruction) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferInCancellationInstruction) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferInCancellationInstruction) AddTransferInToBeCancelled() *model.TransferIn2 {
	t.TransferInToBeCancelled = new(model.TransferIn2)
	return t.TransferInToBeCancelled
}

package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200101 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.002.001.01 Document"`
	Message *TransferOutCancellationInstruction `xml:"sese.002.001.01"`
}

func (d *Document00200101) AddMessage() *TransferOutCancellationInstruction {
	d.Message = new(TransferOutCancellationInstruction)
	return d.Message
}

// Scope
// The TransferOutCancellationInstruction message is sent by an instructing party or an instructing party's designated agent to the executing party.
// This message is used to request the cancellation of a TransferOutInstruction that was previously sent by the instructing party.
// Usage
// The TransferOutCancellationInstruction message is sent by an instructing party to request cancellation of a previously sent TransferOutInstruction.
// This message must contain the reference of the message to be cancelled. The message may also contain all the details of the message to be cancelled, but this is not recommended.
type TransferOutCancellationInstruction struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference2 `xml:"PrvsRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	//
	TransferOutToBeCancelled *model.TransferOut2 `xml:"TrfOutToBeCanc,omitempty"`
}

func (t *TransferOutCancellationInstruction) AddPreviousReference() *model.AdditionalReference2 {
	t.PreviousReference = new(model.AdditionalReference2)
	return t.PreviousReference
}

func (t *TransferOutCancellationInstruction) AddPoolReference() *model.AdditionalReference2 {
	t.PoolReference = new(model.AdditionalReference2)
	return t.PoolReference
}

func (t *TransferOutCancellationInstruction) AddRelatedReference() *model.AdditionalReference2 {
	t.RelatedReference = new(model.AdditionalReference2)
	return t.RelatedReference
}

func (t *TransferOutCancellationInstruction) AddTransferOutToBeCancelled() *model.TransferOut2 {
	t.TransferOutToBeCancelled = new(model.TransferOut2)
	return t.TransferOutToBeCancelled
}

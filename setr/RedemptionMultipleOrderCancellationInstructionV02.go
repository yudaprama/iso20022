package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00500102 struct {
	XMLName xml.Name                                           `xml:"urn:iso:std:iso:20022:tech:xsd:setr.005.001.02 Document"`
	Message *RedemptionMultipleOrderCancellationInstructionV02 `xml:"setr.005.001.02"`
}

func (d *Document00500102) AddMessage() *RedemptionMultipleOrderCancellationInstructionV02 {
	d.Message = new(RedemptionMultipleOrderCancellationInstructionV02)
	return d.Message
}

// Scope
// The RedemptionMultipleOrderCancellationInstruction message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is sent to cancel a previously sent RedemptionMultipleOrder instruction.
// Usage
// The RedemptionMultipleOrderCancellationInstruction message is used to cancel the entire previously sent RedemptionMultipleOrder message and all the individual orders that it contained. There is no amendment, but a cancellation and re-instruct policy.
// This message must contain the reference of the message to be cancelled. This message may also contain all the details of the message to be cancelled, but this is not recommended.
// The deadline and acceptance of a cancellation instruction is subject to a service level agreement (SLA). This cancellation message is a cancellation instruction. There is no automatic acceptance of the cancellation instruction.
// The rejection or acceptance of a cancellation message instruction is made using an OrderCancellationStatusReport message.
type RedemptionMultipleOrderCancellationInstructionV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef"`

	// Common information related to all the orders to be cancelled.
	OrderToBeCancelled *model.RedemptionMultipleOrderInstruction1 `xml:"OrdrToBeCanc,omitempty"`
}

func (r *RedemptionMultipleOrderCancellationInstructionV02) AddMasterReference() *model.AdditionalReference3 {
	r.MasterReference = new(model.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionMultipleOrderCancellationInstructionV02) AddPoolReference() *model.AdditionalReference3 {
	r.PoolReference = new(model.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionMultipleOrderCancellationInstructionV02) AddPreviousReference() *model.AdditionalReference3 {
	r.PreviousReference = new(model.AdditionalReference3)
	return r.PreviousReference
}

func (r *RedemptionMultipleOrderCancellationInstructionV02) AddOrderToBeCancelled() *model.RedemptionMultipleOrderInstruction1 {
	r.OrderToBeCancelled = new(model.RedemptionMultipleOrderInstruction1)
	return r.OrderToBeCancelled
}

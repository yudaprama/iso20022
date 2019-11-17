package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100102 struct {
	XMLName xml.Name                                             `xml:"urn:iso:std:iso:20022:tech:xsd:setr.011.001.02 Document"`
	Message *SubscriptionMultipleOrderCancellationInstructionV02 `xml:"setr.011.001.02"`
}

func (d *Document01100102) AddMessage() *SubscriptionMultipleOrderCancellationInstructionV02 {
	d.Message = new(SubscriptionMultipleOrderCancellationInstructionV02)
	return d.Message
}

// Scope
// The SubscriptionMultipleOrderCancellationInstruction message is sent by an instructing party, eg, an investment manager or its authorised representative , to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is sent to cancel a previously sent SubscriptionMultipleOrder instruction.
// Usage
// The SubscriptionMultipleOrderCancellationInstruction message is used to cancel the entire previously sent SubscriptionMultipleOrder message and all the individual orders that it contained. There is no amendment, but a cancellation and re-instruct policy.
// This message must contain the reference of the message to be cancelled. This message may also contain all the details of the message to be cancelled, but this is not recommended.
// The deadline and acceptance of a cancellation instruction is subject to a service level agreement (SLA). This cancellation message is a cancellation instruction. There is no automatic acceptance of the cancellation instruction.
// The rejection or acceptance of a cancellation message instruction is made using an OrderCancellationStatusReport message.
type SubscriptionMultipleOrderCancellationInstructionV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef"`

	// Common information related to all the orders to be cancelled.
	OrderToBeCancelled *model.SubscriptionMultipleOrderInstruction1 `xml:"OrdrToBeCanc,omitempty"`
}

func (s *SubscriptionMultipleOrderCancellationInstructionV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SubscriptionMultipleOrderCancellationInstructionV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionMultipleOrderCancellationInstructionV02) AddPreviousReference() *model.AdditionalReference3 {
	s.PreviousReference = new(model.AdditionalReference3)
	return s.PreviousReference
}

func (s *SubscriptionMultipleOrderCancellationInstructionV02) AddOrderToBeCancelled() *model.SubscriptionMultipleOrderInstruction1 {
	s.OrderToBeCancelled = new(model.SubscriptionMultipleOrderInstruction1)
	return s.OrderToBeCancelled
}

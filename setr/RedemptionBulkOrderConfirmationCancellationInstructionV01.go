package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05300101 struct {
	XMLName xml.Name                                                   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.053.001.01 Document"`
	Message *RedemptionBulkOrderConfirmationCancellationInstructionV01 `xml:"RedBlkOrdrConfCxlInstrV01"`
}

func (d *Document05300101) AddMessage() *RedemptionBulkOrderConfirmationCancellationInstructionV01 {
	d.Message = new(RedemptionBulkOrderConfirmationCancellationInstructionV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent sends the RedemptionBulkOrderConfirmationCancellationInstruction message to the instructing party, for example, an investment manager or its authorised representative to cancel a previously sent RedemptionBulkOrderConfirmation.
// Usage
// The RedemptionBulkOrderConfirmationCancellationInstruction message is used to cancel one or more previously sent subscription order confirmations. The amendment indicator element is used to specify whether the subscription order confirmation cancellation is to be followed by a RedemptionBulkOrderConfirmationAmendment.
// The RedemptionBulkOrderConfirmationCancellationInstruction message is used to either:
// - cancel an entire RedemptionBulkOrderConfirmation message, that is, all the individual order confirmations that it contained, or,
// - request the cancellation of one or more individual confirmations.
// There are two ways to use the message.
// (1) When the RedemptionBulkOrderConfirmationCancellationInstruction message is used to cancel an entire message, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of all the individual order confirmations listed in the SubscriptionOrderConfirmation message, or,
// - quoting the details of all the individual order confirmations (this includes the OrderReference and DealReference) listed in SubscriptionOrderConfirmation message but this is not recommended.
// The message identification of the RedemptionBulkOrderConfirmation message may also be quoted in PreviousReference.
// It is also possible to instruct the cancellation of an entire confirmation message by quoting its message identification in PreviousReference, but this is not recommended.
// (2) When the RedemptionBulkOrderConfirmationCancellationInstruction message is used to cancel one or more individual order confirmations, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of each individual order confirmation listed in the RedemptionBulkOrderConfirmation message, or,
// - quoting the details of each individual order execution (this includes the OrderReference and DealReference) listed in RedemptionBulkOrderConfirmation message but this is not recommended.
// The message identification of the RedemptionBulkOrderConfirmation message in which the individual order confirmation was conveyed may also be quoted in PreviousReference.
type RedemptionBulkOrderConfirmationCancellationInstructionV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// References of the orders to be cancelled.
	CancellationByReference *model.InvestmentFundOrderExecution1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the orders to be cancelled.
	CancellationByOrderConfirmationDetails *model.OrderConfirmationDetails1 `xml:"CxlByOrdrConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddPoolReference() *model.AdditionalReference3 {
	r.PoolReference = new(model.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddRelatedReference() *model.AdditionalReference3 {
	r.RelatedReference = new(model.AdditionalReference3)
	return r.RelatedReference
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddCancellationByReference() *model.InvestmentFundOrderExecution1 {
	r.CancellationByReference = new(model.InvestmentFundOrderExecution1)
	return r.CancellationByReference
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddCancellationByOrderConfirmationDetails() *model.OrderConfirmationDetails1 {
	r.CancellationByOrderConfirmationDetails = new(model.OrderConfirmationDetails1)
	return r.CancellationByOrderConfirmationDetails
}

func (r *RedemptionBulkOrderConfirmationCancellationInstructionV01) AddCopyDetails() *model.CopyInformation2 {
	r.CopyDetails = new(model.CopyInformation2)
	return r.CopyDetails
}

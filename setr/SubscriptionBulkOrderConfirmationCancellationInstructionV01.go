package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04900101 struct {
	XMLName xml.Name                                                     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.049.001.01 Document"`
	Message *SubscriptionBulkOrderConfirmationCancellationInstructionV01 `xml:"SbcptBlkOrdrConfCxlInstrV01"`
}

func (d *Document04900101) AddMessage() *SubscriptionBulkOrderConfirmationCancellationInstructionV01 {
	d.Message = new(SubscriptionBulkOrderConfirmationCancellationInstructionV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent sends the SubscriptionBulkOrderConfirmationCancellationInstruction message to the instructing party, for example, an investment manager or its authorised representative to cancel a previously sent SubscriptionBulkOrderConfirmation.
// Usage
// The SubscriptionBulkOrderConfirmationCancellationInstruction message is used to cancel one or more previously sent subscription bulk order confirmations.
// The amendment indicator element is used to specify whether the subscription bulk order confirmation cancellation is to be followed by a SubscriptionBulkOrderConfirmationAmendment.
// The SubscriptionBulkOrderConfirmationCancellationInstruction message is used to either:
// - cancel an entire SubscriptionBulkOrderConfirmation message, that is, all the individual order confirmations that it contained, or,
// - request the cancellation of one or more individual confirmations.
// There are two ways to use the message.
// (1) When the SubscriptionBulkOrderConfirmationCancellationInstruction message is used to cancel an entire message, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of all the individual order confirmations listed in the SubscriptionBulkOrderConfirmation message, or,
// - quoting the details of all the individual order confirmations (this includes the OrderReference and DealReference) listed in SubscriptionBulkOrderConfirmation message but this is not recommended.
// The message identification of the SubscriptionOrderConfirmation message may also be quoted in PreviousReference.
// It is also possible to instruct the cancellation of an entire confirmation message by quoting its message identification in PreviousReference, but this is not recommended.
// (2) When the SubscriptionBulkOrderConfirmationCancellationInstruction message is used to cancel one or more individual order confirmations, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of each individual order confirmation listed in the SubscriptionBulkOrderConfirmation message, or,
// - quoting the details of each individual order execution (this includes the OrderReference and DealReference) listed in SubscriptionBulkOrderConfirmation message but this is not recommended.
// The message identification of the SubscriptionBulkOrderConfirmation message in which the individual order confirmation was conveyed may also be quoted in PreviousReference.
// The rejection or acceptance of a SubscriptionBulkOrderConfirmationCancellationInstruction is made using an OrderConfirmationStatusReport message.
type SubscriptionBulkOrderConfirmationCancellationInstructionV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// References of the orders confirmations to be cancelled.
	CancellationByReference *model.InvestmentFundOrderExecution1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the orders confirmations to be cancelled.
	CancellationByOrderConfirmationDetails *model.SubscriptionBulkOrderConfirmation1 `xml:"CxlByOrdrConfDtls,omitempty"`

	// Message is a copy.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddCancellationByReference() *model.InvestmentFundOrderExecution1 {
	s.CancellationByReference = new(model.InvestmentFundOrderExecution1)
	return s.CancellationByReference
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddCancellationByOrderConfirmationDetails() *model.SubscriptionBulkOrderConfirmation1 {
	s.CancellationByOrderConfirmationDetails = new(model.SubscriptionBulkOrderConfirmation1)
	return s.CancellationByOrderConfirmationDetails
}

func (s *SubscriptionBulkOrderConfirmationCancellationInstructionV01) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}

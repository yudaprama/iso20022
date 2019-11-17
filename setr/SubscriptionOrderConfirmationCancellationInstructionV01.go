package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04700101 struct {
	XMLName xml.Name                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.047.001.01 Document"`
	Message *SubscriptionOrderConfirmationCancellationInstructionV01 `xml:"SbcptOrdrConfCxlInstrV01"`
}

func (d *Document04700101) AddMessage() *SubscriptionOrderConfirmationCancellationInstructionV01 {
	d.Message = new(SubscriptionOrderConfirmationCancellationInstructionV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SubscriptionOrderConfirmationCancellationInstruction message to the instructing party, for example, an investment manager or its authorised representative to cancel a previously sent SubscriptionOrderConfirmation.
// Usage
// The SubscriptionOrderConfirmationCancellationInstruction message is used to cancel one or more previously sent subscription order confirmations. The amendment indicator element is used to specify whether the subscription order confirmation cancellation is to be followed by a SubscriptionOrderConfirmationAmendment.
// The SubscriptionOrderConfirmationCancellationInstruction message is used to either:
// - cancel an entire SubscriptionOrderConfirmation message, that is, all the individual order confirmations that it contained, or,
// - request the cancellation of one or more individual confirmations.
// There are two ways to use the message.
// (1) When the SubscriptionOrderConfirmationCancellationInstruction message is used to cancel an entire message, this can be done by either:
// - quoting the business references, for example, OrderReference, Deal Reference, of all the individual order confirmations listed in the SubscriptionOrderConfirmation message, or,
// - quoting the details of all the individual order confirmations (this includes the OrderReference and DealReference) listed in SubscriptionOrderConfirmation message but this is not recommended.
// The message identification of the SubscriptionOrderConfirmation message may also be quoted in PreviousReference.
// It is also possible to instruct the cancellation of an entire confirmation message by quoting its message identification in PreviousReference.
// (2) When the SubscriptionOrderConfirmationCancellationInstruction message is used to cancel one or more individual order confirmations, this can be done by either:
// - quoting the "business" references, for example, OrderReference, Deal Reference, of each individual order confirmation listed in the SubscriptionOrderConfirmation message, or,
// - quoting the details of each individual order execution (this includes the OrderReference and DealReference) listed in SubscriptionOrderConfirmation message but this is not recommended.
// The message identification of the SubscriptionOrderConfirmation message in which the individual order confirmation was conveyed may also be quoted in PreviousReference.
// The rejection or acceptance of a SubscriptionOrderConfirmationCancellationInstruction is made using an OrderConfirmationStatusReport message.
type SubscriptionOrderConfirmationCancellationInstructionV01 struct {

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
	CancellationByOrderConfirmationDetails *model.SubscriptionOrderConfirmation1 `xml:"CxlByOrdrConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddCancellationByReference() *model.InvestmentFundOrderExecution1 {
	s.CancellationByReference = new(model.InvestmentFundOrderExecution1)
	return s.CancellationByReference
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddCancellationByOrderConfirmationDetails() *model.SubscriptionOrderConfirmation1 {
	s.CancellationByOrderConfirmationDetails = new(model.SubscriptionOrderConfirmation1)
	return s.CancellationByOrderConfirmationDetails
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV01) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

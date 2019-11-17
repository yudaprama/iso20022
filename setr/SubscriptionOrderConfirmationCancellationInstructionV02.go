package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04700102 struct {
	XMLName xml.Name                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.047.001.02 Document"`
	Message *SubscriptionOrderConfirmationCancellationInstructionV02 `xml:"SbcptOrdrConfCxlInstr"`
}

func (d *Document04700102) AddMessage() *SubscriptionOrderConfirmationCancellationInstructionV02 {
	d.Message = new(SubscriptionOrderConfirmationCancellationInstructionV02)
	return d.Message
}

// Scope
// The SubscriptionOrderConfirmationCancellationInstruction message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to cancel a previously sent SubscriptionOrderConfirmation.
// Usage
// To request the cancellation of one or more individual order executions, the order reference and deal reference of each individual order execution in the original SubscriptionOrderConfirmation are specified in the order reference and deal reference elements respectively.  The message identification of the SubscriptionOrderConfirmation message in which the individual order execution was conveyed may also be quoted in PreviousReference but this is not recommended. The AmendmentIndicator is used to specify whether the subscription order confirmation cancellation is to be followed by an amendment  An amendment of a subscription order confirmation is carried out by sending a SubscriptionOrderConfirmation message in which the AmendmentIndicator contains the value ‘true’.
type SubscriptionOrderConfirmationCancellationInstructionV02 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *model.YesNoIndicator `xml:"AmdmntInd"`

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the individual order confirmation to be cancelled.
	//
	OrderReferences []*model.InvestmentFundOrder11 `xml:"OrdrRefs"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) AddPoolReference() *model.AdditionalReference9 {
	s.PoolReference = new(model.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) AddRelatedReference() *model.AdditionalReference8 {
	s.RelatedReference = new(model.AdditionalReference8)
	return s.RelatedReference
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) SetAmendmentIndicator(value string) {
	s.AmendmentIndicator = (*model.YesNoIndicator)(&value)
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) SetMasterReference(value string) {
	s.MasterReference = (*model.Max35Text)(&value)
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) AddOrderReferences() *model.InvestmentFundOrder11 {
	newValue := new(model.InvestmentFundOrder11)
	s.OrderReferences = append(s.OrderReferences, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationCancellationInstructionV02) AddCopyDetails() *model.CopyInformation4 {
	s.CopyDetails = new(model.CopyInformation4)
	return s.CopyDetails
}

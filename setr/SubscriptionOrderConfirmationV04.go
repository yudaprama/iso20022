package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.012.001.04 Document"`
	Message *SubscriptionOrderConfirmationV04 `xml:"SbcptOrdrConf"`
}

func (d *Document01200104) AddMessage() *SubscriptionOrderConfirmationV04 {
	d.Message = new(SubscriptionOrderConfirmationV04)
	return d.Message
}

// Scope
// The SubscriptionOrderConfirmation message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to confirm the details of the execution of a SubscriptionOrder instruction.
// Usage
// The SubscriptionOrderConfirmation message is used to confirm the execution of one or more individual orders.
// A SubscriptionOrder message containing more than one individual order may be responded to by more than one SubscriptionOrderConfirmation message, as the valuation cycle of the financial instruments in each individual order may be different. When a SubscriptionOrderConfirmation message contains fewer confirmations that originally instructed in the original SubscriptionOrder message, there is no specification indication in the confirmation for this. Reconciliation must be based on the references.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the SubscriptionOrder message in which the individual orders was conveyed may also be quoted in RelatedReference but this is not recommended.
//
// A SubscriptionOrder must in all cases be responded to by a SubscriptionOrderConfirmation message and in no circumstances by a SubscriptionBulkOrderConfirmation message.
// If the executing party needs to confirm one or more subscription orders for the same financial instrument, then a SubscriptionBulkOrderConfirmation message must be used.
// When the message is used to convey a confirmation amendment/s, the AmendmentIndicator must be present with the value ‘true’ or ‘1’. When this is the case, the message must only contain a confirmation amendment/s and not contain both a confirmation amendment/s and a ‘new’ confirmation/s.
type SubscriptionOrderConfirmationV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// General information related to the execution of the orders.
	MultipleExecutionDetails *model.SubscriptionMultipleExecution5 `xml:"MltplExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionOrderConfirmationV04) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderConfirmationV04) AddPoolReference() *model.AdditionalReference9 {
	s.PoolReference = new(model.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionOrderConfirmationV04) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationV04) AddRelatedReference() *model.AdditionalReference8 {
	s.RelatedReference = new(model.AdditionalReference8)
	return s.RelatedReference
}

func (s *SubscriptionOrderConfirmationV04) AddMultipleExecutionDetails() *model.SubscriptionMultipleExecution5 {
	s.MultipleExecutionDetails = new(model.SubscriptionMultipleExecution5)
	return s.MultipleExecutionDetails
}

func (s *SubscriptionOrderConfirmationV04) AddCopyDetails() *model.CopyInformation4 {
	s.CopyDetails = new(model.CopyInformation4)
	return s.CopyDetails
}

func (s *SubscriptionOrderConfirmationV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

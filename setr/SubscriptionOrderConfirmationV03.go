package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.012.001.03 Document"`
	Message *SubscriptionOrderConfirmationV03 `xml:"SbcptOrdrConfV03"`
}

func (d *Document01200103) AddMessage() *SubscriptionOrderConfirmationV03 {
	d.Message = new(SubscriptionOrderConfirmationV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SubscriptionOrderConfirmation message to the instructing party, for example, an investment manager or its authorised representative to confirm the details of the execution of a SubscriptionOrder instruction.
// Usage
// The SubscriptionOrderConfirmation message is used to confirm the execution of one or more individual orders.
// A SubscriptionOrder message containing more than one individual order may be responded to by more than one SubscriptionOrderConfirmation message, as the valuation cycle of the financial instruments in each individual order may be different.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the SubscriptionOrder message in which the individual orders was conveyed may also be quoted in RelatedReference.
// When the executing party sends several confirmations, there is no specific indication in the message that it is an incomplete confirmation. Reconciliation must be based on the references.
// A SubscriptionOrder must in all cases be responded to by a SubscriptionOrderConfirmation message and in no circumstances by a SubscriptionBulkOrderConfirmation message.
// If the executing party needs to confirm a SubscriptionBulkOrder message, then a SubscriptionBulkOrderConfirmation message must be used.
type SubscriptionOrderConfirmationV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// General information related to the execution of investment fund order.
	MultipleExecutionDetails *model.SubscriptionMultipleExecution3 `xml:"MltplExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*model.Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionOrderConfirmationV03) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderConfirmationV03) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionOrderConfirmationV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationV03) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionOrderConfirmationV03) AddMultipleExecutionDetails() *model.SubscriptionMultipleExecution3 {
	s.MultipleExecutionDetails = new(model.SubscriptionMultipleExecution3)
	return s.MultipleExecutionDetails
}

func (s *SubscriptionOrderConfirmationV03) AddRelatedPartyDetails() *model.Intermediary9 {
	newValue := new(model.Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationV03) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionOrderConfirmationV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

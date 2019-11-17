package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900104 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.009.001.04 Document"`
	Message *SubscriptionBulkOrderConfirmationV04 `xml:"SbcptBlkOrdrConf"`
}

func (d *Document00900104) AddMessage() *SubscriptionBulkOrderConfirmationV04 {
	d.Message = new(SubscriptionBulkOrderConfirmationV04)
	return d.Message
}

// Scope
// The SubscriptionBulkOrderConfirmation message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to confirm the details of the execution of a SubscriptionBulkOrder instruction.
// Usage
// The SubscriptionBulkOrderConfirmation message is used to confirm the execution of all individual orders.
// There is usually one bulk confirmation message for one bulk order message.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the SubscriptionBulkOrder message in which the individual order was conveyed may also be quoted in RelatedReference.
// A SubscriptionBulkOrder must in all cases be responded to by a SubscriptionBulkOrderConfirmation and in no circumstances by a SubscriptionOrderConfirmation.
// If the executing party needs to confirm a SubscriptionOrder instruction, then the SubscriptionOrderConfirmation must be used.
// When the message is used to convey a confirmation amendment/s, the AmendmentIndicator must be present with the value ‘true’ or ‘1’. When this is the case, the message must only contain a confirmation amendment/s and not contain both a confirmation amendment/s and a ‘new’ confirmation/s.
type SubscriptionBulkOrderConfirmationV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// General information related to the execution of the orders.
	BulkExecutionDetails *model.SubscriptionBulkExecution4 `xml:"BlkExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderConfirmationV04) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderConfirmationV04) AddPoolReference() *model.AdditionalReference9 {
	s.PoolReference = new(model.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderConfirmationV04) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationV04) AddRelatedReference() *model.AdditionalReference8 {
	s.RelatedReference = new(model.AdditionalReference8)
	return s.RelatedReference
}

func (s *SubscriptionBulkOrderConfirmationV04) AddBulkExecutionDetails() *model.SubscriptionBulkExecution4 {
	s.BulkExecutionDetails = new(model.SubscriptionBulkExecution4)
	return s.BulkExecutionDetails
}

func (s *SubscriptionBulkOrderConfirmationV04) AddCopyDetails() *model.CopyInformation4 {
	s.CopyDetails = new(model.CopyInformation4)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderConfirmationV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

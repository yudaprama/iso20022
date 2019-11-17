package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200102 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.012.001.02 Document"`
	Message *SubscriptionMultipleOrderConfirmationV02 `xml:"setr.012.001.02"`
}

func (d *Document01200102) AddMessage() *SubscriptionMultipleOrderConfirmationV02 {
	d.Message = new(SubscriptionMultipleOrderConfirmationV02)
	return d.Message
}

// Scope
// The SubscriptionMultipleOrderConfirmation message is sent by an executing party, eg, a transfer agent, to the instructing party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the executing party and the instruction party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to confirm the details of the execution of a SubscriptionMultipleOrder message.
// Usage
// The SubscriptionMultipleOrderConfirmation message is sent, after the price has been determined, to confirm the execution of the individual orders.
// A SubscriptionMultipleOrder may be responded to by more than one SubscriptionMultipleOrderConfirmation message, as the valuation cycle of the financial instruments in each individual order may be different.
// When the executing party sends several confirmations, there is no specific indication in the message that it is an incomplete confirmation. Reconciliation must be based on the references.
// A SubscriptionMultipleOrder must in all cases be responded to by a SubscriptionMultipleOrderConfirmation message/s and in no circumstances by a SubscriptionBulkOrderConfirmation message/s.
// If the executing party needs to confirm a SubscriptionBulkOrder message, then a SubscriptionBulkOrderConfirmation message must be used.
type SubscriptionMultipleOrderConfirmationV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef"`

	// General information related to the execution of investment fund order.
	MultipleExecutionDetails *model.SubscriptionMultipleExecution2 `xml:"MltplExctnDtls"`

	// Confirmation of the information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddMultipleExecutionDetails() *model.SubscriptionMultipleExecution2 {
	s.MultipleExecutionDetails = new(model.SubscriptionMultipleExecution2)
	return s.MultipleExecutionDetails
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

func (s *SubscriptionMultipleOrderConfirmationV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

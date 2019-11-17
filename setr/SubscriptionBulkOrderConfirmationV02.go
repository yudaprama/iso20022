package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.009.001.02 Document"`
	Message *SubscriptionBulkOrderConfirmationV02 `xml:"setr.009.001.02"`
}

func (d *Document00900102) AddMessage() *SubscriptionBulkOrderConfirmationV02 {
	d.Message = new(SubscriptionBulkOrderConfirmationV02)
	return d.Message
}

// Scope
// The SubscriptionBulkOrderConfirmation message is sent by an executing party, eg, a transfer agent, to an instructing party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the executing party and the instructing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to confirm the details of the execution of a SubscriptionBulkOrder message.
// Usage
// The SubscriptionBulkOrderConfirmation message is sent, after the price has been determined, to confirm the execution of all individual orders.
// There is usually one bulk confirmation message for one bulk order message.
// A SubscriptionBulkOrder must in all cases be responded to by a SubscriptionBulkOrderConfirmation and in no circumstances by a SubscriptionMultipleOrderConfirmation.
// If the executing party needs to confirm a SubscriptionMultipleOrder message, then the SubscriptionMultipleOrderConfirmation message must be used.
type SubscriptionBulkOrderConfirmationV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef"`

	// General information related to the execution of investment orders.
	BulkExecutionDetails *model.SubscriptionBulkExecution2 `xml:"BlkExctnDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderConfirmationV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SubscriptionBulkOrderConfirmationV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderConfirmationV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationV02) AddRelatedReference() *model.AdditionalReference3 {
	s.RelatedReference = new(model.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionBulkOrderConfirmationV02) AddBulkExecutionDetails() *model.SubscriptionBulkExecution2 {
	s.BulkExecutionDetails = new(model.SubscriptionBulkExecution2)
	return s.BulkExecutionDetails
}

func (s *SubscriptionBulkOrderConfirmationV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderConfirmationV02) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderConfirmationV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

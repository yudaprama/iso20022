package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000102 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:setr.010.001.02 Document"`
	Message *SubscriptionMultipleOrderV02 `xml:"setr.010.001.02"`
}

func (d *Document01000102) AddMessage() *SubscriptionMultipleOrderV02 {
	d.Message = new(SubscriptionMultipleOrderV02)
	return d.Message
}

// Scope
// The SubscriptionMultipleOrder message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to instruct the executing party to subscribe to one or more financial instruments, for the same account.
// Usage
// The SubscriptionMultipleOrder message is used for multiple orders. It may also be used for single orders, ie, a message containing one order for one financial instrument and related to one investment account. For a single subscription order, the SubscriptionMultipleOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for the same financial instrument but for different accounts, then the SubscriptionBulkOrder message must be used.
type SubscriptionMultipleOrderV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	MultipleOrderDetails *model.SubscriptionMultipleOrder2 `xml:"MltplOrdrDtls"`

	// Information related to the intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionMultipleOrderV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SubscriptionMultipleOrderV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionMultipleOrderV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderV02) AddMultipleOrderDetails() *model.SubscriptionMultipleOrder2 {
	s.MultipleOrderDetails = new(model.SubscriptionMultipleOrder2)
	return s.MultipleOrderDetails
}

func (s *SubscriptionMultipleOrderV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleOrderV02) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

func (s *SubscriptionMultipleOrderV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

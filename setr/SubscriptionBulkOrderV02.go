package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.007.001.02 Document"`
	Message *SubscriptionBulkOrderV02 `xml:"setr.007.001.02"`
}

func (d *Document00700102) AddMessage() *SubscriptionBulkOrderV02 {
	d.Message = new(SubscriptionBulkOrderV02)
	return d.Message
}

// Scope
// The SubscriptionBulkOrder message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to instruct the executing party to subscribe to a financial instrument, for two or more accounts.
// Usage
// The SubscriptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, ie, account owners, but are related to the same financial instrument. This message will be typically be used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single subscription order, the SubscriptionMultipleOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for different financial instruments but for the same account, then the SubscriptionMultipleOrder must be used.
type SubscriptionBulkOrderV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	BulkOrderDetails *model.SubscriptionBulkOrder2 `xml:"BlkOrdrDtls"`

	// The information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderV02) AddMasterReference() *model.AdditionalReference3 {
	s.MasterReference = new(model.AdditionalReference3)
	return s.MasterReference
}

func (s *SubscriptionBulkOrderV02) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV02) AddBulkOrderDetails() *model.SubscriptionBulkOrder2 {
	s.BulkOrderDetails = new(model.SubscriptionBulkOrder2)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV02) AddCopyDetails() *model.CopyInformation1 {
	s.CopyDetails = new(model.CopyInformation1)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

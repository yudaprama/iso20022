package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700104 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.007.001.04 Document"`
	Message *SubscriptionBulkOrderV04 `xml:"SbcptBlkOrdr"`
}

func (d *Document00700104) AddMessage() *SubscriptionBulkOrderV04 {
	d.Message = new(SubscriptionBulkOrderV04)
	return d.Message
}

// Scope
// The SubscriptionBulkOrder message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to instruct a subscription to a financial instrument for two or more accounts.
// Usage
// The SubscriptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are for the same financial instrument. This message will typically be used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
type SubscriptionBulkOrderV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// General information related to the orders.
	BulkOrderDetails *model.SubscriptionBulkOrder5 `xml:"BlkOrdrDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderV04) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderV04) AddPoolReference() *model.AdditionalReference9 {
	s.PoolReference = new(model.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderV04) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV04) AddBulkOrderDetails() *model.SubscriptionBulkOrder5 {
	s.BulkOrderDetails = new(model.SubscriptionBulkOrder5)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderV04) AddCopyDetails() *model.CopyInformation4 {
	s.CopyDetails = new(model.CopyInformation4)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:setr.007.001.03 Document"`
	Message *SubscriptionBulkOrderV03 `xml:"SbcptBlkOrdrV03"`
}

func (d *Document00700103) AddMessage() *SubscriptionBulkOrderV03 {
	d.Message = new(SubscriptionBulkOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative sends the SubscriptionBulkOrder message to the executing party, for example, a transfer agent, to instruct a subscription to a financial instrument for two or more accounts.
// Usage
// The SubscriptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are related to the same financial instrument. This message will typically be used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for different financial instruments but for the same account, then the SubscriptionOrder must be used.
type SubscriptionBulkOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	BulkOrderDetails *model.SubscriptionBulkOrder4 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*model.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionBulkOrderV03) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderV03) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV03) AddBulkOrderDetails() *model.SubscriptionBulkOrder4 {
	s.BulkOrderDetails = new(model.SubscriptionBulkOrder4)
	return s.BulkOrderDetails
}

func (s *SubscriptionBulkOrderV03) AddRelatedPartyDetails() *model.Intermediary8 {
	newValue := new(model.Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderV03) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionBulkOrderV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

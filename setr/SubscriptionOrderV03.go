package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01000103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:setr.010.001.03 Document"`
	Message *SubscriptionOrderV03 `xml:"SbcptOrdrV03"`
}

func (d *Document01000103) AddMessage() *SubscriptionOrderV03 {
	d.Message = new(SubscriptionOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the SubscriptionOrder message to the executing party, for example, a transfer agent, to instruct the subscription of one or more financial instruments for one investment fund account.
// Usage
// The SubscriptionOrder message is used to instruct single subscription orders, that is, a message containing one order for one financial instrument and related to one investment account. The SubscriptionOrder message may also be used for multiple orders, that is, a message containing several orders related to the same investment account for different financial instruments.
// For a single subscription order, the SubscriptionOrder message, not the SubscriptionBulkOrder message, must be used.
// If there are subscription orders for the same financial instrument but for different accounts, then the SubscriptionBulkOrder message must be used.
type SubscriptionOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	MultipleOrderDetails *model.SubscriptionMultipleOrder4 `xml:"MltplOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*model.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionOrderV03) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderV03) AddPoolReference() *model.AdditionalReference3 {
	s.PoolReference = new(model.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionOrderV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderV03) AddMultipleOrderDetails() *model.SubscriptionMultipleOrder4 {
	s.MultipleOrderDetails = new(model.SubscriptionMultipleOrder4)
	return s.MultipleOrderDetails
}

func (s *SubscriptionOrderV03) AddRelatedPartyDetails() *model.Intermediary8 {
	newValue := new(model.Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrderV03) AddCopyDetails() *model.CopyInformation2 {
	s.CopyDetails = new(model.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionOrderV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800104 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:setr.008.001.04 Document"`
	Message *SubscriptionBulkOrderCancellationRequestV04 `xml:"SbcptBlkOrdrCxlReq"`
}

func (d *Document00800104) AddMessage() *SubscriptionBulkOrderCancellationRequestV04 {
	d.Message = new(SubscriptionBulkOrderCancellationRequestV04)
	return d.Message
}

// Scope
// The SubscriptionBulkOrderCancellationRequest message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to request the cancellation of a previously sent SubscriptionBulkOrder.
// Usage
// The SubscriptionBulkOrderCancellationRequest message is used to request the cancellation of one or more individual orders.
// There is no amendment, but a cancellation and re-instruct policy.
// To request the cancellation of one or more individual orders, the order reference of each individual order listed in the original SubscriptionBulkOrder message is specified in the order reference element. The message identification of the SubscriptionBulkOrder message which contains the individual orders to be cancelled may also be quoted in PreviousReference but this is not recommended.
// The deadline and acceptance of a cancellation request is subject to a service level agreement (SLA). This cancellation message is a cancellation request. There is no automatic acceptance of the cancellation.
// The rejection or acceptance of a SubscriptionBulkOrderCancellationRequest is made using an OrderCancellationStatusReport message.
type SubscriptionBulkOrderCancellationRequestV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the individual order to be cancelled.
	OrderReferences []*model.InvestmentFundOrder9 `xml:"OrdrRefs"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`
}

func (s *SubscriptionBulkOrderCancellationRequestV04) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionBulkOrderCancellationRequestV04) AddPoolReference() *model.AdditionalReference9 {
	s.PoolReference = new(model.AdditionalReference9)
	return s.PoolReference
}

func (s *SubscriptionBulkOrderCancellationRequestV04) AddPreviousReference() *model.AdditionalReference8 {
	s.PreviousReference = new(model.AdditionalReference8)
	return s.PreviousReference
}

func (s *SubscriptionBulkOrderCancellationRequestV04) SetMasterReference(value string) {
	s.MasterReference = (*model.Max35Text)(&value)
}

func (s *SubscriptionBulkOrderCancellationRequestV04) AddOrderReferences() *model.InvestmentFundOrder9 {
	newValue := new(model.InvestmentFundOrder9)
	s.OrderReferences = append(s.OrderReferences, newValue)
	return newValue
}

func (s *SubscriptionBulkOrderCancellationRequestV04) AddCopyDetails() *model.CopyInformation4 {
	s.CopyDetails = new(model.CopyInformation4)
	return s.CopyDetails
}

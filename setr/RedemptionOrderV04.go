package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400104 struct {
	XMLName xml.Name            `xml:"urn:iso:std:iso:20022:tech:xsd:setr.004.001.04 Document"`
	Message *RedemptionOrderV04 `xml:"RedOrdr"`
}

func (d *Document00400104) AddMessage() *RedemptionOrderV04 {
	d.Message = new(RedemptionOrderV04)
	return d.Message
}

// Scope
// The RedemptionOrder message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to instruct the redemption of one or more financial instruments for one investment fund account.
// Usage
// The RedemptionOrder message is used to instruct single redemption orders, that is, a message containing one order for one financial instrument and related to one investment account. The RedemptionOrder message may also be used for multiple orders, that is, a message containing several orders related to the same investment account for different financial instruments.
// For a single redemption order, the RedemptionOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for the same financial instrument but for different accounts that are to be communicated in a single message, then the RedemptionBulkOrder message must be used.
type RedemptionOrderV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// General information related to the orders.
	MultipleOrderDetails *model.RedemptionMultipleOrder6 `xml:"MltplOrdrDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionOrderV04) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionOrderV04) AddPoolReference() *model.AdditionalReference9 {
	r.PoolReference = new(model.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionOrderV04) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionOrderV04) AddMultipleOrderDetails() *model.RedemptionMultipleOrder6 {
	r.MultipleOrderDetails = new(model.RedemptionMultipleOrder6)
	return r.MultipleOrderDetails
}

func (r *RedemptionOrderV04) AddCopyDetails() *model.CopyInformation4 {
	r.CopyDetails = new(model.CopyInformation4)
	return r.CopyDetails
}

func (r *RedemptionOrderV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

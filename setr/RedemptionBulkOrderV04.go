package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100104 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.001.001.04 Document"`
	Message *RedemptionBulkOrderV04 `xml:"RedBlkOrdr"`
}

func (d *Document00100104) AddMessage() *RedemptionBulkOrderV04 {
	d.Message = new(RedemptionBulkOrderV04)
	return d.Message
}

// Scope
// The RedemptionBulkOrder message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to instruct a redemption from a financial instrument for two or more accounts.
// Usage
// The RedemptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are for the same financial instrument. The RedemptionBulkOrder can result in one single bulk cash settlement or several individual cash settlements.
// This message will be typically used by a party collecting orders, that is, a concentrator, bulking these individual orders into one bulk order before sending it to an executing party.
// For a single redemption order, the RedemptionOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for different financial instruments but for the same account, then the RedemptionOrder must be used.
type RedemptionBulkOrderV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// General information related to the orders.
	BulkOrderDetails *model.RedemptionBulkOrder6 `xml:"BlkOrdrDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderV04) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderV04) AddPoolReference() *model.AdditionalReference9 {
	r.PoolReference = new(model.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionBulkOrderV04) AddPreviousReference() *model.AdditionalReference8 {
	newValue := new(model.AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV04) AddBulkOrderDetails() *model.RedemptionBulkOrder6 {
	r.BulkOrderDetails = new(model.RedemptionBulkOrder6)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderV04) AddCopyDetails() *model.CopyInformation4 {
	r.CopyDetails = new(model.CopyInformation4)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

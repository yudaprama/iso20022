package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.001.001.03 Document"`
	Message *RedemptionBulkOrderV03 `xml:"RedBlkOrdrV03"`
}

func (d *Document00100103) AddMessage() *RedemptionBulkOrderV03 {
	d.Message = new(RedemptionBulkOrderV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative sends the RedemptionBulkOrder message to the executing party, for example, a transfer agent, to instruct a redemption from a financial instrument for two or more accounts.
// Usage
// The RedemptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, that is, account owners, but are for the same financial instrument. The RedemptionBulkOrder can result in one single bulk cash settlement or several individual cash settlements.
// This message will be typically used by a party collecting orders, that is, a concentrator, bulking these individual orders into one bulk order before sending it to an executing party.
// For a single redemption order, the RedemptionOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for different financial instruments but for the same account, then the RedemptionOrder must be used.
type RedemptionBulkOrderV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Common information related to all the orders.
	BulkOrderDetails *model.RedemptionBulkOrder4 `xml:"BlkOrdrDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*model.Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderV03) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderV03) AddPoolReference() *model.AdditionalReference3 {
	r.PoolReference = new(model.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderV03) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV03) AddBulkOrderDetails() *model.RedemptionBulkOrder4 {
	r.BulkOrderDetails = new(model.RedemptionBulkOrder4)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderV03) AddRelatedPartyDetails() *model.Intermediary8 {
	newValue := new(model.Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV03) AddCopyDetails() *model.CopyInformation2 {
	r.CopyDetails = new(model.CopyInformation2)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

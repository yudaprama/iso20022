package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.001.001.02 Document"`
	Message *RedemptionBulkOrderV02 `xml:"setr.001.001.02"`
}

func (d *Document00100102) AddMessage() *RedemptionBulkOrderV02 {
	d.Message = new(RedemptionBulkOrderV02)
	return d.Message
}

// Scope
// The RedemptionBulkOrder message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to instruct the executing party to redeem a financial instrument for two or more accounts.
// Usage
// The RedemptionBulkOrder message is used to bulk several individual orders into one bulk order. The individual orders come from different instructing parties, ie, account owners, but are the same financial instrument. The RedemptionBulkOrder can result in one single bulk cash settlement or several individual cash settlements.
// This message will be typically used by a party collecting orders and bulking these individual orders into one bulk order before sending it to another party.
// For a single redemption order, the RedemptionMultipleOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for different financial instruments but for the same account, then the RedemptionMultipleOrder must be used.
type RedemptionBulkOrderV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Common information related to all the orders.
	BulkOrderDetails *model.RedemptionBulkOrder2 `xml:"BlkOrdrDtls"`

	// The information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderV02) AddMasterReference() *model.AdditionalReference3 {
	r.MasterReference = new(model.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionBulkOrderV02) AddPoolReference() *model.AdditionalReference3 {
	r.PoolReference = new(model.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV02) AddBulkOrderDetails() *model.RedemptionBulkOrder2 {
	r.BulkOrderDetails = new(model.RedemptionBulkOrder2)
	return r.BulkOrderDetails
}

func (r *RedemptionBulkOrderV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionBulkOrderV02) AddCopyDetails() *model.CopyInformation1 {
	r.CopyDetails = new(model.CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

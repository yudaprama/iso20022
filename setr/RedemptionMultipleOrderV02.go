package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:setr.004.001.02 Document"`
	Message *RedemptionMultipleOrderV02 `xml:"setr.004.001.02"`
}

func (d *Document00400102) AddMessage() *RedemptionMultipleOrderV02 {
	d.Message = new(RedemptionMultipleOrderV02)
	return d.Message
}

// Scope
// The RedemptionMultipleOrder message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to instruct the executing party to redeem to one or more financial instruments, for the same account.
// Usage
// The RedemptionMultipleOrder message is used for multiple orders. It may also be used for single orders, ie, a message containing one order for one financial instrument and related to one investment account. For a single redemption order, the RedemptionMultipleOrder message, not the RedemptionBulkOrder message, must be used.
// If there are redemption orders for the same financial instrument but for different accounts, then the RedemptionBulkOrder must be used.
type RedemptionMultipleOrderV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// General information related to the order.
	MultipleOrderDetails *model.RedemptionMultipleOrder2 `xml:"MltplOrdrDtls"`

	// The information related to an intermediary.
	IntermediaryDetails []*model.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *model.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionMultipleOrderV02) AddMasterReference() *model.AdditionalReference3 {
	r.MasterReference = new(model.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionMultipleOrderV02) AddPoolReference() *model.AdditionalReference3 {
	r.PoolReference = new(model.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionMultipleOrderV02) AddPreviousReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionMultipleOrderV02) AddMultipleOrderDetails() *model.RedemptionMultipleOrder2 {
	r.MultipleOrderDetails = new(model.RedemptionMultipleOrder2)
	return r.MultipleOrderDetails
}

func (r *RedemptionMultipleOrderV02) AddIntermediaryDetails() *model.Intermediary4 {
	newValue := new(model.Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrderV02) AddCopyDetails() *model.CopyInformation1 {
	r.CopyDetails = new(model.CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionMultipleOrderV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

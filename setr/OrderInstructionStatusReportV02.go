package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:setr.016.001.02 Document"`
	Message *OrderInstructionStatusReportV02 `xml:"setr.016.001.02"`
}

func (d *Document01600102) AddMessage() *OrderInstructionStatusReportV02 {
	d.Message = new(OrderInstructionStatusReportV02)
	return d.Message
}

// Scope
// The OrderInstructionStatusReport is sent by an executing party, eg, a transfer agent, to an instructing party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the executing party and the instructing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message reports the status of an order from the time the executing party receives the order until the order is executed.
// Usage
// The OrderInstructionStatusReport message is sent by an executing party to the instructing party to report on the status of a subscription, redemption or a switch order.
// The message can be used to report one of the following:
// - a received status, or
// - an accepted status, or
// - a sent to next party status, or
// - an already executed status, or
// - a cancelled status, or
// - a conditionally accepted status, or
// - a rejected status, or
// - a suspended status, or
// - an in-repair status (at the individual order level only), or
// - repaired conditions (at the individual order level only).
// For subscription and redemption orders, the OrderInstructionStatusReport message covers both bulk and multiple categories of orders, and this message may provide the status either at the bulk or at the individual level.
// For a switch order, this message provides the status of the whole order, ie, it is not possible to accept one leg and to reject the other leg, the entire switch order has to be rejected. In order to identify which leg within the switch is causing a problem, the redemption or subscription leg identification is used.
type OrderInstructionStatusReportV02 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference []*model.AdditionalReference3 `xml:"OthrRef"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef"`

	// Reference to a multiple order or bulk order that represents the common reference of several individual orders.
	MasterReference *model.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Status report details of a bulk or multiple or switch order that was previously received.
	OrderDetailsReport *model.OrderStatusAndReason3 `xml:"OrdrDtlsRpt"`

	// Status report details of the individual orders of a bulk or multiple order that was previously received.
	IndividualOrderDetailsReport []*model.IndividualOrderStatusAndReason1 `xml:"IndvOrdrDtlsRpt"`
}

func (o *OrderInstructionStatusReportV02) AddOtherReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	o.OtherReference = append(o.OtherReference, newValue)
	return newValue
}

func (o *OrderInstructionStatusReportV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	o.RelatedReference = append(o.RelatedReference, newValue)
	return newValue
}

func (o *OrderInstructionStatusReportV02) AddMasterReference() *model.AdditionalReference3 {
	o.MasterReference = new(model.AdditionalReference3)
	return o.MasterReference
}

func (o *OrderInstructionStatusReportV02) AddOrderDetailsReport() *model.OrderStatusAndReason3 {
	o.OrderDetailsReport = new(model.OrderStatusAndReason3)
	return o.OrderDetailsReport
}

func (o *OrderInstructionStatusReportV02) AddIndividualOrderDetailsReport() *model.IndividualOrderStatusAndReason1 {
	newValue := new(model.IndividualOrderStatusAndReason1)
	o.IndividualOrderDetailsReport = append(o.IndividualOrderDetailsReport, newValue)
	return newValue
}

package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600104 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:setr.016.001.04 Document"`
	Message *OrderInstructionStatusReportV04 `xml:"OrdrInstrStsRpt"`
}

func (d *Document01600104) AddMessage() *OrderInstructionStatusReportV04 {
	d.Message = new(OrderInstructionStatusReportV04)
	return d.Message
}

// Scope
// The OrderInstructionStatusReport message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to report the status of an order from the time the executing party receives the order until the time the order is executed.
// Usage
// The OrderInstructionStatusReport message is used to report on the status of a subscription, redemption or a switch order.
// The OrderInstructionStatusReport message may be used to give the status of:
// - one order message by using OrderDetailsReport or,
// - one or more individual order instructions by using IndividualOrderDetailsReport, or
// - one or more switch orders by using SwitchOrderDetailsReport.
// If the OrderInstructionStatusReport message is used to report the status of an individual order, then IndividualOrderDetailsReport is used and the order reference of the individual order is quoted in OrderReference. The message identification of the message in which the individual order was conveyed may also be quoted in RelatedReference but this is not recommended.
// If the OrderInstructionStatusReport message is used to report the status of a switch order, then SwitchOrderDetailsReport is used and the order reference of the switch order is quoted in OrderReference. The message identification of the message in which the switch order was conveyed may also be quoted in RelatedReference but this is not recommended.
// If the OrderInstructionStatusReport message is used to report the status of an entire order message, for example, the SubscriptionBulkOrder, or a SubscriptionOrder containing several orders, then OrderDetailsReport is used and the message identification of the order message is quoted in RelatedReference. All the orders within the message must have the same status.
// One of the following statuses can be reported:
// - an accepted status , or,
// - an already executed status, or,
// - a sent to next party status, or,
// - a received status, or,
// - a settled status, or,
// - a communication problem with next party status, or,
// - a confirmation amendment status, or,
// - a done for the day status, or,
// - a partially done status, or,
// - an open status, or,
// - a cancelled status, or
// - a conditionally accepted status, or,
// - a rejected status, or,
// - a suspended status, or,
// - a partially settled status, or,
// - an in-repair status (only for an individual or switch order).
// For a switch order, the OrderInstructionStatusReport message provides the status of the whole switch order, that is, it is not possible to accept one leg and to reject the other leg: the entire switch order has to be rejected. In order to identify the legs within the switch that are causing the problem, the leg is identified in either the RedemptionLegIdentification or SubscriptionLegIdentification elements.
// When the OrderInstructionStatusReport is used to give the status of an individual or a switch order, the following can be specified:
// - repaired conditions (for a switch, this is at the level of a leg),
// - information related to the order, for example, settlement amount, number of units, expected trade date.
type OrderInstructionStatusReportV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to the message or communication that was previously received.
	Reference *model.References61Choice `xml:"Ref,omitempty"`

	// Status of the order.
	StatusReport *model.Status24Choice `xml:"StsRpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (o *OrderInstructionStatusReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	o.MessageIdentification = new(model.MessageIdentification1)
	return o.MessageIdentification
}

func (o *OrderInstructionStatusReportV04) AddReference() *model.References61Choice {
	o.Reference = new(model.References61Choice)
	return o.Reference
}

func (o *OrderInstructionStatusReportV04) AddStatusReport() *model.Status24Choice {
	o.StatusReport = new(model.Status24Choice)
	return o.StatusReport
}

func (o *OrderInstructionStatusReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	o.Extension = append(o.Extension, newValue)
	return newValue
}

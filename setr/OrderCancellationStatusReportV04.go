package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700104 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.017.001.04 Document"`
	Message *OrderCancellationStatusReportV04 `xml:"OrdrCxlStsRpt"`
}

func (d *Document01700104) AddMessage() *OrderCancellationStatusReportV04 {
	d.Message = new(OrderCancellationStatusReportV04)
	return d.Message
}

// Scope
// The OrderCancellationStatusReport message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to report the status of an order cancellation request that was previously received.
// Usage
// The OrderCancellationStatusReport message is used to provide the status of:
// - one or more individual order cancellation requests by using IndividualCancellationStatusReport, or,
// - an order cancellation request message by using CancellationStatusReport.
// If the OrderCancellationStatusReport message is used to report the status of an individual order cancellation request, then the repetitive IndividualCancellationStatusReport sequence is used and the order reference of the individual order is quoted in OrderReference. The message identification of the message in which the individual order was conveyed may also be quoted in RelatedReference but this is not recommended.
// If the OrderCancellationStatusReport message is used to report the status of an entire order cancellation request message, for example, the SubscriptionBulkOrderCancellationRequest, or a SubscriptionOrderCancellationRequest containing several orders, then the CancellationStatusReport sequence is used. The message identification of the order cancellation request message may also be quoted in RelatedReference but this is not recommended. All the order cancellation requests within the message must have the same status.
// One of the following statuses can be reported:
// - the order cancellation is pending, or,
// - the order cancellation request is rejected, or,
// - the order is cancelled.
// When the cancellation is rejected, the reason for the rejection must be specified.
type OrderCancellationStatusReportV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to the message or communication that was previously received.
	Reference *model.References61Choice `xml:"Ref,omitempty"`

	// Status of the order cancellation.
	StatusReport *model.Status26Choice `xml:"StsRpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (o *OrderCancellationStatusReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	o.MessageIdentification = new(model.MessageIdentification1)
	return o.MessageIdentification
}

func (o *OrderCancellationStatusReportV04) AddReference() *model.References61Choice {
	o.Reference = new(model.References61Choice)
	return o.Reference
}

func (o *OrderCancellationStatusReportV04) AddStatusReport() *model.Status26Choice {
	o.StatusReport = new(model.Status26Choice)
	return o.StatusReport
}

func (o *OrderCancellationStatusReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	o.Extension = append(o.Extension, newValue)
	return newValue
}

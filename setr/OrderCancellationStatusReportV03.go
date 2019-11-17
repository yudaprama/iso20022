package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700103 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.017.001.03 Document"`
	Message *OrderCancellationStatusReportV03 `xml:"OrdrCxlStsRptV03"`
}

func (d *Document01700103) AddMessage() *OrderCancellationStatusReportV03 {
	d.Message = new(OrderCancellationStatusReportV03)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the OrderCancellationStatusReport to the instructing party, for example, an investment manager or its authorised representative, to report the status of an order cancellation request that was previously received.
// Usage
// The OrderCancellationStatusReport message is used to provide the status of:
// - one or more individual order cancellation requests by using IndividualCancellationStatusReport, or,
// - an order cancellation request message by using CancellationStatusReport.
// If the OrderCancellationStatusReport message is used to report the status of an individual order cancellation request, then the repetitive IndividualCancellationStatusReport sequence is used and the order reference of the individual order is quoted in OrderReference. The message identification of the message in which the individual order was conveyed may also be quoted in RelatedReference.
// If the OrderCancellationStatusReport message is used to report the status of an entire order cancellation request message, for example, the SubscriptionBulkOrderCancellationRequest, or a SubscriptionOrderCancellationRequest containing several orders, then the CancellationStatusReport sequence. is used and the message identification of the order cancellation request message is quoted in RelatedReference. All the order cancellation requests within the message must have the same status.
// One of the following statuses can be reported:
// - the order cancellation is pending, or,
// - the order cancellation request is rejected, or,
// - the order is cancelled.
// When the cancellation is rejected, the reason for the rejection must be specified.
type OrderCancellationStatusReportV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference []*model.AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Status report details of a bulk or multiple or switch order cancellation message.
	CancellationStatusReport *model.OrderStatusAndReason8 `xml:"CxlStsRpt"`

	// Status report details of one or more individual orders from a bulk or multiple or switch order cancellation request.
	IndividualCancellationStatusReport []*model.IndividualOrderStatusAndReason4 `xml:"IndvCxlStsRpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (o *OrderCancellationStatusReportV03) AddMessageIdentification() *model.MessageIdentification1 {
	o.MessageIdentification = new(model.MessageIdentification1)
	return o.MessageIdentification
}

func (o *OrderCancellationStatusReportV03) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	o.RelatedReference = append(o.RelatedReference, newValue)
	return newValue
}

func (o *OrderCancellationStatusReportV03) AddOtherReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	o.OtherReference = append(o.OtherReference, newValue)
	return newValue
}

func (o *OrderCancellationStatusReportV03) AddCancellationStatusReport() *model.OrderStatusAndReason8 {
	o.CancellationStatusReport = new(model.OrderStatusAndReason8)
	return o.CancellationStatusReport
}

func (o *OrderCancellationStatusReportV03) AddIndividualCancellationStatusReport() *model.IndividualOrderStatusAndReason4 {
	newValue := new(model.IndividualOrderStatusAndReason4)
	o.IndividualCancellationStatusReport = append(o.IndividualCancellationStatusReport, newValue)
	return newValue
}

func (o *OrderCancellationStatusReportV03) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	o.Extension = append(o.Extension, newValue)
	return newValue
}

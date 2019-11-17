package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:setr.017.001.02 Document"`
	Message *OrderCancellationStatusReportV02 `xml:"setr.017.001.02"`
}

func (d *Document01700102) AddMessage() *OrderCancellationStatusReportV02 {
	d.Message = new(OrderCancellationStatusReportV02)
	return d.Message
}

// Scope
// The OrderCancellationStatusReport is sent by an executing party, eg, a transfer agent, to the instructing party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the executing party and the instructing party. The intermediary party is, for example, an intermediary or a concentrator.
// The message gives the status of an order cancellation instruction message that was previously sent by the instructing party.
// Usage
// The OrderCancellationStatusReport message is used to report the status of an order cancellation instruction message that was previously sent by the instructing party. The message can be used to report one of the following:
// - the cancellation is accepted for further processing, or
// - the cancellation is rejected, or
// - the order has been cancelled.
// When the cancellation is rejected, the reason for the rejection must be specified.
type OrderCancellationStatusReportV02 struct {

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference3 `xml:"RltdRef"`

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference []*model.AdditionalReference3 `xml:"OthrRef"`

	// Status report details of a bulk or multiple or switch order cancellation instruction that was previously received.
	CancellationStatusReport *model.OrderStatusAndReason4 `xml:"CxlStsRpt"`
}

func (o *OrderCancellationStatusReportV02) AddRelatedReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	o.RelatedReference = append(o.RelatedReference, newValue)
	return newValue
}

func (o *OrderCancellationStatusReportV02) AddOtherReference() *model.AdditionalReference3 {
	newValue := new(model.AdditionalReference3)
	o.OtherReference = append(o.OtherReference, newValue)
	return newValue
}

func (o *OrderCancellationStatusReportV02) AddCancellationStatusReport() *model.OrderStatusAndReason4 {
	o.CancellationStatusReport = new(model.OrderStatusAndReason4)
	return o.CancellationStatusReport
}

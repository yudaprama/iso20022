package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800104 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.04 Document"`
	Message *RequestForOrderStatusReportV04 `xml:"ReqForOrdrStsRpt"`
}

func (d *Document01800104) AddMessage() *RequestForOrderStatusReportV04 {
	d.Message = new(RequestForOrderStatusReportV04)
	return d.Message
}

// Scope
// The RequestForOrderStatusReport message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to request the status of one or more order instructions or order cancellation requests.
// Usage
// The RequestForOrderStatusReport message is used to request the status of:
// - one or several individual order instructions, or,
// - one or several order messages, or
// - one or several individual order cancellation requests, or,
// - one or several order cancellation request messages.
// The response to a RequestForOrderStatusReport message is the OrderInstructionStatusReport message or OrderCancellationStatusReport message.
// If the RequestForOrderStatusReport message is used to request the status of several individual order instructions or one or more order instruction messages, then the instructing party may receive several OrderInstructionStatusReport messages from the executing party.
// If the RequestForOrderStatusReport message is used to request the status of several individual order cancellation requests or one or more order cancellation messages, then the instructing party may receive several OrderCancellationStatusReport messages from the executing party.
// When the RequestForOrderStatusReport is used to request the status of one or more individual orders or order cancellations, each individual order is identified with its order reference. The investment account and/or financial instrument related to the order can also be identified. The message identification of the message in which the individual order was conveyed may also be quoted in PreviousReference.
// When the RequestForOrderStatusReport is used to request the status of an order message or an order cancellations request message, then the message identification of the order or cancellation message is identified in PreviousReference.
// The RequestForOrderStatusReport message may not be used to request the status of an investment account, a transfer or the status of a financial instrument.
type RequestForOrderStatusReportV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Information to identify the order(s) for which the status is requested.
	RequestDetails []*model.MessageAndBusinessReference10 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RequestForOrderStatusReportV04) AddMessageIdentification() *model.MessageIdentification1 {
	r.MessageIdentification = new(model.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RequestForOrderStatusReportV04) AddRequestDetails() *model.MessageAndBusinessReference10 {
	newValue := new(model.MessageAndBusinessReference10)
	r.RequestDetails = append(r.RequestDetails, newValue)
	return newValue
}

func (r *RequestForOrderStatusReportV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}

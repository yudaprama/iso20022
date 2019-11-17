package setr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:setr.018.001.02 Document"`
	Message *RequestForOrderStatusReportV02 `xml:"setr.018.001.02"`
}

func (d *Document01800102) AddMessage() *RequestForOrderStatusReportV02 {
	d.Message = new(RequestForOrderStatusReportV02)
	return d.Message
}

// Scope
// The RequestForOrderStatusReport is sent by an instructing party, eg, an investment manager or its authorised representative, to the executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party.
// This message requests the status of one or several order instruction or order cancellation messages.
// Usage
// The RequestForOrderStatusReport message is used to request the status of:
// - one or several order messages,
// - one or several cancellation messages,
// - one or several individual orders within a bulk or multiple order message.
// If the RequestForOrderStatusReport message is used to request the status of several messages, then the instructing party will receive several reply messages from the executing party, ie, several OrderInstructionStatusReport messages and/or OrderCancellationStatusReport messages. The number of reply messages will match the number of references stated in the RequestForOrderStatusReport message.
// The RequestForOrderStatusReport message may not be used to request the status of an investment account, a transfer or the status of a financial instrument.
type RequestForOrderStatusReportV02 struct {

	// Information to identify the order(s) for which the status is requested.
	RequestDetails []*model.MessageAndBusinessReference2 `xml:"ReqDtls"`
}

func (r *RequestForOrderStatusReportV02) AddRequestDetails() *model.MessageAndBusinessReference2 {
	newValue := new(model.MessageAndBusinessReference2)
	r.RequestDetails = append(r.RequestDetails, newValue)
	return newValue
}

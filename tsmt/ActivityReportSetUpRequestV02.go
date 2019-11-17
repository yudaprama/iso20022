package tsmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.004.001.02 Document"`
	Message *ActivityReportSetUpRequestV02 `xml:"ActvtyRptSetUpReq"`
}

func (d *Document00400102) AddMessage() *ActivityReportSetUpRequestV02 {
	d.Message = new(ActivityReportSetUpRequestV02)
	return d.Message
}

// Scope
// The ActivityReportSetUpRequest message is sent by any party involved in a transaction to the matching application.
// The ActivityReportSetUpRequest message can be sent to request the reset of the pre-determined time at which the daily production of the activity report should take place.
// Usage
// This message is sent to the matching application by a bank, in order to set the UTC offset specifying the hour when the matching application will generate every day an activity report covering the last 24 hours and send it. By default, this offset is equal to 0.
type ActivityReportSetUpRequestV02 struct {

	// Identifies the request message.
	RequestIdentification *model.MessageIdentification1 `xml:"ReqId"`

	// Specifies the parameters to calculate the local reporting time.
	UTCOffset *model.UTCOffset1 `xml:"UTCOffset"`
}

func (a *ActivityReportSetUpRequestV02) AddRequestIdentification() *model.MessageIdentification1 {
	a.RequestIdentification = new(model.MessageIdentification1)
	return a.RequestIdentification
}

func (a *ActivityReportSetUpRequestV02) AddUTCOffset() *model.UTCOffset1 {
	a.UTCOffset = new(model.UTCOffset1)
	return a.UTCOffset
}

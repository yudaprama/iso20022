package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03300103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.033.001.03 Document"`
	Message *RequestForDuplicateV03 `xml:"ReqForDplct"`
}

func (d *Document03300103) AddMessage() *RequestForDuplicateV03 {
	d.Message = new(RequestForDuplicateV03)
	return d.Message
}

// Scope
// The Request For Duplicate message is sent by the case assignee to the case creator or case assigner.
// This message is used to request a copy of the original payment instruction considered in the case.
// Usage
// The Request For Duplicate message:
// - must be answered with a Duplicate message
// - must be used when a case assignee requests a copy of the original payment instruction. This occurs, for example, when the case assignee cannot trace the payment instruction based on the elements mentioned in the case assignment message
// - covers one and only one instruction at a time. If several payment instruction copies are needed by the case assignee, then multiple Request For Duplicate messages must be sent
// - must be used exclusively between the case assignee and its case creator/case assigner
type RequestForDuplicateV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case2 `xml:"Case"`
}

func (r *RequestForDuplicateV03) AddAssignment() *model.CaseAssignment2 {
	r.Assignment = new(model.CaseAssignment2)
	return r.Assignment
}

func (r *RequestForDuplicateV03) AddCase() *model.Case2 {
	r.Case = new(model.Case2)
	return r.Case
}

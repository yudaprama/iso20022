package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03700103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.037.001.03 Document"`
	Message *DebitAuthorisationRequestV03 `xml:"DbtAuthstnReq"`
}

func (d *Document03700103) AddMessage() *DebitAuthorisationRequestV03 {
	d.Message = new(DebitAuthorisationRequestV03)
	return d.Message
}

// Scope
// The Debit Authorisation Request message is sent by an account servicing institution to an account owner. This message is used to request authorisation to debit an account.
// Usage
// The Debit Authorisation Request message must be answered with a Debit Authorisation Response message.
// The Debit Authorisation Request message can be used to request debit authorisation in a:
// - request to modify payment case (in the case of a lower final amount or change of creditor)
// - request to cancel payment case (full amount)
// - unable to apply case (the creditor whose account has been credited is not the intended beneficiary)
// - claim non receipt case (the creditor whose account has been credited is not the intended beneficiary)
// The Debit Authorisation Request message covers one and only one payment instruction at a time. If an account servicing institution needs to request debit authorisation for several instructions, then multiple Debit Authorisation Request messages must be sent.
// The Debit Authorisation Request must be used exclusively between the account servicing institution and the account owner. It must not be used in place of a Request To Modify Payment or Request To Cancel Payment message between subsequent agents.
type DebitAuthorisationRequestV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case2 `xml:"Case"`

	// Identifies the underlying payment instructrion.
	Underlying *model.UnderlyingTransaction1Choice `xml:"Undrlyg"`

	// Detailed information about the request.
	Detail *model.DebitAuthorisationDetails3 `xml:"Dtl"`
}

func (d *DebitAuthorisationRequestV03) AddAssignment() *model.CaseAssignment2 {
	d.Assignment = new(model.CaseAssignment2)
	return d.Assignment
}

func (d *DebitAuthorisationRequestV03) AddCase() *model.Case2 {
	d.Case = new(model.Case2)
	return d.Case
}

func (d *DebitAuthorisationRequestV03) AddUnderlying() *model.UnderlyingTransaction1Choice {
	d.Underlying = new(model.UnderlyingTransaction1Choice)
	return d.Underlying
}

func (d *DebitAuthorisationRequestV03) AddDetail() *model.DebitAuthorisationDetails3 {
	d.Detail = new(model.DebitAuthorisationDetails3)
	return d.Detail
}

package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03700104 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.037.001.04 Document"`
	Message *DebitAuthorisationRequestV04 `xml:"DbtAuthstnReq"`
}

func (d *Document03700104) AddMessage() *DebitAuthorisationRequestV04 {
	d.Message = new(DebitAuthorisationRequestV04)
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
type DebitAuthorisationRequestV04 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Identifies the underlying payment instructrion.
	Underlying *model.UnderlyingTransaction2Choice `xml:"Undrlyg"`

	// Detailed information about the request.
	Detail *model.DebitAuthorisation1 `xml:"Dtl"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DebitAuthorisationRequestV04) AddAssignment() *model.CaseAssignment3 {
	d.Assignment = new(model.CaseAssignment3)
	return d.Assignment
}

func (d *DebitAuthorisationRequestV04) AddCase() *model.Case3 {
	d.Case = new(model.Case3)
	return d.Case
}

func (d *DebitAuthorisationRequestV04) AddUnderlying() *model.UnderlyingTransaction2Choice {
	d.Underlying = new(model.UnderlyingTransaction2Choice)
	return d.Underlying
}

func (d *DebitAuthorisationRequestV04) AddDetail() *model.DebitAuthorisation1 {
	d.Detail = new(model.DebitAuthorisation1)
	return d.Detail
}

func (d *DebitAuthorisationRequestV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03700101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.037.001.01 Document"`
	Message *DebitAuthorisationRequest `xml:"camt.037.001.01"`
}

func (d *Document03700101) AddMessage() *DebitAuthorisationRequest {
	d.Message = new(DebitAuthorisationRequest)
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
type DebitAuthorisationRequest struct {

	// Identifies the case assignment.
	Assignment *model.CaseAssignment `xml:"Assgnmt"`

	// Identifies the case.
	Case *model.Case `xml:"Case"`

	// Identifies the underlying payment instructrion.
	Underlying *model.PaymentInstructionExtract `xml:"Undrlyg"`

	// Detailed information about the request.
	Detail *model.DebitAuthorisationDetails `xml:"Dtl"`
}

func (d *DebitAuthorisationRequest) AddAssignment() *model.CaseAssignment {
	d.Assignment = new(model.CaseAssignment)
	return d.Assignment
}

func (d *DebitAuthorisationRequest) AddCase() *model.Case {
	d.Case = new(model.Case)
	return d.Case
}

func (d *DebitAuthorisationRequest) AddUnderlying() *model.PaymentInstructionExtract {
	d.Underlying = new(model.PaymentInstructionExtract)
	return d.Underlying
}

func (d *DebitAuthorisationRequest) AddDetail() *model.DebitAuthorisationDetails {
	d.Detail = new(model.DebitAuthorisationDetails)
	return d.Detail
}

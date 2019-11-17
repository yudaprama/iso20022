package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03600101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.036.001.01 Document"`
	Message *DebitAuthorisationResponse `xml:"camt.036.001.01"`
}

func (d *Document03600101) AddMessage() *DebitAuthorisationResponse {
	d.Message = new(DebitAuthorisationResponse)
	return d.Message
}

// Scope
// The Debit Authorisation Response message is sent by an account owner to its account servicing institution. This message is used to approve or reject a debit authorisation request.
// Usage
// The Debit Authorisation Response message is used to reply to a Debit Authorisation Request message.
// The Debit Authorisation Response message covers one and only one payment instruction at a time. If an account owner needs to reply to several Debit Authorisation Request messages, then multiple Debit Authorisation Response messages must be sent.
// The Debit Authorisation Response message indicates whether the account owner agrees with the request by means of a code. It also allows further details to be given about the debit authorisation, such as acceptable amount and value date for the debit.
// The Debit Authorisation Response message must be used exclusively between the account owner and the account servicing institution. It must not be used in place of a Resolution Of Investigation message between subsequent agents.
type DebitAuthorisationResponse struct {

	// Identifies an assignment.
	Assignment *model.CaseAssignment `xml:"Assgnmt"`

	// Identifies a case.
	Case *model.Case `xml:"Case"`

	// Indicates if the debit authorisation is granted or not.
	Confirmation *model.DebitAuthorisationConfirmation `xml:"Conf"`
}

func (d *DebitAuthorisationResponse) AddAssignment() *model.CaseAssignment {
	d.Assignment = new(model.CaseAssignment)
	return d.Assignment
}

func (d *DebitAuthorisationResponse) AddCase() *model.Case {
	d.Case = new(model.Case)
	return d.Case
}

func (d *DebitAuthorisationResponse) AddConfirmation() *model.DebitAuthorisationConfirmation {
	d.Confirmation = new(model.DebitAuthorisationConfirmation)
	return d.Confirmation
}

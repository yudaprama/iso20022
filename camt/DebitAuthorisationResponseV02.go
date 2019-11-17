package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03600102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.036.001.02 Document"`
	Message *DebitAuthorisationResponseV02 `xml:"DbtAuthstnRspn"`
}

func (d *Document03600102) AddMessage() *DebitAuthorisationResponseV02 {
	d.Message = new(DebitAuthorisationResponseV02)
	return d.Message
}

// Scope
// The Debit Authorisation Response message is sent by an account owner to its account servicing institution. This message is used to approve or reject a debit authorisation request.
// Usage
// The Debit Authorisation Response message is used to reply to a Debit Authorisation Request message.
// The Debit Authorisation Response message covers one and only one payment instruction at a time. If an account owner needs to reply to several Debit Authorisation Request messages, then multiple Debit Authorisation Response messages must be sent.
// The Debit Authorisation Response message indicates whether the account owner agrees with the request by means of a code. It also allows further details to be given about the debit authorisation, such as acceptable amount and value date for the debit.
// The Debit Authorisation Response message must be used exclusively between the account owner and the account servicing institution. It must not be used in place of a Resolution Of Investigation message between subsequent agents.
type DebitAuthorisationResponseV02 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case2 `xml:"Case"`

	// Indicates if the debit authorisation is granted or not.
	Confirmation *model.DebitAuthorisationConfirmation2 `xml:"Conf"`
}

func (d *DebitAuthorisationResponseV02) AddAssignment() *model.CaseAssignment2 {
	d.Assignment = new(model.CaseAssignment2)
	return d.Assignment
}

func (d *DebitAuthorisationResponseV02) AddCase() *model.Case2 {
	d.Case = new(model.Case2)
	return d.Case
}

func (d *DebitAuthorisationResponseV02) AddConfirmation() *model.DebitAuthorisationConfirmation2 {
	d.Confirmation = new(model.DebitAuthorisationConfirmation2)
	return d.Confirmation
}

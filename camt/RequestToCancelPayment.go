package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800201 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.008.002.01 Document"`
	Message *RequestToCancelPayment `xml:"camt.008.002.01"`
}

func (d *Document00800201) AddMessage() *RequestToCancelPayment {
	d.Message = new(RequestToCancelPayment)
	return d.Message
}

// Scope
// The Request To Cancel Payment message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the cancellation of an original payment instruction.
// Usage
// The Request To Cancel Payment message must be answered with a:
// - Resolution Of Investigation message with a positive final outcome when the case assignee can perform the requested cancellation
// - Resolution Of Investigation message with a negative final outcome when the case assignee may perform the requested cancellation but fails to do so (too late, irrevocable instruction, ...)
// - Reject Case Assignment message when the case assignee is unable or not authorised to perform the requested cancellation
// - Notification Of Case Assignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// A Request To Cancel Payment message concerns one and only one original payment instruction at a time. If several original payment instructions need to be cancelled, then multiple Request To Cancel Payment messages must be sent.
// When a case assignee successfully performs a cancellation, it must return the corresponding funds to the case assigner. It may provide some details about the return in the Resolution Of Investigation message.
// The processing of a request to cancel payment case may end with a Debit Authorisation Request message sent to the creditor by its account servicing institution.
// The Request To Cancel Payment message may be used to escalate a case after an unsuccessful request to modify the payment. In this scenario, the case identification remains the same as in the original Request To Modify Payment message and the element ReopenCaseIndication is set to 'Yes' or 'true'.
type RequestToCancelPayment struct {

	// Identifies the assignment of a case from an assigner to an assignee.
	Assignment *model.CaseAssignment `xml:"Assgnmt"`

	// Identifies the case.
	Case *model.Case `xml:"Case"`

	// Identifies the payment instruction to be cancelled.
	Underlying *model.PaymentInstructionExtract `xml:"Undrlyg"`

	// Defines the reason for requesting the cancellation.
	Justification *model.DebitAuthorisationDetails `xml:"Justfn"`
}

func (r *RequestToCancelPayment) AddAssignment() *model.CaseAssignment {
	r.Assignment = new(model.CaseAssignment)
	return r.Assignment
}

func (r *RequestToCancelPayment) AddCase() *model.Case {
	r.Case = new(model.Case)
	return r.Case
}

func (r *RequestToCancelPayment) AddUnderlying() *model.PaymentInstructionExtract {
	r.Underlying = new(model.PaymentInstructionExtract)
	return r.Underlying
}

func (r *RequestToCancelPayment) AddJustification() *model.DebitAuthorisationDetails {
	r.Justification = new(model.DebitAuthorisationDetails)
	return r.Justification
}

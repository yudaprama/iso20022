package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02600101 struct {
	XMLName xml.Name       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.01 Document"`
	Message *UnableToApply `xml:"camt.026.001.01"`
}

func (d *Document02600101) AddMessage() *UnableToApply {
	d.Message = new(UnableToApply)
	return d.Message
}

// Scope
// The Unable To Apply message is sent by a case creator or a case assigner to a case assignee. This message is used to initiate an investigation of a payment instruction that cannot be executed or reconciled.
// Usage
// The Unable To Apply case occurs in two situations:
// - an agent cannot execute the payment instruction due to missing or incorrect information
// - a creditor cannot reconcile the payment entry as it is received unexpectedly, or missing or incorrect information prevents reconciliation
// The Unable To Apply message:
// - always follows the reverse route of the original payment instruction
// - must be forwarded to the preceding agents in the payment processing chain, where appropriate
// - covers one and only one payment instruction (or payment entry) at a time; if several payment instructions cannot be executed or several payment entries cannot be reconciled, then multiple Unable To Apply messages must be sent.
// Depending on what stage the payment is and what has been done to it, for example incorrect routing, errors/omissions when processing the instruction or even the absence of any error, the unable to apply case may lead to a:
// - Additional Payment Information message, sent to the case creator/case assigner, if a truncation or omission in a payment instruction prevented reconciliation.
// - Request To Cancel Payment message, sent to the subsequent agent in the payment processing chain, if the original payment instruction has been incorrectly routed through the chain of agents (this also entails a new corrected payment instruction being issued). Prior to sending the payment cancellation request, the agent should first send a resolution indicating that a cancellation will follow (CWFW).
// - Request To Modify Payment message, sent to the subsequent agent in the payment processing chain, if a truncation or omission has occurred during the processing of the original payment instruction. Prior to sending the modify payment request, the agent should first send a resolution indicating that a modification will follow (MWFW).
// - Debit Authorisation Request message, sent to the case creator/case assigner, if the payment instruction has reached an incorrect creditor.
type UnableToApply struct {

	// Identifies the assignment.
	Assignment *model.CaseAssignment `xml:"Assgnmt"`

	// Identifies the case.
	Case *model.Case `xml:"Case"`

	// References the Payment Instruction that a Party is unable to execute or unable to reconcile.
	Underlying *model.PaymentInstructionExtract `xml:"Undrlyg"`

	// Explains the reason why unable to apply.
	Justification *model.UnableToApplyJustificationChoice `xml:"Justfn"`
}

func (u *UnableToApply) AddAssignment() *model.CaseAssignment {
	u.Assignment = new(model.CaseAssignment)
	return u.Assignment
}

func (u *UnableToApply) AddCase() *model.Case {
	u.Case = new(model.Case)
	return u.Case
}

func (u *UnableToApply) AddUnderlying() *model.PaymentInstructionExtract {
	u.Underlying = new(model.PaymentInstructionExtract)
	return u.Underlying
}

func (u *UnableToApply) AddJustification() *model.UnableToApplyJustificationChoice {
	u.Justification = new(model.UnableToApplyJustificationChoice)
	return u.Justification
}

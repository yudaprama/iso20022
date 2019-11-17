package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02600103 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.03 Document"`
	Message *UnableToApplyV03 `xml:"UblToApply"`
}

func (d *Document02600103) AddMessage() *UnableToApplyV03 {
	d.Message = new(UnableToApplyV03)
	return d.Message
}

// Scope
// The Unable To Apply message is sent by a case creator or a case assigner to a case assignee. This message is used to initiate an investigation of a payment instruction that cannot be executed or reconciled.
// Usage
// The Unable To Apply case occurs in two situations:
// - an agent cannot execute the payment instruction due to missing or incorrect information
// - a creditor cannot reconcile the payment entry as it is received unexpectedly, or missing or incorrect information prevents reconciliation
// The Unable To Apply message
// - always follows the reverse route of the original payment instruction
// - must be forwarded to the preceding agents in the payment processing chain, where appropriate
// - covers one and only one payment instruction (or payment entry) at a time; if several payment instructions cannot be executed or several payment entries cannot be reconciled, then multiple Unable To Apply messages must be sent.
// Depending on what stage the payment is and what has been done to it, for example incorrect routing, errors/omissions when processing the instruction or even the absence of any error, the unable to apply case may lead to a:
// - Additional Payment Information message, sent to the case creator/case assigner, if a truncation or omission in a payment instruction prevented reconciliation.
// - Request To Cancel Payment message, sent to the subsequent agent in the payment processing chain, if the original payment instruction has been incorrectly routed through the chain of agents (this also entails a new corrected payment instruction being issued). Prior to sending the payment cancellation request, the agent should first send a resolution indicating that a cancellation will follow (CWFW).
// - Request To Modify Payment message, sent to the subsequent agent in the payment processing chain, if a truncation or omission has occurred during the processing of the original payment instruction. Prior to sending the modify payment request, the agent should first send a resolution indicating that a modification will follow (MWFW).
// - Debit Authorisation Request message, sent to the case creator/case assigner, if the payment instruction has reached an incorrect creditor.
// The UnableToApply message has the following main characteristics:
// - Case Identification and Reason Code:
// The case creator (the instructed party/creditor of a payment instruction) assigns a unique case identification and optionally
// the reason code for the Unable To Apply message. This information will be passed unchanged to all subsequent case
// assignees.
// - Underlying Payment Instruction Identification:
// The case creator specifies the identification of the underlying payment instruction. This identification needs to be updated
// by the subsequent case assigner(s) in order to match the one used with their case assignee(s).
// - Unable To Apply Justification:
// The Unable To Apply Justification element allows the assigner to indicate whether a specific element causes the unable
// to apply situation (incorrect element and/or mismatched element can be listed) or whether any supplementary information
// needs to be forwarded.
type UnableToApplyV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *model.Case2 `xml:"Case"`

	// References the payment instruction or statement entry that a party is unable to execute or unable to reconcile.
	Underlying *model.UnderlyingTransaction1Choice `xml:"Undrlyg"`

	// Explains the reason why the case creator is unable to apply the instruction.
	Justification *model.UnableToApplyJustification2Choice `xml:"Justfn"`
}

func (u *UnableToApplyV03) AddAssignment() *model.CaseAssignment2 {
	u.Assignment = new(model.CaseAssignment2)
	return u.Assignment
}

func (u *UnableToApplyV03) AddCase() *model.Case2 {
	u.Case = new(model.Case2)
	return u.Case
}

func (u *UnableToApplyV03) AddUnderlying() *model.UnderlyingTransaction1Choice {
	u.Underlying = new(model.UnderlyingTransaction1Choice)
	return u.Underlying
}

func (u *UnableToApplyV03) AddJustification() *model.UnableToApplyJustification2Choice {
	u.Justification = new(model.UnableToApplyJustification2Choice)
	return u.Justification
}

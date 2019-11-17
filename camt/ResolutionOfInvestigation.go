package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.01 Document"`
	Message *ResolutionOfInvestigation `xml:"camt.029.001.01"`
}

func (d *Document02900101) AddMessage() *ResolutionOfInvestigation {
	d.Message = new(ResolutionOfInvestigation)
	return d.Message
}

// Scope
// The Resolution Of Investigation message is sent by a case assignee to a case creator/case assigner.
// This message is used to inform of the resolution of a case, and optionally provides details about .
// - the corrective action undertaken by the case assignee
// - information on the return where applicable
// Usage
// The Resolution Of Investigation message is used by the case assignee to inform a case creator or case assigner about the resolution of a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// The Resolution Of Investigation message covers one and only one case at a time. If the case assignee needs to communicate about several cases, then several Resolution Of Investigation messages must be sent.
// The Resolution Of Investigation message provides:
// - the final outcome of the case, whether positive or negative
// - optionally, the details of the corrective action undertaken by the case assignee and the information of the return
// Whenever a payment instruction has been generated to solve the case under investigation, the optional CorrectionTransaction component present in the message must be completed.
// Whenever the action of modifying or cancelling a payment results in funds being returned, an investigating agent may attached some details in this message. These details facilitates the account reconciliations at the initiating bank and the intermediaries. It must be stressed that returning of funds is outside the scope of this Exceptions and Investigation service. The features given here is only meant to transmit the information of returns when it is available.
// The Resolution Of Investigation message must
// - be forwarded by all subsequent case assignee(s) until it reaches the case creator
// - not be used in place of a Reject Case Assignment or Case Status Report or Notification Of Case Assignment message.
// Take note of an exceptional rule that allows the use of Resolution Of Investigation in lieu of a Case Status Report. Case Status Report is a response-message to a Case Status Report Request. The latter which is sent when the assigner has waited too long (by his standard) for an answer. However it may happen that when the Request arrives, the investigating agent has just obtained a resolution. In such a situation, it would be redundant to send a Case Status Report when then followed immediately by a Resolution Of Investigation. It is therefore quite acceptable for the investigating agent, the assignee, to skip the Case Status Report and send the Resolution Of Investigation message directly.
type ResolutionOfInvestigation struct {

	// Note: the Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment `xml:"Assgnmt"`

	// Identifies a resolved case.
	ResolvedCase *model.Case `xml:"RslvdCase"`

	// Indicates the status of the investigation.
	Status *model.InvestigationStatusChoice `xml:"Sts,omitempty"`

	// References a transaction intitiated to fix the case under investigation.
	CorrectionTransaction *model.PaymentInstructionExtract `xml:"CrrctnTx,omitempty"`
}

func (r *ResolutionOfInvestigation) AddAssignment() *model.CaseAssignment {
	r.Assignment = new(model.CaseAssignment)
	return r.Assignment
}

func (r *ResolutionOfInvestigation) AddResolvedCase() *model.Case {
	r.ResolvedCase = new(model.Case)
	return r.ResolvedCase
}

func (r *ResolutionOfInvestigation) AddStatus() *model.InvestigationStatusChoice {
	r.Status = new(model.InvestigationStatusChoice)
	return r.Status
}

func (r *ResolutionOfInvestigation) AddCorrectionTransaction() *model.PaymentInstructionExtract {
	r.CorrectionTransaction = new(model.PaymentInstructionExtract)
	return r.CorrectionTransaction
}

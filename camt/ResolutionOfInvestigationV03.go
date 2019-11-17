package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900103 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.03 Document"`
	Message *ResolutionOfInvestigationV03 `xml:"RsltnOfInvstgtn"`
}

func (d *Document02900103) AddMessage() *ResolutionOfInvestigationV03 {
	d.Message = new(ResolutionOfInvestigationV03)
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
// Whenever a payment instruction has been generated to solve the case under investigation following a claim non receipt or an unable to apply, the optional CorrectionTransaction component present in the message must be completed.
// Whenever the action of modifying or cancelling a payment results in funds being returned or reversed, an investigating agent may provide the details in the resolution related investigation component, to identify the return or reversal transaction. These details will facilitate the account reconciliations at the initiating bank and the intermediaries. It must be stressed that the return or reversal of funds is outside the scope of this Exceptions and Investigation service. The features given here is only meant to transmit the information of return or reversal when it is available through the resolution of the case.
// The Resolution Of Investigation message must:
// - be forwarded by all subsequent case assignee(s) until it reaches the case creator
// - not be used in place of a Reject Case Assignment or Case Status Report or Notification Of Case Assignment message
// Take note of an exceptional rule that allows the use of Resolution Of Investigation in lieu of a Case Status Report. Case Status Report is a response-message to a Case Status Report Request. The latter which is sent when the assigner has reached its own time-out threshold to receive a response. However it may happen that when the Request arrives, the investigating agent has just obtained a resolution. In such a situation, it would be redundant to send a Case Status Report when then followed immediately by a Resolution Of Investigation. It is therefore quite acceptable for the investigating agent, the assignee, to skip the Case Status Report and send the Resolution Of Investigation message directly.
// The Resolution Of Investigation message should be the sole message to respond to a cancellation request. Details of the underlying transactions and the related statuses for which the cancellation request has been issued may be provided in the Cancellation Details component.
type ResolutionOfInvestigationV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies a resolved case.
	ResolvedCase *model.Case2 `xml:"RslvdCase,omitempty"`

	// Indicates the status of the investigation.
	Status *model.InvestigationStatus2Choice `xml:"Sts"`

	// Specifies the details of the underlying transactions being cancelled.
	CancellationDetails []*model.UnderlyingTransaction3 `xml:"CxlDtls,omitempty"`

	// Details on the underlying statement entry.
	StatementDetails *model.StatementResolutionEntry1 `xml:"StmtDtls,omitempty"`

	// References a transaction initiated to fix the case under investigation.
	CorrectionTransaction *model.CorrectiveTransaction1Choice `xml:"CrrctnTx,omitempty"`

	// Reference of a return or a reversal initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *model.ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`
}

func (r *ResolutionOfInvestigationV03) AddAssignment() *model.CaseAssignment2 {
	r.Assignment = new(model.CaseAssignment2)
	return r.Assignment
}

func (r *ResolutionOfInvestigationV03) AddResolvedCase() *model.Case2 {
	r.ResolvedCase = new(model.Case2)
	return r.ResolvedCase
}

func (r *ResolutionOfInvestigationV03) AddStatus() *model.InvestigationStatus2Choice {
	r.Status = new(model.InvestigationStatus2Choice)
	return r.Status
}

func (r *ResolutionOfInvestigationV03) AddCancellationDetails() *model.UnderlyingTransaction3 {
	newValue := new(model.UnderlyingTransaction3)
	r.CancellationDetails = append(r.CancellationDetails, newValue)
	return newValue
}

func (r *ResolutionOfInvestigationV03) AddStatementDetails() *model.StatementResolutionEntry1 {
	r.StatementDetails = new(model.StatementResolutionEntry1)
	return r.StatementDetails
}

func (r *ResolutionOfInvestigationV03) AddCorrectionTransaction() *model.CorrectiveTransaction1Choice {
	r.CorrectionTransaction = new(model.CorrectiveTransaction1Choice)
	return r.CorrectionTransaction
}

func (r *ResolutionOfInvestigationV03) AddResolutionRelatedInformation() *model.ResolutionInformation1 {
	r.ResolutionRelatedInformation = new(model.ResolutionInformation1)
	return r.ResolutionRelatedInformation
}

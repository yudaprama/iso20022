package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900106 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.06 Document"`
	Message *ResolutionOfInvestigationV06 `xml:"RsltnOfInvstgtn"`
}

func (d *Document02900106) AddMessage() *ResolutionOfInvestigationV06 {
	d.Message = new(ResolutionOfInvestigationV06)
	return d.Message
}

// Scope
// The ResolutionOfInvestigation message is sent by a case assignee to a case creator/case assigner.
// This message is used to inform of the resolution of a case, and optionally provides details about .
// - the corrective action undertaken by the case assignee
// - information on the return where applicable
// Usage
// The ResolutionOfInvestigation message is used by the case assignee to inform a case creator or case assigner about the resolution of a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// The ResolutionOfInvestigation message covers one and only one case at a time. If the case assignee needs to communicate about several cases, then several Resolution Of Investigation messages must be sent.
// The ResolutionOfInvestigation message provides:
// - the final outcome of the case, whether positive or negative
// - optionally, the details of the corrective action undertaken by the case assignee and the information of the return
// Whenever a payment instruction has been generated to solve the case under investigation following a claim non receipt or an unable to apply, the optional CorrectionTransaction component present in the message must be completed.
// Whenever the action of modifying or cancelling a payment results in funds being returned or reversed, an investigating agent may provide the details in the resolution related investigation component, to identify the return or reversal transaction. These details will facilitate the account reconciliations at the initiating bank and the intermediaries. It must be stressed that the return or reversal of funds is outside the scope of this Exceptions and Investigation service. The features given here is only meant to transmit the information of return or reversal when it is available through the resolution of the case.
// The ResolutionOfInvestigation message must:
// - be forwarded by all subsequent case assignee(s) until it reaches the case creator
// - not be used in place of a RejectCaseAssignment or CaseStatusReport or NotificationOfCaseAssignment message
// Take note of an exceptional rule that allows the use of ResolutionOfInvestigation in lieu of a CaseStatusReport. CaseStatusReport is a response-message to a CaseStatusReportRequest. The latter which is sent when the assigner has reached its own time-out threshold to receive a response. However it may happen that when the request arrives, the investigating agent has just obtained a resolution. In such a situation, it would be redundant to send a CaseStatusReport when then followed immediately by a ResolutionOfInvestigation. It is therefore quite acceptable for the investigating agent, the assignee, to skip the Case Status Report and send the ResolutionOfInvestigation message directly.
// The ResolutionOfInvestigation message should be the sole message to respond to a cancellation request. Details of the underlying transactions and the related statuses for which the cancellation request has been issued may be provided in the CancellationDetails component.
type ResolutionOfInvestigationV06 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies a resolved case.
	ResolvedCase *model.Case3 `xml:"RslvdCase,omitempty"`

	// Indicates the status of the investigation.
	Status *model.InvestigationStatus3Choice `xml:"Sts"`

	// Specifies the details of the underlying transactions being cancelled.
	CancellationDetails []*model.UnderlyingTransaction14 `xml:"CxlDtls,omitempty"`

	// Details on the underlying statement entry.
	StatementDetails *model.StatementResolutionEntry2 `xml:"StmtDtls,omitempty"`

	// References a transaction initiated to fix the case under investigation.
	CorrectionTransaction *model.CorrectiveTransaction1Choice `xml:"CrrctnTx,omitempty"`

	// Reference of a return or a reversal initiated to fix the case under investigation as part of the resolution.
	ResolutionRelatedInformation *model.ResolutionInformation1 `xml:"RsltnRltdInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *ResolutionOfInvestigationV06) AddAssignment() *model.CaseAssignment3 {
	r.Assignment = new(model.CaseAssignment3)
	return r.Assignment
}

func (r *ResolutionOfInvestigationV06) AddResolvedCase() *model.Case3 {
	r.ResolvedCase = new(model.Case3)
	return r.ResolvedCase
}

func (r *ResolutionOfInvestigationV06) AddStatus() *model.InvestigationStatus3Choice {
	r.Status = new(model.InvestigationStatus3Choice)
	return r.Status
}

func (r *ResolutionOfInvestigationV06) AddCancellationDetails() *model.UnderlyingTransaction14 {
	newValue := new(model.UnderlyingTransaction14)
	r.CancellationDetails = append(r.CancellationDetails, newValue)
	return newValue
}

func (r *ResolutionOfInvestigationV06) AddStatementDetails() *model.StatementResolutionEntry2 {
	r.StatementDetails = new(model.StatementResolutionEntry2)
	return r.StatementDetails
}

func (r *ResolutionOfInvestigationV06) AddCorrectionTransaction() *model.CorrectiveTransaction1Choice {
	r.CorrectionTransaction = new(model.CorrectiveTransaction1Choice)
	return r.CorrectionTransaction
}

func (r *ResolutionOfInvestigationV06) AddResolutionRelatedInformation() *model.ResolutionInformation1 {
	r.ResolutionRelatedInformation = new(model.ResolutionInformation1)
	return r.ResolutionRelatedInformation
}

func (r *ResolutionOfInvestigationV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}

package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000101 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.030.001.01 Document"`
	Message *AgentCADeactivationStatusAdviceV01 `xml:"AgtCADeactvtnStsAdvc"`
}

func (d *Document03000101) AddMessage() *AgentCADeactivationStatusAdviceV01 {
	d.Message = new(AgentCADeactivationStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to report the status, or a change in status, of a corporate action deactivation instruction or the status of a deactivation cancellation request.
// Usage
// This message is used to provide a status on the deactivation instruction, especially to confirm the deactivation of a Corporate Action event or option.
type AgentCADeactivationStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Deactivation Instruction for which a status is given.
	AgentCADeactivationInstructionIdentification *model.DocumentIdentification8 `xml:"AgtCADeactvtnInstrId"`

	// Identification of the linked Agent CA Deactivation Cancellation Request for which a status is given.
	AgentCADeactivationCancellationRequestIdentification *model.DocumentIdentification8 `xml:"AgtCADeactvtnCxlReqId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Status of the deactivation instruction sent by the issuer (agent).
	DeactivationInstructionStatus []*model.CorporateActionDeactivationInstructionStatus1 `xml:"DeactvtnInstrSts"`

	// Status of the deactivation cancellation request sent by the issuer (agent).
	DeactivationCancellationRequestStatus *model.CorporateActionDeactivationCancellationStatus1Choice `xml:"DeactvtnCxlReqSts"`
}

func (a *AgentCADeactivationStatusAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCADeactivationStatusAdviceV01) AddAgentCADeactivationInstructionIdentification() *model.DocumentIdentification8 {
	a.AgentCADeactivationInstructionIdentification = new(model.DocumentIdentification8)
	return a.AgentCADeactivationInstructionIdentification
}

func (a *AgentCADeactivationStatusAdviceV01) AddAgentCADeactivationCancellationRequestIdentification() *model.DocumentIdentification8 {
	a.AgentCADeactivationCancellationRequestIdentification = new(model.DocumentIdentification8)
	return a.AgentCADeactivationCancellationRequestIdentification
}

func (a *AgentCADeactivationStatusAdviceV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCADeactivationStatusAdviceV01) AddDeactivationInstructionStatus() *model.CorporateActionDeactivationInstructionStatus1 {
	newValue := new(model.CorporateActionDeactivationInstructionStatus1)
	a.DeactivationInstructionStatus = append(a.DeactivationInstructionStatus, newValue)
	return newValue
}

func (a *AgentCADeactivationStatusAdviceV01) AddDeactivationCancellationRequestStatus() *model.CorporateActionDeactivationCancellationStatus1Choice {
	a.DeactivationCancellationRequestStatus = new(model.CorporateActionDeactivationCancellationStatus1Choice)
	return a.DeactivationCancellationRequestStatus
}

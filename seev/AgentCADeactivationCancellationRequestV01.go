package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02900101 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.029.001.01 Document"`
	Message *AgentCADeactivationCancellationRequestV01 `xml:"AgtCADeactvtnCxlReq"`
}

func (d *Document02900101) AddMessage() *AgentCADeactivationCancellationRequestV01 {
	d.Message = new(AgentCADeactivationCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to the CSD to request the cancellation of a previously sent corporate action deactivation instruction.
// Usage
// This message is used to request the cancellation of a pending deactivation instruction. The cancellation must apply to exactly the same level as the original instruction, ie to the entire CA event or to an option as per the original instruction.
// This message must be sent before the deactivation execution date.
// In case a corporate action or option is already deactivated, this message can not be used to reactivate the corporate action entire event or option; the notification advice message must be used to reactivate a corporate action or option.
type AgentCADeactivationCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Deactivation Instruction to be cancelled.
	AgentCADeactivationInstructionIdentification *model.DocumentIdentification8 `xml:"AgtCADeactvtnInstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Details of the deactivation instruction to be cancelled.
	DeactivationInstructionDetails *model.CorporateActionDeactivationInstruction1 `xml:"DeactvtnInstrDtls,omitempty"`
}

func (a *AgentCADeactivationCancellationRequestV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCADeactivationCancellationRequestV01) AddAgentCADeactivationInstructionIdentification() *model.DocumentIdentification8 {
	a.AgentCADeactivationInstructionIdentification = new(model.DocumentIdentification8)
	return a.AgentCADeactivationInstructionIdentification
}

func (a *AgentCADeactivationCancellationRequestV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCADeactivationCancellationRequestV01) AddDeactivationInstructionDetails() *model.CorporateActionDeactivationInstruction1 {
	a.DeactivationInstructionDetails = new(model.CorporateActionDeactivationInstruction1)
	return a.DeactivationInstructionDetails
}

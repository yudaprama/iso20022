package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02000101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.020.001.01 Document"`
	Message *AgentCAMovementCancellationRequestV01 `xml:"AgtCAMvmntCxlReq"`
}

func (d *Document02000101) AddMessage() *AgentCAMovementCancellationRequestV01 {
	d.Message = new(AgentCAMovementCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to request the cancellation of (a) movement(s) previously sent via an Agent Corporate Action Movement Instruction.
// Usage
// This message may be used to cancel an entire Agent Corporate Action Movement Instruction message that was previously sent by the issuer (or its agent) or specific movements.
// This message must contain the identification of the Agent Corporate Action Movement Instruction containing the movement(s) to be cancelled, the agent identification and the corporate action references. This message must also contain details of the movement(s) to be cancelled.
type AgentCAMovementCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Movement Instruction to be cancelled.
	AgentCAMovementInstructionIdentification *model.DocumentIdentification8 `xml:"AgtCAMvmntInstrId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Details of the movement instructions to be cancelled.
	MovementDetails *model.MovementInstruction1 `xml:"MvmntDtls,omitempty"`
}

func (a *AgentCAMovementCancellationRequestV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAMovementCancellationRequestV01) AddAgentCAMovementInstructionIdentification() *model.DocumentIdentification8 {
	a.AgentCAMovementInstructionIdentification = new(model.DocumentIdentification8)
	return a.AgentCAMovementInstructionIdentification
}

func (a *AgentCAMovementCancellationRequestV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAMovementCancellationRequestV01) AddMovementDetails() *model.MovementInstruction1 {
	a.MovementDetails = new(model.MovementInstruction1)
	return a.MovementDetails
}

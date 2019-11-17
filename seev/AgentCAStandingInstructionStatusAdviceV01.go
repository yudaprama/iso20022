package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02700101 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.027.001.01 Document"`
	Message *AgentCAStandingInstructionStatusAdviceV01 `xml:"AgtCAStgInstrStsAdvc"`
}

func (d *Document02700101) AddMessage() *AgentCAStandingInstructionStatusAdviceV01 {
	d.Message = new(AgentCAStandingInstructionStatusAdviceV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to the CSD to report the status, or a change in status, of a standing instruction request or the status of a standing instruction cancellation request.
// Usage
// When this message is used to report the status of a standing instruction request, the building block Standing Instruction Request Identification must be present.
// When this message is used to report the status of a standing instruction cancellation request, the building block Standing Instruction Cancellation Request Identification must be present.
type AgentCAStandingInstructionStatusAdviceV01 struct {

	// Identification assigned by the Sender to unambiguously identify the status advice.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the linked Agent CA Standing Instruction Request for which a status is given.
	AgentCAStandingInstructionRequestIdentification *model.DocumentIdentification8 `xml:"AgtCAStgInstrReqId"`

	// Identification of the linked Agent CA Standing Instruction  Cancellation Request for which a status is given.
	//
	AgentCAStandingInstructionCancellationRequestIdentification *model.DocumentIdentification8 `xml:"AgtCAStgInstrCxlReqId"`

	// General information about the standing instruction.
	StandingInstructionGeneralInformation *model.CorporateActionStandingInstructionGeneralInformation1 `xml:"StgInstrGnlInf"`

	// Status of the standing instruction request.
	StandingInstructionRequestStatus *model.StandingInstructionStatus1Choice `xml:"StgInstrReqSts"`

	// Provides information about the status of a standing instruction cancellation request.
	StandingInstructionCancellationRequestStatus *model.StandingInstructionCancellationStatus1Choice `xml:"StgInstrCxlReqSts"`
}

func (a *AgentCAStandingInstructionStatusAdviceV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAStandingInstructionStatusAdviceV01) AddAgentCAStandingInstructionRequestIdentification() *model.DocumentIdentification8 {
	a.AgentCAStandingInstructionRequestIdentification = new(model.DocumentIdentification8)
	return a.AgentCAStandingInstructionRequestIdentification
}

func (a *AgentCAStandingInstructionStatusAdviceV01) AddAgentCAStandingInstructionCancellationRequestIdentification() *model.DocumentIdentification8 {
	a.AgentCAStandingInstructionCancellationRequestIdentification = new(model.DocumentIdentification8)
	return a.AgentCAStandingInstructionCancellationRequestIdentification
}

func (a *AgentCAStandingInstructionStatusAdviceV01) AddStandingInstructionGeneralInformation() *model.CorporateActionStandingInstructionGeneralInformation1 {
	a.StandingInstructionGeneralInformation = new(model.CorporateActionStandingInstructionGeneralInformation1)
	return a.StandingInstructionGeneralInformation
}

func (a *AgentCAStandingInstructionStatusAdviceV01) AddStandingInstructionRequestStatus() *model.StandingInstructionStatus1Choice {
	a.StandingInstructionRequestStatus = new(model.StandingInstructionStatus1Choice)
	return a.StandingInstructionRequestStatus
}

func (a *AgentCAStandingInstructionStatusAdviceV01) AddStandingInstructionCancellationRequestStatus() *model.StandingInstructionCancellationStatus1Choice {
	a.StandingInstructionCancellationRequestStatus = new(model.StandingInstructionCancellationStatus1Choice)
	return a.StandingInstructionCancellationRequestStatus
}

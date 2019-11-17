package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.014.001.01 Document"`
	Message *AgentCAElectionCancellationRequestV01 `xml:"AgtCAElctnCxlReq"`
}

func (d *Document01400101) AddMessage() *AgentCAElectionCancellationRequestV01 {
	d.Message = new(AgentCAElectionCancellationRequestV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to the issuer (or its agent) to request the cancellation of a previously sent Agent Corporate Action Election Advice message.
// Usage
// This message may only be used to cancel an entire Agent Corporate Action Election Advice message that was previously sent by the CSD. No partial cancellation is allowed.
// This message must contain the identification of the Agent Corporate Action Election Advice to be cancelled, the agent identification and the corporate action references. This message may also contain details of the election advice to be cancelled, but this is not recommended.
type AgentCAElectionCancellationRequestV01 struct {

	// Identification assigned by the Sender to unambiguously identify the cancellation request.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Election Advice to be cancelled.
	AgentCAElectionAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCAElctnAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides information about the election advice to be cancelled.
	ElectionDetails *model.CorporateActionElection3 `xml:"ElctnDtls,omitempty"`
}

func (a *AgentCAElectionCancellationRequestV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAElectionCancellationRequestV01) AddAgentCAElectionAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCAElectionAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCAElectionAdviceIdentification
}

func (a *AgentCAElectionCancellationRequestV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAElectionCancellationRequestV01) AddElectionDetails() *model.CorporateActionElection3 {
	a.ElectionDetails = new(model.CorporateActionElection3)
	return a.ElectionDetails
}

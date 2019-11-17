package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02100101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.021.001.01 Document"`
	Message *AgentCAMovementConfirmationV01 `xml:"AgtCAMvmntConf"`
}

func (d *Document02100101) AddMessage() *AgentCAMovementConfirmationV01 {
	d.Message = new(AgentCAMovementConfirmationV01)
	return d.Message
}

// Scope
// This message is sent by a CSD to an issuer (or its agent) to confirm the settlement of resource movements.
// Usage
// This message is used to confirm the settlement of the movements resulting from an:
// - Agent Corporate Action Movement Instruction message. It confirms the settlement of the exercised resources and/or proceeds movements in which case the building block Agent Corporate Action Movement Instruction Identification must be present. An Agent Corporate Action Movement Instruction message may be responded to by more than one Agent Corporate Action Movement Confirmation messages.
// - Agent Corporate Action Global Distribution Status Advice authorising the global distribution. It confirms the settlement of the exercised resources and/or proceeds movements. The building block Agent Corporate Action Global Distribution Status Advice Identification must be present. An Agent Corporate Action Global Distribution Status Advice message may be responded to by more than one Agent Corporate Action Movement Confirmation messages.
// - Agent Corporate Action Election Status Advice that rejects an election advice. It confirms the return of the exercised resources. The building block Agent Corporate Action Election Status Advice Identification must be present. An Agent Corporate Action Election Status Advice message may be responded to by more than one Agent Corporate Action Movement Confirmation messages.
// - Agent Corporate Action Election Status Advice where an election cancellation request has been accepted. It confirms the return of the exercised resources. The building block Agent Corporate Action Election Status Advice Identification must be present. An Agent Corporate Action Election Status Advice message may be responded to by more than one Agent Corporate Action Movement Confirmation messages.
type AgentCAMovementConfirmationV01 struct {

	// Identification assigned by the Sender to unambiguously identify the confirmation.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA Movement Instruction that triggered the movement(s).
	AgentCAMovementInstructionIdentification *model.DocumentIdentification8 `xml:"AgtCAMvmntInstrId"`

	// Identification of the Agent CA Election Status Advice that triggered the movement(s).
	AgentCAElectionStatusAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCAElctnStsAdvcId"`

	// Identification of the Agent CA Global Distribution Status Advice that triggered the movement(s).
	AgentCAGlobalDistributionStatusAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCAGblDstrbtnStsAdvcId"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Information about the securities movements.
	SecuritiesMovementDetails []*model.CorporateActionSecuritiesMovement1 `xml:"SctiesMvmntDtls,omitempty"`

	// Information about the cash movement.
	CashMovementDetails []*model.CashMovement3 `xml:"CshMvmntDtls,omitempty"`
}

func (a *AgentCAMovementConfirmationV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAMovementConfirmationV01) AddAgentCAMovementInstructionIdentification() *model.DocumentIdentification8 {
	a.AgentCAMovementInstructionIdentification = new(model.DocumentIdentification8)
	return a.AgentCAMovementInstructionIdentification
}

func (a *AgentCAMovementConfirmationV01) AddAgentCAElectionStatusAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCAElectionStatusAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCAElectionStatusAdviceIdentification
}

func (a *AgentCAMovementConfirmationV01) AddAgentCAGlobalDistributionStatusAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCAGlobalDistributionStatusAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCAGlobalDistributionStatusAdviceIdentification
}

func (a *AgentCAMovementConfirmationV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAMovementConfirmationV01) AddSecuritiesMovementDetails() *model.CorporateActionSecuritiesMovement1 {
	newValue := new(model.CorporateActionSecuritiesMovement1)
	a.SecuritiesMovementDetails = append(a.SecuritiesMovementDetails, newValue)
	return newValue
}

func (a *AgentCAMovementConfirmationV01) AddCashMovementDetails() *model.CashMovement3 {
	newValue := new(model.CashMovement3)
	a.CashMovementDetails = append(a.CashMovementDetails, newValue)
	return newValue
}

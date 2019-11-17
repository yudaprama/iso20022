package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.019.001.01 Document"`
	Message *AgentCAMovementInstructionV01 `xml:"AgtCAMvmntInstr"`
}

func (d *Document01900101) AddMessage() *AgentCAMovementInstructionV01 {
	d.Message = new(AgentCAMovementInstructionV01)
	return d.Message
}

// Scope
// This message is sent by an issuer (or its agent) to a CSD to order:
// - the global or individual debit of exercised resources (cash and/or securities), per event and optionally per option and per resource for all or individual CSD client's accounts;
// - and/or the individual credits of the outturn resources per event and optionally per option and per resource for a given CSD client's account.
// Usage
// This message is used to instruct:
// - the global debit of the exercised resources from the CSD client's available or sequestered balance, in which case, the order type must be 'global debit order';
// - the individual debits and credits:
// - the individual debit of the exercised resources from the CSD client's available or sequestered balance and/or
// - the individual credit of the outturn resources to the CSD client's account.
// The order type must be 'individual order';
// - a return order, in the case of a scaleback, i.e. the return of the exercised resources to the CSD client's account. The order type must be either 'global return order' or 'individual return order'; and
// change of option, e.g. in the case of the closure of an option, by moving the exercised resources from one option to another option within the sequestered balances in accordance to the new option conditions. The order type must be 'option change order'.
type AgentCAMovementInstructionV01 struct {

	// Identification assigned by the Sender to unambiguously identify the instruction.
	Identification *model.DocumentIdentification8 `xml:"Id"`

	// Identification of the Agent CA ElectionAdvice when the movements are the result of an ElectionAdvice.
	AgentCAElectionAdviceIdentification *model.DocumentIdentification8 `xml:"AgtCAElctnAdvcId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionInformation1 `xml:"CorpActnGnlInf"`

	// Provides general information about the movement.
	MovementGeneralInformation *model.CorporateActionMovement1 `xml:"MvmntGnlInf"`

	// Information related to the movement of the underlying securities.
	UnderlyingSecuritiesMovementDetails []*model.UnderlyingSecurityMovement1 `xml:"UndrlygSctiesMvmntDtls,omitempty"`

	// Information related to the movement of the underlying cash.
	UnderlyingCashMovementDetails []*model.CashMovement2 `xml:"UndrlygCshMvmntDtls,omitempty"`

	// Information related to the movement of the CA proceeds.
	ProceedsMovementDetails *model.ProceedsMovement1 `xml:"PrcdsMvmntDtls,omitempty"`
}

func (a *AgentCAMovementInstructionV01) AddIdentification() *model.DocumentIdentification8 {
	a.Identification = new(model.DocumentIdentification8)
	return a.Identification
}

func (a *AgentCAMovementInstructionV01) AddAgentCAElectionAdviceIdentification() *model.DocumentIdentification8 {
	a.AgentCAElectionAdviceIdentification = new(model.DocumentIdentification8)
	return a.AgentCAElectionAdviceIdentification
}

func (a *AgentCAMovementInstructionV01) AddCorporateActionGeneralInformation() *model.CorporateActionInformation1 {
	a.CorporateActionGeneralInformation = new(model.CorporateActionInformation1)
	return a.CorporateActionGeneralInformation
}

func (a *AgentCAMovementInstructionV01) AddMovementGeneralInformation() *model.CorporateActionMovement1 {
	a.MovementGeneralInformation = new(model.CorporateActionMovement1)
	return a.MovementGeneralInformation
}

func (a *AgentCAMovementInstructionV01) AddUnderlyingSecuritiesMovementDetails() *model.UnderlyingSecurityMovement1 {
	newValue := new(model.UnderlyingSecurityMovement1)
	a.UnderlyingSecuritiesMovementDetails = append(a.UnderlyingSecuritiesMovementDetails, newValue)
	return newValue
}

func (a *AgentCAMovementInstructionV01) AddUnderlyingCashMovementDetails() *model.CashMovement2 {
	newValue := new(model.CashMovement2)
	a.UnderlyingCashMovementDetails = append(a.UnderlyingCashMovementDetails, newValue)
	return newValue
}

func (a *AgentCAMovementInstructionV01) AddProceedsMovementDetails() *model.ProceedsMovement1 {
	a.ProceedsMovementDetails = new(model.ProceedsMovement1)
	return a.ProceedsMovementDetails
}

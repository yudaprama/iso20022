package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300106 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.003.001.06 Document"`
	Message *AccountModificationInstructionV06 `xml:"AcctModInstr"`
}

func (d *Document00300106) AddMessage() *AccountModificationInstructionV06 {
	d.Message = new(AccountModificationInstructionV06)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent, sends the AccountModificationInstruction message to the account servicer, for example, a registrar, transfer agent, custodian bank or securities depository to modify, that is, create, update or delete specific details of an existing account.
// Usage
// The AccountModificationInstruction message is used to modify the details of an existing account.
// The AccountModificationInstruction message has three specific uses:
// - to maintain/update any of the existing account details, for example, to update the address of the beneficiary or modify the preference to income from distribution to capitalisation, or,
// - to add/create specific details to the existing account when these details were not yet recorded at the time of account creation, for example, to add a second address or to establish new cash settlement standing instructions, or,
// - to delete specific account details, for example, delete cash standing instructions.
// This message cannot be used to delete an entire account, as institution specific and regulatory rules pertaining to account deletion are diverse.
// The usage of this message may be subject to service level agreement (SLA) between the counterparties.
// Execution of the AccountModificationInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountModificationInstructionV06 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Information about the modification instruction.
	InstructionDetails *model.InvestmentAccountModification2 `xml:"InstrDtls,omitempty"`

	// Identifies the account for which the information is modified.
	InvestmentAccountSelection *model.AccountSelection1Choice `xml:"InvstmtAcctSelctn"`

	// Information related to general characteristics of the account to be inserted, updated or deleted.
	ModifiedInvestmentAccount *model.InvestmentAccount51 `xml:"ModfdInvstmtAcct,omitempty"`

	// Information related to an account  party to be inserted, updated or deleted.
	ModifiedAccountParties []*model.AccountParties14 `xml:"ModfdAcctPties,omitempty"`

	// Information related to intermediaries to be inserted, updated or deleted.
	ModifiedIntermediaries []*model.ModificationScope26 `xml:"ModfdIntrmies,omitempty"`

	// Information related to referral information to be inserted, updated or deleted
	ModifiedPlacement *model.ModificationScope33 `xml:"ModfdPlcmnt,omitempty"`

	// Eligibility conditions related to allocation of new issues to be inserted, updated or deleted.
	ModifiedIssueAllocation *model.ModificationScope21 `xml:"ModfdIsseAllcn,omitempty"`

	// Information related to a savings plan to be either inserted, updated or deleted.
	ModifiedSavingsInvestmentPlan []*model.ModificationScope28 `xml:"ModfdSvgsInvstmtPlan,omitempty"`

	// Information related to a withdrawal plan to be either inserted, updated or deleted.
	ModifiedWithdrawalInvestmentPlan []*model.ModificationScope28 `xml:"ModfdWdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction to be either inserted or deleted.
	ModifiedCashSettlement []*model.CashSettlement2 `xml:"ModfdCshSttlm,omitempty"`

	// Information related to documents to be added, deleted or updated.
	//
	ModifiedServiceLevelAgreement []*model.ModificationScope31 `xml:"ModfdSvcLvlAgrmt,omitempty"`

	// Additional information concerning limitations and restrictions on the account to be inserted, updated or deleted.
	ModifiedAdditionalInformation []*model.ModificationScope30 `xml:"ModfdAddtlInf,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountModificationInstructionV06) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountModificationInstructionV06) AddPreviousReference() *model.AdditionalReference6 {
	a.PreviousReference = new(model.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountModificationInstructionV06) AddInstructionDetails() *model.InvestmentAccountModification2 {
	a.InstructionDetails = new(model.InvestmentAccountModification2)
	return a.InstructionDetails
}

func (a *AccountModificationInstructionV06) AddInvestmentAccountSelection() *model.AccountSelection1Choice {
	a.InvestmentAccountSelection = new(model.AccountSelection1Choice)
	return a.InvestmentAccountSelection
}

func (a *AccountModificationInstructionV06) AddModifiedInvestmentAccount() *model.InvestmentAccount51 {
	a.ModifiedInvestmentAccount = new(model.InvestmentAccount51)
	return a.ModifiedInvestmentAccount
}

func (a *AccountModificationInstructionV06) AddModifiedAccountParties() *model.AccountParties14 {
	newValue := new(model.AccountParties14)
	a.ModifiedAccountParties = append(a.ModifiedAccountParties, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddModifiedIntermediaries() *model.ModificationScope26 {
	newValue := new(model.ModificationScope26)
	a.ModifiedIntermediaries = append(a.ModifiedIntermediaries, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddModifiedPlacement() *model.ModificationScope33 {
	a.ModifiedPlacement = new(model.ModificationScope33)
	return a.ModifiedPlacement
}

func (a *AccountModificationInstructionV06) AddModifiedIssueAllocation() *model.ModificationScope21 {
	a.ModifiedIssueAllocation = new(model.ModificationScope21)
	return a.ModifiedIssueAllocation
}

func (a *AccountModificationInstructionV06) AddModifiedSavingsInvestmentPlan() *model.ModificationScope28 {
	newValue := new(model.ModificationScope28)
	a.ModifiedSavingsInvestmentPlan = append(a.ModifiedSavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddModifiedWithdrawalInvestmentPlan() *model.ModificationScope28 {
	newValue := new(model.ModificationScope28)
	a.ModifiedWithdrawalInvestmentPlan = append(a.ModifiedWithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddModifiedCashSettlement() *model.CashSettlement2 {
	newValue := new(model.CashSettlement2)
	a.ModifiedCashSettlement = append(a.ModifiedCashSettlement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddModifiedServiceLevelAgreement() *model.ModificationScope31 {
	newValue := new(model.ModificationScope31)
	a.ModifiedServiceLevelAgreement = append(a.ModifiedServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddModifiedAdditionalInformation() *model.ModificationScope30 {
	newValue := new(model.ModificationScope30)
	a.ModifiedAdditionalInformation = append(a.ModifiedAdditionalInformation, newValue)
	return newValue
}

func (a *AccountModificationInstructionV06) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountModificationInstructionV06) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

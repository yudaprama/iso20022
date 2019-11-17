package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100107 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.001.001.07 Document"`
	Message *AccountOpeningInstructionV07 `xml:"AcctOpngInstr"`
}

func (d *Document00100107) AddMessage() *AccountOpeningInstructionV07 {
	d.Message = new(AccountOpeningInstructionV07)
	return d.Message
}

// Scope
// The AccountOpeningInstruction message is sent by an account owner, for example, an investor or its designated agent to the account servicer, for example, a registrar, transfer agent, custodian or securities depository, to instruct the opening of an account or the opening of an account and the establishment of an investment plan.
// Usage
// The AccountOpeningInstruction is used to open an account directly or indirectly with the account servicer or an intermediary.
// In some markets, for example, Australia, and for some products in the United Kingdom, a first order (also known as a deposit instruction) is placed at the same time as the account opening. To cater for this scenario, an order message can be linked (via references in the message) to the AccountOpeningInstruction message when needed.
// Execution of the AccountOpeningInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountOpeningInstructionV07 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order or settlement transaction.
	OrderReference *model.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Information about the opening instruction.
	InstructionDetails *model.InvestmentAccountOpening3 `xml:"InstrDtls"`

	// Detailed information about the account to be opened.
	InvestmentAccount *model.InvestmentAccount61 `xml:"InvstmtAcct"`

	// Information related to parties that are related to the account, for example, primary account owner.
	AccountParties *model.AccountParties15 `xml:"AcctPties"`

	// Intermediary or other party related to the management of the account.
	Intermediaries []*model.Intermediary36 `xml:"Intrmies,omitempty"`

	// Referral information.
	Placement *model.ReferredAgent2 `xml:"Plcmnt,omitempty"`

	// Eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *model.NewIssueAllocation2 `xml:"NewIsseAllcn,omitempty"`

	// Plan that allows individuals to set aside a fixed amount of money at specified intervals, usually for a special purpose, for example, retirement.
	SavingsInvestmentPlan []*model.InvestmentPlan14 `xml:"SvgsInvstmtPlan,omitempty"`

	// Plan through which holdings are depleted through regular withdrawals at specified intervals.
	WithdrawalInvestmentPlan []*model.InvestmentPlan14 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction associated to  transactions on the account.
	CashSettlement []*model.CashSettlement1 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*model.DocumentToSend3 `xml:"SvcLvlAgrmt,omitempty"`

	// Additional information such as remarks or notes that must be conveyed about the account management activity and or any limitations and restrictions.
	AdditionalInformation []*model.AdditiononalInformation12 `xml:"AddtlInf,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountOpeningInstructionV07) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountOpeningInstructionV07) AddOrderReference() *model.InvestmentFundOrder4 {
	a.OrderReference = new(model.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountOpeningInstructionV07) AddPreviousReference() *model.AdditionalReference6 {
	a.PreviousReference = new(model.AdditionalReference6)
	return a.PreviousReference
}

func (a *AccountOpeningInstructionV07) AddInstructionDetails() *model.InvestmentAccountOpening3 {
	a.InstructionDetails = new(model.InvestmentAccountOpening3)
	return a.InstructionDetails
}

func (a *AccountOpeningInstructionV07) AddInvestmentAccount() *model.InvestmentAccount61 {
	a.InvestmentAccount = new(model.InvestmentAccount61)
	return a.InvestmentAccount
}

func (a *AccountOpeningInstructionV07) AddAccountParties() *model.AccountParties15 {
	a.AccountParties = new(model.AccountParties15)
	return a.AccountParties
}

func (a *AccountOpeningInstructionV07) AddIntermediaries() *model.Intermediary36 {
	newValue := new(model.Intermediary36)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV07) AddPlacement() *model.ReferredAgent2 {
	a.Placement = new(model.ReferredAgent2)
	return a.Placement
}

func (a *AccountOpeningInstructionV07) AddNewIssueAllocation() *model.NewIssueAllocation2 {
	a.NewIssueAllocation = new(model.NewIssueAllocation2)
	return a.NewIssueAllocation
}

func (a *AccountOpeningInstructionV07) AddSavingsInvestmentPlan() *model.InvestmentPlan14 {
	newValue := new(model.InvestmentPlan14)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV07) AddWithdrawalInvestmentPlan() *model.InvestmentPlan14 {
	newValue := new(model.InvestmentPlan14)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV07) AddCashSettlement() *model.CashSettlement1 {
	newValue := new(model.CashSettlement1)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV07) AddServiceLevelAgreement() *model.DocumentToSend3 {
	newValue := new(model.DocumentToSend3)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV07) AddAdditionalInformation() *model.AdditiononalInformation12 {
	newValue := new(model.AdditiononalInformation12)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV07) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountOpeningInstructionV07) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

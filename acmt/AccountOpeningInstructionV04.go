package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100104 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.001.001.04 Document"`
	Message *AccountOpeningInstructionV04 `xml:"AcctOpngInstr"`
}

func (d *Document00100104) AddMessage() *AccountOpeningInstructionV04 {
	d.Message = new(AccountOpeningInstructionV04)
	return d.Message
}

// Scope
// An account owner, for example, an investor or its designated agent sends the AccountOpeningInstruction message to the account servicer, for example, a registrar, transfer agent or custodian to instruct the opening of an account or the opening of an account and establishing an investment plan.
// Usage
// The AccountOpeningInstruction is used to open an account directly or indirectly with the account servicer or an intermediary.
// In some markets, for example, Australia, and for some products in the United Kingdom, a first order (also known as a deposit instruction) is placed at the same time as the account opening. To cater for this scenario, an order message can be linked (via references in the message) to the AccountOpeningInstruction message when needed.
// Execution of the AccountOpeningInstruction is confirmed via an AccountDetailsConfirmation message.
type AccountOpeningInstructionV04 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order.
	OrderReference *model.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Provides detailed information about the opening instruction.
	InstructionDetails *model.InvestmentAccountOpeningDetails `xml:"InstrDtls"`

	// Detailed information about the investment account to be opened.
	InvestmentAccount *model.InvestmentAccount37 `xml:"InvstmtAcct"`

	// Information related to parties who are related to an investment account, for example, primary account owner.
	AccountParties *model.AccountParties9 `xml:"AcctPties"`

	// Information related to an intermediary.
	Intermediaries []*model.Intermediary24 `xml:"Intrmies,omitempty"`

	// Placement agent for the hedge fund industry.
	Placement *model.ReferredAgent1 `xml:"Plcmnt,omitempty"`

	// Eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *model.NewIssueAllocation1 `xml:"NewIsseAllcn,omitempty"`

	// Plan that allows individuals to set aside a fixed amount of money at specified intervals, usually for a special purpose, for example, retirement.
	SavingsInvestmentPlan []*model.InvestmentPlan9 `xml:"SvgsInvstmtPlan,omitempty"`

	// Plan through which an investment fund investor's holdings are depleted through regular withdrawals at specified intervals.
	WithdrawalInvestmentPlan []*model.InvestmentPlan9 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Cash settlement standing instruction associated to the investment fund transaction.
	CashSettlement []*model.InvestmentFundCashSettlementInformation7 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*model.DocumentToSend1 `xml:"SvcLvlAgrmt,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountOpeningInstructionV04) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountOpeningInstructionV04) AddOrderReference() *model.InvestmentFundOrder4 {
	a.OrderReference = new(model.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountOpeningInstructionV04) AddPreviousReference() *model.AdditionalReference3 {
	a.PreviousReference = new(model.AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountOpeningInstructionV04) AddInstructionDetails() *model.InvestmentAccountOpeningDetails {
	a.InstructionDetails = new(model.InvestmentAccountOpeningDetails)
	return a.InstructionDetails
}

func (a *AccountOpeningInstructionV04) AddInvestmentAccount() *model.InvestmentAccount37 {
	a.InvestmentAccount = new(model.InvestmentAccount37)
	return a.InvestmentAccount
}

func (a *AccountOpeningInstructionV04) AddAccountParties() *model.AccountParties9 {
	a.AccountParties = new(model.AccountParties9)
	return a.AccountParties
}

func (a *AccountOpeningInstructionV04) AddIntermediaries() *model.Intermediary24 {
	newValue := new(model.Intermediary24)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV04) AddPlacement() *model.ReferredAgent1 {
	a.Placement = new(model.ReferredAgent1)
	return a.Placement
}

func (a *AccountOpeningInstructionV04) AddNewIssueAllocation() *model.NewIssueAllocation1 {
	a.NewIssueAllocation = new(model.NewIssueAllocation1)
	return a.NewIssueAllocation
}

func (a *AccountOpeningInstructionV04) AddSavingsInvestmentPlan() *model.InvestmentPlan9 {
	newValue := new(model.InvestmentPlan9)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV04) AddWithdrawalInvestmentPlan() *model.InvestmentPlan9 {
	newValue := new(model.InvestmentPlan9)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV04) AddCashSettlement() *model.InvestmentFundCashSettlementInformation7 {
	newValue := new(model.InvestmentFundCashSettlementInformation7)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV04) AddServiceLevelAgreement() *model.DocumentToSend1 {
	newValue := new(model.DocumentToSend1)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountOpeningInstructionV04) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountOpeningInstructionV04) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

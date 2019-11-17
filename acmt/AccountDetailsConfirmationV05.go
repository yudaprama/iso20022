package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200105 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.002.001.05 Document"`
	Message *AccountDetailsConfirmationV05 `xml:"AcctDtlsConf"`
}

func (d *Document00200105) AddMessage() *AccountDetailsConfirmationV05 {
	d.Message = new(AccountDetailsConfirmationV05)
	return d.Message
}

// Scope
// An account servicer, for example, a registrar, transfer agent or custodian bank sends the AccountDetailsConfirmation message to the account owner, for example, an investor to confirm the opening of an investment fund account, execution of an AccountModificationInstruction or to return information requested in a GetAccountDetails message.
// Usage
// The AccountDetailsConfirmation message is used to confirm the opening of an account, modification of an account or the provision of information requested in a previously sent GetAccountDetails message. The message contains detailed information relevant to the opened account.
// When the AccountDetailsConfirmation is used to confirm execution of an AccountModificationInstruction message, it contains the modified subsets of account details that were specified in the AccountModificationInstruction.
// When the AccountDetailsConfirmation is used to reply to a GetAccountDetails message, it returns the selected subsets of account details that were specified in the GetAccountDetails message.
type AccountDetailsConfirmationV05 struct {

	// Identifies the message.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order.
	OrderReference *model.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Provides detailed information about the request or instruction which triggered this confirmation.
	ConfirmationDetails *model.AccountManagementConfirmation2 `xml:"ConfDtls"`

	// Confirmation of the information related to a selected investment account.
	InvestmentAccount *model.InvestmentAccount46 `xml:"InvstmtAcct,omitempty"`

	// Confirmation of information related to parties who are related to a selected investment account.
	AccountParties *model.AccountParties12 `xml:"AcctPties,omitempty"`

	// Confirmation of information related to intermediaries who are related to a selected investment account.
	Intermediaries []*model.Intermediary24 `xml:"Intrmies,omitempty"`

	// Placement agent for the hedge fund industry.
	Placement *model.ReferredAgent1 `xml:"Plcmnt,omitempty"`

	// Eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *model.NewIssueAllocation2 `xml:"NewIsseAllcn,omitempty"`

	// Confirmation of the information related to a savings plan that is related to a selected investment account.
	SavingsInvestmentPlan []*model.InvestmentPlan10 `xml:"SvgsInvstmtPlan,omitempty"`

	// Confirmation of the information related to a withdrawal plan that is related to a selected investment account.
	WithdrawalInvestmentPlan []*model.InvestmentPlan10 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Confirmation of the cash settlement standing instruction associated to the investment fund transaction.
	CashSettlement []*model.InvestmentFundCashSettlementInformation7 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*model.DocumentToSend2 `xml:"SvcLvlAgrmt,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountDetailsConfirmationV05) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountDetailsConfirmationV05) AddOrderReference() *model.InvestmentFundOrder4 {
	a.OrderReference = new(model.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountDetailsConfirmationV05) AddRelatedReference() *model.AdditionalReference3 {
	a.RelatedReference = new(model.AdditionalReference3)
	return a.RelatedReference
}

func (a *AccountDetailsConfirmationV05) AddConfirmationDetails() *model.AccountManagementConfirmation2 {
	a.ConfirmationDetails = new(model.AccountManagementConfirmation2)
	return a.ConfirmationDetails
}

func (a *AccountDetailsConfirmationV05) AddInvestmentAccount() *model.InvestmentAccount46 {
	a.InvestmentAccount = new(model.InvestmentAccount46)
	return a.InvestmentAccount
}

func (a *AccountDetailsConfirmationV05) AddAccountParties() *model.AccountParties12 {
	a.AccountParties = new(model.AccountParties12)
	return a.AccountParties
}

func (a *AccountDetailsConfirmationV05) AddIntermediaries() *model.Intermediary24 {
	newValue := new(model.Intermediary24)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV05) AddPlacement() *model.ReferredAgent1 {
	a.Placement = new(model.ReferredAgent1)
	return a.Placement
}

func (a *AccountDetailsConfirmationV05) AddNewIssueAllocation() *model.NewIssueAllocation2 {
	a.NewIssueAllocation = new(model.NewIssueAllocation2)
	return a.NewIssueAllocation
}

func (a *AccountDetailsConfirmationV05) AddSavingsInvestmentPlan() *model.InvestmentPlan10 {
	newValue := new(model.InvestmentPlan10)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV05) AddWithdrawalInvestmentPlan() *model.InvestmentPlan10 {
	newValue := new(model.InvestmentPlan10)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV05) AddCashSettlement() *model.InvestmentFundCashSettlementInformation7 {
	newValue := new(model.InvestmentFundCashSettlementInformation7)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV05) AddServiceLevelAgreement() *model.DocumentToSend2 {
	newValue := new(model.DocumentToSend2)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV05) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountDetailsConfirmationV05) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

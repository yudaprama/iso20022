package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200106 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.002.001.06 Document"`
	Message *AccountDetailsConfirmationV06 `xml:"AcctDtlsConf"`
}

func (d *Document00200106) AddMessage() *AccountDetailsConfirmationV06 {
	d.Message = new(AccountDetailsConfirmationV06)
	return d.Message
}

// Scope
// An account servicer, for example, a registrar, transfer agent, custodian bank or securities depository  sends the AccountDetailsConfirmation message to the account owner, for example, an investor to confirm the opening of an account, execution of an AccountModificationInstruction or to return information requested in a GetAccountDetails message.
// Usage
// The AccountDetailsConfirmation message is used to confirm the opening of an account, modification of an account or the provision of information requested in a previously sent GetAccountDetails message. The message contains detailed information relevant to the opened account.
// When the AccountDetailsConfirmation is used to confirm execution of an AccountModificationInstruction message, it contains the modified subsets of account details that were specified in the AccountModificationInstruction.
// When the AccountDetailsConfirmation is used to reply to a GetAccountDetails message, it returns the selected subsets of account details that were specified in the GetAccountDetails message.
type AccountDetailsConfirmationV06 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identifies a related order or settlement transaction.
	OrderReference *model.InvestmentFundOrder4 `xml:"OrdrRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference6 `xml:"RltdRef,omitempty"`

	// Information about the request or instruction which triggered this confirmation.
	ConfirmationDetails *model.AccountManagementConfirmation3 `xml:"ConfDtls"`

	// Confirmation of the information related to the investment account.
	InvestmentAccount *model.InvestmentAccount50 `xml:"InvstmtAcct,omitempty"`

	// Confirmation of information related to parties that are related to the account, for example, primary account owner.
	AccountParties *model.AccountParties13 `xml:"AcctPties,omitempty"`

	// Confirmation of an intermediary or other party related to the management of the account. In some markets, when this intermediary is a party acting on behalf of the investor for which it has opened an account at, for example, a central securities depository or international central securities depository, this party is known by the investor as the 'account controller'.
	Intermediaries []*model.Intermediary36 `xml:"Intrmies,omitempty"`

	// Confirmation of referral information.
	Placement *model.ReferredAgent2 `xml:"Plcmnt,omitempty"`

	// Confirmation of eligibility conditions applicable when there is an allocation of new issues for hedge fund account opening.
	NewIssueAllocation *model.NewIssueAllocation2 `xml:"NewIsseAllcn,omitempty"`

	// Confirmation of the information related to a savings plan that is related to the account.
	SavingsInvestmentPlan []*model.InvestmentPlan12 `xml:"SvgsInvstmtPlan,omitempty"`

	// Confirmation of the information related to a withdrawal plan that is related to the account.
	WithdrawalInvestmentPlan []*model.InvestmentPlan12 `xml:"WdrwlInvstmtPlan,omitempty"`

	// Confirmation of a cash settlement standing instruction associated to  transactions on the account.
	CashSettlement []*model.CashSettlement1 `xml:"CshSttlm,omitempty"`

	// Identifies documents to be provided for the account opening.
	ServiceLevelAgreement []*model.DocumentToSend3 `xml:"SvcLvlAgrmt,omitempty"`

	// Additional information concerning limitations and restrictions on the account.
	AdditionalInformation []*model.AccountRestrictions1 `xml:"AddtlInf,omitempty"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountDetailsConfirmationV06) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountDetailsConfirmationV06) AddOrderReference() *model.InvestmentFundOrder4 {
	a.OrderReference = new(model.InvestmentFundOrder4)
	return a.OrderReference
}

func (a *AccountDetailsConfirmationV06) AddRelatedReference() *model.AdditionalReference6 {
	a.RelatedReference = new(model.AdditionalReference6)
	return a.RelatedReference
}

func (a *AccountDetailsConfirmationV06) AddConfirmationDetails() *model.AccountManagementConfirmation3 {
	a.ConfirmationDetails = new(model.AccountManagementConfirmation3)
	return a.ConfirmationDetails
}

func (a *AccountDetailsConfirmationV06) AddInvestmentAccount() *model.InvestmentAccount50 {
	a.InvestmentAccount = new(model.InvestmentAccount50)
	return a.InvestmentAccount
}

func (a *AccountDetailsConfirmationV06) AddAccountParties() *model.AccountParties13 {
	a.AccountParties = new(model.AccountParties13)
	return a.AccountParties
}

func (a *AccountDetailsConfirmationV06) AddIntermediaries() *model.Intermediary36 {
	newValue := new(model.Intermediary36)
	a.Intermediaries = append(a.Intermediaries, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV06) AddPlacement() *model.ReferredAgent2 {
	a.Placement = new(model.ReferredAgent2)
	return a.Placement
}

func (a *AccountDetailsConfirmationV06) AddNewIssueAllocation() *model.NewIssueAllocation2 {
	a.NewIssueAllocation = new(model.NewIssueAllocation2)
	return a.NewIssueAllocation
}

func (a *AccountDetailsConfirmationV06) AddSavingsInvestmentPlan() *model.InvestmentPlan12 {
	newValue := new(model.InvestmentPlan12)
	a.SavingsInvestmentPlan = append(a.SavingsInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV06) AddWithdrawalInvestmentPlan() *model.InvestmentPlan12 {
	newValue := new(model.InvestmentPlan12)
	a.WithdrawalInvestmentPlan = append(a.WithdrawalInvestmentPlan, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV06) AddCashSettlement() *model.CashSettlement1 {
	newValue := new(model.CashSettlement1)
	a.CashSettlement = append(a.CashSettlement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV06) AddServiceLevelAgreement() *model.DocumentToSend3 {
	newValue := new(model.DocumentToSend3)
	a.ServiceLevelAgreement = append(a.ServiceLevelAgreement, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV06) AddAdditionalInformation() *model.AccountRestrictions1 {
	newValue := new(model.AccountRestrictions1)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}

func (a *AccountDetailsConfirmationV06) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	a.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return a.MarketPracticeVersion
}

func (a *AccountDetailsConfirmationV06) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

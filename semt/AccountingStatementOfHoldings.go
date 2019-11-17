package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.003.001.01 Document"`
	Message *AccountingStatementOfHoldings `xml:"semt.003.001.01"`
}

func (d *Document00300101) AddMessage() *AccountingStatementOfHoldings {
	d.Message = new(AccountingStatementOfHoldings)
	return d.Message
}

// Scope
// The AccountingStatementOfHoldings message is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a local agent acting on behalf of its global custodian customer, a custodian acting on behalf of an investment management institution or a broker/dealer, a fund administrator or fund intermediary, trustee or registrar, etc.
// This message provides, at a specified moment in time, valuations of the portfolio together with details of each financial instrument holding.
// The information in the message can be audited or un-audited.
// Usage
// The AccountingStatementOfHoldings message can be sent:
// - At a frequency agreed bi-laterally between the Sender and the Receiver
// - As a response to a request for statement sent by the account owner. The request for statement message will be developed at a later stage.
// This message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, the message can be used to either specify holdings at
// - the main account level, or
// - the sub-account level.
// This message can be used to report where the securities are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// The AccountingStatementOfHoldings message must not be used to reconcile the books of the account owner and the account servicer. The CustodyStatementOfHoldings message is used for reconciliation purposes.
// The AccountingStatementOfHoldings message must not be used for trading purposes.
// Since a SWIFT message as sent is restricted to the maximum input message length, several messages may be needed to accommodate all the information.
type AccountingStatementOfHoldings struct {

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// General information related to the accounting statement of holdings.
	StatementGeneralDetails *model.Statement4 `xml:"StmtGnlDtls"`

	// The safekeeping or investment account.
	AccountDetails *model.SafekeepingAccount1 `xml:"AcctDtls"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation2 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification2 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *model.TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountingStatementOfHoldings) AddPreviousReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	a.PreviousReference = append(a.PreviousReference, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	a.RelatedReference = append(a.RelatedReference, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings) AddMessagePagination() *model.Pagination {
	a.MessagePagination = new(model.Pagination)
	return a.MessagePagination
}

func (a *AccountingStatementOfHoldings) AddStatementGeneralDetails() *model.Statement4 {
	a.StatementGeneralDetails = new(model.Statement4)
	return a.StatementGeneralDetails
}

func (a *AccountingStatementOfHoldings) AddAccountDetails() *model.SafekeepingAccount1 {
	a.AccountDetails = new(model.SafekeepingAccount1)
	return a.AccountDetails
}

func (a *AccountingStatementOfHoldings) AddBalanceForAccount() *model.AggregateBalanceInformation2 {
	newValue := new(model.AggregateBalanceInformation2)
	a.BalanceForAccount = append(a.BalanceForAccount, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings) AddSubAccountDetails() *model.SubAccountIdentification2 {
	newValue := new(model.SubAccountIdentification2)
	a.SubAccountDetails = append(a.SubAccountDetails, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings) AddTotalValues() *model.TotalValueInPageAndStatement {
	a.TotalValues = new(model.TotalValueInPageAndStatement)
	return a.TotalValues
}

func (a *AccountingStatementOfHoldings) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

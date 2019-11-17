package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.003.001.02 Document"`
	Message *AccountingStatementOfHoldingsV02 `xml:"AcctgStmtOfHldgsV02"`
}

func (d *Document00300102) AddMessage() *AccountingStatementOfHoldingsV02 {
	d.Message = new(AccountingStatementOfHoldingsV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent sends the AccountStatementOfHoldings message to the account owner, for example, a fund manager or an account owner's designated agent to provide detailed holdings of the portfolio at a specified moment in time.
// The message provides, at a moment in time, valuations of the portfolio together with details of each financial instrument holding.
// The message can be sent either audited or un-audited and may be provided on a trade date or settlement date basis.
// Usage
// The AccountingStatementOfHoldings message is used to provide valuation detail for each financial instrument held in a portfolio. The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner.
// This message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, the message can be used to either specify holdings at
// - the main account level, or,
// - the sub-account level.
// This message can be used to report where the financial instruments are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// The AccountingStatementOfHoldings message should not be used for trading purposes.
type AccountingStatementOfHoldingsV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// General information related to the accounting statement of holdings.
	StatementGeneralDetails *model.Statement6 `xml:"StmtGnlDtls"`

	// The safekeeping or investment account.
	AccountDetails *model.SafekeepingAccount2 `xml:"AcctDtls"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation3 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification3 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *model.TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountingStatementOfHoldingsV02) AddMessageIdentification() *model.MessageIdentification1 {
	a.MessageIdentification = new(model.MessageIdentification1)
	return a.MessageIdentification
}

func (a *AccountingStatementOfHoldingsV02) AddPreviousReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	a.PreviousReference = append(a.PreviousReference, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	a.RelatedReference = append(a.RelatedReference, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddMessagePagination() *model.Pagination {
	a.MessagePagination = new(model.Pagination)
	return a.MessagePagination
}

func (a *AccountingStatementOfHoldingsV02) AddStatementGeneralDetails() *model.Statement6 {
	a.StatementGeneralDetails = new(model.Statement6)
	return a.StatementGeneralDetails
}

func (a *AccountingStatementOfHoldingsV02) AddAccountDetails() *model.SafekeepingAccount2 {
	a.AccountDetails = new(model.SafekeepingAccount2)
	return a.AccountDetails
}

func (a *AccountingStatementOfHoldingsV02) AddBalanceForAccount() *model.AggregateBalanceInformation3 {
	newValue := new(model.AggregateBalanceInformation3)
	a.BalanceForAccount = append(a.BalanceForAccount, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddSubAccountDetails() *model.SubAccountIdentification3 {
	newValue := new(model.SubAccountIdentification3)
	a.SubAccountDetails = append(a.SubAccountDetails, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldingsV02) AddTotalValues() *model.TotalValueInPageAndStatement {
	a.TotalValues = new(model.TotalValueInPageAndStatement)
	return a.TotalValues
}

func (a *AccountingStatementOfHoldingsV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:semt.002.001.02 Document"`
	Message *CustodyStatementOfHoldingsV02 `xml:"CtdyStmtOfHldgsV02"`
}

func (d *Document00200102) AddMessage() *CustodyStatementOfHoldingsV02 {
	d.Message = new(CustodyStatementOfHoldingsV02)
	return d.Message
}

// Scope
// An account servicer, for example, a transfer agent sends the CustodyStatementOfHoldings message to the account owner, for example, a fund manager or an account owner's designated agent to provide detailed holdings of the portfolio at a specified moment in time.
// The message provides, at a moment in time, the quantity and identification of the financial instruments that the account servicer holds for the account owner. The message can also include availability and the location of holdings to facilitate trading and minimise settlement issues.
// The message reports all information per financial instrument, ie, when a financial instrument is held at multiple places of safekeeping, the total holdings for all locations can be provided.
// Usage
// The CustodyStatementOfHoldings message is used to provide detailed quantity and availability information for financial instrument holdings of a portfolio. The message should be sent at a frequency agreed bi-laterally between the account servicer and the account owner.
// This message can be also be used to report where the financial instruments are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// This message can reflect all outstanding holding information or may only contain changes since the previously sent statement.
// The CustodyStatementOfHoldings message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, this message can be used to either specify holdings at
// - the main account level, or,
// - the sub-account level.
type CustodyStatementOfHoldingsV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// General information related to the custody statement of holdings.
	StatementGeneralDetails *model.Statement7 `xml:"StmtGnlDtls"`

	// The safekeeping or investment account.
	AccountDetails *model.SafekeepingAccount2 `xml:"AcctDtls"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation4 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification5 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *model.TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (c *CustodyStatementOfHoldingsV02) AddMessageIdentification() *model.MessageIdentification1 {
	c.MessageIdentification = new(model.MessageIdentification1)
	return c.MessageIdentification
}

func (c *CustodyStatementOfHoldingsV02) AddPreviousReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	c.PreviousReference = append(c.PreviousReference, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	c.RelatedReference = append(c.RelatedReference, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddMessagePagination() *model.Pagination {
	c.MessagePagination = new(model.Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldingsV02) AddStatementGeneralDetails() *model.Statement7 {
	c.StatementGeneralDetails = new(model.Statement7)
	return c.StatementGeneralDetails
}

func (c *CustodyStatementOfHoldingsV02) AddAccountDetails() *model.SafekeepingAccount2 {
	c.AccountDetails = new(model.SafekeepingAccount2)
	return c.AccountDetails
}

func (c *CustodyStatementOfHoldingsV02) AddBalanceForAccount() *model.AggregateBalanceInformation4 {
	newValue := new(model.AggregateBalanceInformation4)
	c.BalanceForAccount = append(c.BalanceForAccount, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddSubAccountDetails() *model.SubAccountIdentification5 {
	newValue := new(model.SubAccountIdentification5)
	c.SubAccountDetails = append(c.SubAccountDetails, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldingsV02) AddTotalValues() *model.TotalValueInPageAndStatement {
	c.TotalValues = new(model.TotalValueInPageAndStatement)
	return c.TotalValues
}

func (c *CustodyStatementOfHoldingsV02) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	c.Extension = append(c.Extension, newValue)
	return newValue
}

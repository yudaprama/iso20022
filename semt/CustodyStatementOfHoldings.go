package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.002.001.01 Document"`
	Message *CustodyStatementOfHoldings `xml:"semt.002.001.01"`
}

func (d *Document00200101) AddMessage() *CustodyStatementOfHoldings {
	d.Message = new(CustodyStatementOfHoldings)
	return d.Message
}

// Scope
// The CustodyStatementOfHoldings message is sent by an account servicer to the account owner or the account owner's designated agent. The account servicer may be a local agent acting on behalf of its global custodian customer, a custodian acting on behalf of an investment management institution or a broker/dealer, a fund administrator or fund intermediary, trustee or registrar, etc.
// This message reports, at a specified moment in time, the quantity and identification of financial instruments that an account servicer holds for the account owner.
// This message is used to reconcile the books of the account owner and the account servicer for the specified account or sub-account.
// This message can also report availability and/or the location of security holdings to facilitate trading and minimise settlement issues. The reporting is per financial instrument, ie, when a financial instrument is held at multiple places of safekeeping, the total holding for all locations can be provided.
// Usage
// The CustodyStatementOfHoldings message can be sent:
// - At a frequency agreed bi-laterally between the Sender and the Receiver
// - As a response to a request for statement sent by the account owner.
// This message can reflect all outstanding holding information or may only contain changes since the previously sent statement.
// The CustodyStatementOfHoldings message can only be used to list the holdings of a single (master) account. However, it is possible to break down these holdings into one or several sub-accounts. Therefore, this message can be used to either specify holdings at
// - the main account level, or
// - the sub-account level.
// This message can be also be used to report where the securities are safe-kept, physically or notionally. If a security is held in more than one safekeeping place, this can also be indicated.
// This message must not be used to report audited positions. Audited positions are reported using the AccountingStatementOfHoldings message.
// Since a SWIFT message as sent is restricted to the maximum input message length, several messages may be needed to accommodate all the information.
type CustodyStatementOfHoldings struct {

	// Reference to a linked message that was previously sent.
	PreviousReference []*model.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*model.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *model.Pagination `xml:"MsgPgntn"`

	// General information related to the custody statement of holdings.
	StatementGeneralDetails *model.Statement3 `xml:"StmtGnlDtls"`

	// The safekeeping or investment account.
	AccountDetails *model.SafekeepingAccount1 `xml:"AcctDtls"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*model.AggregateBalanceInformation1 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*model.SubAccountIdentification1 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *model.TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (c *CustodyStatementOfHoldings) AddPreviousReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	c.PreviousReference = append(c.PreviousReference, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings) AddRelatedReference() *model.AdditionalReference2 {
	newValue := new(model.AdditionalReference2)
	c.RelatedReference = append(c.RelatedReference, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings) AddMessagePagination() *model.Pagination {
	c.MessagePagination = new(model.Pagination)
	return c.MessagePagination
}

func (c *CustodyStatementOfHoldings) AddStatementGeneralDetails() *model.Statement3 {
	c.StatementGeneralDetails = new(model.Statement3)
	return c.StatementGeneralDetails
}

func (c *CustodyStatementOfHoldings) AddAccountDetails() *model.SafekeepingAccount1 {
	c.AccountDetails = new(model.SafekeepingAccount1)
	return c.AccountDetails
}

func (c *CustodyStatementOfHoldings) AddBalanceForAccount() *model.AggregateBalanceInformation1 {
	newValue := new(model.AggregateBalanceInformation1)
	c.BalanceForAccount = append(c.BalanceForAccount, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings) AddSubAccountDetails() *model.SubAccountIdentification1 {
	newValue := new(model.SubAccountIdentification1)
	c.SubAccountDetails = append(c.SubAccountDetails, newValue)
	return newValue
}

func (c *CustodyStatementOfHoldings) AddTotalValues() *model.TotalValueInPageAndStatement {
	c.TotalValues = new(model.TotalValueInPageAndStatement)
	return c.TotalValues
}

func (c *CustodyStatementOfHoldings) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	c.Extension = append(c.Extension, newValue)
	return newValue
}

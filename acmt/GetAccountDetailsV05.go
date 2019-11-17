package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400105 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.004.001.05 Document"`
	Message *GetAccountDetailsV05 `xml:"GetAcctDtls"`
}

func (d *Document00400105) AddMessage() *GetAccountDetailsV05 {
	d.Message = new(GetAccountDetailsV05)
	return d.Message
}

// Scope
// The GetAccountDetails message is sent by an account owner, for example, an investor or its designated agent to the account servicer, for example, a registrar, transfer agent, custodian bank or securities depository to query the details of an existing account.
// Usage
// The GetAccountDetails message is used to query all or some of the account details for a given account held with an account servicer.
// The message is used prior to an AccountModificationInstruction in order to validate account information before requesting a modification.
// The GetAccountDetails message may also be used to collect account information for general account reconciliation purposes.
// The response to a GetAccountDetails message is via an AccountDetailsConfirmation message.
type GetAccountDetailsV05 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identifies the account for which query is sent.
	InvestmentAccountSelection *model.AccountSelection2Choice `xml:"InvstmtAcctSelctn"`

	// Identifies the type of information that is to be returned in via an  AccountDetailsConfirmation message.
	SelectedInformationType *model.InvestmentAccountInformationType1 `xml:"SelctdInfTp"`
}

func (g *GetAccountDetailsV05) AddMessageIdentification() *model.MessageIdentification1 {
	g.MessageIdentification = new(model.MessageIdentification1)
	return g.MessageIdentification
}

func (g *GetAccountDetailsV05) AddInvestmentAccountSelection() *model.AccountSelection2Choice {
	g.InvestmentAccountSelection = new(model.AccountSelection2Choice)
	return g.InvestmentAccountSelection
}

func (g *GetAccountDetailsV05) AddSelectedInformationType() *model.InvestmentAccountInformationType1 {
	g.SelectedInformationType = new(model.InvestmentAccountInformationType1)
	return g.SelectedInformationType
}

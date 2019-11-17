package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05300101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.053.001.01 Document"`
	Message *BankToCustomerStatementV01 `xml:"BkToCstmrStmtV01"`
}

func (d *Document05300101) AddMessage() *BankToCustomerStatementV01 {
	d.Message = new(BankToCustomerStatementV01)
	return d.Message
}

// Scope
// The Bank-to-Customer Statement message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It is used to inform the account owner, or authorised party, of the entries booked to the account, and to provide the owner with balance information on the account at a given point in time.
// Usage
// The Bank-to-Customer Statement message can contain reports for more than 1 account. It provides information for cash management and/or reconciliation.
// It contains information on booked entries only.
// It can include underlying details of transactions that have been included in the entry.
// The message is exchanged as defined between the account servicer and the account owner. It provides information on items that have been booked to the account (and therefore are "binding" and also balance information. Depending on services agreed between banks and their customers, "binding" statements can be generated and exchanged intraday. Depending on legal requirements in local jurisdictions, "end-of-day" statements may need to be mandatorily generated and exchanged.
// It is possible that the receiver of the message is not the account owner, but a party entitled through arrangement with the account owner to receive the account information (also known as recipient).
type BankToCustomerStatementV01 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader23 `xml:"GrpHdr"`

	// Reports on booked entries and balances for a cash account.
	Statement []*model.AccountStatement1 `xml:"Stmt"`
}

func (b *BankToCustomerStatementV01) AddGroupHeader() *model.GroupHeader23 {
	b.GroupHeader = new(model.GroupHeader23)
	return b.GroupHeader
}

func (b *BankToCustomerStatementV01) AddStatement() *model.AccountStatement1 {
	newValue := new(model.AccountStatement1)
	b.Statement = append(b.Statement, newValue)
	return newValue
}

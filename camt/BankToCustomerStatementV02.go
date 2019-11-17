package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05300102 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.053.001.02 Document"`
	Message *BankToCustomerStatementV02 `xml:"BkToCstmrStmt"`
}

func (d *Document05300102) AddMessage() *BankToCustomerStatementV02 {
	d.Message = new(BankToCustomerStatementV02)
	return d.Message
}

// Scope
// The BankToCustomerStatement message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It is used to inform the account owner, or authorised party, of the entries booked to the account, and to provide the owner with balance information on the account at a given point in time.
// Usage
// The BankToCustomerStatement message can contain reports for more than one account. It provides information for cash management and/or reconciliation.
// It contains information on booked entries only.
// It can include underlying details of transactions that have been included in the entry.
// The message is exchanged as defined between the account servicer and the account owner. It provides information on items that have been booked to the account and also balance information. Depending on services and schedule agreed between banks and their customers, statements may be generated and exchanged accordingly, for example for intraday or prior day periods.
// It is possible that the receiver of the message is not the account owner, but a party entitled through arrangement with the account owner to receive the account information (also known as recipient).
type BankToCustomerStatementV02 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader42 `xml:"GrpHdr"`

	// Reports on booked entries and balances for a cash account.
	Statement []*model.AccountStatement2 `xml:"Stmt"`
}

func (b *BankToCustomerStatementV02) AddGroupHeader() *model.GroupHeader42 {
	b.GroupHeader = new(model.GroupHeader42)
	return b.GroupHeader
}

func (b *BankToCustomerStatementV02) AddStatement() *model.AccountStatement2 {
	newValue := new(model.AccountStatement2)
	b.Statement = append(b.Statement, newValue)
	return newValue
}

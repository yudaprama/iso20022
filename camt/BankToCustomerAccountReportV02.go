package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05200102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.02 Document"`
	Message *BankToCustomerAccountReportV02 `xml:"BkToCstmrAcctRpt"`
}

func (d *Document05200102) AddMessage() *BankToCustomerAccountReportV02 {
	d.Message = new(BankToCustomerAccountReportV02)
	return d.Message
}

// Scope
// The BankToCustomerAccountReport message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It can be used to inform the account owner, or authorised party, of the entries reported to the account, and/or to provide the owner with balance information on the account at a given point in time.
// Usage
// The BankToCustomerAccountReport message can contain reports for more than one account. It provides information for cash management and/or reconciliation. It can be used to:
// - report pending and booked items;
// - provide balance information.
// It can include underlying details of transactions that have been included in the entry.
// It is possible that the receiver of the message is not the account owner, but a party entitled by the account owner to receive the account information (also known as recipient).
// For a statement, the Bank-to-Customer Account Statement message should be used.
type BankToCustomerAccountReportV02 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader42 `xml:"GrpHdr"`

	// Reports on a cash account.
	Report []*model.AccountReport11 `xml:"Rpt"`
}

func (b *BankToCustomerAccountReportV02) AddGroupHeader() *model.GroupHeader42 {
	b.GroupHeader = new(model.GroupHeader42)
	return b.GroupHeader
}

func (b *BankToCustomerAccountReportV02) AddReport() *model.AccountReport11 {
	newValue := new(model.AccountReport11)
	b.Report = append(b.Report, newValue)
	return newValue
}

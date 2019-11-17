package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05200101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.01 Document"`
	Message *BankToCustomerAccountReportV01 `xml:"BkToCstmrAcctRptV01"`
}

func (d *Document05200101) AddMessage() *BankToCustomerAccountReportV01 {
	d.Message = new(BankToCustomerAccountReportV01)
	return d.Message
}

// Scope
// The Bank-to-Customer Account Report message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It can be used to inform the account owner, or authorised party, of the entries reported to the account, and/or to provide the owner with balance information on the account at a given point in time.
// Usage
// The Bank-to-Customer Account Report message can contain reports for more than 1 account. It provides information for cash management and/or reconciliation. It can be used to :
// - report pending and booked items;
// - provide balance information
// It can include underlying details of transactions that have been included in the entry.
// It is possible that the receiver of the message is not the account owner, but a party entitled by the account owner to receive the account information (also known as recipient).
// For a statement that is required due to local legal stipulations, the Bank-to-Customer Account Statement message should be used.
type BankToCustomerAccountReportV01 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader23 `xml:"GrpHdr"`

	// Reports on a cash account.
	Report []*model.AccountReport9 `xml:"Rpt"`
}

func (b *BankToCustomerAccountReportV01) AddGroupHeader() *model.GroupHeader23 {
	b.GroupHeader = new(model.GroupHeader23)
	return b.GroupHeader
}

func (b *BankToCustomerAccountReportV01) AddReport() *model.AccountReport9 {
	newValue := new(model.AccountReport9)
	b.Report = append(b.Report, newValue)
	return newValue
}

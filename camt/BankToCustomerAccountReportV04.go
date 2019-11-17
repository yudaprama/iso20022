package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05200104 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.052.001.04 Document"`
	Message *BankToCustomerAccountReportV04 `xml:"BkToCstmrAcctRpt"`
}

func (d *Document05200104) AddMessage() *BankToCustomerAccountReportV04 {
	d.Message = new(BankToCustomerAccountReportV04)
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
type BankToCustomerAccountReportV04 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader58 `xml:"GrpHdr"`

	// Reports on a cash account.
	Report []*model.AccountReport16 `xml:"Rpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (b *BankToCustomerAccountReportV04) AddGroupHeader() *model.GroupHeader58 {
	b.GroupHeader = new(model.GroupHeader58)
	return b.GroupHeader
}

func (b *BankToCustomerAccountReportV04) AddReport() *model.AccountReport16 {
	newValue := new(model.AccountReport16)
	b.Report = append(b.Report, newValue)
	return newValue
}

func (b *BankToCustomerAccountReportV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}

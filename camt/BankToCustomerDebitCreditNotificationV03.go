package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05400103 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.054.001.03 Document"`
	Message *BankToCustomerDebitCreditNotificationV03 `xml:"BkToCstmrDbtCdtNtfctn"`
}

func (d *Document05400103) AddMessage() *BankToCustomerDebitCreditNotificationV03 {
	d.Message = new(BankToCustomerDebitCreditNotificationV03)
	return d.Message
}

// Scope
// The BankToCustomerDebitCreditNotification message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It can be used to inform the account owner, or authorised party, of single or multiple debit and/or credit entries reported to the account.
// Usage
// The BankToCustomerDebitCreditNotification message can contain reports for more than one account. It provides information for cash management and/or reconciliation.
// It can be used to :
// - report pending and booked items;
// - notify one or more debit entries;
// - notify one or more credit entries;
// - notify a combination of debit and credit entries.
// It can include underlying details of transactions that have been included in the entry.
// It is possible that the receiver of the message is not the account owner, but a party entitled by the account owner to receive the account information (also known as recipient).
// It does not contain balance information.
type BankToCustomerDebitCreditNotificationV03 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader58 `xml:"GrpHdr"`

	// Notifies debit and credit entries for the account.
	Notification []*model.AccountNotification5 `xml:"Ntfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (b *BankToCustomerDebitCreditNotificationV03) AddGroupHeader() *model.GroupHeader58 {
	b.GroupHeader = new(model.GroupHeader58)
	return b.GroupHeader
}

func (b *BankToCustomerDebitCreditNotificationV03) AddNotification() *model.AccountNotification5 {
	newValue := new(model.AccountNotification5)
	b.Notification = append(b.Notification, newValue)
	return newValue
}

func (b *BankToCustomerDebitCreditNotificationV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}

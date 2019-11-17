package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05400102 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.054.001.02 Document"`
	Message *BankToCustomerDebitCreditNotificationV02 `xml:"BkToCstmrDbtCdtNtfctn"`
}

func (d *Document05400102) AddMessage() *BankToCustomerDebitCreditNotificationV02 {
	d.Message = new(BankToCustomerDebitCreditNotificationV02)
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
type BankToCustomerDebitCreditNotificationV02 struct {

	// Common information for the message.
	GroupHeader *model.GroupHeader42 `xml:"GrpHdr"`

	// Notifies debit and credit entries for the account.
	Notification []*model.AccountNotification2 `xml:"Ntfctn"`
}

func (b *BankToCustomerDebitCreditNotificationV02) AddGroupHeader() *model.GroupHeader42 {
	b.GroupHeader = new(model.GroupHeader42)
	return b.GroupHeader
}

func (b *BankToCustomerDebitCreditNotificationV02) AddNotification() *model.AccountNotification2 {
	newValue := new(model.AccountNotification2)
	b.Notification = append(b.Notification, newValue)
	return newValue
}

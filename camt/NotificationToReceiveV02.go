package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05700102 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.057.001.02 Document"`
	Message *NotificationToReceiveV02 `xml:"NtfctnToRcv"`
}

func (d *Document05700102) AddMessage() *NotificationToReceiveV02 {
	d.Message = new(NotificationToReceiveV02)
	return d.Message
}

// Scope
// The NotificationToReceive message is sent by an account owner or by a party acting on the account owner's behalf to one of the account owner's account servicing institutions. It is an advance notice that the account servicing institution will receive funds to be credited to the account of the account owner.
// Usage
// The NotificationToReceive message is used to advise the account servicing institution of funds that the account owner expects to have credited to its account. The message can be used in either a direct or a relay scenario.
type NotificationToReceiveV02 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *model.GroupHeader43 `xml:"GrpHdr"`

	// Set of elements used to provide further details on the account notification.
	Notification *model.AccountNotification4 `xml:"Ntfctn"`
}

func (n *NotificationToReceiveV02) AddGroupHeader() *model.GroupHeader43 {
	n.GroupHeader = new(model.GroupHeader43)
	return n.GroupHeader
}

func (n *NotificationToReceiveV02) AddNotification() *model.AccountNotification4 {
	n.Notification = new(model.AccountNotification4)
	return n.Notification
}

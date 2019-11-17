package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05900102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.02 Document"`
	Message *NotificationToReceiveStatusReportV02 `xml:"NtfctnToRcvStsRpt"`
}

func (d *Document05900102) AddMessage() *NotificationToReceiveStatusReportV02 {
	d.Message = new(NotificationToReceiveStatusReportV02)
	return d.Message
}

// Scope
// The NotificationToReceiveStatusReport message is sent by an account servicing institution to an account owner or to a party acting on the account owner's behalf. It is used to notify the account owner about the status of one or more expected payments that were advised in a previous NotificationToReceive message.
// Usage
// The NotificationToReceiveStatusReport message is sent in response to a NotificationToReceive message and can be used in either a direct or a relay scenario. It is used to advise the account owner of receipt of the funds as expected. It is also used to notify the account owner of non-receipt of funds or of discrepancies in the funds received.
type NotificationToReceiveStatusReportV02 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *model.GroupHeader44 `xml:"GrpHdr"`

	// Set of elements used to identify the original notification and to provide the status.
	OriginalNotificationAndStatus *model.OriginalNotification3 `xml:"OrgnlNtfctnAndSts"`
}

func (n *NotificationToReceiveStatusReportV02) AddGroupHeader() *model.GroupHeader44 {
	n.GroupHeader = new(model.GroupHeader44)
	return n.GroupHeader
}

func (n *NotificationToReceiveStatusReportV02) AddOriginalNotificationAndStatus() *model.OriginalNotification3 {
	n.OriginalNotificationAndStatus = new(model.OriginalNotification3)
	return n.OriginalNotificationAndStatus
}

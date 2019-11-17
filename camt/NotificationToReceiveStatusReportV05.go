package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05900105 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.059.001.05 Document"`
	Message *NotificationToReceiveStatusReportV05 `xml:"NtfctnToRcvStsRpt"`
}

func (d *Document05900105) AddMessage() *NotificationToReceiveStatusReportV05 {
	d.Message = new(NotificationToReceiveStatusReportV05)
	return d.Message
}

// Scope
// The NotificationToReceiveStatusReport message is sent by an account servicing institution to an account owner or to a party acting on the account owner's behalf. It is used to notify the account owner about the status of one or more expected payments that were advised in a previous NotificationToReceive message.
// Usage
// The NotificationToReceiveStatusReport message is sent in response to a NotificationToReceive message and can be used in either a direct or a relay scenario. It is used to advise the account owner of receipt of the funds as expected. It is also used to notify the account owner of non-receipt of funds or of discrepancies in the funds received.
type NotificationToReceiveStatusReportV05 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *model.GroupHeader60 `xml:"GrpHdr"`

	// Set of elements used to identify the original notification and to provide the status.
	OriginalNotificationAndStatus *model.OriginalNotification9 `xml:"OrgnlNtfctnAndSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NotificationToReceiveStatusReportV05) AddGroupHeader() *model.GroupHeader60 {
	n.GroupHeader = new(model.GroupHeader60)
	return n.GroupHeader
}

func (n *NotificationToReceiveStatusReportV05) AddOriginalNotificationAndStatus() *model.OriginalNotification9 {
	n.OriginalNotificationAndStatus = new(model.OriginalNotification9)
	return n.OriginalNotificationAndStatus
}

func (n *NotificationToReceiveStatusReportV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}

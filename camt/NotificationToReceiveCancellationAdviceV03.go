package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document05800103 struct {
	XMLName xml.Name                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.058.001.03 Document"`
	Message *NotificationToReceiveCancellationAdviceV03 `xml:"NtfctnToRcvCxlAdvc"`
}

func (d *Document05800103) AddMessage() *NotificationToReceiveCancellationAdviceV03 {
	d.Message = new(NotificationToReceiveCancellationAdviceV03)
	return d.Message
}

// Scope
// The NotificationToReceiveCancellationAdvice message is sent by an account owner or by a party acting on the account owner's behalf to one of the account owner's account servicing institutions. It is used to advise the account servicing institution about the cancellation of one or more notifications in a previous NotificationToReceive message.
// Usage
// The NotificationToReceiveCancellationAdvice message is used to advise the account servicing institution that the funds are no longer expected. The message can be used in either a direct or a relay scenario.
type NotificationToReceiveCancellationAdviceV03 struct {

	// Set of elements used to provide further details on the message.
	GroupHeader *model.GroupHeader59 `xml:"GrpHdr"`

	// Set of elements used to identify the original notification, to which the cancellation advice refers.
	OriginalNotification *model.OriginalNotification6 `xml:"OrgnlNtfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NotificationToReceiveCancellationAdviceV03) AddGroupHeader() *model.GroupHeader59 {
	n.GroupHeader = new(model.GroupHeader59)
	return n.GroupHeader
}

func (n *NotificationToReceiveCancellationAdviceV03) AddOriginalNotification() *model.OriginalNotification6 {
	n.OriginalNotification = new(model.OriginalNotification6)
	return n.OriginalNotification
}

func (n *NotificationToReceiveCancellationAdviceV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}

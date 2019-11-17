package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document03000104 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.04 Document"`
	Message *NotificationOfCaseAssignmentV04 `xml:"NtfctnOfCaseAssgnmt"`
}

func (d *Document03000104) AddMessage() *NotificationOfCaseAssignmentV04 {
	d.Message = new(NotificationOfCaseAssignmentV04)
	return d.Message
}

// Scope
// The Notification Of Case Assignment message is sent by a case assignee to a case creator/case assigner.
// This message is used to inform the case assigner that:
// - the assignee is reassigning the case to the next agent in the transaction processing chain for further action
// - the assignee will work on the case himself, without re-assigning it to another party, and therefore indicating that the re-assignment has reached its end-point
// Usage
// The Notification Of Case Assignment message is used to notify the case creator or case assigner of further action undertaken by the case assignee in a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// The Notification Of Case Assignment message
// - covers one and only one case at a time. If the case assignee needs to inform a case creator or case assigner about several cases, then multiple Notification Of Case Assignment messages must be sent
// - except in the case where it is used to indicate that an agent is doing the correction himself, this message must be forwarded by all subsequent case assigner(s) until it reaches the case creator
// - must not be used in place of a Resolution Of Investigation or a Case Status Report message
// When the assignee does not reassign the case to another party (that is responding with a Notification Of Case Assignment message with notification MINE - The case is processed by the assignee), the case assignment should contain the case assignment elements as received in the original query.
type NotificationOfCaseAssignmentV04 struct {

	// Specifies generic information about the notification.
	// The receiver of a notification must be the party which assigned the case to the sender.
	Header *model.ReportHeader4 `xml:"Hdr"`

	// Identifies the investigation case.
	Case *model.Case3 `xml:"Case"`

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *model.CaseAssignment3 `xml:"Assgnmt"`

	// Information about the type of action taken.
	Notification *model.CaseForwardingNotification3 `xml:"Ntfctn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NotificationOfCaseAssignmentV04) AddHeader() *model.ReportHeader4 {
	n.Header = new(model.ReportHeader4)
	return n.Header
}

func (n *NotificationOfCaseAssignmentV04) AddCase() *model.Case3 {
	n.Case = new(model.Case3)
	return n.Case
}

func (n *NotificationOfCaseAssignmentV04) AddAssignment() *model.CaseAssignment3 {
	n.Assignment = new(model.CaseAssignment3)
	return n.Assignment
}

func (n *NotificationOfCaseAssignmentV04) AddNotification() *model.CaseForwardingNotification3 {
	n.Notification = new(model.CaseForwardingNotification3)
	return n.Notification
}

func (n *NotificationOfCaseAssignmentV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}

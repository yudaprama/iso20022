package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.012.001.03 Document"`
	Message *MandateAcceptanceReportV03 `xml:"MndtAccptncRpt"`
}

func (d *Document01200103) AddMessage() *MandateAcceptanceReportV03 {
	d.Message = new(MandateAcceptanceReportV03)
	return d.Message
}

// Scope
// The MandateAcceptanceReport message is sent from the agent of the receiver (debtor or creditor) of the MandateRequest message (initiation, amendment or cancellation) to the agent of the initiator of the MandateRequest message (debtor or creditor).
// A MandateAcceptanceReport message is used to confirm the acceptance or rejection of a MandateRequest message. Where acceptance is part of the full process flow, a MandateRequest message only becomes valid after a confirmation of acceptance is received through a MandateAcceptanceReport message from the agent of the receiver.
// Usage
// The MandateAcceptanceReport message can contain one or more confirmation(s) of acceptance or rejection of a specific Mandate Request.
// The messages can be exchanged between debtor agent and creditor agent and between debtor agent and debtor and creditor agent and creditor.
// The MandateAcceptanceReport message can be used in domestic and cross-border scenarios.
type MandateAcceptanceReportV03 struct {

	// Set of characteristics to identify the message and parties playing a role in the mandate acceptance, but which are not part of the mandate.
	GroupHeader *model.GroupHeader47 `xml:"GrpHdr"`

	// Set of elements used to provide information on the acception or rejection of the mandate request.
	UnderlyingAcceptanceDetails []*model.MandateAcceptance3 `xml:"UndrlygAccptncDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateAcceptanceReportV03) AddGroupHeader() *model.GroupHeader47 {
	m.GroupHeader = new(model.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateAcceptanceReportV03) AddUnderlyingAcceptanceDetails() *model.MandateAcceptance3 {
	newValue := new(model.MandateAcceptance3)
	m.UnderlyingAcceptanceDetails = append(m.UnderlyingAcceptanceDetails, newValue)
	return newValue
}

func (m *MandateAcceptanceReportV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

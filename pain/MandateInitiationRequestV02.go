package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00900102 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:pain.009.001.02 Document"`
	Message *MandateInitiationRequestV02 `xml:"MndtInitnReq"`
}

func (d *Document00900102) AddMessage() *MandateInitiationRequestV02 {
	d.Message = new(MandateInitiationRequestV02)
	return d.Message
}

// Scope
// The MandateInitiationRequest message is sent by the initiator of the request to his agent. The initiator can either be the debtor or the creditor.
// The MandateInitiationRequest message is forwarded by the agent of the initiator to the agent of the counterparty.
// The MandateInitiationRequest message is used to set-up the instruction that allows the debtor agent to accept instructions from the creditor, through the creditor agent, to debit the account of the debtor.
// Usage
// The MandateInitiationRequest message can contain only one request to set-up one specific mandate.
// The messages can be exchanged between creditor and creditor agent or debtor and debtor agent and between creditor agent and debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor or debtor.
// The MandateInitiationRequest message can be used in domestic and cross-border scenarios.
type MandateInitiationRequestV02 struct {

	// Set of characteristics to identify the message and parties playing a role in the mandate initiation, but which are not part of the mandate.
	GroupHeader *model.GroupHeader47 `xml:"GrpHdr"`

	// Set of elements used to provide the details of the mandate signed between the (ultimate) creditor and the (ultimate) debtor.
	Mandate *model.Mandate2 `xml:"Mndt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateInitiationRequestV02) AddGroupHeader() *model.GroupHeader47 {
	m.GroupHeader = new(model.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateInitiationRequestV02) AddMandate() *model.Mandate2 {
	m.Mandate = new(model.Mandate2)
	return m.Mandate
}

func (m *MandateInitiationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

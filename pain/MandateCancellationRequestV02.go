package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100102 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.02 Document"`
	Message *MandateCancellationRequestV02 `xml:"MndtCxlReq"`
}

func (d *Document01100102) AddMessage() *MandateCancellationRequestV02 {
	d.Message = new(MandateCancellationRequestV02)
	return d.Message
}

// Scope
// The MandateCancellationRequest message is sent by the initiator of the request to his agent. The initiator can either be the debtor or the creditor.
// The MandateCancellationRequest message is forwarded by the agent of the initiator to the agent of the counterparty.
// A MandateCancellationRequest message is used to request the cancellation of an existing mandate. If accepted, this MandateCancellationRequest message together with the MandateAcceptanceReport message confirming the acceptance will be considered a valid cancellation of an existing mandate, agreed upon by all parties.
// Usage
// The MandateCancellationRequest message can contain only one request to cancel one specific mandate.
// The messages can be exchanged between creditor and creditor agent or debtor and debtor agent and between creditor agent and debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor or debtor.
// The MandateCancellationRequest message can be used in domestic and cross-border scenarios.
type MandateCancellationRequestV02 struct {

	// Set of characteristics to identify the message and parties playing a role in the cancellation of the mandate, but which are not part of the mandate.
	GroupHeader *model.GroupHeader47 `xml:"GrpHdr"`

	// Set of elements used to provide details on the cancellation request.
	UnderlyingCancellationDetails *model.MandateCancellation2 `xml:"UndrlygCxlDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCancellationRequestV02) AddGroupHeader() *model.GroupHeader47 {
	m.GroupHeader = new(model.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateCancellationRequestV02) AddUnderlyingCancellationDetails() *model.MandateCancellation2 {
	m.UnderlyingCancellationDetails = new(model.MandateCancellation2)
	return m.UnderlyingCancellationDetails
}

func (m *MandateCancellationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

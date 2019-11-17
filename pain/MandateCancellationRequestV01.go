package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01100101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:pain.011.001.01 Document"`
	Message *MandateCancellationRequestV01 `xml:"MndtCxlReq"`
}

func (d *Document01100101) AddMessage() *MandateCancellationRequestV01 {
	d.Message = new(MandateCancellationRequestV01)
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
type MandateCancellationRequestV01 struct {

	// Set of characteristics to identify the message and parties playing a role in the cancellation of the mandate, but which are not part of the mandate.
	GroupHeader *model.GroupHeader31 `xml:"GrpHdr"`

	// Set of elements used to provide details on the cancellation request.
	UnderlyingCancellationDetails *model.MandateCancellation1 `xml:"UndrlygCxlDtls"`
}

func (m *MandateCancellationRequestV01) AddGroupHeader() *model.GroupHeader31 {
	m.GroupHeader = new(model.GroupHeader31)
	return m.GroupHeader
}

func (m *MandateCancellationRequestV01) AddUnderlyingCancellationDetails() *model.MandateCancellation1 {
	m.UnderlyingCancellationDetails = new(model.MandateCancellation1)
	return m.UnderlyingCancellationDetails
}

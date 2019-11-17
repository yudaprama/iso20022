package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name               `xml:"urn:iso:std:iso:20022:tech:xsd:pain.017.001.01 Document"`
	Message *MandateCopyRequestV01 `xml:"MndtCpyReq"`
}

func (d *Document01700101) AddMessage() *MandateCopyRequestV01 {
	d.Message = new(MandateCopyRequestV01)
	return d.Message
}

// Scope
// The MandateCopyRequest message is sent by the initiator of the request to his agent. The initiator can either be the debtor or the creditor.
// The MandateCopyRequest message is forwarded by the agent of the initiator to the agent of the counterparty.
// A MandateCopyRequest message is used to request a copy of an existing mandate. If accepted, the mandate copy can be sent using the MandateAcceptanceReport message.
// Usage
// The MandateCopyRequest message can contain one or more copy requests.
// The messages can be exchanged between creditor and creditor agent or debtor and debtor agent and between creditor agent and debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor or debtor.
// The MandateCopyRequest message can be used in domestic and cross-border scenarios.
//
type MandateCopyRequestV01 struct {

	// Set of characteristics to identify the message and parties playing a role in the mandate copy request, but which are not part of the mandate.
	GroupHeader *model.GroupHeader47 `xml:"GrpHdr"`

	// Set of information used to identify the mandate for which a copy is requested.
	UnderlyingCopyRequestDetails []*model.MandateCopy1 `xml:"UndrlygCpyReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MandateCopyRequestV01) AddGroupHeader() *model.GroupHeader47 {
	m.GroupHeader = new(model.GroupHeader47)
	return m.GroupHeader
}

func (m *MandateCopyRequestV01) AddUnderlyingCopyRequestDetails() *model.MandateCopy1 {
	newValue := new(model.MandateCopy1)
	m.UnderlyingCopyRequestDetails = append(m.UnderlyingCopyRequestDetails, newValue)
	return newValue
}

func (m *MandateCopyRequestV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

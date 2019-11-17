package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.003.001.01 Document"`
	Message *InvoiceFinancingCancellationRequestV01 `xml:"InvcFincgCxlReq"`
}

func (d *Document00300101) AddMessage() *InvoiceFinancingCancellationRequestV01 {
	d.Message = new(InvoiceFinancingCancellationRequestV01)
	return d.Message
}

// Scope
// The InvoiceFinancingCancellationRequest message is sent by the Financing Requestor to the Intermediary Agent (relay scenario) or First Agent (direct scenario). It is used to request the cancellation of a previously sent financing request.
// Usage
// The InvoiceFinancingCancellationRequest message is used by the Financing Requestor to request the cancellation of a previously sent financing request.
// It is not possible to send a cancellation request for a single invoice contained in a bulk invoice financing request.
// The InvoiceFinancingCancellationRequest message contains references (original group identification and original creation date and time) of the original financing request message to which is referred.
// As for InvoiceFinancingRequest, the message can be used in a direct or a relay scenario:
// - In a direct scenario, the message is sent directly to the First Agent. The First Agent is the account servicer of the Financing Requestor.
// - In a relay scenario, the message is sent to an Intermediary Agent. The Intermediary Agent forwards the InvoiceFinancingCancellingRequest message to the First Agent.
type InvoiceFinancingCancellationRequestV01 struct {

	// Unique and unambiguous identification of the message.
	CancellationRequestIdentification *model.MessageIdentification1 `xml:"CxlReqId"`

	// Set of information related to the cancellation request, such as actors involved and identification of the original invoice financing request to which the cancellation request refers.
	CancellationRequestInformation *model.CancellationRequestInformation1 `xml:"CxlReqInf"`
}

func (i *InvoiceFinancingCancellationRequestV01) AddCancellationRequestIdentification() *model.MessageIdentification1 {
	i.CancellationRequestIdentification = new(model.MessageIdentification1)
	return i.CancellationRequestIdentification
}

func (i *InvoiceFinancingCancellationRequestV01) AddCancellationRequestInformation() *model.CancellationRequestInformation1 {
	i.CancellationRequestInformation = new(model.CancellationRequestInformation1)
	return i.CancellationRequestInformation
}

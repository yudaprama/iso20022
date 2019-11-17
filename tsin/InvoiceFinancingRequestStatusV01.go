package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00200101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.002.001.01 Document"`
	Message *InvoiceFinancingRequestStatusV01 `xml:"InvcFincgReqSts"`
}

func (d *Document00200101) AddMessage() *InvoiceFinancingRequestStatusV01 {
	d.Message = new(InvoiceFinancingRequestStatusV01)
	return d.Message
}

// Scope
// The InvoiceFinancingRequestStatus message is sent by the First Agent to the Financing Requestor, alternatively through an Intermediary Agent (relay scenario). It is used to inform the Financing Requestor about the positive or negative status of a financing request or a financing cancellation request.
// Usage
// The InvoiceFinancingRequestStatus message flows from the First Agent to the Financing Requestor (alternatively through an Intermediary Agent) to provide status information about a request previously sent.
// Its usage will always be governed by a bilateral agreement between the First Agent and the Financing Requestor.
// The InvoiceFinancingRequestStatus message can be used two fold:
// - to provide information about the reception status (eg rejection, acceptance) of a request message. In this case the status message is the result of a technical validation performed by the First Agent on the request message received;
// - to inform the Financing Requestor about the business status of the financing process initiated. In this case the First Agent can:
// * communicate that a single financing request has been granted, is pending or has not been granted at all;
// * inform that a financing cancellation request has been allowed or denied.
// Note.
// If the Financing Requestor requests financing for more than one instalment related to the same invoice, the First Agent can decide to finance only some of the instalments. In such case the status message contains details and status of every single instalment (financed, not financed).
// The message can be used in a direct or in a relay scenario:
// - In a direct scenario, the message is sent directly by the First Agent to the Financing Requestor;
// - In a relay scenario, the message is sent first by the First Agent to the Intermediary Agent, who forwards it to the Financing Requestor.
// The InvoiceFinancingRequestStatus message refers to the original request(s) by means of references and a set of data elements included into the original request.
type InvoiceFinancingRequestStatusV01 struct {

	// General information that unambiguously identify the invoice financing status report, such as status identification, creation date time.
	StatusIdentification *model.MessageIdentification1 `xml:"StsId"`

	// Set of summary information that unambiguously identifies the original invoice financing (or cancellation) request to which the status is referred. The status of the original request is also given in this block.
	//
	OriginalRequestInformationAndStatus *model.OriginalRequestInformation1 `xml:"OrgnlReqInfAndSts"`

	// Information concerning the business status of a financing request.
	FinancingInformationAndStatus *model.FinancingInformationAndStatus1 `xml:"FincgInfAndSts,omitempty"`
}

func (i *InvoiceFinancingRequestStatusV01) AddStatusIdentification() *model.MessageIdentification1 {
	i.StatusIdentification = new(model.MessageIdentification1)
	return i.StatusIdentification
}

func (i *InvoiceFinancingRequestStatusV01) AddOriginalRequestInformationAndStatus() *model.OriginalRequestInformation1 {
	i.OriginalRequestInformationAndStatus = new(model.OriginalRequestInformation1)
	return i.OriginalRequestInformationAndStatus
}

func (i *InvoiceFinancingRequestStatusV01) AddFinancingInformationAndStatus() *model.FinancingInformationAndStatus1 {
	i.FinancingInformationAndStatus = new(model.FinancingInformationAndStatus1)
	return i.FinancingInformationAndStatus
}

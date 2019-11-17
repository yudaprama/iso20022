package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.001.001.01 Document"`
	Message *InvoiceFinancingRequestV01 `xml:"InvcFincgReq"`
}

func (d *Document00100101) AddMessage() *InvoiceFinancingRequestV01 {
	d.Message = new(InvoiceFinancingRequestV01)
	return d.Message
}

// Scope
// The InvoiceFinancingRequest message is sent by Financing Requestor to the Intermediary Agent or First agent. It is used to request financing of a set of invoices, referenced in the request message itself. If the whole financing request (or a selection of single invoice requests included) is accepted, the amount financed by the First Agent will be credited either to the account specified in the financing request or to another account held by Financing Requestor to First Agent.
// Usage
// The InvoiceFinancingRequest message is issued by the Financing Requestor and represents a bulk financing request since it can contain one or more single financing requests, each request related to an invoice.
// For every invoice it is always possible to identify a supplier and a buyer.
// The subject playing the role of supplier can be different from the Financing Requestor; in this case the Financing Requestor is allowed to send the request message on behalf of the supplier itself.
// This caters for example in the scenario of a collection agency initiating all requests on behalf of a large corporate.
// In instances where an invoice is going to be paid by means of instalments, the Financing Requestor can request financing for one or more instalments related to the invoice payment. In this case, together with the general information related to the invoice, references about instalments to be financed are specified into the request message. The request message must contain information only about the instalments that the Financing Requestor wants to be financed.
// The InvoiceFinancingRequest message is used to exchange:
// - One instance of general information related to the invoice financing request;
// - One instance of information for each single invoice financing request;
// - Optionally, one instance of information for each single instalment to be financed.
// The message can be used in a direct or a relay scenario:
// - In a direct scenario, the message is sent directly to the First Agent. The First Agent is the account servicer of the Financing Requestor;
// - In a relay scenario, the message is sent to an Intermediary Agent. The Intermediary Agent acts as an access point that will forward the InvoiceFinancingRequest message to the First Agent.
type InvoiceFinancingRequestV01 struct {

	// Specifies a set of characteristics that unambiguously identify the invoice financing request, such as group identification, creation date time, number of single invoice financing requests.
	RequestGroupInformation *model.RequestGroupInformation1 `xml:"ReqGrpInf"`

	// Set of characteristics that unambiguously identify the single invoice financing request related to the entire invoice or a specific instalment of the invoice settlement, such as actors involved, invoice totals or payment method.
	InvoiceRequestInformation []*model.InvoiceRequestInformation1 `xml:"InvcReqInf"`
}

func (i *InvoiceFinancingRequestV01) AddRequestGroupInformation() *model.RequestGroupInformation1 {
	i.RequestGroupInformation = new(model.RequestGroupInformation1)
	return i.RequestGroupInformation
}

func (i *InvoiceFinancingRequestV01) AddInvoiceRequestInformation() *model.InvoiceRequestInformation1 {
	newValue := new(model.InvoiceRequestInformation1)
	i.InvoiceRequestInformation = append(i.InvoiceRequestInformation, newValue)
	return newValue
}

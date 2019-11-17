package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.006.001.01 Document"`
	Message *InvoiceAssignmentRequestV01 `xml:"InvcAssgnmtReq"`
}

func (d *Document00600101) AddMessage() *InvoiceAssignmentRequestV01 {
	d.Message = new(InvoiceAssignmentRequestV01)
	return d.Message
}

// The InvoiceAssignmentRequest message is sent from a factoring client to a factoring service provider and, optionally, to an interested party. It indicates the transfer of payment obligations concerning financial documents.
// The message contains a list of financing requests together with data that are necessary to transfer the related rights for example regarding legal references for example jurisdiction, language or country. Furthermore, the message can reference related messages and can include data from other messages.
// A factoring client combines a set of financial documents with same characteristics and assigns them to a factoring service. The client can send several assignments in one message and combine them according to different criteria for example for different clients or different currencies.
type InvoiceAssignmentRequestV01 struct {

	// Set of characteristics that unambiguously identify the invoice assigment request, such as group identification, creation date time, number of single invoice financing requests, totals and subtotals.
	Header *model.BusinessLetter1 `xml:"Hdr"`

	// List of assignments of financial items.
	AssignmentList []*model.FinancingItemList1 `xml:"AssgnmtList"`

	// Number of assignments.
	AssignmentCount *model.Max15NumericText `xml:"AssgnmtCnt,omitempty"`

	// Total number of individual items in all assignments.
	ItemCount *model.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *model.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*model.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (i *InvoiceAssignmentRequestV01) AddHeader() *model.BusinessLetter1 {
	i.Header = new(model.BusinessLetter1)
	return i.Header
}

func (i *InvoiceAssignmentRequestV01) AddAssignmentList() *model.FinancingItemList1 {
	newValue := new(model.FinancingItemList1)
	i.AssignmentList = append(i.AssignmentList, newValue)
	return newValue
}

func (i *InvoiceAssignmentRequestV01) SetAssignmentCount(value string) {
	i.AssignmentCount = (*model.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentRequestV01) SetItemCount(value string) {
	i.ItemCount = (*model.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentRequestV01) SetControlSum(value string) {
	i.ControlSum = (*model.DecimalNumber)(&value)
}

func (i *InvoiceAssignmentRequestV01) AddAttachedMessage() *model.EncapsulatedBusinessMessage1 {
	newValue := new(model.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

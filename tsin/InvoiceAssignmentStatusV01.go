package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.007.001.01 Document"`
	Message *InvoiceAssignmentStatusV01 `xml:"InvcAssgnmtSts"`
}

func (d *Document00700101) AddMessage() *InvoiceAssignmentStatusV01 {
	d.Message = new(InvoiceAssignmentStatusV01)
	return d.Message
}

// The message InvoiceAssignmentStatus is sent by a factoring service provider to a factoring client and, optionally, to an interested party as a response to assignments requests.
// The factoring service provider returns a copy of items of corresponding requests together with an information about the status of treatment, for example acceptance, rejection or treatment not yet finished. A rejection can be the result of bad message syntax, but also for other motives such as risk, compliance or covenants.
// For each reported financial item, the factoring service provider includes a reference to the corresponding item of the InvoiceFinancingRequest message and may include the referenced item as well as data from other related and referenced messages.
// The message contains information about other parties to be notified and whether these parties are required to acknowledge the assignment.
// The message can carry digital signatures if required by context.
type InvoiceAssignmentStatusV01 struct {

	// Set of characteristics that unambiguously identify the assignment status, common parameters, documents and identifications.
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

func (i *InvoiceAssignmentStatusV01) AddHeader() *model.BusinessLetter1 {
	i.Header = new(model.BusinessLetter1)
	return i.Header
}

func (i *InvoiceAssignmentStatusV01) AddAssignmentList() *model.FinancingItemList1 {
	newValue := new(model.FinancingItemList1)
	i.AssignmentList = append(i.AssignmentList, newValue)
	return newValue
}

func (i *InvoiceAssignmentStatusV01) SetAssignmentCount(value string) {
	i.AssignmentCount = (*model.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentStatusV01) SetItemCount(value string) {
	i.ItemCount = (*model.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentStatusV01) SetControlSum(value string) {
	i.ControlSum = (*model.DecimalNumber)(&value)
}

func (i *InvoiceAssignmentStatusV01) AddAttachedMessage() *model.EncapsulatedBusinessMessage1 {
	newValue := new(model.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

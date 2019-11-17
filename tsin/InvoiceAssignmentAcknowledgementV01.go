package tsin

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.013.001.01 Document"`
	Message *InvoiceAssignmentAcknowledgementV01 `xml:"InvcAssgnmtAck"`
}

func (d *Document01300101) AddMessage() *InvoiceAssignmentAcknowledgementV01 {
	d.Message = new(InvoiceAssignmentAcknowledgementV01)
	return d.Message
}

// The InvoiceAssignmentAcknowledgement message is sent from a trade partner to communicate the status of payment obligations related to financial items.  The message can be sent independently or as a response to an InvoiceAssignmentNotification message.
// Depending on legal contexts the message may be required as a response to an InvoiceAssignmentNotification message in order for the assignment to become effective.
// The trade party may include references to the corresponding items of an InvoiceAssignmentRequest, InvoiceAssignmentStatus or InvoiceAssignmentNotification or other messages and may include referenced data.
// The message can carry digital signatures if required by context.
type InvoiceAssignmentAcknowledgementV01 struct {

	// Set of characteristics that unambiguously identify the status, common parameters, documents and identifications.
	Header *model.BusinessLetter1 `xml:"Hdr"`

	// List of payment status information.
	PaymentStatusList []*model.FinancingItemList1 `xml:"PmtStsList"`

	// Number of payment information lists as control value.
	PaymentStatusCount *model.Max15NumericText `xml:"PmtStsCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *model.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *model.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*model.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (i *InvoiceAssignmentAcknowledgementV01) AddHeader() *model.BusinessLetter1 {
	i.Header = new(model.BusinessLetter1)
	return i.Header
}

func (i *InvoiceAssignmentAcknowledgementV01) AddPaymentStatusList() *model.FinancingItemList1 {
	newValue := new(model.FinancingItemList1)
	i.PaymentStatusList = append(i.PaymentStatusList, newValue)
	return newValue
}

func (i *InvoiceAssignmentAcknowledgementV01) SetPaymentStatusCount(value string) {
	i.PaymentStatusCount = (*model.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentAcknowledgementV01) SetItemCount(value string) {
	i.ItemCount = (*model.Max15NumericText)(&value)
}

func (i *InvoiceAssignmentAcknowledgementV01) SetControlSum(value string) {
	i.ControlSum = (*model.DecimalNumber)(&value)
}

func (i *InvoiceAssignmentAcknowledgementV01) AddAttachedMessage() *model.EncapsulatedBusinessMessage1 {
	newValue := new(model.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}

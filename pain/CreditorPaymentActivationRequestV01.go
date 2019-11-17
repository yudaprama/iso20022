package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.01 Document"`
	Message *CreditorPaymentActivationRequestV01 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300101) AddMessage() *CreditorPaymentActivationRequestV01 {
	d.Message = new(CreditorPaymentActivationRequestV01)
	return d.Message
}

// Scope
// This message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents.
// It is used to initiate a creditor payment activation request.
type CreditorPaymentActivationRequestV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*model.PaymentInstruction5 `xml:"PmtInf"`
}

func (c *CreditorPaymentActivationRequestV01) AddGroupHeader() *model.GroupHeader45 {
	c.GroupHeader = new(model.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV01) AddPaymentInformation() *model.PaymentInstruction5 {
	newValue := new(model.PaymentInstruction5)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

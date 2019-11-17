package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300103 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.03 Document"`
	Message *CreditorPaymentActivationRequestV03 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300103) AddMessage() *CreditorPaymentActivationRequestV03 {
	d.Message = new(CreditorPaymentActivationRequestV03)
	return d.Message
}

// The CreditorPaymentActivationRequest message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents. It is used to initiate a creditor payment activation request.
type CreditorPaymentActivationRequestV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*model.PaymentInstruction11 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestV03) AddGroupHeader() *model.GroupHeader45 {
	c.GroupHeader = new(model.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV03) AddPaymentInformation() *model.PaymentInstruction11 {
	newValue := new(model.PaymentInstruction11)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

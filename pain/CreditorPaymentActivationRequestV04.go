package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.04 Document"`
	Message *CreditorPaymentActivationRequestV04 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300104) AddMessage() *CreditorPaymentActivationRequestV04 {
	d.Message = new(CreditorPaymentActivationRequestV04)
	return d.Message
}

// The CreditorPaymentActivationRequest message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents. It is used to initiate a creditor payment activation request.
type CreditorPaymentActivationRequestV04 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*model.PaymentInstruction17 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestV04) AddGroupHeader() *model.GroupHeader45 {
	c.GroupHeader = new(model.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV04) AddPaymentInformation() *model.PaymentInstruction17 {
	newValue := new(model.PaymentInstruction17)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

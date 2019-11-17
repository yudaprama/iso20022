package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.02 Document"`
	Message *CreditorPaymentActivationRequestV02 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300102) AddMessage() *CreditorPaymentActivationRequestV02 {
	d.Message = new(CreditorPaymentActivationRequestV02)
	return d.Message
}

// Scope
// This message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents.
// It is used to initiate a creditor payment activation request.
type CreditorPaymentActivationRequestV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*model.PaymentInstruction8 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestV02) AddGroupHeader() *model.GroupHeader45 {
	c.GroupHeader = new(model.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV02) AddPaymentInformation() *model.PaymentInstruction8 {
	newValue := new(model.PaymentInstruction8)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

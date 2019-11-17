package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100106 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.06 Document"`
	Message *CustomerCreditTransferInitiationV06 `xml:"CstmrCdtTrfInitn"`
}

func (d *Document00100106) AddMessage() *CustomerCreditTransferInitiationV06 {
	d.Message = new(CustomerCreditTransferInitiationV06)
	return d.Message
}

// Scope
// The CustomerCreditTransferInitiation message is sent by the initiating party to the forwarding agent or debtor agent. It is used to request movement of funds from the debtor account to a creditor.
// Usage
// The CustomerCreditTransferInitiation message can contain one or more customer credit transfer instructions.
// The CustomerCreditTransferInitiation message is used to exchange:
// - One or more instances of a credit transfer initiation;
// - Payment transactions that result in book transfers at the debtor agent or payments to another financial institution;
// - Payment transactions that result in an electronic cash transfer to the creditor account or in the emission of a cheque.
// The message can be used in a direct or a relay scenario:
// - In a direct scenario, the message is sent directly to the debtor agent. The debtor agent is the account servicer of the debtor.
// - In a relay scenario, the message is sent to a forwarding agent. The forwarding agent acts as a concentrating financial institution. It will forward the CustomerCreditTransferInitiation message to the debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the debtor. This caters for example for the scenario of a payments factory initiating all payments on behalf of a large corporate.
// The CustomerCreditTransferInitiation message can be used in domestic and cross-border scenarios.
// The CustomerCreditTransferInitiation message must not be used by the debtor agent to execute the credit transfer instruction(s). The FIToFICustomerCreditTransfer message must be used instead.
type CustomerCreditTransferInitiationV06 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader48 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the credit transfer initiation.
	PaymentInformation []*model.PaymentInstruction16 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CustomerCreditTransferInitiationV06) AddGroupHeader() *model.GroupHeader48 {
	c.GroupHeader = new(model.GroupHeader48)
	return c.GroupHeader
}

func (c *CustomerCreditTransferInitiationV06) AddPaymentInformation() *model.PaymentInstruction16 {
	newValue := new(model.PaymentInstruction16)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CustomerCreditTransferInitiationV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.02 Document"`
	Message *CustomerCreditTransferInitiationV02 `xml:"pain.001.001.02"`
}

func (d *Document00100102) AddMessage() *CustomerCreditTransferInitiationV02 {
	d.Message = new(CustomerCreditTransferInitiationV02)
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
// If it is agreed to include the payment information related to the debit side only once (i.e. Grouped mode), the PaymentInformation block will be present only once.
// If it is agreed to repeat the payment information related to the debit side (i.e. Single mode), the PaymentInformation block must be present once per occurrence of the CreditTransferTransactionInformation block.
// The CustomerCreditTransferInitiation message also allows for a Mixed mode where the PaymentInformation block can be repeated and each PaymentInformation block can contain one or several CreditTransferTransactionInformation block(s).
// Single
// When grouping is set to Single, information for each individual instruction is included separately. This means the
// PaymentInformation block is repeated, and present for each occurrence of the CreditTransferTransactionInformation block.
// Grouped
// When grouping is set to Grouped, the PaymentInformation block will be present once and the CreditTransferTransactionInformation block will be repeated.
// Mixed
// When grouping is set to Mixed, the PaymentInformation block may be present once or may be repeated. Each sequence
// of the PaymentInformation block may contain one or several CreditTransferTransactionInformation block(s).
type CustomerCreditTransferInitiationV02 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader1 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the credit transfer initiation.
	PaymentInformation []*model.PaymentInstructionInformation1 `xml:"PmtInf"`
}

func (c *CustomerCreditTransferInitiationV02) AddGroupHeader() *model.GroupHeader1 {
	c.GroupHeader = new(model.GroupHeader1)
	return c.GroupHeader
}

func (c *CustomerCreditTransferInitiationV02) AddPaymentInformation() *model.PaymentInstructionInformation1 {
	newValue := new(model.PaymentInstructionInformation1)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

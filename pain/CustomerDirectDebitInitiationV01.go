package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:pain.008.001.01 Document"`
	Message *CustomerDirectDebitInitiationV01 `xml:"pain.008.001.01"`
}

func (d *Document00800101) AddMessage() *CustomerDirectDebitInitiationV01 {
	d.Message = new(CustomerDirectDebitInitiationV01)
	return d.Message
}

// Scope
// The CustomerDirectDebitInitiation message is sent by the initiating party to the forwarding agent or creditor agent. It is used to request single or bulk collection(s) of funds from one or various debtor's account(s) for a creditor.
// Usage
// The CustomerDirectDebitInitiation message can contain one or more direct debit instructions.
// The message can be used in a direct or a relay scenario:
// - In a direct scenario, the message is sent directly to the creditor agent. The creditor agent is the account servicer of the creditor.
// - In a relay scenario, the message is sent to a forwarding agent. The forwarding agent acts as a concentrating financial institution. It will forward the CustomerDirectDebitInitiation message to the creditor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the creditor. This caters for example for the scenario of a payments factory initiating all payments on behalf of a large corporate.
// The CustomerDirectDebitInitiation message can be used in domestic and cross-border scenarios.
// The CustomerDirectDebitInitiation may or may not contain mandate related information, i.e. extracts from a mandate, such as MandateIdentification or DateOfSignature. The CustomerDirectDebitInitiation message must not be considered as a mandate.
// The CustomerDirectDebitInitiation message must not be used by the creditor agent to execute the direct debit instruction(s). The FIToFICustomerDirectDebit message must be used instead.
// If it is agreed to include the payment information related to the credit side only once (i.e. Grouped mode), the PaymentInformation block will be present only once. If it is agreed to repeat the payment information related to the credit side (i.e. Single mode), the PaymentInformation block must be present once per occurrence of the DirectDebitTransactionInformation block. The CustomerDirectDebitInitiation message also allows for a Mixed mode where the PaymentInformation block can be repeated and each PaymentInformation block can contain one or several DirectDebitTransactionInformation block(s).
// Single
// When grouping is set to Single, information for each individual instruction is included separately. This means the
// PaymentInformation block is repeated, and present for each occurrence of the Direct Debit TransactionInformation block.
// Grouped
// When grouping is set to Grouped, the PaymentInformation block will be present once and the Direct Debit
// TransactionInformation block will be repeated.
// Mixed
// When grouping is set to Mixed, the PaymentInformation block may be present once or may be repeated. Each sequence
// of the PaymentInformation block may contain one or several Direct Debit TransactionInformation block(s).
type CustomerDirectDebitInitiationV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader1 `xml:"GrpHdr"`

	// Set of characteristics that apply to the credit side of the payment transactions included in the direct debit transaction initiation.
	PaymentInformation []*model.PaymentInstructionInformation2 `xml:"PmtInf"`
}

func (c *CustomerDirectDebitInitiationV01) AddGroupHeader() *model.GroupHeader1 {
	c.GroupHeader = new(model.GroupHeader1)
	return c.GroupHeader
}

func (c *CustomerDirectDebitInitiationV01) AddPaymentInformation() *model.PaymentInstructionInformation2 {
	newValue := new(model.PaymentInstructionInformation2)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

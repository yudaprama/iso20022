package pain

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00700105 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:pain.007.001.05 Document"`
	Message *CustomerPaymentReversalV05 `xml:"CstmrPmtRvsl"`
}

func (d *Document00700105) AddMessage() *CustomerPaymentReversalV05 {
	d.Message = new(CustomerPaymentReversalV05)
	return d.Message
}

// Scope
// The CustomerPaymentReversal message is sent by the initiating party to the next party in the payment chain. It is used to reverse a payment previously executed.
// Usage
// The CustomerPaymentReversal message is exchanged between a non-financial institution customer and an agent to reverse a CustomerDirectDebitInitiation message that has been settled. The result will be a credit on the debtor account.
// The CustomerPaymentReversal message refers to the original CustomerDirectDebitInitiation message by means of references only or by means of references and a set of elements from the original instruction.
// The CustomerPaymentReversal message can be used in domestic and cross-border scenarios.
type CustomerPaymentReversalV05 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader56 `xml:"GrpHdr"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *model.OriginalGroupHeader3 `xml:"OrgnlGrpInf"`

	// Information concerning the original payment information, to which the reversal message refers.
	OriginalPaymentInformationAndReversal []*model.OriginalPaymentInstruction11 `xml:"OrgnlPmtInfAndRvsl,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CustomerPaymentReversalV05) AddGroupHeader() *model.GroupHeader56 {
	c.GroupHeader = new(model.GroupHeader56)
	return c.GroupHeader
}

func (c *CustomerPaymentReversalV05) AddOriginalGroupInformation() *model.OriginalGroupHeader3 {
	c.OriginalGroupInformation = new(model.OriginalGroupHeader3)
	return c.OriginalGroupInformation
}

func (c *CustomerPaymentReversalV05) AddOriginalPaymentInformationAndReversal() *model.OriginalPaymentInstruction11 {
	newValue := new(model.OriginalPaymentInstruction11)
	c.OriginalPaymentInformationAndReversal = append(c.OriginalPaymentInformationAndReversal, newValue)
	return newValue
}

func (c *CustomerPaymentReversalV05) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

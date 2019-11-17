package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800103 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.03 Document"`
	Message *FIToFICustomerCreditTransferV03 `xml:"FIToFICstmrCdtTrf"`
}

func (d *Document00800103) AddMessage() *FIToFICustomerCreditTransferV03 {
	d.Message = new(FIToFICustomerCreditTransferV03)
	return d.Message
}

// Scope
// The FinancialInstitutionToFinancialInstitutionCustomerCreditTransfer message is sent by the debtor agent to the creditor agent, directly or through other agents and/or a payment clearing and settlement system. It is used to move funds from a debtor account to a creditor.
// Usage
// The FIToFICustomerCreditTransfer message is exchanged between agents and can contain one or more customer credit transfer instructions.
// The FIToFICustomerCreditTransfer message does not allow for grouping: a CreditTransferTransactionInformation block must be present for each credit transfer transaction.
// The FIToFICustomerCreditTransfer message can be used in different ways:
// - If the instructing agent and the instructed agent wish to use their direct account relationship in the currency of the transfer then the message contains both the funds for the customer transfer(s) as well as the payment details;
// - If the instructing agent and the instructed agent have no direct account relationship in the currency of the transfer, or do not wish to use their account relationship, then other (reimbursement) agents will be involved to cover for the customer transfer(s). The FIToFICustomerCreditTransfer contains only the payment details and the instructing agent must cover the customer transfer by sending a FinancialInstitutionCreditTransfer to a reimbursement agent. This payment method is called the Cover method;
// - If more than two financial institutions are involved in the payment chain and if the FIToFICustomerCreditTransfer is sent from one financial institution to the next financial institution in the payment chain, then the payment method is called the Serial method.
// The FIToFICustomerCreditTransfer message can be used in domestic and cross-border scenarios.
type FIToFICustomerCreditTransferV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader49 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*model.CreditTransferTransaction2 `xml:"CdtTrfTxInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (f *FIToFICustomerCreditTransferV03) AddGroupHeader() *model.GroupHeader49 {
	f.GroupHeader = new(model.GroupHeader49)
	return f.GroupHeader
}

func (f *FIToFICustomerCreditTransferV03) AddCreditTransferTransactionInformation() *model.CreditTransferTransaction2 {
	newValue := new(model.CreditTransferTransaction2)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}

func (f *FIToFICustomerCreditTransferV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	f.SupplementaryData = append(f.SupplementaryData, newValue)
	return newValue
}

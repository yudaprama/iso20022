package pacs

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00800101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.01 Document"`
	Message *FIToFICustomerCreditTransferV01 `xml:"pacs.008.001.01"`
}

func (d *Document00800101) AddMessage() *FIToFICustomerCreditTransferV01 {
	d.Message = new(FIToFICustomerCreditTransferV01)
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
//
type FIToFICustomerCreditTransferV01 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *model.GroupHeader2 `xml:"GrpHdr"`

	// Set of elements providing information specific to the individual credit transfer(s).
	CreditTransferTransactionInformation []*model.CreditTransferTransactionInformation2 `xml:"CdtTrfTxInf"`
}

func (f *FIToFICustomerCreditTransferV01) AddGroupHeader() *model.GroupHeader2 {
	f.GroupHeader = new(model.GroupHeader2)
	return f.GroupHeader
}

func (f *FIToFICustomerCreditTransferV01) AddCreditTransferTransactionInformation() *model.CreditTransferTransactionInformation2 {
	newValue := new(model.CreditTransferTransactionInformation2)
	f.CreditTransferTransactionInformation = append(f.CreditTransferTransactionInformation, newValue)
	return newValue
}

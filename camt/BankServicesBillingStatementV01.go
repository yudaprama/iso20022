package camt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document08600101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.086.001.01 Document"`
	Message *BankServicesBillingStatementV01 `xml:"BkSvcsBllgStmt"`
}

func (d *Document08600101) AddMessage() *BankServicesBillingStatementV01 {
	d.Message = new(BankServicesBillingStatementV01)
	return d.Message
}

// Scope
// The BankServicesBillingStatement message is used to send from a Financial Institution (FI) to its wholesale customers (corporations, governments, institutions, etc.), information describing the FI’s billing of services rendered in the form of an electronic statement in a standardised format. The BankServicesBillingStatement is a periodic (usually end of month) recounting of all service chargeable events that occurred during a reporting cycle, typically a calendar month, along with detailed tax and currency translation information. Account balance information, although strongly recommended, is not required.
// Usage
// The BankServicesBillingStatement message is designed to provide details related to invoices (or an advice of debit) which a financial institution may supply to its customers. The BankServicesBillingStatement is not expressly designed to be an invoice, nor to replace invoices currently in use. The message may be used as an invoice by agreement between the sender and the receiver. No regulatory or legislative requirements were considered when creating this message standard. Users of the BankServicesBillingStatment message are cautioned to be aware of any regulatory or legal requirement for invoices before replacing existing invoices.
// The BankServicesBillingStatement message can supply the detail supporting separate invoices or debits but it is not the
// invoice or advice of debit of record. The BankServicesBillingStatement message must accurately reflect all the charge and tax related events that occurred during the calendar month and how the FI and taxing authorities were compensated
// for these events. The BSB does not ask the Financial Institution to revise its established pricing and billing procedures.
// How, when and what the customer is actually charged for remains in place. The BankServicesBillingStatement message asks the Financial Institution to aggregate and report what actually happened during the calendar month.
// The BankServicesBillingStatement message is intended for use with the ISO 20022 Business Application Header.
type BankServicesBillingStatementV01 struct {

	// Provides header details on the billing statement report.
	ReportHeader *model.ReportHeader3 `xml:"RptHdr"`

	// Group of bank services billing statements with the same sender and receiver characteristics.
	BillingStatementGroup []*model.StatementGroup1 `xml:"BllgStmtGrp"`
}

func (b *BankServicesBillingStatementV01) AddReportHeader() *model.ReportHeader3 {
	b.ReportHeader = new(model.ReportHeader3)
	return b.ReportHeader
}

func (b *BankServicesBillingStatementV01) AddBillingStatementGroup() *model.StatementGroup1 {
	newValue := new(model.StatementGroup1)
	b.BillingStatementGroup = append(b.BillingStatementGroup, newValue)
	return newValue
}

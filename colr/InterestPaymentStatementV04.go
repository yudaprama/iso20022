package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01500104 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:colr.015.001.04 Document"`
	Message *InterestPaymentStatementV04 `xml:"IntrstPmtStmt"`
}

func (d *Document01500104) AddMessage() *InterestPaymentStatementV04 {
	d.Message = new(InterestPaymentStatementV04)
	return d.Message
}

// Scope
// The InterestPaymentStatement message is sent by either the collateral giver or its collateral manager to the collateral taker or its collateral manager. It is used to report the interest amounts calculated based on the effective posted collateral amount, over a specific period of time agreed by both parties.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The InterestPaymentStatement message is used for reporting the interest per period on collateral held.
type InterestPaymentStatementV04 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *model.Max35Text `xml:"TxId"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation4 `xml:"Oblgtn"`

	// Provides general information on the report such as the statement identification.
	StatementParameters *model.Statement32 `xml:"StmtParams"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn,omitempty"`

	// Provides details on the interest statement.
	InterestStatement *model.InterestStatement4 `xml:"IntrstStmt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InterestPaymentStatementV04) SetTransactionIdentification(value string) {
	i.TransactionIdentification = (*model.Max35Text)(&value)
}

func (i *InterestPaymentStatementV04) AddAgreement() *model.Agreement4 {
	i.Agreement = new(model.Agreement4)
	return i.Agreement
}

func (i *InterestPaymentStatementV04) AddObligation() *model.Obligation4 {
	i.Obligation = new(model.Obligation4)
	return i.Obligation
}

func (i *InterestPaymentStatementV04) AddStatementParameters() *model.Statement32 {
	i.StatementParameters = new(model.Statement32)
	return i.StatementParameters
}

func (i *InterestPaymentStatementV04) AddPagination() *model.Pagination {
	i.Pagination = new(model.Pagination)
	return i.Pagination
}

func (i *InterestPaymentStatementV04) AddInterestStatement() *model.InterestStatement4 {
	i.InterestStatement = new(model.InterestStatement4)
	return i.InterestStatement
}

func (i *InterestPaymentStatementV04) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}

package secl

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00300103 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:secl.003.001.03 Document"`
	Message *TradeLegStatementV03 `xml:"TradLegStmt"`
}

func (d *Document00300103) AddMessage() *TradeLegStatementV03 {
	d.Message = new(TradeLegStatementV03)
	return d.Message
}

// Scope
// The TradeLegStatement message is sent by the central counterparty (CCP) to a clearing member to report all trades that have been executed by the trading platform.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The TradeLegStatement message may be either sent:
// - during the day (to report trades execution by batch) or
// - as an end of day report.
type TradeLegStatementV03 struct {

	// Provides various statement parameters such as the statement identification, the statement date and time or the statement frequency.
	StatementParameters *model.Statement31 `xml:"StmtParams"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides the identification of the account owner, that is the clearing member (individual clearing member or general clearing member).
	ClearingMember *model.PartyIdentification35Choice `xml:"ClrMmb"`

	// Identifies the clearing member account at the Central counterparty through which the trade must be cleared (sometimes called position account).
	ClearingAccount *model.SecuritiesAccount18 `xml:"ClrAcct,omitempty"`

	// Provides the statement details.
	StatementDetails []*model.TradeLegStatement3 `xml:"StmtDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TradeLegStatementV03) AddStatementParameters() *model.Statement31 {
	t.StatementParameters = new(model.Statement31)
	return t.StatementParameters
}

func (t *TradeLegStatementV03) AddPagination() *model.Pagination {
	t.Pagination = new(model.Pagination)
	return t.Pagination
}

func (t *TradeLegStatementV03) AddClearingMember() *model.PartyIdentification35Choice {
	t.ClearingMember = new(model.PartyIdentification35Choice)
	return t.ClearingMember
}

func (t *TradeLegStatementV03) AddClearingAccount() *model.SecuritiesAccount18 {
	t.ClearingAccount = new(model.SecuritiesAccount18)
	return t.ClearingAccount
}

func (t *TradeLegStatementV03) AddStatementDetails() *model.TradeLegStatement3 {
	newValue := new(model.TradeLegStatement3)
	t.StatementDetails = append(t.StatementDetails, newValue)
	return newValue
}

func (t *TradeLegStatementV03) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

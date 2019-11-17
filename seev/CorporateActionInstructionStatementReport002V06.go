package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04200206 struct {
	XMLName xml.Name                                         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.002.06 Document"`
	Message *CorporateActionInstructionStatementReport002V06 `xml:"CorpActnInstrStmtRpt"`
}

func (d *Document04200206) AddMessage() *CorporateActionInstructionStatementReport002V06 {
	d.Message = new(CorporateActionInstructionStatementReport002V06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionStatementReport message to an account owner or its designated agent to report balances at the safekeeping account level for one or more corporate action events or at the corporate action event level for one or more safekeeping accounts.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionStatementReport002V06 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General characteristics related to a statement which reports information.
	StatementGeneralDetails *model.Statement48 `xml:"StmtGnlDtls"`

	// Account information and detailed account holdings information report for corporate action events.
	AccountAndStatementDetails []*model.AccountIdentification41 `xml:"AcctAndStmtDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionStatementReport002V06) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CorporateActionInstructionStatementReport002V06) AddStatementGeneralDetails() *model.Statement48 {
	c.StatementGeneralDetails = new(model.Statement48)
	return c.StatementGeneralDetails
}

func (c *CorporateActionInstructionStatementReport002V06) AddAccountAndStatementDetails() *model.AccountIdentification41 {
	newValue := new(model.AccountIdentification41)
	c.AccountAndStatementDetails = append(c.AccountAndStatementDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatementReport002V06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04200101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.01 Document"`
	Message *CorporateActionInstructionStatementReportV01 `xml:"CorpActnInstrStmtRpt"`
}

func (d *Document04200101) AddMessage() *CorporateActionInstructionStatementReportV01 {
	d.Message = new(CorporateActionInstructionStatementReportV01)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionStatementReport message to an account owner or its designated agent to report balances at the safekeeping account level for one or more corporate action events or at the corporate action event level for one or more safekeeping accounts.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionStatementReportV01 struct {

	// Information that unambiguously identifies a CorporateActionInstructionStatementReport message as know by the account servicer. When the report has multiple pages, one message equals one page. Therefore, Identification uniquely identifies the page.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General characteristics related to a statement which reports information.
	StatementGeneralDetails *model.Statement12 `xml:"StmtGnlDtls"`

	// Account information and detailed account holdings information report for corporate action events.
	AccountAndStatementDetails []*model.AccountIdentification6 `xml:"AcctAndStmtDtls"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*model.Extension2 `xml:"Xtnsn,omitempty"`
}

func (c *CorporateActionInstructionStatementReportV01) AddIdentification() *model.DocumentIdentification11 {
	c.Identification = new(model.DocumentIdentification11)
	return c.Identification
}

func (c *CorporateActionInstructionStatementReportV01) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CorporateActionInstructionStatementReportV01) AddStatementGeneralDetails() *model.Statement12 {
	c.StatementGeneralDetails = new(model.Statement12)
	return c.StatementGeneralDetails
}

func (c *CorporateActionInstructionStatementReportV01) AddAccountAndStatementDetails() *model.AccountIdentification6 {
	newValue := new(model.AccountIdentification6)
	c.AccountAndStatementDetails = append(c.AccountAndStatementDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatementReportV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	c.MessageOriginator = new(model.PartyIdentification10Choice)
	return c.MessageOriginator
}

func (c *CorporateActionInstructionStatementReportV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	c.MessageRecipient = new(model.PartyIdentification10Choice)
	return c.MessageRecipient
}

func (c *CorporateActionInstructionStatementReportV01) AddExtension() *model.Extension2 {
	newValue := new(model.Extension2)
	c.Extension = append(c.Extension, newValue)
	return newValue
}

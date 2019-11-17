package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600103 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.03 Document"`
	Message *IntraPositionMovementPostingReportV03 `xml:"IntraPosMvmntPstngRpt"`
}

func (d *Document01600103) AddMessage() *IntraPositionMovementPostingReportV03 {
	d.Message = new(IntraPositionMovementPostingReportV03)
	return d.Message
}

// Scope
// An account servicer sends an IntraPositionMovementPostingReport to an account owner to provide the details of increases and decreases in securities with a given status within a holding, ie, intra-position transfers, which occurred during a specified period, for all or selected securities in a specified safekeeping account which the account servicer holds for the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementPostingReportV03 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *model.Statement15 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrument []*model.FinancialInstrumentDetails10 `xml:"FinInstrm,omitempty"`
}

func (i *IntraPositionMovementPostingReportV03) AddPagination() *model.Pagination {
	i.Pagination = new(model.Pagination)
	return i.Pagination
}

func (i *IntraPositionMovementPostingReportV03) AddStatementGeneralDetails() *model.Statement15 {
	i.StatementGeneralDetails = new(model.Statement15)
	return i.StatementGeneralDetails
}

func (i *IntraPositionMovementPostingReportV03) AddAccountOwner() *model.PartyIdentification36Choice {
	i.AccountOwner = new(model.PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementPostingReportV03) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	i.SafekeepingAccount = new(model.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementPostingReportV03) AddFinancialInstrument() *model.FinancialInstrumentDetails10 {
	newValue := new(model.FinancialInstrumentDetails10)
	i.FinancialInstrument = append(i.FinancialInstrument, newValue)
	return newValue
}

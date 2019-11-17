package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600206 struct {
	XMLName xml.Name                                  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.002.06 Document"`
	Message *IntraPositionMovementPostingReport002V06 `xml:"IntraPosMvmntPstngRpt"`
}

func (d *Document01600206) AddMessage() *IntraPositionMovementPostingReport002V06 {
	d.Message = new(IntraPositionMovementPostingReport002V06)
	return d.Message
}

// Scope
// An account servicer sends an IntraPositionMovementPostingReport to an account owner to provide the details of increases and decreases in securities with a given status within a holding, that is, intra-position transfers, which occurred during a specified period, for all or selected securities in a specified safekeeping account which the account servicer holds for the account owner.
//
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// :
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type IntraPositionMovementPostingReport002V06 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *model.Statement49 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrument []*model.FinancialInstrumentDetails26 `xml:"FinInstrm,omitempty"`
}

func (i *IntraPositionMovementPostingReport002V06) AddPagination() *model.Pagination {
	i.Pagination = new(model.Pagination)
	return i.Pagination
}

func (i *IntraPositionMovementPostingReport002V06) AddStatementGeneralDetails() *model.Statement49 {
	i.StatementGeneralDetails = new(model.Statement49)
	return i.StatementGeneralDetails
}

func (i *IntraPositionMovementPostingReport002V06) AddAccountOwner() *model.PartyIdentification103Choice {
	i.AccountOwner = new(model.PartyIdentification103Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementPostingReport002V06) AddSafekeepingAccount() *model.SecuritiesAccount27 {
	i.SafekeepingAccount = new(model.SecuritiesAccount27)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementPostingReport002V06) AddFinancialInstrument() *model.FinancialInstrumentDetails26 {
	newValue := new(model.FinancialInstrumentDetails26)
	i.FinancialInstrument = append(i.FinancialInstrument, newValue)
	return newValue
}

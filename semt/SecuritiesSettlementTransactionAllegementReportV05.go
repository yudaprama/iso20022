package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900105 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.019.001.05 Document"`
	Message *SecuritiesSettlementTransactionAllegementReportV05 `xml:"SctiesSttlmTxAllgmtRpt"`
}

func (d *Document01900105) AddMessage() *SecuritiesSettlementTransactionAllegementReportV05 {
	d.Message = new(SecuritiesSettlementTransactionAllegementReportV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementReport to an account owner to provide, at a specified time, the status and details of pending settlement allegements, for all or selected securities in a specified safekeeping account.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionAllegementReportV05 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *model.Statement39 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Details of the allegement.
	AllegementDetails []*model.SecuritiesTradeDetails49 `xml:"AllgmtDtls,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddStatementGeneralDetails() *model.Statement39 {
	s.StatementGeneralDetails = new(model.Statement39)
	return s.StatementGeneralDetails
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddAccountOwner() *model.PartyIdentification98 {
	s.AccountOwner = new(model.PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddSafekeepingAccount() *model.SecuritiesAccount24 {
	s.SafekeepingAccount = new(model.SecuritiesAccount24)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAllegementReportV05) AddAllegementDetails() *model.SecuritiesTradeDetails49 {
	newValue := new(model.SecuritiesTradeDetails49)
	s.AllegementDetails = append(s.AllegementDetails, newValue)
	return newValue
}

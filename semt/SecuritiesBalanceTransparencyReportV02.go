package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04100102 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.041.001.02 Document"`
	Message *SecuritiesBalanceTransparencyReportV02 `xml:"SctiesBalTrnsprncyRpt"`
}

func (d *Document04100102) AddMessage() *SecuritiesBalanceTransparencyReportV02 {
	d.Message = new(SecuritiesBalanceTransparencyReportV02)
	return d.Message
}

// Scope
// The SecuritiesBalanceTransparencyReport message is sent by an account servicer, such as a custodian, central securities depository or international central securities depository, to the account owner to provide holdings information for the accounts that it services, to disclose underlying details of holdings on an omnibus account that the sender owns or operates at the receiver. The receiver may also be a custodian, central securities depository, international central securities depository, and the ultimate receiver may be a registrar, transfer agent, fund company, official agent of the reported instrument(s) and/or other parties.
// The SecuritiesBalanceTransparencyReport message provides transparency of holdings through layers of custody chains in a consolidated statement, to allow for an efficient gathering of investor data, which, in turn, may be used to measure marketing effectiveness, validation of compliance with prospectuses and regulatory requirements, and the calculation of trailer fees and other retrocessions.
// Usage
// The SecuritiesBalanceTransparencyReport message is used to provide aggregated holdings information and a breakdown of holdings information.
// A sender of the SecuritiesBalanceTransparencyReport message will identify its own safekeeping account (for example, an omnibus account in the ledger of the receiver) and holdings information at the level of account(s) for which the sender is the account servicer (that is, in the ledger of the sender). When relevant, the sender will aggregate its holdings information with  holdings information of one or more sub levels and sub-sub levels of accounts, that is, with holdings information the sender has received from the owner(s) of the account(s)  for which the sender is the account servicer.
// A sender of the SecuritiesBalanceTransparencyReport message may also use it to send statements to its account owning customers, and these can be enrichments of statements that the respective account owners have previously provided to the sender.
// Ultimately, the statement reaches the relevant fund company, for example, the transfer agent, that may use it for obtaining information about the custodians, distributors and commercial agreement references associated with holdings on an omnibus account at the ultimate place of safekeeping, for example, a central securities depository (CSD) or a register of shareholders.
// When the message is sent by the owner of the account specified in SafekeepingAccountAndHoldings/AccountIdentification, the message will disclose holding details of the underlying owner(s) of the sender’s holdings with the receiver. This direction is commonly referred to as ‘downstream’.
// When the sender is the account servicer of an account owned by the receiver, for example, the account in AccountSubLevel1/AccountIdentification or AccountSubLevel2/AccountIdentification, the message is providing a statement of the receiver’s holdings with sender. This direction is commonly referred to as ‘upstream’, and the safekeeping account should identify the ultimate place of safekeeping (for example, an account in a transfer agent's register of shareholders).
type SecuritiesBalanceTransparencyReportV02 struct {

	// Unique and unambiguous identification of the message. When the report has multiple pages, one message equals one page. Therefore, the MessageIdentification uniquely identifies the page.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	// Identification of the party that is the sender of the message.
	SenderIdentification *model.PartyIdentification100 `xml:"SndrId"`

	// Identification of the party that is the receiver of the message.
	ReceiverIdentification *model.PartyIdentification100 `xml:"RcvrId,omitempty"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// Provides general information on the statement.
	StatementGeneralDetails *model.Statement59 `xml:"StmtGnlDtls"`

	// Details of the account, account sub-levels and the holdings.
	SafekeepingAccountAndHoldings []*model.SafekeepingAccount7 `xml:"SfkpgAcctAndHldgs,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesBalanceTransparencyReportV02) AddMessageIdentification() *model.MessageIdentification1 {
	s.MessageIdentification = new(model.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SecuritiesBalanceTransparencyReportV02) AddSenderIdentification() *model.PartyIdentification100 {
	s.SenderIdentification = new(model.PartyIdentification100)
	return s.SenderIdentification
}

func (s *SecuritiesBalanceTransparencyReportV02) AddReceiverIdentification() *model.PartyIdentification100 {
	s.ReceiverIdentification = new(model.PartyIdentification100)
	return s.ReceiverIdentification
}

func (s *SecuritiesBalanceTransparencyReportV02) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesBalanceTransparencyReportV02) AddStatementGeneralDetails() *model.Statement59 {
	s.StatementGeneralDetails = new(model.Statement59)
	return s.StatementGeneralDetails
}

func (s *SecuritiesBalanceTransparencyReportV02) AddSafekeepingAccountAndHoldings() *model.SafekeepingAccount7 {
	newValue := new(model.SafekeepingAccount7)
	s.SafekeepingAccountAndHoldings = append(s.SafekeepingAccountAndHoldings, newValue)
	return newValue
}

func (s *SecuritiesBalanceTransparencyReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

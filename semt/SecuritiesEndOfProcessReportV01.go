package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02300101 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.023.001.01 Document"`
	Message *SecuritiesEndOfProcessReportV01 `xml:"SctiesEndOfPrcRpt"`
}

func (d *Document02300101) AddMessage() *SecuritiesEndOfProcessReportV01 {
	d.Message = new(SecuritiesEndOfProcessReportV01)
	return d.Message
}

// Scope
// Sent by an executing party or an instructing party to the custodian or an affirming party to notify that all the necessary SecuritiesTradeConfirmation message or any other notification of the process have been sent.
// It may also be sent through Central Matching Utility (CMU).
// The instructing party is typically the investment manager or an intermediary system/vendor communicating on behalf of the investment manager.
// The executing party is typically the broker/dealer or an intermediary system/vendor communicating on behalf of the broker/dealer.
// The custodian or an affirming party is typically the custodian, trustee, financial institution, intermediary system/vendor communicating on behalf of them, or their agent.
// The ISO 20022 Business Application Header must be used
// Usage
// Initiator: Executing Party, CMU or Instructing Party
// Respondent: Custodian or an affirming party does not need to respond.
type SecuritiesEndOfProcessReportV01 struct {

	// Number used to sequence pages when it is not possible for data to be conveyed in a single message and the data has to be split across several pages (messages).
	Pagination []*model.Pagination `xml:"Pgntn,omitempty"`

	// Notifies the type of report transmitted.
	ReportGeneralDetails *model.Report3 `xml:"RptGnlDtls"`

	// Parties involved in the confirmation of the details of a trade.
	ConfirmationParties []*model.ConfirmationParties2 `xml:"ConfPties,omitempty"`

	// Party that identifies the underlying investor.
	Investor []*model.PartyIdentificationAndAccount79 `xml:"Invstr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesEndOfProcessReportV01) AddPagination() *model.Pagination {
	newValue := new(model.Pagination)
	s.Pagination = append(s.Pagination, newValue)
	return newValue
}

func (s *SecuritiesEndOfProcessReportV01) AddReportGeneralDetails() *model.Report3 {
	s.ReportGeneralDetails = new(model.Report3)
	return s.ReportGeneralDetails
}

func (s *SecuritiesEndOfProcessReportV01) AddConfirmationParties() *model.ConfirmationParties2 {
	newValue := new(model.ConfirmationParties2)
	s.ConfirmationParties = append(s.ConfirmationParties, newValue)
	return newValue
}

func (s *SecuritiesEndOfProcessReportV01) AddInvestor() *model.PartyIdentificationAndAccount79 {
	newValue := new(model.PartyIdentificationAndAccount79)
	s.Investor = append(s.Investor, newValue)
	return newValue
}

func (s *SecuritiesEndOfProcessReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02400101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.024.001.01 Document"`
	Message *TotalPortfolioValuationReportV01 `xml:"TtlPrtflValtnRpt"`
}

func (d *Document02400101) AddMessage() *TotalPortfolioValuationReportV01 {
	d.Message = new(TotalPortfolioValuationReportV01)
	return d.Message
}

// Scope
// An account servicer sends a TotalPortfolioValuationReport to an account owner to provide detailed valuation information for a portfolio.
// Usage
// The TotalPortfolioValuationReport will be sent by the account servicer to the account owner on an agreed basis. The report may also be requested using a SecuritiesStatementQuery.
// The TotalPortfolioValuationReport is used to report on a portfolio without Investment Funds; or to report on a portfolio when an investment fund is regarded as a portfolio containing, only one or multiple investment funds
// The TotalPortfolioValuationReport may also be used to:
// - re-send a message previously sent (the CopyDuplicate value is DUPL),
// - provide a third party with a copy of a message for information (the CopyDuplicate value is COPY),
// - re-send to a third party a copy of a message for information (the CopyDuplicate value is CODU).
type TotalPortfolioValuationReportV01 struct {

	// Page number of the message (within the report) and continuation indicator to indicate that the report is to continue or that the message is the last page of the report.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General information related to the total portfolio valuation report.
	ReportGeneralDetails *model.Report4 `xml:"RptGnlDtls"`

	// Details of the account. The account may represent an investment portfolio or a fund.
	AccountDetails *model.SecuritiesAccount21 `xml:"AcctDtls"`

	// Valuation information of the portfolio. In some markets a fund of funds or an investment fund is regarded as a portfolio.
	TotalPortfolioValuation *model.TotalPortfolioValuation1 `xml:"TtlPrtflValtn"`

	// Balance breakdown information.
	Balance *model.PortfolioBalance1 `xml:"Bal,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData *model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TotalPortfolioValuationReportV01) AddPagination() *model.Pagination {
	t.Pagination = new(model.Pagination)
	return t.Pagination
}

func (t *TotalPortfolioValuationReportV01) AddReportGeneralDetails() *model.Report4 {
	t.ReportGeneralDetails = new(model.Report4)
	return t.ReportGeneralDetails
}

func (t *TotalPortfolioValuationReportV01) AddAccountDetails() *model.SecuritiesAccount21 {
	t.AccountDetails = new(model.SecuritiesAccount21)
	return t.AccountDetails
}

func (t *TotalPortfolioValuationReportV01) AddTotalPortfolioValuation() *model.TotalPortfolioValuation1 {
	t.TotalPortfolioValuation = new(model.TotalPortfolioValuation1)
	return t.TotalPortfolioValuation
}

func (t *TotalPortfolioValuationReportV01) AddBalance() *model.PortfolioBalance1 {
	t.Balance = new(model.PortfolioBalance1)
	return t.Balance
}

func (t *TotalPortfolioValuationReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	t.SupplementaryData = new(model.SupplementaryData1)
	return t.SupplementaryData
}

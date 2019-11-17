package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.012.001.01 Document"`
	Message *MoneyMarketSecuredMarketStatisticalReportV01 `xml:"MnyMktScrdMktSttstclRpt"`
}

func (d *Document01200101) AddMessage() *MoneyMarketSecuredMarketStatisticalReportV01 {
	d.Message = new(MoneyMarketSecuredMarketStatisticalReportV01)
	return d.Message
}

// The MoneyMarketSecuredMarketStatisticalReport message is sent by the reporting agents to the relevant competent authority, to report all relevant secured money market transactions.
type MoneyMarketSecuredMarketStatisticalReportV01 struct {

	// Provides the elements specific to the report.
	ReportHeader *model.MoneyMarketReportHeader1 `xml:"RptHdr"`

	// Provides the reason why no activity is reported or the required list of transactions for the secured market segment.
	SecuredMarketReport *model.SecuredMarketReport3Choice `xml:"ScrdMktRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MoneyMarketSecuredMarketStatisticalReportV01) AddReportHeader() *model.MoneyMarketReportHeader1 {
	m.ReportHeader = new(model.MoneyMarketReportHeader1)
	return m.ReportHeader
}

func (m *MoneyMarketSecuredMarketStatisticalReportV01) AddSecuredMarketReport() *model.SecuredMarketReport3Choice {
	m.SecuredMarketReport = new(model.SecuredMarketReport3Choice)
	return m.SecuredMarketReport
}

func (m *MoneyMarketSecuredMarketStatisticalReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

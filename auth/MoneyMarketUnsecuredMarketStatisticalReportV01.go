package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300101 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.013.001.01 Document"`
	Message *MoneyMarketUnsecuredMarketStatisticalReportV01 `xml:"MnyMktUscrdMktSttstclRpt"`
}

func (d *Document01300101) AddMessage() *MoneyMarketUnsecuredMarketStatisticalReportV01 {
	d.Message = new(MoneyMarketUnsecuredMarketStatisticalReportV01)
	return d.Message
}

// The MoneyMarketUnsecuredMarketStatisticalReport message is sent by the reporting agents to the relevant competent authority, to report all relevant unsecured money market transactions.
type MoneyMarketUnsecuredMarketStatisticalReportV01 struct {

	// Provides the elements specific to the report.
	ReportHeader *model.MoneyMarketReportHeader1 `xml:"RptHdr"`

	// Provides the reason why no activity is reported or the required list of transactions for the unsecured market segment.
	UnsecuredMarketReport *model.UnsecuredMarketReport3Choice `xml:"UscrdMktRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MoneyMarketUnsecuredMarketStatisticalReportV01) AddReportHeader() *model.MoneyMarketReportHeader1 {
	m.ReportHeader = new(model.MoneyMarketReportHeader1)
	return m.ReportHeader
}

func (m *MoneyMarketUnsecuredMarketStatisticalReportV01) AddUnsecuredMarketReport() *model.UnsecuredMarketReport3Choice {
	m.UnsecuredMarketReport = new(model.UnsecuredMarketReport3Choice)
	return m.UnsecuredMarketReport
}

func (m *MoneyMarketUnsecuredMarketStatisticalReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

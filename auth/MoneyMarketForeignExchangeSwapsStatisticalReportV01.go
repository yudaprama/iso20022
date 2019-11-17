package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400101 struct {
	XMLName xml.Name                                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.014.001.01 Document"`
	Message *MoneyMarketForeignExchangeSwapsStatisticalReportV01 `xml:"MnyMktFXSwpsSttstclRpt"`
}

func (d *Document01400101) AddMessage() *MoneyMarketForeignExchangeSwapsStatisticalReportV01 {
	d.Message = new(MoneyMarketForeignExchangeSwapsStatisticalReportV01)
	return d.Message
}

// The MoneyMarketSecuredMarketStatisticalReport message is sent by the reporting agents  to the relevant competent authority, to report all daily Foreign Exchange Swaps (FX Swaps) transactions.
type MoneyMarketForeignExchangeSwapsStatisticalReportV01 struct {

	// Provides the elements specific to the report.
	ReportHeader *model.MoneyMarketReportHeader1 `xml:"RptHdr"`

	// Provides the reason why no activity is reported or the required list of transactions for the foreign exchange swaps segment.
	ForeignExchangeSwapsReport *model.ForeignExchangeSwap2Choice `xml:"FXSwpsRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MoneyMarketForeignExchangeSwapsStatisticalReportV01) AddReportHeader() *model.MoneyMarketReportHeader1 {
	m.ReportHeader = new(model.MoneyMarketReportHeader1)
	return m.ReportHeader
}

func (m *MoneyMarketForeignExchangeSwapsStatisticalReportV01) AddForeignExchangeSwapsReport() *model.ForeignExchangeSwap2Choice {
	m.ForeignExchangeSwapsReport = new(model.ForeignExchangeSwap2Choice)
	return m.ForeignExchangeSwapsReport
}

func (m *MoneyMarketForeignExchangeSwapsStatisticalReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

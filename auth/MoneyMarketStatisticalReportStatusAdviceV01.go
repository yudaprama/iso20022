package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02800101 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.028.001.01 Document"`
	Message *MoneyMarketStatisticalReportStatusAdviceV01 `xml:"MnyMktSttstclRptStsAdvc"`
}

func (d *Document02800101) AddMessage() *MoneyMarketStatisticalReportStatusAdviceV01 {
	d.Message = new(MoneyMarketStatisticalReportStatusAdviceV01)
	return d.Message
}

// The MoneyMarketStatisticalReportStatusAdvice message is sent by the relevant competent authority to the reporting agents to provide the status on the reported transactions.
type MoneyMarketStatisticalReportStatusAdviceV01 struct {

	// Provides the status on the global report.
	StatusReportHeader *model.MoneyMarketStatusReportHeader1 `xml:"StsRptHdr"`

	// Provides the status on an individual transaction and the related reason if required.
	TransactionStatus []*model.MoneyMarketTransactionStatus2 `xml:"TxSts,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (m *MoneyMarketStatisticalReportStatusAdviceV01) AddStatusReportHeader() *model.MoneyMarketStatusReportHeader1 {
	m.StatusReportHeader = new(model.MoneyMarketStatusReportHeader1)
	return m.StatusReportHeader
}

func (m *MoneyMarketStatisticalReportStatusAdviceV01) AddTransactionStatus() *model.MoneyMarketTransactionStatus2 {
	newValue := new(model.MoneyMarketTransactionStatus2)
	m.TransactionStatus = append(m.TransactionStatus, newValue)
	return newValue
}

func (m *MoneyMarketStatisticalReportStatusAdviceV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	m.SupplementaryData = append(m.SupplementaryData, newValue)
	return newValue
}

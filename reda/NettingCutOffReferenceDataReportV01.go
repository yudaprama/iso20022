package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document06100101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.061.001.01 Document"`
	Message *NettingCutOffReferenceDataReportV01 `xml:"NetgCutOffRefDataRpt"`
}

func (d *Document06100101) AddMessage() *NettingCutOffReferenceDataReportV01 {
	d.Message = new(NettingCutOffReferenceDataReportV01)
	return d.Message
}

// The Netting Cut Off Reference Data Report message is sent to a participant by a central system to provide details of scheduled, changed, existing and previous netting cut off data held at a central system.
type NettingCutOffReferenceDataReportV01 struct {

	// Specifies the meta data for the netting cut off report including message pagination.
	ReportData *model.NettingCutOffReportData1 `xml:"RptData"`

	// Provides the latest information related to the status of a netting cut off held at a central system.
	ParticipantNettingCutOffData []*model.CutOffData1 `xml:"PtcptNetgCutOffData"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (n *NettingCutOffReferenceDataReportV01) AddReportData() *model.NettingCutOffReportData1 {
	n.ReportData = new(model.NettingCutOffReportData1)
	return n.ReportData
}

func (n *NettingCutOffReferenceDataReportV01) AddParticipantNettingCutOffData() *model.CutOffData1 {
	newValue := new(model.CutOffData1)
	n.ParticipantNettingCutOffData = append(n.ParticipantNettingCutOffData, newValue)
	return newValue
}

func (n *NettingCutOffReferenceDataReportV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	n.SupplementaryData = append(n.SupplementaryData, newValue)
	return newValue
}

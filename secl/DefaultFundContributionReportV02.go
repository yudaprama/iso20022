package secl

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00600102 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:secl.006.001.02 Document"`
	Message *DefaultFundContributionReportV02 `xml:"DfltFndCntrbtnRpt"`
}

func (d *Document00600102) AddMessage() *DefaultFundContributionReportV02 {
	d.Message = new(DefaultFundContributionReportV02)
	return d.Message
}

// Scope
// The DefaultFundContributionReport message is sent by the central counterparty (CCP) to a Clearing member to report on the calculation of the default fund contribution and the potential net excess or deficit.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// CCPs require participants to post assets in a clearing fund that can be used in the event of a default by a participant, to compensate non-defaulting participants for losses they suffer due to this default. The CCP evaluates each participant risk level based on their positions (monthly or daily) and calculate the excess of deficit of the default fund contribution. The DefaultFundContributionReport is usually sent on a monthly basis.
type DefaultFundContributionReportV02 struct {

	// Provides details about the report such as the report identification, the calculation date, the value date.
	ReportParameters *model.ReportParameters2 `xml:"RptParams"`

	// Provides the identification of the account owner, that is the clearing member (individual clearing member or general clearing member).
	ClearingMember *model.PartyIdentification35Choice `xml:"ClrMmb"`

	// Provides details on the default fund report.
	ReportDetails []*model.DefaultFundReport1 `xml:"RptDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DefaultFundContributionReportV02) AddReportParameters() *model.ReportParameters2 {
	d.ReportParameters = new(model.ReportParameters2)
	return d.ReportParameters
}

func (d *DefaultFundContributionReportV02) AddClearingMember() *model.PartyIdentification35Choice {
	d.ClearingMember = new(model.PartyIdentification35Choice)
	return d.ClearingMember
}

func (d *DefaultFundContributionReportV02) AddReportDetails() *model.DefaultFundReport1 {
	newValue := new(model.DefaultFundReport1)
	d.ReportDetails = append(d.ReportDetails, newValue)
	return newValue
}

func (d *DefaultFundContributionReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}

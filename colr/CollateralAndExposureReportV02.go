package colr

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600102 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.02 Document"`
	Message *CollateralAndExposureReportV02 `xml:"CollAndXpsrRpt"`
}

func (d *Document01600102) AddMessage() *CollateralAndExposureReportV02 {
	d.Message = new(CollateralAndExposureReportV02)
	return d.Message
}

// Scope
// This message is sent:
// - either by the collateral giver, or its collateral manager, to the collateral taker, or its collateral manager, or
// - or by the collateral taker, or its collateral manager to the collateral giver, or its collateral manager
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralAndExposureReport is used to provide the details of the valuation of the collateral, that is, the valuation of securities collateral, cash collateral or other type of collateral, posted at a specific calculation date.
type CollateralAndExposureReportV02 struct {

	// Provides information about the report such as the report identification, the report date and time or the report frequency.
	ReportParameters *model.ReportParameters2 `xml:"RptParams"`

	// Specifies the page number and an indicator of whether it is the only or last page, or if there are additional pages.
	//
	Pagination *model.Pagination `xml:"Pgntn,omitempty"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *model.Obligation3 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *model.Agreement2 `xml:"Agrmt,omitempty"`

	// Provides details on the collateral report.
	CollateralReport []*model.Collateral9 `xml:"CollRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralAndExposureReportV02) AddReportParameters() *model.ReportParameters2 {
	c.ReportParameters = new(model.ReportParameters2)
	return c.ReportParameters
}

func (c *CollateralAndExposureReportV02) AddPagination() *model.Pagination {
	c.Pagination = new(model.Pagination)
	return c.Pagination
}

func (c *CollateralAndExposureReportV02) AddObligation() *model.Obligation3 {
	c.Obligation = new(model.Obligation3)
	return c.Obligation
}

func (c *CollateralAndExposureReportV02) AddAgreement() *model.Agreement2 {
	c.Agreement = new(model.Agreement2)
	return c.Agreement
}

func (c *CollateralAndExposureReportV02) AddCollateralReport() *model.Collateral9 {
	newValue := new(model.Collateral9)
	c.CollateralReport = append(c.CollateralReport, newValue)
	return newValue
}

func (c *CollateralAndExposureReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

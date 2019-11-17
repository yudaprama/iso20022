package reda

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00400102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.004.001.02 Document"`
	Message *FundProcessingPassportReportV02 `xml:"FndPrcgPsptRpt"`
}

func (d *Document00400102) AddMessage() *FundProcessingPassportReportV02 {
	d.Message = new(FundProcessingPassportReportV02)
	return d.Message
}

// Scope
// A report provider, for example, a fund promoter, fund management company, transfer agent, or market data provider, sends the FundProcessingPassportReport message to the report recipient, for, a professional investor, investment fund distributor, market data provider, regulator or other interested party to provide the key reference data for financial instruments to facilitate trading.
// Usage
// A unique FundProcessingPassportReport should be prepared for each class of unit/share (for which an individual ISIN should have been allocated), in respect of its "home" market.
// The FundProcessingPassportReport may be used in various models or environments:
// - stand alone environment, for example, initiated by the Report Provider (fund promoter, fund manager and / or reference data vendors) sent on a regular frequency, or when changes are needed.
// - in a request / response environment, with the InvestmentFundReportRequest, for example, initiated by report users (data vendors, professional investors, regulators or investment fund distributors) in enabling the user to control the flow and updates of information.
// - in a reference data vendor environment, for example, market infrastructure and reference data providers may collate and store all fund processing passport information centrally for access via database or regular distribution information. A reference data vendor may assume the role of both report provider and report user.
type FundProcessingPassportReportV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *model.MessageIdentification1 `xml:"MsgId"`

	//     Fund Processing Passsport (FPP) is a fully harmonised document with all key operational information that fund promoters
	//     should provide on their investment funds in order to facilitate their trading.
	FundProcessingPassport []*model.FundProcessingPassport1 `xml:"FPP"`
}

func (f *FundProcessingPassportReportV02) AddMessageIdentification() *model.MessageIdentification1 {
	f.MessageIdentification = new(model.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundProcessingPassportReportV02) AddFundProcessingPassport() *model.FundProcessingPassport1 {
	newValue := new(model.FundProcessingPassport1)
	f.FundProcessingPassport = append(f.FundProcessingPassport, newValue)
	return newValue
}

package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01400102 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.014.001.02 Document"`
	Message *AccountReportV02 `xml:"AcctRpt"`
}

func (d *Document01400102) AddMessage() *AccountReportV02 {
	d.Message = new(AccountReportV02)
	return d.Message
}

// The AccountReport message is sent from a financial institution to an organisation for reporting purposes. It can be sent unsolicited as part of opening, maintenance, or closing process, or it can be sent as response to an AccountReportRequest message.
type AccountReportV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References5 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *model.Organisation12 `xml:"Org"`

	// Account report.
	Report []*model.AccountReport15 `xml:"Rpt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountReportV02) AddReferences() *model.References5 {
	a.References = new(model.References5)
	return a.References
}

func (a *AccountReportV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountReportV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountReportV02) AddOrganisation() *model.Organisation12 {
	a.Organisation = new(model.Organisation12)
	return a.Organisation
}

func (a *AccountReportV02) AddReport() *model.AccountReport15 {
	newValue := new(model.AccountReport15)
	a.Report = append(a.Report, newValue)
	return newValue
}

func (a *AccountReportV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountReportV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

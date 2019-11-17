package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01300102 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.013.001.02 Document"`
	Message *AccountReportRequestV02 `xml:"AcctRptReq"`
}

func (d *Document01300102) AddMessage() *AccountReportRequestV02 {
	d.Message = new(AccountReportRequestV02)
	return d.Message
}

// The AccountReportRequest message is sent from an organisation to a financial institution for reporting purposes. It is a request for an account report. This message can be sent at any time outside of account opening, maintenance or closing processes.
type AccountReportRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*model.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the report.
	OrganisationIdentification *model.OrganisationIdentification8 `xml:"OrgId"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountReportRequestV02) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountReportRequestV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountReportRequestV02) AddAccountIdentification() *model.AccountForAction1 {
	newValue := new(model.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountReportRequestV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountReportRequestV02) AddOrganisationIdentification() *model.OrganisationIdentification8 {
	a.OrganisationIdentification = new(model.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountReportRequestV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountReportRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

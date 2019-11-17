package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01600102 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.016.001.02 Document"`
	Message *AccountExcludedMandateMaintenanceAmendmentRequestV02 `xml:"AcctExcldMndtMntncAmdmntReq"`
}

func (d *Document01600102) AddMessage() *AccountExcludedMandateMaintenanceAmendmentRequestV02 {
	d.Message = new(AccountExcludedMandateMaintenanceAmendmentRequestV02)
	return d.Message
}

// The AccountExcludedMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information. Usage: this update is about account details excluding any mandate information.
// If modification codes are not used: the organisation will specify under the “Account” and “Organisation” tags the complete information as it should be in the financial institution’s records after processing the update request.
// If modification codes are used (in that case, they must be used everywhere): the organisation will specify under the “Account” and “Organisation” tags which elements must be added, deleted, modified, or if they are unchanged.
type AccountExcludedMandateMaintenanceAmendmentRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution. OrganisationIdentification6
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Specifies target dates.
	ContractDates *model.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	Account *model.CustomerAccountModification1 `xml:"Acct"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Organised structure that is set up for a particular purpose, for example, a business, government body, department, charity, or financial institution.
	Organisation *model.OrganisationModification1 `xml:"Org"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddAccount() *model.CustomerAccountModification1 {
	a.Account = new(model.CustomerAccountModification1)
	return a.Account
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddOrganisation() *model.OrganisationModification1 {
	a.Organisation = new(model.OrganisationModification1)
	return a.Organisation
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountExcludedMandateMaintenanceAmendmentRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

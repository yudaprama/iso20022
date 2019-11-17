package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01800101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.018.001.01 Document"`
	Message *AccountMandateMaintenanceAmendmentRequestV01 `xml:"AcctMndtMntncAmdmntReq"`
}

func (d *Document01800101) AddMessage() *AccountMandateMaintenanceAmendmentRequestV01 {
	d.Message = new(AccountMandateMaintenanceAmendmentRequestV01)
	return d.Message
}

// Scope
// The AccountMandateMaintenanceAmendmentRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is sent in response to a request from the financial institution to send additional information. This update is only about mandate information.
// Usage
// This message should only be sent in response to a request from the financial institution. This update is only about mandate information. The organisation will specify under the Mandate tag the complete information as it should be in the financial institutions records after processing the update request. It is not possible to update the account characteristics or organisation information with this message.
// It is possible to request to update the mandate information of several accounts, if these accounts must have exactly the same mandates.
// This message could be sent together with other related documents.
type AccountMandateMaintenanceAmendmentRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Specifies target dates.
	ContractDates *model.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*model.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification []*model.OrganisationIdentification6 `xml:"OrgId"`

	// Information specifying the account mandate.
	Mandate []*model.OperationMandate1 `xml:"Mndt,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddAccountIdentification() *model.AccountForAction1 {
	newValue := new(model.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddOrganisationIdentification() *model.OrganisationIdentification6 {
	newValue := new(model.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddMandate() *model.OperationMandate1 {
	newValue := new(model.OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceAmendmentRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

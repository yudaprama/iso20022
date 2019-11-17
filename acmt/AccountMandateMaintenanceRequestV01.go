package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01700101 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.017.001.01 Document"`
	Message *AccountMandateMaintenanceRequestV01 `xml:"AcctMndtMntncReq"`
}

func (d *Document01700101) AddMessage() *AccountMandateMaintenanceRequestV01 {
	d.Message = new(AccountMandateMaintenanceRequestV01)
	return d.Message
}

// Scope
// The AccountMandateMaintenanceRequest message is sent from an organisation to a financial institution as part of the account maintenance process. It is the initial request message to update an account. This update is only about mandate information.
// Usage
// This message should only be used for initiating the maintenance process of mandate information. This update is only about mandate information. The organisation will specify under the Mandate tag the complete information as it should be in the financial institutions records after processing the update request. It is not possible to update the account characteristics or organisation information with this message.
// It is possible to request to update the mandate information of several accounts, if these accounts must have exactly the same mandates.
// This message could be sent together with other related documents.
type AccountMandateMaintenanceRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References4 `xml:"Refs"`

	// Specifies target dates.
	ContractDates *model.AccountContract2 `xml:"CtrctDts,omitempty"`

	// Account contract established between the organisation or the group to which the organisation belongs, and the account servicer. This contract has to be applied for the new account to be opened and maintained.
	UnderlyingMasterAgreement *model.ContractDocument1 `xml:"UndrlygMstrAgrmt,omitempty"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*model.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Identification of the organisation requesting the change.
	OrganisationIdentification []*model.OrganisationIdentification6 `xml:"OrgId"`

	// Information specifying the account mandate.
	Mandate []*model.OperationMandate1 `xml:"Mndt"`

	// Contains additional information related to the message.
	AdditionalMessageInformation *model.AdditionalInformation5 `xml:"AddtlMsgInf,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountMandateMaintenanceRequestV01) AddReferences() *model.References4 {
	a.References = new(model.References4)
	return a.References
}

func (a *AccountMandateMaintenanceRequestV01) AddContractDates() *model.AccountContract2 {
	a.ContractDates = new(model.AccountContract2)
	return a.ContractDates
}

func (a *AccountMandateMaintenanceRequestV01) AddUnderlyingMasterAgreement() *model.ContractDocument1 {
	a.UnderlyingMasterAgreement = new(model.ContractDocument1)
	return a.UnderlyingMasterAgreement
}

func (a *AccountMandateMaintenanceRequestV01) AddAccountIdentification() *model.AccountForAction1 {
	newValue := new(model.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceRequestV01) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountMandateMaintenanceRequestV01) AddOrganisationIdentification() *model.OrganisationIdentification6 {
	newValue := new(model.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceRequestV01) AddMandate() *model.OperationMandate1 {
	newValue := new(model.OperationMandate1)
	a.Mandate = append(a.Mandate, newValue)
	return newValue
}

func (a *AccountMandateMaintenanceRequestV01) AddAdditionalMessageInformation() *model.AdditionalInformation5 {
	a.AdditionalMessageInformation = new(model.AdditionalInformation5)
	return a.AdditionalMessageInformation
}

func (a *AccountMandateMaintenanceRequestV01) AddDigitalSignature() *model.PartyAndSignature1 {
	newValue := new(model.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02100102 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.021.001.02 Document"`
	Message *AccountClosingAdditionalInformationRequestV02 `xml:"AcctClsgAddtlInfReq"`
}

func (d *Document02100102) AddMessage() *AccountClosingAdditionalInformationRequestV02 {
	d.Message = new(AccountClosingAdditionalInformationRequestV02)
	return d.Message
}

// The AccountClosingAdditionalInformationRequest message is sent from a financial institution to an organisation as part of the account closing process. This message is sent in response to a closing request message from the organisation, if the business content is valid, but additional information is required.
type AccountClosingAdditionalInformationRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References3 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Identifier for an organisation.
	OrganisationIdentification *model.OrganisationIdentification8 `xml:"OrgId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification *model.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Identification of the account to which the remaining positive balance of the account to be closed must be transferred or account from which funds can be moved to the account to be closed and which balance is negative. This account must be held in the same financial institution as the account to be closed if the transfer account is used to compensate a negative balance. For a positive balance to be transferred, an account in another financial institution might be used. In that case the account servicer is mandatory.
	BalanceTransferAccount *model.AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme, that is the servicer of the transfer account.
	TransferAccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"TrfAcctSvcrId,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountClosingAdditionalInformationRequestV02) AddReferences() *model.References3 {
	a.References = new(model.References3)
	return a.References
}

func (a *AccountClosingAdditionalInformationRequestV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountClosingAdditionalInformationRequestV02) AddOrganisationIdentification() *model.OrganisationIdentification8 {
	a.OrganisationIdentification = new(model.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountClosingAdditionalInformationRequestV02) AddAccountIdentification() *model.AccountForAction1 {
	a.AccountIdentification = new(model.AccountForAction1)
	return a.AccountIdentification
}

func (a *AccountClosingAdditionalInformationRequestV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountClosingAdditionalInformationRequestV02) AddBalanceTransferAccount() *model.AccountForAction1 {
	a.BalanceTransferAccount = new(model.AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountClosingAdditionalInformationRequestV02) AddTransferAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.TransferAccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.TransferAccountServicerIdentification
}

func (a *AccountClosingAdditionalInformationRequestV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountClosingAdditionalInformationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}

package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.001.001.01 Document"`
	Message *InformationRequestOpeningV01 `xml:"InfReqOpng"`
}

func (d *Document00100101) AddMessage() *InformationRequestOpeningV01 {
	d.Message = new(InformationRequestOpeningV01)
	return d.Message
}

// This message is sent by the authorities (police, customs, tax authorities, enforcement authorities) to a financial institution to request account and other banking and financial information. Requested information can relate to accounts, their signatories and beneficiaries and co-owners as well as movements plus positions on these accounts.
//
// Requests are underpinned by specific legal texts.
type InformationRequestOpeningV01 struct {

	// Unique identification for the specific investigation as known by the requesting party.
	InvestigationIdentification *model.Max35Text `xml:"InvstgtnId"`

	// Provides details on the legal basis of the request.
	LegalMandateBasis *model.LegalMandate1 `xml:"LglMndtBsis"`

	// Specifies the confidentiality status of the investigation.
	ConfidentialityStatus *model.YesNoIndicator `xml:"CnfdtltySts"`

	// Specifies the date by when the financial institutiion needs to provide a response.
	DueDate *model.DueDate1 `xml:"DueDt,omitempty"`

	// Specifies the dates between which period the authority requests that the financial institution provides a response to the information request.
	InvestigationPeriod *model.DateOrDateTimePeriodChoice `xml:"InvstgtnPrd"`

	// Specifies the the search criteria for the financial institution to perform the search on. The search criteria can be an account, a customer identification or a payment instrument type.
	SearchCriteria *model.SearchCriteria1Choice `xml:"SchCrit"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InformationRequestOpeningV01) SetInvestigationIdentification(value string) {
	i.InvestigationIdentification = (*model.Max35Text)(&value)
}

func (i *InformationRequestOpeningV01) AddLegalMandateBasis() *model.LegalMandate1 {
	i.LegalMandateBasis = new(model.LegalMandate1)
	return i.LegalMandateBasis
}

func (i *InformationRequestOpeningV01) SetConfidentialityStatus(value string) {
	i.ConfidentialityStatus = (*model.YesNoIndicator)(&value)
}

func (i *InformationRequestOpeningV01) AddDueDate() *model.DueDate1 {
	i.DueDate = new(model.DueDate1)
	return i.DueDate
}

func (i *InformationRequestOpeningV01) AddInvestigationPeriod() *model.DateOrDateTimePeriodChoice {
	i.InvestigationPeriod = new(model.DateOrDateTimePeriodChoice)
	return i.InvestigationPeriod
}

func (i *InformationRequestOpeningV01) AddSearchCriteria() *model.SearchCriteria1Choice {
	i.SearchCriteria = new(model.SearchCriteria1Choice)
	return i.SearchCriteria
}

func (i *InformationRequestOpeningV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}

package tsrv

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document00100101 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.001.001.01 Document"`
	Message *UndertakingIssuanceV01 `xml:"UdrtkgIssnc"`
}

func (d *Document00100101) AddMessage() *UndertakingIssuanceV01 {
	d.Message = new(UndertakingIssuanceV01)
	return d.Message
}

// The UndertakingIssuance message is sent (and is thus issued) by the party issuing the undertaking to the beneficiary. The message may be sent either to the beneficiary directly or via an advising party. The undertaking could be a demand guarantee, standby letter of credit, or counter-undertaking (counter-guarantee or counter-standby). It contains details on the applicable rules, expiry date, the amount, required documents, and terms and conditions of the undertaking. The message constitutes an operative financial instrument.
// Under the United Nations Convention on Independent Guarantees and Stand-by Letters of Credit (http://www.uncitral.org), 1996, Article 2, "an undertaking is an independent commitment, known in international practice as an independent guarantee or as a standby letter of credit, given by a bank or other institution or person ('guarantor/issuer') to pay to the beneficiary a certain or determinable amount upon simple demand or upon demand accompanied by other documents, in conformity with the terms and any documentary conditions of the undertaking, indicating, or from which it is to be inferred, that payment is due because of a default in the performance of an obligation, or because of another contingency, or for money borrowed or advanced, or on account of any mature indebtedness undertaken by the principal/applicant or another person".
type UndertakingIssuanceV01 struct {

	// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be collected on the presentation of documents that comply with its terms and conditions.
	UndertakingIssuanceDetails *model.Undertaking3 `xml:"UdrtkgIssncDtls"`

	// Additional information specific to the bank-to-beneficiary communication.
	BankToBeneficiaryInformation []*model.Max2000Text `xml:"BkToBnfcryInf,omitempty"`

	// Additional information specific to the bank-to-bank communication.
	BankToBankInformation []*model.Max2000Text `xml:"BkToBkInf,omitempty"`

	// Digital signature of the undertaking.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingIssuanceV01) AddUndertakingIssuanceDetails() *model.Undertaking3 {
	u.UndertakingIssuanceDetails = new(model.Undertaking3)
	return u.UndertakingIssuanceDetails
}

func (u *UndertakingIssuanceV01) AddBankToBeneficiaryInformation(value string) {
	u.BankToBeneficiaryInformation = append(u.BankToBeneficiaryInformation, (*model.Max2000Text)(&value))
}

func (u *UndertakingIssuanceV01) AddBankToBankInformation(value string) {
	u.BankToBankInformation = append(u.BankToBankInformation, (*model.Max2000Text)(&value))
}

func (u *UndertakingIssuanceV01) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	u.DigitalSignature = append(u.DigitalSignature, newValue)
	return newValue
}

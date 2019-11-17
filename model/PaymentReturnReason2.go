package model

// Provides further details on the reason of the return of the transaction.
type PaymentReturnReason2 struct {

	// Bank transaction code included in the original entry for the transaction.
	OriginalBankTransactionCode *BankTransactionCodeStructure4 `xml:"OrgnlBkTxCd,omitempty"`

	// Party that issues the return.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the return.
	Reason *ReturnReason5Choice `xml:"Rsn,omitempty"`

	// Further details on the return reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentReturnReason2) AddOriginalBankTransactionCode() *BankTransactionCodeStructure4 {
	p.OriginalBankTransactionCode = new(BankTransactionCodeStructure4)
	return p.OriginalBankTransactionCode
}

func (p *PaymentReturnReason2) AddOriginator() *PartyIdentification43 {
	p.Originator = new(PartyIdentification43)
	return p.Originator
}

func (p *PaymentReturnReason2) AddReason() *ReturnReason5Choice {
	p.Reason = new(ReturnReason5Choice)
	return p.Reason
}

func (p *PaymentReturnReason2) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max105Text)(&value))
}

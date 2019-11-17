package model

// Set of elements used to provide information on the reason of the return of the transaction.
type ReturnReasonInformation10 struct {

	// Bank transaction code included in the original entry for the transaction.
	OriginalBankTransactionCode *BankTransactionCodeStructure4 `xml:"OrgnlBkTxCd,omitempty"`

	// Party that issues the return.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the return.
	Reason *ReturnReason5Choice `xml:"Rsn,omitempty"`

	// Further details on the return reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (r *ReturnReasonInformation10) AddOriginalBankTransactionCode() *BankTransactionCodeStructure4 {
	r.OriginalBankTransactionCode = new(BankTransactionCodeStructure4)
	return r.OriginalBankTransactionCode
}

func (r *ReturnReasonInformation10) AddOriginator() *PartyIdentification32 {
	r.Originator = new(PartyIdentification32)
	return r.Originator
}

func (r *ReturnReasonInformation10) AddReason() *ReturnReason5Choice {
	r.Reason = new(ReturnReason5Choice)
	return r.Reason
}

func (r *ReturnReasonInformation10) AddAdditionalInformation(value string) {
	r.AdditionalInformation = append(r.AdditionalInformation, (*Max105Text)(&value))
}

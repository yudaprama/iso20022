package model

// Further information on the return reason of the transaction.
type ReturnReasonInformation5 struct {

	// Bank transaction code included in the original entry for the transaction.
	OriginalBankTransactionCode *BankTransactionCodeStructure1 `xml:"OrgnlBkTxCd,omitempty"`

	// Party issuing the return.
	ReturnOriginator *PartyIdentification8 `xml:"RtrOrgtr,omitempty"`

	// Specifies the reason for the return.
	ReturnReason *ReturnReason1Choice `xml:"RtrRsn,omitempty"`

	// Further details on the return reason.
	AdditionalReturnReasonInformation []*Max105Text `xml:"AddtlRtrRsnInf,omitempty"`
}

func (r *ReturnReasonInformation5) AddOriginalBankTransactionCode() *BankTransactionCodeStructure1 {
	r.OriginalBankTransactionCode = new(BankTransactionCodeStructure1)
	return r.OriginalBankTransactionCode
}

func (r *ReturnReasonInformation5) AddReturnOriginator() *PartyIdentification8 {
	r.ReturnOriginator = new(PartyIdentification8)
	return r.ReturnOriginator
}

func (r *ReturnReasonInformation5) AddReturnReason() *ReturnReason1Choice {
	r.ReturnReason = new(ReturnReason1Choice)
	return r.ReturnReason
}

func (r *ReturnReasonInformation5) AddAdditionalReturnReasonInformation(value string) {
	r.AdditionalReturnReasonInformation = append(r.AdditionalReturnReasonInformation, (*Max105Text)(&value))
}

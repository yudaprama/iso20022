package model

// Specifies the withdrawal reason code and optionally a withdrawal reason sub code.
type WithdrawalReason1 struct {

	// Withdrawal reason expressed as a code.
	WithdrawalReasonCode *WithdrawalReason1Code `xml:"WdrwlRsnCd"`

	// Further withdrawal reason information expressed as a code.
	WithdrawalReasonSubCode *Max4Text `xml:"WdrwlRsnSubCd,omitempty"`
}

func (w *WithdrawalReason1) SetWithdrawalReasonCode(value string) {
	w.WithdrawalReasonCode = (*WithdrawalReason1Code)(&value)
}

func (w *WithdrawalReason1) SetWithdrawalReasonSubCode(value string) {
	w.WithdrawalReasonSubCode = (*Max4Text)(&value)
}

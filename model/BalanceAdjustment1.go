package model

// Specifies the balance adjustments for a specific service.
type BalanceAdjustment1 struct {

	// Identifies the type of adjustment.
	Type *BalanceAdjustmentType1Code `xml:"Tp"`

	// Free-form description and clarification of the adjustment.
	Description *Max105Text `xml:"Desc"`

	// Amount of the adjustment. If the amount would reduce the underlying balance then the amount should be negatively signed. Expressed in the Account currency.
	BalanceAmount *AmountAndDirection34 `xml:"BalAmt"`

	// Day-weighted net amount of the adjustment to the average collected balance over the statement period.
	AverageAmount *AmountAndDirection34 `xml:"AvrgAmt,omitempty"`

	// Date on which the error occurred in the underlying cash account.
	ErrorDate *ISODate `xml:"ErrDt,omitempty"`

	// Date on which the error was corrected in the cash account. If the date is not know then use the last day of the month in which the error was corrected.
	PostingDate *ISODate `xml:"PstngDt"`

	// Number of days within the period to which the adjustment applies.
	Days *DecimalNumber `xml:"Days,omitempty"`

	// Earnings credit adjustment, debit or credit, resulting from this adjustment’s effect on the average collected balance. If the amount would reduce the credit due then the amount should be negatively signed.
	EarningsAdjustmentAmount *AmountAndDirection34 `xml:"EarngsAdjstmntAmt,omitempty"`
}

func (b *BalanceAdjustment1) SetType(value string) {
	b.Type = (*BalanceAdjustmentType1Code)(&value)
}

func (b *BalanceAdjustment1) SetDescription(value string) {
	b.Description = (*Max105Text)(&value)
}

func (b *BalanceAdjustment1) AddBalanceAmount() *AmountAndDirection34 {
	b.BalanceAmount = new(AmountAndDirection34)
	return b.BalanceAmount
}

func (b *BalanceAdjustment1) AddAverageAmount() *AmountAndDirection34 {
	b.AverageAmount = new(AmountAndDirection34)
	return b.AverageAmount
}

func (b *BalanceAdjustment1) SetErrorDate(value string) {
	b.ErrorDate = (*ISODate)(&value)
}

func (b *BalanceAdjustment1) SetPostingDate(value string) {
	b.PostingDate = (*ISODate)(&value)
}

func (b *BalanceAdjustment1) SetDays(value string) {
	b.Days = (*DecimalNumber)(&value)
}

func (b *BalanceAdjustment1) AddEarningsAdjustmentAmount() *AmountAndDirection34 {
	b.EarningsAdjustmentAmount = new(AmountAndDirection34)
	return b.EarningsAdjustmentAmount
}

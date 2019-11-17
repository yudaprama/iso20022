package model

// Choice of formats to  express the lottery type.
type LotteryType1FormatChoice struct {

	// Standard code to  specify the lottery type.
	Code *LotteryType1Code `xml:"Cd"`

	// Proprietary code to  express the lottery type.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (l *LotteryType1FormatChoice) SetCode(value string) {
	l.Code = (*LotteryType1Code)(&value)
}

func (l *LotteryType1FormatChoice) AddProprietary() *GenericIdentification13 {
	l.Proprietary = new(GenericIdentification13)
	return l.Proprietary
}

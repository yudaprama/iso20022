package model

// Choice between a standard code or proprietary code to specify the type of lottery.
type LotteryTypeFormat4Choice struct {

	// Standard code to specify the type of lottery announced.
	Code *LotteryType1Code `xml:"Cd"`

	// Proprietary identification of the type of lottery announced.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (l *LotteryTypeFormat4Choice) SetCode(value string) {
	l.Code = (*LotteryType1Code)(&value)
}

func (l *LotteryTypeFormat4Choice) AddProprietary() *GenericIdentification30 {
	l.Proprietary = new(GenericIdentification30)
	return l.Proprietary
}

package model

// Choice between a standard code or proprietary code to specify the type of lottery.
type LotteryTypeFormat5Choice struct {

	// Standard code to specify the type of lottery announced.
	Code *LotteryType1Code `xml:"Cd"`

	// Proprietary identification of the type of lottery announced.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (l *LotteryTypeFormat5Choice) SetCode(value string) {
	l.Code = (*LotteryType1Code)(&value)
}

func (l *LotteryTypeFormat5Choice) AddProprietary() *GenericIdentification47 {
	l.Proprietary = new(GenericIdentification47)
	return l.Proprietary
}

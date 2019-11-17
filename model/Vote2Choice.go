package model

// Determines how the voting instructions are specified.
type Vote2Choice struct {

	// Instruction specifying the instructed quantity of voting rights per resolution. Split votes can be indicated. If only one type of decision is indicated, the number of votes cast must not be adjusted if the position of the voting party increases.
	VoteInstruction []*Vote4 `xml:"VoteInstr"`

	// Instruction specifiying a vote instruction per resolution for the entire entitlement.
	GlobalVoteInstruction []*Vote3 `xml:"GblVoteInstr"`
}

func (v *Vote2Choice) AddVoteInstruction() *Vote4 {
	newValue := new(Vote4)
	v.VoteInstruction = append(v.VoteInstruction, newValue)
	return newValue
}

func (v *Vote2Choice) AddGlobalVoteInstruction() *Vote3 {
	newValue := new(Vote3)
	v.GlobalVoteInstruction = append(v.GlobalVoteInstruction, newValue)
	return newValue
}

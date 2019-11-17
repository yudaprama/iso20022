package model

// Determines how the voting instructions are specified.
type Vote3Choice struct {

	// Instruction specifying the instructed quantity of voting rights per resolution. Split votes can be indicated. If only one type of decision is indicated, the number of votes cast must not be adjusted if the position of the voting party increases.
	VoteInstruction []*Vote8 `xml:"VoteInstr"`

	// Instruction specifying a vote instruction per resolution for the entire entitlement.
	GlobalVoteInstruction []*Vote9 `xml:"GblVoteInstr"`
}

func (v *Vote3Choice) AddVoteInstruction() *Vote8 {
	newValue := new(Vote8)
	v.VoteInstruction = append(v.VoteInstruction, newValue)
	return newValue
}

func (v *Vote3Choice) AddGlobalVoteInstruction() *Vote9 {
	newValue := new(Vote9)
	v.GlobalVoteInstruction = append(v.GlobalVoteInstruction, newValue)
	return newValue
}

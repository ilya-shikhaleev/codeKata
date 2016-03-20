package battle

type WantedPlayer struct {
	damage float64
	name   string
	gender Gender
}

type CasePlayer struct {
	in   Character
	want WantedPlayer
}

type CaseFight struct {
	in   [2]Character
	want bool
}

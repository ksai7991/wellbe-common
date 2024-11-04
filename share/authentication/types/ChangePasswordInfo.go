package types

type ChangePasswordInfo struct {
	PreviousPassword string
	ProposedPassword string
	AccessToken      string
}
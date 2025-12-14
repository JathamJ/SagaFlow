package core

type Role struct {
	Name             string
	VoicePrompt      string
	AppearancePrompt string // Appearance prompt ensure consistency in generating characters across multiple paragraphs
}

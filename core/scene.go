package core

type Scene struct {
	BackGroundPrompt string
	Prompt           string
	Roles            []*Role
	Storyboard       []*StoryBoard
	Next             []*Scene
}

package core

import "time"

type Drama struct {
	Title          string
	Style          string
	Language       string
	Prompt         string
	SceneTimeLimit time.Duration
	Roles          map[string]*Role
}

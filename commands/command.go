package commands

import (
	"github.com/uhppoted/uhppote-cli/config"
	"github.com/uhppoted/uhppote-core/uhppote"
)

type Context struct {
	uhppote *uhppote.UHPPOTE
	config  *config.Config
}

type Command interface {
	Execute(context Context) error
	CLI() string
	Description() string
	Usage() string
	Help()
}

func NewContext(u *uhppote.UHPPOTE, c *config.Config) Context {
	return Context{u, c}
}

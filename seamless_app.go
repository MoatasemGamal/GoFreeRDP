package gofreerdp

import (
	"fmt"
	"strings"
)

type SeamlessApp struct {
	program string
	cmd     string
	file    string
	guid    string
	icon    string
	name    string
	workdir string
}

func (app *SeamlessApp) toString() string {
	var parts []string

	// Check each field and append to parts if non-empty
	if app.program != "" {
		parts = append(parts, fmt.Sprintf("program:'%s'", app.program))
	}
	if app.cmd != "" {
		parts = append(parts, fmt.Sprintf("cmd:'%s'", app.cmd))
	}
	if app.file != "" {
		parts = append(parts, fmt.Sprintf("file:'%s'", app.file))
	}
	if app.guid != "" {
		parts = append(parts, fmt.Sprintf("guid:'%s'", app.guid))
	}
	if app.icon != "" {
		parts = append(parts, fmt.Sprintf("icon:'%s'", app.icon))
	}
	if app.name != "" {
		parts = append(parts, fmt.Sprintf("name:'%s'", app.name))
	}
	if app.workdir != "" {
		parts = append(parts, fmt.Sprintf("workdir:'%s'", app.workdir))
	}

	// Join all parts into a single string, separated by commas
	return strings.Join(parts, ",")
}

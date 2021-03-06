// NB this code was based on https://github.com/hashicorp/packer/blob/370b67497e90785b71be1b6fcc6430de487d644e/provisioner/powershell/elevated.go

package update

import (
	"text/template"
)

type elevatedOptions struct {
	Username        string
	Password        string
	TaskName        string
	TaskDescription string
	Command         string
}

var elevatedTemplate = template.Must(
	template.New("Elevated").Parse(
		string(MustAsset("elevated-template.ps1"))))

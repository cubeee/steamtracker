package engine

import (
	"github.com/flosch/pongo2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func RegisterFilters() {
	pongo2.RegisterFilter("formatnumber", filterFormatNumber)
}

func filterFormatNumber(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	p := message.NewPrinter(language.English)
	return pongo2.AsValue(p.Sprintf(param.String(), in.Interface())), nil
}

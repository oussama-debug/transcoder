package sections

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

func CreateHelperSection() *tview.TextView {
	maxLength := 0
	var sb strings.Builder

	header := tview.NewTextView().SetDynamicColors(true).SetWordWrap(true)

	values := map[string]string{
		"version": "0.0.1",
		"author":  "Oussama Jaaouani <ojaaouani@x2omedia.com>",
	}

	keys := []string{"version", "author"}
	for key := range values {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}

	for _, key := range keys {
		value := values[key]
		padding := maxLength - len(key)
		fmt.Fprintf(&sb, "[orange]%s:[white]%*s%s\n", key, padding+1, "", value)
	}

	header.SetText(sb.String())
	return header
}

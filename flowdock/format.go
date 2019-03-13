package flowdock

import (
	"fmt"
	"org2fd/orgmode"
	"strings"
)

// Format formats the OrgMode AST as Flowdock notes
func Format(orgMode *orgmode.OrgMode) string {
	var builder strings.Builder

	for _, headline := range orgMode.Headlines {
		builder.WriteString(fmt.Sprintf("- %s\n", *headline.Line.String))
	}

	return builder.String()
}

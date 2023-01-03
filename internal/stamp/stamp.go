package stamp

import (
	"github.com/bmf-san/akashigo/types"
)

var typeStampTextAliasMap = map[string]string{
	"hello": types.TextAttendance,
	"bye":   types.TextLeaveWork,
}

// GetTypeStampByAliasMap gets type stamp text using alias map.
// If the given text does not exist in the alias map, the given text will be returned without any changes.
func GetTypeStampTextByAlias(alias string) string {
	typeStampText, ok := typeStampTextAliasMap[alias]
	if !ok {
		return alias
	}
	return typeStampText
}

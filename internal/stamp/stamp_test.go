package stamp

import (
	"fmt"
	"testing"

	"github.com/bmf-san/akashigo/types"
)

func Test_getTypeStampTextByAlias(t *testing.T) {
	tests := []struct {
		alias string
		want  string
	}{
		{
			alias: "hello",
			want:  types.TextAttendance,
		},
		{
			alias: "bye",
			want:  types.TextLeaveWork,
		},
		{
			alias: "non existent",
			want:  "non existent",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v - %v", tt.want, tt.alias), func(t *testing.T) {
			if got := GetTypeStampTextByAlias(tt.alias); got != tt.want {
				t.Errorf("GetTypeStampTextByAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

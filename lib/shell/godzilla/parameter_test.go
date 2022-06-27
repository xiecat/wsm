package godzilla

import (
	"reflect"
	"testing"
)

func TestSplitArgs(t *testing.T) {
	type args struct {
		input                    string
		maxParts                 int
		removeAllEscapeSequences bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "base 测试",
			args: args{
				input:                    `cmd /c "cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami"`,
				maxParts:                 100000,
				removeAllEscapeSequences: false,
			},
			want: []string{"cmd", "/c", `cd /d "D:/Jdk/apache-tomcat-7.0.109/bin/"&whoami`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitArgs(tt.args.input, tt.args.maxParts, tt.args.removeAllEscapeSequences); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

package utils

import (
	"reflect"
	"testing"
)

func TestSecretKey(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "生成符合标准的秘钥",
			args: args{pwd: "rebeyond"},
			want: []byte("e45e329feb5d925b"),
		},
		{
			name: "空值",
			args: args{pwd: ""},
			want: []byte("d41d8cd98f00b204"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecretKey(tt.args.pwd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("secretKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

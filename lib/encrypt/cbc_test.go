package encrypt

import (
	"reflect"
	"testing"
)

func TestAESCBCDecrypt(t *testing.T) {
	type args struct {
		src []byte
		key []byte
		iv  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test empty",
			args: args{
				src: []byte(""),
				key: []byte("0123456789abcdef"),
				iv:  []byte("0123456789abcdef"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AESCBCDecrypt(tt.args.src, tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("AESCBCDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AESCBCDecrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

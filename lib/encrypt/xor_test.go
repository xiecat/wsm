package encrypt

import (
	"reflect"
	"testing"
)

func TestXor(t *testing.T) {
	type args struct {
		content []byte
		key     []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Xor 加密完再还原",
			args: args{
				content: []byte("SJKFDHKBERUYBXBVOPMFNSLFDLIOERHKBBZKJXVHGFKWKHKJABDFASBDFK"),
				key:     []byte("0123456789abcdef"),
			},
			want: []byte("SJKFDHKBERUYBXBVOPMFNSLFDLIOERHKBBZKJXVHGFKWKHKJABDFASBDFK"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Xor 加密
			got1 := Xor(tt.args.content, tt.args.key)
			// 还原
			if got := Xor(got1, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}

package encrypt

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

func TestXor(t *testing.T) {
	src, _ := base64.StdEncoding.DecodeString("AkQCHVA6UzljMVZWEhsFE0FyXx8oXAoWAkFUUVpRTEcMWwZRFhVWF1MfBAILGwNaDRlHFhVQDlgOWBUdX1EVWwxSK1EPXWM5YzE1UV9QLVoNU2cLYjhhWg5VFR1RFENQBxZKVEIaJQNMe1FZHVURUgBeAB0WVwxaAkUYBRwETwJTD0pSC1ZOG0VGXV1TWQgTTBQEQgUVUTtgMTUyUVkFUhFRSAFgOmE5Yx5WU0BTEnAMQwtEYDlhOWMCUEpXVxRHAlQJVSRRDVxhMjUyMlcMVwZOAFMXTABbD1R0QFVHYwRjNmUfARhDWgcRGlYSFiUJTHwBW01ZEVgAWVAfRlsMUAJCSAdMCE8IUwgaUFtaThFFQQ1fA1UIGUwTW1dZYSR3W2UvSQxrKwgwcmRBXm0NQwgDN3g2dlhsWnQCcHlnLAEERi14LE0ZcCtLcXUFVgZcM041SiB5VldRBVJ6MHZhM2NYAFs3fSUBMHtMXGF+UGAgZxZcO1QRUlZjfWZ8DTQKJgEnezF1U14TeX18R0woexlyIgcAXw5pG2FPcHMDDwFXUS0=")

	fmt.Println(string(Xor(src, []byte("3c6e0b8a9c15224a"))))

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

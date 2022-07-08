package wsm

import (
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/behinder"
	"testing"
)

const (
	// JspIndexShellUrl 测试输出存在干扰字符的情况
	JspIndexShellUrl = "http://139.196.175.86:8080/bxindex.jsp"
	JspShellUrl      = "http://139.196.175.86:8080/bx.jsp"
	CsharpShellUrl   = "http://139.196.175.86:8081/bx.aspx"
	AspShellUrl      = "http://139.196.175.86:8081/bx.asp"
	PhpShellUrl      = "http://139.196.175.86/bx.php"
)

func TestBehinderInfo_Ping(t *testing.T) {
	type fields struct {
		BaseShell   BaseShell
		secretKey   []byte
		encryptMode int
		prefixLen   int
		suffixLen   int
	}
	type args struct {
		p []shell.IParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "冰蝎 JSP 结果不加密测试",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspShellUrl,
					Password: "rebeyond",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
					Headers:  nil,
				},
			},
			args: args{
				p: []shell.IParams{
					&behinder.PingParams{
						// response 结果不加密测试
						OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: false, NotEncrypt: true},
						Content:        "xxxxxxx",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "冰蝎 JSP 结果存在干扰字符测试",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspIndexShellUrl,
					Password: "rebeyond",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
					Headers:  nil,
				},
			},
			args: args{
				p: []shell.IParams{
					&behinder.PingParams{
						// out.println("干扰字符") 结果不加密测试
						OnlyJavaParams: behinder.OnlyJavaParams{ForcePrint: true, NotEncrypt: false},
						Content:        "xxxxxxx",
					},
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "冰蝎 JSP jdk 1.8.0_181",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspShellUrl,
					Password: "rebeyond",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
					Headers:  nil,
				},
			},
			args:    args{nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "冰蝎 ASPX shell .NET CLR v4.0 环境下",
			fields: fields{
				BaseShell: BaseShell{
					Url:      CsharpShellUrl,
					Password: "rebeyond",
					Script:   shell.CsharpScript,
					Proxy:    "http://127.0.0.1:9999",
					Headers:  nil,
				},
			},
			args:    args{nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "冰蝎 ASP",
			fields: fields{
				BaseShell: BaseShell{
					Url:      AspShellUrl,
					Password: "rebeyond",
					Script:   shell.AspScript,
					Proxy:    "http://127.0.0.1:9999",
					Headers:  nil,
				},
			},
			args:    args{nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "冰蝎 PHP php_5.4.45",
			fields: fields{
				BaseShell: BaseShell{
					Url:      PhpShellUrl,
					Password: "rebeyond",
					Script:   shell.PhpScript,
					Proxy:    "http://127.0.0.1:9999",
					Headers:  nil,
				},
			},
			args:    args{nil},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BehinderInfo{
				BaseShell:   tt.fields.BaseShell,
				secretKey:   tt.fields.secretKey,
				encryptMode: tt.fields.encryptMode,
				prefixLen:   tt.fields.prefixLen,
				suffixLen:   tt.fields.suffixLen,
			}
			b, err := NewBehinder(b)
			if err != nil {
				t.Errorf("NewBehinder(b) error = %v, wantErr %v", err, nil)
				return
			}
			got, err := b.Ping(tt.args.p...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
		})
	}
}

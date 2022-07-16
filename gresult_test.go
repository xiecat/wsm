package wsm

import (
	"fmt"
	"testing"
)

func Test_parserDatabaseOptToMap(t *testing.T) {
	raw := ` ok
U0NIRU1BX05BTUU=
aW5mb3JtYXRpb25fc2NoZW1h
Z29kemlsbGE=
bXlzcWw=
cGVyZm9ybWFuY2Vfc2NoZW1h
c3lz
`
	_ = raw
	raw2 := `ok
SWQ=	c3RyaW5n	cmF3	
MQ==	1eLKx9K7zPWy4srUyv2+3S0tLS0tZ29kemlsbGE=	MTIzMzQ1NDM1NDM1NA==	
Mg==	dHR0dHR0dHR0dHR0	YmVoaW5kZXI=	
`
	_ = raw2
	a, _ := parserDatabaseOptToMap(raw)
	fmt.Printf("%#+v\n", a)
	//type args struct {
	//	raw string
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    map[string]string
	//	wantErr bool
	//}{
	//	{
	//		name: "godzilla 数据库管理回显处理",
	//		args: args{raw: raw},
	//		want: map[string]string{"": ""},
	//		wantErr: false,
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := parserDatabaseOptToMap(tt.args.raw)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("parserDatabaseOptToMap() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("parserDatabaseOptToMap() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

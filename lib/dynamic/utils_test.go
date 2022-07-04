package dynamic

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestInStrSlice(t *testing.T) {
	type args struct {
		array []string
		str   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "存在",
			args: args{
				array: []string{"1", "2", "3"},
				str:   "2",
			},
			want: true,
		},
		{
			name: "不存在",
			args: args{
				array: []string{"1", "2", "3"},
				str:   "4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InStrSlice(tt.args.array, tt.args.str); got != tt.want {
				t.Errorf("InStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToBytes(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "int 转 bytes",
			args: args{value: 30},
			want: []byte{30, 0, 0, 0},
		},
		{
			name: "int 转 bytes",
			args: args{value: 256},
			want: []byte{0, 1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToBytes(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeBytes(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeBytes(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPrefixLenAndSuffixLen(t *testing.T) {
	src, _ := base64.StdEncoding.DecodeString("bUFVWUx6bXFuNVFQRGt5STVsdlNwMGZqaUJ1MWU3MDQ3WWpmY3p3WTZqNUduaVdTNUwxOHRNMlhWSmxEQTM5aFAySEZLZkRnelJ0cTJWaXpHQU1UaHUrNEpueWhJdW1aa3FiTktvMEQzMUJodHlWYVhXcERYQ2NGZGMyU3g5ZUQ0SEoyc3FPbUp5N0xrNHRYdjE2OW1uV0c5azQ2RHRzNzBZQW1vVlgvZGlYeGU1ZktSakNSdzE5QmlNbk9STThvTUVYeXhtc0k5VGpxZVBrMzBPMmgxeVZtK0pVbjdab1l2UFA5d3dSdEt1Q3lKd2dGb2dISk1TbTBWSHhnb1B5cGEwUkV1cmE3MTZUNml1SXpPZmw2dmc2RFRHU1o3dUEvdVlDZHBkNmYvWjQ9")
	sub1, _ := base64.StdEncoding.DecodeString("SXo0TURRZnlqWkdhV1hRajBrQVVldVR3RWQyMEJDK0pJWGpoandPblJNMWFJNGJDUDhDVkJDd05EVDRPeUVadXFPWVBCQVJMeVlIRGRodTdEb2VIRE1IOHk3cGI3Z2k3N2FIdlBtTDlPQnR2aW00RzMwazNScTA5OVJLVUsxRkVMRGtsZENZYWgzZGgzWS9NbVJ5SGdPLy9zaC81YVNaMzdXenFnMElmaW13UzIzdWJmTFhnMU10Y0VwRjdRSmEvQXlpR2ZOclQ3QjN6UVRQTy9JVUdvS3J2WFYwODlzN04xRUdySWpmL2VOaVpTVHdLbnNPbkt2Mm54MGdKTXVHZkNDNjZ2Z3FxYmUweTQ3QWJ4eVcrTXFYNk9iTk0vYUxuTzZyT1ZtSi92b1k9")
	sub2, _ := base64.StdEncoding.DecodeString("bUFVWUx6bXFuNVFQRGt5STVsdlNwMGZqaUJ1MWU3MDQ3WWpmY3p3WTZqNUduaVdTNUwxOHRNMlhWSmxEQTM5aFAySEZLZkRnelJ0cTJWaXpHQU1UaHUrNEpueWhJdW1aa3FiTktvMEQzMUJodHlWYVhXcERYQ2NGZGMyU3g5ZUQ0SEoyc3FPbUp5N0xrNHRYdjE2OW1uV0c5azQ2RHRzNzBZQW1vVlgvZGlYeGU1ZktSakNSdzE5QmlNbk9STThvTUVYeXhtc0k5VGpxZVBrMzBPMmgxeVZtK0pVbjdab1l2UFA5d3dSdEt1Q3lKd2dGb2dISk1TbTBWSHhnb1B5cGEwUkV1cmE3MTZUNml1SXpPZmw2dmc2RFRHU1o3dUEvdVlDZHBkNmYvWjQ9")

	type args struct {
		src    []byte
		substr [][]byte
	}
	tests := []struct {
		name         string
		args         args
		wantIndex    int
		wantEndIndex int
	}{
		{
			name: "test success",
			args: args{
				src: src,
				substr: [][]byte{
					sub1,
					sub2,
				},
			},
			wantIndex:    0,
			wantEndIndex: 0,
		},
		{
			name: "test success",
			args: args{
				src: []byte("aaaaaaaaaa"),
				substr: [][]byte{
					[]byte("bbbbbbbbbb"),
					[]byte("aaaaaaaaaa"),
				},
			},
			wantIndex:    0,
			wantEndIndex: 0,
		},
		{
			name: "test success",
			args: args{
				src: []byte("aaaaaaaaaa"),
				substr: [][]byte{
					[]byte("bbbbbbbbbb"),
					[]byte("cccccccccc"),
				},
			},
			wantIndex:    -1,
			wantEndIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotEndIndex := GetPrefixLenAndSuffixLen(tt.args.src, tt.args.substr...)
			if gotIndex != tt.wantIndex {
				t.Errorf("GetPrefixLenAndSuffixLen() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotEndIndex != tt.wantEndIndex {
				t.Errorf("GetPrefixLenAndSuffixLen() gotEndIndex = %v, want %v", gotEndIndex, tt.wantEndIndex)
			}
		})
	}
}

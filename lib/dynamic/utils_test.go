package dynamic

import (
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

func TestGetIndexAndLastIndex(t *testing.T) {
	type args struct {
		src    []byte
		substr []byte
	}
	tests := []struct {
		name          string
		args          args
		wantIndex     int
		wantLastIndex int
	}{
		{
			name: "base test",
			args: args{
				src:    []byte("AKJSHDFKdsjhfbahsbfjhysdfjadsbf"),
				substr: []byte("SHDFKdsjhfbahsbfjhysdfj"),
			},
			wantIndex:     3,
			wantLastIndex: 3 + len("SHDFKdsjhfbahsbfjhysdfj"),
		},
		{
			name: "base test2",
			args: args{
				src:    []byte("AKJSHSHDFKds!@#$%^@C≈ç√∫∂jhfbahsbfjhysdfjysdfjadsbf"),
				substr: []byte("SHDFKds!@#$%^@C≈ç√∫∂jhfbahsbfjhysdfj"),
			},
			wantIndex:     5,
			wantLastIndex: 5 + len("SHDFKds!@#$%^@C≈ç√∫∂jhfbahsbfjhysdfj"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotLastIndex := GetIndexAndLastIndex(tt.args.src, tt.args.substr)
			if gotIndex != tt.wantIndex {
				t.Errorf("GetIndexAndLastIndex() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotLastIndex != tt.wantLastIndex {
				t.Errorf("GetIndexAndLastIndex() gotLastIndex = %v, want %v", gotLastIndex, tt.wantLastIndex)
			}
		})
	}
}

package helpers

import (
	"reflect"
	"testing"
)

func TestGetQueryParams(t *testing.T) {
	type args struct {
		offset string
		limit  string
	}
	tests := []struct {
		name    string
		args    args
		want    *QueryParams
		wantErr bool
	}{
		{
			name: "return {offset: 10, limit: 100}",
			args: args{
				offset: "10",
				limit:  "100",
			},
			want: &QueryParams{
				Offset: 10,
				Limit:  100,
			},
			wantErr: false,
		},
		{
			name: "return error",
			args: args{
				offset: "asdasdasd",
				limit:  "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "return default limit",
			args: args{
				offset: "10",
				limit:  "",
			},
			want: &QueryParams{
				Offset: 10,
				Limit:  50,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetQueryParams(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueryParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueryParams() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return abcd",
			args: args{
				word: "bcad",
			},
			want: "abcd",
		},
		{
			name: "return elov",
			args: args{
				word: "love",
			},
			want: "elov",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortWord(tt.args.word); got != tt.want {
				t.Errorf("SortWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

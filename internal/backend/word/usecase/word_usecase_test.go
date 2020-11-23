package usecase

import (
	"github.com/Elephmoon/anagramDictionary/internal/models"
	"reflect"
	"testing"
)

func Test_generateAnagramAnswer(t *testing.T) {
	type args struct {
		searchWord string
		words      []*models.Word
	}
	tests := []struct {
		name string
		args args
		want models.AnagramResponse
	}{
		{
			name: "test generate AnagramResponse",
			args: args{
				searchWord: "love",
				words: []*models.Word{
					{
						Word: "test",
					},
					{
						Word: "afasfasf",
					},
				},
			},
			want: models.AnagramResponse{
				Word:     "love",
				Anagrams: []string{"test", "afasfasf"},
			},
		},
		{
			name: "test generate AnagramResponse with empty words",
			args: args{
				searchWord: "test",
				words:      []*models.Word{},
			},
			want: models.AnagramResponse{
				Word:     "test",
				Anagrams: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateAnagramResponse(tt.args.searchWord, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateAnagramResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateCreateReq(t *testing.T) {
	type args struct {
		words *models.CreateReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no error",
			args:    args{words: &models.CreateReq{Words: []string{"asdasdasdasd"}}},
			wantErr: false,
		},
		{
			name:    "return validator err",
			args:    args{&models.CreateReq{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateReq(tt.args.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCreateReq() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

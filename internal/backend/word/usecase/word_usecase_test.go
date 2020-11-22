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

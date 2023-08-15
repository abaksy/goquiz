package quiz

import (
	"reflect"
	"testing"
)

func TestParseQuestions(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Read CSV file",
			args: args{fileName: "../tests/test1.csv"},
			want: [][]string{{"5+5", "10"}, {"7+3", "10"}, {"1+1", "2"}, {"8+3", "11"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseQuestions(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}

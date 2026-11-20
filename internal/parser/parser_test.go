package parser

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []int
		wantErr bool
	}{
		{
			name: "valid numbers",
			args: []string{"1", "2", "3"},
			want: []int{1, 2, 3},
		},
		{
			name: "space separated numbers",
			args: []string{"1 2 3"},
			want: []int{1, 2, 3},
		},
		{
			name: "negative numbers",
			args: []string{"-5", "0", "8"},
			want: []int{-5, 0, 8},
		},
		{
			name: "mixed arguments",
			args: []string{"1 2", "3 4"},
			want: []int{1, 2, 3, 4},
		},
		{
			name:    "invalid text",
			args:    []string{"abc"},
			wantErr: true,
		},
		{
			name:    "mixed valid and invalid",
			args:    []string{"1", "abc", "3"},
			wantErr: true,
		},
		{
			name:    "duplicate numbers",
			args:    []string{"1", "2", "1"},
			wantErr: true,
		},
		{
			name:    "empty input",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "integer overflow",
			args:    []string{"999999999999999999999999999999"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ParseArgs(test.args)

			if test.wantErr {
				if err == nil {
					t.Fatal("expected an error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("expected %v, got %v", test.want, got)
			}
		})
	}
}

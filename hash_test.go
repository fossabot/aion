package aion

import "testing"

func Test_hash(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "empty string",
			args: args{input: ""},
			want: 14695981039346656037,
		},
		{
			name: "short string",
			args: args{input: "hello world"},
			want: 9065573210506989167,
		},
		{
			name: "long string",
			args: args{input: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."},
			want: 6310023833031058075,
		},
		{
			name: "number as string",
			args: args{input: "1"},
			want: 12638153115695167470,
		},
		{
			name: "Special characters",
			args: args{input: "@!#+*?="},
			want: 14488640670009027472,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.input); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

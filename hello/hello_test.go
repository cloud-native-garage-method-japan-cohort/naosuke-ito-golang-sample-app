package hello

import "testing"

func TestSayHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "正常系",
			want: "こんにちは、世界",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayHello(); got != tt.want {
				t.Errorf("SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

package decorator

import "testing"

func TestColorSquare_Draw(t *testing.T) {
	type fields struct {
		square IDraw
		color  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case", fields{color: "test", square: Square{}}, "this is a square, color is test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ColorSquare{
				square: tt.fields.square,
				color:  tt.fields.color,
			}
			if got := c.Draw(); got != tt.want {
				t.Errorf("Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}

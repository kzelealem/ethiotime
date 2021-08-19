package ethiopiandate

import (
	"testing"
	"time"
)

func TestTime_Format(t *testing.T) {
	tt := time.Now().Add(time.Hour * -12)
	ett := Date(tt)
	et := Now()
	type args struct {
		layout string
	}
	tests := []struct {
		name string
		tr   Time
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test", et, args{"January 02 pm 03:04"}, ""},
		{"test", ett, args{"January 02 PM 03:04"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Format(tt.args.layout); got != tt.want {
				t.Errorf("Time.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

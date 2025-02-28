package ethiopiandate

import (
	"testing"
	"time"
)

func TestTime_Format(t *testing.T) {
	testTime := time.Date(2024, 1, 1, 14, 30, 0, 0, time.UTC)
	ethTime := Date(testTime)

	type args struct {
		layout string
	}
	tests := []struct {
		name string
		tr   Time
		args args
		want string
	}{
		{"full date and time", ethTime, args{"January 02, 2006 03:04 PM"}, "ታኅሣሥ 22, 2016 08:30 ከሰአት"},
		{"date only", ethTime, args{"January 02, 2006"}, "ታኅሣሥ 22, 2016"},
		{"time only", ethTime, args{"03:04 PM"}, "08:30 ከሰአት"},
		{"weekday and date", ethTime, args{"Monday, January 02"}, "ሰኞ, ታኅሣሥ 22"},
		{"full format with weekday", ethTime, args{"Monday, January 02, 2006 03:04 PM"}, "ሰኞ, ታኅሣሥ 22, 2016 08:30 ከሰአት"},
		{"numeric date", ethTime, args{"01/02/2006"}, "04/22/2016"},
		{"24-hour time", ethTime, args{"15:04"}, "08:30"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Format(tt.args.layout); got != tt.want {
				t.Errorf("Time.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

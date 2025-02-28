package ethiopiandate

import (
	"testing"
	"time"
)

func TestTime_Clock(t *testing.T) {
	tests := []struct {
		name   string
		tr     Time
		wantHh int
		wantMm int
		wantSs int
	}{
		{
			name:   "morning conversion",
			tr:     Date(time.Date(2024, 1, 1, 8, 30, 15, 0, time.UTC)),
			wantHh: 2,
			wantMm: 30,
			wantSs: 15,
		},
		{
			name:   "afternoon conversion",
			tr:     Date(time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC)),
			wantHh: 8,
			wantMm: 0,
			wantSs: 0,
		},
		{
			name:   "midnight conversion",
			tr:     Date(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			wantHh: 18,
			wantMm: 0,
			wantSs: 0,
		},
		{
			name:   "pre-midnight",
			tr:     Date(time.Date(2024, 1, 1, 23, 59, 59, 0, time.UTC)),
			wantHh: 17,
			wantMm: 59,
			wantSs: 59,
		},
		{
			name:   "early morning",
			tr:     Date(time.Date(2024, 1, 1, 6, 0, 0, 0, time.UTC)),
			wantHh: 0,
			wantMm: 0,
			wantSs: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHh, gotMm, gotSs := tt.tr.Clock()
			if gotHh != tt.wantHh {
				t.Errorf("Clock() hour = %v, want %v", gotHh, tt.wantHh)
			}
			if gotMm != tt.wantMm {
				t.Errorf("Clock() minute = %v, want %v", gotMm, tt.wantMm)
			}
			if gotSs != tt.wantSs {
				t.Errorf("Clock() second = %v, want %v", gotSs, tt.wantSs)
			}
		})
	}
}

func TestTime_clock(t *testing.T) {
	tests := []struct {
		name   string
		tr     Time
		wantHh int
		wantMm int
		wantSs int
	}{
		{
			name:   "standard morning",
			tr:     Date(time.Date(2024, 1, 1, 10, 15, 30, 0, time.UTC)),
			wantHh: 4,
			wantMm: 15,
			wantSs: 30,
		},
		{
			name:   "around midnight",
			tr:     Date(time.Date(2024, 1, 1, 5, 45, 0, 0, time.UTC)),
			wantHh: 23,
			wantMm: 45,
			wantSs: 0,
		},
		{
			name:   "noon time",
			tr:     Date(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)),
			wantHh: 6,
			wantMm: 0,
			wantSs: 0,
		},
		{
			name:   "late evening",
			tr:     Date(time.Date(2024, 1, 1, 20, 30, 45, 0, time.UTC)),
			wantHh: 14,
			wantMm: 30,
			wantSs: 45,
		},
		{
			name:   "exact 6 AM",
			tr:     Date(time.Date(2024, 1, 1, 6, 0, 0, 0, time.UTC)),
			wantHh: 0,
			wantMm: 0,
			wantSs: 0,
		},
		{
			name:   "pre 6 AM",
			tr:     Date(time.Date(2024, 1, 1, 5, 59, 59, 0, time.UTC)),
			wantHh: 23,
			wantMm: 59,
			wantSs: 59,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHh, gotMm, gotSs := tt.tr.clock()
			if gotHh != tt.wantHh {
				t.Errorf("clock() hour = %v, want %v", gotHh, tt.wantHh)
			}
			if gotMm != tt.wantMm {
				t.Errorf("clock() minute = %v, want %v", gotMm, tt.wantMm)
			}
			if gotSs != tt.wantSs {
				t.Errorf("clock() second = %v, want %v", gotSs, tt.wantSs)
			}
		})
	}
}

func TestTime_String(t *testing.T) {
	tests := []struct {
		name string
		tr   Time
		want string
	}{
		{
			name: "morning time",
			tr:   Date(time.Date(2024, 1, 1, 8, 30, 0, 0, time.UTC)),
			want: "ታኅሣሥ 22, 2016 02:30",
		},
		{
			name: "afternoon time",
			tr:   Date(time.Date(2024, 1, 1, 14, 45, 0, 0, time.UTC)),
			want: "ታኅሣሥ 22, 2016 08:45",
		},
		{
			name: "midnight",
			tr:   Date(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)),
			want: "ታኅሣሥ 22, 2016 18:00",
		},
		{
			name: "pre-midnight",
			tr:   Date(time.Date(2024, 1, 1, 23, 59, 0, 0, time.UTC)),
			want: "ታኅሣሥ 22, 2016 17:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

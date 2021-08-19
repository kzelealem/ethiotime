package ethiopiandate

import (
	"testing"
	"time"
)

func TestTime_Clock(t *testing.T) {
	tt := time.Now().Add(-8 * time.Hour)
	tr := Date(tt)
	tests := []struct {
		name   string
		tr     Time
		wantHh int
		wantMm int
		wantSs int
	}{
		// TODO: Add test cases.
		{"test", tr, 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHh, gotMm, gotSs := tt.tr.Clock()
			if gotHh != tt.wantHh {
				t.Errorf("Time.Clock() gotHh = %v, want %v", gotHh, tt.wantHh)
			}
			if gotMm != tt.wantMm {
				t.Errorf("Time.Clock() gotMm = %v, want %v", gotMm, tt.wantMm)
			}
			if gotSs != tt.wantSs {
				t.Errorf("Time.Clock() gotSs = %v, want %v", gotSs, tt.wantSs)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHh, gotMm, gotSs := tt.tr.clock()
			if gotHh != tt.wantHh {
				t.Errorf("Time.clock() gotHh = %v, want %v", gotHh, tt.wantHh)
			}
			if gotMm != tt.wantMm {
				t.Errorf("Time.clock() gotMm = %v, want %v", gotMm, tt.wantMm)
			}
			if gotSs != tt.wantSs {
				t.Errorf("Time.clock() gotSs = %v, want %v", gotSs, tt.wantSs)
			}
		})
	}
}

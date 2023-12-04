package main

import "testing"

func TestCalc(t *testing.T) {
	tests := map[string]struct {
		arg  uint64
		want string
	}{
		"2023 = 卯年": {2023, "卯年"},
		"2024 = 辰年": {2024, "辰年"},
		"2025 = 巳年": {2025, "巳年"},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Calc(tt.arg); got != tt.want {
				t.Errorf("Calc(%d) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

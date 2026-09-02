package qimen

import "testing"

func TestSmoke(t *testing.T) {
	r, e := Calculate(map[string]any{"year": 2026, "month": 4, "day": 10, "hour": 14, "minute": 30, "timezone": "Asia/Shanghai"})
	if e != nil {
		t.Fatal(e)
	}
	if len(r["palaces"].([]any)) != 9 {
		t.Fatal(r)
	}
}

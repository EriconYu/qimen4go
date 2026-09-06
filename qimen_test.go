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
	if r["algorithmVersion"] != "qimen-zhuanpan-chaibu-v1" || r["juNumber"] != float64(1) {
		t.Fatal(r)
	}
	if r["zhiShi"].(map[string]any)["palace"] != float64(8) {
		t.Fatalf("wrong zhi shi palace: %#v", r["zhiShi"])
	}
	canonical, e := Canonical(map[string]any{"year": 2026, "month": 4, "day": 10, "hour": 14, "timezone": "Asia/Shanghai"})
	if e != nil || canonical["九宫盘"] == nil {
		t.Fatalf("canonical failed: %#v %v", canonical, e)
	}
}

func TestInvalidInput(t *testing.T) {
	if _, err := Calculate(map[string]any{"year": 2026, "month": 2, "day": 30, "hour": 12}); err == nil {
		t.Fatal("expected invalid date error")
	}
}

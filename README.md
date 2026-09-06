# qimen4go

Go 奇门遁甲排盘 SDK，使用 `go:embed` 内置经过版本锁定的排盘引擎。运行时要求 Node.js 20+。

```go
input := map[string]any{
    "year": 2026, "month": 4, "day": 10, "hour": 14, "minute": 0,
    "timezone": "Asia/Shanghai",
}
chart, err := qimen.Calculate(input)
canonical, err := qimen.Canonical(input)
```

`Calculate` 返回完整结构化盘面，`Canonical` 返回适合跨语言交换和展示的规范 JSON。当前算法版本为 `qimen-zhuanpan-chaibu-v1`。

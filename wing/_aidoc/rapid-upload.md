# 快传

添加一个快传interface
```go

type RapidUploader interface {
	// 只进行快传操作
    RapidUpload(ctx context.Context, dstDir model.Obj, file model.FileStreamer, up UpdateProgress) (model.Obj, error)
}
```

* 检查 @drivers 目录，查看哪些driver有快传功能，并汇总成表格
* 将支持快传的driver，实现一下RapidUploader
# 快传

在

添加一个快传interface
```go

type RapidUploader interface {
	// 只进行快传操作
    RapidUpload(ctx context.Context, dstDir model.Obj, hash filehash.FileHashMetadata) (model.Obj, error)
}
```

* 检查 @drivers 目录，driver有快传功能:
  * baidu_netdisk
  * 123
* 将支持快传的driver，实现一下RapidUploader, 代码放在 `drivers/xxx/rapid.go`
* 无需写测试
* 在 @server/router.go `func _fs(g *gin.RouterGroup)` 中添加接口：
  * 快传接口 `/put_rapid`， 背后是 RapidUpload
    * 可以只接受某个hash信息，查寻文件名和大小，再上传
    * 也可以传入文件名、大小、hash，查询后，对比大小，再使用传入的文件名进行上传

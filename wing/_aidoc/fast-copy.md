# 快速复制

通过网盘文件的元数据，查找本地数据库，然后在目标平台快传。

### 本地数据库

使用 https://github.com/cockroachdb/pebble 存储文件的元数据， 数据文件 `filehash`

* `K:MD5:{md5_binary}` 以文件md5二进制作为主key， value=`name={xxx}|size=xxx|md5={md5_binary}|sha1={sha1_binary}` 元数据
* `K:SHA1:{sha1_binary}` 以文件sha1二进制作为key， value={md5_binary}
* `K:GCID:{gcid_binary}` ， value={gcid_binary}
* `C:{FILE_FULLPATH}|{FILE_SIZE}`： 保存本地文件的信息，避免重复计算， value={md5_binary}

通过多次查找找到文件元信息

## 功能

### 计算文件hash

在 @cmd 文件夹下 创建一个hash命令:
* 输入一个文件路径，自动计算文件元数据
* 输入一个文件夹路径，自动计算文件夹內所有文件的元数据，并保存到数据库
* 注意给元数据添加一个函数输出人类可读的信息，比如md5_binary显示为md5_hex
* `--json output.json` 表示将结果输出到一个json文件



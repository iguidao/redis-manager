# redis-manager
这是一个可以操作redis集群/codis的web管理平台
该功能还在开发中

### 功能列表
1. 支持进行[string/list/hash/set/zset]类型的key的查询
2. 支持查询大key/热key/慢key/查询1万key

### 启动方式
- mysql创建 `redis_manager` 数据库
- 执行 go mod 初始化
- 执行 go run main.go 运行代码

### API列表
Http | API | Introduce
--- | --- | --- 
GET |   /redis-manager/base/v1/health | 健康检查
POST |   /redis-manager/cli/v1/querykey | 查key命令
POST |   /redis-manager/cli/v1/bigkey | 查大key
POST |   /redis-manager/cli/v1/hotkey | 查热key
POST |   /redis-manager/cli/v1/allkey | 查1万key
POST |   /redis-manager/cli/v1/slowkey | 查慢key
POST |   /redis-manager/cli/v1/delkey | 删key

### 联系方式
mail: xiaohui920@sina.cn



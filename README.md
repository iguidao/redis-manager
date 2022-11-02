# redis-manager
这是一个可以操作redis集群/codis的web管理平台
#### 注：该功能代码还在开发中，目前只是支持codis的一些操作

### 功能列表
1. 支持进行codis集群的[string/list/hash/set/zset]类型的key的查询
2. 支持查询codis集群的大key/热key/慢key/查询1万key

### 启动方式
- mysql创建 `redis_manager` 数据库
- 执行 go mod 初始化
- 执行 go run main.go 运行代码

### API列表
Http | API | Introduce
--- | --- | --- 
GET |    /redis-manager/base/v1/health | 健康检查
POST | /redis-manager/codis/v1/add | 添加codis的manager地址
GET | /redis-manager/codis/v1/list | 列出codis的manager地址
GET | /redis-manager/codis/v1/cluster | 查看单个codis manager下的集群信息
GET | /redis-manager/codis/v1/group | 查看单个codis manager下的group信息
POST | /redis-manager/cli/v1/opkey | 进行【查key/查大key/查热key/查1万key/查慢key/删除key】等操作，目前只支持codis
POST | /redis-manager/cli/v1/analysisrdb | 分析redis的dump文件中大key的top10

### 联系方式
mail: xiaohui920@sina.cn



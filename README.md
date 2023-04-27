# redis-manager
这是一个可以操作[redis cluster/codis/腾讯云redis/阿里云redis]的web管理平台。
_________________

## 功能简介
1. **Cluster操作界面：** 支持Redis Cluster集群的添加，可以查看Redis Cluster的集群状态
2. **Codis操作界面：** 支持嵌入Codis Dashboard平台，可以查看codis平台信息
3. **云Redis操作界面：** 支持腾讯云Redis的导入，可以查看腾讯Redis的基本信息
4. **云Redis操作界面：** 支持阿里云Redis的导入，可以查看阿里Redis的基本信息（开发中...）
5. **数据查询界面：** 支持[string/list/hash/set/zset]类型的key的查询，以及查询[大key/热key/慢key/查询1万key]等功能
6. **用户界面：**  支持用户的添加删除，可以管理平台用户
7. **系统设置界面：** 支持设置全局配置以及用户权限配置，可以管理平台系统配置
8. **历史记录洁面：** 支持记载变更操作记录，方便审核回溯


## 项目启动
### 环境依赖
- mysql数据库 8.0版本以上
- 云redis或者codis，非cluster，2.8版本以上
### 启动
- mysql需要创建 `redis_manager` 数据库
- 复制 yaml/dev.yaml 到 yaml/config.yaml
- 编辑 yaml/config.yaml文件的[mysql、redis]配置
- 执行 sh start-docker.sh

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



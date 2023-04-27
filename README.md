# redis-manager
这是一个可以操作[redis cluster/codis/腾讯云redis/阿里云redis]的web管理平台。
_________________

## 功能简介
1. **Cluster操作界面：** 支持Redis Cluster集群的添加，可以查看Redis Cluster的集群状态
2. **Codis操作界面：** 支持嵌入Codis Dashboard平台，可以查看codis平台信息
3. **云Redis操作界面：** 支持腾讯云Redis的导入，可以查看腾讯Redis的基本信息
4. **云Redis操作界面：** 支持阿里云Redis的导入，可以查看阿里Redis的基本信息（开发中...）
5. **数据查询界面：** 支持[string/list/hash/set/zset]类型的key的查询，以及查询[大key/热key/慢key/查询1万key]等功能，[阿里云redis暂时不支持]
6. **用户界面：**  支持用户的添加删除，可以管理平台用户
7. **系统设置界面：** 支持设置全局配置以及用户权限配置，可以管理平台系统配置
8. **历史记录界面：** 支持记载变更操作记录，方便审核回溯


## 项目启动
### 依赖语言
- Golang1.19 + Vue3
### 环境依赖
- mysql数据库 8.0版本以上
- 云redis或者codis，或者单Redis，非cluster，2.8版本以上
### 启动
- mysql需要创建 `redis_manager` 数据库
- 复制 yaml/dev.yaml 到 yaml/config.yaml
- 编辑 yaml/config.yaml文件的[mysql、redis]配置
- 执行 sh start-docker.sh

## 项目依赖
### 大key查询依赖
- 如果需要在自建Cluster和Codis上查询大key，则需要在每个redis机器上安装Agent，检测bgsave的dump文件，上传到cos里，才能进行使用；项目地址：[redis-agent](https://github.com/iguidao/redis-agent)


## 主要功能界面介绍

### 登陆
> 账号密码登陆
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/login.png"/>

### 概览  
> 展示集群基本信息
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/board.png"/>

### Cluster集群页面
> 展示Redis Cluster集群的信息
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/cluster.png"/>

### Codis集群页面
> 展示Codis集群信息
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/codis.png"/>

### 腾讯Redis集群界面
> 展示腾讯Redis集群信息
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/tx-redis.png"/>>

### 阿里Redis集群界面
> 展示阿里Redis集群信息[未测试]
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/ali-redis.png"/>

### 数据查询界面
> 进行redis的key操作，支持[查key/大key/热key/慢key/查询1万key/删key]功能
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/cli.png"/>

### 用户管理界面
> 进行用户的创建删除，身份更改功能
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/user.png"/>

### 系统设置界面
> 进行全局配置
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/setting-cfg.png"/>
> 进行权限配置
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/setting-rule.png"/>

### 历史记录界面
> 所有变更操作都会记录
<img src="https://raw.githubusercontent.com/iguidao/img-folder/master/redis-manager/history.png"/>

## 联系方式
暂无



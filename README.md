# breeze-api

### 实现 web 快速开发的脚手架

### 使用说明
1. ###### 克隆项目
```
git clone https://github.com/mengxiangyu996/breeze-api.git
```
2. ###### 进入目录
```
cd breeze-api
```
3. ###### 修改配置文件
```
mv env.json.example env.json
```
4. ###### 安装依赖
```
go mod tidy
```
5. ###### 复制依赖到本项目
```
go mod vendor
```
6. ###### 启动服务
```
go run main.go
```

### API文档
* [点击查看api文档](https://melodious-centaur-e00019.netlify.app/)

### 阿里云流水线
* #### 构建
###### 对于常规构建，可以使用以下命令
```
# 默认使用goproxy.cn
export GOPROXY=https://goproxy.cn
# input your command here
go build main.go
```
###### 如果需要为特定平台（例如 arm64 架构的 Linux）构建，可以使用以下命令
```
export GOPROXY=https://goproxy.cn
# input your command here
GOARCH=arm64 GOOS=linux go build -o main main.go
```
###### 构建物上传配置
![image](https://github.com/user-attachments/assets/8ffe3e99-3f09-4326-9abd-df807b1ac021)


* #### 部署
###### 部署命令
```
tar zxvf /opt/1panel/apps/openresty/openresty/www/sites/breeze/package.tgz -C /opt/1panel/apps/openresty/openresty/www/sites/breeze/
cd /opt/1panel/apps/openresty/openresty/www/sites/breeze
bash service.sh restart
```

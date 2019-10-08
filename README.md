# heroes-service
准备工作
```
1.golang
2.docker
3.docker-compose
4.hyperledger fabric
5.hyperledger fabric-sdk-go
6.beego 安装流程查看beego官网https://beego.me/
```
拉取源代码
```
mkdir -p $GOPATH/src/github.com/dddengyunjie
cd $GOPATH/src/github.com/dddengyunjie
git clone https://github.com/dddengyunjie/heroes-service.git
```
启动镜像
```
cd $GOPATH/src/github.com/dddengyunjie/heroes-service/fixtures
./startup_couchdb.sh
```
等待所有镜像正常启动后再继续
启动go web服务
```
cd $GOPATH/src/github.com/dddengyunjie/heroes-service/web
bee run 
```
启动成功后输出如下语句
http server Running on http://:8080
Admin server Running on :8088
可在浏览器中查看http://localhost:8080/ 
管理界面http://localhost:8088/ 

接口调用
1.初始化marble,示例初始化一个名称为marble1，颜色为blue，大小为10，所有者为tom的marble:
http://localhost:8080/initMarble/marble1/blue/10/tom
2.根据颜色交易marble
http://localhost:8080/transferMarblesBasedOnColor/blue/lily
这些接口都是直接用 hyperledger/fabric-samples/chaincode/marbles02/ 里的，具体可查看https://github.com/hyperledger/fabric-samples/tree/release-1.4/chaincode/marbles02
关于其它接口可查看web/main.go 文件，接口这块还很粗糙，后续再改进。

# Go Project Template
[`Project Template`](http://tp-km.tutorabc.com/display/TTK/02+project+template)


# 準備
[`完成環境設置`](http://tp-km.tutorabc.com/display/TTK/01+env)



### 使用方式
1. Windows中，限使用`git bash` 同linux shell
2. `git clone "http://scm.tutorabc.com/jimzhong/shop-payment"`
3. cd shop-payment
4. `./initialization_project.sh ${project_name}`
5. cd ..
6. cd ${project_name}


### 測試Run
1. make
2. make run
3. 以上沒出錯，則沒問題


## Go 目錄結構

### `/conf`

放config的地方，默認**service.conf**
不同環境命名為**service.conf.${ENV}**，例如**service.conf.prod**

### `/internationalization`
放置 **i18n**多語系的地方，詳細操作可見Makefile


### `/routers`
放置 API router的地方，用到的middleware放置  `/middleware`


## * `/shop-payment` *
主要放置專案邏輯的地方
### `/shop-payment/client`
創建client，例如 redisClient, dbClient, httpClient...
### `/shop-payment/controller`
API controller
### `/shop-payment/controller_model`
定義 reqeust / response Struct <少用....盡量使用idl>
### `/shop-payment/dao`:
dao邏輯
### `/shop-payment/dao_model`
dao Struct ，創建方式參考[`Xorm tool`](http://tp-km.tutorabc.com/display/TTK/Xorm+Tool+-+Gen+struct+from+Table)
### `/shop-payment/rpc`
call 其他服務邏輯
### `/shop-payment/rpc_model`
其他服務 requeset / response
### `/shop-payment/schedule`
如果有排程
### `/shop-payment/service`
主要商業邏輯地方 （Business Logic Layer）
### `/shop-payment/service_model`
service用到的 Struct
### `/shop-payment/sutil`
使用到的util

## 部署
Docker裡面 直接用makefile 命令build，保證開發與部署一致
部署時運行control.sh 切換config file && 啟動服務

## API Interface 使用  Interface description language（IDL）

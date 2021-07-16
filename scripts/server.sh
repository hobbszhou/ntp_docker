#!/bin/bash

myargv=`echo $IDC`
confPrefix="auth-api"
confSuffix="yaml"
APP_ARGS_TMP="auth-api.yaml"


## 创建日志目录
mkdir -p $LOG_PATH

# 修改目录权限
chmod 777 -R $LOG_PATH
if [[ -n $myargv ]]
then
  APP_ARGS_TMP="$confPrefix""-""$myargv"".""$confSuffix"
fi

## 运行程序
$APP_PATH/$APP_NAME -f ./etc/$APP_ARGS_TMP
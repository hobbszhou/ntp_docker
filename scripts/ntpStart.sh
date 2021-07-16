#!/bin/bash

ntpAddr=$1
ntpConfPath=$2

#echo $ntpConfPath

#1. 替换配置
num=0
for date in $(sed -n '/^server/p' $ntpConfPath)
do
  if [ $num == 1 ]
  then
    sed -i -e "s/$date/$ntpAddr/g" $ntpConfPath
  fi
  let num=$num+1
#  echo $num
done

#ntpdate 120.25.115.20
#2. 杀死正在运行的ntpd服务
num=0
#echo $num
for date in $(ps -aux |grep ntpd | awk -F " " '{print $2}')
do
    kill -9 $date

done
#3. 启动ntpd
#/usr/sbin/ntpd -c /etc/ntp.conf -p /tmp/ntpd.pid -d
service ntpd restart
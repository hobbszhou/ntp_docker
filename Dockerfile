FROM dockerhub.venuscloud.cn/base/alpine:3.8

LABEL maintainer="zhou_hongbo@venusgroup.com.cn"

# 环境变量设置
ENV APP_NAME ntp-server
ENV APP_ARGS ./etc/auth-api.yaml
ENV APP_ROOT /var/www
ENV APP_PATH $APP_ROOT/$APP_NAME
ENV LOG_PATH $APP_ROOT/$APP_NAME/logs

# 执行入口文件添加
WORKDIR $APP_PATH

ADD ./cmd/api/etc/* $APP_PATH/etc/
ADD $APP_NAME $APP_PATH/
ADD scripts/* $APP_PATH/scripts
ADD ntp.conf /data/

RUN chmod +x $APP_PATH/scripts/* \
&& yum install ntp ntpdata -y \
&& yum install initscripts -y \
&& rm -rf /etc/ntp.conf \
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& cp /data/ntp.conf /etc/ \
&& mkdir docs \
&& cp $APP_PATH/scripts/server.sh ./
CMD $APP_PATH/server.sh && tail -f $APP_PATH/start.sh


#ADD ntp.conf /data/
#ADD start.sh /root/zhb/
#
#RUN yum install ntp ntpdata -y && rm -rf /etc/ntp.conf \
#&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
#&& chmod +x /root/zhb/start.sh \
#&& cp /data/ntp.conf /etc/
#CMD /root/zhb/start.sh && tail -f /root/zhb/start.sh
#########################################################################
# File Name: start.sh
# Description: start.sh
# Author: liangjf
# mail: liangjf_dada@163.com
# Created Time: 2019-06-04 23:53:16
#########################################################################
#!/bin/bash
echo '删除日志文件'
rm -f log/nsqlookupd.log
rm -f log/nsqd1.log
rm -f log/nsqd2.log
rm -f log/nsqadmin.log

echo '启动nsq服务'
nohup nsqlookupd >log/nsqlookupd.log 2>&1&
sleep 1
echo '启动nsqd服务'
nohup nsqd --lookupd-tcp-address=0.0.0.0:4160 -tcp-address="0.0.0.0:4151" -http-address="0.0.0.0:4153" --data-path=/home/liangjf/study/go_home/src/myself/nsq_data/nsq1 >log/nsqd1.log 2>&1&
sleep 1
nohup nsqd --lookupd-tcp-address=0.0.0.0:4160 -tcp-address="0.0.0.0:4152" -http-address="0.0.0.0:4154" --data-path=/home/liangjf/study/go_home/src/myself/nsq_data/nsq2 >log/nsqd2.log 2>&1&
sleep 1
echo '启动nsqdadmin服务'
nohup nsqadmin --lookupd-http-address=0.0.0.0:4161 >log/nsqadmin.log 2>&1&

ps -aux | grep nsq

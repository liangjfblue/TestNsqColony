#########################################################################
# File Name: stop.sh
# Description: stop.sh
# Author: liangjf
# mail: liangjf_dada@163.com
# Created Time: 2019-06-04 23:39:36
#########################################################################
#!/bin/bash
ps -ef | grep nsq| grep -v grep | awk '{print $2}' | xargs kill -2

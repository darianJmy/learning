#!/bin/bash

backDir=/opt
backTime=$(date +%Y-%m-%d_%H-%M-%S)
backuplog=etcd_backup.log

export ETCDCTL_API=3 && /usr/local/bin/etcdctl --endpoints 127.0.0.1:2379  --cacert="/etc/etcd/ssl/ca.crt" --cert="/etc/etcd/ssl/client.crt" --key="/etc/etcd/ssl/client.key" snapshot save $backDir/etcd_$backTime.db >> /dev/null 2>&1
gzip $backDir/etcd_$backTime.db
echo "备份完成 $backDir/etcd_$backTime.db" >> $backDir/$backuplog 2>&1

if [ $(echo `find $backDir -name '*.db.gz' -mtime +30` |awk '{print length($0)}') -eq 0 ]; then
    echo "没有要删除的备份文件" >> $backDir/$backuplog 2>&1
    exit
fi

rm -rf `find $backDir -name '*.db.gz' -mtime +30` >> $backupDir/$backuplog 2>&1
echo "删除完成 `find $backDir -name '*.db.gz' -mtime +30`" >> $backDir/$backuplog 2>&1

## crontab -e
## 0 3 * * *   /usr/sbin/mysql-bakup.sh
## 每天3点定时执行
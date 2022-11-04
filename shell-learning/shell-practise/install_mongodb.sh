#!/usr/bin/env bash

set -x

server_package=mongodb-org-server-4.4.8-1.el8.x86_64.rpm
shell_package=mongodb-org-shell-4.4.8-1.el8.x86_64.rpm
file_path=mongodb

function _ensure_file_path {
    if  [[ ! -d /service/ ]]; then
        mkdir -p /service/
    fi

    if ! ls /service/{logs,software,backup} > /dev/null 2>&1; then
        mkdir -p /service/{logs,software,backup}
    fi

    if ! type wget >/dev/null 2>&1; then
        yum -y install wget
    fi
}

function _download_mongo {
    _ensure_file_path

    if [[ ! -f /service/software/$server_package ]]; then
        cd /service/software && wget http://software.voneyun.com/$server_package
        if [ $? -ne 0 ]; then
            echo "wget mongodb-org-server package failed"
            exit 1
        fi
    fi

    if [[ ! -f /service/software/$shell_package ]]; then
        cd /service/software && wget http://software.voneyun.com/$shell_package
        if [ $? -ne 0 ]; then
            echo "wget mongodb-org-shell package failed"
            exit 1
        fi
    fi
}


function install_mongodb {
    _download_mongo

    cd /service/software/ && rpm -ivh mongodb-org-*

    if [[ ! -d /service/logs/$file_path ]]; then
        mkdir -p /service/logs/$file_path
    fi

    if [[ ! -d /service/software/$file_path ]]; then
        mkdir -p /service/software/$file_path
    fi

    if [[ ! -d /service/software/$file_path/data ]]; then
        mkdir -p /service/software/$file_path/data
    fi

    if [[ ! -d /service/software/$file_path/keyfile ]]; then
        mkdir -p /service/software/$file_path/keyfile
        openssl rand -base64 745 > /service/software/$file_path/keyfile/mongo-keyfile
        chmod 600 /service/software/$file_path/keyfile/mongo-keyfile
    fi

    chown -R mongod:mongod /service/software/mongodb
    chown -R mongod:mongod /service/logs/mongodb    

    change_config
    systemctl start mongod
}

function change_config {
    cat > /etc/mongod.conf << EOF
# mongod.conf

# for documentation of all options, see:
#   http://docs.mongodb.org/manual/reference/configuration-options/

# where to write logging data.
systemLog:
  destination: file
  logAppend: true
  path: /service/logs/mongodb/mongod.log

# Where and how to store data.
storage:
  dbPath: /service/software/mongodb
  journal:
    enabled: true
#  engine:
#  wiredTiger:

# how the process runs
processManagement:
  fork: true  # fork and run in background
  pidFilePath: /var/run/mongodb/mongod.pid  # location of pidfile
  timeZoneInfo: /usr/share/zoneinfo

# network interfaces
net:
  port: 27017
  bindIp: 127.0.0.1  # Enter 0.0.0.0,:: to bind to all IPv4 and IPv6 addresses or, alternatively, use the net.bindIpAll setting.


security:
  keyFile: /service/software/mongodb/keyfile/mongo-keyfile

#operationProfiling:

#replication:

#sharding:

## Enterprise-Only Options

#auditLog:

#snmp:
EOF
}

function create_admin {
    mongo --host 127.0.0.1 --port 27017 admin --eval 'printjson(db.createUser({ user: "admin", pwd: "scmongo@SX.1", roles: [{ role: "root", db: "admin" }] }))'
}

install_mongodb
create_admin

#!/usr/bin/env bash

set -x

package=redis_5.0.6_cluster.tar.gz
file_path=redis_5.0.6

function _ensure_file_path {
    if  [[ ! -d /service/ ]]; then
        mkdir -p /service/
    fi

    if ! ls /service/{logs,software,backup} > /dev/null 2>&1; then
        mkdir -p /service/{logs,software,backup}
    fi

    if [[ ! -d /service/logs/redis ]]; then
        mkdir -p /service/logs/redis
    fi

    if [ ! type wget >/dev/null 2>&1 ]; then
        yum -y install wget
    fi
}

function _download_redis {
    _ensure_file_path

    if [[ ! -f /service/software/$package ]]; then
        cd /service/software && wget http://software.voneyun.com/redis/$package
        if [ $? -ne 0 ]; then
            echo "wget redis package failed"
            exit 1
        fi
    fi
}


function install_redis {
    _download_redis

    if [[ -d /service/software/$file_path ]]; then
        rm -rf /service/software/$file_path
        rm -rf /bin/redis-*
    fi

    tar -xf /service/software/$package -C /service/software
    sed -i "s/^cluster-enabled/#cluster-enabled/g" /service/software/$file_path/conf/redis.conf
    ln -s /service/software/$file_path/bin/* /bin/    
}


function create_system {
    cat > /etc/systemd/system/redis.service << EOF
[Unit]
Description=Redis In-Memory Data Store
After=network.target

[Service]
User=root
Group=root
Type=forking
ExecStart=/service/software/redis_5.0.6/bin/redis-server  /service/software/redis_5.0.6/conf/redis.conf
ExecStop=/service/software/redis_5.0.6/bin/redis-cli shutdown
Restart=always

[Install]
WantedBy=multi-user.target    
EOF

    if [[ -f /etc/systemd/system/redis.service ]]; then
        systemctl daemon-reload
        systemctl restart redis
        systemctl enable redis
    else
        echo "start redis failed"    
        exit 1
    fi    
    
}


install_redis
create_system

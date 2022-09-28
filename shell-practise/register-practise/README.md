## Delete-Register-Images

### 功能
**删除镜像仓库 Register 里面的Image**

**脚本有三个参数需要完善：**

* 镜像仓库地址与Port
* 镜像名称
* 镜像标签

```
registry_host_port=$1
image_name=$2
image_tag=$3
```


**可以通过 delete-register-images.sh 删除镜像里面image**
```
./delete-register-images.sh 192.168.100.100:5000  centos latest
```
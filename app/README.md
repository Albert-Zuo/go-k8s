### 安装环境（略）

1. Docker

2. Kubectl

3. Golang

4. Kind

   ```shell
   ❯ GO111MODULE="on" go get sigs.k8s.io/kind@v0.10.0
   
   ❯ kind version
   kind v0.10.0 go1.16 linux/amd64
   
   ❯ kind create cluster
   Creating cluster "kind" ...
    ✓ Ensuring node image (kindest/node:v1.20.2) 🖼
    ✓ Preparing nodes 📦
    ✓ Writing configuration 📜
    ✓ Starting control-plane 🕹️
    ✓ Installing CNI 🔌
    ✓ Installing StorageClass 💾
   Set kubectl context to "kind-kind"
   You can now use your cluster with:
   
   kubectl cluster-info --context kind-kind
   
   Thanks for using kind! 😊
   ```



### 一、容器

1、构建Docker镜像

```shell
[root@localhost ~]# git clone https://github.com/Albert-Zuo/go-k8s.git
[root@localhost ~]# cd go-k8s/app
[root@localhost app]# docker build -t monitor:v1.0.0 .
[root@localhost app]# docker images
REPOSITORY           TAG           IMAGE ID       CREATED         SIZE
albertzuo/monitor    v1.0.0        135decd1bd5f   23 hours ago    304MB
golang               1.16-alpine   1b35785aa3c4   2 weeks ago     302MB
```



2、push 镜像仓库（hub.docker 公有仓库为例）

https://hub.docker.com/ 注册账号

```shell
## 登录docker公共仓库，输入账号密码
[root@localhost app]# docker login
Login Succeeded

[root@localhost app]# docker_username=`docker info | grep 'Username' | awk '{print $2}'`
[root@localhost app]# docker push ${docker_username}/monitor:v1.0.0
```



3、构建容器

```shell
[root@localhost app]# cd k8s
[root@localhost k8s]# sed -i "s/albertzuo/${docker_username}/" monitor.yml
[root@localhost k8s]# kubectl apply -f monitor.yml
pod/monitor created
[root@localhost k8s]# kubectl get po
NAME                                READY   STATUS    RESTARTS   AGE
monitor                             1/1     Running   0          31s
```

4、操作容器

查看日志

```shell
[root@localhost k8s]# kubectl logs -f  monitor
0, Hello World!
1, Hello World!
2, Hello World!
3, Hello World!
4, Hello World!
5, Hello World!
6, Hello World!
7, Hello World!
8, Hello World!
9, Hello World!
10, Hello World!
11, Hello World!
```



Pod详情(容器生成失败，可以查看错误日志)

```shell
[root@localhost app]# kubectl describe po monitor
```


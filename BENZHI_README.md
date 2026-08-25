Go 实现的近轴光学 ABCD 矩阵分析服务，默认启动 HTTP 在 :8080 提供光线传输矩阵计算与望远镜分析接口，也可通过 trace/telescope 子命令在命令行对 JSON 元件序列做离线计算。

## 构建与启动

```bash
go build -o abcd-optics .
./abcd-optics                              # 启动 HTTP 服务 :8080
./abcd-optics trace example/telescope.json # 命令行分析
```

## 评测镜像

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh abcd-optics
```

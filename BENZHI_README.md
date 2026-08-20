# abcd-optics

abcd-optics 是近轴（paraxial）光学的 ABCD 光线传输核算内核：输入一个元件序列（自由空间、薄透镜、球面折射面）与物距，用 Go 做 2×2 矩阵连乘得到系统矩阵 M = M_n…M_1，然后输出 M 的 A,B,C,D、行列式、有效焦距 f_eff = −1/C，以及给定物距 s 的像距 s' 和横向放大率 m。

## 构建 / 运行 / 测试

```text
go build ./...     # 编译
go run . trace example/telescope.json
go test ./...      # 测试
```

## 评测镜像

本目录评测专用文件（勿覆盖项目自带 Dockerfile/README）：

- `benzhi.Dockerfile`
- `build_benzhi_docker.sh`
- `BENZHI_README.md`（本文件）

两种架构都要构建并进容器验证：

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh <image-name> linux/arm64
./build_benzhi_docker.sh <image-name> linux/amd64
docker run -it <image-name>:latest
```

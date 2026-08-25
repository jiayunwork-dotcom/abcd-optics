# abcd-optics

abcd-optics 是近轴（paraxial）光学的 ABCD 光线传输核算内核：输入一个元件序列（自由空间、薄透镜、球面折射面）与物距，用 Go 做 2×2 矩阵连乘得到系统矩阵 M = M_n…M_1，然后输出 M 的 A,B,C,D、行列式、有效焦距 f_eff = −1/C，以及给定物距 s 的像距 s' 和横向放大率 m。约定列向量为 [y, u]，u 是近轴斜率角（Helmholtz 不变量用角度约定，空气折射率=1）；自由空间 [[1,L],[0,1]]、薄透镜 [[1,0],[−1/f,1]]、球面折射 [[1,0],[(n1−n2)/(R·n2), n1/n2]]。边界：负传播长度、零焦距、非正折射率、未知元件一律 error 并以 stderr + 非零退出码结束；本仓只做单序列近轴成像，不做波动光学、不做相机滤镜、不做颜色科学。

- 输入：一个 JSON 算例文件，`elements` 是有序元件数组（第一个元件光线先遇到），`object_distance` 是物平面到第一面的距离，可选 `ray` 为入射光线；见 `example/telescope.json`
- 输出：系统矩阵 A,B,C,D、det(M)、有效焦距、物距、像距 s'、放大率 m、逐元件出射光线、交叉规则检查（det≈1、薄透镜 1/s+1/s'=1/f、光线正走后再逆序求逆回到入射）
- 边界：空间 L<0、薄透镜 f=0、折射面 R=0 或 n≤0、未知元件 kind、缺必填字段、物距≤0、系统空序列均报错；无焦系统 C≈0 时有效焦距与主平面未定义
- 副命令 `telescope f1 f2 spacing`：两薄透镜间距分析，间距 f1+f2 时系统无焦（C→0），放大率逼近 −f2/f1

## 复现命令

```text
go run . trace example/telescope.json
```

`example/telescope.json` 是两片薄透镜（f1=100, f2=50，间距 150=f1+f2）构成的近无焦望远镜，物距 500。期望输出：C≈0（无焦）、放大率 m≈−0.5（即 −f2/f1，负号表示倒像）、det(M)=1、round-trip 检查 PASS。

## 构建 / 运行 / 测试

```text
go build ./...                    # 编译（纯标准库）
go test ./...                     # 全部测试（matrix / element / system / server）
go run .                          # 启动 HTTP 服务 :8080
go run . trace example/telescope.json
go run . telescope 100 50 150     # 无焦望远镜：C=0, m=-0.5
go run . telescope 100 50 170     # 间距偏离：出现焦度 C≠0
```

## HTTP API
默认无参启动 HTTP 服务，提供以下接口：
- `POST /api/trace` — 提交 JSON 光学序列，返回系统矩阵与成像报告
- `GET /api/telescope?f1=100\&f2=-25\&spacing=75` — 分析两薄透镜望远镜
- `GET /health` — 健康检查
## 非法输入示例

```text
go run . trace <(echo '{"object_distance":50,"elements":[{"kind":"thin_lens","focal":0}]}')
# → stderr 报 zero focal length，退出码 1
```

### JSON 元件格式

| kind | 字段 | 矩阵 |
|------|------|------|
| `space` | `length`（L≥0） | [[1,L],[0,1]] |
| `thin_lens` | `focal`（f≠0） | [[1,0],[−1/f,1]] |
| `refraction` | `radius`, `n1`, `n2`（n>0, R≠0） | [[1,0],[(n1−n2)/(R·n2), n1/n2]] |

成像：从物平面（距第一面 s）到像平面（距最后一面 s'）的总矩阵 B 分量置零，得
s' = −(A·s + B)/(C·s + D)，放大率 m = A + C·s'（等于 det(M)/(C·s+D)，空气到空气 det=1）。平行光入射（物在无穷远）像在 −A/C 处，f_eff = −1/C。

## 目录

```text
main.go                    # 入口：默认启动 HTTP / trace / telescope / help
internal/matrix/           # 2×2 矩阵、光线、求逆、往返、特征、稳定性
internal/element/          # 元件模型、JSON 解析、校验、介质连续性
internal/system/           # 系统组装、成像、焦距、望远镜、交叉规则、报告
internal/server/           # HTTP 服务：/api/trace, /api/telescope, /health
example/telescope.json     # 无焦望远镜算例
```

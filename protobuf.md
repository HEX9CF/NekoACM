# ProtoBuf

## 生成 pb 文件

```bash
protoc --go_out=pkg/pb --go-grpc_out=pkg/pb api/proto/*.proto
```
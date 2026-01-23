# uniladin-backend-libs

Bộ thư viện dùng chung cho các service Go trong hệ thống Uniladin. Mục tiêu là tái sử dụng hạ tầng, thống nhất contract giao tiếp giữa các service, và giảm trùng lặp code.

## Các thư viện hiện có

### rabbitmq
Thư viện RabbitMQ dùng chung cho việc kết nối, publish/consume, chuẩn hóa exchange/queue/routing và định nghĩa event giữa các service.
- Mã nguồn: `libs/rabbitmq`
- Contract sự kiện: `libs/rabbitmq/services/<service>`
- Hướng dẫn chi tiết: `libs/rabbitmq/README.md`

## Cách link submodule

```bash
git submodule add https://github.com/Uniladin/backend-libs-go.git libs
git submodule update --init --recursive
```

Cập nhật submodule khi cần:

```bash
git submodule update --remote libs
```

## Thiết lập go.mod

Trong service, thêm `replace` và `require` để dùng libs từ thư mục `./libs`:

```go
replace uniladin-backend-libs => ./libs

require (
    uniladin-backend-libs v0.0.0-00010101000000-000000000000
)
```

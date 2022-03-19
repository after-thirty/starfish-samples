# 案例使用帮助

1. 创建 starfish_order 数据库 并导入脚本
2. 启动 order_svc
3. 创建 starfish_product 数据库 并导入脚本
4. 启动 product_svc
5. 直接执行 aggregation_svc main.go
6. 调用 aggregation_svc API /createSoCommit

## 测试流程
测试过程中 可以把 order_svc 或者 product_svc 任何一个停掉，全局事务都会回滚

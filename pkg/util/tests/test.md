#### 对于微服务来说，只进行pkg实现测试
* 记录：同一个包不同文件下方法无法访问，测试时需要输入例如: go test -v snowflake_test.go snowflake.go  

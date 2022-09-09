## 测试
```shell
 ✗ go test --benchmem -bench="." -v           
=== RUN   TestGetInstance
--- PASS: TestGetInstance (0.00s)
=== RUN   TestGetLazySingleton
--- PASS: TestGetLazySingleton (0.00s)
goos: darwin
goarch: amd64
pkg: golang_design_pattern/1.singleton
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkGetInstance
BenchmarkGetInstance-12                 1000000000               0.1418 ns/op          0 B/op          0 allocs/op
BenchmarkGetLazySingleton
BenchmarkGetLazySingleton-12            1000000000               0.4367 ns/op          0 B/op          0 allocs/op
PASS
ok      golang_design_pattern/1.singleton       3.730s

```
可以看到，饿汉式的单例要比懒汉式的单例的性能要好一点。
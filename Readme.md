
# 为什么选fiber

参考[go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark)

# Error相关处理
1. [pkg/errors](https://github.com/pkg/errors)停止维护，可替换为[cockroachdb/errors](https://github.com/cockroachdb/errors)

2. 何时打印调用栈
官方库无法 `wrap` 调用栈，所以 `fmt.Errorf %w` 不如 `pkg/errors` 库实用，但是 `errors.Wrap` 最好保证只调用一次，否则全是重复调用栈

我们项目的使用情况是 `log error` 级别的打印栈，warn 和 info 都不打印，当然 case by case 还得看实际使用情况

如果`Wrap`多次可以看这里 [为什么不允许处处使用 errors.Wrap](https://lailin.xyz/post/go-training-03.html#%E4%B8%BA%E4%BB%80%E4%B9%88%E4%B8%8D%E5%85%81%E8%AE%B8%E5%A4%84%E5%A4%84%E4%BD%BF%E7%94%A8-errors-Wrap)


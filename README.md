### 一个Go的并发安全基本数据结构的实践

### 使用

```
import ds "github.com/shameby/data-structures"
```

### Set
初始化

```
NewSet(isConcurrency bool) Set
```

初始化参数

|参数名|类型|说明|
|:---|:----- |----|
|isConcurrency|bool|true为并发安全，反之则不是|

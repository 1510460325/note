## Go map
***数据结构***
```go
type hmap struct {
    count     int
    flags     uint8
    // bucket的数量是2^B, 最多可以放 loadFactor * 2^B 个元素，再多就要 hashGrow 了
    B         uint8
    //overflow 的 bucket 近似数
    noverflow uint16
    hash0     uint32 // hash seed
    //2^B 大小的数组，如果 count == 0 的话，可能是 nil
    buckets    unsafe.Pointer
    // 扩容的时候，buckets 长度会是 oldbuckets 的两倍,只有在 growing 时候为空。
    oldbuckets unsafe.Pointer
    // 指示扩容进度，小于此地址的 buckets 迁移完成
    nevacuate  uintptr
}

type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}

```
***插入*** :
* 根据 key 算落到 bucket 的位置
* 遍历高八位数组，如果为空则直接插入，如果高八位相等，则查看是否 key 相等，相等则更新
* 8 个满了，则通过 overflow 找到下个桶
* 如果最后一个桶也满了，则新建桶，插入数据

***扩容***: 增量扩容  
当 count 达到 6.5 * 2 ^ B 次方或者 overflow 过多，会触发扩容操作
* 新建一个 bucket 数组
* 写操作时如果发现扩容还在进行中（old != nil），将会 growWork 搬迁对应的 bucket
* 如果所有 bucket 都已经迁移完毕，则删除 old 指针
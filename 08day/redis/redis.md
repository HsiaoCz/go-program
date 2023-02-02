# redis 基础知识和go操作redis

redis 默认 16个数据库 默认使用第0个
DBSIZE：查看数据库的大小
select 选择数据库
keys *：查看所有的键
flushdb:清除当前的数据库
flushall:清除所有的数据库

reids单线程，基于内存操作

redis的五个基本数据类型:
字符串、哈希表、列表、集合、有序集合

## 1、key

keys *：查看所有的key
EXISTS key:判断一个key是否存在
move key:移除一个键值对
EXPIRE name:设置10秒钟过期
ttl name:查看还有多久过期
type key:查看类型
remove key:移除key

## 2、字符串

set key value:设置键值对
get key:获取值
EXISTS key:判断一个键是否存在
APPEND key value:追加值，键不存在相当于set
STRLEN key:获取字符串的长度
incr key:相当于自增
decr key:相当于自减
INCRBY key 10:指定增量 +=10
DECRBY key 10:指定减量 -=10
GETRANGE key 0 3 :截取字符串，相当于从0截取到下标
GETRANGE key 0 -1:截取字符串，截取全部
setrange key 1 a :替换字符串，以指定开始的位置
setex key 10 value:设置过期时间
setnx key value :设置值，没有会设置，再次使用不会生效
mset k1 v1 k2 v2 k3 v3:一次设置多个值
mget k1 k2 k3 :一次获取多个值

set user:{name:zhangsan,age:3} :设置一个对象，值为json字符串
getset key value:不存在返回nil,存在设置值，返回之前的值

## 3、List

列表：list看成栈或者队列 后进的在前面
list命令以l开头

lpush list one:将一个值或者多个值插入到列表的头部
lrange list 0 -1:以区间取值 左右都包含
lpop list :移除列表的第一个元素
rpop:移除列表的最后一个元素
lindex list 1:通过下标获取值
Llen list :获取列表长度
lrem list 1 one :移除指定的值 和个数
ltrim list 0 1 :截取list 指定下标的值，这个截取之后会改变list，整个list只剩下下标的值还存在
rpoplpush list1 list2 :将list2的最后一个值移动到list1中
lset list 0 item:将list指定下标的值进行替换
linsert list before value1 value2:在value1之前插入一个新的值
linsert list after value1 value2：在value1之后插入value2

list实际上是一个列表,key不存在会创建新的链表
key存在，新增内容,移除所有的值，空链表也代表不存在
两边插入或者改动值，效率最高，中间效率低

## 4、set集合

集合中的元素不允许重复,且无序
sadd myset value :添加值
smembers myset :查看指定的值
sismember myset hello:判断一个值是否存在
scard myset ：获取set中的值
srem myset hello ：移除set中的值
srandmember myset 2 ：随机抽取两个元素
srandmember myset :随机抽取一个元素
spop myset :随机删除一个集合中的元素
smove myset1 myset2 "hello" ：将指定值移动到另一个集合
sdiff s1 s2:两个集合的差集
sinter s1 s2:两个集合的交集
sunion s1 s2:两个集合的并集

集合可以用来做关注、粉丝什么的

## 5、hsah

hash一个map集合,值是map
hset myhash filed1 hello:设置key value
hget myhash filed1 ：获取字段值
hmst myhash filed1 hah filed2 hahah:一次设置多个值
hmget myhash filed1 filed2:一次获取多个值
hgetall myhash :一次获取全部的值
hdel myhash filed1:删除指定的key和相应的值
hlen myhash ：查看长度
hexists myhash filed1：判断hash中指定的字段是否存在
hkeys myhash：只获取所有的value
hvals myhash :只获取value
hincrby myhash filed1 1:给字段值加1
hdecrby myhash filed1 1:给字段值减1
hsetnx myhash filed2 hello:不存在设置值

hash适合对象存储,string 适合存储字符串

## 6、zset

有序集合
zadd myset 1 one :设置一个值
zadd myset 2 two 3 three :设置多个值
zadd salary 20000 zhangsan
zadd salary 30000 lisi
zrangebycore salary -inf +inf  withscores:查看数据
zrangebycore salary -inf 3000 withscores :查看数据到3000
zrevrange salary 0 -1:从小到大排序
zrem salary zhangsan :去除
zcard salary ：获取有序集合的元素个数
zcount myset 1 3：获取指定区间的成员数量

有序集合可以用来设置排行榜，成绩单

## 7、特殊数据类型:geospatial 地理位置

定位，附近的人，打车距离计算

geoadd china:city 纬度 经度  城市名称
geopos china:city beijing :获取指定城市的经纬度
geodist china:city beijing shanghai km :查看指定城市到城市的距离
georadius china:city 110 30 1000 km :在给定的经纬度的情况下，查找半径1000km的城市
geohash china:city beijing chongqi :返回二维的经纬度转化成字符串，两个字符串越接近，距离越近
georadiusmemery :找出指定范围内的元素，中心点是指定的

## 8、hyperloglog 基数

基数：不重复的元素

pfadd mykey a b c d e f g h i :添加
pfcount mykey :统计基数数量
prmerge mykey3 mykey1 mykey2:合并两组

## 9、bitmaps

位存储，统计用户信息，只有两种状态 0/1
setbit sign 0 1
setbit sign 1 0
打卡,周一打了，周二没打
getbit sign 1:查看状态
bitcount sign :统计打卡数 就是统计1的个数

## 10、redis事物

redis事物不是原子的，redis单条命令是原子的
redis事物没有隔离级别
所有的命令在事物中，没有被直接执行，只有发起执行命令的时候才会执行
multi:开启事物
set name zhangsan;
get name
exec: 执行事物
discard:放弃事物

监控:
悲观锁:无论做什么都会加锁
乐观锁:更新数据的时候判断一下，在此期间是否有人修改数据
获取version 更新的时候比较version

set money 100
set out 0
watch money  监控money对象
multi 开启事物
decrby money 20
incrby out 20
exec
在这个事物过程中，数据没有发生变动，事物能够执行成功
如果这时有别的线程修改了数据，会执行失败，执行失败可以先UNWATCH
再watch,之后再次执行事物

## 11、go-redis

安装依赖
go get -u github.com/go-redis/redis

```go
// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
redisDb = redis.NewClient(&redis.Options{
Addr:     "localhost:6379", // redis地址
Password: "",               // redis密码，没有则留空
DB:       0,                // 默认数据库，默认是0
})

//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
_, err = redisDb.Ping().Result()
if err != nil {
return err
}
return nil
}

func main() {
err := initClient()
if err != nil {
//redis连接错误
panic(err)
}
fmt.Println("Redis连接成功")
}
```

redis连接参数解析:

```go
type Options struct {
// 网络类型 tcp 或者 unix.
// 默认是 tcp.
Network string
// redis地址，格式 host:port
Addr string

// 新建一个redis连接的时候，会回调这个函数
OnConnect func(*Conn) error

// redis密码，redis server没有设置可以为空。
Password string
// redis数据库，序号从0开始，默认是0，可以不用设置
DB int

// redis操作失败最大重试次数，默认不重试。
MaxRetries int

// 最小重试时间间隔.
// 默认是 8ms ; -1 表示关闭.
MinRetryBackoff time.Duration

// 最大重试时间间隔
// 默认是 512ms; -1 表示关闭.
MaxRetryBackoff time.Duration

// redis连接超时时间.
// 默认是 5 秒.
DialTimeout time.Duration
// socket读取超时时间
// 默认 3 秒.
ReadTimeout time.Duration

// socket写超时时间
WriteTimeout time.Duration

// redis连接池的最大连接数.
// 默认连接池大小等于 cpu个数 * 10
PoolSize int

// redis连接池最小空闲连接数.
MinIdleConns int
// redis连接最大的存活时间，默认不会关闭过时的连接.
MaxConnAge time.Duration

// 当你从redis连接池获取一个连接之后，连接池最多等待这个拿出去的连接多长时间。
// 默认是等待 ReadTimeout + 1 秒.
PoolTimeout time.Duration
// redis连接池多久会关闭一个空闲连接.
// 默认是 5 分钟. -1 则表示关闭这个配置项
IdleTimeout time.Duration
// 多长时间检测一下，空闲连接
// 默认是 1 分钟. -1 表示关闭空闲连接检测
IdleCheckFrequency time.Duration

// 只读设置，如果设置为true， redis只能查询缓存不能更新。
readOnly bool
}
```

string操作

```go
type Cmdable interface {
    //给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
    Set(key string, value interface{}, expiration time.Duration) *StatusCmd
    //查询数据库中名称为key的value值
    Get(key string) *StringCmd
    //设置一个key的值，并返回这个key的旧值
    GetSet(key string, value interface{}) *StringCmd
    //如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
    SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd
    //批量查询key的值。比如redisDb.MGet("name1","name2","name3")
    MGet(keys ...string) *SliceCmd
    //批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
    MSet(pairs ...interface{}) *StatusCmd
    //Incr函数每次加一,key对应的值必须是整数或nil
    //否则会报错incr key1: ERR value is not an integer or out of range
    Incr(key string) *IntCmd
    // IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
    IncrBy(key string, value int64) *IntCmd
    // IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
    IncrByFloat(key string, value float64) *FloatCmd
    // Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
    Decr(key string) *IntCmd
    //DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
 DecrBy(key string, decrement int64) *IntCmd
    //删除key操作,支持批量删除 redisDb.Del("key1","key2","key3")
    Del(keys ...string) *IntCmd
    //设置key的过期时间,单位秒
    Expire(key string, expiration time.Duration) *BoolCmd
    //给数据库中名称为key的string值追加value
    Append(key, value string) *IntCmd
}

```

list操作

```go
type Cmdable interface {
    //从列表左边插入数据,list不存在则新建一个继续插入数据
 LPush(key string, values ...interface{}) *IntCmd
    //跟LPush的区别是，仅当列表存在的时候才插入数据
 LPushX(key string, value interface{}) *IntCmd
    //返回名称为 key 的 list 中 start 至 end 之间的元素
    //返回从0开始到-1位置之间的数据，意思就是返回全部数据
 LRange(key string, start, stop int64) *StringSliceCmd
    //返回列表的长度大小
 LLen(key string) *IntCmd
    //截取名称为key的list的数据，list的数据为截取后的值
 LTrim(key string, start, stop int64) *StatusCmd
    //根据索引坐标，查询列表中的数据
    LIndex(key string, index int64) *StringCmd
    //给名称为key的list中index位置的元素赋值
 LSet(key string, index int64, value interface{}) *StatusCmd
    //在指定位置插入数据。op为"after或者before"
 LInsert(key, op string, pivot, value interface{}) *IntCmd
    //在指定位置前面插入数据
 LInsertBefore(key string, pivot, value interface{}) *IntCmd
    //在指定位置后面插入数据
 LInsertAfter(key string, pivot, value interface{}) *IntCmd
    //从列表左边删除第一个数据，并返回删除的数据
 LPop(key string) *StringCmd
    //删除列表中的数据。删除count个key的list中值为value 的元素。
 LRem(key string, count int64, value interface{}) *IntCmd
   }
```

集合操作

```go
type Cmdable interface {
    //向名称为key的set中添加元素member
    SAdd(key string, members ...interface{}) *IntCmd
    //获取集合set元素个数
 SCard(key string) *IntCmd
    //判断元素member是否在集合set中
 SIsMember(key string, member interface{}) *BoolCmd
    //返回名称为 key 的 set 的所有元素
 SMembers(key string) *StringSliceCmd
    //求差集
 SDiff(keys ...string) *StringSliceCmd
    //求差集并将差集保存到 destination 的集合
 SDiffStore(destination string, keys ...string) *IntCmd
    //求交集
 SInter(keys ...string) *StringSliceCmd
    //求交集并将交集保存到 destination 的集合
 SInterStore(destination string, keys ...string) *IntCmd
    //求并集
 SUnion(keys ...string) *StringSliceCmd
    //求并集并将并集保存到 destination 的集合
 SUnionStore(destination string, keys ...string) *IntCmd
    //随机返回集合中的一个元素，并且删除这个元素
 SPop(key string) *StringCmd
    // 随机返回集合中的count个元素，并且删除这些元素
 SPopN(key string, count int64) *StringSliceCmd
    //删除名称为 key 的 set 中的元素 member,并返回删除的元素个数
 SRem(key string, members ...interface{}) *IntCmd
    //随机返回名称为 key 的 set 的一个元素
 SRandMember(key string) *StringCmd
    //随机返回名称为 key 的 set 的count个元素
 SRandMemberN(key string, count int64) *StringSliceCmd
    //把集合里的元素转换成map的key
 SMembersMap(key string) *StringStructMapCmd
    //移动集合source中的一个member元素到集合destination中去
 SMove(source, destination string, member interface{}) *BoolCmd
}
```

hash操作

```go
type Cmdable interface {
    //根据key和字段名，删除hash字段，支持批量删除hash字段
    HDel(key string, fields ...string) *IntCmd
    //检测hash字段名是否存在。
 HExists(key, field string) *BoolCmd
    //根据key和field字段，查询field字段的值
 HGet(key, field string) *StringCmd
    //根据key查询所有字段和值
 HGetAll(key string) *StringStringMapCmd
    //根据key和field字段，累加数值。
 HIncrBy(key, field string, incr int64) *IntCmd
    //根据key和field字段，累加数值。
 HIncrByFloat(key, field string, incr float64) *FloatCmd
    //根据key返回所有字段名
 HKeys(key string) *StringSliceCmd
    //根据key，查询hash的字段数量
 HLen(key string) *IntCmd
    //根据key和多个字段名，批量查询多个hash字段值
 HMGet(key string, fields ...string) *SliceCmd
    //根据key和多个字段名和字段值，批量设置hash字段值
 HMSet(key string, fields map[string]interface{}) *StatusCmd
    //根据key和field字段设置，field字段的值
 HSet(key, field string, value interface{}) *BoolCmd
    //根据key和field字段，查询field字段的值
 HSetNX(key, field string, value interface{}) *BoolCmd
}
```

有序集合操作

```go
type Cmdable interface {
    // 添加一个或者多个元素到集合，如果元素已经存在则更新分数
    ZAdd(key string, members ...Z) *IntCmd
 ZAddNX(key string, members ...Z) *IntCmd
 ZAddXX(key string, members ...Z) *IntCmd
 ZAddCh(key string, members ...Z) *IntCmd
 ZAddNXCh(key string, members ...Z) *IntCmd
    // 添加一个或者多个元素到集合，如果元素已经存在则更新分数
 ZAddXXCh(key string, members ...Z) *IntCmd
    //增加元素的分数
 ZIncr(key string, member Z) *FloatCmd
 ZIncrNX(key string, member Z) *FloatCmd
 ZIncrXX(key string, member Z) *FloatCmd
    //增加元素的分数，增加的分数必须是float64类型
 ZIncrBy(key string, increment float64, member string) *FloatCmd
    // 存储增加分数的元素到destination集合
 ZInterStore(destination string, store ZStore, keys ...string) *IntCmd
    //返回集合元素个数
    ZCard(key string) *IntCmd
    //统计某个分数范围内的元素个数
 ZCount(key, min, max string) *IntCmd
    //返回集合中某个索引范围的元素，根据分数从小到大排序
 ZRange(key string, start, stop int64) *StringSliceCmd
    //ZRevRange的结果是按分数从大到小排序。
    ZRevRange(key string, start, stop int64) *StringSliceCmd
 //根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
 ZRangeByScore(key string, opt ZRangeBy) *StringSliceCmd
    //根据分数范围返回集合元素，用法类似ZRangeByScore，区别是元素根据分数从大到小排序。
    ZRemRangeByScore(key, min, max string) *IntCmd
    //用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
    ZRangeWithScores(key string, start, stop int64) *ZSliceCmd
    //根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
 ZRank(key, member string) *IntCmd
    //ZRevRank的作用跟ZRank一样，区别是ZRevRank是按分数从大到小排序。
    ZRevRank(key, member string) *IntCmd 
    //查询元素对应的分数
 ZScore(key, member string) *FloatCmd
    //删除集合元素
 ZRem(key string, members ...interface{}) *IntCmd
    //根据索引范围删除元素。从最低分到高分的（stop-start）个元素
 ZRemRangeByRank(key string, start, stop int64) *IntCmd
}
```

事物操作

```go
type Pipeliner interface {
 StatefulCmdable
 Do(args ...interface{}) *Cmd
 Process(cmd Cmder) error
 Close() error
 Discard() error
 Exec() ([]Cmder, error)
}
```

事物的常用函数
//以Pipeline的方式操作事务
TxPipeline() Pipeliner
Watch - redis乐观锁支持

事物的示例:

```go
func main() {
 err := initClient()
 if err != nil {
  //redis连接错误
  panic(err)
 }
 //统计开发语言排行榜
 zsetKey := "language_rank"
 // 开启一个TxPipeline事务
 pipe := redisDb.TxPipeline()

 // 执行事务操作，可以通过pipe读写redis
 incr := pipe.Incr(zsetKey)
 pipe.Expire(zsetKey, time.Hour)

 // 通过Exec函数提交redis事务
 _, err = pipe.Exec()

 // 提交事务后，我们可以查询事务操作的结果
 // 前面执行Incr函数，在没有执行exec函数之前，实际上还没开始运行。
 fmt.Println(incr.Val(), err)
}

```

watch 操作

```go
func main() {
 err := initClient()
 if err != nil {
  //redis连接错误
  panic(err)
 }
 // 定义一个回调函数，用于处理事务逻辑
 fn := func(tx *redis.Tx) error {
  // 先查询下当前watch监听的key的值
  v, err := tx.Get("key").Result()
  if err != nil && err != redis.Nil {
   return err
  }

  // 这里可以处理业务
  fmt.Println(v)

  // 如果key的值没有改变的话，Pipelined函数才会调用成功
  _, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
   // 在这里给key设置最新值
   pipe.Set("key", "new value", 0)
   return nil
  })
  return err
 }

 // 使用Watch监听一些Key, 同时绑定一个回调函数fn, 监听Key后的逻辑写在fn这个回调函数里面
 // 如果想监听多个key，可以这么写：client.Watch(fn, "key1", "key2", "key3")
 redisDb.Watch(fn, "key")
}
```

具体的示例可以查看:
<https://juejin.cn/post/7027347979065360392>

go-redis文档:
<https://godoc.org/github.com/go-redis/redis>
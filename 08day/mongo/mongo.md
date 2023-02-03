# mongo

这里用docker 玩一玩算了
装起来真麻烦,有点哈批

```docker
// 拉取镜像
docker pull mongo:4.2.5

// 启动容器 这里可以给mongoDB设置一个用于连接的账号
docker run -p 27017:27017 --name mongo \
-v /mydata/mongo/db:/data/db \
-d mongo:4.2.5 --auth

// 进入容器
docker exec -it mongo mongo

// 在admin集合中创建账号
use admin
db.createUser({ 
    user: 'hsiaocz', 
    pwd: 'shaw123', 
    roles: [ { role: "root", db: "admin" } ] });

// 创建完成之后验证是否可以登录
db.auth("hsiaocz","shaw123")

//这里登录成功，会报个1

```

这里每次进去，先使用admin,然后使用账号登录，就可以正常操作了

```mongo
use admin
db.auth("hsiaocz","shaw123")
show dbs
```

mongoBD的一些概念
database 数据库  相当于mysql的database
collection 集合  相当于mysql的table
document   文档  相当于mysql的row 一行 也叫记录
field      域    相当于mysql数据字段column
index     索引   相当于mysql的索引
primary key  主键
mongodb会自动的将_id字段设置成主键

## 1、数据库操作

```mongo
// 创建数据库 使用use创建 插入一条数据时创建
use test;

// 查看数据库
show dbs

// 删除数据库
db.dropDatabase()
```

## 2、集合操作

所谓的集合就相当于mysql里的表

有一点要注意，集合操作基于数据库对象

```mongo
// 创建集合
use test
db.createCollection("article")

// 查看集合
show collections

// 删除集合
db.article.drop()
```

## 3、文档操作

文档操作基于集合对象

```mongo
// 插入文档
db.collection.insert(document)
例子：
db.article.insert({title:"mongo",description:"mongodb 是一个Nosql数据库",by:'bob',url:"http://www.mongo.com",tags:["mongo","mysql","nosql"],likes:100})

// 获取文档
db.article.find()

// 更新文档
db.collection.update(
   <query>,
   <update>,
   {
     multi: <boolean>
   }
)
# query：修改的查询条件，类似于SQL中的WHERE部分
# update：更新属性的操作符，类似与SQL中的SET部分
# multi：设置为true时会更新所有符合条件的文档，默认为false只更新找到的第一条

将article集合中所有title为mongo 教程的改为mongoDB
db.article.update({'title':"mongo 教程"},{$set:{"title":"mongoDB"}},{multi:true})

使用update()方法可以替换已有的文档
db.collection.save(document)

save方法要传递一个文档
db.article.save({
    "_id" : ObjectId("5e9943661379a112845e4056"),
    "title" : "MongoDB 教程",
    "description" : "MongoDB 是一个 Nosql 数据库",
    "by" : "Andy",
    "url" : "https://www.mongodb.com/",
    "tags" : [ 
        "mongodb", 
        "database", 
        "NoSQL"
    ],
    "likes" : 100.0
})

//删除文档 使用remove方法
db.collection.remove(
   <query>,
   {
     justOne: <boolean>
   }
)
# query：删除的查询条件，类似于SQL中的WHERE部分
# justOne：设置为true只删除一条记录，默认为false删除所有记录

// 删除article中集合为MongoDB的所有文档
db.article.remove({'title':'mongoDB'})

// 查询文档
db.collection.find(query, projection)
# query：查询条件，类似于SQL中的WHERE部分
# projection：可选，使用投影操作符指定返回的键

使用db.collection.find()方法可以返回所有的值，表示查询所有


```

### 3.1、关于查询操作中的条件查询

操作
格式
SQL中的类似语句

等于
{<key>:<value>}
where title = 'MongoDB 教程'

小于
{<key>:{$lt:<value>}}
where likes < 50

小于或等于
{<key>:{$lte:<value>}}
where likes <= 50

大于
{<key>:{$gt:<value>}}
where likes > 50

大于或等于
{<key>:{$gte:<value>}}
where likes >= 50

不等于
{<key>:{$ne:<value>}}
where likes != 50

通过$符号指定，前面的是key，后面的是条件

例子：

```mongo

// 查询title为mongoDB教程的所有文档
db.article.find({'title':'mongoDB 教程'})

// 条件查询，查询likes大于50的所有文章
db.article.find({'likes':{$gt:50}})

// 如果是and条件，可以在find后面更多个条件，使用,隔开来实现
db.article.find({'title':"mongoDB 教程",'by':'alex'})

// or条件使用$or操作符实现,查询title为Redis教程或者MongoDB教程的所有文章
db.article.find({$or:[{'title':'redis 教程'},{'title':'mongo 教程'}]})

//AND 和 OR条件的联合使用，例如查询likes大于50，并且title为Redis 教程或者"MongoDB 教程的所有文档。
db.article.find({"likes": {$gt:50}, $or: [{"title": "Redis 教程"},{"title": "MongoDB 教程"}]})
```

### 3.2、一些其他的操作

**limit和Skip操作**
1-读取指定数量的文档,可以使用limit()方法
db.collection.find().limit(number)

例如只查询集合中的两条数据
db.article.find().limit(2)

跳过指定数量的文档来读取，可以使用skip()方法，语法如下
db.collection.find().limmit(number).skip(number)

从第二条开始,查询artilce集合中的两条数据
db.article.find().limit(2).skip(1)

**排序方法**
使用sort()方法对数据进行排序sort()方法通过参数指定排序的字段，并使用1和-1来指定排序方式,1为升序，-1为降序

`db.collection.find().sort({KEY:1})`

按照article集合中文档的likes字段降序排列
`db.article.find().sort({likes:-1})`

**索引**
索引通常能够极大的提高查询的效率，如果没有索引，MongoDB在读取数据时必须扫描集合中的每个文件并选取那些符合条件的条件

mongo使用createIndex()方法创建索引，语法格式：

```mongo
db.collection.createIndex(keys, options)
# background：建索引过程会阻塞其它数据库操作，设置为true表示后台创建，默认为false
# unique：设置为true表示创建唯一索引
# name：指定索引名称，如果没有指定会自动生成

给title和description字段创建索引，1表示升序索引，-1表示降序索引，指定以后台方式创建
db.article.createIndex({"title":1,"description":-1}, {background: true})

查看集合中已经创建的索引
db.article.getIndexes()

```

**分组操作**
mongoDB的分组使用aggregate()方法
`db.collection.aggregate(AGGREGATE_OPERATION)`

聚合操作中常用的操作符如下:
`$sum`  计算总和
`$avg`  计算平均值
`$min`  计算最小值
`$max`  计算最大值

根据by字段聚合文档并计算文档数量,类似于sql中的count()函数
`db.article.aggregate([{$group:{_id:"$by",sum_count:{"$sum:1"}}}])`

根据by字段聚合文档并计算likes字段的平局值，类似与SQL中的avg()语句:
`db.article.aggregate([{$group : {_id : "$by", avg_likes : {$avg : "$likes"}}}])`

正则表达式:
mongoDB使用$regex操作符来设置匹配字符串的正则表达式，可以用来模糊查询，类似于SQL中的like操作

例如查询title中包含教程的文档
`db.article.find({title:{$regex:"教程"}})`
不区分大小写的模糊查询，使用$options
`db.article.find({title:{$regex:'elasticsearch',$options:'$i'}})`

关于文档部分，更多的看这
[https://www.mongodb.com/docs/manual/]官方文档

## 4、看看go操作mongoDB

首先当然是库

```go
go get github.com/mongodb/mongo-go-driver
```

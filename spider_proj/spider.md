# 爬虫项目

- 通用爬虫
- 聚焦爬虫

- go 语言爬虫库/框架

- 使用 ElasticSearch 作为数据存储
- 使用 Go 语言 标准模板库实现 http 数据展示部分

爬虫的主题

- 爬取内容

## 14.2 正则表达式

### 获取城市名称和链接

- 使用 css 选择器
- 使用 xpath
- 使用正则表达式

- [随机伪造 User-Agent](https://blog.csdn.net/tyBaoErGe/article/details/50375802)

## Elastic search 
```sh
docker pull elasticsearch:5.6.16
docker run -d -p 9200:9200  elasticsearch:5.6.16
```
``` sh
go get gopkg.in/olivere/elastic.v5
```
```go
import elastic "gopkg.in/olivere/elastic.v5"
```
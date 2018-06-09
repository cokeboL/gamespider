编译：
	gamespider放到GOPATH/src下
	go build -o gamespider.bin gamespider
	./gamespider.bin

配置：
	需要爬取的配置文件放到 res/src 目录下，会遍历该目录下的文件进行解析爬取，只遍历一层

结果：
	爬取结果存放在 res/dst 目录中，每个配置文件对应一个相应名字目录

参数：
	每个 url 最多尝试次数：3
	每个 url 爬取超市时间：time.Second * 10
	每个 url 每次爬取延时：time.Second / 10
	所有 url 爬取起始延时：time.Second / 10，每增加一个 url，此值会加上每次爬取的延时控制爬取速度

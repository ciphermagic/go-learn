module ciphermagic.cn

go 1.12

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/olivere/elastic/v7 v7.0.15
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20190311212946-11955173bddd
)

replace github.com/olivere/elastic/v7 => gopkg.in/olivere/elastic.v7 v7.0.15

module ciphermagic.cn

go 1.12

require (
	fyne.io/fyne/v2 v2.3.1
	github.com/facebookarchive/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.5.0
	github.com/olivere/elastic/v7 v7.0.15
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4
	golang.org/x/text v0.3.7
	golang.org/x/tools v0.1.8-0.20211022200916-316ba0b74098
)

replace github.com/olivere/elastic/v7 => gopkg.in/olivere/elastic.v7 v7.0.15

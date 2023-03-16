module ciphermagic.cn

go 1.12

require (
	fyne.io/fyne/v2 v2.3.1
	github.com/BurntSushi/toml v1.2.1
	github.com/bytedance/gopkg v0.0.0-20230224073017-0b6876860a2f
	github.com/facebookarchive/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.5.0
	github.com/olivere/elastic/v7 v7.0.15
	golang.org/x/net v0.7.0
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	golang.org/x/text v0.7.0
	golang.org/x/tools v0.1.12
)

replace github.com/olivere/elastic/v7 => gopkg.in/olivere/elastic.v7 v7.0.15

package mysql

import (
	docker "drone/drone-docker"
	"testing"
)

func TestTagInset(t *testing.T) {
	tagStoreMysql,_ := docker.InitTagStore("mysql",
		docker.WithUrl("root:5ziEppim@tcp(mysql-2580-0.tripanels.com:2580)/tags?charset=utf8"),
	)

	tagStoreMysql.TagInset()

}
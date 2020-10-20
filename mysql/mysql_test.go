package mysql

import (
	docker "drone/drone-docker"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestTagInset(t *testing.T) {
	tagStoreMysql,err := docker.InitTagStore("mysql",
		docker.WithUrl("root:5ziEppim@tcp(mysql-2580-0.tripanels.com:2580)/tags?charset=utf8"),
	)
	if err != nil {
		logrus.Fatal(err)
		//panic("init registry failed")
	}

	tagStoreMysql.TagInset()

}
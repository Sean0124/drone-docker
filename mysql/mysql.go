package mysql

import (
	"database/sql"
	docker "drone/drone-docker"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type TagStoreMysql struct {
	options *docker.Options
	client  *sql.DB
}

func init() {
	var mysqlinfo *TagStoreMysql = &TagStoreMysql{
		options: nil,
		client:  nil,
	}

	docker.RegisterTagStorePlugin(mysqlinfo)
}

func (m *TagStoreMysql) Name() string {
	return "mysql"
}

func (m *TagStoreMysql) Init(opts ...docker.Option) (err error) {
	m.options = &docker.Options{}
	for _, opt := range opts {
		opt(m.options)
	}
	m.client, err = sql.Open("mysql", m.options.Url)

	if err != nil {
		err = fmt.Errorf("init mysqlstroe failed, err:%v", err)
		return
	}

	return
}

func (m *TagStoreMysql) TagInset() {

	stmt, err := m.client.Prepare("INSERT drone SET DRONE_REPO=?,DRONE_BRANCH=?,TAG=?")
	checkErr(err)
	DRONE_REPO := os.Getenv("DRONE_REPO")
	DRONE_BRANCH := os.Getenv("DRONE_BRANCH")
	_, err = stmt.Exec(DRONE_REPO, DRONE_BRANCH, "0.0.0")
	checkErr(err)
}

func (m *TagStoreMysql) TagUpdate(tag string) {
	stmt, err := m.client.Prepare("update drone set TAG=? where DRONE_REPO=? and DRONE_BRANCH=?")
	checkErr(err)

	DRONE_REPO := os.Getenv("DRONE_REPO")
	DRONE_BRANCH := os.Getenv("DRONE_BRANCH")
	//fmt.Println("MysqlUpdate tag:", tag)
	logrus.Printf("MysqlUpdate tag: %s", tag)
	_, err = stmt.Exec(tag, DRONE_REPO, DRONE_BRANCH)
	checkErr(err)
}

func (m *TagStoreMysql) TagFind() (tag string) {
	DRONE_REPO := os.Getenv("DRONE_REPO")
	DRONE_BRANCH := os.Getenv("DRONE_BRANCH")
	//DRONE_REPO := "cloudcdlusters-websites/cloudclusters"
	//DRONE_BRANCH := "devedlop"
	m.client.QueryRow("SELECT TAG FROM drone where DRONE_REPO=? and DRONE_BRANCH=?", DRONE_REPO, DRONE_BRANCH).Scan(&tag)
	return tag
}

func checkErr(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

package docker

import (
	"database/sql"
	"fmt"
	"github.com/coreos/go-semver/semver"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

// DefaultTagSuffix returns a set of default suggested tags
// based on the commit ref with an attached suffix.
func DefaultTagSuffix(ref, suffix string) []string {
	tags := DefaultTags(ref)
	if len(suffix) == 0 {
		return tags
	}
	for i, tag := range tags {
		if tag == "latest" {
			tags[i] = suffix
		} else {
			tags[i] = fmt.Sprintf("%s-%s", tag, suffix)
		}
	}
	return tags
}

func splitOff(input string, delim string) string {
	parts := strings.SplitN(input, delim, 2)

	if len(parts) == 2 {
		return parts[0]
	}

	return input
}

// DefaultTags returns a set of default suggested tags based on
// the commit ref.
func DefaultTags(ref string) []string {
	if !strings.HasPrefix(ref, "refs/tags/") {
		return []string{"latest"}
	}
	v := stripTagPrefix(ref)
	version, err := semver.NewVersion(v)
	if err != nil {
		return []string{"latest"}
	}
	if version.PreRelease != "" || version.Metadata != "" {
		return []string{
			version.String(),
		}
	}

	v = stripTagPrefix(ref)
	v = splitOff(splitOff(v, "+"), "-")
	dotParts := strings.SplitN(v, ".", 3)

	if version.Major == 0 {
		return []string{
			fmt.Sprintf("%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor),
			fmt.Sprintf("%0*d.%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor, len(dotParts[2]), version.Patch),
		}
	}
	return []string{
		fmt.Sprintf("%0*d", len(dotParts[0]), version.Major),
		fmt.Sprintf("%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor),
		fmt.Sprintf("%0*d.%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor, len(dotParts[2]), version.Patch),
	}
}

// UseDefaultTag for keep only default branch for latest tag
func UseDefaultTag(ref, defaultBranch string) bool {
	if strings.HasPrefix(ref, "refs/tags/") {
		return true
	}
	if stripHeadPrefix(ref) == defaultBranch {
		return true
	}
	return false
}

func stripHeadPrefix(ref string) string {
	return strings.TrimPrefix(ref, "refs/heads/")
}

func stripTagPrefix(ref string) string {
	ref = strings.TrimPrefix(ref, "refs/tags/")
	ref = strings.TrimPrefix(ref, "v")
	return ref
}

func MysqlCont() *sql.DB {
	db, err := sql.Open("mysql", "root:5ziEppim@tcp(mysql-2580-0.tripanels.com:2580)/tags?charset=utf8")
	checkErr(err)

	return db
}

func MysqlInset(db *sql.DB) {
	stmt, err := db.Prepare("INSERT drone SET DRONE_REPO=?,DRONE_BRANCH=?,TAG=?")
	checkErr(err)
	DRONE_REPO := os.Getenv("DRONE_REPO")
	DRONE_BRANCH := os.Getenv("DRONE_BRANCH")
	_, err = stmt.Exec(DRONE_REPO, DRONE_BRANCH, "0.0.1")
	checkErr(err)

}

func MysqlUpdate(db *sql.DB, tag string) {
	stmt, err := db.Prepare("update drone set TAG=? where DRONE_REPO=? and DRONE_BRANCH=?")
	checkErr(err)

	DRONE_REPO := os.Getenv("DRONE_REPO")
	DRONE_BRANCH := os.Getenv("DRONE_BRANCH")
	fmt.Println("MysqlUpdate tag:", tag)
	_, err = stmt.Exec(tag, DRONE_REPO, DRONE_BRANCH)
	checkErr(err)
}

func MysqlFind(db *sql.DB) (TAG string) {
	DRONE_REPO := os.Getenv("DRONE_REPO")
	DRONE_BRANCH := os.Getenv("DRONE_BRANCH")
	//DRONE_REPO := "cloudcdlusters-websites/cloudclusters"
	//DRONE_BRANCH := "devedlop"
	db.QueryRow("SELECT TAG FROM drone where DRONE_REPO=? and DRONE_BRANCH=?", DRONE_REPO, DRONE_BRANCH).Scan(&TAG)
	return TAG

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TagTemplateInit(tagTemplate string) (string,error) {
	//**-d.d.d-**
	dockerTag, err := semver.NewVersion(tagTemplate)
	if err !=nil {
		fmt.Println(err)
		return "",err
	}
	dockerTag.Minor=0
	dockerTag.Patch=0
	dockerTag.Major=0

	return  dockerTag.String(),nil
}

func TagTemplateParse(tagTemplate string) (*semver.Version,error) {
	return semver.NewVersion(tagTemplate)
}


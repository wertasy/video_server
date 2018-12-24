package dbops

import (
	"testing"
	"time"
	"strconv"
	"fmt"
)

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}
func TestVideo(t *testing.T) {
	t.Run("Add", testAddNewVideo)
	t.Run("Get", testGetVideoInfo)
	t.Run("Delete", testDeleteVideoInfo)
	t.Run("Reget", testRegetVideoInfo)
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComment", testAddComment)
	t.Run("ListComments", testListComments)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("wert", "1234")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("wert")
	if pwd != "1234" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("wert", "1234")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("wert")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting User test failed")
	}
}
func testAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "video1")
	if err != nil {
		t.Errorf("Error of AddNewVideo: %v", err)
	}
	if video == nil {
		t.Errorf("Adding Video test failed")
	} else {
		tempvid = video.Id
	}
}

func testGetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	} else {
		tempvid = video.Id
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempvid)
	if err != nil || video != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func testAddComment(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddComment(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComment: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v\n", i, ele)
	}
}

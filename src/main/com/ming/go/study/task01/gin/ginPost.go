package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Teacher struct {
	TID      string
	TName    string
	Age      int
	PhoneNum string
}

func main() {

	t := Teacher{
		TID:      "T000001",
		TName:    "张老师",
		Age:      25,
		PhoneNum: "1872313134",
	}

	tJson, err := json.Marshal(t)
	if err != nil {
		fmt.Println("转换失败")
	}

	url := "http://localhost:8000/api/gin/testInfo"
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(tJson))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "ELADMIN-TOEKN=Bearer%20eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiI0ZWYxMzU0M2FiMWU0MTBiYTVlMWFmMjYwYmE3OWU0NiIsInVzZXIiOiJhZG1pbiIsInN1YiI6ImFkbWluIn0.YbD-7JqiWA1BZfSok6Q4x7E7YGzVWp3UlCUNY0YBPNf6_gQeEj9Tuy2CXwZGJAa2ywFq8qWbQ7qkNzmdDZtkJA")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiI0ZWYxMzU0M2FiMWU0MTBiYTVlMWFmMjYwYmE3OWU0NiIsInVzZXIiOiJhZG1pbiIsInN1YiI6ImFkbWluIn0.YbD-7JqiWA1BZfSok6Q4x7E7YGzVWp3UlCUNY0YBPNf6_gQeEj9Tuy2CXwZGJAa2ywFq8qWbQ7qkNzmdDZtkJA")
	resp, rer := client.Do(req)
	defer resp.Body.Close()
	if rer != nil {
		fmt.Println("请求失败: ", rer)
	}
	body, ber := io.ReadAll(resp.Body)
	if ber != nil {
		fmt.Println("返回结果body读取失败：", ber)
	}
	fmt.Println("请求成功，服务端返回结果：", string(body))
}

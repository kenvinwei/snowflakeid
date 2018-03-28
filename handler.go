package snowflakeid

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type resp struct {
	Status string `json:"status"`
	Data   []ID   `json:"data"`
}

func getUniqueId(c echo.Context) error {

	n, _ := strconv.Atoi(c.Param("num"))

	r := &resp{}

	r.Status = "ok"
	r.Data = MakeID(n)
	respData, _ := json.Marshal(r)
	log.Print(string(respData))
	return c.String(http.StatusOK, string(respData))
}

func MakeID(num int) []ID {
	req := []ID{}

	ch := make(chan ID)

	node, err := NewNode(InitNode)

	m := make(map[ID]int)

	if err != nil {
		// err code ...
	}

	for i := 0; i < num; i++ {
		go func() {
			id := node.Generate()
			ch <- id
		}()
	}

	defer close(ch)

	for i := 0; i < num; i++ {
		id := <-ch

		_, ok := m[id]
		if ok {
			fmt.Printf("ID is not unique!\n")
			break
		}
		// 将 id 作为 key 存入 map
		m[id] = i

		// 将 id 作为 key 存入 map
		req = append(req, id)
	}
	//log.Print(req)
	return req

}

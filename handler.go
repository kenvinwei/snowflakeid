package snowflakeid

import (
	"encoding/json"
	"fmt"
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
	r.Data = makeID(n)

	jdata, _ := json.Marshal(r)

	fmt.Print("%v", r)

	return c.String(http.StatusOK, string(jdata))
}

func makeID(num int) []ID {
	req := []ID{}

	ch := make(chan ID)

	node, err := NewNode(InitNode)

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

		// 将 id 作为 key 存入 map
		req = append(req, id)
	}

	return req

}

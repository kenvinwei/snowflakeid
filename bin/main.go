package main

import (
	"fmt" // 测试、打印

	"snowflakeid"
)

func test() {
	// 生成节点实例
	node, err := snowflakeid.NewNode(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan snowflakeid.ID)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := node.Generate()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[snowflakeid.ID]int)
	for i := 0; i < count; i++ {
		id := <-ch
		fmt.Println(id)
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		fmt.Print(m)
		if ok {
			fmt.Printf("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All ", count, " snowflake ID generate successed!\n")
}

func main() {

	snowflakeid.InitNode = 1
	snowflakeid.Start()

	//For Test
	//fmt.Printf("%v", snowflakeid.MakeID(5))

}

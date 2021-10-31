package main

import (
	"go-err/dao"
	"log"
)

func main()  {
	defer func() {
		dao.Db.Close()
	}()


	err := dao.GetPerson()

	//业务层处理空数据，其他异常交给数据层处理
	if err != nil && dao.IsErrNoRows(err) {
		log.Println("GetPerson No Data")
	}
}

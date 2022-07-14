package main

import (
	"encoding/json"
	"fmt"
	jsonCore "main/core"
	"os"

	"github.com/kataras/iris/v12"
)

var base64 = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAAA+BAMAAACcvUmpAAAAG1BMVEXz+/5FM3nd4u3Hydycl7taTIlwZZqxsMyGfqqkyBVpAAAACXBIWXMAAA7EAAAOxAGVKw4bAAACE0lEQVRYhe2Wz0/CMBTHl45tXJ+6wXEoeMaf8QgmA48MjGemJHBk8QcenRrjn+1rO4jrdBSWejD9JiSs7d6n71dXw9DS0tL6LzrwlSNIAJ5ySgAAY8WMQ2TAjloGCSlkVy2kBX8AiaBmqg6XA9AmAAOlEAvqBnrSUQo5gpphA/hKIQkMEVRft6zcJsI3/K3N+znM2iUgc8OoAswL15AeFnkZCCoGr2h60aPtWivHIMVdUmXNCi+5iX1/A4hVHK2YMURfyTO6N5HHJFhb5OW3WewhD+0NBUbC0bIUtDIwI7HjV7Fw+sdmLu0pQ/7IwwPyJBK+JxchuNPVUwNgL/vOGQVM6HfIN6SUsHBkITY1sjIcimmv0OlXNlFc+0uZ3O+OSEaN+H8sLzc/6/VpWUpCLMa4E0Yr14tomYhYLHGb76velA5XnLqeUyW1beb2u8z6yIkkWzQSGA51igQjGiWWlZaYdocG6/6JBVHurMF9upffB2IMHa3QGxqQdBfjzCsWdy2UP84q4HYyA/juLenyeNTYAnD9zIq0oqVLC5PoZlcSVmkkYpAx80zsuAZLli3fJIYtuEwh3BDANH0WNmzRODphrkM3UMSMdpemzatQ/GrScnN/Pphl1UpLZ2WC5NL7yDO2vSPphbKwl0nAetHfHsIPJii+Wpx+zPolEKgm9cWVLc9tRR4+30teG7S0tLS0tMrrCyG6TICDHJeyAAAAAElFTkSuQmCC"

type Data struct {
	Key     string `json:"key"`
	Captcha string `json:"captcha"`
}
type Meta struct {
	Hide     json.Number `json:"hide"`
	Title    string      `json:"title"`
	Noslider bool        `json:"noslider"`
	Icon     string      `json:"icon"`
}

type Website struct {
	Path     string      `json:"path"`
	Id       json.Number `json:"id"`
	Meta     *Meta       `json:"meta"`
	Redirect string      `json:"redirect"`
	Children []Website   `json:"children"`
}

// iris 框架
func main() {

	filePtr, err := os.Open("./menu.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var info []Website
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}

	// 1 .创建 app结构体对象
	app := iris.Default()

	// 处理请求 验证码
	app.Get("/login/captcha", func(ctx iris.Context) {
		var mapslist map[string]string
		mapslist = map[string]string{"key": "0da8cd64-2a35-42e1-99f4-9776924a84da", "captcha": base64}

		ctx.JSON(jsonCore.GenSuccessData(mapslist))
	})

	app.Post("/login/login", func(ctx iris.Context) {
		var mapslist map[string]string
		mapslist = map[string]string{"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJqd3RfeWgiLCJleHAiOjE2NTMyNzU4MTQsInN1YiI6IllIIiwiYXVkIjoiZXZlcnkiLCJuYmYiOjE2NTI2NzEwMTQsImlhdCI6MTY1MjY3MTAxNCwianRpIjoxMDAwMSwidWlkIjoxfQ.uCuwaO1w86mDMWfU8MFr-HXX-WVbiRtyiFYU_Uxk3Us"}

		ctx.JSON(jsonCore.GenSuccessData(mapslist))
	})

	app.Get("/index/getUserInfo", func(ctx iris.Context) {
		var mapslist map[string]string
		mapslist = map[string]string{
			// "id":               1,
			"realname": "管理员",
			"nickname": "王能1",
			// "gender":           1,
			"avatar": "",
			"mobile": "13222222222",
			"email":  "222222@qq.com",
			// "birthday":         null,
			// "dept_id":          1,
			// "level_id":         2,
			// "position_id":      7,
			"province_code": "2791",
			"city_code":     "2792",
			"district_code": "2793",
			"address":       "江苏省",
			"city_name":     "",
			"username":      "admin",
			"password":      "062a8e08dfa50dccb57b7fbd37b48790",
			"salt":          "",
			"intro":         "个人简介",
			// "status":           1,
			"note": "暂无备注",
			// "sort":             124,
			// "login_num":        0,
			// "login_ip":         null,
			// "login_time":       0,
			// "create_user":      1,
			// "create_time":      1621998864,
			// "update_user":      1,
			// "update_time":      1651227746,
			// "mark":             1,
			"create_user_name": "管理员",
			"update_user_name": "管理员",
			"gender_name":      "男",
			// "city": [
			// 		"2791",
			// 		"2792",
			// 		"2793"
			// ],
			"level_name":    "总裁办",
			"position_name": "设备科长",
			"dept_name":     "陕西煤业股份有限公司",
			// "roles": [],
			// "authorities": [],
			// "permissionList": [	]

		}

		ctx.JSON(jsonCore.GenSuccessData(mapslist))
	})

	app.Get("/index/getMenuList", func(ctx iris.Context) {

		ctx.JSON(jsonCore.GenSuccessData(info))
	})

	app.Get("/login/logout", func(ctx iris.Context) {

		ctx.JSON(jsonCore.GenSuccessData('1'))
	})
	//  2.监听端口
	app.Run(iris.Addr(":8089"))
}

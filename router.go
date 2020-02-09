package beego2ts

import "github.com/astaxie/beego/swagger"

type Router struct {
	Name   string `json:"name" bson:"name"`
	Path   string `json:"path" bson:"_id"`
	Method string `json:"method" bson:"method"`
	Desc   string `json:"desc" bson:"desc"`
}

func newRouter(Name, Path, Method, Desc string) Router {
	return Router{
		Name:   Name,
		Path:   Path,
		Method: Method,
		Desc:   Desc,
	}
}

func getRouters(data swagger.Swagger) (route []Router) {
	for k, v := range data.Paths {
		if v.Get != nil {
			route = append(route, newRouter(v.Get.OperationID, data.BasePath+k, "get", v.Get.Description))
		}
		if v.Post != nil {
			route = append(route, newRouter(v.Post.OperationID, data.BasePath+k, "post", v.Post.Description))
		}
		if v.Put != nil {
			route = append(route, newRouter(v.Put.OperationID, data.BasePath+k, "put", v.Put.Description))
		}
		if v.Delete != nil {
			route = append(route, newRouter(v.Delete.OperationID, data.BasePath+k, "delete", v.Delete.Description))
		}
	}
	return
}

package beego2ts

import (
	"fmt"
	"github.com/astaxie/beego/swagger"
	"os"
	"regexp"
)

type Method string

const (
	GET    Method = "get"
	POST   Method = "post"
	PUT    Method = "put"
	DELETE Method = "delete"
)

func writeApiFile(filepath string, data swagger.Swagger) {

	api := []string{"import request, { query } from './request'", ""}
	for k, v := range data.Paths {
		if v.Get != nil {
			f := WriteFunction(v.Get.OperationID, data.BasePath+k, GET, v.Get.Description, v.Get.Parameters)
			api = append(api, f...)
		}
		if v.Post != nil {
			f := WriteFunction(v.Post.OperationID, data.BasePath+k, POST, v.Post.Description, v.Post.Parameters)
			api = append(api, f...)
		}
		if v.Put != nil {
			f := WriteFunction(v.Put.OperationID, data.BasePath+k, PUT, v.Put.Description, v.Put.Parameters)
			api = append(api, f...)
		}
		if v.Delete != nil {
			f := WriteFunction(v.Delete.OperationID, data.BasePath+k, DELETE, v.Delete.Description, v.Delete.Parameters)
			api = append(api, f...)
		}
	}
	f, err := os.Create(filepath + "/api.ts")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	for _, v := range api {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

const formSuffix = "Form"

func WriteFunction(name, url string, method Method, desc string, param []swagger.Parameter) (ss []string) {
	switch method {
	case GET:
		ss = createGetFunc(name, desc, url, param)
	case POST:
		ss = createPostFunc(name, desc, url, param)
	case PUT:
		ss = createPutFunc(name, desc, url, param)
	case DELETE:
		ss = createDeleteFunc(name, desc, url, param)
	}

	form := createInterface(name+formSuffix, param)
	ss = append(ss, form...)

	return
}

func createInterface(name string, param []swagger.Parameter) []string {
	if len(param) == 0 {
		return []string{}
	}
	if len(param) == 1 && param[0].In == "path" {
		return []string{}
	}
	d := []string{
		"export interface " + name + " {",
	}
	for _, v := range param {
		if v.In == "path" && v.Name == "id" {
			continue
		}
		s := "	" + v.Name
		if !v.Required {
			s = s + "?"
		}
		s = s + ": " + type2ts(v.Type)

		if v.Description != "" {
			s = s + " // " + v.Description
		}
		d = append(d, s)
	}
	d = append(d, "}")
	return d
}
func createGetFunc(name, desc, url string, param []swagger.Parameter) []string {
	var funcHeader = ""
	if len(param) == 0 {
		funcHeader = "export function " + name + "() {"
	} else if len(param) == 1 {
		if ok, _ := regexp.MatchString("/{", url); ok && param[0].In == "path" && param[0].Name == "id" {
			funcHeader = "export function " + name + "(id:" + type2ts(param[0].Type) + ") {"
			re, _ := regexp.Compile("/{")
			url = re.ReplaceAllString(url, "/${")
		} else {
			funcHeader = "export function " + name + "(params?: " + name + formSuffix + ") {"
			url = url + "?${query(params)}"
		}
	} else {
		funcHeader = "export function " + name + "(params?: " + name + formSuffix + ") {"
		url = url + "?${query(params)}"
		if ok, _ := regexp.MatchString("/{", url); ok {
			for _, v := range param {
				if v.In == "path" && v.Name == "id" {
					funcHeader = "export function " + name + "(id: string, params?: " + name + formSuffix + ") {"
					re, _ := regexp.Compile("/{")
					url = re.ReplaceAllString(url, "/${")
					break
				}
			}
		}

	}
	d := []string{
		"",
		"// " + desc,
		funcHeader,
		"	return request({",
		"		url: `" + url + "`,",
		"		method: 'get',",
		"	})",
		"}",
	}
	return d
}
func createPostFunc(name, desc, url string, param []swagger.Parameter) []string {
	var funcHeader = ""
	if len(param) == 0 {
		funcHeader = "export function " + name + "() {"
	} else {
		funcHeader = "export function " + name + "(params?: " + name + formSuffix + ") {"
	}
	d := []string{
		"",
		"// " + desc,
		funcHeader,
		"	return request({",
		"		url: `" + url + "`,",
		"		method: 'post',",
		"		data: query(params),",
		"	})",
		"}",
	}
	return d
}
func createPutFunc(name, desc, url string, param []swagger.Parameter) []string {
	var funcHeader = ""
	if len(param) == 0 {
		funcHeader = "export function " + name + "() {"
	} else {
		funcHeader = "export function " + name + "(params?: " + name + formSuffix + ") {"
	}
	d := []string{
		"",
		"// " + desc,
		funcHeader,
		"	return request({",
		"		url: `" + url + "`,",
		"		method: 'put',",
		"		data: query(params),",
		"	})",
		"}",
	}
	return d
}
func createDeleteFunc(name, desc, url string, param []swagger.Parameter) []string {
	var funcHeader = ""
	if len(param) == 0 {
		funcHeader = "export function " + name + "() {"
	} else if len(param) == 1 {
		if ok, _ := regexp.MatchString("/{", url); ok && param[0].In == "path" && param[0].Name == "id" {
			funcHeader = "export function " + name + "(id:" + type2ts(param[0].Type) + ") {"
			re, _ := regexp.Compile("/{")
			url = re.ReplaceAllString(url, "/${")
		} else {
			funcHeader = "export function " + name + "(params?: " + name + formSuffix + ") {"
			url = url + "?${query(params)}"
		}
	} else {
		funcHeader = "export function " + name + "(params?: " + name + formSuffix + ") {"
		url = url + "?${query(params)}"
		if ok, _ := regexp.MatchString("/{", url); ok {
			for _, v := range param {
				if v.In == "path" && v.Name == "id" {
					funcHeader = "export function " + name + "(id: string, params?: " + name + formSuffix + ") {"
					re, _ := regexp.Compile("/{")
					url = re.ReplaceAllString(url, "/${")
					break
				}
			}
		}

	}
	d := []string{
		"",
		"// " + desc,
		funcHeader,
		"	return request({",
		"		url: `" + url + "`,",
		"		method: 'delete',",
		"	})",
		"}",
	}
	return d
}

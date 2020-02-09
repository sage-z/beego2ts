package beego2ts

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path"
)

func BuildApi(outputDir string) {
	currpath, _ := os.Getwd()
	if !path.IsAbs(outputDir) {
		outputDir = path.Join(currpath, outputDir)
	}

	if !isExist(outputDir) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		createInitFile(outputDir)
		// todo 创建 git 仓库
	}

	data := GenerateDocs(currpath)
	writeApiFile(outputDir, data)
	//todo  修改版本 提交 推送
}

func BuildRouters() []Router {
	currpath, _ := os.Getwd()
	data := GenerateDocs(currpath)
	return getRouters(data)
}

func BuildSwagger(outputDir string) {
	currpath, _ := os.Getwd()
	if !path.IsAbs(outputDir) {
		outputDir = path.Join(currpath, outputDir)
	}

	data := GenerateDocs(currpath)

	os.Mkdir(path.Join(outputDir, "swagger"), 0755)
	fd, err := os.Create(path.Join(outputDir, "swagger", "swagger.json"))
	if err != nil {
		panic(err)
	}
	fdyml, err := os.Create(path.Join(outputDir, "swagger", "swagger.yml"))
	if err != nil {
		panic(err)
	}
	defer fdyml.Close()
	defer fd.Close()
	dt, err := json.MarshalIndent(data, "", "    ")
	dtyml, erryml := yaml.Marshal(data)
	if err != nil || erryml != nil {
		panic(err)
	}
	_, err = fd.Write(dt)
	_, erryml = fdyml.Write(dtyml)
	if err != nil || erryml != nil {
		panic(err)
	}
}

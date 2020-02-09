package beego2ts

import (
	"os"
)

func createInitFile(filepath string) {
	r, err := os.Create(filepath + "/request.js")
	if err != nil {
		panic(err)
	}
	p, err := os.Create(filepath + "/package.json")
	if err != nil {
		panic(err)
	}
	g, err := os.Create(filepath + "/.gitignore")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	defer p.Close()
	defer g.Close()

	_, err = r.Write([]byte(request))
	_, errp := p.Write([]byte(packagejson))
	_, errg := g.Write([]byte(gitignore))

	if err != nil || errp != nil || errg != nil {
		panic(err)
	}
}

const request = `import axios from 'axios'

// 创建 axios 实例
const service = axios.create({
  // baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // 超时时间
})

// 请求拦截器
service.interceptors.request.use(
  config => {

    // if (store.getters.token) {
    //   // let each request carry token
    //   // ['X-Token'] is a custom headers key
    //   // please modify it according to the actual situation
    //   config.headers['X-Token'] = Token.get()
    // }
    config.headers["Content-Type"] = "application/x-www-form-urlencoded;charset=utf-8"
    return config
  },
  error => {
    // 当请求错误的时候
    console.log(error)
    return Promise.reject(error)
  }
)

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data

    if (res.code !== 0) {

      return Promise.reject(new Error(res.msg || 'Error'))
    } else {
      return res.data
    }
  },
  error => {
    console.log('err ' + error) // for debug
    // notification['error']({
    //   message: 'Error',
    //   description: error+"" || 'Error',
    // });
    return Promise.reject(error)
  }
)

export default service

function query(object){
  let s = ""
  for (const key in object) {
    if (object.hasOwnProperty(key) && object[key] != undefined) {
      const element = object[key];
      s = s + "&" + key + '=' + element
    }
  }
  return s.slice(1)
}

export {
  query
}
`
const packagejson = `{
  "name": "api",
  "version": "0.0.0",
  "description": "接口SDK",
  "main": "index.js",
  "dependencies": {
    "tslib": "^1.10.0",
    "axios": "^0.19.0"
  },
  "author": "z",
  "license": "ISC"
}
`

const gitignore = `.DS_Store
node_modules
/dist

# local env files
.env.local
.env.*.local

# Log files
npm-debug.log*
yarn-debug.log*
yarn-error.log*

# Editor directories and files
.idea
.vscode
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?`

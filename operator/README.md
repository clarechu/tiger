# operator

quick start

##### 安装 operator-sdk
operator sdk 安装方法非常多，我们可以直接在 github 上面下载需要使用的版本，然后放置到 PATH 环境下面即可，当然也可以将源码 clone 到本地手动编译安装即可，如果你是 Mac，当然还可以使用常用的 brew 工具进行安装：


```bash
brew install operator-sdk

➜  tigger git:(master) ✗ operator-sdk version
operator-sdk version: "v0.16.0", commit: "55f1446c5f472e7d8e308dcdf36d0d7fc44fc4fd", go version: "go1.14 darwin/amd64"
    
```

##### 创建项目

```bash
operator-sdk new operator
```

#### 添加 API

上面我们添加自定义的 API，接下来可以添加对应的自定义 API 的具体实现 Controller，同样在项目根目录下面执行如下命令：


```bash
operator-sdk add controller --api-version=app.example.com/v1 --kind=AppService
```

打开源文件pkg/apis/app/v1/appservice_types.go，
需要我们根据我们的需求去自定义结构体 AppServiceSpec，我们最上面预定义的资源清单中就有 size、image、ports 这些属性，所有我们需要用到的属性都需要在这个结构体中进行定义：

```go
// AppServiceSpec defines the desired state of AppService
type AppServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Size      *int32 `json:"size"`
	Image     string `json:"image"`
	Resources string `json:"resources,omitempty"`
	Envs      string `json:"envs,omitempty"`
	Ports     string `json:"ports,omitempty"`
}
```

然后一个比较重要的结构体AppServiceStatus用来描述资源的状态，当然我们可以根据需要去自定义状态的描述，我这里就偷懒直接使用 Deployment 的状态了：

```go

// AppServiceStatus defines the observed state of AppService
type AppServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Status int `json:"status"`
}
```

定义完成后，在项目根目录下面执行如下命令：

```bash
$ operator-sdk generate k8s

```

改命令是用来根据我们自定义的 API 描述来自动生成一些代码，目录pkg/apis/app/v1/下面以zz_generated开头的文件就是自动生成的代码，里面的内容并不需要我们去手动编写。

这样我们就算完成了对自定义资源对象的 API 的声明。


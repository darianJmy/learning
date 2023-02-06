package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"path/filepath"
	"time"
)

// 定义一个 Controller 结构体
type Controller struct {
	indexer cache.Indexer
	queue workqueue.RateLimitingInterface
	informer cache.Controller
}



func NewController(indexer cache.Indexer,queue workqueue.RateLimitingInterface,informer cache.Controller) *Controller {
	return &Controller{
		indexer: indexer,
		queue: queue,
		informer: informer,
	}
}

// 携程执行初始化函数
func (c *Controller) Run(threadiness int, stopCh chan struct{}) {

	defer runtime.HandleCrash()
	// 关闭queue
	defer  c.queue.ShutDown()


	fmt.Printf("Start Custom Controller")

	//启动Informer
	go c.informer.Run(stopCh)

	//等待Informer刷新缓存
	if !cache.WaitForCacheSync(stopCh,c.informer.HasSynced) {
		fmt.Printf("Time out waiting caches to sync")
		return
	}

	//携程处理 queue 的程序数量
	for i:=0;i<threadiness;i++ {
		go wait.Until(c.runWorker,time.Second,stopCh)
	}


	//chan 使 run 卡在这
	<-stopCh

	fmt.Printf("Stop Custom Controller")
}

func (c *Controller) runWorker() {
	// 一直for 循环拿 queue 里面的 key，如果queue 数组没东西了，会卡住知道有数据进入 queue
	for c.ProcessItem(){}

}

// 获取 key
func (c *Controller) ProcessItem() bool{
	// 从 queue 中拿 key
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	// 函数执行结束删除这个 key
	defer c.queue.Done(key)

	//执行函数的具体处理功能，如果执行失败，就放进queue 等待下一次取出 queue执行
	if err := c.HanderObject(key.(string)); err != nil {
		if c.queue.NumRequeues(key) < 5 {
			c.queue.Add(key)
		}
	}

	return true
}


// 这边就是通过 queue 里面的 key 获取 indexer 里面的 object
func (c *Controller) HanderObject(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		fmt.Printf("This Object %s is error", key)
		return err
	}

	if !exists {
		fmt.Printf("This object %s does not found",key)
	} else {
		fmt.Printf(obj.(*v1.Pod).GetName(),obj.(*v1.Pod).GetNamespace())
		fmt.Printf("hello1")
	}

	return nil

}

func initClient() (*kubernetes.Clientset, error) {
	var err error
	var config *rest.Config

	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(可选) kubeconfig 文件的绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 文件的绝对路径")
	}
	flag.Parse()

	if config, err = rest.InClusterConfig(); err != nil {
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}
	return kubernetes.NewForConfig(config)
}


func main() {

	clientset, err := initClient()
	if err != nil {
		klog.Fatal(err)
	}

	//生成默认的 queue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())


	//生成特定资源与命名空间下的 listWatcher
	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(),"pods","default",fields.Everything())

	//生成indexer与informer
	indexer, informer := cache.NewIndexerInformer(podListWatcher,&v1.Pod{},0,cache.ResourceEventHandlerFuncs{
		// 资源事件处理函数
		// Add 后会被处理添加到 Queue
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			fmt.Printf("AddFunc %s",key)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			fmt.Printf("UpdateFunc %s",key)
			if err == nil {
				queue.Add(key)
			}

		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			fmt.Printf("DeleteFunc %s",key)
			if err == nil {
				queue.Add(key)
			}
		},
	},cache.Indexers{})


	//实例化
	controller := NewController(indexer,queue,informer)

	stopCh := make(chan struct{})
	//执行 Run
	go controller.Run(1,stopCh)

	defer close(stopCh)
	//select 卡住函数
	select {}
}


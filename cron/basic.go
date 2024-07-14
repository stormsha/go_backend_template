package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"reflect"
	"runtime"
	"stormsha.com/gbt/utils"
	"sync"
	"time"
)

var logger = utils.GetLogger()

// 秒级定时器，cron.New()默认是分钟级定时器
var crontab = cron.New(cron.WithSeconds())

func Cron() *cron.Cron {
	return crontab
}

func StartCron() {
	crontab.Start()
	fmt.Println("cron start.", time.Now())
}

func StopCron() {
	crontab.Stop()
	fmt.Println("cron stop.", time.Now())
}

// DeferFunc 避免函数抛异常
func DeferFunc(f func()) func() {
	return func() {
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("%v 执行异常 : %v", funcName, err)
			}
		}()
		logger.Infof("任务开始 : %v", funcName)
		f()
		logger.Infof("任务结束 : %v", funcName)
	}
}

var m = new(sync.Map)

// LockedFunc 避免函数抛异常+同时只执行一个任务
func LockedFunc(fs ...func()) func() {
	return func() {
		for _, f := range fs {
			func() {
				funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
				defer func() {
					if err := recover(); err != nil {
						logger.Errorf("%v 执行异常 : %v", funcName, err)
					}
				}()

				_, loaded := m.LoadOrStore(funcName, 1)
				if loaded {
					logger.Infof("任务跳过 : %v \n", funcName)
				} else {
					defer m.Delete(funcName)
					logger.Infof("任务开始 : %v\n", funcName)
					t1 := time.Now().UnixNano()
					f()
					t2 := time.Now().UnixNano()
					logger.Infof("任务结束 : %v, cost %v ms\n", funcName, (t2-t1)/1000000)
				}
			}()
		}
	}
}

func InitCorn() {
	// 用户手动调用初始化cron方法
	initArticleData()
	fmt.Println("___________Cron 初始化成功!___________")
}

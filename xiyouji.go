package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	// 限制 CPU 使用数，避免过载
	runtime.GOMAXPROCS(1)
	// 开启对锁调用的跟踪
	runtime.SetMutexProfileFraction(1)
	// 开启对阻塞操作的跟踪
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range AllPeoples {
			v.Live()
		}
		time.Sleep(time.Second)
	}
}

// 常量定义
const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
	Ti = Ki * Gi
	Pi = Ki * Ti
)

// 抽象人物
type People interface {
	Name() string
	Live()

	Eat()   // 吃
	Drink() // 喝
	Shit()  // 拉
	Sleep() // 睡
}

var (
	AllPeoples = []People{
		&WuKong{},     // 悟空
		&TangSeng{},   // 唐僧
		&ZhuBaJie{},   // 猪八戒
		&ShaHeShang{}, // 沙和尚
	}
)

// ========= 悟空 ========
type WuKong struct {
}

func (d *WuKong) Name() string {
	return "WuKong"
}

func (d *WuKong) Live() {
	d.Eat()
	d.Drink()
	d.Shit()
	d.Sleep()
}

func (d *WuKong) Eat() {
	log.Println(d.Name(), "eat")
}

func (d *WuKong) Drink() {
	log.Println(d.Name(), "drink")
}

func (d *WuKong) Shit() {
	log.Println(d.Name(), "shit")
	// 拉出很多废物
	_ = make([]byte, 16*Mi)
}

func (d *WuKong) Sleep() {
	log.Println(d.Name(), "sleep")
}

// ========= 唐僧 ========
type TangSeng struct {
}

func (d *TangSeng) Name() string {
	return "TangSeng"
}

func (d *TangSeng) Live() {
	d.Eat()
	d.Drink()
	d.Shit()
	d.Sleep()
}

func (d *TangSeng) Eat() {
	log.Println(d.Name(), "eat")
}

func (d *TangSeng) Drink() {
	log.Println(d.Name(), "drink")
}

func (d *TangSeng) Shit() {
	log.Println(d.Name(), "shit")
	// 锁问题
	m := &sync.Mutex{}
	m.Lock()
	go func() {
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
}

func (d *TangSeng) Sleep() {
	log.Println(d.Name(), "sleep")
	// 假睡觉
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(30 * time.Second)
		}()
	}
}

// ========= 猪八戒 ========
type ZhuBaJie struct {
}

func (d *ZhuBaJie) Name() string {
	return "ZhuBaJie"
}

func (d *ZhuBaJie) Live() {
	d.Eat()
	d.Drink()
	d.Shit()
	d.Sleep()
}

func (d *ZhuBaJie) Eat() {
	log.Println(d.Name(), "eat")
}

func (d *ZhuBaJie) Drink() {
	log.Println(d.Name(), "drink")
	// 空喝
	loop := 10000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}

func (d *ZhuBaJie) Shit() {
	log.Println(d.Name(), "shit")
}

func (d *ZhuBaJie) Sleep() {
	log.Println(d.Name(), "sleep")
	// 多睡一点
	<-time.After(time.Second)
}

// ========= 沙和尚 ========
type ShaHeShang struct {
	buffer [][Mi]byte
}

func (d *ShaHeShang) Name() string {
	return "ShaHeShang"
}

func (d *ShaHeShang) Live() {
	d.Eat()
	d.Drink()
	d.Shit()
	d.Sleep()
}

func (d *ShaHeShang) Eat() {
	log.Println(d.Name(), "eat")
	max := Gi
	for len(d.buffer)*Mi < max {
		d.buffer = append(d.buffer, [Mi]byte{})
	}
}

func (d *ShaHeShang) Drink() {
	log.Println(d.Name(), "drink")
}

func (d *ShaHeShang) Shit() {
	log.Println(d.Name(), "shit")
}

func (d *ShaHeShang) Sleep() {
	log.Println(d.Name(), "sleep")
}

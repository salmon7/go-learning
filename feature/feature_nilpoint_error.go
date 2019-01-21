package main

import "fmt"

type TA struct {
}

type IA interface {
	MethodA()
}

func (t TA) MethodA() {
	fmt.Printf("A MethodA\n")
}

type TB struct {
}

type IB interface {
	MethodB()
}

func (t TB) MethodB() {
	fmt.Printf("B MethodB\n")
}

type IC interface {
	MethodA()
	MethodB()
}

type TC struct {
	IA
	IB
}

func (t TC) MethodA() {
	fmt.Printf("C MethodA\n")
}

func (t TC) MethodB() {
	fmt.Printf("C MethodB\n")
}

func main() {
	ta := &TA{}
	ta.MethodA()

	tb := &TB{}
	tb.MethodB()

	var ib IB = tb
	if ib, ok := ib.(IB); ok {
		fmt.Printf("convert from ib to TB ok, %v\n", ib)
		ib.MethodB()
	} else {
		fmt.Printf("convert from ib to TB err, %v\n", ib)
	}

	tc := &TC{}
	var icb IB = tc
	if icb, ok := icb.(IB); ok {
		fmt.Printf("convert from ib to TC ok, %v\n", icb)
		icb.MethodB()
	} else {
		fmt.Printf("convert from ib to TC err, %v\n", icb)
	}

	method2()
	method3()
}

func method2() {
	fmt.Println("================")
	ta := &TA{}
	ta.MethodA()

	tb := &TB{}
	tb.MethodB()

	tc := &TC{}

	var ibc IC = tc
	if icb, ok := ibc.(IC); ok {
		fmt.Printf("convert from ib to TC ok, %v\n", icb)
		icb.MethodB()
	} else {
		fmt.Printf("convert from ib to TC err, %v\n", icb)
	}
}

func method3() {
	fmt.Println("================")

	rti := method4()
	if rti == nil {
		fmt.Printf("method3 rti ==nil \n")
	} else {
		fmt.Printf("method3 rti !=nil \n")
	}

	if temp, ok := rti.(lowLevelError); ok {
		fmt.Printf("method3 convert from rti to lowLevelError ok, %v\n", temp)
		if temp == nil {
			fmt.Printf("method3 convert from rti to lowLevelError ok, temp ==nil \n")
		} else {
			fmt.Printf("method3 convert from rti to lowLevelError ok, temp !=nil \n")
			temp.LowLevelError()
		}
	} else {
		fmt.Printf("method3 convert from rti to lowLevelError err, %v\n", temp)
	}
}

func method4() RTI{
	var rti RTI
	rti = returnNilRTI()
	if rti !=nil{
		fmt.Println("method4 returnNilRTI error")
		return rti
	}else{
		fmt.Println("method4 returnNilRTI ok")
	}

	rti = returnNilACE()
	if rti !=nil{
		fmt.Println("method4 returnNilACE error")
		return rti
	}else {
		fmt.Println("method4 returnNilACE ok")
	}
	return returnRTI()

}

func returnRTI() RTI {
	rti := &ACE{}
	return rti
}

func returnNilRTI() RTI {
	return nil
}

func returnNilACE() *ACE {
	/*各种业务逻辑，有错则return error，没错则return nil*/
	return nil
}

type ACE struct {
	RT
}

type RTI interface {
	Method()
}
type RT struct {
}

func (RT) Method() {
	fmt.Println("RT MethodRT")
}

func (a ACE) LowLevelError() {
	fmt.Println("ACE MethodLowLevelError")
}

type lowLevelError interface {
	LowLevelError()
}

/*
重点看method3()和method4()的代码
假设method4()是业务逻辑的代码，在没有错误时返回nil，有错误时则返回具体的RTI
最近写代码遇到一个奇怪的问题，在子方法中返回一个nil的struct类型【指针】A给父方法中的一个interface变量B后，父方法无法判断B是否为空
导致使用类型断言后，出现了空指针引用。

看输出可以看到即使returnNilACE()返回的是nil（意味着没有错误），程序依然输出 "method4 returnNilACE error"，最后导致了method4()的panic
原因在于 rti = returnNilACE()，这行代码导致rit的value值为nil，但是它的type值依然为RTI，rti为接口，只有当它的value和type值都为nil时，接口才为空

避免方法：
1.rti = returnNilACE() 改为 rtix := returnNilACE()
2.修改returnNilRTI返回为*ACE

 */
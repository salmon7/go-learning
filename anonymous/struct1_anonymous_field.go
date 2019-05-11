package main

import "fmt"

type Parent1 struct {
	Child1
}
type Child1 struct {
	Name string
}

type Parent2 struct {
	*Child2
}
type Child2 struct {
	Name string
}

type Parent3 struct {
	C Child3
}
type Child3 struct {
	Name string
}

type Parent4 struct {
	C *Child4
}
type Child4 struct {
	Name string
}


func main(){
	parent1 := Parent1{Child1{}}
	parent2 := Parent2{&Child2{}}
	parent3 := Parent3{Child3{}}
	parent4 := Parent4{&Child4{}}

	cp_parent1 := parent1
	cp_parent2 := parent2
	cp_parent3 := parent3
	cp_parent4 := parent4

	fmt.Printf("&parent1=%p, &cp_parent1=%p, &parent1.Child1=%p, &cp_parent1.Child1=%p\n", &parent1, &cp_parent1, &parent1.Child1, &cp_parent1.Child1)

	fmt.Printf("&parent2=%p, &cp_parent2=%p, parent2.Child2=%p, cp_parent2.Child2=%p\n", &parent2, &cp_parent2, parent2.Child2, cp_parent2.Child2)

	fmt.Printf("&parent3=%p, &cp_parent3=%p, &parent3.C=%p, &cp_parent3.C=%p\n", &parent3, &cp_parent3, &parent3.C, &cp_parent3.C)

	fmt.Printf("&parent4=%p, &cp_parent4=%p, parent4.C=%p, cp_parent4.C=%p\n", &parent4, &cp_parent4, parent4.C, cp_parent4.C)

}

/*
output:
&parent1=0xc42007c1c0, &cp_parent1=0xc42007c200, &parent1.Child1=0xc42007c1c0, &cp_parent1.Child1=0xc42007c200
&parent2=0xc42008a018, &cp_parent2=0xc42008a028, parent2.Child2=0xc42007c1d0, cp_parent2.Child2=0xc42007c1d0
&parent3=0xc42007c1e0, &cp_parent3=0xc42007c210, &parent3.C=0xc42007c1e0, &cp_parent3.C=0xc42007c210
&parent4=0xc42008a020, &cp_parent4=0xc42008a030, parent4.C=0xc42007c1f0, cp_parent4.C=0xc42007c1f0

可以看到parent2和parent4的子类地址在复制品中的是相同的

凡是使用地址类型，无论在结构体中直接嵌入另一个结构体，还是在结构体中定义一个变量，如果这个被嵌入的结构体/变量是个指针类型，那么赋值时，只会赋值地址，不会分配新的空间。（这句话把地址类型换成 map，slice，chan同样适用）

ps：
1.有一个诡异现象，如果把Child中的Name去掉，会发现以上的规则会被破坏
2.最近开发时有个一个新需求，在一个实例中，需要实现一个数据库连接要被多个数据库使用，即一个数据库变量要实现操作多个数据库。由于本身进行了部分封装，需要对现有的代码添加dbname变量，刚好就遇到了本文所描述的问题。实现了一个数据连接变量，对应多个dbname

参考：
Go语言传值和深浅复制问题
http://kchu.me/2016/03/27/Go%E8%AF%AD%E8%A8%80%E4%BC%A0%E5%80%BC%E5%92%8C%E6%B7%B1%E6%B5%85%E5%A4%8D%E5%88%B6%E9%97%AE%E9%A2%98/
https://blog.csdn.net/lolimostlovely/article/details/80717701
 */
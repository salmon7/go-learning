package main

import (
	"log"
)

type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	admin := &Admin{
		User: User{
			Name:  "AriesDevil",
			Email: "ariesdevil@xxoo.com",
		},
		Level: "super",
	}

	SendNotification(admin)
}


/*
考虑一个问题，加一个interface定义了多个方法，一个struct实现了这些方法，在赋值给一个非空的interface时，什么时候要传这个struct的值，什么时候远传这个struct的地址

讨论这个问题，我们需要了解一个struct/*struct所包含的方法集:
1.类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
2.类型 T 的可调用方法集包含接受者为 T 的所有方法集
3.类型 T 的可调用方法集不包含接受者为 *T 的方法

即
*T -> T, *T
T -> T

如何是子类实现了接口的方法呢？子类父类中为一个值和指针时有什么不同？
1.如果 S 包含一个匿名字段 T，S 和 *S 的方法集都包含接受者为 T 的方法提升。
2.对于 *S 类型的方法集包含接受者为 *T 的方法提升
3.如果 S 包含一个匿名字段 *T，S 和 *S 的方法集都包含接受者为 T 或者 *T 的方法提升

即
S(T) -> T
*S(T) -> T, *T
S(*T) -> T, *T
*S(*S) -> T, *T
可以总结为只要S和T其中一个为指针，那么S和*S就会包含T和*T的方法集

参考：https://www.jianshu.com/p/d87c69ac6ce7
 */

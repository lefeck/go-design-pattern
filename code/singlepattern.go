package main

import "fmt"

type IdentityCardNo struct {
	instance *IdentityCardNo
	no       string
}

func (i *IdentityCardNo) getinstance() *IdentityCardNo {
	no, ok := i.getIdentityCardNo()

	if ok {
		fmt.Printf("身份证ID: %s 已经存在\n", no)
		fmt.Println("重复办理身份证，获取旧号码！")
	}
	if i.instance == nil {
		fmt.Println("第一次办理身份证，分配新号码！")
		i.setIdentityCardNo("No1234135324253")
	}
	return i.instance
}

func (i *IdentityCardNo) setIdentityCardNo(no string) {
	fmt.Printf("身份证ID: %s \n", no)
	i.no = no
}

func (i *IdentityCardNo) getIdentityCardNo() (string, bool) {
	if i.no != "" {
		return i.no, true
	}
	return i.no, false
}

func main() {
	for i := 0; i < 1; i++ {
		c := &IdentityCardNo{no: "No123413533425"}
		c.getinstance()
	}

}

package memory

import (
	"fmt"
	"unsafe"
)

//Memory allocated to variable of type struct will depend on the way
//struct fields are arrange. That is happening because in 64 bit machine for a one CPU cycle
//CPU can only access 64bit size of data only
// Since software architecture in 64, virtual memory addresses can hold only 64 bits,
//In one computer cycle, CPU will access only one memory access
func Memory_struct_demo(){
	type demo_high_memory struct {
		age int16
		name string
		checked bool
	}
	type demo_low_memory struct {
		name string
		age int16
		checked bool
	}
	var a = demo_high_memory{
		10,
		"ss",
		true,
	}
	var b = demo_low_memory{
		"ss",
		10,
		true,
	}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}

# go-memorydll
a go wrapper for memory dll 

it's the same as system's loadlibrary but it load dll content from memory
usage :
dll,err:=NewDLL(testdll,"example.dll");
proc,err:=dll.FindProc("gcd")
result,_,_:=proc.Call(uintptr(4),uintptr(8))

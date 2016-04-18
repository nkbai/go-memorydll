package memorydll

//#include<stdlib.h>
// extern void * MemoryLoadLibrary(const void *);
// extern void * MemoryGetProcAddress(void *, char *);
// extern void MemoryFreeLibrary(void *);
import "C"
import (
	"unsafe"
	"errors"
	"syscall"
	"fmt"
)

type Handle uintptr
// A DLL implements access to a single DLL.
type DLL struct {
	Name   string
	Handle Handle
}

// DLLError describes reasons for DLL load failures.
type DLLError struct {
	Err     error
	ObjName string
	Msg     string
}

func (e *DLLError) Error() string { return e.Msg }
// FindProc searches DLL d for procedure named name and returns *Proc
// if found. It returns an error if search fails.
func (d *DLL) FindProc(name string) (proc *Proc, err error) {
	return  memoryGetProcAddress(d, name)
}

// MustFindProc is like FindProc but panics if search fails.
func (d *DLL) MustFindProc(name string) *Proc {
	p, e := d.FindProc(name)
	if e != nil {
		panic(e)
	}
	return p
}

// Release unloads DLL d from memory.
func (d *DLL) Release() {
	memoryFreeLibrary(d)
}

// A Proc implements access to a procedure inside a DLL.
type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
}

// Addr returns the address of the procedure represented by p.
// The return value can be passed to Syscall to run the procedure.
func (p *Proc) Addr() uintptr {
	return p.addr
}

// Call executes procedure p with arguments a. It will panic, if more then 15 arguments
// are supplied.
//
// The returned error is always non-nil, constructed from the result of GetLastError.
// Callers must inspect the primary return value to decide whether an error occurred
// (according to the semantics of the specific function being called) before consulting
// the error. The error will be guaranteed to contain syscall.Errno.
func (p *Proc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	switch len(a) {
	case 0:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), 0, 0, 0)
	case 1:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), a[0], 0, 0)
	case 2:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), a[0], a[1], 0)
	case 3:
		return syscall.Syscall(p.Addr(), uintptr(len(a)), a[0], a[1], a[2])
	case 4:
		return syscall.Syscall6(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], 0, 0)
	case 5:
		return syscall.Syscall6(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], 0)
	case 6:
		return syscall.Syscall6(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5])
	case 7:
		return syscall.Syscall9(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0)
	case 8:
		return syscall.Syscall9(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0)
	case 9:
		return syscall.Syscall9(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
	case 10:
		return syscall.Syscall12(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0)
	case 11:
		return syscall.Syscall12(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0)
	case 12:
		return syscall.Syscall12(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
	case 13:
		return syscall.Syscall15(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], 0, 0)
	case 14:
		return syscall.Syscall15(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], 0)
	case 15:
		return syscall.Syscall15(p.Addr(), uintptr(len(a)), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14])
	default:
		panic("Call " + p.Name + " with too many arguments " + fmt.Sprintf("%d",len(a)) + ".")
	}
	return
}

//remembe to release when this dll is useless
func NewDLL(dlldata []byte,dllname string) (*DLL,error){
	dlldataPointer:=unsafe.Pointer(&dlldata[0])
	handle:=C.MemoryLoadLibrary(dlldataPointer);
	if handle!=nil{
		return &DLL{
			Name:dllname,
			Handle:Handle(handle),
		},nil
	}else{
			e:=errors.New("dll data error");
			return nil, &DLLError{
			Err:     e,
			ObjName: dllname,
			Msg:     "Failed to load " + dllname + ": " + e.Error(),
		}
	}

}

func memoryGetProcAddress(dll *DLL, procname string ) (proc *Proc, err error) {
	cname:=C.CString(procname)
	defer C.free(unsafe.Pointer(cname))
	addr:=C.MemoryGetProcAddress(unsafe.Pointer(dll.Handle),cname);
	if addr!=nil{
		return  &Proc{
			Dll:dll,
			Name:procname,
			addr:uintptr(addr),
		},nil
	}
	e:=errors.New("no such function");
	return nil, &DLLError{
			Err:     e,
			ObjName: procname,
			Msg:     "Failed to find " + procname + " procedure in " + dll.Name + ": " + e.Error(),
		}
}
//remember free!
func memoryFreeLibrary(dll * DLL){
	C.MemoryFreeLibrary(unsafe.Pointer(dll.Handle));
}
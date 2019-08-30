// Code generated by 'go generate'; DO NOT EDIT.

package namespaceapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procCreateBoundaryDescriptorW  = modkernel32.NewProc("CreateBoundaryDescriptorW")
	procDeleteBoundaryDescriptor   = modkernel32.NewProc("DeleteBoundaryDescriptor")
	procAddSIDToBoundaryDescriptor = modkernel32.NewProc("AddSIDToBoundaryDescriptor")
	procCreatePrivateNamespaceW    = modkernel32.NewProc("CreatePrivateNamespaceW")
	procOpenPrivateNamespaceW      = modkernel32.NewProc("OpenPrivateNamespaceW")
	procClosePrivateNamespace      = modkernel32.NewProc("ClosePrivateNamespace")
)

func createBoundaryDescriptor(name *uint16, flags uint32) (handle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateBoundaryDescriptorW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(flags), 0)
	handle = windows.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func deleteBoundaryDescriptor(boundaryDescriptor windows.Handle) {
	syscall.Syscall(procDeleteBoundaryDescriptor.Addr(), 1, uintptr(boundaryDescriptor), 0, 0)
	return
}

func addSIDToBoundaryDescriptor(boundaryDescriptor windows.Handle, requiredSid *windows.SID) (err error) {
	r1, _, e1 := syscall.Syscall(procAddSIDToBoundaryDescriptor.Addr(), 2, uintptr(boundaryDescriptor), uintptr(unsafe.Pointer(requiredSid)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func createPrivateNamespace(privateNamespaceAttributes *windows.SecurityAttributes, boundaryDescriptor windows.Handle, aliasPrefix *uint16) (handle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreatePrivateNamespaceW.Addr(), 3, uintptr(unsafe.Pointer(privateNamespaceAttributes)), uintptr(boundaryDescriptor), uintptr(unsafe.Pointer(aliasPrefix)))
	handle = windows.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func openPrivateNamespace(boundaryDescriptor windows.Handle, aliasPrefix *uint16) (handle windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procOpenPrivateNamespaceW.Addr(), 2, uintptr(boundaryDescriptor), uintptr(unsafe.Pointer(aliasPrefix)), 0)
	handle = windows.Handle(r0)
	if handle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func closePrivateNamespace(handle windows.Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procClosePrivateNamespace.Addr(), 2, uintptr(handle), uintptr(flags), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
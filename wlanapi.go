package wlanapi

import (
    "fmt"
    "golang.org/x/sys/windows"
    "syscall"
    "unsafe"
)

func retValToError(errNo uintptr) error {

    return fmt.Errorf("return code: %x", errNo)
}


func OpenHandle() (handle *windows.Handle, err error) {

    clientVer := uintptr(2)
    var h windows.Handle

    r1, _, _ := fWlanOpenHandle.Call(
        clientVer,
        uintptr(0),
        uintptr(unsafe.Pointer(&clientVer)),
        uintptr(unsafe.Pointer(&h)),
        )

    if r1 != S_OK {
        return nil, retValToError(r1)
    }

    return &h, nil
}

func CloseHandle(handle *windows.Handle) (err error) {

    r1, _, _ := fWlanCloseHandle.Call(
        uintptr(unsafe.Pointer(handle)),
        uintptr(0),
        )

    if r1 != S_OK {
        return retValToError(r1)
    }

    return nil
}

func EnumInterfaces(handle *windows.Handle) (interfaceInfoList *WLAN_INTERFACE_INFO_LIST, err error) {

    var iil *WLAN_INTERFACE_INFO_LIST

    r1, _, _ := fWlanEnumInterfaces.Call(
        uintptr(unsafe.Pointer(handle)),
        uintptr(0),
        uintptr(unsafe.Pointer(&iil)),
        )

    if r1 != S_OK {
        return nil, retValToError(r1)
    }

    return iil, nil

    //TODO: Need to rework to cleanup unmanaged memory WlanFreeMemory(uintptr(unsafe.Pointer(iil)))
}

func defaultInterface(handle *windows.Handle) (guid *syscall.GUID, description string, err error) {

    iil, err := EnumInterfaces(handle)
    if err != nil {
        return guid, "", err
    }

    wli := iil.InterfaceInfo[iil.dwIndex]

    return &wli.InterfaceGuid, windows.UTF16ToString(wli.strInterfaceDescription[:]), nil
}
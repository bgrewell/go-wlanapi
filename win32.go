package wlanapi

import "syscall"

var (
    wlanapi = syscall.NewLazyDLL("wlanapi.dll")
    fWlanOpenHandle = wlanapi.NewProc("WlanOpenHandle")
    fWlanCloseHandle = wlanapi.NewProc("WlanCloseHandle")
    fWlanEnumInterfaces = wlanapi.NewProc("WlanEnumInterfaces")
)

package wlanapi

import (
    "golang.org/x/sys/windows"
    "syscall"
)

type Manager struct {
    handle *windows.Handle
    interfaceGuid *syscall.GUID
    InterfaceName string
}

func (m *Manager) Initialize() (err error) {

    // Get a handle to the wlan api
    m.handle, err = OpenHandle()
    if err != nil {
        return err
    }

    // Find the first interface
    // TODO: should allow the user to select
    m.interfaceGuid, m.InterfaceName, err = defaultInterface(m.handle)
    if err != nil {
        return err
    }

    return nil
}

package wlanapi

import "syscall"

const (
    MAX_INDEX = 1000
    S_OK = 0
)

type DOT11_SSID struct {
    uSSIDLength uint32
    ucSSID      [32]byte
}

type WLAN_AVAILABLE_NETWORK_LIST struct {
    dwNumberOfItems uint32
    dwIndex         uint32
    Network         [MAX_INDEX + 1]WLAN_AVAILABLE_NETWORK
}

type WLAN_AVAILABLE_NETWORK struct {
    strProfileName              [256]uint16
    dot11Ssid                   DOT11_SSID
    dot11BssType                uint32
    uNumberOfBssids             uint32
    bNetworkConnectable         int32
    wlanNotConnectableReason    uint32
    uNumberOfPhyTypes           uint32
    dot11PhyTypes               [8]uint32
    bMorePhyTypes               int32
    wlanSignalQuality           uint32
    bSecurityEnabled            int32
    dot11DefaultAuthAlgorithm   uint32
    dot11DefaultCipherAlgorithm uint32
    dwFlags                     uint32
    dwReserved                  uint32
}

type WLAN_BSS_LIST struct {
    dwTotalSize     uint32
    dwNumberOfItems uint32
    wlanBssEntries  [MAX_INDEX + 1]WLAN_BSS_ENTRY
}

type WLAN_BSS_ENTRY struct {
    dot11Ssid               DOT11_SSID
    uPhyId                  uint32
    dot11Bssid              [6]byte
    dot11BssType            uint32
    dot11BssPhyType         uint32
    lRssi                   int32
    uLinkQuality            uint32
    bInRegDomain            int32
    usBeaconPeriod          uint16
    ullTimestamp            uint64
    ullHostTimestamp        uint64
    usCapabilityInformation uint16
    ulChCenterFrequency     uint32
    wlanRateSet             WLAN_RATE_SET
    ulIeOffset              uint32
    ulIeSize                uint32
}

type WLAN_RATE_SET struct {
    uRateSetLength uint32
    usRateSet      [126]uint16
}

type WLAN_INTERFACE_INFO struct {
    InterfaceGuid           syscall.GUID
    strInterfaceDescription [256]uint16
    isState                 uint32
}

type WLAN_INTERFACE_INFO_LIST struct {
    dwNumberOfItems uint32
    dwIndex         uint32
    InterfaceInfo   [MAX_INDEX + 1]WLAN_INTERFACE_INFO
}

type WLAN_PROFILE_INFO struct{
    ProfileName uint8
    Flags uint32
}

type WLAN_PROFILE_INFO_LIST struct{
    NumberOfItems uint32
    Index uint32
    ProfileInfo WLAN_PROFILE_INFO
}
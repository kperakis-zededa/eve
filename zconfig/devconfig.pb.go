// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devconfig.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SWAdapterType int32

const (
	SWAdapterType_IGNORE SWAdapterType = 0
	SWAdapterType_VLAN   SWAdapterType = 1
	SWAdapterType_BOND   SWAdapterType = 2
)

var SWAdapterType_name = map[int32]string{
	0: "IGNORE",
	1: "VLAN",
	2: "BOND",
}
var SWAdapterType_value = map[string]int32{
	"IGNORE": 0,
	"VLAN":   1,
	"BOND":   2,
}

func (x SWAdapterType) String() string {
	return proto.EnumName(SWAdapterType_name, int32(x))
}
func (SWAdapterType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type MapServer struct {
	NameOrIp   string `protobuf:"bytes,1,opt,name=NameOrIp" json:"NameOrIp,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=Credential" json:"Credential,omitempty"`
}

func (m *MapServer) Reset()                    { *m = MapServer{} }
func (m *MapServer) String() string            { return proto.CompactTextString(m) }
func (*MapServer) ProtoMessage()               {}
func (*MapServer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *MapServer) GetNameOrIp() string {
	if m != nil {
		return m.NameOrIp
	}
	return ""
}

func (m *MapServer) GetCredential() string {
	if m != nil {
		return m.Credential
	}
	return ""
}

type ZedServer struct {
	HostName string   `protobuf:"bytes,1,opt,name=HostName" json:"HostName,omitempty"`
	EID      []string `protobuf:"bytes,2,rep,name=EID" json:"EID,omitempty"`
}

func (m *ZedServer) Reset()                    { *m = ZedServer{} }
func (m *ZedServer) String() string            { return proto.CompactTextString(m) }
func (*ZedServer) ProtoMessage()               {}
func (*ZedServer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ZedServer) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ZedServer) GetEID() []string {
	if m != nil {
		return m.EID
	}
	return nil
}

type DeviceLispDetails struct {
	LispMapServers         []*MapServer `protobuf:"bytes,1,rep,name=LispMapServers" json:"LispMapServers,omitempty"`
	LispInstance           uint32       `protobuf:"varint,2,opt,name=LispInstance" json:"LispInstance,omitempty"`
	EID                    string       `protobuf:"bytes,4,opt,name=EID" json:"EID,omitempty"`
	EIDHashLen             uint32       `protobuf:"varint,5,opt,name=EIDHashLen" json:"EIDHashLen,omitempty"`
	ZedServers             []*ZedServer `protobuf:"bytes,6,rep,name=ZedServers" json:"ZedServers,omitempty"`
	EidAllocationPrefix    []byte       `protobuf:"bytes,8,opt,name=EidAllocationPrefix,proto3" json:"EidAllocationPrefix,omitempty"`
	EidAllocationPrefixLen uint32       `protobuf:"varint,9,opt,name=EidAllocationPrefixLen" json:"EidAllocationPrefixLen,omitempty"`
	ClientAddr             string       `protobuf:"bytes,10,opt,name=ClientAddr" json:"ClientAddr,omitempty"`
	Experimental           bool         `protobuf:"varint,20,opt,name=Experimental" json:"Experimental,omitempty"`
}

func (m *DeviceLispDetails) Reset()                    { *m = DeviceLispDetails{} }
func (m *DeviceLispDetails) String() string            { return proto.CompactTextString(m) }
func (*DeviceLispDetails) ProtoMessage()               {}
func (*DeviceLispDetails) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DeviceLispDetails) GetLispMapServers() []*MapServer {
	if m != nil {
		return m.LispMapServers
	}
	return nil
}

func (m *DeviceLispDetails) GetLispInstance() uint32 {
	if m != nil {
		return m.LispInstance
	}
	return 0
}

func (m *DeviceLispDetails) GetEID() string {
	if m != nil {
		return m.EID
	}
	return ""
}

func (m *DeviceLispDetails) GetEIDHashLen() uint32 {
	if m != nil {
		return m.EIDHashLen
	}
	return 0
}

func (m *DeviceLispDetails) GetZedServers() []*ZedServer {
	if m != nil {
		return m.ZedServers
	}
	return nil
}

func (m *DeviceLispDetails) GetEidAllocationPrefix() []byte {
	if m != nil {
		return m.EidAllocationPrefix
	}
	return nil
}

func (m *DeviceLispDetails) GetEidAllocationPrefixLen() uint32 {
	if m != nil {
		return m.EidAllocationPrefixLen
	}
	return 0
}

func (m *DeviceLispDetails) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

func (m *DeviceLispDetails) GetExperimental() bool {
	if m != nil {
		return m.Experimental
	}
	return false
}

// Device Operational Commands Semantic
// For rebooting device,  command=Reset, counter = counter+delta, desiredState = on
// For poweroff device,  command=Reset, counter = counter+delta, desiredState = off
// For backup at midnight, command=Backup, counter = counter+delta, desiredState=n/a, opsTime = mm/dd/yy:hh:ss
// Current implementation does support only single command outstanding for each type
// In future can be extended to have more scheduled events
//
type DeviceOpsCmd struct {
	Counter      uint32 `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
	DesiredState bool   `protobuf:"varint,3,opt,name=desiredState" json:"desiredState,omitempty"`
	// FIXME: change to timestamp, once we move to gogo proto
	OpsTime string `protobuf:"bytes,4,opt,name=opsTime" json:"opsTime,omitempty"`
}

func (m *DeviceOpsCmd) Reset()                    { *m = DeviceOpsCmd{} }
func (m *DeviceOpsCmd) String() string            { return proto.CompactTextString(m) }
func (*DeviceOpsCmd) ProtoMessage()               {}
func (*DeviceOpsCmd) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DeviceOpsCmd) GetCounter() uint32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *DeviceOpsCmd) GetDesiredState() bool {
	if m != nil {
		return m.DesiredState
	}
	return false
}

func (m *DeviceOpsCmd) GetOpsTime() string {
	if m != nil {
		return m.OpsTime
	}
	return ""
}

type SWAdapterParams struct {
	AType SWAdapterType `protobuf:"varint,1,opt,name=aType,enum=SWAdapterType" json:"aType,omitempty"`
	// vlan
	UnderlayInterface string `protobuf:"bytes,8,opt,name=underlayInterface" json:"underlayInterface,omitempty"`
	VlanId            uint32 `protobuf:"varint,9,opt,name=vlanId" json:"vlanId,omitempty"`
	// OR : repeated physical interfaces for bond0
	Bondgroup []string `protobuf:"bytes,10,rep,name=bondgroup" json:"bondgroup,omitempty"`
}

func (m *SWAdapterParams) Reset()                    { *m = SWAdapterParams{} }
func (m *SWAdapterParams) String() string            { return proto.CompactTextString(m) }
func (*SWAdapterParams) ProtoMessage()               {}
func (*SWAdapterParams) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *SWAdapterParams) GetAType() SWAdapterType {
	if m != nil {
		return m.AType
	}
	return SWAdapterType_IGNORE
}

func (m *SWAdapterParams) GetUnderlayInterface() string {
	if m != nil {
		return m.UnderlayInterface
	}
	return ""
}

func (m *SWAdapterParams) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *SWAdapterParams) GetBondgroup() []string {
	if m != nil {
		return m.Bondgroup
	}
	return nil
}

type SystemAdapter struct {
	// name of the adapter; hardware-specific e.g., eth0
	Name         string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	AllocDetails *SWAdapterParams `protobuf:"bytes,20,opt,name=allocDetails" json:"allocDetails,omitempty"`
	// this is part of the freelink group
	FreeUplink bool `protobuf:"varint,2,opt,name=freeUplink" json:"freeUplink,omitempty"`
	// this is part of the uplink group
	Uplink bool `protobuf:"varint,3,opt,name=uplink" json:"uplink,omitempty"`
	// attach this network config for this adapter
	NetworkUUID string `protobuf:"bytes,4,opt,name=networkUUID" json:"networkUUID,omitempty"`
	// if its static network we need ip address
	Addr string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty"`
	// alias/logical name which will be reported to zedcloud
	// and used for app instances
	LogicalName string `protobuf:"bytes,6,opt,name=logicalName" json:"logicalName,omitempty"`
}

func (m *SystemAdapter) Reset()                    { *m = SystemAdapter{} }
func (m *SystemAdapter) String() string            { return proto.CompactTextString(m) }
func (*SystemAdapter) ProtoMessage()               {}
func (*SystemAdapter) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *SystemAdapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SystemAdapter) GetAllocDetails() *SWAdapterParams {
	if m != nil {
		return m.AllocDetails
	}
	return nil
}

func (m *SystemAdapter) GetFreeUplink() bool {
	if m != nil {
		return m.FreeUplink
	}
	return false
}

func (m *SystemAdapter) GetUplink() bool {
	if m != nil {
		return m.Uplink
	}
	return false
}

func (m *SystemAdapter) GetNetworkUUID() string {
	if m != nil {
		return m.NetworkUUID
	}
	return ""
}

func (m *SystemAdapter) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SystemAdapter) GetLogicalName() string {
	if m != nil {
		return m.LogicalName
	}
	return ""
}

type EdgeDevConfig struct {
	Id                 *UUIDandVersion          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DevConfigSha256    []byte                   `protobuf:"bytes,2,opt,name=devConfigSha256,proto3" json:"devConfigSha256,omitempty"`
	DevConfigSignature []byte                   `protobuf:"bytes,3,opt,name=devConfigSignature,proto3" json:"devConfigSignature,omitempty"`
	Apps               []*AppInstanceConfig     `protobuf:"bytes,4,rep,name=apps" json:"apps,omitempty"`
	Networks           []*NetworkConfig         `protobuf:"bytes,5,rep,name=networks" json:"networks,omitempty"`
	Datastores         []*DatastoreConfig       `protobuf:"bytes,6,rep,name=datastores" json:"datastores,omitempty"`
	LispInfo           *DeviceLispDetails       `protobuf:"bytes,7,opt,name=lispInfo" json:"lispInfo,omitempty"`
	Base               []*BaseOSConfig          `protobuf:"bytes,8,rep,name=base" json:"base,omitempty"`
	Reboot             *DeviceOpsCmd            `protobuf:"bytes,9,opt,name=reboot" json:"reboot,omitempty"`
	Backup             *DeviceOpsCmd            `protobuf:"bytes,10,opt,name=backup" json:"backup,omitempty"`
	ConfigItems        []*ConfigItem            `protobuf:"bytes,11,rep,name=configItems" json:"configItems,omitempty"`
	SystemAdapterList  []*SystemAdapter         `protobuf:"bytes,12,rep,name=systemAdapterList" json:"systemAdapterList,omitempty"`
	Services           []*ServiceInstanceConfig `protobuf:"bytes,13,rep,name=services" json:"services,omitempty"`
	// Override dmidecode info if set
	Manufacturer     string                   `protobuf:"bytes,14,opt,name=manufacturer" json:"manufacturer,omitempty"`
	ProductName      string                   `protobuf:"bytes,15,opt,name=productName" json:"productName,omitempty"`
	NetworkInstances []*NetworkInstanceConfig `protobuf:"bytes,16,rep,name=networkInstances" json:"networkInstances,omitempty"`
}

func (m *EdgeDevConfig) Reset()                    { *m = EdgeDevConfig{} }
func (m *EdgeDevConfig) String() string            { return proto.CompactTextString(m) }
func (*EdgeDevConfig) ProtoMessage()               {}
func (*EdgeDevConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *EdgeDevConfig) GetId() *UUIDandVersion {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EdgeDevConfig) GetDevConfigSha256() []byte {
	if m != nil {
		return m.DevConfigSha256
	}
	return nil
}

func (m *EdgeDevConfig) GetDevConfigSignature() []byte {
	if m != nil {
		return m.DevConfigSignature
	}
	return nil
}

func (m *EdgeDevConfig) GetApps() []*AppInstanceConfig {
	if m != nil {
		return m.Apps
	}
	return nil
}

func (m *EdgeDevConfig) GetNetworks() []*NetworkConfig {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *EdgeDevConfig) GetDatastores() []*DatastoreConfig {
	if m != nil {
		return m.Datastores
	}
	return nil
}

func (m *EdgeDevConfig) GetLispInfo() *DeviceLispDetails {
	if m != nil {
		return m.LispInfo
	}
	return nil
}

func (m *EdgeDevConfig) GetBase() []*BaseOSConfig {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *EdgeDevConfig) GetReboot() *DeviceOpsCmd {
	if m != nil {
		return m.Reboot
	}
	return nil
}

func (m *EdgeDevConfig) GetBackup() *DeviceOpsCmd {
	if m != nil {
		return m.Backup
	}
	return nil
}

func (m *EdgeDevConfig) GetConfigItems() []*ConfigItem {
	if m != nil {
		return m.ConfigItems
	}
	return nil
}

func (m *EdgeDevConfig) GetSystemAdapterList() []*SystemAdapter {
	if m != nil {
		return m.SystemAdapterList
	}
	return nil
}

func (m *EdgeDevConfig) GetServices() []*ServiceInstanceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *EdgeDevConfig) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *EdgeDevConfig) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *EdgeDevConfig) GetNetworkInstances() []*NetworkInstanceConfig {
	if m != nil {
		return m.NetworkInstances
	}
	return nil
}

// Timers and other per-device policy which relates to the interaction
// with zedcloud. Note that the timers are randomized on the device
// to avoid synchronization with other devices. Random range is between
// between .5 and 1.5 of these nominal values. If not set (i.e. zero),
// it means the default value of 60 seconds.
type ConfigItem struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	// Below section will be deprecated
	//
	// Types that are valid to be assigned to ConfigItemValue:
	//	*ConfigItem_BoolValue
	//	*ConfigItem_Uint32Value
	//	*ConfigItem_Uint64Value
	//	*ConfigItem_FloatValue
	//	*ConfigItem_StringValue
	ConfigItemValue isConfigItem_ConfigItemValue `protobuf_oneof:"configItemValue"`
}

func (m *ConfigItem) Reset()                    { *m = ConfigItem{} }
func (m *ConfigItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigItem) ProtoMessage()               {}
func (*ConfigItem) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

type isConfigItem_ConfigItemValue interface {
	isConfigItem_ConfigItemValue()
}

type ConfigItem_BoolValue struct {
	BoolValue bool `protobuf:"varint,3,opt,name=boolValue,oneof"`
}
type ConfigItem_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,4,opt,name=uint32Value,oneof"`
}
type ConfigItem_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,5,opt,name=uint64Value,oneof"`
}
type ConfigItem_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,6,opt,name=floatValue,oneof"`
}
type ConfigItem_StringValue struct {
	StringValue string `protobuf:"bytes,7,opt,name=stringValue,oneof"`
}

func (*ConfigItem_BoolValue) isConfigItem_ConfigItemValue()   {}
func (*ConfigItem_Uint32Value) isConfigItem_ConfigItemValue() {}
func (*ConfigItem_Uint64Value) isConfigItem_ConfigItemValue() {}
func (*ConfigItem_FloatValue) isConfigItem_ConfigItemValue()  {}
func (*ConfigItem_StringValue) isConfigItem_ConfigItemValue() {}

func (m *ConfigItem) GetConfigItemValue() isConfigItem_ConfigItemValue {
	if m != nil {
		return m.ConfigItemValue
	}
	return nil
}

func (m *ConfigItem) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ConfigItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ConfigItem) GetBoolValue() bool {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *ConfigItem) GetUint32Value() uint32 {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (m *ConfigItem) GetUint64Value() uint64 {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *ConfigItem) GetFloatValue() float32 {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *ConfigItem) GetStringValue() string {
	if x, ok := m.GetConfigItemValue().(*ConfigItem_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConfigItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConfigItem_OneofMarshaler, _ConfigItem_OneofUnmarshaler, _ConfigItem_OneofSizer, []interface{}{
		(*ConfigItem_BoolValue)(nil),
		(*ConfigItem_Uint32Value)(nil),
		(*ConfigItem_Uint64Value)(nil),
		(*ConfigItem_FloatValue)(nil),
		(*ConfigItem_StringValue)(nil),
	}
}

func _ConfigItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConfigItem)
	// configItemValue
	switch x := m.ConfigItemValue.(type) {
	case *ConfigItem_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ConfigItem_Uint32Value:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint32Value))
	case *ConfigItem_Uint64Value:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *ConfigItem_FloatValue:
		b.EncodeVarint(6<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case *ConfigItem_StringValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case nil:
	default:
		return fmt.Errorf("ConfigItem.ConfigItemValue has unexpected type %T", x)
	}
	return nil
}

func _ConfigItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConfigItem)
	switch tag {
	case 3: // configItemValue.boolValue
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConfigItemValue = &ConfigItem_BoolValue{x != 0}
		return true, err
	case 4: // configItemValue.uint32Value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConfigItemValue = &ConfigItem_Uint32Value{uint32(x)}
		return true, err
	case 5: // configItemValue.uint64Value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ConfigItemValue = &ConfigItem_Uint64Value{x}
		return true, err
	case 6: // configItemValue.floatValue
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.ConfigItemValue = &ConfigItem_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	case 7: // configItemValue.stringValue
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ConfigItemValue = &ConfigItem_StringValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _ConfigItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConfigItem)
	// configItemValue
	switch x := m.ConfigItemValue.(type) {
	case *ConfigItem_BoolValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += 1
	case *ConfigItem_Uint32Value:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint32Value))
	case *ConfigItem_Uint64Value:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *ConfigItem_FloatValue:
		n += proto.SizeVarint(6<<3 | proto.WireFixed32)
		n += 4
	case *ConfigItem_StringValue:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ConfigRequest struct {
	ConfigHash string `protobuf:"bytes,1,opt,name=configHash" json:"configHash,omitempty"`
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *ConfigRequest) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

type ConfigResponse struct {
	Config     *EdgeDevConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	ConfigHash string         `protobuf:"bytes,2,opt,name=configHash" json:"configHash,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *ConfigResponse) GetConfig() *EdgeDevConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigResponse) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

func init() {
	proto.RegisterType((*MapServer)(nil), "MapServer")
	proto.RegisterType((*ZedServer)(nil), "ZedServer")
	proto.RegisterType((*DeviceLispDetails)(nil), "DeviceLispDetails")
	proto.RegisterType((*DeviceOpsCmd)(nil), "DeviceOpsCmd")
	proto.RegisterType((*SWAdapterParams)(nil), "sWAdapterParams")
	proto.RegisterType((*SystemAdapter)(nil), "SystemAdapter")
	proto.RegisterType((*EdgeDevConfig)(nil), "EdgeDevConfig")
	proto.RegisterType((*ConfigItem)(nil), "ConfigItem")
	proto.RegisterType((*ConfigRequest)(nil), "ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "ConfigResponse")
	proto.RegisterEnum("SWAdapterType", SWAdapterType_name, SWAdapterType_value)
}

func init() { proto.RegisterFile("devconfig.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xed, 0x6e, 0xdb, 0x36,
	0x17, 0x8e, 0x1d, 0xc7, 0xb5, 0x8f, 0x3f, 0xe2, 0xf0, 0x2d, 0x0a, 0xa1, 0x78, 0xd7, 0x7a, 0xc2,
	0x36, 0x04, 0xc5, 0xa6, 0x14, 0x6e, 0x57, 0x60, 0xc0, 0x7e, 0x2c, 0x89, 0x83, 0xc6, 0x40, 0x96,
	0x14, 0x4c, 0xdb, 0x0d, 0xfd, 0xc7, 0x48, 0xb4, 0x4b, 0x44, 0x22, 0x35, 0x92, 0xca, 0x9a, 0xde,
	0xca, 0x6e, 0x6d, 0x37, 0xb0, 0x0b, 0x18, 0xf6, 0x77, 0xe0, 0x87, 0x64, 0xc9, 0xc9, 0xfe, 0xf1,
	0x3c, 0xcf, 0xc3, 0xe3, 0xc3, 0xf3, 0x25, 0xc3, 0x6e, 0x42, 0x6f, 0x62, 0xc1, 0x97, 0x6c, 0x15,
	0xe5, 0x52, 0x68, 0xf1, 0xd8, 0x01, 0x59, 0x26, 0x78, 0x09, 0x90, 0x3c, 0x6f, 0x28, 0xd0, 0x15,
	0x51, 0x54, 0xa8, 0xe6, 0x2d, 0x4e, 0x75, 0x03, 0x18, 0x29, 0x2d, 0x24, 0x59, 0xd1, 0xca, 0xa4,
	0xf2, 0x86, 0xc5, 0x95, 0xc9, 0xa9, 0x66, 0x5c, 0x69, 0x67, 0x86, 0xaf, 0xa1, 0xff, 0x33, 0xc9,
	0x2f, 0xa9, 0xbc, 0xa1, 0x12, 0x3d, 0x86, 0xde, 0x39, 0xc9, 0xe8, 0x85, 0x5c, 0xe4, 0x41, 0x6b,
	0xda, 0xda, 0xef, 0xe3, 0xca, 0x46, 0x4f, 0x00, 0x8e, 0x25, 0x4d, 0x28, 0xd7, 0x8c, 0xa4, 0x41,
	0xdb, 0xb2, 0x35, 0x24, 0xfc, 0x01, 0xfa, 0x1f, 0x68, 0xb2, 0x76, 0x74, 0x2a, 0x94, 0x36, 0x97,
	0x4b, 0x47, 0xa5, 0x8d, 0x26, 0xb0, 0x7d, 0xb2, 0x98, 0x07, 0xed, 0xe9, 0xf6, 0x7e, 0x1f, 0x9b,
	0x63, 0xf8, 0x4f, 0x1b, 0xf6, 0xe6, 0xd4, 0xc4, 0x78, 0xc6, 0x54, 0x3e, 0xa7, 0x9a, 0xb0, 0x54,
	0xa1, 0x19, 0x8c, 0x8d, 0x59, 0x45, 0xa7, 0x82, 0xd6, 0x74, 0x7b, 0x7f, 0x30, 0x83, 0xa8, 0x82,
	0xf0, 0x86, 0x02, 0x85, 0x30, 0x34, 0xc8, 0x82, 0x2b, 0x4d, 0x78, 0x4c, 0x6d, 0x98, 0x23, 0xdc,
	0xc0, 0xca, 0xdf, 0xef, 0xd8, 0xb0, 0xcc, 0xd1, 0x3c, 0xed, 0x64, 0x31, 0x3f, 0x25, 0xea, 0xe3,
	0x19, 0xe5, 0xc1, 0x8e, 0xbd, 0x53, 0x43, 0xd0, 0x33, 0x80, 0xea, 0x69, 0x2a, 0xe8, 0xfa, 0x28,
	0x2a, 0x08, 0xd7, 0x58, 0xf4, 0x1c, 0xfe, 0x77, 0xc2, 0x92, 0xc3, 0x34, 0x15, 0x31, 0xd1, 0x4c,
	0xf0, 0x37, 0x92, 0x2e, 0xd9, 0xa7, 0xa0, 0x37, 0x6d, 0xed, 0x0f, 0xf1, 0x7d, 0x14, 0x7a, 0x05,
	0x8f, 0xee, 0x81, 0x4d, 0x24, 0x7d, 0x1b, 0xc9, 0x7f, 0xb0, 0xb6, 0x20, 0x29, 0xa3, 0x5c, 0x1f,
	0x26, 0x89, 0x0c, 0xc0, 0x17, 0xa4, 0x42, 0x4c, 0x2e, 0x4e, 0x3e, 0xe5, 0x54, 0xb2, 0x8c, 0x72,
	0x4d, 0xd2, 0xe0, 0xe1, 0xb4, 0xb5, 0xdf, 0xc3, 0x0d, 0x2c, 0x5c, 0xc2, 0xd0, 0x25, 0xfe, 0x22,
	0x57, 0xc7, 0x59, 0x82, 0x02, 0x78, 0x10, 0x8b, 0x82, 0x6b, 0x2a, 0x7d, 0xea, 0x4a, 0xd3, 0x78,
	0x4b, 0xa8, 0x62, 0x92, 0x26, 0x97, 0x9a, 0x68, 0x1a, 0x6c, 0x3b, 0x6f, 0x75, 0xcc, 0xdc, 0x16,
	0xb9, 0x7a, 0xcb, 0x32, 0xea, 0xb3, 0x5b, 0x9a, 0xe1, 0x1f, 0x2d, 0xd8, 0x55, 0xbf, 0x1c, 0x26,
	0x24, 0xd7, 0x54, 0xbe, 0x21, 0x92, 0x64, 0x0a, 0x7d, 0x05, 0x3b, 0xe4, 0xed, 0x6d, 0xee, 0x1a,
	0x64, 0x3c, 0x1b, 0x47, 0x95, 0xc0, 0xa0, 0xd8, 0x91, 0xe8, 0x5b, 0xd8, 0x2b, 0x78, 0x42, 0x65,
	0x4a, 0x6e, 0x17, 0x26, 0x90, 0x25, 0x89, 0xa9, 0xcd, 0x66, 0x1f, 0xdf, 0x25, 0xd0, 0x23, 0xe8,
	0xde, 0xa4, 0x84, 0x2f, 0x12, 0x9f, 0x3b, 0x6f, 0xa1, 0xff, 0x43, 0xff, 0x4a, 0xf0, 0x64, 0x25,
	0x45, 0x91, 0x07, 0x60, 0x3b, 0x6f, 0x0d, 0x84, 0x7f, 0xb5, 0x60, 0x74, 0x79, 0xab, 0x34, 0xcd,
	0x7c, 0x00, 0x08, 0x41, 0x87, 0xaf, 0x7b, 0xd7, 0x9e, 0xd1, 0x4b, 0x18, 0x12, 0x53, 0x06, 0xdf,
	0x9f, 0x36, 0x9f, 0x83, 0xd9, 0x24, 0xda, 0x78, 0x17, 0x6e, 0xa8, 0x4c, 0x95, 0x96, 0x92, 0xd2,
	0x77, 0x79, 0xca, 0xf8, 0xb5, 0x4d, 0x6a, 0x0f, 0xd7, 0x10, 0x13, 0x71, 0xe1, 0x38, 0x97, 0x51,
	0x6f, 0xa1, 0x29, 0x0c, 0x38, 0xd5, 0xbf, 0x0b, 0x79, 0xfd, 0xee, 0x5d, 0xd5, 0xad, 0x75, 0xc8,
	0xc4, 0x48, 0x4c, 0xe5, 0x77, 0x5c, 0x8c, 0xe6, 0x6c, 0x6e, 0xa5, 0x62, 0xc5, 0x62, 0x92, 0xda,
	0xd1, 0xeb, 0xba, 0x5b, 0x35, 0x28, 0xfc, 0x73, 0x07, 0x46, 0x27, 0xc9, 0x8a, 0xce, 0xe9, 0xcd,
	0xb1, 0x5d, 0x1a, 0xe8, 0x29, 0xb4, 0x59, 0x62, 0x5f, 0x3a, 0x98, 0xed, 0x46, 0xc6, 0x35, 0xe1,
	0xc9, 0x7b, 0x2a, 0x15, 0x13, 0x1c, 0xb7, 0x59, 0x82, 0xf6, 0xed, 0xa6, 0x72, 0xea, 0xcb, 0x8f,
	0x64, 0xf6, 0xfd, 0x2b, 0xfb, 0x8e, 0x21, 0xde, 0x84, 0x51, 0x04, 0x68, 0x0d, 0xb1, 0x15, 0x27,
	0xba, 0x90, 0xae, 0x55, 0x86, 0xf8, 0x1e, 0x06, 0x7d, 0x03, 0x1d, 0x92, 0xe7, 0x2a, 0xe8, 0xd8,
	0x91, 0x42, 0xd1, 0x61, 0x5e, 0x8d, 0xa9, 0x93, 0x62, 0xcb, 0xa3, 0x67, 0xd0, 0xf3, 0x2f, 0x57,
	0xc1, 0x8e, 0xd5, 0x8e, 0xa3, 0x73, 0x07, 0x78, 0x5d, 0xc5, 0xa3, 0xe7, 0x00, 0x09, 0xd1, 0xc4,
	0xec, 0x40, 0x5a, 0x0e, 0xeb, 0x24, 0x9a, 0x97, 0x90, 0xd7, 0xd7, 0x34, 0x28, 0x82, 0x5e, 0x6a,
	0x17, 0xc4, 0x52, 0x04, 0x0f, 0x6c, 0x1a, 0x50, 0x74, 0x67, 0x1d, 0xe1, 0x4a, 0x83, 0xbe, 0x84,
	0x8e, 0x59, 0xc3, 0x41, 0xcf, 0xfa, 0x1e, 0x45, 0x47, 0x44, 0xd1, 0x8b, 0xcb, 0x32, 0x60, 0x43,
	0xa1, 0xaf, 0xa1, 0x2b, 0xe9, 0x95, 0x10, 0xda, 0xf6, 0xa1, 0x11, 0xd5, 0xc7, 0x0c, 0x7b, 0xd2,
	0xc8, 0xae, 0x48, 0x7c, 0x6d, 0x7b, 0xf2, 0x3e, 0x99, 0x23, 0xd1, 0x77, 0x30, 0x70, 0x0b, 0x7e,
	0xa1, 0x69, 0xa6, 0x82, 0x81, 0xfd, 0xdd, 0x41, 0x74, 0x5c, 0x61, 0xb8, 0xce, 0xa3, 0x1f, 0x61,
	0x4f, 0xd5, 0xbb, 0xf9, 0x8c, 0x29, 0x1d, 0x0c, 0x7d, 0xda, 0x1a, 0x7d, 0x8e, 0xef, 0x0a, 0xd1,
	0x0c, 0x7a, 0xfe, 0x83, 0xa1, 0x82, 0x91, 0xbd, 0xf4, 0x28, 0xba, 0x74, 0xc0, 0x46, 0x6d, 0x2a,
	0x9d, 0x59, 0x0e, 0x19, 0xe1, 0xc5, 0x92, 0xc4, 0xa6, 0xac, 0x32, 0x18, 0xdb, 0xbe, 0x6b, 0x60,
	0xa6, 0x35, 0x73, 0x29, 0x92, 0x22, 0x76, 0x5f, 0x85, 0x5d, 0xd7, 0x9a, 0x35, 0x08, 0x1d, 0xc1,
	0xc4, 0x57, 0xb1, 0xfc, 0x21, 0x15, 0x4c, 0x7c, 0x04, 0xe7, 0x4d, 0xc2, 0x47, 0x70, 0x47, 0x1f,
	0xfe, 0xdd, 0x02, 0x58, 0xe7, 0xc5, 0xec, 0xfa, 0x6b, 0x7a, 0xeb, 0xc7, 0xd8, 0x1c, 0xd1, 0x43,
	0xd8, 0xb9, 0x21, 0x69, 0x41, 0xfd, 0x17, 0xcc, 0x19, 0xe8, 0x89, 0xd9, 0x0f, 0x22, 0x7d, 0x6f,
	0x19, 0x3b, 0x88, 0xa7, 0x5b, 0x78, 0x0d, 0xa1, 0x10, 0x06, 0x05, 0xe3, 0xfa, 0xc5, 0xcc, 0x29,
	0xcc, 0x34, 0x8e, 0x4e, 0xb7, 0x70, 0x1d, 0x2c, 0x35, 0xaf, 0x5e, 0x3a, 0x8d, 0x19, 0xcb, 0x4e,
	0xa9, 0xf1, 0x20, 0x9a, 0x02, 0x2c, 0x53, 0x41, 0xb4, 0x93, 0x98, 0xf1, 0x6c, 0x9f, 0x6e, 0xe1,
	0x1a, 0x66, 0xbc, 0x28, 0x2d, 0x19, 0x5f, 0x39, 0x89, 0xe9, 0xc7, 0xbe, 0xf1, 0x52, 0x03, 0x8f,
	0xf6, 0x60, 0x77, 0x5d, 0x6f, 0x0b, 0x85, 0x07, 0x30, 0xf2, 0x39, 0xa1, 0xbf, 0x15, 0x54, 0x69,
	0xb3, 0x77, 0x9c, 0xc6, 0x7c, 0xc4, 0x7c, 0x02, 0x6a, 0x48, 0xf8, 0x2b, 0x8c, 0xcb, 0x0b, 0x2a,
	0x17, 0x5c, 0x99, 0x61, 0xec, 0x3a, 0xde, 0xef, 0x82, 0x71, 0xd4, 0xd8, 0x13, 0xd8, 0xb3, 0x1b,
	0x9e, 0xdb, 0x9b, 0x9e, 0x9f, 0x1d, 0xc0, 0xa8, 0xb1, 0xc9, 0x11, 0x40, 0x77, 0xf1, 0xfa, 0xfc,
	0x02, 0x9f, 0x4c, 0xb6, 0x50, 0x0f, 0x3a, 0xef, 0xcf, 0x0e, 0xcf, 0x27, 0x2d, 0x73, 0x3a, 0xba,
	0x38, 0x9f, 0x4f, 0xda, 0x47, 0x3f, 0xc1, 0xd3, 0x58, 0x64, 0xd1, 0x67, 0x9a, 0xd0, 0x84, 0x44,
	0x71, 0x2a, 0x8a, 0x24, 0x2a, 0x1a, 0x7f, 0x5a, 0x3e, 0x7c, 0xb1, 0x62, 0xfa, 0x63, 0x71, 0x15,
	0xc5, 0x22, 0x3b, 0x70, 0xba, 0x03, 0x92, 0xb3, 0x83, 0xcf, 0xee, 0x67, 0xaf, 0xba, 0x56, 0xf5,
	0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x2c, 0xde, 0x00, 0x52, 0x09, 0x00, 0x00,
}

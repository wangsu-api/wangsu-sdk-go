package appadomain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryAppaDomainForTerraformRequest struct {
}

func (s QueryAppaDomainForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformRequest) GoString() string {
	return s.String()
}

type QueryAppaDomainForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.", "zh_CN":"接口响应数据"}
	Data *QueryAppaDomainForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryAppaDomainForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponse) SetCode(v string) *QueryAppaDomainForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponse) SetMessage(v string) *QueryAppaDomainForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponse) SetData(v *QueryAppaDomainForTerraformResponseData) *QueryAppaDomainForTerraformResponse {
	s.Data = v
	return s
}

type QueryAppaDomainForTerraformResponseData struct {
	// {"en":"Domain ID", "zh_CN":"加速域名ID"}
	DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"Domain name you want to accelerate.A generic domain name is supported, starting with the symbol '.', such as .example.com.", "zh_CN":"加速域名。支持泛域名，以符号“.”开头，例如：.example.com。"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
	// {"en":"The service type of the accelerated domain name. The value can be:
	// appa: Application Acceleration", "zh_CN":"加速域名的服务类型，可选值为：
	// appa：应用加速/s-appa：应用安全加速解决方案"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"cname", "zh_CN":" 服务域名"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
	// {"en":"status", "zh_CN":" 状态"}
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// {"en":"is enabled", "zh_CN":" 是否启用"}
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
	// {"en":"Origin configuration.
	// Example:
	// 'originConfig':[
	//   {'level':1,'strategy':'robin','origin':
	//     [
	//       {'originIp':'1.1.1.1','weight':10},
	//       {'originIp':'2.2.2.2','weight':20}
	//     ]
	//   },
	//   {'level':2,'strategy':'quick','origin':
	//    [{'originIp':'3.3.3.3','weight':10}]
	//   }
	// ]", "zh_CN":"源站配置
	// 示例
	// 'originConfig':[
	//   {'level':1,'strategy':'robin','origin':
	//     [
	//       {'originIp':'1.1.1.1','weight':10},
	//       {'originIp':'2.2.2.2','weight':20}
	//     ]
	//   },
	//   {'level':2,'strategy':'quick','origin':
	//    [{'originIp':'3.3.3.3','weight':10}]
	//   }
	// ]"}
	OriginConfig []*QueryAppaDomainForTerraformResponseDataOriginConfig `json:"originConfig,omitempty" xml:"originConfig,omitempty" require:"true" type:"Repeated"`
	// {'en':'HTTP port. The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:httpPorts:["9001"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port and TCP port must be unique.3. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'HTTP端口，取值范围为1-65535的整数，可配置多个，格式如：httpPorts:["9001"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口不能重复。3、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	HttpPorts []*string `json:"httpPorts,omitempty" xml:"httpPorts,omitempty" require:"true" type:"Repeated"`
	// {'en':'HTTPS port. The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:httpPorts:["9002","9003"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port and TCP port must be unique.3. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'HTTPS端口，取值范围为1-65535的整数，可配置多个，格式如：httpsPorts:["9002","9003"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口不能重复。3、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	HttpsPorts []*string `json:"httpsPorts,omitempty" xml:"httpsPorts,omitempty" require:"true" type:"Repeated"`
	// {'en':'TCP port.The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:tcpPorts:["9005-9007"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port and TCP port must be unique.3. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'TCP端口，取值范围为1-65535的整数，可配置多个，格式如：tcpPorts:["9005-9007"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口不能重复。3、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	TcpPorts []*string `json:"tcpPorts,omitempty" xml:"tcpPorts,omitempty" type:"Repeated"`
	// {'en':'UDP port.The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:udpPorts:["9008-9009"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'UDP端口，取值范围为1-65535的整数，可配置多个，格式如：udpPorts:["9008-9009"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	UdpPorts []*string `json:"udpPorts,omitempty" xml:"udpPorts,omitempty" type:"Repeated"`
	// {"en":"Carry client IP configuration.", "zh_CN":"携带用户IP回源配置配置。"}
	CarryClientIp *QueryAppaDomainForTerraformResponseDataCarryClientIp `json:"carryClientIp,omitempty" xml:"carryClientIp,omitempty" require:"true" type:"Struct"`
}

func (s QueryAppaDomainForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponseData) SetDomainId(v int64) *QueryAppaDomainForTerraformResponseData {
	s.DomainId = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetDomainName(v string) *QueryAppaDomainForTerraformResponseData {
	s.DomainName = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetServiceType(v string) *QueryAppaDomainForTerraformResponseData {
	s.ServiceType = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetCname(v string) *QueryAppaDomainForTerraformResponseData {
	s.Cname = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetStatus(v string) *QueryAppaDomainForTerraformResponseData {
	s.Status = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetEnabled(v bool) *QueryAppaDomainForTerraformResponseData {
	s.Enabled = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetOriginConfig(v []*QueryAppaDomainForTerraformResponseDataOriginConfig) *QueryAppaDomainForTerraformResponseData {
	s.OriginConfig = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetHttpPorts(v []*string) *QueryAppaDomainForTerraformResponseData {
	s.HttpPorts = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetHttpsPorts(v []*string) *QueryAppaDomainForTerraformResponseData {
	s.HttpsPorts = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetTcpPorts(v []*string) *QueryAppaDomainForTerraformResponseData {
	s.TcpPorts = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetUdpPorts(v []*string) *QueryAppaDomainForTerraformResponseData {
	s.UdpPorts = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseData) SetCarryClientIp(v *QueryAppaDomainForTerraformResponseDataCarryClientIp) *QueryAppaDomainForTerraformResponseData {
	s.CarryClientIp = v
	return s
}

type QueryAppaDomainForTerraformResponseDataOriginConfig struct {
	// {"en":"The level of the origin, which value can be an integer ranging from 1 to 5. Note:1. Must be configured level by level start from level 1. The same level cannot be configured repeatedly.2. The lower the value, the higher the priority.", "zh_CN":"层级，可选值为1-5的整数。注意：1、必须从层级1开始逐级配置，相同层级不能重复配置。2、数值越低，优先级越高。"}
	Level *int32 `json:"level,omitempty" xml:"level,omitempty" require:"true"`
	// {"en":"Origin selection strategy supports fast, robin and hash. The value can be: fast: Fast strategy, robin: Robin strategy,hash: Hash strategy", "zh_CN":"选源策略，支持快速、轮询、哈希，可选值为:fast：快速,robin：轮询,hash：哈希"}
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty" require:"true"`
	// {"en":"Origin information of a certain level. A level can be configured with multiple origin IP addresses or domain names.Example:'origin':[{'originIp':'1.1.1.1','weight':10},{'originIp':'2.2.2.2','weight':20}]", "zh_CN":"某个层级的源信息。一个层级可以配置多个回源IP/域名。示例：'origin':[{'originIp':'1.1.1.1','weight':10},{'originIp':'2.2.2.2','weight':20}]"}
	Origin []*QueryAppaDomainForTerraformResponseDataOriginConfigOrigin `json:"origin,omitempty" xml:"origin,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppaDomainForTerraformResponseDataOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponseDataOriginConfig) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponseDataOriginConfig) SetLevel(v int32) *QueryAppaDomainForTerraformResponseDataOriginConfig {
	s.Level = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataOriginConfig) SetStrategy(v string) *QueryAppaDomainForTerraformResponseDataOriginConfig {
	s.Strategy = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataOriginConfig) SetOrigin(v []*QueryAppaDomainForTerraformResponseDataOriginConfigOrigin) *QueryAppaDomainForTerraformResponseDataOriginConfig {
	s.Origin = v
	return s
}

type QueryAppaDomainForTerraformResponseDataOriginConfigOrigin struct {
	// {"en":"Origin address, which can be an IP or domain name.", "zh_CN":"回源IP/域名，可配置一个IP或域名。"}
	OriginIp *string `json:"originIp,omitempty" xml:"originIp,omitempty" require:"true"`
	// {"en":"Weight, which is only useful for robin strategy. The value is an integer ranging from 1 to 10000. If this parameter is not specified, the default value is 10.", "zh_CN":"权重，只对轮询策略有用。取值范围为1-10000的整数，不填默认为10。"}
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s QueryAppaDomainForTerraformResponseDataOriginConfigOrigin) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponseDataOriginConfigOrigin) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponseDataOriginConfigOrigin) SetOriginIp(v string) *QueryAppaDomainForTerraformResponseDataOriginConfigOrigin {
	s.OriginIp = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataOriginConfigOrigin) SetWeight(v int32) *QueryAppaDomainForTerraformResponseDataOriginConfigOrigin {
	s.Weight = &v
	return s
}

type QueryAppaDomainForTerraformResponseDataCarryClientIp struct {
	// {"en":"Carry client IP configuration for TCP.", "zh_CN":"TCP协议携带用户IP回源功能配置。"}
	TcpCarryClientIp *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp `json:"tcpCarryClientIp,omitempty" xml:"tcpCarryClientIp,omitempty" require:"true" type:"Struct"`
	// {"en":"Carry client IP configuration for HTTP.", "zh_CN":"HTTP协议携带用户IP回源功能。"}
	HttpCarryClientIp *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp `json:"httpCarryClientIp,omitempty" xml:"httpCarryClientIp,omitempty" require:"true" type:"Struct"`
}

func (s QueryAppaDomainForTerraformResponseDataCarryClientIp) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponseDataCarryClientIp) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIp) SetTcpCarryClientIp(v *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) *QueryAppaDomainForTerraformResponseDataCarryClientIp {
	s.TcpCarryClientIp = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIp) SetHttpCarryClientIp(v *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp) *QueryAppaDomainForTerraformResponseDataCarryClientIp {
	s.HttpCarryClientIp = v
	return s
}

type QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp struct {
	// {"en":"The function switch of TCP carry client IP.
	// true: Enabled
	// false: Disabled", "zh_CN":"TCP协议携带用户IP回源功能开关。
	// true：开启
	// false：关闭"}
	CarryClientIpEnabled *string `json:"carryClientIpEnabled,omitempty" xml:"carryClientIpEnabled,omitempty" require:"true"`
	// {"en":"The protocol of TCP carry client IP.
	// ws: The client IP is carried through TCP Option 78. The origin can get the client IP through the SDK provided by Wangsu or the F5 device.
	// toa: Carry client IP through TOA.
	// pp: Carry client IP through Proxy Protocol.", "zh_CN":"TCP协议携带用户IP回源所使用的协议。
	// ws：通过tcp option 78携带用户IP，源站可通过F5设备或网宿提供的SDK获取用户IP。
	// toa：通过TOA携带用户IP回源。
	// pp：通过Proxy Protocol携带用户IP回源。"}
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
	// {"en":"The number of packets carrying client IP other than SYN packet, when using ws or toa to carry client IP.", "zh_CN":"使用ws或toa携带用户IP回源时，除SYN包外携带用户IP的数据包个数。"}
	PacketNum *string `json:"packetNum,omitempty" xml:"packetNum,omitempty" require:"true"`
	// {"en":"The mode of TCP carry client IP by ws or toa.
	// 0: Only SYN packet carry client IP.
	// 1: SYN packet and the first N packets carry client IP,N is specified by packetNum.
	// 2: SYN packet, the first N packets, RST, PSH, URG and FIN/ACK packet carry client IP, N is specified by packetNum.", "zh_CN":"使用ws或toa携带用户IP回源的模式。
	// 0：仅SYN包携带用户IP回源。
	// 1：SYN包及前N个数据包携带用户IP回源，N由carryClientIpNum指定。
	// 2：SYN包、前N个数据包 、RST、PSH、URG、FIN/ACK包携带用户IP回源，N由carryClientIpNum指定。"}
	Mode *int `json:"mode,omitempty" xml:"mode,omitempty" require:"true"`
	// {"en":"The opcode to carry client IP  by TOA.", "zh_CN":"通过TOA携带用户IP回源时所使用的opcode。"}
	TcpOptionCode *string `json:"tcpOptionCode,omitempty" xml:"tcpOptionCode,omitempty" require:"true"`
}

func (s QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) SetCarryClientIpEnabled(v string) *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp {
	s.CarryClientIpEnabled = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) SetProtocol(v string) *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp {
	s.Protocol = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) SetPacketNum(v string) *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp {
	s.PacketNum = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) SetMode(v int) *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp {
	s.Mode = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp) SetTcpOptionCode(v string) *QueryAppaDomainForTerraformResponseDataCarryClientIpTcpCarryClientIp {
	s.TcpOptionCode = &v
	return s
}

type QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp struct {
	// {"en":"The function switch of carry client IP through the HTTP header.
	// true: Enabled.
	// false: Disabled.", "zh_CN":"HTTP协议通过HTTP头部携带用户IP回源的功能开关。
	// true：开启
	// false：关闭"}
	CarryClientIpEnabled *string `json:"carryClientIpEnabled,omitempty" xml:"carryClientIpEnabled,omitempty" require:"true"`
	// {"en":"Specify which acceleration ports to carry client IP, the function of carry client IP through the HTTP header is enabled.", "zh_CN":"当开启通过HTTP头部携带用户IP回源时，指定哪些加速端口传递用户IP回源。"}
	Ports []*string `json:"ports,omitempty" xml:"ports,omitempty" require:"true" type:"Repeated"`
	// {"en":"The field name of HTTP header to carry client IP, such as X-Forwarded-For, Cdn-Src-Ip.", "zh_CN":"HTTP头部携带IP字段的名称，如X-Forwarded-For、Cdn-Src-Ip。"}
	HttpHeaderName *string `json:"httpHeaderName,omitempty" xml:"httpHeaderName,omitempty" require:"true"`
}

func (s QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp) String() string {
	return tea.Prettify(s)
}

func (s QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp) GoString() string {
	return s.String()
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp) SetCarryClientIpEnabled(v string) *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp {
	s.CarryClientIpEnabled = &v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp) SetPorts(v []*string) *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp {
	s.Ports = v
	return s
}

func (s *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp) SetHttpHeaderName(v string) *QueryAppaDomainForTerraformResponseDataCarryClientIpHttpCarryClientIp {
	s.HttpHeaderName = &v
	return s
}

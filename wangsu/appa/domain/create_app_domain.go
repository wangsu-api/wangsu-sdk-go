package appadomain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type AddAppaDomainForTerraformRequest struct {
	// {"en":"Domain name you want to accelerate.A generic domain name is supported, starting with the symbol '.', such as .example.com.", "zh_CN":"加速域名。支持泛域名，以符号“.”开头，例如：.example.com。"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
	// {"en":"The service type of the accelerated domain name. The value can be:
	// appa: Application Acceleration", "zh_CN":"加速域名的服务类型，可选值为：
	// appa：应用加速/s-appa：应用安全加速解决方案"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
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
	OriginConfig []*AddAppaDomainForTerraformRequestOriginConfig `json:"originConfig,omitempty" xml:"originConfig,omitempty" require:"true" type:"Repeated"`
	// {'en':'HTTP port. The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:httpPorts:["9001"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port and TCP port must be unique.3. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'HTTP端口，取值范围为1-65535的整数，可配置多个，格式如：httpPorts:["9001"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口不能重复。3、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	HttpPorts []*string `json:"httpPorts,omitempty" xml:"httpPorts,omitempty" type:"Repeated"`
	// {'en':'HTTPS port. The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:httpPorts:["9002","9003"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port and TCP port must be unique.3. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'HTTPS端口，取值范围为1-65535的整数，可配置多个，格式如：httpsPorts:["9002","9003"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口不能重复。3、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	HttpsPorts []*string `json:"httpsPorts,omitempty" xml:"httpsPorts,omitempty" type:"Repeated"`
	// {'en':'TCP port.The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:tcpPorts:["9005-9007"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port and TCP port must be unique.3. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'TCP端口，取值范围为1-65535的整数，可配置多个，格式如：tcpPorts:["9005-9007"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口不能重复。3、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	TcpPorts []*string `json:"tcpPorts,omitempty" xml:"tcpPorts,omitempty" type:"Repeated"`
	// {'en':'UDP port.The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:udpPorts:["9008-9009"]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. At least one HTTP port or HTTPS port or TCP port or UDP port must be configured.', 'zh_CN':'UDP端口，取值范围为1-65535的整数，可配置多个，格式如：udpPorts:["9008-9009"]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口和TCP端口和UDP端口必须至少配置一个。'}
	UdpPorts []*string `json:"udpPorts,omitempty" xml:"udpPorts,omitempty" type:"Repeated"`
}

func (s AddAppaDomainForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppaDomainForTerraformRequest) GoString() string {
	return s.String()
}

func (s *AddAppaDomainForTerraformRequest) SetDomainName(v string) *AddAppaDomainForTerraformRequest {
	s.DomainName = &v
	return s
}

func (s *AddAppaDomainForTerraformRequest) SetServiceType(v string) *AddAppaDomainForTerraformRequest {
	s.ServiceType = &v
	return s
}

func (s *AddAppaDomainForTerraformRequest) SetOriginConfig(v []*AddAppaDomainForTerraformRequestOriginConfig) *AddAppaDomainForTerraformRequest {
	s.OriginConfig = v
	return s
}

func (s *AddAppaDomainForTerraformRequest) SetHttpPorts(v []*string) *AddAppaDomainForTerraformRequest {
	s.HttpPorts = v
	return s
}

func (s *AddAppaDomainForTerraformRequest) SetHttpsPorts(v []*string) *AddAppaDomainForTerraformRequest {
	s.HttpsPorts = v
	return s
}

func (s *AddAppaDomainForTerraformRequest) SetTcpPorts(v []*string) *AddAppaDomainForTerraformRequest {
	s.TcpPorts = v
	return s
}

func (s *AddAppaDomainForTerraformRequest) SetUdpPorts(v []*string) *AddAppaDomainForTerraformRequest {
	s.UdpPorts = v
	return s
}

type AddAppaDomainForTerraformRequestOriginConfig struct {
	// {"en":"The level of the origin, which value can be an integer ranging from 1 to 5. Note:1. Must be configured level by level start from level 1. The same level cannot be configured repeatedly.2. The lower the value, the higher the priority.", "zh_CN":"层级，可选值为1-5的整数。注意：1、必须从层级1开始逐级配置，相同层级不能重复配置。2、数值越低，优先级越高。"}
	Level *int32 `json:"level,omitempty" xml:"level,omitempty" require:"true"`
	// {"en":"Origin selection strategy supports fast, robin and hash. The value can be: fast: Fast strategy, robin: Robin strategy,hash: Hash strategy", "zh_CN":"选源策略，支持快速、轮询、哈希，可选值为:fast：快速,robin：轮询,hash：哈希"}
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty" require:"true"`
	// {"en":"Origin information of a certain level. A level can be configured with multiple origin IP addresses or domain names.Example:'origin':[{'originIp':'1.1.1.1','weight':10},{'originIp':'2.2.2.2','weight':20}]", "zh_CN":"某个层级的源信息。一个层级可以配置多个回源IP/域名。示例：'origin':[{'originIp':'1.1.1.1','weight':10},{'originIp':'2.2.2.2','weight':20}]"}
	Origin []*AddAppaDomainForTerraformRequestOriginConfigOrigin `json:"origin,omitempty" xml:"origin,omitempty" require:"true" type:"Repeated"`
}

func (s AddAppaDomainForTerraformRequestOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s AddAppaDomainForTerraformRequestOriginConfig) GoString() string {
	return s.String()
}

func (s *AddAppaDomainForTerraformRequestOriginConfig) SetLevel(v int32) *AddAppaDomainForTerraformRequestOriginConfig {
	s.Level = &v
	return s
}

func (s *AddAppaDomainForTerraformRequestOriginConfig) SetStrategy(v string) *AddAppaDomainForTerraformRequestOriginConfig {
	s.Strategy = &v
	return s
}

func (s *AddAppaDomainForTerraformRequestOriginConfig) SetOrigin(v []*AddAppaDomainForTerraformRequestOriginConfigOrigin) *AddAppaDomainForTerraformRequestOriginConfig {
	s.Origin = v
	return s
}

type AddAppaDomainForTerraformRequestOriginConfigOrigin struct {
	// {"en":"Origin address, which can be an IP or domain name.", "zh_CN":"回源IP/域名，可配置一个IP或域名。"}
	OriginIp *string `json:"originIp,omitempty" xml:"originIp,omitempty" require:"true"`
	// {"en":"Weight, which is only useful for robin strategy. The value is an integer ranging from 1 to 10000. If this parameter is not specified, the default value is 10.", "zh_CN":"权重，只对轮询策略有用。取值范围为1-10000的整数，不填默认为10。"}
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s AddAppaDomainForTerraformRequestOriginConfigOrigin) String() string {
	return tea.Prettify(s)
}

func (s AddAppaDomainForTerraformRequestOriginConfigOrigin) GoString() string {
	return s.String()
}

func (s *AddAppaDomainForTerraformRequestOriginConfigOrigin) SetOriginIp(v string) *AddAppaDomainForTerraformRequestOriginConfigOrigin {
	s.OriginIp = &v
	return s
}

func (s *AddAppaDomainForTerraformRequestOriginConfigOrigin) SetWeight(v int32) *AddAppaDomainForTerraformRequestOriginConfigOrigin {
	s.Weight = &v
	return s
}

type AddAppaDomainForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.", "zh_CN":"接口响应数据"}
	Data *AddAppaDomainForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s AddAppaDomainForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppaDomainForTerraformResponse) GoString() string {
	return s.String()
}

func (s *AddAppaDomainForTerraformResponse) SetCode(v string) *AddAppaDomainForTerraformResponse {
	s.Code = &v
	return s
}

func (s *AddAppaDomainForTerraformResponse) SetMessage(v string) *AddAppaDomainForTerraformResponse {
	s.Message = &v
	return s
}

func (s *AddAppaDomainForTerraformResponse) SetData(v *AddAppaDomainForTerraformResponseData) *AddAppaDomainForTerraformResponse {
	s.Data = v
	return s
}

type AddAppaDomainForTerraformResponseData struct {
	// {"en":"domain ID", "zh_CN":"域名ID"}
	DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"cname", "zh_CN":"服务域名"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
}

func (s AddAppaDomainForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddAppaDomainForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *AddAppaDomainForTerraformResponseData) SetDomainId(v int64) *AddAppaDomainForTerraformResponseData {
	s.DomainId = &v
	return s
}

func (s *AddAppaDomainForTerraformResponseData) SetCname(v string) *AddAppaDomainForTerraformResponseData {
	s.Cname = &v
	return s
}

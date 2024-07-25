package appadomain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type UpdateAppaDomainForTerraformRequest struct {
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
	OriginConfig []*UpdateAppaDomainForTerraformRequestOriginConfig `json:"originConfig" xml:"originConfig" type:"Repeated"`
	// {"en":"HTTP port. The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:httpPorts:[1000,1001]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port must be unique.3. At least one HTTP port or HTTPS port must be configured.", "zh_CN":"HTTP端口，取值范围为1-65535的整数，可配置多个，格式如：httpPorts:[1000,1001]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口不能重复。3、HTTP端口和HTTPS端口必须至少配置一个。"}
	HttpPorts []*string `json:"httpPorts" xml:"httpPorts" type:"Repeated"`
	// {"en":"HTTPS port. The value is an integer ranging from 1 to 65535. Multiple ports are supported and can be configured in the following format:httpPorts:[1000,1001]Note: 1. Ports 2012, 2323, 2443, 4031, 12012, 20121, 57891, 62016, 65383, and 65529 do not support.2. The HTTP port and HTTPS port must be unique.3. At least one HTTP port or HTTPS port must be configured.", "zh_CN":"HTTPS端口，取值范围为1-65535的整数，可配置多个，格式如：httpsPorts:[2000,2001]注意：1、端口2012、2323、2443、4031、12012、20121、57891、62016、65383、65529不支持配置。2、HTTP端口和HTTPS端口不能重复。3、HTTP端口和HTTPS端口必须至少配置一个。"}
	HttpsPorts []*string `json:"httpsPorts" xml:"httpsPorts" type:"Repeated"`
}

func (s UpdateAppaDomainForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppaDomainForTerraformRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppaDomainForTerraformRequest) SetOriginConfig(v []*UpdateAppaDomainForTerraformRequestOriginConfig) *UpdateAppaDomainForTerraformRequest {
	s.OriginConfig = v
	return s
}

func (s *UpdateAppaDomainForTerraformRequest) SetHttpPorts(v []*string) *UpdateAppaDomainForTerraformRequest {
	s.HttpPorts = v
	return s
}

func (s *UpdateAppaDomainForTerraformRequest) SetHttpsPorts(v []*string) *UpdateAppaDomainForTerraformRequest {
	s.HttpsPorts = v
	return s
}

type UpdateAppaDomainForTerraformRequestOriginConfig struct {
	// {"en":"The level of the origin, which value can be an integer ranging from 1 to 5. Note:1. Must be configured level by level start from level 1. The same level cannot be configured repeatedly.2. The lower the value, the higher the priority.", "zh_CN":"层级，可选值为1-5的整数。注意：1、必须从层级1开始逐级配置，相同层级不能重复配置。2、数值越低，优先级越高。"}
	Level *int32 `json:"level" xml:"level" require:"true"`
	// {"en":"Origin selection strategy supports fast, robin and hash. The value can be:fast: Fast strategy, robin: Robin strategy,hash: Hash strategy", "zh_CN":"选源策略，支持快速、轮询、哈希，可选值为:fast：快速,robin：轮询,hash：哈希"}
	Strategy *string `json:"strategy" xml:"strategy" require:"true"`
	// {"en":"Origin information of a certain level. A level can be configured with multiple origin IP addresses or domain names.Example:'origin':[{'originIp':'1.1.1.1','weight':10},{'originIp':'2.2.2.2','weight':20}]", "zh_CN":"某个层级的源信息。一个层级可以配置多个回源IP/域名。示例：'origin':[{'originIp':'1.1.1.1','weight':10},{'originIp':'2.2.2.2','weight':20}]"}
	Origin []*UpdateAppaDomainForTerraformRequestOriginConfigOrigin `json:"origin" xml:"origin" require:"true" type:"Repeated"`
}

func (s UpdateAppaDomainForTerraformRequestOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppaDomainForTerraformRequestOriginConfig) GoString() string {
	return s.String()
}

func (s *UpdateAppaDomainForTerraformRequestOriginConfig) SetLevel(v int32) *UpdateAppaDomainForTerraformRequestOriginConfig {
	s.Level = &v
	return s
}

func (s *UpdateAppaDomainForTerraformRequestOriginConfig) SetStrategy(v string) *UpdateAppaDomainForTerraformRequestOriginConfig {
	s.Strategy = &v
	return s
}

func (s *UpdateAppaDomainForTerraformRequestOriginConfig) SetOrigin(v []*UpdateAppaDomainForTerraformRequestOriginConfigOrigin) *UpdateAppaDomainForTerraformRequestOriginConfig {
	s.Origin = v
	return s
}

type UpdateAppaDomainForTerraformRequestOriginConfigOrigin struct {
	// {"en":"Origin address, which can be an IP or domain name.", "zh_CN":"回源IP/域名，可配置一个IP或域名。"}
	OriginIp *string `json:"originIp" xml:"originIp" require:"true"`
	// {"en":"Weight, which is only useful for robin strategy. The value is an integer ranging from 1 to 10000. If this parameter is not specified, the default value is 10.", "zh_CN":"权重，只对轮询策略有用。取值范围为1-10000的整数，不填默认为10。"}
	Weight *int32 `json:"weight" xml:"weight"`
}

func (s UpdateAppaDomainForTerraformRequestOriginConfigOrigin) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppaDomainForTerraformRequestOriginConfigOrigin) GoString() string {
	return s.String()
}

func (s *UpdateAppaDomainForTerraformRequestOriginConfigOrigin) SetOriginIp(v string) *UpdateAppaDomainForTerraformRequestOriginConfigOrigin {
	s.OriginIp = &v
	return s
}

func (s *UpdateAppaDomainForTerraformRequestOriginConfigOrigin) SetWeight(v int32) *UpdateAppaDomainForTerraformRequestOriginConfigOrigin {
	s.Weight = &v
	return s
}

type UpdateAppaDomainForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code" xml:"code" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message" xml:"message" require:"true"`
	// {"en":"Response data.", "zh_CN":"接口响应数据"}
	Data *string `json:"data" xml:"data" require:"true"`
}

func (s UpdateAppaDomainForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppaDomainForTerraformResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppaDomainForTerraformResponse) SetCode(v string) *UpdateAppaDomainForTerraformResponse {
	s.Code = &v
	return s
}

func (s *UpdateAppaDomainForTerraformResponse) SetMessage(v string) *UpdateAppaDomainForTerraformResponse {
	s.Message = &v
	return s
}

func (s *UpdateAppaDomainForTerraformResponse) SetData(v string) *UpdateAppaDomainForTerraformResponse {
	s.Data = &v
	return s
}

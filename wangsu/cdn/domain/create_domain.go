// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDomainRequest struct {
	// {"en":"Version code , the current version is 1.0.0", "zh_CN":"版本号，当前版本号1.0.0"}
	Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"Need to access the domain name of the CDN. a generic domain name is supported, starting with the symbol '.', such as.example.com, which also contains a multilevel 'a.b.example.com'.If example.com is filed, the domain name xx.example.com does not need to be filed.", "zh_CN":"需要接入CDN的域名。支持泛域名，以符号“.”开头，如：.example.com，泛域名也包含多级“a.b.example.com”。
	// 如果example.com已备案，那么域名xx.example.com则不需要备案。"}
	DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
	// {"en":"The service type of the accelerated domain name (only one service type can be submitted at a time):
	//   web/web-https: Web page acceleration/Web page acceleration-https
	// wsa/Wsa-https: Full-station acceleration/full-station acceleration-https
	// vodstream/vod-https: on-demand acceleration/on-demand acceleration-https
	// download/dl-https: Download Acceleration/Download Acceleration-https
	// livestream/live-https/cloudv-live: livestream acceleration
	// v6sa/osv6: IPv6 Security&Acceleration Solution/IPv6 One-stop Solution
	// Note:
	// 1. the https in the code, such as web-https does not represent immediate support for https access, you need to upload the certificate to support https.
	//   ", "zh_CN":"加速域名的服务类型（一次只能提交一个服务类型）：
	// web/web-https：网页加速/网页加速-https
	// wsa/wsa-https：全站加速/全站加速-https
	// vodstream/vod-https：点播加速/点播加速-https
	// download/dl-https：下载加速/下载加速-https
	// livestream/live-https/cloudv-live：直播加速
	// v6sa/osv6：ipv6安全加速解决方案/IPv6一体化解决方案
	// 注意：
	// 1、service-type中的https不代表立即开启https，比如web-https中的https并不代表立刻支持https访问，需上传完证书后才可以支持https，切记！"}
	ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty"`
	// {"en":"The acceleration area of the acceleration domain, if the resource coverage needs to be limited according to the area, the acceleration area needs to be specified.
	//     When no acceleration area is specified, we will provide acceleration services with optimal resource coverage according to the service area opened by the customer.
	//     Multiple regions are separated by semicolons, and the supported regions are as follows: cn (Mainland China), am (Americas),
	//     emea (Europe, Middle East, Africa), apac (Asia-Pacific region).",
	// "zh_CN":"加速域名的加速区域，如果有需要根据区域限定资源覆盖时，才需要指定加速区域。未指定加速区域时，我们将按照客户开通的服务区域，以最优的资源覆盖提供加速服务。多个区域以分号分隔，支持配置的区域如下：cn（中国大陆）、am（美洲）、emea（欧洲、中东、非洲）、apac（亚太地区）"}
	ServiceAreas *string `json:"service-areas,omitempty" xml:"service-areas,omitempty"`
	// {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// {"en":"Configuration template, if you want to add the a domain using some specified configuration by default, you can specify the template id. For more detail, please contract the technical support.",
	//     "zh_CN":"配置单模板，特定的使用场景下，如果希望新增的加速域名参照某些指定配置时，可以指定配置单模板，具体使用请咨询对应的客户负责人。"}
	ConfigFormId *int `json:"config-form-id,omitempty" xml:"config-form-id,omitempty"`
	// {"en":"Refer to the configuration of the specified domain.
	// Note:
	// 1. If the referenced domain uses a certificate, the new domain should be in the 'DNS name' of the certificate.
	// 2. If the referenced domain has no China ICP, while the new domain name has, it may affect the cover resources and service quality.
	// 3. If the referenced domain has China ICP, while the new domain name doesn't, then the cover resources may be re-selected if it does not meet the policy requirements.
	// 4. It is not allowed to reference a domain which is traffic-free.", "zh_CN":"参照指定域名的配置，来创建加速域名。
	// 注意：
	// 1.参照域名如果有使用证书，新增域名也要在对应证书授权范围内。
	// 2.参照未备案域名，新增的域名如果已备案，可能影响资源使用和服务质量。
	// 3.参照备案域名，新增的域名如果未备案，若资源不满足政策要求，可能重选。
	// 4.不允许参照免流域名创建新域名。"}
	ReferencedDomainName *string `json:"referenced-domain-name,omitempty" xml:"referenced-domain-name,omitempty"`
	// {"en":"If you need to share a CNAME between domains, you can use this parameter. This parameter is a unique label for a public CNAME. Domains with the same cname-label will have the same CNAME.
	// Note:
	// 1. Domains with the same cname-label have the same coverage.
	// 2. Constraints of sharing a CNAME: consistent service-type, consistent certificate-id (if there is a certificate), consistent service-areas
	// 3. Multiple http domains can share a CNAME, multiple sni https domains can share a CNAME too.
	// 4. When a cname-label is used by a single domain, then the domain can be canceled acceleration. While a cname-label using by more then one domains, they can not be canceled acceleration.
	// 5. Support the purpose of modifying cname by modifying cname-label. )
	// ", "zh_CN":"共用一级标签，若有多个加速域名需要共用一级域名，则可以使用该参数。即拥有相同cname-label的一组域名，共用一级cname。
	// 注意：
	// 1、拥有相同cname-label的域名共用一级cname，且有完全一致的dns覆盖
	// 2、共用一级的约束：加速类型一致(service-type)、证书id一致（certificate-id,如果有证书）、加速区域一致(service-areas)
	// 3、多个http域名可共用一级，多个sni https域名可共用一级
	// 4、单个域名使用cname-label时，域名可cancel；多个域名共用一级时，不允许cancel这些域名
	// 5、支持通过修改cname-label达到修改cname的目的。）
	// "}
	CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
	// {"en":"The first level of cname prefix, true, indicates that the domain cname is used as the cname prefix, otherwise the 14-bit random string (number + letter) is used as the cname prefix.
	// Note: When the prefix is a generic domain name, a wsall is added as a prefix. Such as... Baidu.com.wscloudcdn.com, which will generate wsall.Baidu.com.wscloudcdn.com", "zh_CN":"一级cname前缀，true表示使用域名名称作为cname前缀，否则，使用14位随机串（数字+字母）作为cname前缀。
	//   注意：当前缀是泛域名时，则再增加wsall作为前缀。如.baidu.com.wscloudcdn.com，会生成wsall.baidu.com.wscloudcdn.com"}
	CnameWithCustomizedPrefix *string `json:"cname-with-customized-prefix,omitempty" xml:"cname-with-customized-prefix,omitempty"`
	// {"en":"Back-to-origin policy setting, which is used to set the origin site information and the back-to-origin policy of the none-live accelerated domain", "zh_CN":"回源策略设置(非直播域名使用)，用于设置加速域名的源站信息和回源策略。"}
	OriginConfig *CreateDomainRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" type:"Struct"`
	// {"en":"Live domain configuration, used to set the push flow of rtmp live acceleration domain (use required)
	// Note: In addition to the API call permission, you need to contact the dedicated customer service to apply for the corresponding API client template.", "zh_CN":"直播域名配置，用于设置rtmp直播加速域名的推拉流（使用需申请）
	// 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
	LiveConfig *CreateDomainRequestLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
	// {"en":"Identifies whether a domain name is fully overseas accelerated.
	// Whether the default is false
	// True: indicates that the client domain name is a pure overseas acceleration
	// False: Indicates that the client domain name has accelerated in China", "zh_CN":"标识域名是否是纯海外加速的。
	// 默认是否（false）
	// true ：表示客户域名纯海外加速
	// false：表示客户域名有在中国加速"}
	AccelerateNoChina *string `json:"accelerate-no-china,omitempty" xml:"accelerate-no-china,omitempty"`
	// {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip, X-Forwarded-For and ori_X-Forwarded-For.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip、X-Forwarded-For、ori_X-Forwarded-For
	// 1） Cdn-Src-Ip： 回源头部名称为Cdn-Src-Ip，获取与节点进行建联的IP作为客户端IP传递回源。
	// 2） X-Forwarded-For： 回源头部名称为X-Forwarded-For，携带的客户端IP值是Cdn-Src-Ip获取到的建联IP。
	// 3） ori_X-Forwarded-For：客户端请求CDN节点时会自带X-Forwarded-For，则CDN透传此头部和值回源。"}
	HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
	// {"en":"The live streaming domain which is pull domian ,and  directly returned to the source to verify the configuration.
	// which can be an IP or a domain name.
	// Can be IP or domain name. Ip and domain names can only be one. Multiple input parameters are not supported.", "zh_CN":"直播拉流域名，直接回源校验配置。
	// 可以是IP或域名。ip和域名只能一种。不支持多个入参。"}
	UpstreamHost *string `json:"upstream-host,omitempty" xml:"upstream-host,omitempty"`
	// {"en":"Set the publishing point of the live push-pull domain name
	// note:
	// 1. Pull flow and corresponding push flow domain name must be configured with the same publishing point.
	// 2. do not want to modify the publishing point, do not pass the node and the following parameters
	// 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
	// 注意：
	// 1、拉流和对应的推流域名，必须配置相同的发布点；
	// 2、不想修改发布点时，不要传入该节点及以下入参；
	// 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
	PublishPoints []*CreateDomainRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
	// {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your  certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate]", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】"}
	Ssl *CreateDomainRequestSsl `json:"ssl,omitempty" xml:"ssl,omitempty" type:"Struct"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetVersion(v string) *CreateDomainRequest {
	s.Version = &v
	return s
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDomainRequest) SetServiceType(v string) *CreateDomainRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateDomainRequest) SetServiceAreas(v string) *CreateDomainRequest {
	s.ServiceAreas = &v
	return s
}

func (s *CreateDomainRequest) SetComment(v string) *CreateDomainRequest {
	s.Comment = &v
	return s
}

func (s *CreateDomainRequest) SetConfigFormId(v int) *CreateDomainRequest {
	s.ConfigFormId = &v
	return s
}

func (s *CreateDomainRequest) SetReferencedDomainName(v string) *CreateDomainRequest {
	s.ReferencedDomainName = &v
	return s
}

func (s *CreateDomainRequest) SetCnameLabel(v string) *CreateDomainRequest {
	s.CnameLabel = &v
	return s
}

func (s *CreateDomainRequest) SetCnameWithCustomizedPrefix(v string) *CreateDomainRequest {
	s.CnameWithCustomizedPrefix = &v
	return s
}

func (s *CreateDomainRequest) SetOriginConfig(v *CreateDomainRequestOriginConfig) *CreateDomainRequest {
	s.OriginConfig = v
	return s
}

func (s *CreateDomainRequest) SetLiveConfig(v *CreateDomainRequestLiveConfig) *CreateDomainRequest {
	s.LiveConfig = v
	return s
}

func (s *CreateDomainRequest) SetAccelerateNoChina(v string) *CreateDomainRequest {
	s.AccelerateNoChina = &v
	return s
}

func (s *CreateDomainRequest) SetHeaderOfClientip(v string) *CreateDomainRequest {
	s.HeaderOfClientip = &v
	return s
}

func (s *CreateDomainRequest) SetUpstreamHost(v string) *CreateDomainRequest {
	s.UpstreamHost = &v
	return s
}

func (s *CreateDomainRequest) SetPublishPoints(v []*CreateDomainRequestPublishPoints) *CreateDomainRequest {
	s.PublishPoints = v
	return s
}

func (s *CreateDomainRequest) SetSsl(v *CreateDomainRequestSsl) *CreateDomainRequest {
	s.Ssl = v
	return s
}

type CreateDomainRequestOriginConfig struct {
	// {"en":"Origin address, which can be an IP or domain name.
	// 1. Multiple IPs are supported, separated by semicolons.
	// 2. Only one domain name is allowed. IP and domain name cannot exist at the same time.
	// 3. The length cannot exceed 500 characters.
	// 4. The number of IPs cannot exceed 15.", "zh_CN":"回源地址，可以是IP或域名。
	// 1、IP以分号分隔，支持多个。
	// 2、域名只能输入一个。IP与域名不能同时输入。
	// 3、限制最大不能超过500个字符长度。
	// 4、源IP个数不能超过15个。"}
	OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
	// {"en":"The Origin HOST for changing the HOST field in the return source HTTP request header. The supported domain name formats, each segement separated by a dot, does not exceed 62 characters, the total length should not exceed 128 characters.
	// .", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: 域名，每段（点号分隔）长度小于等于62，域名总长度小于等于128。"}
	DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
}

func (s CreateDomainRequestOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestOriginConfig) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestOriginConfig) SetOriginIps(v string) *CreateDomainRequestOriginConfig {
	s.OriginIps = &v
	return s
}

func (s *CreateDomainRequestOriginConfig) SetDefaultOriginHostHeader(v string) *CreateDomainRequestOriginConfig {
	s.DefaultOriginHostHeader = &v
	return s
}

type CreateDomainRequestLiveConfig struct {
	// {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
	StreamType *string `json:"stream-type,omitempty" xml:"stream-type,omitempty"`
	// {"en":"The push-pull domain name is used to set the push-flow domain name corresponding to the rtmp live streaming domain name. When the stream-type is pull, at least one of the source IP address and the corresponding push-stream domain name is not empty. When the stream-type is push, Incoming.", "zh_CN":"配套推流域名，用于设置rtmp直播拉流域名对应的推流域名，当stream-type为pull时，源站IP和配套推流域名至少一个不为空；当stream-type为push时，无需传入。"}
	OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty"`
	// {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.
	// 1. If it is a push-pull flow package, fill in 127.0.0.1, and the system will also default to 127.0.0.1.
	// 2. If it is directly returning to the source, fill in the source IP of the source pull stream.
	// ", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
	// 1、如果是推拉流配套，则填写127.0.0.1，不传系统也默认为127.0.0.1
	// 2、如果是直接回源拉流，则填写回源拉流的源站IP"}
	LiveConfigOriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
}

func (s CreateDomainRequestLiveConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestLiveConfig) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestLiveConfig) SetStreamType(v string) *CreateDomainRequestLiveConfig {
	s.StreamType = &v
	return s
}

func (s *CreateDomainRequestLiveConfig) SetOriginPushHost(v string) *CreateDomainRequestLiveConfig {
	s.OriginPushHost = &v
	return s
}

func (s *CreateDomainRequestLiveConfig) SetLiveConfigOriginIps(v string) *CreateDomainRequestLiveConfig {
	s.LiveConfigOriginIps = &v
	return s
}

type CreateDomainRequestPublishPoints struct {
	// {"en":"Livestream domain settings. Publish point, support multiple, do not pass the system by default to generate a publishing point uri for [/]", "zh_CN":"发布点，支持多个，不传系统默认生成一条发布点uri为“/”"}
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s CreateDomainRequestPublishPoints) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestPublishPoints) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestPublishPoints) SetUri(v string) *CreateDomainRequestPublishPoints {
	s.Uri = &v
	return s
}

type CreateDomainRequestSsl struct {
	// {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
	UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty"`
	// {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用合用证书（暂不支持）"}
	UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty"`
	// {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID，use-ssl为true时，才能传ssl-certificate-id。"}
	SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty"`
}

func (s CreateDomainRequestSsl) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestSsl) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestSsl) SetUseSsl(v string) *CreateDomainRequestSsl {
	s.UseSsl = &v
	return s
}

func (s *CreateDomainRequestSsl) SetUseForSni(v string) *CreateDomainRequestSsl {
	s.UseForSni = &v
	return s
}

func (s *CreateDomainRequestSsl) SetSslCertificateId(v int) *CreateDomainRequestSsl {
	s.SslCertificateId = &v
	return s
}

type CreateDomainResponse struct {
	// {"en":"If httpstatus=202, the interface is successfully invoked. And the current deployment of the new domain can be viewed using x-cnc-request-id in the header", "zh_CN":"httpstatus=202;   表示成功调用新增域名接口，可使用header中的x-cnc-request-id查看当前新增域名的部署情况"}
	HttpStatus *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
	// {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
	XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
	// {"en":"A URL used to access the domain name information, where the domain-id is a unique identifier generated by the network cloud platform for the domain name, and the value is a string.", "zh_CN":"用于访问该域名信息的URL，其中domain-id为我司云平台为该域名生成的唯一标示，其值为字符串。"}
	Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
	// {"en":"The name of the service domain automatically generated by the My company, for example: xxxx.cdn30.com", "zh_CN":"由我司自动生成的服务域名名称，例如：xxxx.cdn30.com"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
	// {"en":"The error code, when HTTPStatus is not 202, indicates the type of error the current request is calling.", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response information, when success is successful", "zh_CN":"响应信息，成功时为success"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHttpStatus(v int) *CreateDomainResponse {
	s.HttpStatus = &v
	return s
}

func (s *CreateDomainResponse) SetXCncRequestId(v string) *CreateDomainResponse {
	s.XCncRequestId = &v
	return s
}

func (s *CreateDomainResponse) SetLocation(v string) *CreateDomainResponse {
	s.Location = &v
	return s
}

func (s *CreateDomainResponse) SetCname(v string) *CreateDomainResponse {
	s.Cname = &v
	return s
}

func (s *CreateDomainResponse) SetCode(v string) *CreateDomainResponse {
	s.Code = &v
	return s
}

func (s *CreateDomainResponse) SetMessage(v string) *CreateDomainResponse {
	s.Message = &v
	return s
}

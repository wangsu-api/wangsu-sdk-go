// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type EditDomainConfigRequest struct {
	// {"en":"Version, the current version is 1.0.0", "zh_CN":"版本号，当前版本号1.0.0"}
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// {"en":"Remarks, up to 1000 characters.", "zh_CN":"备注信息，最大限制1000个字符"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// {"en":"The acceleration area of the acceleration domain, if the resource coverage needs to be limited according to the area, the acceleration area needs to be specified. When no acceleration area is specified, we will provide acceleration services with optimal resource coverage according to the service area opened by the customer. Multiple regions are separated by semicolons, and the supported regions are as follows: cn (Mainland China), am (Americas), emea (Europe, Middle East, Africa), apac (Asia-Pacific region).",
	//     "zh_CN":"加速域名的加速区域，如果有需要根据区域限定资源覆盖时，才需要指定加速区域。未指定加速区域时，我们将按照客户开通的服务区域，以最优的资源覆盖提供加速服务。多个区域以分号分隔，支持配置的区域如下：cn（中国大陆）、am（美洲）、emea（欧洲、中东、非洲）、apac（亚太地区）"}
	ServiceAreas *string `json:"service-areas,omitempty" xml:"service-areas,omitempty"`
	// {"en":"If you need to share a CNAME between domains, you can use this parameter. This parameter is a unique label for a public CNAME. Domains with the same cname-label will have the same CNAME.
	// Note:
	// 1. Domains with the same cname-label have the same coverage.
	// 2. Constraints of sharing a CNAME: consistent service-type, consistent certificate-id (if there is a certificate), consistent service-areas
	// 3. Multiple http domains can share a CNAME, multiple sni https domains can share a CNAME too.
	// 4. When a cname-label is used by a single domain, then the domain can be canceled acceleration. While a cname-label using by more then one domains, they can not be canceled acceleration.
	// 5. Support the purpose of modifying cname by modifying cname-label. )",
	// "zh_CN":"共用一级标签，若有多个加速域名需要共用一级域名，则可以使用该参数。即拥有相同cname-label的一组域名，共用一级cname。
	// 注意：
	// 1、拥有相同cname-label的域名共用一级cname，且有完全一致的dns覆盖
	// 2、共用一级的约束：加速类型一致(service-type)、证书id一致（certificate-id,如果有证书）、加速区域一致(service-areas)
	// 3、多个http域名可共用一级，多个sni https域名可共用一级
	// 4、单个域名使用cname-label时，域名可cancel；多个域名共用一级时，不允许cancel这些域名
	// 5、支持通过修改cname-label达到修改cname的目的。"}
	CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty"`
	// {"en":"Back to origin policy settings for setting source site information and return source policies for accelerated domain names", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
	OriginConfig *EditDomainConfigRequestOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" type:"Struct"`
	// {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your  certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate]", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】"}
	Ssl *EditDomainConfigRequestSsl `json:"ssl,omitempty" xml:"ssl,omitempty" type:"Struct"`
	// {"en":"Cache file HOST.
	// Cache rules for caching HOST domain names and accelerated domain names must be consistent.", "zh_CN":"缓存文件HOST。缓存HOST域名和加速域名的缓存规则必须一致。"}
	CacheHost *string `json:"cache-host,omitempty" xml:"cache-host,omitempty"`
	// {"en":"Enable httpdns settings.
	// The optional values are true and false, true means enabled; false means off. This function is not enabled by default. If you need it, please contact the technical support to apply for this feature.", "zh_CN":"启用httpdns设置（使用需申请）
	// 可选值为true和false，true表示启用；false表示关闭
	// 注意：该功能默认不开启，若您有需要，请联系专属客服申请开通。"}
	EnableHttpdns *string `json:"enable-httpdns,omitempty" xml:"enable-httpdns,omitempty"`
	// {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
	HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty"`
	// {"en":"Live domain name configuration, used to set the push flow of live acceleration domain name", "zh_CN":"直播域名配置，用于设置直播加速域名的推拉流"}
	LiveConfig *EditDomainConfigRequestLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" type:"Struct"`
	// {"en":"Set the publishing point of the live push-pull domain.
	// Note:
	// 1. The pull stream and the corresponding push stream domain must be configured with the same publishing point.
	// 2. If you are not going to modify the publishing point, please do not pass this param.
	// 3. The publishing point adopts the overlay update. Each time you modify, you need to submit all the publishing points. You cannot submit only the parts that need to be modified.", "zh_CN":"设置直播推拉流域名的发布点
	// 注意：
	// 1、拉流和对应的推流域名，必须配置相同的发布点；
	// 2、不想修改发布点时，不要传入该节点及以下入参；
	// 3、发布点采用覆盖式更新，每次修改时，需要提交全部发布点，不能仅提交需要修改的部分。"}
	PublishPoints []*EditDomainConfigRequestPublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" type:"Repeated"`
}

func (s EditDomainConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s EditDomainConfigRequest) GoString() string {
	return s.String()
}

func (s *EditDomainConfigRequest) SetVersion(v string) *EditDomainConfigRequest {
	s.Version = &v
	return s
}

func (s *EditDomainConfigRequest) SetComment(v string) *EditDomainConfigRequest {
	s.Comment = &v
	return s
}

func (s *EditDomainConfigRequest) SetServiceAreas(v string) *EditDomainConfigRequest {
	s.ServiceAreas = &v
	return s
}

func (s *EditDomainConfigRequest) SetCnameLabel(v string) *EditDomainConfigRequest {
	s.CnameLabel = &v
	return s
}

func (s *EditDomainConfigRequest) SetOriginConfig(v *EditDomainConfigRequestOriginConfig) *EditDomainConfigRequest {
	s.OriginConfig = v
	return s
}

func (s *EditDomainConfigRequest) SetSsl(v *EditDomainConfigRequestSsl) *EditDomainConfigRequest {
	s.Ssl = v
	return s
}

func (s *EditDomainConfigRequest) SetCacheHost(v string) *EditDomainConfigRequest {
	s.CacheHost = &v
	return s
}

func (s *EditDomainConfigRequest) SetEnableHttpdns(v string) *EditDomainConfigRequest {
	s.EnableHttpdns = &v
	return s
}

func (s *EditDomainConfigRequest) SetHeaderOfClientip(v string) *EditDomainConfigRequest {
	s.HeaderOfClientip = &v
	return s
}

func (s *EditDomainConfigRequest) SetLiveConfig(v *EditDomainConfigRequestLiveConfig) *EditDomainConfigRequest {
	s.LiveConfig = v
	return s
}

func (s *EditDomainConfigRequest) SetPublishPoints(v []*EditDomainConfigRequestPublishPoints) *EditDomainConfigRequest {
	s.PublishPoints = v
	return s
}

type EditDomainConfigRequestOriginConfig struct {
	// {"en":"Origin address, which can be an IP or domain name.
	// 1. Multiple IPs are supported, separated by semicolons.
	// 2. Only one domain name is allowed. IP and domain name cannot exist at the same time.
	// 3. The length cannot exceed 500 characters.
	// 4. The number of IPs cannot exceed 15.
	// ", "zh_CN":"回源地址，可以是IP或域名。
	// 1、IP以分号分隔，支持多个。
	// 2、域名只能输入一个。IP与域名不能同时输入。
	// 3、限制最大不能超过500个字符长度。
	// 4、IP最多15个。"}
	OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
	// {"en":"Back-to-origin HOST, used to change the HOST field in the back-to-origin HTTP request header. The supported formats are: ① domain name ③ ip
	// Note:
	// 1. Must comply with the ip/domain name format specification. If it is a domain name, the length of the domain name must be less than or equal to 128 characters.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: ①域名；②ip；
	// 注意：
	// 1、必须符合ip/域名格式规范。如果是域名，则域名长度小于等于128。"}
	DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty"`
}

func (s EditDomainConfigRequestOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s EditDomainConfigRequestOriginConfig) GoString() string {
	return s.String()
}

func (s *EditDomainConfigRequestOriginConfig) SetOriginIps(v string) *EditDomainConfigRequestOriginConfig {
	s.OriginIps = &v
	return s
}

func (s *EditDomainConfigRequestOriginConfig) SetDefaultOriginHostHeader(v string) *EditDomainConfigRequestOriginConfig {
	s.DefaultOriginHostHeader = &v
	return s
}

type EditDomainConfigRequestSsl struct {
	// {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
	UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty"`
	// {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用合用证书（暂不支持）"}
	UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty"`
	// {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID
	// use-ssl为true时，才能传ssl-certificate-id。"}
	SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty"`
}

func (s EditDomainConfigRequestSsl) String() string {
	return tea.Prettify(s)
}

func (s EditDomainConfigRequestSsl) GoString() string {
	return s.String()
}

func (s *EditDomainConfigRequestSsl) SetUseSsl(v string) *EditDomainConfigRequestSsl {
	s.UseSsl = &v
	return s
}

func (s *EditDomainConfigRequestSsl) SetUseForSni(v string) *EditDomainConfigRequestSsl {
	s.UseForSni = &v
	return s
}

func (s *EditDomainConfigRequestSsl) SetSslCertificateId(v int) *EditDomainConfigRequestSsl {
	s.SslCertificateId = &v
	return s
}

type EditDomainConfigRequestLiveConfig struct {
	// {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.
	// 1. If it is a push-pull flow package, fill in 127.0.0.1, and the system will also default to 127.0.0.1.
	// 2. If it is directly returning to the source, fill in the source IP of the source pull stream.
	// ", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
	// 1、如果是推拉流配套，则填写127.0.0.1，不传系统也默认为127.0.0.1
	// 2、如果是直接回源拉流，则填写回源拉流的源站IP"}
	LiveConfigOriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty"`
	// {"en":"A matching push domain name is used to set up a current domain name corresponding to the live streaming domain name. When the stream-type is pull, the source station IP and the supporting current domain name are at least one empty; when stream-type is push, it does not need to be introduced.", "zh_CN":"配套推流域名，用于设置直播拉流域名对应的推流域名，当stream-type为pull时，源站IP和配套推流域名至少一个不为空；当stream-type为push时，无需传入。"}
	OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty"`
}

func (s EditDomainConfigRequestLiveConfig) String() string {
	return tea.Prettify(s)
}

func (s EditDomainConfigRequestLiveConfig) GoString() string {
	return s.String()
}

func (s *EditDomainConfigRequestLiveConfig) SetLiveConfigOriginIps(v string) *EditDomainConfigRequestLiveConfig {
	s.LiveConfigOriginIps = &v
	return s
}

func (s *EditDomainConfigRequestLiveConfig) SetOriginPushHost(v string) *EditDomainConfigRequestLiveConfig {
	s.OriginPushHost = &v
	return s
}

type EditDomainConfigRequestPublishPoints struct {
	// {"en":"Livestream domain settings. Publish point, support multiple, do not pass the system by default to generate a publishing point uri for [/]", "zh_CN":"发布点，支持多个，不传系统默认生成一条发布点uri为“/”"}
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s EditDomainConfigRequestPublishPoints) String() string {
	return tea.Prettify(s)
}

func (s EditDomainConfigRequestPublishPoints) GoString() string {
	return s.String()
}

func (s *EditDomainConfigRequestPublishPoints) SetUri(v string) *EditDomainConfigRequestPublishPoints {
	s.Uri = &v
	return s
}

type EditDomainConfigResponse struct {
	// {"en":"If httpstatus=202, the interface is successfully invoked. You can use the x-cnc-request-id in the header to check the deployment of domain.", "zh_CN":"httpstatus=202;   表示成功调用接口，可使用header中的x-cnc-request-id查看域名的部署情况"}
	HttpStatusCode *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
	// {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
	XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
	// {"en":"The corresponding deployment version number of this modification", "zh_CN":"本次修改对应的部署版本号"}
	XCncDeployVersion *string `json:"x-cnc-deploy-version,omitempty" xml:"x-cnc-deploy-version,omitempty" require:"true"`
	// {"en":"Error code, which appears when HTTPStatus is not 202, represents the error type of the current request call", "zh_CN":"错误代码，当HTTPStatus不为202时出现，表示当前请求调用的错误类型"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response information, success when successful", "zh_CN":"响应信息，成功时为success"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s EditDomainConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s EditDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *EditDomainConfigResponse) SetHttpStatusCode(v int) *EditDomainConfigResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *EditDomainConfigResponse) SetXCncRequestId(v string) *EditDomainConfigResponse {
	s.XCncRequestId = &v
	return s
}

func (s *EditDomainConfigResponse) SetXCncDeployVersion(v string) *EditDomainConfigResponse {
	s.XCncDeployVersion = &v
	return s
}

func (s *EditDomainConfigResponse) SetCode(v string) *EditDomainConfigResponse {
	s.Code = &v
	return s
}

func (s *EditDomainConfigResponse) SetMessage(v string) *EditDomainConfigResponse {
	s.Message = &v
	return s
}

// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type GetBasicConfigurationOfDomainRequest struct {
}

func (s GetBasicConfigurationOfDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainRequest) GoString() string {
	return s.String()
}

//write a function to convert a json reponse to GetBasicConfigurationOfDomainResponse accodrding the  key in the JSON response

type GetBasicConfigurationOfDomainResponse struct {
	// {"en":"Acceleration domain ID returned by the system.", "zh_CN":"系统返回的加速域名ID"}
	DomainId *string `json:"domain-id,omitempty" xml:"domain-id,omitempty" require:"true"`
	// {"en":"Accelerated domain name.", "zh_CN":"加速域名的名称"}
	DomainName *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
	// {"en":"Domain name creation time,
	// Format: week, dd month yyyy hh:mm:ss GMT +8:00", "zh_CN":"域名创建时间，格式: week, dd month yyyy hh:mm:ss GMT +8:00
	// 如：Mon, 18 Feb 2019 02:54:19 GMT +8:00"}
	CreatedDate *string `json:"created-date,omitempty" xml:"created-date,omitempty" require:"true"`
	// {"en":"Domain name last modified time,
	// Format: week, dd month yyyy hh:mm:ss GMT +8:00
	// Example: Mon, 18 Feb 2019 02:54:19 GMT +8 p.m", "zh_CN":"域名最近修改时间，格式: week, dd month yyyy hh:mm:ss GMT +8:00
	// 如：Mon, 18 Feb 2019 02:54:19 GMT +8:00"}
	LastModified *string `json:"last-modified,omitempty" xml:"last-modified,omitempty" require:"true"`
	// {"en":"Speed up domain name service types, including the following:
	// Web/web-https: web acceleration / web acceleration - https
	// Wsa/wsa-https: Total Station Acceleration / Total Station Acceleration - https
	// Vodstream/vod-https: on-demand acceleration/on-demand acceleration-https
	// Download/dl-https: Download acceleration/download acceleration - https", "zh_CN":"加速域名的服务类型，包括如下：
	// web/web-https：网页加速/网页加速-https
	// wsa/wsa-https：全站加速/全站加速-https
	// vodstream/vod-https：点播加速/点播加速-https
	// download/dl-https：下载加速/下载加速-https"}
	ServiceType *string `json:"service-type,omitempty" xml:"service-type,omitempty" require:"true"`
	// {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
	// {"en":"Service areas of the domain name", "zh_CN":"域名的加速区域."}
	ServiceAreas *string `json:"service-areas,omitempty" xml:"service-areas,omitempty" require:"true"`
	// {"en":"CNAME of the domain you queried, for example: 7nt6mrh7sdkslj.cdn30.com.",
	//     "zh_CN":"加速域名对应的CNAME域名，例如：7nt6mrh7sdkslj.cdn30.com。"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
	// {"en":"The deployment status of the domain name. Deployed indicates that the domain configuration is distributed successfully. InProgress indicates that the deployment task of the domain configuration is still in progress, and may be in queue, or failed.",
	//     "zh_CN":"加速域名的部署状态，Deployed表示该加速域名配置完成部署；InProgress表示该加速域名配置的部署任务还在进行中，可能处于排队、部署中或失败任意一种状态。"}
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// {"en":"Accelerate status of the domain in our CDN, true means the CDN acceleration is normal; false means all request will back to origin directly in DNS.", "zh_CN":"加速域名的CDN服务状态，true表示启用CDN加速服务；false表示取消CDN加速服务。"}
	CdnServiceStatus *string `json:"cdn-service-status,omitempty" xml:"cdn-service-status,omitempty" require:"true"`
	// {"en":"Activation of the domain. It is false when the domain service is disabled, and true when the domain service is enabled.","zh_CN":"加速域名的启用状态，当禁用加速域名服务后，此项为false；当启用加速域名服务后，此项为true"}
	Enabled *string `json:"enabled,omitempty" xml:"enabled,omitempty" require:"true"`
	// {"en":"Shared first level alias", "zh_CN":"共用一级别名"}
	CnameLabel *string `json:"cname-label,omitempty" xml:"cname-label,omitempty" require:"true"`
	// {"en":"Return source policy setting, used to set the source station information and return source policy of the accelerated domain name", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
	OriginConfig *GetBasicConfigurationOfDomainResponseOriginConfig `json:"origin-config,omitempty" xml:"origin-config,omitempty" require:"true" type:"Struct"`
	// {"en":"Ssl certificate settings, used to set the ssl certificate configuration for the accelerated domain name", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置"}
	Ssl *GetBasicConfigurationOfDomainResponseSsl `json:"ssl,omitempty" xml:"ssl,omitempty" require:"true" type:"Struct"`
	// {"en":"Cache rule settings for setting cache rules for accelerated domain names", "zh_CN":"查询缓存时间配置，请使用新接口：【查询缓存时间配置】接口"}
	CacheBehaviors map[string]interface{} `json:"cache-behaviors,omitempty" xml:"cache-behaviors,omitempty" require:"true"`
	// {"en":"Cache file HOST (not return by default, application is required to use)", "zh_CN":"缓存文件HOST（默认不返回，使用需申请）
	// 缓存HOST域名和加速域名的”缓存规则”必须一致
	// 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
	CacheHost *string `json:"cache-host,omitempty" xml:"cache-host,omitempty" require:"true"`
	// {"en":"Enable httpdns settings (not return by default, application is required to use)", "zh_CN":"启用httpdns设置（默认不返回，使用需申请）
	// 可选值为true和false，true表示启用；false表示关闭
	// 注意：该节点下的相关参数配置，除开通API调用权限外，还需要联系专属客服申请开通对应的API客户模板"}
	EnableHttpdns *string `json:"enable-httpdns,omitempty" xml:"enable-httpdns,omitempty" require:"true"`
	// {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
	HeaderOfClientip *string `json:"header-of-clientip,omitempty" xml:"header-of-clientip,omitempty" require:"true"`
	// {"en":"The live push-pull stream type, the optional values are pull and push, pull means pull flow; push means push flow.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
	StreamMode *string `json:"stream-mode,omitempty" xml:"stream-mode,omitempty" require:"true"`
	// {"en":"Live domain name configuration, rtmp live acceleration domain name push-pull flow", "zh_CN":"直播域名配置，rtmp直播加速域名的推拉流"}
	LiveConfig *GetBasicConfigurationOfDomainResponseLiveConfig `json:"live-config,omitempty" xml:"live-config,omitempty" require:"true" type:"Struct"`
	// {"en":"The launch point of the live push-pull domain name", "zh_CN":"直播推拉流域名的发布点
	// 注意：拉流和对应的推流域名，发布点是相同的"}
	PublishPoints *GetBasicConfigurationOfDomainResponsePublishPoints `json:"publish-points,omitempty" xml:"publish-points,omitempty" require:"true" type:"Struct"`
}

func (s GetBasicConfigurationOfDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponse) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponse) SetDomainId(v string) *GetBasicConfigurationOfDomainResponse {
	s.DomainId = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetDomainName(v string) *GetBasicConfigurationOfDomainResponse {
	s.DomainName = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCreatedDate(v string) *GetBasicConfigurationOfDomainResponse {
	s.CreatedDate = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetLastModified(v string) *GetBasicConfigurationOfDomainResponse {
	s.LastModified = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetServiceType(v string) *GetBasicConfigurationOfDomainResponse {
	s.ServiceType = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetComment(v string) *GetBasicConfigurationOfDomainResponse {
	s.Comment = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetServiceAreas(v string) *GetBasicConfigurationOfDomainResponse {
	s.ServiceAreas = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCname(v string) *GetBasicConfigurationOfDomainResponse {
	s.Cname = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetStatus(v string) *GetBasicConfigurationOfDomainResponse {
	s.Status = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCdnServiceStatus(v string) *GetBasicConfigurationOfDomainResponse {
	s.CdnServiceStatus = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetEnabled(v string) *GetBasicConfigurationOfDomainResponse {
	s.Enabled = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCnameLabel(v string) *GetBasicConfigurationOfDomainResponse {
	s.CnameLabel = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetOriginConfig(v *GetBasicConfigurationOfDomainResponseOriginConfig) *GetBasicConfigurationOfDomainResponse {
	s.OriginConfig = v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetSsl(v *GetBasicConfigurationOfDomainResponseSsl) *GetBasicConfigurationOfDomainResponse {
	s.Ssl = v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCacheBehaviors(v map[string]interface{}) *GetBasicConfigurationOfDomainResponse {
	s.CacheBehaviors = v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetCacheHost(v string) *GetBasicConfigurationOfDomainResponse {
	s.CacheHost = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetEnableHttpdns(v string) *GetBasicConfigurationOfDomainResponse {
	s.EnableHttpdns = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetHeaderOfClientip(v string) *GetBasicConfigurationOfDomainResponse {
	s.HeaderOfClientip = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetStreamMode(v string) *GetBasicConfigurationOfDomainResponse {
	s.StreamMode = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetLiveConfig(v *GetBasicConfigurationOfDomainResponseLiveConfig) *GetBasicConfigurationOfDomainResponse {
	s.LiveConfig = v
	return s
}

func (s *GetBasicConfigurationOfDomainResponse) SetPublishPoints(v *GetBasicConfigurationOfDomainResponsePublishPoints) *GetBasicConfigurationOfDomainResponse {
	s.PublishPoints = v
	return s
}

type GetBasicConfigurationOfDomainResponseOriginConfig struct {
	// {"en":"Return source address, which can be IP or domain name.", "zh_CN":"回源地址，可以是IP或域名。
	// 1、IP以分号分隔，支持多个。
	// 2、域名只能一个。
	// 3、限制最大不能超过500个字符长度。"}
	OriginIps *string `json:"originIps,omitempty" xml:"originIps,omitempty" require:"true"`
	// {"en":"Back to the source HOST, used to change the HOST field in the source HTTP request header.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。"}
	DefaultOriginHostHeader *string `json:"default-origin-host-header,omitempty" xml:"default-origin-host-header,omitempty" require:"true"`
	// {"en":"advance origin config", "zh_CN":"高级源配置"}
	AdvOriginConfigs *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs `json:"adv-origin-configs,omitempty" xml:"adv-origin-configs,omitempty" require:"true" type:"Struct"`
}

func (s GetBasicConfigurationOfDomainResponseOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseOriginConfig) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfig) SetOriginIps(v string) *GetBasicConfigurationOfDomainResponseOriginConfig {
	s.OriginIps = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfig) SetDefaultOriginHostHeader(v string) *GetBasicConfigurationOfDomainResponseOriginConfig {
	s.DefaultOriginHostHeader = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfig) SetAdvOriginConfigs(v *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) *GetBasicConfigurationOfDomainResponseOriginConfig {
	s.AdvOriginConfigs = v
	return s
}

type GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs struct {
	// {"en":"The advanced source monitors the url, and requests <master-ips> through the url. If the response is not 2**, 3** response, it is considered that the primary source ip is faulty, and <backup-ips> is used at this time.", "zh_CN":"高级源监控url，通过该url请求<master-ips>，如果返回非2**，3**响应时，认为主要回源ip故障，此时使用<backup-ips>。
	// 完整的url，例如：http://a.example.com/test.html"}
	DetectUrl *string `json:"detect-url,omitempty" xml:"detect-url,omitempty" require:"true"`
	// {"en":"Advanced source monitoring period, in seconds, optional as an integer greater than or equal to 0, 0 means no monitoring", "zh_CN":"高级源监控周期，单位秒，可选值为大于等于0的整数，0表示不监控"}
	DetectPeriod *int `json:"detect-period,omitempty" xml:"detect-period,omitempty" require:"true"`
	// {"en":"advance origin config", "zh_CN":"高级源配置"}
	AdvOriginConfig *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig `json:"adv-origin-config,omitempty" xml:"adv-origin-config,omitempty" require:"true" type:"Struct"`
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) SetDetectUrl(v string) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs {
	s.DetectUrl = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) SetDetectPeriod(v int) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs {
	s.DetectPeriod = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs) SetAdvOriginConfig(v *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigs {
	s.AdvOriginConfig = v
	return s
}

type GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig struct {
	// {"en":"The advanced source mainly returns the source IP. Multiple IPs are separated by a semicolon \";\", and the return source IP cannot be repeated.", "zh_CN":"高级源主要回源IP，多个IP用分号“;”分隔，回源IP不能重复"}
	MasterIps *string `json:"master-ips,omitempty" xml:"master-ips,omitempty" require:"true"`
	// {"en":"Advanced source backup source IP, multiple IPs are separated by semicolon \";\", and the return source IP cannot be duplicated.", "zh_CN":"高级源备用回源IP，多个IP用分号“;”分隔，回源IP不能重复"}
	BackupIps *string `json:"backup-ips,omitempty" xml:"backup-ips,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) SetMasterIps(v string) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig {
	s.MasterIps = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig) SetBackupIps(v string) *GetBasicConfigurationOfDomainResponseOriginConfigAdvOriginConfigsAdvOriginConfig {
	s.BackupIps = &v
	return s
}

type GetBasicConfigurationOfDomainResponseSsl struct {
	// {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
	UseSsl *string `json:"use-ssl,omitempty" xml:"use-ssl,omitempty" require:"true"`
	// {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use non-sni traditional certificate", "zh_CN":"使用sni证书，可选值为true和false，true表示使用sni证书，false表示使用非sni的传统证书"}
	UseForSni *string `json:"use-for-sni,omitempty" xml:"use-for-sni,omitempty" require:"true"`
	// {"en":"Certificate ID, the certificate ID returned by the system after the new certificate is successfully added.", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID"}
	SslCertificateId *int `json:"ssl-certificate-id,omitempty" xml:"ssl-certificate-id,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseSsl) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseSsl) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseSsl) SetUseSsl(v string) *GetBasicConfigurationOfDomainResponseSsl {
	s.UseSsl = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseSsl) SetUseForSni(v string) *GetBasicConfigurationOfDomainResponseSsl {
	s.UseForSni = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseSsl) SetSslCertificateId(v int) *GetBasicConfigurationOfDomainResponseSsl {
	s.SslCertificateId = &v
	return s
}

type GetBasicConfigurationOfDomainResponseLiveConfig struct {
	// {"en":"The live push-pull stream type, the optional values are pull and push.", "zh_CN":"直播推拉流类型，可选值为pull和push，pull表示拉流；   push表示推流。"}
	StreamType *string `json:"stream-type,omitempty" xml:"stream-type,omitempty" require:"true"`
	// {"en":"Source station IP. When the stream-type is pull, at least one of the source station IP and the companion push stream domain name is not empty.", "zh_CN":"源站IP，当stream-type为pull时，源站IP和配套推流域名至少一个不为空。
	// 1、如果是推拉流配套，则返回127.0.0.1
	// 2、如果是直接回源拉流，则返回源站IP"}
	OriginIps *string `json:"origin-ips,omitempty" xml:"origin-ips,omitempty" require:"true"`
	// {"en":"The corresponding push domain name, the rtmp live streaming domain name corresponding to the push domain name", "zh_CN":"配套推流域名，rtmp直播拉流域名对应的推流域名"}
	OriginPushHost *string `json:"origin-push-host,omitempty" xml:"origin-push-host,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponseLiveConfig) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponseLiveConfig) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponseLiveConfig) SetStreamType(v string) *GetBasicConfigurationOfDomainResponseLiveConfig {
	s.StreamType = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseLiveConfig) SetOriginIps(v string) *GetBasicConfigurationOfDomainResponseLiveConfig {
	s.OriginIps = &v
	return s
}

func (s *GetBasicConfigurationOfDomainResponseLiveConfig) SetOriginPushHost(v string) *GetBasicConfigurationOfDomainResponseLiveConfig {
	s.OriginPushHost = &v
	return s
}

type GetBasicConfigurationOfDomainResponsePublishPoints struct {
	// {"en":"Release point, support multiple, the system default is \"/\"", "zh_CN":"发布点，支持多个，系统默认值为“/”"}
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
}

func (s GetBasicConfigurationOfDomainResponsePublishPoints) String() string {
	return tea.Prettify(s)
}

func (s GetBasicConfigurationOfDomainResponsePublishPoints) GoString() string {
	return s.String()
}

func (s *GetBasicConfigurationOfDomainResponsePublishPoints) SetUri(v string) *GetBasicConfigurationOfDomainResponsePublishPoints {
	s.Uri = &v
	return s
}

type Paths struct {
	// {"en":"Accelerated domain name or domain ID", "zh_CN":"域名名称或域名id"}
	Domain *string `json:"domain-name,omitempty" xml:"domain-name,omitempty" require:"true"`
}

func (s *Paths) SetDomain(v string) *Paths {
	s.Domain = &v
	return s
}

type Parameters struct {
}

type ResponseHeader struct {
	// {"en":"Httpstatus=200 indicates that the interface is successfully invoked.", "zh_CN":"httpstatus=200;   表示成功调用接口"}
	HttpStatusCode *int `json:"http status code,omitempty" xml:"http status code,omitempty" require:"true"`
	// {"en":"Uniquely labeled id for querying each requested task (for all interfaces)", "zh_CN":"唯一标示的id，用于查询每次请求的任务 （适用全部接口）"}
	XCncRequestId *string `json:"x-cnc-request-id,omitempty" xml:"x-cnc-request-id,omitempty" require:"true"`
	// {"en":"The corresponding deployment version number of this modification", "zh_CN":"本次修改对应的部署版本号"}
	XCncDeployVersion *string `json:"x-cnc-deploy-version,omitempty" xml:"x-cnc-deploy-version,omitempty" require:"true"`
}

func (s ResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s ResponseHeader) GoString() string {
	return s.String()
}

func (s *ResponseHeader) SetHttpStatusCode(v int) *ResponseHeader {
	s.HttpStatusCode = &v
	return s
}

func (s *ResponseHeader) SetXCncRequestId(v string) *ResponseHeader {
	s.XCncRequestId = &v
	return s
}

func (s *ResponseHeader) SetXCncDeployVersion(v string) *ResponseHeader {
	s.XCncDeployVersion = &v
	return s
}

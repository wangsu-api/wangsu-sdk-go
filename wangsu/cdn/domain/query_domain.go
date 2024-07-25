// This file is auto-generated, don't edit it. Thanks.
package domain

import (
	"github.com/alibabacloud-go/tea/tea"
)

type QueryDomainForTerraformRequest struct {
}

func (s QueryDomainForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformRequest) GoString() string {
	return s.String()
}

type QueryDomainForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.", "zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.", "zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.", "zh_CN":"接口响应数据"}
	Data *QueryDomainForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDomainForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponse) SetCode(v string) *QueryDomainForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryDomainForTerraformResponse) SetMessage(v string) *QueryDomainForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryDomainForTerraformResponse) SetData(v *QueryDomainForTerraformResponseData) *QueryDomainForTerraformResponse {
	s.Data = v
	return s
}

type QueryDomainForTerraformResponseData struct {
	// {"en":"domain ID", "zh_CN":"域名ID"}
	DomainId *int64 `json:"domainId,omitempty" xml:"domainId,omitempty" require:"true"`
	// {"en":"cname", "zh_CN":"服务域名"}
	Cname *string `json:"cname,omitempty" xml:"cname,omitempty" require:"true"`
	// {"en":"Need to access the domain name of the CDN. a generic domain name is supported, starting with the symbol '.', such as.example.com, which also contains a multilevel 'a.b.example.com'.If example.com is filed, the domain name xx.example.com does not need to be filed.", "zh_CN":"需要接入CDN的域名。支持泛域名，以符号“.”开头，如：.example.com，泛域名也包含多级“a.b.example.com”。
	// 如果example.com已备案，那么域名xx.example.com则不需要备案。"}
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty" require:"true"`
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
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"The acceleration area of the acceleration domain, if the resource coverage needs to be limited according to the area, the acceleration area needs to be specified.
	//     When no acceleration area is specified, we will provide acceleration services with optimal resource coverage according to the service area opened by the customer.
	//     Multiple regions are separated by semicolons, and the supported regions are as follows: cn (Mainland China), am (Americas),
	//     emea (Europe, Middle East, Africa), apac (Asia-Pacific region).",
	// "zh_CN":"加速域名的加速区域，如果有需要根据区域限定资源覆盖时，才需要指定加速区域。未指定加速区域时，我们将按照客户开通的服务区域，以最优的资源覆盖提供加速服务。多个区域以分号分隔，支持配置的区域如下：cn（中国大陆）、am（美洲）、emea（欧洲、中东、非洲）、apac（亚太地区）"}
	ServiceAreas *string `json:"serviceAreas,omitempty" xml:"serviceAreas,omitempty" require:"true"`
	// {"en":"Remarks, up to 1000 characters", "zh_CN":"备注信息，最大限制1000个字符"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
	// {"en":"Pass the response header of client IP. The optional values are Cdn-Src-Ip and X-Forwarded-For. The default value is Cdn-Src-Ip.", "zh_CN":"传递客户端ip的响应头部，可选值为Cdn-Src-Ip和X-Forwarded-For，默认值为Cdn-Src-Ip"}
	HeaderOfClientIp *string `json:"headerOfClientIp,omitempty" xml:"headerOfClientIp,omitempty" require:"true"`
	// {"en":"Back to origin policy settings for setting source site information and return source policies for accelerated domain names", "zh_CN":"回源策略设置，用于设置加速域名的源站信息和回源策略"}
	OriginConfig *QueryDomainForTerraformResponseDataOriginConfig `json:"originConfig,omitempty" xml:"originConfig,omitempty" require:"true" type:"Struct"`
	// {"en":"SSL settings, to bind a certificate with the accelerated domain. You can use the interface [AddCertificate] to upload your  certificates. If you want to modify a certificate, please use the interface: [UpdateCertificate]", "zh_CN":"ssl证书设置，用于设置加速域名的ssl证书配置。上传证书请使用接口：【新增证书V2】；若要修改证书，请使用接口：【修改证书V2】"}
	Ssl *QueryDomainForTerraformResponseDataSsl `json:"ssl,omitempty" xml:"ssl,omitempty" require:"true" type:"Struct"`
	// {"en":"Cache time configuration
	// note:
	// 1. When you need to cancel the cache time configuration setting, you can pass in the empty node <cache-time-behaviors></cache-time-behaviors>.
	// 2. When it is required to set the cache time configuration, this item is required.", "zh_CN":"缓存时间配置
	// 注意：
	// 1. 需要取消缓存时间配置设置时，可以传入空节点<cache-time-behaviors></cache-time-behaviors>。
	// 2. 表示需要设置缓存时间配置时，此项必填"}
	CacheTimeBehaviors []*QueryDomainForTerraformResponseDataCacheTimeBehaviors `json:"cacheTimeBehaviors,omitempty" xml:"cacheTimeBehaviors,omitempty" type:"Repeated"`
	// {"en":"Custom Cachekey Configuration, parent node
	// 1. When you need to configure the cachekey rules, this must be filled in.
	// 2. Configuration of clearing for <cacheKeyRules/>.", "zh_CN":"配置自定义缓存key功能。
	// 1.需要设置自定义缓存key配置时，此项必填
	// 2.为<cacheKeyRules/>时清空自定义缓存key配置"}
	CacheKeyRules []*QueryDomainForTerraformResponseDataCacheKeyRules `json:"cacheKeyRules,omitempty" xml:"cacheKeyRules,omitempty" require:"true" type:"Repeated"`
	// {"en":"Query String Settings Configuration, parent node
	// 1. When you need to configure the query string, this must be filled in.
	// 2. Configuration of clearing query string settings for <query-string-settings/>.", "zh_CN":"查询串设置配置，父标签
	// 1.需要设置查询串配置时，此项必填
	// 2.为<query-string-settings/>时清空查询串设置的配置"}
	QueryStringSettings []*QueryDomainForTerraformResponseDataQueryStringSettings `json:"queryStringSettings,omitempty" xml:"queryStringSettings,omitempty" type:"Repeated"`
	// {"en":"Cache the file according to the response header content", "zh_CN":"根据响应头内容缓存文件
	//     注意：
	// 1、需要取消根据响应头内容缓存文件配置设置时，可以传入空节点<cache-by-repheaers></cache-by-repheaers>。
	//     2、表示需要设置根据响应头内容缓存文件配置时，此项必填"}
	CacheByRespHeaders []*QueryDomainForTerraformResponseDataCacheByRespHeaders `json:"cacheByRespHeaders,omitempty" xml:"cacheByRespHeaders,omitempty" require:"true" type:"Repeated"`
	// {"en":"Status Code Caching Rule Configuration, parent node
	// 1. When you need to set status code caching rules, this must be filled in.
	// 2. Configuration of Clear Status Code Caching Rules for <http-code-cache-rules/>.", "zh_CN":"状态码缓存规则配置，父标签
	// 1.需要设置状态码缓存规则时，此项必填
	// 2.为<http-code-cache-rules/>时清空状态码缓存规则配置"}
	HttpCodeCacheRules []*QueryDomainForTerraformResponseDataHttpCodeCacheRules `json:"httpCodeCacheRules,omitempty" xml:"httpCodeCacheRules,omitempty" require:"true" type:"Repeated"`
	// {"en":"Ignore protocol caching and push configuration, parent tags
	//
	// 1. This must be filled when protocol cache and push configuration need to be ignored
	// 2.<ignore-protocol-rules/>:Clear the configuration ignore about protocol cache and pushing", "zh_CN":"忽略协议缓存和推送配置，父标签
	// 1.需要设置忽略协议缓存和推送配置时，此项必填
	// 2.为<ignore-protocol-rules/>时清空忽略协议缓存和推送的配置"}
	IgnoreProtocolRules []*QueryDomainForTerraformResponseDataIgnoreProtocolRules `json:"ignoreProtocolRules,omitempty" xml:"ignoreProtocolRules,omitempty" require:"true" type:"Repeated"`
	// {"en":"Http2.0 settings, used to enable or disable http2.0, parent node.", "zh_CN":"http2.0设置，用于设置http2.0的开启或关闭，父标签"}
	Http2Settings *QueryDomainForTerraformResponseDataHttp2Settings `json:"http2Settings,omitempty" xml:"http2Settings,omitempty" type:"Struct"`
	// {"en":"Http header settings
	// note:
	// 1. When you need to cancel the http header setting, you can pass in the empty node <header-modify-rules></header-modify-rules>.
	// 2. indicating that you need to set the http header, this field is required", "zh_CN":"http头设置
	// 注意：
	// 1. 需要取消http头设置时，可以传入空节点<header-modify-rules></header-modify-rules>。
	// 2. 表示需要设置http头，此项必填"}
	HeaderModifyRules []*QueryDomainForTerraformResponseDataHeaderModifyRules `json:"headerModifyRules,omitempty" xml:"headerModifyRules,omitempty" type:"Repeated"`
	// {"en":"redirection function
	// note:
	// 1. Define a set of internal redirected content. If there is internal redirected content, this field is required.
	// 2. need to clear the content redirection content under the domain name, you can pass the empty node <rewrite-rule-settings></rewrite-rule-settings>", "zh_CN":"一级服务改写替换&mdash;二级服务 内部重定向
	// 注意：
	// 1. 定义一组内部重定向内容，，如果有使用内部重定向内容，此项必填
	// 2. 需要清空域名下的内容重定向内容，可以传入空节点<rewrite-rule-settings></rewrite-rule-settings>
	// 3. 如果有开启其他高级配置（如防盗链配置），有些配置可能会有配置冲突，建议先与技术支持人员确认"}
	RewriteRuleSettings []*QueryDomainForTerraformResponseDataRewriteRuleSettings `json:"rewriteRuleSettings,omitempty" xml:"rewriteRuleSettings,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseData) SetDomainId(v int64) *QueryDomainForTerraformResponseData {
	s.DomainId = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetCname(v string) *QueryDomainForTerraformResponseData {
	s.Cname = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetDomainName(v string) *QueryDomainForTerraformResponseData {
	s.DomainName = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetServiceType(v string) *QueryDomainForTerraformResponseData {
	s.ServiceType = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetServiceAreas(v string) *QueryDomainForTerraformResponseData {
	s.ServiceAreas = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetComment(v string) *QueryDomainForTerraformResponseData {
	s.Comment = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetHeaderOfClientIp(v string) *QueryDomainForTerraformResponseData {
	s.HeaderOfClientIp = &v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetOriginConfig(v *QueryDomainForTerraformResponseDataOriginConfig) *QueryDomainForTerraformResponseData {
	s.OriginConfig = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetSsl(v *QueryDomainForTerraformResponseDataSsl) *QueryDomainForTerraformResponseData {
	s.Ssl = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetCacheTimeBehaviors(v []*QueryDomainForTerraformResponseDataCacheTimeBehaviors) *QueryDomainForTerraformResponseData {
	s.CacheTimeBehaviors = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetCacheKeyRules(v []*QueryDomainForTerraformResponseDataCacheKeyRules) *QueryDomainForTerraformResponseData {
	s.CacheKeyRules = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetQueryStringSettings(v []*QueryDomainForTerraformResponseDataQueryStringSettings) *QueryDomainForTerraformResponseData {
	s.QueryStringSettings = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetCacheByRespHeaders(v []*QueryDomainForTerraformResponseDataCacheByRespHeaders) *QueryDomainForTerraformResponseData {
	s.CacheByRespHeaders = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetHttpCodeCacheRules(v []*QueryDomainForTerraformResponseDataHttpCodeCacheRules) *QueryDomainForTerraformResponseData {
	s.HttpCodeCacheRules = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetIgnoreProtocolRules(v []*QueryDomainForTerraformResponseDataIgnoreProtocolRules) *QueryDomainForTerraformResponseData {
	s.IgnoreProtocolRules = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetHttp2Settings(v *QueryDomainForTerraformResponseDataHttp2Settings) *QueryDomainForTerraformResponseData {
	s.Http2Settings = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetHeaderModifyRules(v []*QueryDomainForTerraformResponseDataHeaderModifyRules) *QueryDomainForTerraformResponseData {
	s.HeaderModifyRules = v
	return s
}

func (s *QueryDomainForTerraformResponseData) SetRewriteRuleSettings(v []*QueryDomainForTerraformResponseDataRewriteRuleSettings) *QueryDomainForTerraformResponseData {
	s.RewriteRuleSettings = v
	return s
}

type QueryDomainForTerraformResponseDataOriginConfig struct {
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
	OriginIps *string `json:"originIps,omitempty" xml:"originIps,omitempty" require:"true"`
	// {"en":"Back-to-origin HOST, used to change the HOST field in the back-to-origin HTTP request header. The supported formats are: ① domain name ③ ip
	// Note:
	// 1. Must comply with the ip/domain name format specification. If it is a domain name, the length of the domain name must be less than or equal to 128 characters.", "zh_CN":"回源HOST，用于更改回源HTTP请求头中的HOST字段。支持格式为: ①域名；②ip；
	// 注意：
	// 1、必须符合ip/域名格式规范。如果是域名，则域名长度小于等于128。"}
	DefaultOriginHostHeader *string `json:"defaultOriginHostHeader,omitempty" xml:"defaultOriginHostHeader,omitempty" require:"true"`
	// {"en":"useRange", "zh_CN":"设置range回源功能的开启或关闭，可选值为true、false
	// true：开启range回源
	// false：关闭range回源
	// 不传：表示不修改
	// 注意：
	// 1、Range是Http请求头，用于文件指定部分的请求。如：Range: bytes=0-999   就是请求该文件的前1000个字节。开启Range回源配置能够有效提高大文件分发效率，提升响应速度。源站需要支持range请求，否则会导致回源失败。
	// 2、下载加速的域名默认range回源为开启状态，网页加速的域名默认为关闭状态"}
	UseRange *string `json:"useRange,omitempty" xml:"useRange,omitempty" require:"true"`
	// {"en":"follow301", "zh_CN":"设置回源跟随301。当节点回源请求返回301状态码时，CDN节点会直接向跳转地址请求资源而不返回301给用户。可选值为true、false
	// true：开启301跟随
	// false：关闭301跟随
	// 不传：表示不修改"}
	Follow301 *string `json:"follow301,omitempty" xml:"follow301,omitempty" require:"true"`
	// {"en":"follow302", "zh_CN":"设置回源跟随302。当节点回源请求返回302状态码时，CDN节点会直接向跳转地址请求资源而不返回302给用户。可选值为true、false
	// true：开启302跟随
	// false：关闭302跟随
	// 不传：表示不修改"}
	Follow302 *string `json:"follow302,omitempty" xml:"follow302,omitempty" require:"true"`
	// {"en":"advance origin config", "zh_CN":"高级源配置"}
	AdvSrcSetting *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting `json:"advSrcSetting,omitempty" xml:"advSrcSetting,omitempty" require:"true" type:"Struct"`
}

func (s QueryDomainForTerraformResponseDataOriginConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataOriginConfig) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataOriginConfig) SetOriginIps(v string) *QueryDomainForTerraformResponseDataOriginConfig {
	s.OriginIps = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfig) SetDefaultOriginHostHeader(v string) *QueryDomainForTerraformResponseDataOriginConfig {
	s.DefaultOriginHostHeader = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfig) SetUseRange(v string) *QueryDomainForTerraformResponseDataOriginConfig {
	s.UseRange = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfig) SetFollow301(v string) *QueryDomainForTerraformResponseDataOriginConfig {
	s.Follow301 = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfig) SetFollow302(v string) *QueryDomainForTerraformResponseDataOriginConfig {
	s.Follow302 = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfig) SetAdvSrcSetting(v *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) *QueryDomainForTerraformResponseDataOriginConfig {
	s.AdvSrcSetting = v
	return s
}

type QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting struct {
	// {"en":"Use advance origin config, the optional values are true and false, true means to use advance origin config, false means not to use advance origin config","zh_CN":"使用高级源，可选值为true和false，true表示使用高级源，false表示不使用高级源"}
	UseAdvSrc *string `json:"useAdvSrc,omitempty" xml:"useAdvSrc,omitempty" require:"true"`
	// {"en":"The advanced source monitors the url, and requests <master-ips> through the url. If the response is not 2**, 3** response, it is considered that the primary source ip is faulty, and <backup-ips> is used at this time.", "zh_CN":"高级源监控url，通过该url请求<master-ips>，如果返回非2**，3**响应时，认为主要回源ip故障，此时使用<backup-ips>。
	// 完整的url，例如：http://a.example.com/test.html"}
	DetectUrl *string `json:"detectUrl,omitempty" xml:"detectUrl,omitempty" require:"true"`
	// {"en":"Advanced source monitoring period, in seconds, optional as an integer greater than or equal to 0, 0 means no monitoring", "zh_CN":"高级源监控周期，单位秒，可选值为大于等于0的整数，0表示不监控"}
	DetectPeriod *string `json:"detectPeriod,omitempty" xml:"detectPeriod,omitempty" require:"true"`
	// {"en":"The advanced source mainly returns the source IP. Multiple IPs are separated by a semicolon \";\", and the return source IP cannot be repeated.", "zh_CN":"高级源主要回源IP，多个IP用分号“;”分隔，回源IP不能重复"}
	MasterIps []*string `json:"masterIps,omitempty" xml:"masterIps,omitempty" require:"true" type:"Repeated"`
	// {"en":"Advanced source backup source IP, multiple IPs are separated by semicolon \";\", and the return source IP cannot be duplicated.", "zh_CN":"高级源备用回源IP，多个IP用分号“;”分隔，回源IP不能重复"}
	BackupIps []*string `json:"backupIps,omitempty" xml:"backupIps,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) SetUseAdvSrc(v string) *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting {
	s.UseAdvSrc = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) SetDetectUrl(v string) *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting {
	s.DetectUrl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) SetDetectPeriod(v string) *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting {
	s.DetectPeriod = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) SetMasterIps(v []*string) *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting {
	s.MasterIps = v
	return s
}

func (s *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting) SetBackupIps(v []*string) *QueryDomainForTerraformResponseDataOriginConfigAdvSrcSetting {
	s.BackupIps = v
	return s
}

type QueryDomainForTerraformResponseDataSsl struct {
	// {"en":"Use a certificate, the optional values are true and false, true means to use the certificate, false means not to use the certificate", "zh_CN":"使用证书，可选值为true和false，true表示使用证书，false表示不使用证书"}
	UseSsl *string `json:"useSsl,omitempty" xml:"useSsl,omitempty" require:"true"`
	// {"en":"Use sni certificate, the optional values are true and false, true means use sni certificate, false means use shared certificate (not supported)", "zh_CN":"证书ID，新增证书成功后，系统返回的证书ID
	// useSsl为true时，才能传sslCertificateId。"}
	SslCertificateId *string `json:"sslCertificateId,omitempty" xml:"sslCertificateId,omitempty" require:"true"`
	// {"en":"Backup certificate ID", "zh_CN":"备用证书ID"}
	BackupCertificateId *string `json:"backupCertificateId,omitempty" xml:"backupCertificateId,omitempty" require:"true"`
	// {"en":"SM2 certificate IDS", "zh_CN":"国密证书ID列表"}
	GmCertificateIds []*string `json:"gmCertificateIds,omitempty" xml:"gmCertificateIds,omitempty" require:"true" type:"Repeated"`
	// {"en":"TLS version. Optional values: SSLv3,TLSv1,TLSv1.1,TLSv1.2,TLSv1.3", "zh_CN":"TLS协议（TLSVersion），支持配置多个，英文分号分隔。可选值为: SSLv3、TLSv1、TLSv1.1、TLSv1.2、TLSv1.3。域名有配置证书时，该配置才有效。"}
	TlsVersion *string `json:"tlsVersion,omitempty" xml:"tlsVersion,omitempty" require:"true"`
	// {"en":"Enable OCSP(Online Certificate Status Protocol).", "zh_CN":"支持OCSP，默认不启用，可选值true、false。域名有配置证书时，该配置才有效。"}
	EnableOcsp *string `json:"enableOcsp,omitempty" xml:"enableOcsp,omitempty" require:"true"`
	// {"en":"This optional object is used to specify a colon separated list of cipher suites which are permitted when clients negotiate security settings to access your content. Cipher suites which you can specify are: LOW, ALL:!LOW, HIGH, !EXPORT, !aNULL, !RC4, !DH, !SHA, !MD5, @STRENGTH,  AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256, AES128-GCM-SHA256, AES256-GCM-SHA384, ECDHE-RSA-AES128-SHA, ECDHE-RSA-AES256-SHA, ECDHE-RSA-AES128-SHA256, ECDHE-RSA-AES256-SHA384, ECDHE-RSA-AES128-GCM-SHA256, and ECDHE-RSA-AES256-GCM-SHA384. These cipher suites are a subset of those supported by OpenSSL, https://www.openssl.org/docs/man1.0.2/man1/ciphers.html. Please note that !MD5 or !SHA must appear after HIGH..", "zh_CN":"设置cipher加密套件，冒号分隔。请注意：系统已有默认的算法，如无特殊修改需要，无需调整。"}
	SslCipherSuite *string `json:"sslCipherSuite,omitempty" xml:"sslCipherSuite,omitempty" require:"true"`
}

func (s QueryDomainForTerraformResponseDataSsl) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataSsl) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataSsl) SetUseSsl(v string) *QueryDomainForTerraformResponseDataSsl {
	s.UseSsl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataSsl) SetSslCertificateId(v string) *QueryDomainForTerraformResponseDataSsl {
	s.SslCertificateId = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataSsl) SetBackupCertificateId(v string) *QueryDomainForTerraformResponseDataSsl {
	s.BackupCertificateId = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataSsl) SetGmCertificateIds(v []*string) *QueryDomainForTerraformResponseDataSsl {
	s.GmCertificateIds = v
	return s
}

func (s *QueryDomainForTerraformResponseDataSsl) SetTlsVersion(v string) *QueryDomainForTerraformResponseDataSsl {
	s.TlsVersion = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataSsl) SetEnableOcsp(v string) *QueryDomainForTerraformResponseDataSsl {
	s.EnableOcsp = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataSsl) SetSslCipherSuite(v string) *QueryDomainForTerraformResponseDataSsl {
	s.SslCipherSuite = &v
	return s
}

type QueryDomainForTerraformResponseDataCacheTimeBehaviors struct {
	// {"en":"dataId is to indicate a specific group configuration when the client has multiple groups of configurations. dataId can be retrieved through a query interface.
	// Note:
	// 1. If dataId is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified.
	// 2. If multiple groups of configurations are included, some of them are configured with dataId and others are not, then the expression of dataId is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of dataId.
	// 3. If the dataId is not passed, it means that the original configuration will be fully covered by this configuration.
	// 4. If no configuration parameter is passed, only domain name and secondary label are passed. It means that all configuration corresponding to this interface is cleared.
	// 5. If there is no specific configuration item in a set of configurations, the dataId must be filled in, and the value should be the actual dataId, which means clearing the value of the corresponding dataId configuration item.
	// 6. Tt is not allowed when neither configuration item nor dataId is specified in a set of configurations.",
	//     "zh_CN":"中文：配置多组配置时，具体某组配置的id。dataId可以通过查询接口获取。
	// 注意：
	// 1、如果有传dataId，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
	// 2、如果入参多组配置，其中有些组配置有传dataId，有些没有传，则有传dataId的表示修改具体某组配置，没有传dataId的表示在原来基础上新增一组配置；
	// 3、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置；
	// 4、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个该接口对应的域名配置；
	// 5、如果一组配置没有具体的配置项，则dataId必填，且值要为实际存在的dataId，此时表示清空这个dataId对应的一组配置；
	// 6、不允许一组配置没有指定具体的配置项也没有dataId。"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：*"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty"`
	// {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
	// E.g: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
	// 客户入参参考：^https?://[^/]+/.*\.m3u8"}
	ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
	// {"en":"Specify common types: Select the domain name that requires the cache  to be all files or the home page. :
	// E.g:
	// All: all files
	// Homepage: homepage", "zh_CN":"指定常用类型：选择缓存域名的是全部文件还是首页。入参参考值：
	// all：全部文件
	// homepage：首页"}
	CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty"`
	// {"en":"File Type: Specify the file type for cache settings.
	// File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
	// If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定需要缓存的文件类型。
	// 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
	// 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
	CustomFileType *string `json:"customFileType,omitempty" xml:"customFileType,omitempty"`
	// {"en":"Specify URL cache: Specify url according to requirements for cache
	// INS format does not support URI format with http(s)://", "zh_CN":"指定URL缓存：根据需求指定url进行缓存
	// 入参不支持含http(s):// 开头的URI格式"}
	SpecifyUrlPattern *string `json:"specifyUrlPattern,omitempty" xml:"specifyUrlPattern,omitempty"`
	// {"en":"Directory: Specify the directory cache.
	// Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录缓存。
	// 输入合法的目录格式。多个以英文分号隔开"}
	Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
	// {"en":"Cache time: set the time corresponding to the cache object
	// Input format: integer plus unit, such as 20s, 30m, 1h, 2d, no cache is set to 0. Do not enter the unit default is seconds
	// There is no upper limit on the cache time theory. This time is set according to the customer's own needs. If the customer feels that some of the files are not changed frequently, then the setting is longer. For example, the text class js, css, html, etc. can be set shorter, the picture, video and audio classes can be set longer (because the cache time will be replaced by the new file due to the file heat algorithm, the longest suggestion Do not exceed one month)", "zh_CN":"缓存时间：设置缓存对象对应的时间
	// 入参格式：整数加单位，比如20s、30m、1h、2d，不缓存设置为0。不输入单位默认是秒
	// 缓存时间理论上没有上限限制，这个时间根据客户自身的需求设定，如果客户觉得其中一些文件，变更不频繁，那么就设置长一点。例如，文本类的js，css，html等可以设置得短一些，图片、视频音频类的可以设置的长一点（因为缓存时间会因文件热度算法，旧文件会被新文件替换掉，最长建议不要超过一个月）"}
	CacheTtl *string `json:"cacheTtl,omitempty" xml:"cacheTtl,omitempty"`
	// {"en":"Ignore the source station does not cache the header. The optional values are true and false, which are used to ignore the two configurations of cache-control in the request header (private, no-cache) and the Authorization set by the client.
	// The ture indicates that the source station's settings for the three are ignored. Enables resources to be cached on the service node in the form of cache-control: public, and then our nodes can cache this type of resource and provide acceleration services.
	// False means that when the source station sets cache-control: private, cache-control: no-cache for a resource or specifies to cache according to authorization, our service node will not cache such files.", "zh_CN":"忽略源站不缓存头。可选值为true和false，用于忽略请求头中cache-control的两种配置（private，no-cache）和客户端设置的Authorization。
	// ture表示会忽略掉源站对于这三者的设定。使得资源能够以cache-control: public的方式缓存在服务节点上，然后我们的节点才能缓存这种类型的资源，提供加速服务。
	// false表示当源站对某种资源设定了cache-control: private,cache-control:no-cache或指定根据authorization进行缓存时，我们的服务节点将不会对此类文件进行缓存。"}
	IgnoreCacheControl *string `json:"ignoreCacheControl,omitempty" xml:"ignoreCacheControl,omitempty"`
	// {"en":"Respect the server: Accelerate whether to prioritize the source cache time.
	// Optional values: true and false
	// True: indicates that the server is time-first
	// False: The cache time of the CDN configuration takes precedence.", "zh_CN":"尊重服务端：加速是否要按源站缓存时间优先。
	// 可选值：true和false
	// true：表示重服务端时间优先
	// false:CDN配置的缓存时间优先"}
	IsRespectServer *string `json:"isRespectServer,omitempty" xml:"isRespectServer,omitempty"`
	// {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
	// When adding a new configuration item, the default is not true.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
	// 新增配置项时，不传默认为 true"}
	IgnoreLetterCase *string `json:"ignoreLetterCase,omitempty" xml:"ignoreLetterCase,omitempty"`
	// {"en":"Reload processing rules, optional: ignore or if-modified-since
	// If-modified-since: indicates that you want to convert to if-modified-since
	// Ignore: means to ignore client refresh", "zh_CN":"reload处理规则，可选项：ignore或者if-modified-since
	// if-modified-since：表示要转成if-modified-since
	// ignore:表示忽略客户端刷新"}
	ReloadManage *string `json:"reloadManage,omitempty" xml:"reloadManage,omitempty"`
	// {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
	// When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
	// 新增配置项时，不传默认为 10
	// 如果传了值，不能为空"}
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// {"en":"You can set it 'true' to cache
	// ignoring the http header 'Authentication'.  If it is empty, the header is not ignored by default.", "zh_CN":"忽略鉴权头部Authentication，可选值为true和false。默认为不忽略。"}
	IgnoreAuthenticationHeader *string `json:"ignoreAuthenticationHeader,omitempty" xml:"ignoreAuthenticationHeader,omitempty"`
}

func (s QueryDomainForTerraformResponseDataCacheTimeBehaviors) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataCacheTimeBehaviors) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetDataId(v int64) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.DataId = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetPathPattern(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetExceptPathPattern(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.ExceptPathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetCustomPattern(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.CustomPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetFileType(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.FileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetCustomFileType(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.CustomFileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetSpecifyUrlPattern(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.SpecifyUrlPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetDirectory(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.Directory = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetCacheTtl(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.CacheTtl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetIgnoreCacheControl(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.IgnoreCacheControl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetIsRespectServer(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.IsRespectServer = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetIgnoreLetterCase(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.IgnoreLetterCase = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetReloadManage(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.ReloadManage = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetPriority(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.Priority = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheTimeBehaviors) SetIgnoreAuthenticationHeader(v string) *QueryDomainForTerraformResponseDataCacheTimeBehaviors {
	s.IgnoreAuthenticationHeader = &v
	return s
}

type QueryDomainForTerraformResponseDataCacheKeyRules struct {
	// {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty" require:"true"`
	// {"en":"Specify a uri, such as /test/specifyurl", "zh_CN":"指定具体的url，如/test/specifyurl"}
	SpecifyUrl *string `json:"specifyUrl,omitempty" xml:"specifyUrl,omitempty" require:"true"`
	// {"en":"Whether to match specifyUrl exactly or not, you can select true and false.
	// True:means match exactly. False: means fuzzy match, such as specifyUrl='/test/uri', and  request for /test/uri?p=1 will be matched.", "zh_CN":"是否完全匹配specifyUrl，可选择为true和false。
	// 为true则完全匹配；为false则模糊匹配，如指定/test/uri，请求/test/uri?p=1也会匹配"}
	FullMatch4SpecifyUrl *string `json:"fullMatch4SpecifyUrl,omitempty" xml:"fullMatch4SpecifyUrl,omitempty" require:"true"`
	// {"en":"Specify common types: Select the domain name that requires the cache to be all files or the home page. :
	// E.g:
	// All: all files
	// Homepage: homepage", "zh_CN":"指定常用类型：选择缓存域名的是全部文件还是首页。入参参考值： all：全部文件 homepage：首页"}
	CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty" require:"true"`
	// {"en":"File Type: Specify the file type for cache settings.
	// File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
	// If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定需要缓存的文件类型。 文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置。"}
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty" require:"true"`
	// {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
	CustomFileType *string `json:"customFileType,omitempty" xml:"customFileType,omitempty" require:"true"`
	// {"en":"Directory: Specify the directory cache.
	// Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录：指定目录缓存。 输入合法的目录格式。多个以英文分号隔开"}
	Directory *string `json:"directory,omitempty" xml:"directory,omitempty" require:"true"`
	// {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
	// When adding a new configuration item, the default is true", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
	IgnoreCase *string `json:"ignoreCase,omitempty" xml:"ignoreCase,omitempty" require:"true"`
	// {"en":"Header name.
	// Example: If you specify a header as &lsquo;lang', Then, if the value of Lang is consistent, one copy will be cached", "zh_CN":"头部名称
	// 例如：指定头部lang，lang的值一致则缓存一份"}
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty" require:"true"`
	// {"en":"Parameter Of the specified Header,
	// Example: Specifies the header as 'cookie', parameterOfHeader as 'name'. Then, if the value of name is consistent, one copy will be cached.", "zh_CN":"头部值的参数名，
	// 例如：指定头部Cookie，头部值的参数名为name。则name的值一致则缓存一份。"}
	ParameterOfHeader *string `json:"parameterOfHeader,omitempty" xml:"parameterOfHeader,omitempty" require:"true"`
	// {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
	// When adding a new configuration item, the default is 10", "zh_CN":"优先级，表示客户多组配置的优先执行顺序。数字越大，优先级越高。不传默认为10，不可清空。"}
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty" require:"true"`
	// {"en":"DataId is to indicate a specific group configuration when the client has multiple groups of configurations. dataId can be retrieved through a query interface. Note: A. If dataId is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with dataId and others are not, then the expression of dataId is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of dataId. C. If the dataId is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the dataId must be filled in, and the value is the actual dataId, which means clearing the value of the corresponding dataId configuration item; it is not allowed that there is no specific configuration item or dataId in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。本功能只支持一组配置。dataId可以通过查询接口获取。 注意： a、如果有传dataId，说明指定修改该组配置项内容； b、如果入参都没有传dataId,表示用本次的配置全量覆盖原先配置； c、如果一组配置没有具体的配置项，则dataId必填，且值为实际存在的dataId，表示清空这个dataId对应配置项的值；不允许一组配置没有具体的配置项也没有dataId。"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryDomainForTerraformResponseDataCacheKeyRules) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataCacheKeyRules) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetPathPattern(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetSpecifyUrl(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.SpecifyUrl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetFullMatch4SpecifyUrl(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.FullMatch4SpecifyUrl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetCustomPattern(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.CustomPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetFileType(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.FileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetCustomFileType(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.CustomFileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetDirectory(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.Directory = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetIgnoreCase(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.IgnoreCase = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetHeaderName(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.HeaderName = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetParameterOfHeader(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.ParameterOfHeader = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetPriority(v string) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.Priority = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheKeyRules) SetDataId(v int64) *QueryDomainForTerraformResponseDataCacheKeyRules {
	s.DataId = &v
	return s
}

type QueryDomainForTerraformResponseDataQueryStringSettings struct {
	// {"en":"The url matching mode. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*
	// 注：url匹配模式、文件类型（自定义文件类型）、常用类型、指定url、目录，有且仅有一项必填"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty"`
	// {"en":"Exception to url matching pattern, except for some urls: abc.jpg, no content redirection", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
	//     客户入参参考：^https?://[^/]+/.*\.m3u8"}
	ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
	// {"en":"File Type: Specify the file type for anti-theft chain settings.
	// File types include: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf
	// If you need all types, pass all directly. Multiples are separated by semicolons, and all and specific file types cannot be configured at the same time.", "zh_CN":"文件类型：指定文件类型进行防盗链设置。可选文件类型包括：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts 如果需要全部类型，则直接传all。多个以分号隔开，all和具体文件类型不能同时配置"}
	FileTypes *string `json:"fileTypes,omitempty" xml:"fileTypes,omitempty"`
	// {"en":"Custom file type: Fill in the appropriate identifiable file type according to your needs outside of the specified file type. Can be used with file-type. If the file-type is also configured, the actual file type is the sum of the two parameters.", "zh_CN":"自定义文件类型：在指定文件类型外根据自身需求，填写适当的可识别文件类型。可以搭配file-type使用。如果file-type也有配置，实际生效的文件类型是两个入参的总和"}
	CustomFileTypes *string `json:"customFileTypes,omitempty" xml:"customFileTypes,omitempty"`
	// {"en":"Specify common types: Select the domain name that requires the anti-theft chain to be all files or the home page. :
	// E.g:
	// All: all files
	// Homepage: homepage", "zh_CN":"常用类型：可选值为homepage和all
	// all：全部文件
	// homepage：首页"}
	CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty"`
	// {"en":"Specify URL cache: Specify url according to requirements for anti-theft chain setting
	// INS format does not support URI format with http(s)://", "zh_CN":"指定url，非http(s):// 开头，多个以换行分隔"}
	SpecifyUrlPattern *string `json:"specifyUrlPattern,omitempty" xml:"specifyUrlPattern,omitempty"`
	// {"en":"Directory: Specify the directory for anti-theft chain settings
	// Enter a legal directory format. Multiple separated by semicolons", "zh_CN":"目录，可输入合法的目录格式。以/开头和结尾，多个以分号隔开。"}
	Directories *string `json:"directories,omitempty" xml:"directories,omitempty"`
	// {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
	// When adding a new configuration item, the default is 10", "zh_CN":"优先级，表示客户多组配置的优先执行顺序。数字越大，优先级越高。不传默认为10，不可清空。"}
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// {"en":"Whether to ignore letter case.", "zh_CN":"是否忽略大小写：允许值为true和false，默认为忽略"}
	IgnoreLetterCase *string `json:"ignoreLetterCase,omitempty" xml:"ignoreLetterCase,omitempty"`
	// {"en":"Define whether to cache with query strings, 'true' means ignore query strings, while 'false' means cache with all query strings.", "zh_CN":"缓存是否忽略查询串，允许值为true和false。
	// true表示忽略查询串，相同拷贝；false表示不忽略，分别缓存，即带全部参数缓存。"}
	IgnoreQueryString *string `json:"ignoreQueryString,omitempty" xml:"ignoreQueryString,omitempty"`
	// {"en":"Cache with the specified query string parameters. If the kept parameter values are the same, one copy will be cached.
	// Note:
	// 1. query-string-kept and query-string-removed are mutually exclusive, and only one of them has a value.
	// 2. query-string-kept and ignore-query-string are mutually exclusive, and only one has a value.", "zh_CN":"缓存保留参数，指定保留的参数值相同，则缓存一份。
	// 注：
	// 1.query-string-kept和query-string-removed两者互斥，只能一个有值。
	// 2.query-string-kept和ignore-query-string两者互斥，只能一个有值。"}
	QueryStringKept *string `json:"queryStringKept,omitempty" xml:"queryStringKept,omitempty"`
	// {"en":"Cache without the specified query string parameters. After deleting the specified parameter, if the other parameter values are the same, one copy will be cached.
	// 1. query-string-kept and query string removed are mutually exclusive, and only one has a value.
	// 2. query-string-removed and ignore-query-string are mutually exclusive.", "zh_CN":"缓存删除参数，删除指定的参数后，其余参数值相同，则缓存一份。
	// 1.query-string-kept和query-string-removed两者互斥，只能一个有值。
	// 2.query-string-removed和ignore-query-string两者互斥，只能一个有值。"}
	QueryStringRemoved *string `json:"queryStringRemoved,omitempty" xml:"queryStringRemoved,omitempty"`
	// {"en":"Whether to use the original URL back source, the allowable values are true and false.
	// When ignore-query-string is true or not set, source-with-query is true to indicate that the source is returned according to the original request, and false to indicate that the question mark is returned.
	// When ignore-query-string is false, this default setting is empty (input is invalid)", "zh_CN":"是否用原始url回源，允许值为true和false。
	// ignore-query-string为true或未设置时，source-with-query为true表示按原始请求回源；为false表示去问号回源。
	// ignore-query-string为false时，此项默认设置为空（输入无效）。"}
	SourceWithQuery *string `json:"sourceWithQuery,omitempty" xml:"sourceWithQuery,omitempty"`
	// {"en":"Return to the source after specifying the reserved parameter value. Please separate them with semicolons, if no parameters reserved, please fill in:- . 1. Source-key-kept and ignore-query-string are mutually exclusive, and only one of them has a value. 2. Source-key-kept and source-key-removed are mutually exclusive, and only one of them has a value.
	// ", "zh_CN":"回源保留参数，指定保留的参数值后回源。多个请以英文分号分隔，任何参数都不保留请填：- 1、source-key-kept和ignore-query-string两者互斥，只能一个有值。 2、source-key-kept和source-key-removed两者互斥，只能一个有值。
	// "}
	SourceKeyKept *string `json:"sourceKeyKept,omitempty" xml:"sourceKeyKept,omitempty"`
	// {"en":"Return to the source after specifying the deleted parameter value. Please separate them with semicolons, and if you do not delete any parameters, please fill in:- . 1. Source-key-removed and ignore-query-string are mutually exclusive, and only one of them has a value. 2. Source-key-kept and source-key-removed are mutually exclusive, and only one of them has a value.
	// ", "zh_CN":"回源删除参数，指定删除的参数值后回源。多个请以英文分号分隔，任何参数都不删除请填：- 1、source-key-removed和ignore-query-string两者互斥，只能一个有值。 2、source-key-kept和source-key-removed两者互斥，只能一个有值。
	// "}
	SourceKeyRemoved *string `json:"sourceKeyRemoved,omitempty" xml:"sourceKeyRemoved,omitempty"`
	// {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note: A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified. B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id. C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration. D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared. E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意： a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；  b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；  c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置；  d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置；  e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s QueryDomainForTerraformResponseDataQueryStringSettings) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataQueryStringSettings) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetPathPattern(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetExceptPathPattern(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.ExceptPathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetFileTypes(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.FileTypes = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetCustomFileTypes(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.CustomFileTypes = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetCustomPattern(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.CustomPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetSpecifyUrlPattern(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.SpecifyUrlPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetDirectories(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.Directories = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetPriority(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.Priority = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetIgnoreLetterCase(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.IgnoreLetterCase = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetIgnoreQueryString(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.IgnoreQueryString = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetQueryStringKept(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.QueryStringKept = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetQueryStringRemoved(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.QueryStringRemoved = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetSourceWithQuery(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.SourceWithQuery = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetSourceKeyKept(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.SourceKeyKept = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetSourceKeyRemoved(v string) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.SourceKeyRemoved = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataQueryStringSettings) SetDataId(v int64) *QueryDomainForTerraformResponseDataQueryStringSettings {
	s.DataId = &v
	return s
}

type QueryDomainForTerraformResponseDataCacheByRespHeaders struct {
	// {"en":"Add grid type identity, represents the customer multi - group configuration, a specific group of configuration", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置
	//         如果是新增一组配置项的值时，不需要传。如果指定修改具体data-id的配置项值时，需要传入对应配置项对应的data-id.可以通过查询接口获取"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// {"en":"The name of the response header", "zh_CN":"响应头头部名称
	//     需要缓存的响应头名称。如Cache-Control"}
	ResponseHeader *string `json:"responseHeader,omitempty" xml:"responseHeader,omitempty" require:"true"`
	// {"en":"Url matching pattern, support fuzzy regular, if all match, the parameter can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty"`
	// {"en":"Exception to url matching pattern, except for some urls: abc.jpg, no content redirection", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
	//     客户入参参考：^https?://[^/]+/.*\.m3u8"}
	ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
	// {"en":"Header values.", "zh_CN":"头部值。"}
	ResponseValue *string `json:"responseValue,omitempty" xml:"responseValue,omitempty" require:"true"`
	// {"en":"Ignore case, the optional value is true or false, true means ignore case;False means that case is not ignored.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
	//     新增配置项时，不传默认为 true"}
	IgnoreLetterCase *string `json:"ignoreLetterCase,omitempty" xml:"ignoreLetterCase,omitempty"`
	// {"en":"Represents the priority execution order for multiple groups of customer redirected content.The larger the number, the higher the priority.", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
	//     新增配置项时，不传默认为 10
	//     如果传了值，为清空配置"}
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// {"en":"Whether the response header content cache file is allowed. An optional value of true or false indicates that the response header content cache file is allowed.False indicates that the response header content cache file is not allowed.", "zh_CN":"是否允许响应头内容缓存文件，可选值为true或false，true表示允许响应头内容缓存文件；false表示不允许响应头内容缓存文件；"}
	IsRespheader *string `json:"isRespheader,omitempty" xml:"isRespheader,omitempty" require:"true"`
}

func (s QueryDomainForTerraformResponseDataCacheByRespHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataCacheByRespHeaders) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetDataId(v int64) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.DataId = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetResponseHeader(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.ResponseHeader = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetPathPattern(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetExceptPathPattern(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.ExceptPathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetResponseValue(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.ResponseValue = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetIgnoreLetterCase(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.IgnoreLetterCase = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetPriority(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.Priority = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataCacheByRespHeaders) SetIsRespheader(v string) *QueryDomainForTerraformResponseDataCacheByRespHeaders {
	s.IsRespheader = &v
	return s
}

type QueryDomainForTerraformResponseDataHttpCodeCacheRules struct {
	// {"en":"Configure HTTP status code, parent node", "zh_CN":"配置http状态码，父标签"}
	HttpCodes []*string `json:"httpCodes,omitempty" xml:"httpCodes,omitempty" require:"true" type:"Repeated"`
	// {"en":"Define the caching time of the specified status code in units s, 0 to indicate no caching", "zh_CN":"配置指定的状态码的缓存时间，单位s，0表示不缓存"}
	CacheTtl *string `json:"cacheTtl,omitempty" xml:"cacheTtl,omitempty" require:"true"`
	// {"en":"Data-id is to indicate a specific group configuration when the client has multiple groups of configurations. Data-id can be retrieved through a query interface. Note:
	// A. If data-id is passed, it means that one group of configuration items is specified to be modified, and no other group configuration items need to be modified.
	// B. If multiple groups of configurations are included, some of them are configured with data-id and others are not, then the expression of data-id is used to modify a specific group of configurations, and a new group of configurations is added on the original basis without the expression of data-id.
	// C. If the data-id is not transmitted, it means that the original configuration will be fully covered by this configuration.
	// D. If no configuration parameter is passed, only domain name and secondary label are passed, which means that all configuration of domain name secondary service corresponding to this interface is cleared.
	// E. If there is no specific configuration item in a set of configurations, the data-id must be filled in, and the value is the actual data-id, which means clearing the value of the corresponding data-id configuration item; it is not allowed that there is no specific configuration item or data-id in a set of configurations.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意：
	// a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
	// b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；
	// c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置；
	// d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置；
	// e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty" require:"true"`
}

func (s QueryDomainForTerraformResponseDataHttpCodeCacheRules) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataHttpCodeCacheRules) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataHttpCodeCacheRules) SetHttpCodes(v []*string) *QueryDomainForTerraformResponseDataHttpCodeCacheRules {
	s.HttpCodes = v
	return s
}

func (s *QueryDomainForTerraformResponseDataHttpCodeCacheRules) SetCacheTtl(v string) *QueryDomainForTerraformResponseDataHttpCodeCacheRules {
	s.CacheTtl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHttpCodeCacheRules) SetDataId(v int64) *QueryDomainForTerraformResponseDataHttpCodeCacheRules {
	s.DataId = &v
	return s
}

type QueryDomainForTerraformResponseDataIgnoreProtocolRules struct {
	// {"en":"Url matching pattern, support regular, if all matches, input parameters can be configured as:.*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty" require:"true"`
	// {"en":"The exception url matches the pattern in the same format as the path-pattern", "zh_CN":"例外的url匹配模式，格式同path-pattern"}
	ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
	// {"en":"Whether to ignore the protocol cache, with allowable values of true and false. True turns on the HTTP/HTTPS Shared cache. Not on by default.", "zh_CN":"是否忽略协议缓存，允许值为true和false。为true则开启http/https共用缓存。默认不开启。"}
	CacheIgnoreProtocol *string `json:"cacheIgnoreProtocol,omitempty" xml:"cacheIgnoreProtocol,omitempty"`
	// {"en":"It is recommended to use with cache-ignore protocol to avoid push failure.
	//
	// Note:
	//
	// 1. Once configured, the global effect is not applied to the matched path-pattern.
	//
	// 2. Directory push does not distinguish protocols, while url push can distinguish protocols", "zh_CN":"是否忽略协议推送，允许值为true和false。为true则推送时忽略协议；为false则区分协议推送。
	// 建议和cache-ignore-protocol配套使用，避免推送失效。
	// 注意：
	// 1.一旦配置，则全局生效，不针对匹配的path-pattern生效。
	// 2.目录推送不区分协议，url推送可区分协议"}
	PurgeIgnoreProtocol *string `json:"purgeIgnoreProtocol,omitempty" xml:"purgeIgnoreProtocol,omitempty"`
	// {"en":"When configuring multiple groups of configurations, specify the id of a particular group of configurations", "zh_CN":"配置多组配置时，具体某组配置的id"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
}

func (s QueryDomainForTerraformResponseDataIgnoreProtocolRules) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataIgnoreProtocolRules) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataIgnoreProtocolRules) SetPathPattern(v string) *QueryDomainForTerraformResponseDataIgnoreProtocolRules {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataIgnoreProtocolRules) SetExceptPathPattern(v string) *QueryDomainForTerraformResponseDataIgnoreProtocolRules {
	s.ExceptPathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataIgnoreProtocolRules) SetCacheIgnoreProtocol(v string) *QueryDomainForTerraformResponseDataIgnoreProtocolRules {
	s.CacheIgnoreProtocol = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataIgnoreProtocolRules) SetPurgeIgnoreProtocol(v string) *QueryDomainForTerraformResponseDataIgnoreProtocolRules {
	s.PurgeIgnoreProtocol = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataIgnoreProtocolRules) SetDataId(v int64) *QueryDomainForTerraformResponseDataIgnoreProtocolRules {
	s.DataId = &v
	return s
}

type QueryDomainForTerraformResponseDataHttp2Settings struct {
	// {"en":"Enable http2.0. The optional values are true and false. If it is empty, the default value is false. True means http2.0 is on; false means http2.0 is off.", "zh_CN":"开启http2.0，可选值为true和false，为空时默认为false。true表示开启http2.0；false表示关闭http2.0"}
	EnableHttp2 *string `json:"enableHttp2,omitempty" xml:"enableHttp2,omitempty" require:"true"`
	// {"en":"Back-to-origin protocol, the optional value is
	// http1.1: Use the HTTP1.1 protocol version to back to source. if not filled, use it as default.
	// follow-request: Same as client request protocol
	// http2.0: Use the HTTP2.0 protocol. version to back to source.", "zh_CN":"回源协议，可选值为
	// http1.1：使用HTTP1.1协议版本回源，不填时默认该协议
	// follow-request：跟随客户端请求协议
	// http2.0：使用HTTP2.0协议版本回源"}
	BackToOriginProtocol *string `json:"backToOriginProtocol,omitempty" xml:"backToOriginProtocol,omitempty" require:"true"`
}

func (s QueryDomainForTerraformResponseDataHttp2Settings) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataHttp2Settings) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataHttp2Settings) SetEnableHttp2(v string) *QueryDomainForTerraformResponseDataHttp2Settings {
	s.EnableHttp2 = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHttp2Settings) SetBackToOriginProtocol(v string) *QueryDomainForTerraformResponseDataHttp2Settings {
	s.BackToOriginProtocol = &v
	return s
}

type QueryDomainForTerraformResponseDataHeaderModifyRules struct {
	// {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置。
	// data-id可以通过查询接口获取。
	// 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：
	// a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
	// b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；
	// c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；
	// d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；
	// e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty"`
	// {"en":"Exception url matching pattern, support regular. Example: ", "zh_CN":"例外的url匹配模式，支持正则。 入参参考："}
	ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
	// {"en":"Matching conditions: specify common types, optional values are all or homepage. 1. all: all files 2. homepage: home page", "zh_CN":"匹配条件：指定常用类型，可选值为all或homepage 1. all：全部文件 2. homepage：首页"}
	CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty"`
	// {"en":"Matching conditions: file type, please separate by semicolon, optional values: gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts.", "zh_CN":"匹配条件：文件类型，多个请以英文;分隔，可选值：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts"}
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// {"en":"Matching condition: Custom file type, separate by semicolon.", "zh_CN":"匹配条件：自定义文件类型，多个请以英文分号分隔。"}
	CustomFileType *string `json:"customFileType,omitempty" xml:"customFileType,omitempty"`
	// {"en":"Directory", "zh_CN":"目录"}
	Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
	// {"en":"Matching Condition: Specify URL.
	// The input parameter does not support the URI format starting with http(s)://", "zh_CN":"匹配条件：指定URL
	// 入参不支持含http(s):// 开头的URI格式"}
	SpecifyUrl *string `json:"specifyUrl,omitempty" xml:"specifyUrl,omitempty"`
	// {"en":"The matching request method, the optional values are: GET, POST, PUT, HEAD, DELETE, OPTIONS, separate by semicolons.", "zh_CN":"匹配的请求方式，可选值为：GET、POST、PUT、HEAD、DELETE、OPTIONS，多个请以英文分号分隔"}
	RequestMethod *string `json:"requestMethod,omitempty" xml:"requestMethod,omitempty"`
	// {"en":"The control direction of the http header, the optional value is cache2visitor/cache2origin/visitor2cache/origin2cache, single-select.
	// Cache2origin refers to the source direction---corresponding to the configuration item return source request;
	// Cache2visitor refers to the direction of the client back - the corresponding configuration item returns to the client response;
	// Visitor2cache refers to receiving client requests
	// Origin2cache refers to the receiving source response", "zh_CN":"http头的控制方向，可选值为cache2visitor/cache2origin/visitor2cache/origin2cache，单选。
	// cache2origin是指回源方向---对应配置项回源请求；
	// cache2visitor是指回客户端方向—对应配置项回客户端应答；
	// visitor2cache是指接收客户端请求
	// origin2cache是指接收源应答
	// 配置接收源应答方向，添加非CACHE control头，无法传递给客户端"}
	HeaderDirection *string `json:"headerDirection,omitempty" xml:"headerDirection,omitempty" require:"true"`
	// {"en":"The control type of the http header supports the addition and deletion of the http header value. The optional value is add|set|delete, which is single-selected. Corresponding to the header-name and header-value parameters.
	// 1. Add: add a header
	// 2. Set: modify the header value
	// 3. Delete: delete the header
	// Note: priority is delete > set > add", "zh_CN":"http头的控制类型，支持http头部的增删改，可选值为add|set|delete，单选。对应header-name、header-value参数
	// 1. add：表示新增一个头部，头部名称为header-name，头部值为header-value
	// 2. set：表示修改指定头部header-name的值为header-value
	// 3. delete：表示删除头部，header-name可同时配置多个
	// 注意：优先级delete>set>add。当源站有对应响应头，则按源站响应的头部响应给客户端，此处新增的无效。"}
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// {"en":"Http header regular match, optional value: true / false.
	// True: indicates that the value of the header-name is handled as a regular match.
	// False: indicates that the value of the header-name is processed according to the actual parameters, and no regular match is made.
	// Do not pass the default is false", "zh_CN":"http头正则匹配，可选值：true/false。
	// true：表示对header-name的值按正则匹配方式处理
	// false:表示对header-name的值按实际入参处理，不做正则匹配。
	// 不传默认是false"}
	AllowRegexp *string `json:"allowRegexp,omitempty" xml:"allowRegexp,omitempty"`
	// {"en":"Http header name, add or modify the http header, only one is allowed; delete the http header to allow multiple entries, separated by a semicolon ';'.
	// Note: The operation of the special http header is limited, and the http header and operation type of the operation are allowed.
	// This item is required and cannot be empty
	// When the action is add: indicates that the header-name header is added.
	// When the action is set: modify the header-name header
	// When the action is delete: delete the header-name header", "zh_CN":"http头名称，新增或修改http头，只允许输入一个；删除http头允许输入多个，以分号“;”隔开。
	// 1.当action为add：表示新增这个header-name头部
	// 2.当action为set：修改这个header-name头部的值
	// 3.当action为delete：删除这个header-name头部
	//
	// 注意：对特殊http头的操作是受限的，允许操作的http头及操作类型请参看【概览】-【附件2： header操作】"}
	HeaderName *string `json:"headerName,omitempty" xml:"headerName,omitempty"`
	// {"en":"The value corresponding to the HTTP header field, for example: mytest.example.com
	// Note:
	// 1. When the action is add or set, the input parameter must be passed a value
	// 2. When the action is delete, the input parameter is not passed
	// Support to get the value of specified variable by keyword, such as client IP, including:
	// Key words: meaning
	// #timestamp: current time, timestamp as 1559124945
	// #request-host: host in the request header
	// #request-url: request url, which contains the full path of the protocol domain name, etc., such as http://aaa.aa.com/a.html
	// #request-uri: request uri, relative path format, such as /index.html
	// #origin- IP: return source IP
	// #cache-ip: edge node IP
	// #server-ip: external service IP
	// #client-ip: client IP, or visitor IP
	// #response-header{XXX} : get the value in the response header, such as #response-header{etag}, get the etag value in response-header
	// #header{XXX} : to get the value in the HTTP header of the request, such as #header{user-agent}, is to get the user-agent value in the header
	// #cookie{XXX} : get the value in the cookie, such as #cookie{account}, is to get the value of the account set in the cookie", "zh_CN":"http头域对应的值，例如：mytest.example.com
	// 注意：
	// 1. 当action为add或set时，该入参必须传值
	// 2. 当action为delete时，该入参不用传
	// 支持通过关键字获取指定变量值，如客户端ip，包含如下：
	// 关键字：含义
	// #timestamp：当前时间，时间戳如1559124945
	// #request-host：请求头中的HOST
	// #request-url：请求url，包含协议域名等的全路径，如http://aaa.aa.com/a.html
	// #request-uri：请求uri，相对路径格式，如/index.html
	// #origin-ip：回源IP
	// #cache-ip：边缘节点IP
	// #server-ip：对外服务IP
	// #client-ip：客户端IP，即访客IP
	// #response-header{xxx}：获取响应头中的值，如#response-header{etag}，获取response-header中的etag值
	// #header{xxx}：获取请求的http header中的值，如#header{User-Agent}，是获取header中的User-Agent值
	// #cookie{xxx}：获取cookie中的值，如#cookie{account}，是获取cookie中设置的account的值  "}
	HeaderValue *string `json:"headerValue,omitempty" xml:"headerValue,omitempty"`
	// {"en":"Match request header, header values support regular, header and header values separated by Spaces, e.g. : Range bytes=[0-9]{9,}", "zh_CN":"2匹配请求头，头部值支持正则，头和头部值用空格隔开，如：Range bytes=[0-9]{9,}"}
	RequestHeader *string `json:"requestHeader,omitempty" xml:"requestHeader,omitempty"`
	// {"en":"Indicates the priority of execution order for multiple sets of configurations. A higher number indicates higher priority. If no parameters are passed, the default value is 10 and cannot be cleared.", "zh_CN":"表示客户多组配置的优先执行顺序。数字越大，优先级越高。 不传参默认为10，不可清空"}
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// {"en":"Exception file type.", "zh_CN":"例外的文件类型	"}
	ExceptFileType *string `json:"exceptFileType,omitempty" xml:"exceptFileType,omitempty"`
	// {"en":"Exception directory.", "zh_CN":"例外的目录"}
	ExceptDirectory *string `json:"exceptDirectory,omitempty" xml:"exceptDirectory,omitempty"`
	// {"en":"Exception request method.", "zh_CN":"例外的请求方式"}
	ExceptRequestMethod *string `json:"exceptRequestMethod,omitempty" xml:"exceptRequestMethod,omitempty"`
	// {"en":"Exception request header.", "zh_CN":"例外的请求头"}
	ExceptRequestHeader *string `json:"exceptRequestHeader,omitempty" xml:"exceptRequestHeader,omitempty"`
}

func (s QueryDomainForTerraformResponseDataHeaderModifyRules) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataHeaderModifyRules) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetDataId(v int64) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.DataId = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetPathPattern(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetExceptPathPattern(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.ExceptPathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetCustomPattern(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.CustomPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetFileType(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.FileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetCustomFileType(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.CustomFileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetDirectory(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.Directory = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetSpecifyUrl(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.SpecifyUrl = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetRequestMethod(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.RequestMethod = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetHeaderDirection(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.HeaderDirection = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetAction(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.Action = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetAllowRegexp(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.AllowRegexp = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetHeaderName(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.HeaderName = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetHeaderValue(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.HeaderValue = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetRequestHeader(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.RequestHeader = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetPriority(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.Priority = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetExceptFileType(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.ExceptFileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetExceptDirectory(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.ExceptDirectory = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetExceptRequestMethod(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.ExceptRequestMethod = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataHeaderModifyRules) SetExceptRequestHeader(v string) *QueryDomainForTerraformResponseDataHeaderModifyRules {
	s.ExceptRequestHeader = &v
	return s
}

type QueryDomainForTerraformResponseDataRewriteRuleSettings struct {
	// {"en":"Add a grid type identifier to indicate a specific group configuration when the client has multiple groups of configurations.", "zh_CN":"添加grid类型标识，表示客户多组配置时，具体某组配置；data-id重复，已入参同个id最后一组为准生效
	// data-id可以通过查询接口获取。
	// 注意：添加grid类型标识：data-id，每一组配置对应一个data-id：
	// a、如果客户有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；
	// b、如果客户入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；
	// c、如果客户入参都没有传data-id,表示用本次的配置全量覆盖原先配置；
	// d、如果客户入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置。（c、d内容和当前方案实现一致）；
	// e、一个gird标签下的入参不能为空，如果，没有具体的配置项，则data-id必填，且值为实际存在的data-id,表示清空这个data-id对应配置项的值；"}
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// {"en":"The url matching mode supports fuzzy regularization. If all matches, the input parameters can be configured as: *", "zh_CN":"url匹配模式，支持正则，客户入参参考：.*
	// 对于匹配到的URL进行内容重定向"}
	PathPattern *string `json:"pathPattern,omitempty" xml:"pathPattern,omitempty"`
	// {"en":"Matching conditions: specify common types, optional values are all or homepage 1. all: all files 2. homepage: home page", "zh_CN":"匹配条件：指定常用类型，可选值为all或homepage 1. all：全部文件 2. homepage：首页"}
	CustomPattern *string `json:"customPattern,omitempty" xml:"customPattern,omitempty"`
	// {"en":"directory", "zh_CN":"目录"}
	Directory *string `json:"directory,omitempty" xml:"directory,omitempty"`
	// {"en":"gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts", "zh_CN":"匹配条件：文件类型，多个请以英文;分隔，可选值：gif png bmp jpeg jpg html htm shtml mp3 wma flv mp4 wmv zip exe rar css txt ico js swf m3u8 xml f4m bootstarp ts"}
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// {"en":"Matching condition: Custom file type, please separate them by semicolon.", "zh_CN":"匹配条件：自定义文件类型，多个请以英文;分隔。"}
	CustomFileType *string `json:"customFileType,omitempty" xml:"customFileType,omitempty"`
	// {"en":"Exceptional url matching mode, except for certain URLs: such as abc.jpg, no content redirection
	// Customer reference: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做内容重定向
	// 客户入参参考：^https?://[^/]+/.*\.m3u8"}
	ExceptPathPattern *string `json:"exceptPathPattern,omitempty" xml:"exceptPathPattern,omitempty"`
	// {"en":"Ignore case, the optional value is true or false, true means to ignore case; false means not to ignore case;
	// When adding a new configuration item, the default is not true.
	// If the client passes a null value: such as <ignore-letter-case></ignore-letter-case>, the configuration is cleared.", "zh_CN":"忽略大小写，可选值为true或false，true表示忽略大小写；false表示不忽略大小写；
	// 新增配置项时，不传默认为 true
	// 如果客户传了空值：如<ignore-letter-case></ignore-letter-case>，则表示清空配置"}
	IgnoreLetterCase *string `json:"ignoreLetterCase,omitempty" xml:"ignoreLetterCase,omitempty"`
	// {"en":"Rewrite the location where the content is generated. The input value is: Cache indicates the node;
	// Other input formats are not supported at this time", "zh_CN":"改写内容的生成位置。可输入值为：Cache表示节点；
	// 暂不支持其他入参格式"}
	PublishType *string `json:"publishType,omitempty" xml:"publishType,omitempty" require:"true"`
	// {"en":"Indicates the priority execution order of multiple sets of redirected content by the customer. The higher the number, the higher the priority.
	// When adding a new configuration item, the default is 10", "zh_CN":"表示客户多组重定向内容的优先执行顺序。数字越大，优先级越高。
	// 新增配置项时，不传默认为 10"}
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// {"en":"Configuration item: old url
	// Indicates the protocol mode before rewriting (that is, the object that needs to be rewritten), such as: ^https://([^/]+/.*)", "zh_CN":"配置项：旧url
	// 表示改写前的协议方式（即需要改写的对象），如：^https://([^/]+/.*)
	// 如果是回源协议改写，则表示客户请求的原始url，配套的参数after-value，表示客户请求需要转换的回源请求。"}
	BeforeValue *string `json:"beforeValue,omitempty" xml:"beforeValue,omitempty" require:"true"`
	// {"en":"Configuration item: new url
	// Indicates the protocol method after rewriting, such as: http://$1", "zh_CN":"配置项：新url
	// 表示改写后的协议方式，如：http://$1
	// 如果请求重定向带状态码则参考入参：301:https://$1
	// 注：如果url含域名，则域名需要是本身。"}
	AfterValue *string `json:"afterValue,omitempty" xml:"afterValue,omitempty" require:"true"`
	// {"en":"Redirection type; support for input:
	// before: before the anti-theft chain
	// after: after the anti-theft chain", "zh_CN":"重定向类型；支持入参：
	// before：防盗链之前
	// after：防盗链之后"}
	RewriteType *string `json:"rewriteType,omitempty" xml:"rewriteType,omitempty" require:"true"`
	// {"en":"Matching condition: Request header", "zh_CN":"匹配条件：请求头"}
	RequestHeader *string `json:"requestHeader,omitempty" xml:"requestHeader,omitempty"`
	// {"en":"Matching condition: Exception request header", "zh_CN":"匹配条件：例外的请求头"}
	ExceptionRequestHeader *string `json:"exceptionRequestHeader,omitempty" xml:"exceptionRequestHeader,omitempty"`
}

func (s QueryDomainForTerraformResponseDataRewriteRuleSettings) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainForTerraformResponseDataRewriteRuleSettings) GoString() string {
	return s.String()
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetDataId(v int64) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.DataId = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetPathPattern(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.PathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetCustomPattern(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.CustomPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetDirectory(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.Directory = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetFileType(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.FileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetCustomFileType(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.CustomFileType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetExceptPathPattern(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.ExceptPathPattern = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetIgnoreLetterCase(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.IgnoreLetterCase = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetPublishType(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.PublishType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetPriority(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.Priority = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetBeforeValue(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.BeforeValue = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetAfterValue(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.AfterValue = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetRewriteType(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.RewriteType = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetRequestHeader(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.RequestHeader = &v
	return s
}

func (s *QueryDomainForTerraformResponseDataRewriteRuleSettings) SetExceptionRequestHeader(v string) *QueryDomainForTerraformResponseDataRewriteRuleSettings {
	s.ExceptionRequestHeader = &v
	return s
}

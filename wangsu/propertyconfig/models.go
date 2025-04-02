package propertyconfig

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type GetASchemaForARuleFormatRequest struct {
}

func (s GetASchemaForARuleFormatRequest) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatRequest) GoString() string {
  return s.String()
}

type GetASchemaForARuleFormatRequestHeader struct {
}

func (s GetASchemaForARuleFormatRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatRequestHeader) GoString() string {
  return s.String()
}

type GetASchemaForARuleFormatPaths struct {
}

func (s GetASchemaForARuleFormatPaths) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatPaths) GoString() string {
  return s.String()
}

type GetASchemaForARuleFormatParameters struct {
}

func (s GetASchemaForARuleFormatParameters) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatParameters) GoString() string {
  return s.String()
}

type GetASchemaForARuleFormatResponse struct {
  // {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Response data.","zh_CN":"接口响应数据"}
  Data *GetASchemaForARuleFormatResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
  // {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s GetASchemaForARuleFormatResponse) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatResponse) GoString() string {
  return s.String()
}

func (s *GetASchemaForARuleFormatResponse) SetCode(v string) *GetASchemaForARuleFormatResponse {
  s.Code = &v
  return s
}

func (s *GetASchemaForARuleFormatResponse) SetData(v *GetASchemaForARuleFormatResponseData) *GetASchemaForARuleFormatResponse {
  s.Data = v
  return s
}

func (s *GetASchemaForARuleFormatResponse) SetMessage(v string) *GetASchemaForARuleFormatResponse {
  s.Message = &v
  return s
}

type GetASchemaForARuleFormatResponseData struct {
  // {"en":"property configuration json schema.","zh_CN":"property配置的json schema"}
  Schema *string `json:"schema,omitempty" xml:"schema,omitempty" require:"true"`
}

func (s GetASchemaForARuleFormatResponseData) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatResponseData) GoString() string {
  return s.String()
}

func (s *GetASchemaForARuleFormatResponseData) SetSchema(v string) *GetASchemaForARuleFormatResponseData {
  s.Schema = &v
  return s
}

type GetASchemaForARuleFormatResponseHeader struct {
}

func (s GetASchemaForARuleFormatResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetASchemaForARuleFormatResponseHeader) GoString() string {
  return s.String()
}




type QueryHWAntiHotlinkingConfigRequest struct {
}

func (s QueryHWAntiHotlinkingConfigRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigRequest) GoString() string {
  return s.String()
}

type QueryHWAntiHotlinkingConfigResponse struct {
  // {"en":"Anti-theft chain configuration
  // note:
  // 1. When you need to cancel the anti-theft chain configuration settings, you can pass in the empty node <cache-time-behaviors></cache-time-behaviors>.
  // 2. When it is necessary to set the anti-theft chain configuration, this item is required.", "zh_CN":"防盗链配置
  // 注意：
  // 1. 需要取消防盗链配置设置时，可以传入空节点<cache-time-behaviors></cache-time-behaviors>。
  // 2. 表示需要设置防盗链配置时，此项必填"}
  HuaweiVisitControlRules []*QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules `json:"huawei-visit-control-rules,omitempty" xml:"huawei-visit-control-rules,omitempty" require:"true" type:"Repeated"`
}

func (s QueryHWAntiHotlinkingConfigResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigResponse) GoString() string {
  return s.String()
}

func (s *QueryHWAntiHotlinkingConfigResponse) SetHuaweiVisitControlRules(v []*QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) *QueryHWAntiHotlinkingConfigResponse {
  s.HuaweiVisitControlRules = v
  return s
}

type QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules struct     {
  // {"en":"When configuring multiple configuration sets, the specific configuration set's ID. The data-id can be obtained through the query interface. Note: a. If data-id is provided, it indicates the modification of a specific Configuration Item in one of the configuration sets. No modification is needed for other configuration sets. b. If multiple configuration sets are provided as input, and some have data-id while others do not, then those with data-id represent modifications to specific configuration sets, whereas those without data-id represent new configurations added on top of existing ones. c. If none of the inputs have data-id, it means the current configuration completely overrides the previous configuration. d. If no configuration parameters are provided and only the domain and secondary tag are transmitted, it indicates clearing all configurations corresponding to the domain's secondary service for this interface. e. If a configuration set has no specific Configuration Item, then data-id is required with an actual existing data-id value, indicating the clearing of the Configuration Item corresponding to this data-id. A configuration set with no specific Configuration Item and no data-id is not allowed.", "zh_CN":"配置多组配置时，具体某组配置的id。data-id可以通过查询接口获取。 注意： a、如果有传data-id，说明指定修改其中一组配置项内容，不需求修改其他组配置内容不需要入参；  b、如果入参多组配置，其中有些组配置有传data-id，有些没有传，则有传data-id的表示修改具体某组配置，没有传data-id的表示在原来基础上新增一组配置；  c、如果入参都没有传data-id,表示用本次的配置全量覆盖原先配置；  d、如果入参没有传任何配置项参数，只传了域名和二级标签，表示清空这个接口对应域名二级服务所有配置；  e、如果一组配置没有具体的配置项，则data-id必填，且值为实际存在的data-id，表示清空这个data-id对应配置项的值；不允许一组配置没有具体的配置项也没有data-id。"}
  DataId *int64 `json:"data-id,omitempty" xml:"data-id,omitempty"`
  // {"en":"The url matching mode supports regularization. If all matches, the input parameters can be configured as: .*", "zh_CN":"url匹配模式，支持正则，如果是全部匹配，入参可以配置为：.*"}
  PathPattern *string `json:"path-pattern,omitempty" xml:"path-pattern,omitempty"`
  // {"en":"Exceptional url matching mode, except for some URLs: such as abc.jpg, do not do anti-theft chain function
  // E.g: ^https?://[^/]+/.*\.m3u8", "zh_CN":"例外的url匹配模式，某些URL除外：如abc.jpg，不做防盗链功能
  // 客户入参参考：^https?://[^/]+/.*\.m3u8"}
  ExceptPathPattern *string `json:"except-path-pattern,omitempty" xml:"except-path-pattern,omitempty"`
  // {"en":"Prohibited IP segment
  // Input parameter limit reference interface limit
  // Forbidden IP and exceptional IP cannot be configured at the same time", "zh_CN":"禁止的IP段
  // 支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2
  // 禁止的IP和例外的IP，只能一个有值"}
  ForbiddenIps *string `json:"forbidden-ips,omitempty" xml:"forbidden-ips,omitempty"`
  // {"en":"The exception IP segment supports input IP or IP segment, and the IP segments are separated by a semicolon (;), such as 1.1.1.0/24; 2.2.2.2, some IP exceptions, no anti-theft chain", "zh_CN":"例外的IP段，支持输入IP或IP段，IP段之间用分号(;)隔开，如1.1.1.0/24;2.2.2.2，某些IP例外，不做防盗链"}
  AllowedIps *string `json:"allowed-ips,omitempty" xml:"allowed-ips,omitempty"`
  // {"en":"Forbidden Method", "zh_CN":"禁止的请求方法,组合入参：allowed-method判断，如果入参有allowed-method内容，则不允许配置forbidden-method。一个域名只能配置一种IP性质配置。 客户实际配置生效的是禁止的还是例外的，由最后一次入参的值决定，接口限制只能入参forbidden-method或allowed-method。如原先客户配置的是禁止的请求方法，本次入参是例外的请求方法，实际生效的是例外的请求方法 多个值用分号隔开 客户入参参考：get;post"}
  ForbiddenMethod *string `json:"forbidden-method,omitempty" xml:"forbidden-method,omitempty"`
  // {"en":"Allowed Method", "zh_CN":"例外的请求方法,组合入参：allowed-method判断，如果入参有allowed-imethod内容，则不允许配置forbidden-imethod。一个域名只能配置一种IP性质配置。 客户实际配置生效的是禁止的还是例外的，由最后一次入参的值决定，接口限制只能入参forbidden-method或allowed-method。如原先客户配置的是禁止的请求方法，本次入参是例外的请求方法，实际生效的是例外的请求方法 多个值用分号隔开 客户入参参考：get"}
  AllowedMethod *string `json:"allowed-method,omitempty" xml:"allowed-method,omitempty"`
  // {"en":"Decrypt Algorithm", "zh_CN":"解密算法，入参支持：空|aes-base64|aes-base64-level。 入参只能选择一种算法， 支持不传：不传就默认客户不用这个防盗链算法 如果传：aes-base64，表示点播算法； 如果传aes-base64-level，表示直播算法"}
  DecryptAlgorithm *string `json:"decrypt-algorithm,omitempty" xml:"decrypt-algorithm,omitempty"`
  // {"en":"Decrypt Key", "zh_CN":"秘钥集合，如果解密算法的入参不为空，则组内容不能为空。如果有多组秘钥信息，需要输入多组。一组内容包括：解密秘钥和秘钥过期时间 示例： <decrypt-key>-----一组秘钥信息    <secret-key>D915581AA2EF37B4</secret-key>    <expiry-time>-1</expiry-time> </decrypt-key> <decrypt-key>---一组秘钥信息    <secret-key>D915581AA2EF37B4</secret-key>    <expiry-time>20180731100000</expiry-time> </decrypt-key>"}
  DecryptKey []*QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey `json:"decrypt-key,omitempty" xml:"decrypt-key,omitempty" type:"Repeated"`
  // {"en":"Cipher Expiry Time", "zh_CN":"防盗链过期时间：精确到秒，例如：5分钟过期，则入参为300。如果不传则是0，表示马上过期，"}
  CipherExpiryTime *string `json:"cipher-expiry-time,omitempty" xml:"cipher-expiry-time,omitempty"`
  // {"en":"Authorize To Origin Rules", "zh_CN":"是否回源鉴权，支持不传，不传为空，如果传： true：表示要回源鉴权 false:不需要回源鉴权"}
  AuthorizeToOriginRules *string `json:"authorize-to-origin-rules,omitempty" xml:"authorize-to-origin-rules,omitempty"`
  // {"en":"Authorize Method", "zh_CN":"回源鉴权方式。支持入参：video、live。"}
  AuthorizeMethod *string `json:"authorize-method,omitempty" xml:"authorize-method,omitempty"`
  // {"en":"Authorize Url", "zh_CN":"回源加密串取uri的url。支持正则输入。 入参参考： 点播:^https?://[^/]+((/?[^/]+/)+).*($|\?) 直播：^https?://([^/]+(/?[^/]+/)+[^\.]+).*($|\?)"}
  AuthorizeUrl *string `json:"authorize-url,omitempty" xml:"authorize-url,omitempty"`
  // {"en":"Authorize Pattern", "zh_CN":"回源加密串取uri的字段，支持正则输入，例如：$1"}
  AuthorizePattern *string `json:"authorize-pattern,omitempty" xml:"authorize-pattern,omitempty"`
  // {"en":"Authorize Key", "zh_CN":"回源鉴权秘钥，入参参考：huawei"}
  AuthorizeKey *string `json:"authorize-key,omitempty" xml:"authorize-key,omitempty"`
  // {"en":"Authorize Cdn", "zh_CN":"回源参数hw_cdn的值：代表配置回源鉴权CDN厂家的标识。参考入参： 如果是解密算法是aes-base64-level，则回源参数hw_cdn的值，建议输入：ws-hw"}
  AuthorizeCdn *string `json:"authorize-cdn,omitempty" xml:"authorize-cdn,omitempty"`
}

func (s QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) GoString() string {
  return s.String()
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetDataId(v int64) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.DataId = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetPathPattern(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.PathPattern = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetExceptPathPattern(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.ExceptPathPattern = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetForbiddenIps(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.ForbiddenIps = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAllowedIps(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AllowedIps = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetForbiddenMethod(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.ForbiddenMethod = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAllowedMethod(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AllowedMethod = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetDecryptAlgorithm(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.DecryptAlgorithm = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetDecryptKey(v []*QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.DecryptKey = v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetCipherExpiryTime(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.CipherExpiryTime = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAuthorizeToOriginRules(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AuthorizeToOriginRules = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAuthorizeMethod(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AuthorizeMethod = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAuthorizeUrl(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AuthorizeUrl = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAuthorizePattern(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AuthorizePattern = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAuthorizeKey(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AuthorizeKey = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules) SetAuthorizeCdn(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules {
  s.AuthorizeCdn = &v
  return s
}

type QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey struct     {
  // {"en":"Secret Key", "zh_CN":"解密秘钥内容"}
  SecretKey *string `json:"Secret Key,omitempty" xml:"Secret Key,omitempty"`
  // {"en":"Secret Key", "zh_CN":"解密秘钥对应的过期时间，-1表示永不过期，入参格式精确到秒，例如：20180731100000"}
  ExpiryTime *string `json:"Expiry Time,omitempty" xml:"Expiry Time,omitempty"`
}

func (s QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey) GoString() string {
  return s.String()
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey) SetSecretKey(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey {
  s.SecretKey = &v
  return s
}

func (s *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey) SetExpiryTime(v string) *QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey {
  s.ExpiryTime = &v
  return s
}

type QueryHWAntiHotlinkingConfigPaths struct {
  // {"en":"The domain name for the acceleration domain.", "zh_CN":"加速域名。"}
  Domain *string `json:"domain,omitempty" xml:"domain,omitempty" require:"true"`
}

func (s QueryHWAntiHotlinkingConfigPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigPaths) GoString() string {
  return s.String()
}

func (s *QueryHWAntiHotlinkingConfigPaths) SetDomain(v string) *QueryHWAntiHotlinkingConfigPaths {
  s.Domain = &v
  return s
}

type QueryHWAntiHotlinkingConfigParameters struct {
}

func (s QueryHWAntiHotlinkingConfigParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigParameters) GoString() string {
  return s.String()
}

type QueryHWAntiHotlinkingConfigRequestHeader struct {
}

func (s QueryHWAntiHotlinkingConfigRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigRequestHeader) GoString() string {
  return s.String()
}

type QueryHWAntiHotlinkingConfigResponseHeader struct {
}

func (s QueryHWAntiHotlinkingConfigResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryHWAntiHotlinkingConfigResponseHeader) GoString() string {
  return s.String()
}





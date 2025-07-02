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

type QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRules struct {
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

type QueryHWAntiHotlinkingConfigResponseHuaweiVisitControlRulesDecryptKey struct {
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

type CreateDeploymentTaskForTerraformRequest struct {
	// {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// {"en":"Indicates whether to deploy to staging or production. Enum: staging,production","zh_CN":"指定部署任务的目标环境，即演练或生产环境。取值范围: staging,production"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	// {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties to the staging or production environments.","zh_CN":"指部署任务所要执行的操作，可以包括项目的部署或卸载操作。"}
	Actions []*CreateDeploymentTaskForTerraformRequestActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s CreateDeploymentTaskForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTaskForTerraformRequest) SetDeploymentName(v string) *CreateDeploymentTaskForTerraformRequest {
	s.DeploymentName = &v
	return s
}

func (s *CreateDeploymentTaskForTerraformRequest) SetTarget(v string) *CreateDeploymentTaskForTerraformRequest {
	s.Target = &v
	return s
}

func (s *CreateDeploymentTaskForTerraformRequest) SetActions(v []*CreateDeploymentTaskForTerraformRequestActions) *CreateDeploymentTaskForTerraformRequest {
	s.Actions = v
	return s
}

type CreateDeploymentTaskForTerraformRequestActions struct {
	// {"en":"Describe an action to take. You can deploy a property, remove a property. Enum: deploy_property,remove_property","zh_CN":"指定操作类型，包括部署项目、卸载项目。取值范围: deploy_property,remove_property"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的项目ID。"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Indicates the version of the property to deploy or remove.","zh_CN":"指定要部署或卸载项目的版本。"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s CreateDeploymentTaskForTerraformRequestActions) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformRequestActions) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTaskForTerraformRequestActions) SetAction(v string) *CreateDeploymentTaskForTerraformRequestActions {
	s.Action = &v
	return s
}

func (s *CreateDeploymentTaskForTerraformRequestActions) SetPropertyId(v int) *CreateDeploymentTaskForTerraformRequestActions {
	s.PropertyId = &v
	return s
}

func (s *CreateDeploymentTaskForTerraformRequestActions) SetVersion(v int) *CreateDeploymentTaskForTerraformRequestActions {
	s.Version = &v
	return s
}

type CreateDeploymentTaskForTerraformRequestHeader struct {
}

func (s CreateDeploymentTaskForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformRequestHeader) GoString() string {
	return s.String()
}

type CreateDeploymentTaskForTerraformPaths struct {
}

func (s CreateDeploymentTaskForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformPaths) GoString() string {
	return s.String()
}

type CreateDeploymentTaskForTerraformParameters struct {
}

func (s CreateDeploymentTaskForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformParameters) GoString() string {
	return s.String()
}

type CreateDeploymentTaskForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *CreateDeploymentTaskForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreateDeploymentTaskForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTaskForTerraformResponse) SetCode(v string) *CreateDeploymentTaskForTerraformResponse {
	s.Code = &v
	return s
}

func (s *CreateDeploymentTaskForTerraformResponse) SetMessage(v string) *CreateDeploymentTaskForTerraformResponse {
	s.Message = &v
	return s
}

func (s *CreateDeploymentTaskForTerraformResponse) SetData(v *CreateDeploymentTaskForTerraformResponseData) *CreateDeploymentTaskForTerraformResponse {
	s.Data = v
	return s
}

type CreateDeploymentTaskForTerraformResponseData struct {
	// {"en":"ID of the deployment task.","zh_CN":"部署任务标识"}
	DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
}

func (s CreateDeploymentTaskForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTaskForTerraformResponseData) SetDeploymentId(v int) *CreateDeploymentTaskForTerraformResponseData {
	s.DeploymentId = &v
	return s
}

type CreateDeploymentTaskForTerraformResponseHeader struct {
}

func (s CreateDeploymentTaskForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentTaskForTerraformResponseHeader) GoString() string {
	return s.String()
}

type CreatePropertyForTerraformRequest struct {
	// {"en":"Product Service Type related to your contract.Optional values include wsa, wsa-https..","zh_CN":"产品服务类型。可选值: wsa,wsa-https。请根据您的合同产品服务类型填写。"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"The name of the property. The length must not exceed 256 characters.","zh_CN":"项目的名称。长度不超过256个字符。"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
	PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
	// {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
	VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
	// {"en":"hostnames","zh_CN":"域名列表"}
	Hostnames []*CreatePropertyForTerraformRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
	// {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
	Origins *string `json:"origins,omitempty" xml:"origins,omitempty"`
	// {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
	// {"en":"Rules","zh_CN":"规则配置，详情查看schema接口。"}
	Rules *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s CreatePropertyForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequest) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformRequest) SetServiceType(v string) *CreatePropertyForTerraformRequest {
	s.ServiceType = &v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetPropertyName(v string) *CreatePropertyForTerraformRequest {
	s.PropertyName = &v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetPropertyComment(v string) *CreatePropertyForTerraformRequest {
	s.PropertyComment = &v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetVersionComment(v string) *CreatePropertyForTerraformRequest {
	s.VersionComment = &v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetHostnames(v []*CreatePropertyForTerraformRequestHostnames) *CreatePropertyForTerraformRequest {
	s.Hostnames = v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetOrigins(v string) *CreatePropertyForTerraformRequest {
	s.Origins = &v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetVariables(v string) *CreatePropertyForTerraformRequest {
	s.Variables = &v
	return s
}

func (s *CreatePropertyForTerraformRequest) SetRules(v string) *CreatePropertyForTerraformRequest {
	s.Rules = &v
	return s
}

type CreatePropertyForTerraformRequestHostnames struct {
	// {"en":"default origin","zh_CN":"默认源站"}
	DefaultOrigin *CreatePropertyForTerraformRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
	// {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
	// {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
	Certificates []*CreatePropertyForTerraformRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	// {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
	EdgeHostname *CreatePropertyForTerraformRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s CreatePropertyForTerraformRequestHostnames) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnames) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnames) SetDefaultOrigin(v *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) *CreatePropertyForTerraformRequestHostnames {
	s.DefaultOrigin = v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnames) SetHostname(v string) *CreatePropertyForTerraformRequestHostnames {
	s.Hostname = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnames) SetCertificates(v []*CreatePropertyForTerraformRequestHostnamesCertificates) *CreatePropertyForTerraformRequestHostnames {
	s.Certificates = v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnames) SetEdgeHostname(v *CreatePropertyForTerraformRequestHostnamesEdgeHostname) *CreatePropertyForTerraformRequestHostnames {
	s.EdgeHostname = v
	return s
}

type CreatePropertyForTerraformRequestHostnamesDefaultOrigin struct {
	// {"en":"origin servers","zh_CN":"源站服务器"}
	Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
	// {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
	IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
	// {"en":"http port","zh_CN":"http端口"}
	HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
	// {"en":"origin host","zh_CN":"回源HOST"}
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// {"en":"https port","zh_CN":"https端口"}
	HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
}

func (s CreatePropertyForTerraformRequestHostnamesDefaultOrigin) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnamesDefaultOrigin) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetServers(v []*string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.Servers = v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetIpVersion(v string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.IpVersion = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpPort(v int) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.HttpPort = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHost(v string) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.Host = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *CreatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.HttpsPort = &v
	return s
}

type CreatePropertyForTerraformRequestHostnamesCertificates struct {
	// {"en":"Certificate ID","zh_CN":"证书ID"}
	CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificate usage. data range: default_sni, ssl_bk, gm_sm2_enc, gm_sm2_sign, client_mtls","zh_CN":"证书用途。取值范围：default_sni, dual_sni, gm_sm2_enc, gm_sm2_sign, client_mtls"}
	CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s CreatePropertyForTerraformRequestHostnamesCertificates) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnamesCertificates) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnamesCertificates) SetCertificateId(v int) *CreatePropertyForTerraformRequestHostnamesCertificates {
	s.CertificateId = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesCertificates) SetCertificateUsage(v string) *CreatePropertyForTerraformRequestHostnamesCertificates {
	s.CertificateUsage = &v
	return s
}

type CreatePropertyForTerraformRequestHostnamesEdgeHostname struct {
	// {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
	EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
	// {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
	EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
}

func (s CreatePropertyForTerraformRequestHostnamesEdgeHostname) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHostnamesEdgeHostname) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *CreatePropertyForTerraformRequestHostnamesEdgeHostname {
	s.EdgeHostnamePrefix = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesEdgeHostname) SetComment(v string) *CreatePropertyForTerraformRequestHostnamesEdgeHostname {
	s.Comment = &v
	return s
}

func (s *CreatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *CreatePropertyForTerraformRequestHostnamesEdgeHostname {
	s.EdgeHostnameSuffix = &v
	return s
}

type CreatePropertyForTerraformRequestHeader struct {
}

func (s CreatePropertyForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformRequestHeader) GoString() string {
	return s.String()
}

type CreatePropertyForTerraformPaths struct {
}

func (s CreatePropertyForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformPaths) GoString() string {
	return s.String()
}

type CreatePropertyForTerraformParameters struct {
}

func (s CreatePropertyForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformParameters) GoString() string {
	return s.String()
}

type CreatePropertyForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *CreatePropertyForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s CreatePropertyForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformResponse) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformResponse) SetCode(v string) *CreatePropertyForTerraformResponse {
	s.Code = &v
	return s
}

func (s *CreatePropertyForTerraformResponse) SetMessage(v string) *CreatePropertyForTerraformResponse {
	s.Message = &v
	return s
}

func (s *CreatePropertyForTerraformResponse) SetData(v *CreatePropertyForTerraformResponseData) *CreatePropertyForTerraformResponse {
	s.Data = v
	return s
}

type CreatePropertyForTerraformResponseData struct {
	// {"en":"Property ID","zh_CN":"项目ID"}
	PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Property Name","zh_CN":"项目名称"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"Property Version","zh_CN":"项目版本"}
	PropertyVersion *int64 `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
}

func (s CreatePropertyForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *CreatePropertyForTerraformResponseData) SetPropertyId(v int64) *CreatePropertyForTerraformResponseData {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyForTerraformResponseData) SetPropertyName(v string) *CreatePropertyForTerraformResponseData {
	s.PropertyName = &v
	return s
}

func (s *CreatePropertyForTerraformResponseData) SetPropertyVersion(v int64) *CreatePropertyForTerraformResponseData {
	s.PropertyVersion = &v
	return s
}

type CreatePropertyForTerraformResponseHeader struct {
}

func (s CreatePropertyForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyForTerraformResponseHeader) GoString() string {
	return s.String()
}

type DeletePropertyForTerraformRequest struct {
}

func (s DeletePropertyForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePropertyForTerraformRequest) GoString() string {
	return s.String()
}

type DeletePropertyForTerraformRequestHeader struct {
}

func (s DeletePropertyForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s DeletePropertyForTerraformRequestHeader) GoString() string {
	return s.String()
}

type DeletePropertyForTerraformPaths struct {
	// {"en":"Property ID","zh_CN":"项目ID"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s DeletePropertyForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s DeletePropertyForTerraformPaths) GoString() string {
	return s.String()
}

func (s *DeletePropertyForTerraformPaths) SetPropertyId(v int) *DeletePropertyForTerraformPaths {
	s.PropertyId = &v
	return s
}

type DeletePropertyForTerraformParameters struct {
}

func (s DeletePropertyForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s DeletePropertyForTerraformParameters) GoString() string {
	return s.String()
}

type DeletePropertyForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s DeletePropertyForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePropertyForTerraformResponse) GoString() string {
	return s.String()
}

func (s *DeletePropertyForTerraformResponse) SetCode(v string) *DeletePropertyForTerraformResponse {
	s.Code = &v
	return s
}

func (s *DeletePropertyForTerraformResponse) SetMessage(v string) *DeletePropertyForTerraformResponse {
	s.Message = &v
	return s
}

type DeletePropertyForTerraformResponseHeader struct {
}

func (s DeletePropertyForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s DeletePropertyForTerraformResponseHeader) GoString() string {
	return s.String()
}

type QueryDeploymentForTerraformRequest struct {
}

func (s QueryDeploymentForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformRequest) GoString() string {
	return s.String()
}

type QueryDeploymentForTerraformRequestHeader struct {
}

func (s QueryDeploymentForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformRequestHeader) GoString() string {
	return s.String()
}

type QueryDeploymentForTerraformPaths struct {
	// {"en":"ID of the deployment task","zh_CN":"部署任务ID"}
	DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
}

func (s QueryDeploymentForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformPaths) GoString() string {
	return s.String()
}

func (s *QueryDeploymentForTerraformPaths) SetDeploymentId(v int) *QueryDeploymentForTerraformPaths {
	s.DeploymentId = &v
	return s
}

type QueryDeploymentForTerraformParameters struct {
}

func (s QueryDeploymentForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformParameters) GoString() string {
	return s.String()
}

type QueryDeploymentForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *QueryDeploymentForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryDeploymentForTerraformResponse) SetCode(v string) *QueryDeploymentForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryDeploymentForTerraformResponse) SetMessage(v string) *QueryDeploymentForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryDeploymentForTerraformResponse) SetData(v *QueryDeploymentForTerraformResponseData) *QueryDeploymentForTerraformResponse {
	s.Data = v
	return s
}

type QueryDeploymentForTerraformResponseData struct {
	// {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
	DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
	// {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty" require:"true"`
	// {"en":"Status of Deployment. Enum:PENDING,IN_PROCESS,SUCCESS,FAIL","zh_CN":"任务状态。取值范围:PENDING,IN_PROCESS,SUCCESS,FAIL"}
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"项目的部署环境。取值范围: staging,production"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	// {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
	SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
	// {"en":"An RFC3339 date indicates when the task was last updated.","zh_CN":"RFC 3339格式的日期，表示任务的最近更新时间。"}
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
	// {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
	// {"en":"This array contains all the actions related to a deployment. They can include deployment and removal of properties to the staging or production environments.","zh_CN":"指部署任务所要执行的操作，可以包括项目的部署或卸载操作。"}
	Actions []*QueryDeploymentForTerraformResponseDataActions `json:"actions,omitempty" xml:"actions,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDeploymentForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryDeploymentForTerraformResponseData) SetDeploymentId(v int) *QueryDeploymentForTerraformResponseData {
	s.DeploymentId = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetDeploymentName(v string) *QueryDeploymentForTerraformResponseData {
	s.DeploymentName = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetStatus(v string) *QueryDeploymentForTerraformResponseData {
	s.Status = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetTarget(v string) *QueryDeploymentForTerraformResponseData {
	s.Target = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetSubmissionTime(v string) *QueryDeploymentForTerraformResponseData {
	s.SubmissionTime = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetLastUpdateTime(v string) *QueryDeploymentForTerraformResponseData {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetFinishTime(v string) *QueryDeploymentForTerraformResponseData {
	s.FinishTime = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseData) SetActions(v []*QueryDeploymentForTerraformResponseDataActions) *QueryDeploymentForTerraformResponseData {
	s.Actions = v
	return s
}

type QueryDeploymentForTerraformResponseDataActions struct {
	// {"en":"Describe an action to take. You can deploy a property, remove a property. Enum: deploy_property,remove_property","zh_CN":"指定操作类型，包括部署加速项目、卸载项目。取值范围: deploy_property,remove_property"}
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	// {"en":"ID of the property to deploy or remove from the staging or production environment.","zh_CN":"指定要部署或卸载的项目ID。"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Indicates the version of the property to deploy or remove.","zh_CN":"指定要部署或卸载项目的版本。"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s QueryDeploymentForTerraformResponseDataActions) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponseDataActions) GoString() string {
	return s.String()
}

func (s *QueryDeploymentForTerraformResponseDataActions) SetAction(v string) *QueryDeploymentForTerraformResponseDataActions {
	s.Action = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseDataActions) SetPropertyId(v int) *QueryDeploymentForTerraformResponseDataActions {
	s.PropertyId = &v
	return s
}

func (s *QueryDeploymentForTerraformResponseDataActions) SetVersion(v int) *QueryDeploymentForTerraformResponseDataActions {
	s.Version = &v
	return s
}

type QueryDeploymentForTerraformResponseHeader struct {
}

func (s QueryDeploymentForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentForTerraformResponseHeader) GoString() string {
	return s.String()
}

type QueryDeploymentsForTerraformRequest struct {
}

func (s QueryDeploymentsForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformRequest) GoString() string {
	return s.String()
}

type QueryDeploymentsForTerraformRequestHeader struct {
}

func (s QueryDeploymentsForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformRequestHeader) GoString() string {
	return s.String()
}

type QueryDeploymentsForTerraformPaths struct {
}

func (s QueryDeploymentsForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformPaths) GoString() string {
	return s.String()
}

type QueryDeploymentsForTerraformParameters struct {
	// {"en":"Property ID","zh_CN":"加速项目标识"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" url:"propertyId"`
	// {"en":"Status of Deployment. Enum:PENDING,IN_PROCESS,SUCCESS,FAIL","zh_CN":"任务状态。取值范围:PENDING,IN_PROCESS,SUCCESS,FAIL"}
	Status *string `json:"status,omitempty" xml:"status,omitempty" url:"status"`
	// {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"加速项目的部署环境。取值范围: staging,production"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" url:"target"`
	// {"en":"Indicates the first item to return. The default is '0'.","zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
	Offset *int `json:"offset,omitempty" xml:"offset,omitempty" url:"offset"`
	// {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200","zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
	Limit *int `json:"limit,omitempty" xml:"limit,omitempty" url:"limit"`
	// {"en":"Order of properties to return. Enum: asc,desc Default: desc","zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty" url:"sortOrder"`
	// {"en":"Returns results in sorted order. Enum: submissionTime,lastUpdateTime Default: submissionTime","zh_CN":"返回结果的排序依据。取值范围: submissionTime,lastUpdateTime 默认值: submissionTime"}
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty" url:"sortBy"`
}

func (s QueryDeploymentsForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformParameters) GoString() string {
	return s.String()
}

func (s *QueryDeploymentsForTerraformParameters) SetPropertyId(v int) *QueryDeploymentsForTerraformParameters {
	s.PropertyId = &v
	return s
}

func (s *QueryDeploymentsForTerraformParameters) SetStatus(v string) *QueryDeploymentsForTerraformParameters {
	s.Status = &v
	return s
}

func (s *QueryDeploymentsForTerraformParameters) SetTarget(v string) *QueryDeploymentsForTerraformParameters {
	s.Target = &v
	return s
}

func (s *QueryDeploymentsForTerraformParameters) SetOffset(v int) *QueryDeploymentsForTerraformParameters {
	s.Offset = &v
	return s
}

func (s *QueryDeploymentsForTerraformParameters) SetLimit(v int) *QueryDeploymentsForTerraformParameters {
	s.Limit = &v
	return s
}

func (s *QueryDeploymentsForTerraformParameters) SetSortOrder(v string) *QueryDeploymentsForTerraformParameters {
	s.SortOrder = &v
	return s
}

func (s *QueryDeploymentsForTerraformParameters) SetSortBy(v string) *QueryDeploymentsForTerraformParameters {
	s.SortBy = &v
	return s
}

type QueryDeploymentsForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *QueryDeploymentsForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryDeploymentsForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryDeploymentsForTerraformResponse) SetCode(v string) *QueryDeploymentsForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponse) SetMessage(v string) *QueryDeploymentsForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponse) SetData(v *QueryDeploymentsForTerraformResponseData) *QueryDeploymentsForTerraformResponse {
	s.Data = v
	return s
}

type QueryDeploymentsForTerraformResponseData struct {
	// {"en":"Number of properties.","zh_CN":"部署任务的总数。"}
	Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	// {"en":"List of deployment task summaries.","zh_CN":"部署任务列表。"}
	Deployments []*QueryDeploymentsForTerraformResponseDataDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" require:"true" type:"Repeated"`
}

func (s QueryDeploymentsForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryDeploymentsForTerraformResponseData) SetCount(v int) *QueryDeploymentsForTerraformResponseData {
	s.Count = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseData) SetDeployments(v []*QueryDeploymentsForTerraformResponseDataDeployments) *QueryDeploymentsForTerraformResponseData {
	s.Deployments = v
	return s
}

type QueryDeploymentsForTerraformResponseDataDeployments struct {
	// {"en":"ID representing the deployment task","zh_CN":"部署任务的ID。"}
	DeploymentId *int `json:"deploymentId,omitempty" xml:"deploymentId,omitempty" require:"true"`
	// {"en":"Name representing the deployment task","zh_CN":"部署任务的名称。"}
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty" require:"true"`
	// {"en":"Status of Deployment. Enum:PENDING,IN_PROCESS,SUCCESS,FAIL","zh_CN":"任务状态。取值范围:PENDING,IN_PROCESS,SUCCESS,FAIL"}
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"项目的部署环境。取值范围: staging,production"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	// {"en":"An RFC 3339 format date indicates when the task was submitted.","zh_CN":"RFC 3339格式的日期，表示任务的提交时间。"}
	SubmissionTime *string `json:"submissionTime,omitempty" xml:"submissionTime,omitempty" require:"true"`
	// {"en":"An RFC3339 date indicates when the task was last updated.","zh_CN":"RFC 3339格式的日期，表示任务的最近更新时间。"}
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
	// {"en":"An RFC 3339 date indicates when the task completed.","zh_CN":"RFC 3339格式的日期，表示任务的完成时间。"}
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty" require:"true"`
}

func (s QueryDeploymentsForTerraformResponseDataDeployments) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponseDataDeployments) GoString() string {
	return s.String()
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetDeploymentId(v int) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.DeploymentId = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetDeploymentName(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.DeploymentName = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetStatus(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.Status = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetTarget(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.Target = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetSubmissionTime(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.SubmissionTime = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetLastUpdateTime(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryDeploymentsForTerraformResponseDataDeployments) SetFinishTime(v string) *QueryDeploymentsForTerraformResponseDataDeployments {
	s.FinishTime = &v
	return s
}

type QueryDeploymentsForTerraformResponseHeader struct {
}

func (s QueryDeploymentsForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryDeploymentsForTerraformResponseHeader) GoString() string {
	return s.String()
}

type QueryPropertiesForTerraformRequest struct {
}

func (s QueryPropertiesForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformRequest) GoString() string {
	return s.String()
}

type QueryPropertiesForTerraformRequestHeader struct {
}

func (s QueryPropertiesForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformRequestHeader) GoString() string {
	return s.String()
}

type QueryPropertiesForTerraformPaths struct {
}

func (s QueryPropertiesForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformPaths) GoString() string {
	return s.String()
}

type QueryPropertiesForTerraformParameters struct {
	// {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" url:"serviceType"`
	// {"en":"The value can be 'staging', or 'production' to filter the results based on where the property has been deployed.","zh_CN":"加速项目的部署环境。取值范围: staging,production"}
	Target *string `json:"target,omitempty" xml:"target,omitempty" url:"target"`
	// {"en":"Indicates the first item to return. The default is '0'.","zh_CN":"查询起始位置。默认值: 0 取值范围: >= 0"}
	Offset *int `json:"offset,omitempty" xml:"offset,omitempty" url:"offset"`
	// {"en":"Maximum number of properties to return.  Default: 100 Range: <= 200","zh_CN":"每次查询的最大条数。默认值: 100 取值范围: <= 200"}
	Limit *int `json:"limit,omitempty" xml:"limit,omitempty" url:"limit"`
	// {"en":"Order of properties to return. Enum: asc,desc Default: desc","zh_CN":"返回结果的顺序。默认按最后更新时间降序。取值范围: asc,desc 默认值: desc"}
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty" url:"sortOrder"`
	// {"en":"Returns results in sorted order. Enum: creationTime,lastUpdateTime Default: lastUpdateTime","zh_CN":"返回结果的排序依据。取值范围: creationTime,lastUpdateTime 默认值: lastUpdateTime"}
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty" url:"sortBy"`
	// {"en":"Hostname","zh_CN":"域名"}
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" url:"hostname"`
}

func (s QueryPropertiesForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformParameters) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformParameters) SetServiceType(v string) *QueryPropertiesForTerraformParameters {
	s.ServiceType = &v
	return s
}

func (s *QueryPropertiesForTerraformParameters) SetTarget(v string) *QueryPropertiesForTerraformParameters {
	s.Target = &v
	return s
}

func (s *QueryPropertiesForTerraformParameters) SetOffset(v int) *QueryPropertiesForTerraformParameters {
	s.Offset = &v
	return s
}

func (s *QueryPropertiesForTerraformParameters) SetLimit(v int) *QueryPropertiesForTerraformParameters {
	s.Limit = &v
	return s
}

func (s *QueryPropertiesForTerraformParameters) SetSortOrder(v string) *QueryPropertiesForTerraformParameters {
	s.SortOrder = &v
	return s
}

func (s *QueryPropertiesForTerraformParameters) SetSortBy(v string) *QueryPropertiesForTerraformParameters {
	s.SortBy = &v
	return s
}

func (s *QueryPropertiesForTerraformParameters) SetHostname(v string) *QueryPropertiesForTerraformParameters {
	s.Hostname = &v
	return s
}

type QueryPropertiesForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *QueryPropertiesForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertiesForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponse) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponse) SetCode(v string) *QueryPropertiesForTerraformResponse {
	s.Code = &v
	return s
}

func (s *QueryPropertiesForTerraformResponse) SetMessage(v string) *QueryPropertiesForTerraformResponse {
	s.Message = &v
	return s
}

func (s *QueryPropertiesForTerraformResponse) SetData(v *QueryPropertiesForTerraformResponseData) *QueryPropertiesForTerraformResponse {
	s.Data = v
	return s
}

type QueryPropertiesForTerraformResponseData struct {
	// {"en":"Number of properties.","zh_CN":"项目数量。"}
	Count *int `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	// {"en":"List of properties.","zh_CN":"项目列表。"}
	Properties []*QueryPropertiesForTerraformResponseDataProperties `json:"properties,omitempty" xml:"properties,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponseData) SetCount(v int) *QueryPropertiesForTerraformResponseData {
	s.Count = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseData) SetProperties(v []*QueryPropertiesForTerraformResponseDataProperties) *QueryPropertiesForTerraformResponseData {
	s.Properties = v
	return s
}

type QueryPropertiesForTerraformResponseDataProperties struct {
	// {"en":"Property ID","zh_CN":"项目标识"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Name of the property.","zh_CN":"项目的名称"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"A description of the property.","zh_CN":"项目的描述。"}
	PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
	// {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
	CreationTime *string `json:"creationTime,omitempty" xml:"creationTime,omitempty" require:"true"`
	// {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty" require:"true"`
	// {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
	LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
	// {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
	StagingVersion *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
	ProductionVersion *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
	StagingDeployingVersion *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
	ProductionDeployingVersion *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertiesForTerraformResponseDataProperties) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataProperties) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetPropertyId(v int) *QueryPropertiesForTerraformResponseDataProperties {
	s.PropertyId = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetPropertyName(v string) *QueryPropertiesForTerraformResponseDataProperties {
	s.PropertyName = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetPropertyComment(v string) *QueryPropertiesForTerraformResponseDataProperties {
	s.PropertyComment = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetServiceType(v string) *QueryPropertiesForTerraformResponseDataProperties {
	s.ServiceType = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetCreationTime(v string) *QueryPropertiesForTerraformResponseDataProperties {
	s.CreationTime = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetLastUpdateTime(v string) *QueryPropertiesForTerraformResponseDataProperties {
	s.LastUpdateTime = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetLatestVersion(v int) *QueryPropertiesForTerraformResponseDataProperties {
	s.LatestVersion = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetStagingVersion(v *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) *QueryPropertiesForTerraformResponseDataProperties {
	s.StagingVersion = v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetProductionVersion(v *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) *QueryPropertiesForTerraformResponseDataProperties {
	s.ProductionVersion = v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetStagingDeployingVersion(v *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) *QueryPropertiesForTerraformResponseDataProperties {
	s.StagingDeployingVersion = v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataProperties) SetProductionDeployingVersion(v *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) *QueryPropertiesForTerraformResponseDataProperties {
	s.ProductionDeployingVersion = v
	return s
}

type QueryPropertiesForTerraformResponseDataPropertiesStagingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesStagingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertiesForTerraformResponseDataPropertiesProductionVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesProductionVersion {
	s.Hostnames = v
	return s
}

type QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesStagingDeployingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) SetVersion(v int) *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion) SetHostnames(v []*string) *QueryPropertiesForTerraformResponseDataPropertiesProductionDeployingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertiesForTerraformResponseHeader struct {
}

func (s QueryPropertiesForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertiesForTerraformResponseHeader) GoString() string {
	return s.String()
}

type QueryPropertyConfigForTerrformRequest struct {
}

func (s QueryPropertyConfigForTerrformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformRequest) GoString() string {
	return s.String()
}

type QueryPropertyConfigForTerrformRequestHeader struct {
}

func (s QueryPropertyConfigForTerrformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformRequestHeader) GoString() string {
	return s.String()
}

type QueryPropertyConfigForTerrformPaths struct {
	// {"en":"Property ID","zh_CN":"项目ID"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformPaths) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformPaths) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformPaths) SetPropertyId(v int) *QueryPropertyConfigForTerrformPaths {
	s.PropertyId = &v
	return s
}

type QueryPropertyConfigForTerrformParameters struct {
}

func (s QueryPropertyConfigForTerrformParameters) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformParameters) GoString() string {
	return s.String()
}

type QueryPropertyConfigForTerrformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *QueryPropertyConfigForTerrformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertyConfigForTerrformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponse) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponse) SetCode(v string) *QueryPropertyConfigForTerrformResponse {
	s.Code = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponse) SetMessage(v string) *QueryPropertyConfigForTerrformResponse {
	s.Message = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponse) SetData(v *QueryPropertyConfigForTerrformResponseData) *QueryPropertyConfigForTerrformResponse {
	s.Data = v
	return s
}

type QueryPropertyConfigForTerrformResponseData struct {
	// {"en":"Property ID","zh_CN":"项目标识"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Name of the property.","zh_CN":"项目的名称"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"A description of the property.","zh_CN":"项目的描述。"}
	PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
	// {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
	PropertyCreationTime *string `json:"propertyCreationTime,omitempty" xml:"propertyCreationTime,omitempty" require:"true"`
	// {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
	PropertyLastUpdateTime *string `json:"propertyLastUpdateTime,omitempty" xml:"propertyLastUpdateTime,omitempty" require:"true"`
	// {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
	StagingVersion *QueryPropertyConfigForTerrformResponseDataStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
	ProductionVersion *QueryPropertyConfigForTerrformResponseDataProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
	StagingDeployingVersion *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
	ProductionDeployingVersion *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
	LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
	// {"en":"A property version. It must be an integer value >=1.","zh_CN":"项目的版本，必须是大于0的整数。"}
	CurrentVersion *int `json:"currentVersion,omitempty" xml:"currentVersion,omitempty" require:"true"`
	// {"en":"A description of the version.","zh_CN":"版本描述。"}
	VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty" require:"true"`
	// {"en":"Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.","zh_CN":"该项目版本是否处于冻结状态。当加速项目版本部署后即进入冻结状态，不可再更新该版本。"}
	Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
	// {"en":"RFC3339 format date indicating when the version was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
	VersionCreationTime *string `json:"versionCreationTime,omitempty" xml:"versionCreationTime,omitempty" require:"true"`
	// {"en":"RFC3339 date indicating when the version was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
	VersionLastUpdateTime *string `json:"versionLastUpdateTime,omitempty" xml:"versionLastUpdateTime,omitempty" require:"true"`
	// {"en":"hostnames","zh_CN":"域名列表"}
	Hostnames []*QueryPropertyConfigForTerrformResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
	// {"en":"Rules","zh_CN":"规则"}
	Rules *string `json:"rules,omitempty" xml:"rules,omitempty" require:"true"`
	// {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
	Origins *string `json:"origins,omitempty" xml:"origins,omitempty" require:"true"`
	// {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseData) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyId(v int) *QueryPropertyConfigForTerrformResponseData {
	s.PropertyId = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyName(v string) *QueryPropertyConfigForTerrformResponseData {
	s.PropertyName = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyComment(v string) *QueryPropertyConfigForTerrformResponseData {
	s.PropertyComment = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetServiceType(v string) *QueryPropertyConfigForTerrformResponseData {
	s.ServiceType = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyCreationTime(v string) *QueryPropertyConfigForTerrformResponseData {
	s.PropertyCreationTime = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetPropertyLastUpdateTime(v string) *QueryPropertyConfigForTerrformResponseData {
	s.PropertyLastUpdateTime = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetStagingVersion(v *QueryPropertyConfigForTerrformResponseDataStagingVersion) *QueryPropertyConfigForTerrformResponseData {
	s.StagingVersion = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetProductionVersion(v *QueryPropertyConfigForTerrformResponseDataProductionVersion) *QueryPropertyConfigForTerrformResponseData {
	s.ProductionVersion = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetStagingDeployingVersion(v *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) *QueryPropertyConfigForTerrformResponseData {
	s.StagingDeployingVersion = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetProductionDeployingVersion(v *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) *QueryPropertyConfigForTerrformResponseData {
	s.ProductionDeployingVersion = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetLatestVersion(v int) *QueryPropertyConfigForTerrformResponseData {
	s.LatestVersion = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetCurrentVersion(v int) *QueryPropertyConfigForTerrformResponseData {
	s.CurrentVersion = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVersionComment(v string) *QueryPropertyConfigForTerrformResponseData {
	s.VersionComment = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetFrozen(v bool) *QueryPropertyConfigForTerrformResponseData {
	s.Frozen = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVersionCreationTime(v string) *QueryPropertyConfigForTerrformResponseData {
	s.VersionCreationTime = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVersionLastUpdateTime(v string) *QueryPropertyConfigForTerrformResponseData {
	s.VersionLastUpdateTime = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetHostnames(v []*QueryPropertyConfigForTerrformResponseDataHostnames) *QueryPropertyConfigForTerrformResponseData {
	s.Hostnames = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetRules(v string) *QueryPropertyConfigForTerrformResponseData {
	s.Rules = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetOrigins(v string) *QueryPropertyConfigForTerrformResponseData {
	s.Origins = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseData) SetVariables(v string) *QueryPropertyConfigForTerrformResponseData {
	s.Variables = &v
	return s
}

type QueryPropertyConfigForTerrformResponseDataStagingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataStagingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataStagingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataStagingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataStagingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyConfigForTerrformResponseDataProductionVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataProductionVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataProductionVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataProductionVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataProductionVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataStagingDeployingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) SetVersion(v int) *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion) SetHostnames(v []*string) *QueryPropertyConfigForTerrformResponseDataProductionDeployingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyConfigForTerrformResponseDataHostnames struct {
	// {"en":"default origin","zh_CN":"默认源站"}
	DefaultOrigin *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" require:"true" type:"Struct"`
	// {"en":"hostname","zh_CN":"域名"}
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
	// {"en":"hostname association certificate configuration","zh_CN":"关联证书配置"}
	Certificates []*QueryPropertyConfigForTerrformResponseDataHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
	// {"en":"icp","zh_CN":"域名备案号"}
	Icp *string `json:"icp,omitempty" xml:"icp,omitempty" require:"true"`
	// {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
	EdgeHostname *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnames) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnames) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetDefaultOrigin(v *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) *QueryPropertyConfigForTerrformResponseDataHostnames {
	s.DefaultOrigin = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetHostname(v string) *QueryPropertyConfigForTerrformResponseDataHostnames {
	s.Hostname = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetCertificates(v []*QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) *QueryPropertyConfigForTerrformResponseDataHostnames {
	s.Certificates = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetIcp(v string) *QueryPropertyConfigForTerrformResponseDataHostnames {
	s.Icp = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnames) SetEdgeHostname(v *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) *QueryPropertyConfigForTerrformResponseDataHostnames {
	s.EdgeHostname = v
	return s
}

type QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin struct {
	// {"en":"origin servers","zh_CN":"源站服务器"}
	Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
	// {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
	IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" require:"true"`
	// {"en":"http port","zh_CN":"http端口"}
	HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty" require:"true"`
	// {"en":"origin host","zh_CN":"回源HOST"}
	Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
	// {"en":"https port","zh_CN":"https端口"}
	HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetServers(v []*string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.Servers = v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetIpVersion(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.IpVersion = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpPort(v int) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.HttpPort = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHost(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.Host = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpsPort(v int) *QueryPropertyConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.HttpsPort = &v
	return s
}

type QueryPropertyConfigForTerrformResponseDataHostnamesCertificates struct {
	// {"en":"Certificate ID","zh_CN":"证书ID"}
	CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificate usage. data range: default_sni, ssl_bk, gm_sm2_enc, gm_sm2_sign, client_mtls","zh_CN":"证书用途。取值范围：default_sni, dual_sni, gm_sm2_enc, gm_sm2_sign, client_mtls"}
	CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) SetCertificateId(v int) *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates {
	s.CertificateId = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates) SetCertificateUsage(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesCertificates {
	s.CertificateUsage = &v
	return s
}

type QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname struct {
	// {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
	EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
	// {"en":"dns service status.data range: inactive,active","zh_CN":"DNS服务状态。取值范围：inactive, active。备注：inactive：挂起，active：生效"}
	DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
	// {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
	EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
	// {"en":"edge-hostname","zh_CN":"调度域名"}
	EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) GoString() string {
	return s.String()
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.EdgeHostnamePrefix = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostname(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.EdgeHostname = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetDnsServiceStatus(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.DnsServiceStatus = &v
	return s
}

func (s *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *QueryPropertyConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.EdgeHostnameSuffix = &v
	return s
}

type QueryPropertyConfigForTerrformResponseHeader struct {
}

func (s QueryPropertyConfigForTerrformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyConfigForTerrformResponseHeader) GoString() string {
	return s.String()
}

type QueryPropertyVersionConfigForTerrformRequest struct {
}

func (s QueryPropertyVersionConfigForTerrformRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformRequest) GoString() string {
	return s.String()
}

type QueryPropertyVersionConfigForTerrformRequestHeader struct {
}

func (s QueryPropertyVersionConfigForTerrformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformRequestHeader) GoString() string {
	return s.String()
}

type QueryPropertyVersionConfigForTerrformPaths struct {
	// {"en":"Property ID","zh_CN":"项目ID"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Property Version","zh_CN":"项目版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformPaths) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformPaths) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformPaths) SetPropertyId(v int) *QueryPropertyVersionConfigForTerrformPaths {
	s.PropertyId = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformPaths) SetVersion(v int) *QueryPropertyVersionConfigForTerrformPaths {
	s.Version = &v
	return s
}

type QueryPropertyVersionConfigForTerrformParameters struct {
}

func (s QueryPropertyVersionConfigForTerrformParameters) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformParameters) GoString() string {
	return s.String()
}

type QueryPropertyVersionConfigForTerrformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *QueryPropertyVersionConfigForTerrformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertyVersionConfigForTerrformResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponse) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponse) SetCode(v string) *QueryPropertyVersionConfigForTerrformResponse {
	s.Code = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponse) SetMessage(v string) *QueryPropertyVersionConfigForTerrformResponse {
	s.Message = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponse) SetData(v *QueryPropertyVersionConfigForTerrformResponseData) *QueryPropertyVersionConfigForTerrformResponse {
	s.Data = v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseData struct {
	// {"en":"Property ID","zh_CN":"项目标识"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Name of the property.","zh_CN":"项目的名称"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"A description of the property.","zh_CN":"项目的描述。"}
	PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty" require:"true"`
	// {"en":"Unique identifier for the product. Optional values include wsa, wsa-https.","zh_CN":"服务类型。可选值: wsa, wsa-https"}
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty" require:"true"`
	// {"en":"RFC3339 format date indicating when the property was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
	PropertyCreationTime *string `json:"propertyCreationTime,omitempty" xml:"propertyCreationTime,omitempty" require:"true"`
	// {"en":"RFC3339 date indicating when the property was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
	PropertyLastUpdateTime *string `json:"propertyLastUpdateTime,omitempty" xml:"propertyLastUpdateTime,omitempty" require:"true"`
	// {"en":"Describes the version of the property deployed to staging.","zh_CN":"描述已部署到演练环境的项目版本。"}
	StagingVersion *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion `json:"stagingVersion,omitempty" xml:"stagingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deployed to production.","zh_CN":"描述已部署到生产环境的项目版本。"}
	ProductionVersion *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion `json:"productionVersion,omitempty" xml:"productionVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deploying to staging.","zh_CN":"描述正在部署到演练环境的项目版本。"}
	StagingDeployingVersion *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion `json:"stagingDeployingVersion,omitempty" xml:"stagingDeployingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Describes the version of the property deploying to production.","zh_CN":"描述正在部署到生产环境的项目版本。"}
	ProductionDeployingVersion *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion `json:"productionDeployingVersion,omitempty" xml:"productionDeployingVersion,omitempty" require:"true" type:"Struct"`
	// {"en":"Latest version of the property.","zh_CN":"项目的最新版本。"}
	LatestVersion *int `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" require:"true"`
	// {"en":"A property version. It must be an integer value >=1.","zh_CN":"项目的版本，必须是大于0的整数。"}
	CurrentVersion *int `json:"currentVersion,omitempty" xml:"currentVersion,omitempty" require:"true"`
	// {"en":"A description of the version.","zh_CN":"版本描述。"}
	VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty" require:"true"`
	// {"en":"Indicates whether the property version is frozen or can still be updated. A property version is frozen once it has been deployed.","zh_CN":"该项目版本是否处于冻结状态。当加速项目版本部署后即进入冻结状态，不可再更新该版本。"}
	Frozen *bool `json:"frozen,omitempty" xml:"frozen,omitempty" require:"true"`
	// {"en":"RFC3339 format date indicating when the version was created.","zh_CN":"RFC 3339格式的日期，表示创建项目的时间。"}
	VersionCreationTime *string `json:"versionCreationTime,omitempty" xml:"versionCreationTime,omitempty" require:"true"`
	// {"en":"RFC3339 date indicating when the version was last updated.","zh_CN":"RFC 3339格式的日期，表示项目的最近更新时间。"}
	VersionLastUpdateTime *string `json:"versionLastUpdateTime,omitempty" xml:"versionLastUpdateTime,omitempty" require:"true"`
	// {"en":"hostnames","zh_CN":"域名列表"}
	Hostnames []*QueryPropertyVersionConfigForTerrformResponseDataHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
	// {"en":"Rules","zh_CN":"规则"}
	Rules *string `json:"rules,omitempty" xml:"rules,omitempty" require:"true"`
	// {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
	Origins *string `json:"origins,omitempty" xml:"origins,omitempty" require:"true"`
	// {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseData) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyId(v int) *QueryPropertyVersionConfigForTerrformResponseData {
	s.PropertyId = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyName(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.PropertyName = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyComment(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.PropertyComment = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetServiceType(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.ServiceType = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyCreationTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.PropertyCreationTime = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetPropertyLastUpdateTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.PropertyLastUpdateTime = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetStagingVersion(v *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) *QueryPropertyVersionConfigForTerrformResponseData {
	s.StagingVersion = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetProductionVersion(v *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) *QueryPropertyVersionConfigForTerrformResponseData {
	s.ProductionVersion = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetStagingDeployingVersion(v *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) *QueryPropertyVersionConfigForTerrformResponseData {
	s.StagingDeployingVersion = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetProductionDeployingVersion(v *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) *QueryPropertyVersionConfigForTerrformResponseData {
	s.ProductionDeployingVersion = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetLatestVersion(v int) *QueryPropertyVersionConfigForTerrformResponseData {
	s.LatestVersion = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetCurrentVersion(v int) *QueryPropertyVersionConfigForTerrformResponseData {
	s.CurrentVersion = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVersionComment(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.VersionComment = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetFrozen(v bool) *QueryPropertyVersionConfigForTerrformResponseData {
	s.Frozen = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVersionCreationTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.VersionCreationTime = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVersionLastUpdateTime(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.VersionLastUpdateTime = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetHostnames(v []*QueryPropertyVersionConfigForTerrformResponseDataHostnames) *QueryPropertyVersionConfigForTerrformResponseData {
	s.Hostnames = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetRules(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.Rules = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetOrigins(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.Origins = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseData) SetVariables(v string) *QueryPropertyVersionConfigForTerrformResponseData {
	s.Variables = &v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataStagingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataStagingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataProductionVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataProductionVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataStagingDeployingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion struct {
	// {"en":"Version of the property.","zh_CN":"项目的版本"}
	Version *int `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	// {"en":"hostnames.","zh_CN":"域名列表"}
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) SetVersion(v int) *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion {
	s.Version = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion) SetHostnames(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataProductionDeployingVersion {
	s.Hostnames = v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnames struct {
	// {"en":"default origin","zh_CN":"默认源站"}
	DefaultOrigin *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" require:"true" type:"Struct"`
	// {"en":"hostname","zh_CN":"域名"}
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
	// {"en":"hostname association certificate configuration","zh_CN":"关联证书配置"}
	Certificates []*QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" require:"true" type:"Repeated"`
	// {"en":"icp","zh_CN":"域名备案号"}
	Icp *string `json:"icp,omitempty" xml:"icp,omitempty" require:"true"`
	// {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
	EdgeHostname *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true" type:"Struct"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnames) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnames) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetDefaultOrigin(v *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
	s.DefaultOrigin = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetHostname(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
	s.Hostname = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetCertificates(v []*QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
	s.Certificates = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetIcp(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
	s.Icp = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnames) SetEdgeHostname(v *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) *QueryPropertyVersionConfigForTerrformResponseDataHostnames {
	s.EdgeHostname = v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin struct {
	// {"en":"origin servers","zh_CN":"源站服务器"}
	Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" require:"true" type:"Repeated"`
	// {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
	IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty" require:"true"`
	// {"en":"http port","zh_CN":"http端口"}
	HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty" require:"true"`
	// {"en":"origin host","zh_CN":"回源HOST"}
	Host *string `json:"host,omitempty" xml:"host,omitempty" require:"true"`
	// {"en":"https port","zh_CN":"https端口"}
	HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetServers(v []*string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.Servers = v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetIpVersion(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.IpVersion = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpPort(v int) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.HttpPort = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHost(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.Host = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin) SetHttpsPort(v int) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesDefaultOrigin {
	s.HttpsPort = &v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates struct {
	// {"en":"Certificate ID","zh_CN":"证书ID"}
	CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificate usage. data range: default_sni, ssl_bk, gm_sm2_enc, gm_sm2_sign, client_mtls","zh_CN":"证书用途。取值范围：default_sni, dual_sni, gm_sm2_enc, gm_sm2_sign, client_mtls"}
	CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) SetCertificateId(v int) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates {
	s.CertificateId = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates) SetCertificateUsage(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesCertificates {
	s.CertificateUsage = &v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname struct {
	// {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
	EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
	// {"en":"dns service status.data range: inactive,active","zh_CN":"DNS服务状态。取值范围：inactive, active。备注：inactive：挂起，active：生效"}
	DnsServiceStatus *string `json:"dnsServiceStatus,omitempty" xml:"dnsServiceStatus,omitempty" require:"true"`
	// {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
	EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
	// {"en":"edge-hostname","zh_CN":"调度域名"}
	EdgeHostname *string `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" require:"true"`
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) GoString() string {
	return s.String()
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.EdgeHostnamePrefix = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostname(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.EdgeHostname = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetDnsServiceStatus(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.DnsServiceStatus = &v
	return s
}

func (s *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *QueryPropertyVersionConfigForTerrformResponseDataHostnamesEdgeHostname {
	s.EdgeHostnameSuffix = &v
	return s
}

type QueryPropertyVersionConfigForTerrformResponseHeader struct {
}

func (s QueryPropertyVersionConfigForTerrformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s QueryPropertyVersionConfigForTerrformResponseHeader) GoString() string {
	return s.String()
}

type UpdatePropertyForTerraformRequest struct {
	// {"en":"The name of the property. The length must not exceed 256 characters.","zh_CN":"项目的名称。长度不超过256个字符。"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"Property description.The length must not exceed 256 characters.","zh_CN":"项目的描述。长度不超过256个字符。"}
	PropertyComment *string `json:"propertyComment,omitempty" xml:"propertyComment,omitempty"`
	// {"en":"Property version description.The length must not exceed 256 characters.","zh_CN":"项目版本的描述。长度不超过256个字符。"}
	VersionComment *string `json:"versionComment,omitempty" xml:"versionComment,omitempty"`
	// {"en":"hostnames","zh_CN":"域名列表"}
	Hostnames []*UpdatePropertyForTerraformRequestHostnames `json:"hostnames,omitempty" xml:"hostnames,omitempty" require:"true" type:"Repeated"`
	// {"en":"Specify the hostname and settings used to contact the origin once service begins.","zh_CN":"回源配置，详情查看schema接口。"}
	Origins *string `json:"origins,omitempty" xml:"origins,omitempty"`
	// {"en":"The Variables feature allows you to define variables, assign values to them, and reuse them in functional test cases. A variable consists of a name and a value.","zh_CN":"变量配置，详情查看schema接口。"}
	Variables *string `json:"variables,omitempty" xml:"variables,omitempty"`
	// {"en":"Rules","zh_CN":"规则配置"}
	Rules *string `json:"rules,omitempty" xml:"rules,omitempty"`
}

func (s UpdatePropertyForTerraformRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequest) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformRequest) SetPropertyName(v string) *UpdatePropertyForTerraformRequest {
	s.PropertyName = &v
	return s
}

func (s *UpdatePropertyForTerraformRequest) SetPropertyComment(v string) *UpdatePropertyForTerraformRequest {
	s.PropertyComment = &v
	return s
}

func (s *UpdatePropertyForTerraformRequest) SetVersionComment(v string) *UpdatePropertyForTerraformRequest {
	s.VersionComment = &v
	return s
}

func (s *UpdatePropertyForTerraformRequest) SetHostnames(v []*UpdatePropertyForTerraformRequestHostnames) *UpdatePropertyForTerraformRequest {
	s.Hostnames = v
	return s
}

func (s *UpdatePropertyForTerraformRequest) SetOrigins(v string) *UpdatePropertyForTerraformRequest {
	s.Origins = &v
	return s
}

func (s *UpdatePropertyForTerraformRequest) SetVariables(v string) *UpdatePropertyForTerraformRequest {
	s.Variables = &v
	return s
}

func (s *UpdatePropertyForTerraformRequest) SetRules(v string) *UpdatePropertyForTerraformRequest {
	s.Rules = &v
	return s
}

type UpdatePropertyForTerraformRequestHostnames struct {
	// {"en":"default origin","zh_CN":"默认源站"}
	DefaultOrigin *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin `json:"defaultOrigin,omitempty" xml:"defaultOrigin,omitempty" type:"Struct"`
	// {"en":"hostname, the length must not exceed 128 characters. A wildcard hostname must start with an asterisk (*).","zh_CN":"域名，长度不超过128个字符。泛域名需要以“*”开头。"}
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty" require:"true"`
	// {"en":"hostname association ssl configuration","zh_CN":"关联证书配置"}
	Certificates []*UpdatePropertyForTerraformRequestHostnamesCertificates `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	// {"en":"hostname edge-hostname config","zh_CN":"调度域名配置"}
	EdgeHostname *UpdatePropertyForTerraformRequestHostnamesEdgeHostname `json:"edgeHostname,omitempty" xml:"edgeHostname,omitempty" type:"Struct"`
}

func (s UpdatePropertyForTerraformRequestHostnames) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnames) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetDefaultOrigin(v *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) *UpdatePropertyForTerraformRequestHostnames {
	s.DefaultOrigin = v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetHostname(v string) *UpdatePropertyForTerraformRequestHostnames {
	s.Hostname = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetCertificates(v []*UpdatePropertyForTerraformRequestHostnamesCertificates) *UpdatePropertyForTerraformRequestHostnames {
	s.Certificates = v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnames) SetEdgeHostname(v *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) *UpdatePropertyForTerraformRequestHostnames {
	s.EdgeHostname = v
	return s
}

type UpdatePropertyForTerraformRequestHostnamesDefaultOrigin struct {
	// {"en":"origin servers","zh_CN":"源站服务器"}
	Servers []*string `json:"servers,omitempty" xml:"servers,omitempty" type:"Repeated"`
	// {"en":"IP version. data range: dual,ipv4,ipv6","zh_CN":"IP版本。取值范围：dual，ipv4，ipv6"}
	IpVersion *string `json:"ipVersion,omitempty" xml:"ipVersion,omitempty"`
	// {"en":"http port","zh_CN":"http端口"}
	HttpPort *int `json:"httpPort,omitempty" xml:"httpPort,omitempty"`
	// {"en":"origin host","zh_CN":"回源HOST"}
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// {"en":"https port","zh_CN":"https端口"}
	HttpsPort *int `json:"httpsPort,omitempty" xml:"httpsPort,omitempty"`
}

func (s UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetServers(v []*string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.Servers = v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetIpVersion(v string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.IpVersion = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpPort(v int) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.HttpPort = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHost(v string) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.Host = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin) SetHttpsPort(v int) *UpdatePropertyForTerraformRequestHostnamesDefaultOrigin {
	s.HttpsPort = &v
	return s
}

type UpdatePropertyForTerraformRequestHostnamesCertificates struct {
	// {"en":"Certificate ID","zh_CN":"证书ID"}
	CertificateId *int `json:"certificateId,omitempty" xml:"certificateId,omitempty" require:"true"`
	// {"en":"certificate usage. data range: default_sni, ssl_bk, gm_sm2_enc, gm_sm2_sign, client_mtls","zh_CN":"证书用途。取值范围：default_sni, dual_sni, gm_sm2_enc, gm_sm2_sign, client_mtls"}
	CertificateUsage *string `json:"certificateUsage,omitempty" xml:"certificateUsage,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformRequestHostnamesCertificates) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnamesCertificates) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnamesCertificates) SetCertificateId(v int) *UpdatePropertyForTerraformRequestHostnamesCertificates {
	s.CertificateId = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesCertificates) SetCertificateUsage(v string) *UpdatePropertyForTerraformRequestHostnamesCertificates {
	s.CertificateUsage = &v
	return s
}

type UpdatePropertyForTerraformRequestHostnamesEdgeHostname struct {
	// {"en":"edge-hostname prefix","zh_CN":"调度域名前缀"}
	EdgeHostnamePrefix *string `json:"edgeHostnamePrefix,omitempty" xml:"edgeHostnamePrefix,omitempty" require:"true"`
	// {"en":"edge-hostname comment","zh_CN":"调度域名描述"}
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// {"en":"edge-hostname suffix","zh_CN":"调度域名后缀"}
	EdgeHostnameSuffix *string `json:"edgeHostnameSuffix,omitempty" xml:"edgeHostnameSuffix,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformRequestHostnamesEdgeHostname) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHostnamesEdgeHostname) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnamePrefix(v string) *UpdatePropertyForTerraformRequestHostnamesEdgeHostname {
	s.EdgeHostnamePrefix = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) SetComment(v string) *UpdatePropertyForTerraformRequestHostnamesEdgeHostname {
	s.Comment = &v
	return s
}

func (s *UpdatePropertyForTerraformRequestHostnamesEdgeHostname) SetEdgeHostnameSuffix(v string) *UpdatePropertyForTerraformRequestHostnamesEdgeHostname {
	s.EdgeHostnameSuffix = &v
	return s
}

type UpdatePropertyForTerraformRequestHeader struct {
}

func (s UpdatePropertyForTerraformRequestHeader) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformRequestHeader) GoString() string {
	return s.String()
}

type UpdatePropertyForTerraformPaths struct {
	// {"en":"Property ID","zh_CN":"项目ID"}
	PropertyId *int `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformPaths) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformPaths) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformPaths) SetPropertyId(v int) *UpdatePropertyForTerraformPaths {
	s.PropertyId = &v
	return s
}

type UpdatePropertyForTerraformParameters struct {
}

func (s UpdatePropertyForTerraformParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformParameters) GoString() string {
	return s.String()
}

type UpdatePropertyForTerraformResponse struct {
	// {"en":"Response code, 0 means successful.","zh_CN":"接口响应code，0代表成功。"}
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// {"en":"Response error message if failed.","zh_CN":"接口响应信息，success代表成功，失败则提供失败信息。"}
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	// {"en":"Response data.","zh_CN":"接口响应数据"}
	Data *UpdatePropertyForTerraformResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Struct"`
}

func (s UpdatePropertyForTerraformResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformResponse) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformResponse) SetCode(v string) *UpdatePropertyForTerraformResponse {
	s.Code = &v
	return s
}

func (s *UpdatePropertyForTerraformResponse) SetMessage(v string) *UpdatePropertyForTerraformResponse {
	s.Message = &v
	return s
}

func (s *UpdatePropertyForTerraformResponse) SetData(v *UpdatePropertyForTerraformResponseData) *UpdatePropertyForTerraformResponse {
	s.Data = v
	return s
}

type UpdatePropertyForTerraformResponseData struct {
	// {"en":"Property ID","zh_CN":"项目ID"}
	PropertyId *int64 `json:"propertyId,omitempty" xml:"propertyId,omitempty" require:"true"`
	// {"en":"Property Name","zh_CN":"项目名称"}
	PropertyName *string `json:"propertyName,omitempty" xml:"propertyName,omitempty" require:"true"`
	// {"en":"Property Version","zh_CN":"项目版本"}
	PropertyVersion *int64 `json:"propertyVersion,omitempty" xml:"propertyVersion,omitempty" require:"true"`
}

func (s UpdatePropertyForTerraformResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformResponseData) GoString() string {
	return s.String()
}

func (s *UpdatePropertyForTerraformResponseData) SetPropertyId(v int64) *UpdatePropertyForTerraformResponseData {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyForTerraformResponseData) SetPropertyName(v string) *UpdatePropertyForTerraformResponseData {
	s.PropertyName = &v
	return s
}

func (s *UpdatePropertyForTerraformResponseData) SetPropertyVersion(v int64) *UpdatePropertyForTerraformResponseData {
	s.PropertyVersion = &v
	return s
}

type UpdatePropertyForTerraformResponseHeader struct {
}

func (s UpdatePropertyForTerraformResponseHeader) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyForTerraformResponseHeader) GoString() string {
	return s.String()
}

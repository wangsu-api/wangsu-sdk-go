package namespace

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type CreatePrimalNamespaceRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreatePrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s CreatePrimalNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceRequest) GoString() string {
  return s.String()
}

func (s *CreatePrimalNamespaceRequest) SetApiVersion(v string) *CreatePrimalNamespaceRequest {
  s.ApiVersion = &v
  return s
}

func (s *CreatePrimalNamespaceRequest) SetKind(v string) *CreatePrimalNamespaceRequest {
  s.Kind = &v
  return s
}

func (s *CreatePrimalNamespaceRequest) SetMetadata(v *CreatePrimalNamespaceObjectMeta) *CreatePrimalNamespaceRequest {
  s.Metadata = v
  return s
}

type CreatePrimalNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"namespace object", "zh_CN":"namespace对象"}
  Data *CreatePrimalNamespaceNamespace `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreatePrimalNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceResponse) GoString() string {
  return s.String()
}

func (s *CreatePrimalNamespaceResponse) SetCode(v int64) *CreatePrimalNamespaceResponse {
  s.Code = &v
  return s
}

func (s *CreatePrimalNamespaceResponse) SetMsg(v string) *CreatePrimalNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *CreatePrimalNamespaceResponse) SetRequestId(v string) *CreatePrimalNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *CreatePrimalNamespaceResponse) SetData(v *CreatePrimalNamespaceNamespace) *CreatePrimalNamespaceResponse {
  s.Data = v
  return s
}

type CreatePrimalNamespacePaths struct {
}

func (s CreatePrimalNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespacePaths) GoString() string {
  return s.String()
}

type CreatePrimalNamespaceParameters struct {
}

func (s CreatePrimalNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceParameters) GoString() string {
  return s.String()
}

type CreatePrimalNamespaceRequestHeader struct {
}

func (s CreatePrimalNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceRequestHeader) GoString() string {
  return s.String()
}

type CreatePrimalNamespaceResponseHeader struct {
}

func (s CreatePrimalNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceResponseHeader) GoString() string {
  return s.String()
}

type CreatePrimalNamespaceNamespace struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreatePrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s CreatePrimalNamespaceNamespace) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceNamespace) GoString() string {
  return s.String()
}

func (s *CreatePrimalNamespaceNamespace) SetApiVersion(v string) *CreatePrimalNamespaceNamespace {
  s.ApiVersion = &v
  return s
}

func (s *CreatePrimalNamespaceNamespace) SetKind(v string) *CreatePrimalNamespaceNamespace {
  s.Kind = &v
  return s
}

func (s *CreatePrimalNamespaceNamespace) SetMetadata(v *CreatePrimalNamespaceObjectMeta) *CreatePrimalNamespaceNamespace {
  s.Metadata = v
  return s
}

type CreatePrimalNamespaceObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"CreatePrimalNamespaceNamespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  CreatePrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*CreatePrimalNamespaceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*CreatePrimalNamespaceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s CreatePrimalNamespaceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceObjectMeta) GoString() string {
  return s.String()
}

func (s *CreatePrimalNamespaceObjectMeta) SetName(v string) *CreatePrimalNamespaceObjectMeta {
  s.Name = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetGenerateName(v string) *CreatePrimalNamespaceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetNamespace(v string) *CreatePrimalNamespaceObjectMeta {
  s.CreatePrimalNamespaceNamespace = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetSelfLink(v string) *CreatePrimalNamespaceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetUid(v string) *CreatePrimalNamespaceObjectMeta {
  s.Uid = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetResourceVersion(v string) *CreatePrimalNamespaceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetGeneration(v int64) *CreatePrimalNamespaceObjectMeta {
  s.Generation = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetCreationTimestamp(v string) *CreatePrimalNamespaceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetDeletionTimestamp(v string) *CreatePrimalNamespaceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *CreatePrimalNamespaceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetLabels(v map[string]*string) *CreatePrimalNamespaceObjectMeta {
  s.Labels = v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetAnnotations(v map[string]*string) *CreatePrimalNamespaceObjectMeta {
  s.Annotations = v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetOwnerReferences(v []*CreatePrimalNamespaceOwnerReference) *CreatePrimalNamespaceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetFinalizers(v []*string) *CreatePrimalNamespaceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetClusterName(v string) *CreatePrimalNamespaceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *CreatePrimalNamespaceObjectMeta) SetManagedFields(v []*CreatePrimalNamespaceManagedFieldsEntry) *CreatePrimalNamespaceObjectMeta {
  s.ManagedFields = v
  return s
}

type CreatePrimalNamespaceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this CreatePrimalNamespaceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'CreatePrimalNamespaceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“CreatePrimalNamespaceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"CreatePrimalNamespaceFieldsV1 holds the first JSON version format as described in the 'CreatePrimalNamespaceFieldsV1' type", "zh_CN":"CreatePrimalNamespaceFieldsV1 包含类型 “CreatePrimalNamespaceFieldsV1” 中描述的第一个 JSON 版本格式"}
  CreatePrimalNamespaceFieldsV1 *CreatePrimalNamespaceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s CreatePrimalNamespaceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetManager(v string) *CreatePrimalNamespaceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetOperation(v string) *CreatePrimalNamespaceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetApiVersion(v string) *CreatePrimalNamespaceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetTime(v string) *CreatePrimalNamespaceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetFieldsType(v string) *CreatePrimalNamespaceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetFieldsV1(v *CreatePrimalNamespaceFieldsV1) *CreatePrimalNamespaceManagedFieldsEntry {
  s.CreatePrimalNamespaceFieldsV1 = v
  return s
}

func (s *CreatePrimalNamespaceManagedFieldsEntry) SetSubresource(v string) *CreatePrimalNamespaceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type CreatePrimalNamespaceFieldsV1 struct {
}

func (s CreatePrimalNamespaceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceFieldsV1) GoString() string {
  return s.String()
}

type CreatePrimalNamespaceOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s CreatePrimalNamespaceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s CreatePrimalNamespaceOwnerReference) GoString() string {
  return s.String()
}

func (s *CreatePrimalNamespaceOwnerReference) SetApiVersion(v string) *CreatePrimalNamespaceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *CreatePrimalNamespaceOwnerReference) SetKind(v string) *CreatePrimalNamespaceOwnerReference {
  s.Kind = &v
  return s
}

func (s *CreatePrimalNamespaceOwnerReference) SetName(v string) *CreatePrimalNamespaceOwnerReference {
  s.Name = &v
  return s
}

func (s *CreatePrimalNamespaceOwnerReference) SetUid(v string) *CreatePrimalNamespaceOwnerReference {
  s.Uid = &v
  return s
}

func (s *CreatePrimalNamespaceOwnerReference) SetController(v bool) *CreatePrimalNamespaceOwnerReference {
  s.Controller = &v
  return s
}

func (s *CreatePrimalNamespaceOwnerReference) SetBlockOwnerDeletion(v bool) *CreatePrimalNamespaceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type ListPrimalNamespaceRequest struct {
}

func (s ListPrimalNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceRequest) GoString() string {
  return s.String()
}

type ListPrimalNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"namespace", "zh_CN":"namespace"}
  Data *ListPrimalNamespaceNamespaceList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListPrimalNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceResponse) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceResponse) SetCode(v int64) *ListPrimalNamespaceResponse {
  s.Code = &v
  return s
}

func (s *ListPrimalNamespaceResponse) SetMsg(v string) *ListPrimalNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *ListPrimalNamespaceResponse) SetRequestId(v string) *ListPrimalNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *ListPrimalNamespaceResponse) SetData(v *ListPrimalNamespaceNamespaceList) *ListPrimalNamespaceResponse {
  s.Data = v
  return s
}

type ListPrimalNamespacePaths struct {
}

func (s ListPrimalNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespacePaths) GoString() string {
  return s.String()
}

type ListPrimalNamespaceParameters struct {
}

func (s ListPrimalNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceParameters) GoString() string {
  return s.String()
}

type ListPrimalNamespaceRequestHeader struct {
}

func (s ListPrimalNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceRequestHeader) GoString() string {
  return s.String()
}

type ListPrimalNamespaceResponseHeader struct {
}

func (s ListPrimalNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceResponseHeader) GoString() string {
  return s.String()
}

type ListPrimalNamespaceNamespaceList struct {
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Standard list metadata", "zh_CN":"标准列表元数据"}
  Metadata *ListPrimalNamespaceListMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"List of ListPrimalNamespaceNamespace", "zh_CN":"ListPrimalNamespaceNamespace 列表"}
  Items []*ListPrimalNamespaceNamespace `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListPrimalNamespaceNamespaceList) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceNamespaceList) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceNamespaceList) SetKind(v string) *ListPrimalNamespaceNamespaceList {
  s.Kind = &v
  return s
}

func (s *ListPrimalNamespaceNamespaceList) SetApiVersion(v string) *ListPrimalNamespaceNamespaceList {
  s.ApiVersion = &v
  return s
}

func (s *ListPrimalNamespaceNamespaceList) SetMetadata(v *ListPrimalNamespaceListMeta) *ListPrimalNamespaceNamespaceList {
  s.Metadata = v
  return s
}

func (s *ListPrimalNamespaceNamespaceList) SetItems(v []*ListPrimalNamespaceNamespace) *ListPrimalNamespaceNamespaceList {
  s.Items = v
  return s
}

type ListPrimalNamespaceListMeta struct {
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system", "zh_CN":"selfLink 表示此对象的 URL，由系统填充，只读。已弃用：selfLink 是一个遗留的只读字段，不再由系统填充。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only", "zh_CN":"标识该对象的服务器内部版本的字符串，客户端可以用该字段来确定对象何时被更改。 该值对客户端是不透明的，并且应该原样传回给服务器。该值由系统填充，只读"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message", "zh_CN":"如果用户对返回的条目数量设置了限制，则 continue 可能被设置，表示服务器有更多可用的数据。 该值是不透明的，可用于向提供此列表服务的端点发出另一个请求，以检索下一组可用的对象。 如果服务器配置已更改或时间已过去几分钟，则可能无法继续提供一致的列表。 除非你在错误消息中收到此令牌（token），否则使用此 continue 值时返回的 resourceVersion 字段应该和第一个响应中的值是相同的"}
  Continue *string `json:"continue,omitempty" xml:"continue,omitempty"`
  // {"en":"remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is estimating the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact", "zh_CN":"remainingItemCount 是列表中未包含在此列表响应中的后续项目的数量。 如果列表请求包含标签或字段选择器，则剩余项目的数量是未知的，并且在序列化期间该字段将保持未设置和省略。 如果列表是完整的（因为它没有分块或者这是最后一个块），那么就没有剩余的项目，并且在序列化过程中该字段将保持未设置和省略。 早于 v1.15 的服务器不设置此字段。remainingItemCount 的预期用途是估计集合的大小。 客户端不应依赖于设置准确的 remainingItemCount"}
  RemainingItemCount *int64 `json:"remainingItemCount,omitempty" xml:"remainingItemCount,omitempty"`
}

func (s ListPrimalNamespaceListMeta) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceListMeta) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceListMeta) SetSelfLink(v string) *ListPrimalNamespaceListMeta {
  s.SelfLink = &v
  return s
}

func (s *ListPrimalNamespaceListMeta) SetResourceVersion(v string) *ListPrimalNamespaceListMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListPrimalNamespaceListMeta) SetContinue(v string) *ListPrimalNamespaceListMeta {
  s.Continue = &v
  return s
}

func (s *ListPrimalNamespaceListMeta) SetRemainingItemCount(v int64) *ListPrimalNamespaceListMeta {
  s.RemainingItemCount = &v
  return s
}

type ListPrimalNamespaceNamespace struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *ListPrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s ListPrimalNamespaceNamespace) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceNamespace) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceNamespace) SetApiVersion(v string) *ListPrimalNamespaceNamespace {
  s.ApiVersion = &v
  return s
}

func (s *ListPrimalNamespaceNamespace) SetKind(v string) *ListPrimalNamespaceNamespace {
  s.Kind = &v
  return s
}

func (s *ListPrimalNamespaceNamespace) SetMetadata(v *ListPrimalNamespaceObjectMeta) *ListPrimalNamespaceNamespace {
  s.Metadata = v
  return s
}

type ListPrimalNamespaceObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"ListPrimalNamespaceNamespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  ListPrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*ListPrimalNamespaceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*ListPrimalNamespaceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s ListPrimalNamespaceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceObjectMeta) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceObjectMeta) SetName(v string) *ListPrimalNamespaceObjectMeta {
  s.Name = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetGenerateName(v string) *ListPrimalNamespaceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetNamespace(v string) *ListPrimalNamespaceObjectMeta {
  s.ListPrimalNamespaceNamespace = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetSelfLink(v string) *ListPrimalNamespaceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetUid(v string) *ListPrimalNamespaceObjectMeta {
  s.Uid = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetResourceVersion(v string) *ListPrimalNamespaceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetGeneration(v int64) *ListPrimalNamespaceObjectMeta {
  s.Generation = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetCreationTimestamp(v string) *ListPrimalNamespaceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetDeletionTimestamp(v string) *ListPrimalNamespaceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *ListPrimalNamespaceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetLabels(v map[string]*string) *ListPrimalNamespaceObjectMeta {
  s.Labels = v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetAnnotations(v map[string]*string) *ListPrimalNamespaceObjectMeta {
  s.Annotations = v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetOwnerReferences(v []*ListPrimalNamespaceOwnerReference) *ListPrimalNamespaceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetFinalizers(v []*string) *ListPrimalNamespaceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetClusterName(v string) *ListPrimalNamespaceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *ListPrimalNamespaceObjectMeta) SetManagedFields(v []*ListPrimalNamespaceManagedFieldsEntry) *ListPrimalNamespaceObjectMeta {
  s.ManagedFields = v
  return s
}

type ListPrimalNamespaceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this ListPrimalNamespaceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'ListPrimalNamespaceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“ListPrimalNamespaceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"ListPrimalNamespaceFieldsV1 holds the first JSON version format as described in the 'ListPrimalNamespaceFieldsV1' type", "zh_CN":"ListPrimalNamespaceFieldsV1 包含类型 “ListPrimalNamespaceFieldsV1” 中描述的第一个 JSON 版本格式"}
  ListPrimalNamespaceFieldsV1 *ListPrimalNamespaceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s ListPrimalNamespaceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetManager(v string) *ListPrimalNamespaceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetOperation(v string) *ListPrimalNamespaceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetApiVersion(v string) *ListPrimalNamespaceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetTime(v string) *ListPrimalNamespaceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetFieldsType(v string) *ListPrimalNamespaceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetFieldsV1(v *ListPrimalNamespaceFieldsV1) *ListPrimalNamespaceManagedFieldsEntry {
  s.ListPrimalNamespaceFieldsV1 = v
  return s
}

func (s *ListPrimalNamespaceManagedFieldsEntry) SetSubresource(v string) *ListPrimalNamespaceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type ListPrimalNamespaceFieldsV1 struct {
}

func (s ListPrimalNamespaceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceFieldsV1) GoString() string {
  return s.String()
}

type ListPrimalNamespaceOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s ListPrimalNamespaceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s ListPrimalNamespaceOwnerReference) GoString() string {
  return s.String()
}

func (s *ListPrimalNamespaceOwnerReference) SetApiVersion(v string) *ListPrimalNamespaceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *ListPrimalNamespaceOwnerReference) SetKind(v string) *ListPrimalNamespaceOwnerReference {
  s.Kind = &v
  return s
}

func (s *ListPrimalNamespaceOwnerReference) SetName(v string) *ListPrimalNamespaceOwnerReference {
  s.Name = &v
  return s
}

func (s *ListPrimalNamespaceOwnerReference) SetUid(v string) *ListPrimalNamespaceOwnerReference {
  s.Uid = &v
  return s
}

func (s *ListPrimalNamespaceOwnerReference) SetController(v bool) *ListPrimalNamespaceOwnerReference {
  s.Controller = &v
  return s
}

func (s *ListPrimalNamespaceOwnerReference) SetBlockOwnerDeletion(v bool) *ListPrimalNamespaceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type ListNamespaceRequest struct {
}

func (s ListNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceRequest) GoString() string {
  return s.String()
}

type ListNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"namespace", "zh_CN":"namespace"}
  Data *ListNamespaceNamespaceList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceResponse) GoString() string {
  return s.String()
}

func (s *ListNamespaceResponse) SetCode(v int64) *ListNamespaceResponse {
  s.Code = &v
  return s
}

func (s *ListNamespaceResponse) SetMsg(v string) *ListNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *ListNamespaceResponse) SetRequestId(v string) *ListNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *ListNamespaceResponse) SetData(v *ListNamespaceNamespaceList) *ListNamespaceResponse {
  s.Data = v
  return s
}

type ListNamespacePaths struct {
}

func (s ListNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s ListNamespacePaths) GoString() string {
  return s.String()
}

type ListNamespaceParameters struct {
  // {"en":"name", "zh_CN":"名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceParameters) GoString() string {
  return s.String()
}

func (s *ListNamespaceParameters) SetName(v string) *ListNamespaceParameters {
  s.Name = &v
  return s
}

type ListNamespaceRequestHeader struct {
}

func (s ListNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceRequestHeader) GoString() string {
  return s.String()
}

type ListNamespaceResponseHeader struct {
}

func (s ListNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceResponseHeader) GoString() string {
  return s.String()
}

type ListNamespaceNamespaceList struct {
  // {"en":"total", "zh_CN":"总条数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
  // {"en":"list of namespace", "zh_CN":"命名空间列表"}
  NamespaceInfoList []*ListNamespaceNamespaceInfo `json:"namespaceInfoList,omitempty" xml:"namespaceInfoList,omitempty" require:"true" type:"Repeated"`
}

func (s ListNamespaceNamespaceList) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceNamespaceList) GoString() string {
  return s.String()
}

func (s *ListNamespaceNamespaceList) SetTotal(v int64) *ListNamespaceNamespaceList {
  s.Total = &v
  return s
}

func (s *ListNamespaceNamespaceList) SetNamespaceInfoList(v []*ListNamespaceNamespaceInfo) *ListNamespaceNamespaceList {
  s.NamespaceInfoList = v
  return s
}

type ListNamespaceNamespaceInfo struct {
  // {"en":"id", "zh_CN":"id"}
  Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
  // {"en":"name of namespace", "zh_CN":"命名空间名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"description of namespace", "zh_CN":"命名空间描述"}
  Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
  // {"en":"label of namespace", "zh_CN":"命名空间标签"}
  Label *string `json:"label,omitempty" xml:"label,omitempty" require:"true"`
  // {"en":"time of update", "zh_CN":"更新时间"}
  UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty" require:"true"`
}

func (s ListNamespaceNamespaceInfo) String() string {
  return tea.Prettify(s)
}

func (s ListNamespaceNamespaceInfo) GoString() string {
  return s.String()
}

func (s *ListNamespaceNamespaceInfo) SetId(v int64) *ListNamespaceNamespaceInfo {
  s.Id = &v
  return s
}

func (s *ListNamespaceNamespaceInfo) SetName(v string) *ListNamespaceNamespaceInfo {
  s.Name = &v
  return s
}

func (s *ListNamespaceNamespaceInfo) SetDescription(v string) *ListNamespaceNamespaceInfo {
  s.Description = &v
  return s
}

func (s *ListNamespaceNamespaceInfo) SetLabel(v string) *ListNamespaceNamespaceInfo {
  s.Label = &v
  return s
}

func (s *ListNamespaceNamespaceInfo) SetUpdateTime(v int64) *ListNamespaceNamespaceInfo {
  s.UpdateTime = &v
  return s
}




type DeletePrimalNamespaceRequest struct {
}

func (s DeletePrimalNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceRequest) GoString() string {
  return s.String()
}

type DeletePrimalNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"status"}
  Data *DeletePrimalNamespaceStatus `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeletePrimalNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceResponse) GoString() string {
  return s.String()
}

func (s *DeletePrimalNamespaceResponse) SetCode(v int64) *DeletePrimalNamespaceResponse {
  s.Code = &v
  return s
}

func (s *DeletePrimalNamespaceResponse) SetMsg(v string) *DeletePrimalNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *DeletePrimalNamespaceResponse) SetRequestId(v string) *DeletePrimalNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *DeletePrimalNamespaceResponse) SetData(v *DeletePrimalNamespaceStatus) *DeletePrimalNamespaceResponse {
  s.Data = v
  return s
}

type DeletePrimalNamespacePaths struct {
  // {"en":"namespace name", "zh_CN":"命名空间名称"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s DeletePrimalNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespacePaths) GoString() string {
  return s.String()
}

func (s *DeletePrimalNamespacePaths) SetNamespace(v string) *DeletePrimalNamespacePaths {
  s.Namespace = &v
  return s
}

type DeletePrimalNamespaceParameters struct {
}

func (s DeletePrimalNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceParameters) GoString() string {
  return s.String()
}

type DeletePrimalNamespaceRequestHeader struct {
}

func (s DeletePrimalNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceRequestHeader) GoString() string {
  return s.String()
}

type DeletePrimalNamespaceResponseHeader struct {
}

func (s DeletePrimalNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceResponseHeader) GoString() string {
  return s.String()
}

type DeletePrimalNamespaceStatus struct {
  // {"en":"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values", "zh_CN":"APIVersion 定义对象表示的版本化模式。 服务器应将已识别的模式转换为最新的内部值，并可能拒绝无法识别的值"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase", "zh_CN":"Kind 是一个字符串值，表示此对象表示的 REST 资源。 服务器可以从客户端提交请求的端点推断出这一点。 无法更新。驼峰式规则"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"DeletePrimalNamespaceStatus of the operation. One of: 'Success' or 'Failure'", "zh_CN":"操作状态。“Success”或“Failure” 之一"}
  DeletePrimalNamespaceStatus *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Suggested HTTP return code for this status, 0 if not set", "zh_CN":"此状态的建议 HTTP 返回代码，如果未设置，则为 0"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Extended data associated with the reason. Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type", "zh_CN":"与原因（Reason）相关的扩展数据。每个原因都可以定义自己的扩展细节。 此字段是可选的，并且不保证返回的数据符合任何模式，除非由原因类型定义"}
  Details *DeletePrimalNamespaceStatusDetails `json:"details,omitempty" xml:"details,omitempty" require:"true"`
}

func (s DeletePrimalNamespaceStatus) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceStatus) GoString() string {
  return s.String()
}

func (s *DeletePrimalNamespaceStatus) SetApiVersion(v string) *DeletePrimalNamespaceStatus {
  s.ApiVersion = &v
  return s
}

func (s *DeletePrimalNamespaceStatus) SetKind(v string) *DeletePrimalNamespaceStatus {
  s.Kind = &v
  return s
}

func (s *DeletePrimalNamespaceStatus) SetStatus(v string) *DeletePrimalNamespaceStatus {
  s.DeletePrimalNamespaceStatus = &v
  return s
}

func (s *DeletePrimalNamespaceStatus) SetCode(v int32) *DeletePrimalNamespaceStatus {
  s.Code = &v
  return s
}

func (s *DeletePrimalNamespaceStatus) SetDetails(v *DeletePrimalNamespaceStatusDetails) *DeletePrimalNamespaceStatus {
  s.Details = v
  return s
}

type DeletePrimalNamespaceStatusDetails struct {
  // {"en":"The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)", "zh_CN":"与状态 StatusReason 关联的资源的名称属性（当有一个可以描述的名称时）"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind", "zh_CN":"与状态 StatusReason 关联的资源的种类属性"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"The group attribute of the resource associated with the status StatusReason", "zh_CN":"与状态 StatusReason 关联的资源的组属性"}
  Group *string `json:"group,omitempty" xml:"group,omitempty" require:"true"`
  // {"en":"UID of the resource. (when there is a single resource which can be described)", "zh_CN":"资源的 UID（当有单个可以描述的资源时）"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty" require:"true"`
}

func (s DeletePrimalNamespaceStatusDetails) String() string {
  return tea.Prettify(s)
}

func (s DeletePrimalNamespaceStatusDetails) GoString() string {
  return s.String()
}

func (s *DeletePrimalNamespaceStatusDetails) SetName(v string) *DeletePrimalNamespaceStatusDetails {
  s.Name = &v
  return s
}

func (s *DeletePrimalNamespaceStatusDetails) SetKind(v string) *DeletePrimalNamespaceStatusDetails {
  s.Kind = &v
  return s
}

func (s *DeletePrimalNamespaceStatusDetails) SetGroup(v string) *DeletePrimalNamespaceStatusDetails {
  s.Group = &v
  return s
}

func (s *DeletePrimalNamespaceStatusDetails) SetUid(v string) *DeletePrimalNamespaceStatusDetails {
  s.Uid = &v
  return s
}




type GetPrimalNamespaceRequest struct {
}

func (s GetPrimalNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceRequest) GoString() string {
  return s.String()
}

type GetPrimalNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"namespace", "zh_CN":"namespace"}
  Data *GetPrimalNamespaceNamespace `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetPrimalNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceResponse) GoString() string {
  return s.String()
}

func (s *GetPrimalNamespaceResponse) SetCode(v int64) *GetPrimalNamespaceResponse {
  s.Code = &v
  return s
}

func (s *GetPrimalNamespaceResponse) SetMsg(v string) *GetPrimalNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *GetPrimalNamespaceResponse) SetRequestId(v string) *GetPrimalNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *GetPrimalNamespaceResponse) SetData(v *GetPrimalNamespaceNamespace) *GetPrimalNamespaceResponse {
  s.Data = v
  return s
}

type GetPrimalNamespacePaths struct {
  // {"en":"namespace name", "zh_CN":"命名空间名称"}
  GetPrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s GetPrimalNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespacePaths) GoString() string {
  return s.String()
}

func (s *GetPrimalNamespacePaths) SetNamespace(v string) *GetPrimalNamespacePaths {
  s.GetPrimalNamespaceNamespace = &v
  return s
}

type GetPrimalNamespaceParameters struct {
}

func (s GetPrimalNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceParameters) GoString() string {
  return s.String()
}

type GetPrimalNamespaceRequestHeader struct {
}

func (s GetPrimalNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceRequestHeader) GoString() string {
  return s.String()
}

type GetPrimalNamespaceResponseHeader struct {
}

func (s GetPrimalNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceResponseHeader) GoString() string {
  return s.String()
}

type GetPrimalNamespaceNamespace struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *GetPrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s GetPrimalNamespaceNamespace) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceNamespace) GoString() string {
  return s.String()
}

func (s *GetPrimalNamespaceNamespace) SetApiVersion(v string) *GetPrimalNamespaceNamespace {
  s.ApiVersion = &v
  return s
}

func (s *GetPrimalNamespaceNamespace) SetKind(v string) *GetPrimalNamespaceNamespace {
  s.Kind = &v
  return s
}

func (s *GetPrimalNamespaceNamespace) SetMetadata(v *GetPrimalNamespaceObjectMeta) *GetPrimalNamespaceNamespace {
  s.Metadata = v
  return s
}

type GetPrimalNamespaceObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"GetPrimalNamespaceNamespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  GetPrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*GetPrimalNamespaceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*GetPrimalNamespaceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s GetPrimalNamespaceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceObjectMeta) GoString() string {
  return s.String()
}

func (s *GetPrimalNamespaceObjectMeta) SetName(v string) *GetPrimalNamespaceObjectMeta {
  s.Name = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetGenerateName(v string) *GetPrimalNamespaceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetNamespace(v string) *GetPrimalNamespaceObjectMeta {
  s.GetPrimalNamespaceNamespace = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetSelfLink(v string) *GetPrimalNamespaceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetUid(v string) *GetPrimalNamespaceObjectMeta {
  s.Uid = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetResourceVersion(v string) *GetPrimalNamespaceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetGeneration(v int64) *GetPrimalNamespaceObjectMeta {
  s.Generation = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetCreationTimestamp(v string) *GetPrimalNamespaceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetDeletionTimestamp(v string) *GetPrimalNamespaceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *GetPrimalNamespaceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetLabels(v map[string]*string) *GetPrimalNamespaceObjectMeta {
  s.Labels = v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetAnnotations(v map[string]*string) *GetPrimalNamespaceObjectMeta {
  s.Annotations = v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetOwnerReferences(v []*GetPrimalNamespaceOwnerReference) *GetPrimalNamespaceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetFinalizers(v []*string) *GetPrimalNamespaceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetClusterName(v string) *GetPrimalNamespaceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *GetPrimalNamespaceObjectMeta) SetManagedFields(v []*GetPrimalNamespaceManagedFieldsEntry) *GetPrimalNamespaceObjectMeta {
  s.ManagedFields = v
  return s
}

type GetPrimalNamespaceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this GetPrimalNamespaceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'GetPrimalNamespaceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“GetPrimalNamespaceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"GetPrimalNamespaceFieldsV1 holds the first JSON version format as described in the 'GetPrimalNamespaceFieldsV1' type", "zh_CN":"GetPrimalNamespaceFieldsV1 包含类型 “GetPrimalNamespaceFieldsV1” 中描述的第一个 JSON 版本格式"}
  GetPrimalNamespaceFieldsV1 *GetPrimalNamespaceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s GetPrimalNamespaceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetManager(v string) *GetPrimalNamespaceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetOperation(v string) *GetPrimalNamespaceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetApiVersion(v string) *GetPrimalNamespaceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetTime(v string) *GetPrimalNamespaceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetFieldsType(v string) *GetPrimalNamespaceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetFieldsV1(v *GetPrimalNamespaceFieldsV1) *GetPrimalNamespaceManagedFieldsEntry {
  s.GetPrimalNamespaceFieldsV1 = v
  return s
}

func (s *GetPrimalNamespaceManagedFieldsEntry) SetSubresource(v string) *GetPrimalNamespaceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type GetPrimalNamespaceFieldsV1 struct {
}

func (s GetPrimalNamespaceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceFieldsV1) GoString() string {
  return s.String()
}

type GetPrimalNamespaceOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s GetPrimalNamespaceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s GetPrimalNamespaceOwnerReference) GoString() string {
  return s.String()
}

func (s *GetPrimalNamespaceOwnerReference) SetApiVersion(v string) *GetPrimalNamespaceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *GetPrimalNamespaceOwnerReference) SetKind(v string) *GetPrimalNamespaceOwnerReference {
  s.Kind = &v
  return s
}

func (s *GetPrimalNamespaceOwnerReference) SetName(v string) *GetPrimalNamespaceOwnerReference {
  s.Name = &v
  return s
}

func (s *GetPrimalNamespaceOwnerReference) SetUid(v string) *GetPrimalNamespaceOwnerReference {
  s.Uid = &v
  return s
}

func (s *GetPrimalNamespaceOwnerReference) SetController(v bool) *GetPrimalNamespaceOwnerReference {
  s.Controller = &v
  return s
}

func (s *GetPrimalNamespaceOwnerReference) SetBlockOwnerDeletion(v bool) *GetPrimalNamespaceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type PutPatchPrimalNamespaceRequest struct {
}

func (s PutPatchPrimalNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceRequest) GoString() string {
  return s.String()
}

type PutPatchPrimalNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"PutPatchPrimalNamespaceNamespace", "zh_CN":"PutPatchPrimalNamespaceNamespace"}
  Data *PutPatchPrimalNamespaceNamespace `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PutPatchPrimalNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceResponse) GoString() string {
  return s.String()
}

func (s *PutPatchPrimalNamespaceResponse) SetCode(v int64) *PutPatchPrimalNamespaceResponse {
  s.Code = &v
  return s
}

func (s *PutPatchPrimalNamespaceResponse) SetMsg(v string) *PutPatchPrimalNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *PutPatchPrimalNamespaceResponse) SetRequestId(v string) *PutPatchPrimalNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *PutPatchPrimalNamespaceResponse) SetData(v *PutPatchPrimalNamespaceNamespace) *PutPatchPrimalNamespaceResponse {
  s.Data = v
  return s
}

type PutPatchPrimalNamespacePaths struct {
  // {"en":"namespace name", "zh_CN":"命名空间名称"}
  PutPatchPrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s PutPatchPrimalNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespacePaths) GoString() string {
  return s.String()
}

func (s *PutPatchPrimalNamespacePaths) SetNamespace(v string) *PutPatchPrimalNamespacePaths {
  s.PutPatchPrimalNamespaceNamespace = &v
  return s
}

type PutPatchPrimalNamespaceParameters struct {
}

func (s PutPatchPrimalNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceParameters) GoString() string {
  return s.String()
}

type PutPatchPrimalNamespaceRequestHeader struct {
}

func (s PutPatchPrimalNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceRequestHeader) GoString() string {
  return s.String()
}

type PutPatchPrimalNamespaceResponseHeader struct {
}

func (s PutPatchPrimalNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceResponseHeader) GoString() string {
  return s.String()
}

type PutPatchPrimalNamespaceNamespace struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *PutPatchPrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s PutPatchPrimalNamespaceNamespace) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceNamespace) GoString() string {
  return s.String()
}

func (s *PutPatchPrimalNamespaceNamespace) SetApiVersion(v string) *PutPatchPrimalNamespaceNamespace {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchPrimalNamespaceNamespace) SetKind(v string) *PutPatchPrimalNamespaceNamespace {
  s.Kind = &v
  return s
}

func (s *PutPatchPrimalNamespaceNamespace) SetMetadata(v *PutPatchPrimalNamespaceObjectMeta) *PutPatchPrimalNamespaceNamespace {
  s.Metadata = v
  return s
}

type PutPatchPrimalNamespaceObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"PutPatchPrimalNamespaceNamespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  PutPatchPrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*PutPatchPrimalNamespaceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*PutPatchPrimalNamespaceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s PutPatchPrimalNamespaceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceObjectMeta) GoString() string {
  return s.String()
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetName(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.Name = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetGenerateName(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetNamespace(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.PutPatchPrimalNamespaceNamespace = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetSelfLink(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetUid(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.Uid = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetResourceVersion(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetGeneration(v int64) *PutPatchPrimalNamespaceObjectMeta {
  s.Generation = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetCreationTimestamp(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetDeletionTimestamp(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *PutPatchPrimalNamespaceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetLabels(v map[string]*string) *PutPatchPrimalNamespaceObjectMeta {
  s.Labels = v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetAnnotations(v map[string]*string) *PutPatchPrimalNamespaceObjectMeta {
  s.Annotations = v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetOwnerReferences(v []*PutPatchPrimalNamespaceOwnerReference) *PutPatchPrimalNamespaceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetFinalizers(v []*string) *PutPatchPrimalNamespaceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetClusterName(v string) *PutPatchPrimalNamespaceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *PutPatchPrimalNamespaceObjectMeta) SetManagedFields(v []*PutPatchPrimalNamespaceManagedFieldsEntry) *PutPatchPrimalNamespaceObjectMeta {
  s.ManagedFields = v
  return s
}

type PutPatchPrimalNamespaceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this PutPatchPrimalNamespaceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'PutPatchPrimalNamespaceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“PutPatchPrimalNamespaceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"PutPatchPrimalNamespaceFieldsV1 holds the first JSON version format as described in the 'PutPatchPrimalNamespaceFieldsV1' type", "zh_CN":"PutPatchPrimalNamespaceFieldsV1 包含类型 “PutPatchPrimalNamespaceFieldsV1” 中描述的第一个 JSON 版本格式"}
  PutPatchPrimalNamespaceFieldsV1 *PutPatchPrimalNamespaceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s PutPatchPrimalNamespaceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetManager(v string) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetOperation(v string) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetApiVersion(v string) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetTime(v string) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetFieldsType(v string) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetFieldsV1(v *PutPatchPrimalNamespaceFieldsV1) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.PutPatchPrimalNamespaceFieldsV1 = v
  return s
}

func (s *PutPatchPrimalNamespaceManagedFieldsEntry) SetSubresource(v string) *PutPatchPrimalNamespaceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type PutPatchPrimalNamespaceFieldsV1 struct {
}

func (s PutPatchPrimalNamespaceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceFieldsV1) GoString() string {
  return s.String()
}

type PutPatchPrimalNamespaceOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s PutPatchPrimalNamespaceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPrimalNamespaceOwnerReference) GoString() string {
  return s.String()
}

func (s *PutPatchPrimalNamespaceOwnerReference) SetApiVersion(v string) *PutPatchPrimalNamespaceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchPrimalNamespaceOwnerReference) SetKind(v string) *PutPatchPrimalNamespaceOwnerReference {
  s.Kind = &v
  return s
}

func (s *PutPatchPrimalNamespaceOwnerReference) SetName(v string) *PutPatchPrimalNamespaceOwnerReference {
  s.Name = &v
  return s
}

func (s *PutPatchPrimalNamespaceOwnerReference) SetUid(v string) *PutPatchPrimalNamespaceOwnerReference {
  s.Uid = &v
  return s
}

func (s *PutPatchPrimalNamespaceOwnerReference) SetController(v bool) *PutPatchPrimalNamespaceOwnerReference {
  s.Controller = &v
  return s
}

func (s *PutPatchPrimalNamespaceOwnerReference) SetBlockOwnerDeletion(v bool) *PutPatchPrimalNamespaceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}




type UpdatePrimalNamespaceRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdatePrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s UpdatePrimalNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceRequest) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespaceRequest) SetApiVersion(v string) *UpdatePrimalNamespaceRequest {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePrimalNamespaceRequest) SetKind(v string) *UpdatePrimalNamespaceRequest {
  s.Kind = &v
  return s
}

func (s *UpdatePrimalNamespaceRequest) SetMetadata(v *UpdatePrimalNamespaceObjectMeta) *UpdatePrimalNamespaceRequest {
  s.Metadata = v
  return s
}

type UpdatePrimalNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"namespace", "zh_CN":"namespace"}
  Data *UpdatePrimalNamespaceNamespace `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdatePrimalNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceResponse) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespaceResponse) SetCode(v int64) *UpdatePrimalNamespaceResponse {
  s.Code = &v
  return s
}

func (s *UpdatePrimalNamespaceResponse) SetMsg(v string) *UpdatePrimalNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *UpdatePrimalNamespaceResponse) SetRequestId(v string) *UpdatePrimalNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *UpdatePrimalNamespaceResponse) SetData(v *UpdatePrimalNamespaceNamespace) *UpdatePrimalNamespaceResponse {
  s.Data = v
  return s
}

type UpdatePrimalNamespacePaths struct {
  // {"en":"namespace name", "zh_CN":"命名空间名称"}
  UpdatePrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s UpdatePrimalNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespacePaths) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespacePaths) SetNamespace(v string) *UpdatePrimalNamespacePaths {
  s.UpdatePrimalNamespaceNamespace = &v
  return s
}

type UpdatePrimalNamespaceParameters struct {
}

func (s UpdatePrimalNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceParameters) GoString() string {
  return s.String()
}

type UpdatePrimalNamespaceRequestHeader struct {
}

func (s UpdatePrimalNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceRequestHeader) GoString() string {
  return s.String()
}

type UpdatePrimalNamespaceResponseHeader struct {
}

func (s UpdatePrimalNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceResponseHeader) GoString() string {
  return s.String()
}

type UpdatePrimalNamespaceNamespace struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdatePrimalNamespaceObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s UpdatePrimalNamespaceNamespace) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceNamespace) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespaceNamespace) SetApiVersion(v string) *UpdatePrimalNamespaceNamespace {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePrimalNamespaceNamespace) SetKind(v string) *UpdatePrimalNamespaceNamespace {
  s.Kind = &v
  return s
}

func (s *UpdatePrimalNamespaceNamespace) SetMetadata(v *UpdatePrimalNamespaceObjectMeta) *UpdatePrimalNamespaceNamespace {
  s.Metadata = v
  return s
}

type UpdatePrimalNamespaceObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"UpdatePrimalNamespaceNamespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  UpdatePrimalNamespaceNamespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system.", "zh_CN":"表示此对象的 URL。由系统填充。只读。已弃用。Kubernetes 将在 1.20 版本中停止传播该字段，并计划在 1.21 版本中删除该字段。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.Populated by the system. Read-only", "zh_CN":"该对象在时间和空间上的唯一值。它通常由服务器在成功创建资源时生成，并且不允许使用 PUT 操作更改。由系统填充。只读"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.Populated by the system. Read-only. Value must be treated as opaque by clients and", "zh_CN":"一个不透明的值，表示此对象的内部版本，客户端可以使用该值来确定对象是否已被更改。 可用于乐观并发、变更检测以及对资源或资源集的监听操作。 客户端必须将这些值视为不透明的，且未更改地传回服务器。 它们可能仅对特定资源或一组资源有效。由系统填充。只读。客户端必须将值视为不透明。"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"A sequence number representing a specific generation of the desired state. Populated by the system. Read-only", "zh_CN":"表示期望状态的特定生成的序列号。由系统填充。只读"}
  Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
  // {"en":"a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.Populated by the system. Read-only. Null for lists", "zh_CN":"一个时间戳，表示创建此对象时的服务器时间。 不能保证在单独的操作中按发生前的顺序设置。 客户端不得设置此值。它以 RFC3339 形式表示，并采用 UTC。由系统填充。只读。列表为空"}
  CreationTimestamp *string `json:"creationTimestamp,omitempty" xml:"creationTimestamp,omitempty"`
  // {"en":"RFC 3339 date and time at which this resource will be deleted", "zh_CN":"删除此资源的 RFC 3339 日期和时间"}
  DeletionTimestamp *string `json:"deletionTimestamp,omitempty" xml:"deletionTimestamp,omitempty"`
  // {"en":"Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only", "zh_CN":"此对象从系统中删除之前允许正常终止的秒数。 仅当设置了 deletionTimestamp 时才设置。 只能缩短。只读"}
  DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty" xml:"deletionGracePeriodSeconds,omitempty"`
  // {"en":"Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services", "zh_CN":"可用于组织和分类（确定范围和选择）对象的字符串键和值的映射。 可以匹配 ReplicationController 和 Service 的选择算符"}
  Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
  // {"en":"Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects", "zh_CN":"annotations 是一个非结构化的键值映射，存储在资源中，可以由外部工具设置以存储和检索任意元数据。 它们不可查询，在修改对象时应保留"}
  Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
  // {"en":"List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller", "zh_CN":"此对象所依赖的对象列表。如果列表中的所有对象都已被删除，则该对象将被垃圾回收。 如果此对象由控制器管理，则此列表中的条目将指向此控制器，controller 字段设置为 true。 管理控制器不能超过一个"}
  OwnerReferences []*UpdatePrimalNamespaceOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
  // {"en":"ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object", "zh_CN":"managedFields 将 workflow-id 和版本映射到由该工作流管理的字段集。 这主要用于内部管理，用户通常不需要设置或理解该字段。 工作流可以是用户名、控制器名或特定应用路径的名称，如 “ci-cd”。 字段集始终存在于修改对象时工作流使用的版本"}
  ManagedFields []*UpdatePrimalNamespaceManagedFieldsEntry `json:"managedFields,omitempty" xml:"managedFields,omitempty" type:"Repeated"`
}

func (s UpdatePrimalNamespaceObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceObjectMeta) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespaceObjectMeta) SetName(v string) *UpdatePrimalNamespaceObjectMeta {
  s.Name = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetGenerateName(v string) *UpdatePrimalNamespaceObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetNamespace(v string) *UpdatePrimalNamespaceObjectMeta {
  s.UpdatePrimalNamespaceNamespace = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetSelfLink(v string) *UpdatePrimalNamespaceObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetUid(v string) *UpdatePrimalNamespaceObjectMeta {
  s.Uid = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetResourceVersion(v string) *UpdatePrimalNamespaceObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetGeneration(v int64) *UpdatePrimalNamespaceObjectMeta {
  s.Generation = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetCreationTimestamp(v string) *UpdatePrimalNamespaceObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetDeletionTimestamp(v string) *UpdatePrimalNamespaceObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetDeletionGracePeriodSeconds(v int64) *UpdatePrimalNamespaceObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetLabels(v map[string]*string) *UpdatePrimalNamespaceObjectMeta {
  s.Labels = v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetAnnotations(v map[string]*string) *UpdatePrimalNamespaceObjectMeta {
  s.Annotations = v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetOwnerReferences(v []*UpdatePrimalNamespaceOwnerReference) *UpdatePrimalNamespaceObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetFinalizers(v []*string) *UpdatePrimalNamespaceObjectMeta {
  s.Finalizers = v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetClusterName(v string) *UpdatePrimalNamespaceObjectMeta {
  s.ClusterName = &v
  return s
}

func (s *UpdatePrimalNamespaceObjectMeta) SetManagedFields(v []*UpdatePrimalNamespaceManagedFieldsEntry) *UpdatePrimalNamespaceObjectMeta {
  s.ManagedFields = v
  return s
}

type UpdatePrimalNamespaceManagedFieldsEntry struct {
  // {"en":"an identifier of the workflow managing these fields", "zh_CN":"管理这些字段的工作流的标识符"}
  Manager *string `json:"manager,omitempty" xml:"manager,omitempty"`
  // {"en":"the type of operation which lead to this UpdatePrimalNamespaceManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'", "zh_CN":"导致创建此 managedFields 表项的操作类型。 此字段的仅有合法值是 “Apply” 和 “Update”"}
  Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
  // {"en":"defines the version of this resource that this field set applies to. The format is \"group\/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted", "zh_CN":"定义此字段集适用的资源的版本。 格式是 “group\/version”，就像顶级 apiVersion 字段一样。 必须跟踪字段集的版本，因为它不能自动转换"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"the timestamp of when the ManagedFields entry was added", "zh_CN":"添加 managedFields 条目时的时间戳"}
  Time *string `json:"time,omitempty" xml:"time,omitempty"`
  // {"en":"the discriminator for the different fields format and version. There is currently only one possible value: 'UpdatePrimalNamespaceFieldsV1'", "zh_CN":"不同字段格式和版本的鉴别器。 目前只有一个可能的值：“UpdatePrimalNamespaceFieldsV1”"}
  FieldsType *string `json:"fieldsType,omitempty" xml:"fieldsType,omitempty"`
  // {"en":"UpdatePrimalNamespaceFieldsV1 holds the first JSON version format as described in the 'UpdatePrimalNamespaceFieldsV1' type", "zh_CN":"UpdatePrimalNamespaceFieldsV1 包含类型 “UpdatePrimalNamespaceFieldsV1” 中描述的第一个 JSON 版本格式"}
  UpdatePrimalNamespaceFieldsV1 *UpdatePrimalNamespaceFieldsV1 `json:"fieldsV1,omitempty" xml:"fieldsV1,omitempty"`
  // {"en":"the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource", "zh_CN":"用于更新该对象的子资源的名称，如果对象是通过主资源更新的，则为空字符串。 该字段的值用于区分管理者，即使他们共享相同的名称。例如，状态更新将不同于使用相同管理者名称的常规更新。 请注意，apiVersion 字段与 subresource 字段无关，它始终对应于主资源的版本"}
  Subresource *string `json:"subresource,omitempty" xml:"subresource,omitempty"`
}

func (s UpdatePrimalNamespaceManagedFieldsEntry) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceManagedFieldsEntry) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetManager(v string) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.Manager = &v
  return s
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetOperation(v string) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.Operation = &v
  return s
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetApiVersion(v string) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetTime(v string) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.Time = &v
  return s
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetFieldsType(v string) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.FieldsType = &v
  return s
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetFieldsV1(v *UpdatePrimalNamespaceFieldsV1) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.UpdatePrimalNamespaceFieldsV1 = v
  return s
}

func (s *UpdatePrimalNamespaceManagedFieldsEntry) SetSubresource(v string) *UpdatePrimalNamespaceManagedFieldsEntry {
  s.Subresource = &v
  return s
}

type UpdatePrimalNamespaceFieldsV1 struct {
}

func (s UpdatePrimalNamespaceFieldsV1) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceFieldsV1) GoString() string {
  return s.String()
}

type UpdatePrimalNamespaceOwnerReference struct {
  // {"en":"API version of the referent", "zh_CN":"被引用资源的 API 版本"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
  // {"en":"Kind of the referent", "zh_CN":"被引用资源的类别"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"Name of the referent", "zh_CN":"被引用资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"UID of the referent", "zh_CN":"被引用资源的 uid"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
  // {"en":"If true, this reference points to the managing controller", "zh_CN":"如果为 true，则此引用指向管理的控制器"}
  Controller *bool `json:"controller,omitempty" xml:"controller,omitempty"`
  // {"en":"If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed", "zh_CN":"如果为 true，**并且** 如果属主具有 “foregroundDeletion” 终结器，则在删除此引用之前，无法从键值存储中删除属主。 默认为 false。要设置此字段，用户需要属主的 “delete” 权限， 否则将返回 422 (Unprocessable Entity)"}
  BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" xml:"blockOwnerDeletion,omitempty"`
}

func (s UpdatePrimalNamespaceOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s UpdatePrimalNamespaceOwnerReference) GoString() string {
  return s.String()
}

func (s *UpdatePrimalNamespaceOwnerReference) SetApiVersion(v string) *UpdatePrimalNamespaceOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePrimalNamespaceOwnerReference) SetKind(v string) *UpdatePrimalNamespaceOwnerReference {
  s.Kind = &v
  return s
}

func (s *UpdatePrimalNamespaceOwnerReference) SetName(v string) *UpdatePrimalNamespaceOwnerReference {
  s.Name = &v
  return s
}

func (s *UpdatePrimalNamespaceOwnerReference) SetUid(v string) *UpdatePrimalNamespaceOwnerReference {
  s.Uid = &v
  return s
}

func (s *UpdatePrimalNamespaceOwnerReference) SetController(v bool) *UpdatePrimalNamespaceOwnerReference {
  s.Controller = &v
  return s
}

func (s *UpdatePrimalNamespaceOwnerReference) SetBlockOwnerDeletion(v bool) *UpdatePrimalNamespaceOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}





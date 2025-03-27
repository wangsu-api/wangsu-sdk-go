package storagemanagement

// This file is auto-generated, don't edit it. Thanks.
import (
  "github.com/alibabacloud-go/tea/tea"
)

type UpdatePvcsRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdatePvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *UpdatePvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s UpdatePvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsRequest) GoString() string {
  return s.String()
}

func (s *UpdatePvcsRequest) SetApiVersion(v string) *UpdatePvcsRequest {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePvcsRequest) SetKind(v string) *UpdatePvcsRequest {
  s.Kind = &v
  return s
}

func (s *UpdatePvcsRequest) SetMetadata(v *UpdatePvcsObjectMeta) *UpdatePvcsRequest {
  s.Metadata = v
  return s
}

func (s *UpdatePvcsRequest) SetSpec(v *UpdatePvcsPersistentVolumeClaimSpec) *UpdatePvcsRequest {
  s.Spec = v
  return s
}

type UpdatePvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc", "zh_CN":"pvc"}
  Data *UpdatePvcsPersistentVolumeClaim `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UpdatePvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsResponse) GoString() string {
  return s.String()
}

func (s *UpdatePvcsResponse) SetCode(v int64) *UpdatePvcsResponse {
  s.Code = &v
  return s
}

func (s *UpdatePvcsResponse) SetMsg(v string) *UpdatePvcsResponse {
  s.Msg = &v
  return s
}

func (s *UpdatePvcsResponse) SetRequestId(v string) *UpdatePvcsResponse {
  s.RequestId = &v
  return s
}

func (s *UpdatePvcsResponse) SetData(v *UpdatePvcsPersistentVolumeClaim) *UpdatePvcsResponse {
  s.Data = v
  return s
}

type UpdatePvcsPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"deployment name", "zh_CN":"pvc 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s UpdatePvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsPaths) GoString() string {
  return s.String()
}

func (s *UpdatePvcsPaths) SetNamespace(v string) *UpdatePvcsPaths {
  s.Namespace = &v
  return s
}

func (s *UpdatePvcsPaths) SetName(v string) *UpdatePvcsPaths {
  s.Name = &v
  return s
}

type UpdatePvcsParameters struct {
}

func (s UpdatePvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsParameters) GoString() string {
  return s.String()
}

type UpdatePvcsRequestHeader struct {
}

func (s UpdatePvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsRequestHeader) GoString() string {
  return s.String()
}

type UpdatePvcsResponseHeader struct {
}

func (s UpdatePvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsResponseHeader) GoString() string {
  return s.String()
}

type UpdatePvcsObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
  OwnerReferences []*UpdatePvcsOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
}

func (s UpdatePvcsObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsObjectMeta) GoString() string {
  return s.String()
}

func (s *UpdatePvcsObjectMeta) SetName(v string) *UpdatePvcsObjectMeta {
  s.Name = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetGenerateName(v string) *UpdatePvcsObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetNamespace(v string) *UpdatePvcsObjectMeta {
  s.Namespace = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetSelfLink(v string) *UpdatePvcsObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetUid(v string) *UpdatePvcsObjectMeta {
  s.Uid = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetResourceVersion(v string) *UpdatePvcsObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetGeneration(v int64) *UpdatePvcsObjectMeta {
  s.Generation = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetCreationTimestamp(v string) *UpdatePvcsObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetDeletionTimestamp(v string) *UpdatePvcsObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetDeletionGracePeriodSeconds(v int64) *UpdatePvcsObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *UpdatePvcsObjectMeta) SetLabels(v map[string]*string) *UpdatePvcsObjectMeta {
  s.Labels = v
  return s
}

func (s *UpdatePvcsObjectMeta) SetAnnotations(v map[string]*string) *UpdatePvcsObjectMeta {
  s.Annotations = v
  return s
}

func (s *UpdatePvcsObjectMeta) SetOwnerReferences(v []*UpdatePvcsOwnerReference) *UpdatePvcsObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *UpdatePvcsObjectMeta) SetFinalizers(v []*string) *UpdatePvcsObjectMeta {
  s.Finalizers = v
  return s
}

func (s *UpdatePvcsObjectMeta) SetClusterName(v string) *UpdatePvcsObjectMeta {
  s.ClusterName = &v
  return s
}

type UpdatePvcsOwnerReference struct {
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

func (s UpdatePvcsOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsOwnerReference) GoString() string {
  return s.String()
}

func (s *UpdatePvcsOwnerReference) SetApiVersion(v string) *UpdatePvcsOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePvcsOwnerReference) SetKind(v string) *UpdatePvcsOwnerReference {
  s.Kind = &v
  return s
}

func (s *UpdatePvcsOwnerReference) SetName(v string) *UpdatePvcsOwnerReference {
  s.Name = &v
  return s
}

func (s *UpdatePvcsOwnerReference) SetUid(v string) *UpdatePvcsOwnerReference {
  s.Uid = &v
  return s
}

func (s *UpdatePvcsOwnerReference) SetController(v bool) *UpdatePvcsOwnerReference {
  s.Controller = &v
  return s
}

func (s *UpdatePvcsOwnerReference) SetBlockOwnerDeletion(v bool) *UpdatePvcsOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type UpdatePvcsPersistentVolumeClaim struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *UpdatePvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *UpdatePvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"represents the current information/status of a persistent volume claim. Read-only", "zh_CN":"表示一个持久卷申领的当前信息/状态。只读"}
  Status *UpdatePvcsPersistentVolumeClaimStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdatePvcsPersistentVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsPersistentVolumeClaim) GoString() string {
  return s.String()
}

func (s *UpdatePvcsPersistentVolumeClaim) SetApiVersion(v string) *UpdatePvcsPersistentVolumeClaim {
  s.ApiVersion = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaim) SetKind(v string) *UpdatePvcsPersistentVolumeClaim {
  s.Kind = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaim) SetMetadata(v *UpdatePvcsObjectMeta) *UpdatePvcsPersistentVolumeClaim {
  s.Metadata = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaim) SetSpec(v *UpdatePvcsPersistentVolumeClaimSpec) *UpdatePvcsPersistentVolumeClaim {
  s.Spec = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaim) SetStatus(v *UpdatePvcsPersistentVolumeClaimStatus) *UpdatePvcsPersistentVolumeClaim {
  s.Status = v
  return s
}

type UpdatePvcsPersistentVolumeClaimSpec struct {
  // {"en":"contains the desired access modes the volume should have", "zh_CN":"包含卷应具备的预期访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"a label query over volumes to consider for binding", "zh_CN":"在绑定时对卷进行选择所执行的标签查询"}
  Selector *UpdatePvcsMetaV1LabelSelector `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim", "zh_CN":"表示卷应拥有的最小资源。 如果启用了 RecoverVolumeExpansionFailure 功能特性，则允许用户指定这些资源要求， 此值必须低于之前的值，但必须高于申领的状态字段中记录的容量"}
  Resources *UpdatePvcsResourceRequirements `json:"resources,omitempty" xml:"resources,omitempty"`
  // {"en":"the binding reference to the PersistentVolume backing this claim", "zh_CN":"对此申领所对应的 PersistentVolume 的绑定引用"}
  VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
  // {"en":"the name of the StorageClass required by the claim", "zh_CN":"此申领所要求的 StorageClass 名称"}
  StorageClassName *string `json:"storageClassName,omitempty" xml:"storageClassName,omitempty"`
  // {"en":"defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec", "zh_CN":"定义申领需要哪种类别的卷。当申领规约中未包含此字段时，意味着取值为 Filesystem"}
  VolumeMode *string `json:"volumeMode,omitempty" xml:"volumeMode,omitempty"`
  // {"en":"dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (UpdatePvcsPersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource", "zh_CN":"dataSource 字段可用于二选一：- 现有的 VolumeSnapshot 对象（snapshot.storage.k8s.io/VolumeSnapshot）- 现有的 PVC (UpdatePvcsPersistentVolumeClaim)。如果制备器或外部控制器可以支持指定的数据源，则它将根据指定数据源的内容创建新的卷。 当 AnyVolumeDataSource 特性门控被启用时，dataSource 内容将被复制到 dataSourceRef， 当 dataSourceRef.namespace 未被指定时，dataSourceRef 内容将被复制到 dataSource。 如果名字空间被指定，则 dataSourceRef 不会被复制到 dataSource"}
  DataSource *UpdatePvcsTypedLocalObjectReference `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
}

func (s UpdatePvcsPersistentVolumeClaimSpec) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsPersistentVolumeClaimSpec) GoString() string {
  return s.String()
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetAccessModes(v []*string) *UpdatePvcsPersistentVolumeClaimSpec {
  s.AccessModes = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetSelector(v *UpdatePvcsMetaV1LabelSelector) *UpdatePvcsPersistentVolumeClaimSpec {
  s.Selector = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetResources(v *UpdatePvcsResourceRequirements) *UpdatePvcsPersistentVolumeClaimSpec {
  s.Resources = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetVolumeName(v string) *UpdatePvcsPersistentVolumeClaimSpec {
  s.VolumeName = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetStorageClassName(v string) *UpdatePvcsPersistentVolumeClaimSpec {
  s.StorageClassName = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetVolumeMode(v string) *UpdatePvcsPersistentVolumeClaimSpec {
  s.VolumeMode = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimSpec) SetDataSource(v *UpdatePvcsTypedLocalObjectReference) *UpdatePvcsPersistentVolumeClaimSpec {
  s.DataSource = v
  return s
}

type UpdatePvcsMetaV1LabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*UpdatePvcsLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s UpdatePvcsMetaV1LabelSelector) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsMetaV1LabelSelector) GoString() string {
  return s.String()
}

func (s *UpdatePvcsMetaV1LabelSelector) SetMatchLabels(v map[string]*string) *UpdatePvcsMetaV1LabelSelector {
  s.MatchLabels = v
  return s
}

func (s *UpdatePvcsMetaV1LabelSelector) SetMatchExpressions(v []*UpdatePvcsLabelSelectorRequirement) *UpdatePvcsMetaV1LabelSelector {
  s.MatchExpressions = v
  return s
}

type UpdatePvcsLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdatePvcsLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *UpdatePvcsLabelSelectorRequirement) SetKey(v string) *UpdatePvcsLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *UpdatePvcsLabelSelectorRequirement) SetOperator(v string) *UpdatePvcsLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *UpdatePvcsLabelSelectorRequirement) SetValues(v []*string) *UpdatePvcsLabelSelectorRequirement {
  s.Values = v
  return s
}

type UpdatePvcsResourceRequirements struct {
  // {"en":"describes the maximum amount of compute resources allowed", "zh_CN":"描述所允许的最大计算资源用量"}
  Limits map[string]*string `json:"limits,omitempty" xml:"limits,omitempty"`
  // {"en":"describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits", "zh_CN":"requests 描述所需的最小计算资源量。如果容器省略了 requests，但明确设定了 limits， 则 requests 默认值为 limits 值，否则为实现定义的值。请求不能超过限制"}
  Requests map[string]*string `json:"requests,omitempty" xml:"requests,omitempty"`
}

func (s UpdatePvcsResourceRequirements) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsResourceRequirements) GoString() string {
  return s.String()
}

func (s *UpdatePvcsResourceRequirements) SetLimits(v map[string]*string) *UpdatePvcsResourceRequirements {
  s.Limits = v
  return s
}

func (s *UpdatePvcsResourceRequirements) SetRequests(v map[string]*string) *UpdatePvcsResourceRequirements {
  s.Requests = v
  return s
}

type UpdatePvcsTypedLocalObjectReference struct {
  // {"en":"the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required", "zh_CN":"被引用资源的组。如果不指定 APIGroup，则指定的 Kind 必须在核心 API 组中。对于任何其它第三方类型，都需要 APIGroup"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
  // {"en":" the type of resource being referenced", "zh_CN":"被引用的资源的类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"the name of resource being referenced", "zh_CN":"被引用的资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdatePvcsTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *UpdatePvcsTypedLocalObjectReference) SetApiGroup(v string) *UpdatePvcsTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

func (s *UpdatePvcsTypedLocalObjectReference) SetKind(v string) *UpdatePvcsTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *UpdatePvcsTypedLocalObjectReference) SetName(v string) *UpdatePvcsTypedLocalObjectReference {
  s.Name = &v
  return s
}

type UpdatePvcsPersistentVolumeClaimStatus struct {
  // {"en":"represents the current phase of UpdatePvcsPersistentVolumeClaim", "zh_CN":"表示 UpdatePvcsPersistentVolumeClaim 的当前阶段"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
  // {"en":"contains the actual access modes the volume backing the PVC has", "zh_CN":"包含支持 PVC 的卷所具有的实际访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"represents the actual resources of the underlying volume", "zh_CN":"表示底层卷的实际资源"}
  Capacity map[string]*int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
  // {"en":"the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'", "zh_CN":"持久卷声明的当前的状况。 如果正在调整底层持久卷的大小，则状况将被设为 “ResizeStarted”"}
  Conditions []*UpdatePvcsPersistentVolumeClaimCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s UpdatePvcsPersistentVolumeClaimStatus) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsPersistentVolumeClaimStatus) GoString() string {
  return s.String()
}

func (s *UpdatePvcsPersistentVolumeClaimStatus) SetPhase(v string) *UpdatePvcsPersistentVolumeClaimStatus {
  s.Phase = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimStatus) SetAccessModes(v []*string) *UpdatePvcsPersistentVolumeClaimStatus {
  s.AccessModes = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimStatus) SetCapacity(v map[string]*int64) *UpdatePvcsPersistentVolumeClaimStatus {
  s.Capacity = v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimStatus) SetConditions(v []*UpdatePvcsPersistentVolumeClaimCondition) *UpdatePvcsPersistentVolumeClaimStatus {
  s.Conditions = v
  return s
}

type UpdatePvcsPersistentVolumeClaimCondition struct {
  // {"en":"type, required", "zh_CN":"类型，必需"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status, required", "zh_CN":"状态，必需"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized", "zh_CN":"reason 是唯一的，它应该是一个机器可理解的简短字符串，指明上次状况转换的原因。 如果它报告 “ResizeStarted”，则意味着正在调整底层持久卷的大小"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":" the human-readable message indicating details about last transition", "zh_CN":"人类可读的消息，指示有关上一次转换的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdatePvcsPersistentVolumeClaimCondition) String() string {
  return tea.Prettify(s)
}

func (s UpdatePvcsPersistentVolumeClaimCondition) GoString() string {
  return s.String()
}

func (s *UpdatePvcsPersistentVolumeClaimCondition) SetType(v string) *UpdatePvcsPersistentVolumeClaimCondition {
  s.Type = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimCondition) SetStatus(v string) *UpdatePvcsPersistentVolumeClaimCondition {
  s.Status = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimCondition) SetReason(v string) *UpdatePvcsPersistentVolumeClaimCondition {
  s.Reason = &v
  return s
}

func (s *UpdatePvcsPersistentVolumeClaimCondition) SetMessage(v string) *UpdatePvcsPersistentVolumeClaimCondition {
  s.Message = &v
  return s
}




type ListStorageClassRequest struct {
}

func (s ListStorageClassRequest) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassRequest) GoString() string {
  return s.String()
}

type ListStorageClassResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"storageClass list", "zh_CN":"storageClass列表"}
  Data []*ListStorageClassScInfo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s ListStorageClassResponse) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassResponse) GoString() string {
  return s.String()
}

func (s *ListStorageClassResponse) SetCode(v int64) *ListStorageClassResponse {
  s.Code = &v
  return s
}

func (s *ListStorageClassResponse) SetMsg(v string) *ListStorageClassResponse {
  s.Msg = &v
  return s
}

func (s *ListStorageClassResponse) SetRequestId(v string) *ListStorageClassResponse {
  s.RequestId = &v
  return s
}

func (s *ListStorageClassResponse) SetData(v []*ListStorageClassScInfo) *ListStorageClassResponse {
  s.Data = v
  return s
}

type ListStorageClassScInfo struct {
  // {"en":"storageClass name", "zh_CN":"storageClass名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"storageClass cn name", "zh_CN":"storageClass中文名称"}
  NameCn *string `json:"nameCn,omitempty" xml:"nameCn,omitempty" require:"true"`
}

func (s ListStorageClassScInfo) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassScInfo) GoString() string {
  return s.String()
}

func (s *ListStorageClassScInfo) SetName(v string) *ListStorageClassScInfo {
  s.Name = &v
  return s
}

func (s *ListStorageClassScInfo) SetNameCn(v string) *ListStorageClassScInfo {
  s.NameCn = &v
  return s
}

type ListStorageClassPaths struct {
}

func (s ListStorageClassPaths) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassPaths) GoString() string {
  return s.String()
}

type ListStorageClassParameters struct {
}

func (s ListStorageClassParameters) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassParameters) GoString() string {
  return s.String()
}

type ListStorageClassRequestHeader struct {
}

func (s ListStorageClassRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassRequestHeader) GoString() string {
  return s.String()
}

type ListStorageClassResponseHeader struct {
}

func (s ListStorageClassResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListStorageClassResponseHeader) GoString() string {
  return s.String()
}




type PagingPvcsRequest struct {
}

func (s PagingPvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsRequest) GoString() string {
  return s.String()
}

type PagingPvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc list", "zh_CN":"pvc列表"}
  Data *PagingPvcsPvcList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PagingPvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsResponse) GoString() string {
  return s.String()
}

func (s *PagingPvcsResponse) SetCode(v int64) *PagingPvcsResponse {
  s.Code = &v
  return s
}

func (s *PagingPvcsResponse) SetMsg(v string) *PagingPvcsResponse {
  s.Msg = &v
  return s
}

func (s *PagingPvcsResponse) SetRequestId(v string) *PagingPvcsResponse {
  s.RequestId = &v
  return s
}

func (s *PagingPvcsResponse) SetData(v *PagingPvcsPvcList) *PagingPvcsResponse {
  s.Data = v
  return s
}

type PagingPvcsPvcList struct {
  // {"en":"pvcs info", "zh_CN":"pvc信息"}
  Items []*PagingPvcsPvcInfo `json:"items,omitempty" xml:"items,omitempty" require:"true" type:"Repeated"`
  // {"en":"total size", "zh_CN":"总条数"}
  Total *int64 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s PagingPvcsPvcList) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsPvcList) GoString() string {
  return s.String()
}

func (s *PagingPvcsPvcList) SetItems(v []*PagingPvcsPvcInfo) *PagingPvcsPvcList {
  s.Items = v
  return s
}

func (s *PagingPvcsPvcList) SetTotal(v int64) *PagingPvcsPvcList {
  s.Total = &v
  return s
}

type PagingPvcsPvcInfo struct {
  // {"en":"pvcs name", "zh_CN":"pvc名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"namespace info", "zh_CN":"pvc namespace"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"pvc io type", "zh_CN":"pvc IO类型"}
  IOType *string `json:"iOType,omitempty" xml:"iOType,omitempty" require:"true"`
  // {"en":"pvc in every cluster info", "zh_CN":"pvc在各集群信息"}
  Clusters []*PagingPvcsCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"pvc workload", "zh_CN":"pvc绑定的负载"}
  WorkLoads []*PagingPvcsWorkLoad `json:"workLoads,omitempty" xml:"workLoads,omitempty" require:"true" type:"Repeated"`
  // {"en":"createTime", "zh_CN":"pvc创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"pvc storage", "zh_CN":"pvc存储大小"}
  Storage *int64 `json:"storage,omitempty" xml:"storage,omitempty" require:"true"`
  // {"en":"pvc storageClass", "zh_CN":"pvc storageClass"}
  StorageClass *string `json:"storageClass,omitempty" xml:"storageClass,omitempty" require:"true"`
}

func (s PagingPvcsPvcInfo) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsPvcInfo) GoString() string {
  return s.String()
}

func (s *PagingPvcsPvcInfo) SetName(v string) *PagingPvcsPvcInfo {
  s.Name = &v
  return s
}

func (s *PagingPvcsPvcInfo) SetNamespace(v string) *PagingPvcsPvcInfo {
  s.Namespace = &v
  return s
}

func (s *PagingPvcsPvcInfo) SetIOType(v string) *PagingPvcsPvcInfo {
  s.IOType = &v
  return s
}

func (s *PagingPvcsPvcInfo) SetClusters(v []*PagingPvcsCluster) *PagingPvcsPvcInfo {
  s.Clusters = v
  return s
}

func (s *PagingPvcsPvcInfo) SetWorkLoads(v []*PagingPvcsWorkLoad) *PagingPvcsPvcInfo {
  s.WorkLoads = v
  return s
}

func (s *PagingPvcsPvcInfo) SetCreateTime(v int64) *PagingPvcsPvcInfo {
  s.CreateTime = &v
  return s
}

func (s *PagingPvcsPvcInfo) SetStorage(v int64) *PagingPvcsPvcInfo {
  s.Storage = &v
  return s
}

func (s *PagingPvcsPvcInfo) SetStorageClass(v string) *PagingPvcsPvcInfo {
  s.StorageClass = &v
  return s
}

type PagingPvcsCluster struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"cluster cn name", "zh_CN":"集群中文名称"}
  NameCn *string `json:"nameCn,omitempty" xml:"nameCn,omitempty" require:"true"`
  // {"en":"pvc status", "zh_CN":"pvc在集群的状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s PagingPvcsCluster) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsCluster) GoString() string {
  return s.String()
}

func (s *PagingPvcsCluster) SetName(v string) *PagingPvcsCluster {
  s.Name = &v
  return s
}

func (s *PagingPvcsCluster) SetNameCn(v string) *PagingPvcsCluster {
  s.NameCn = &v
  return s
}

func (s *PagingPvcsCluster) SetStatus(v string) *PagingPvcsCluster {
  s.Status = &v
  return s
}

type PagingPvcsWorkLoad struct {
  // {"en":"workload kind", "zh_CN":"负载类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"workload name", "zh_CN":"负载名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"workload namespace", "zh_CN":"负载命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s PagingPvcsWorkLoad) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsWorkLoad) GoString() string {
  return s.String()
}

func (s *PagingPvcsWorkLoad) SetKind(v string) *PagingPvcsWorkLoad {
  s.Kind = &v
  return s
}

func (s *PagingPvcsWorkLoad) SetName(v string) *PagingPvcsWorkLoad {
  s.Name = &v
  return s
}

func (s *PagingPvcsWorkLoad) SetNamespace(v string) *PagingPvcsWorkLoad {
  s.Namespace = &v
  return s
}

type PagingPvcsPaths struct {
}

func (s PagingPvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsPaths) GoString() string {
  return s.String()
}

type PagingPvcsParameters struct {
  // {"en":"pvc storageclass", "zh_CN":"pvc storageclass，多个以逗号隔开"}
  Storageclass *string `json:"storageclass,omitempty" xml:"storageclass,omitempty" require:"true"`
  // {"en":"create by sts volumeClaimTemplates", "zh_CN":"是否由sts volumeClaimTemplates直接创建"}
  Createbyedge *string `json:"createbyedge,omitempty" xml:"createbyedge,omitempty" require:"true"`
}

func (s PagingPvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsParameters) GoString() string {
  return s.String()
}

func (s *PagingPvcsParameters) SetStorageclass(v string) *PagingPvcsParameters {
  s.Storageclass = &v
  return s
}

func (s *PagingPvcsParameters) SetCreatebyedge(v string) *PagingPvcsParameters {
  s.Createbyedge = &v
  return s
}

type PagingPvcsRequestHeader struct {
}

func (s PagingPvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsRequestHeader) GoString() string {
  return s.String()
}

type PagingPvcsResponseHeader struct {
}

func (s PagingPvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PagingPvcsResponseHeader) GoString() string {
  return s.String()
}




type PvcInNamespaceRequest struct {
}

func (s PvcInNamespaceRequest) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceRequest) GoString() string {
  return s.String()
}

type PvcInNamespaceResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc list", "zh_CN":"pvc列表"}
  Data []*PvcInNamespacePvcInfo `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s PvcInNamespaceResponse) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceResponse) GoString() string {
  return s.String()
}

func (s *PvcInNamespaceResponse) SetCode(v int64) *PvcInNamespaceResponse {
  s.Code = &v
  return s
}

func (s *PvcInNamespaceResponse) SetMsg(v string) *PvcInNamespaceResponse {
  s.Msg = &v
  return s
}

func (s *PvcInNamespaceResponse) SetRequestId(v string) *PvcInNamespaceResponse {
  s.RequestId = &v
  return s
}

func (s *PvcInNamespaceResponse) SetData(v []*PvcInNamespacePvcInfo) *PvcInNamespaceResponse {
  s.Data = v
  return s
}

type PvcInNamespacePvcInfo struct {
  // {"en":"pvcs name", "zh_CN":"pvc名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"namespace info", "zh_CN":"pvc namespace"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"pvc io type", "zh_CN":"pvc IO类型"}
  IOType *string `json:"iOType,omitempty" xml:"iOType,omitempty" require:"true"`
  // {"en":"pvc in every cluster info", "zh_CN":"pvc在各集群信息"}
  Clusters []*PvcInNamespaceCluster `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true" type:"Repeated"`
  // {"en":"pvc workload", "zh_CN":"pvc绑定的负载"}
  WorkLoads []*PvcInNamespaceWorkLoad `json:"workLoads,omitempty" xml:"workLoads,omitempty" require:"true" type:"Repeated"`
  // {"en":"createTime", "zh_CN":"pvc创建时间"}
  CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty" require:"true"`
  // {"en":"pvc storage", "zh_CN":"pvc存储大小"}
  Storage *int64 `json:"storage,omitempty" xml:"storage,omitempty" require:"true"`
  // {"en":"pvc storageClass", "zh_CN":"pvc storageClass"}
  StorageClass *string `json:"storageClass,omitempty" xml:"storageClass,omitempty" require:"true"`
}

func (s PvcInNamespacePvcInfo) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespacePvcInfo) GoString() string {
  return s.String()
}

func (s *PvcInNamespacePvcInfo) SetName(v string) *PvcInNamespacePvcInfo {
  s.Name = &v
  return s
}

func (s *PvcInNamespacePvcInfo) SetNamespace(v string) *PvcInNamespacePvcInfo {
  s.Namespace = &v
  return s
}

func (s *PvcInNamespacePvcInfo) SetIOType(v string) *PvcInNamespacePvcInfo {
  s.IOType = &v
  return s
}

func (s *PvcInNamespacePvcInfo) SetClusters(v []*PvcInNamespaceCluster) *PvcInNamespacePvcInfo {
  s.Clusters = v
  return s
}

func (s *PvcInNamespacePvcInfo) SetWorkLoads(v []*PvcInNamespaceWorkLoad) *PvcInNamespacePvcInfo {
  s.WorkLoads = v
  return s
}

func (s *PvcInNamespacePvcInfo) SetCreateTime(v int64) *PvcInNamespacePvcInfo {
  s.CreateTime = &v
  return s
}

func (s *PvcInNamespacePvcInfo) SetStorage(v int64) *PvcInNamespacePvcInfo {
  s.Storage = &v
  return s
}

func (s *PvcInNamespacePvcInfo) SetStorageClass(v string) *PvcInNamespacePvcInfo {
  s.StorageClass = &v
  return s
}

type PvcInNamespaceCluster struct {
  // {"en":"cluster name", "zh_CN":"集群名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"cluster cn name", "zh_CN":"集群中文名称"}
  NameCn *string `json:"nameCn,omitempty" xml:"nameCn,omitempty" require:"true"`
  // {"en":"pvc status", "zh_CN":"pvc在集群的状态"}
  Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s PvcInNamespaceCluster) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceCluster) GoString() string {
  return s.String()
}

func (s *PvcInNamespaceCluster) SetName(v string) *PvcInNamespaceCluster {
  s.Name = &v
  return s
}

func (s *PvcInNamespaceCluster) SetNameCn(v string) *PvcInNamespaceCluster {
  s.NameCn = &v
  return s
}

func (s *PvcInNamespaceCluster) SetStatus(v string) *PvcInNamespaceCluster {
  s.Status = &v
  return s
}

type PvcInNamespaceWorkLoad struct {
  // {"en":"workload kind", "zh_CN":"负载类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"workload name", "zh_CN":"负载名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"workload namespace", "zh_CN":"负载命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s PvcInNamespaceWorkLoad) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceWorkLoad) GoString() string {
  return s.String()
}

func (s *PvcInNamespaceWorkLoad) SetKind(v string) *PvcInNamespaceWorkLoad {
  s.Kind = &v
  return s
}

func (s *PvcInNamespaceWorkLoad) SetName(v string) *PvcInNamespaceWorkLoad {
  s.Name = &v
  return s
}

func (s *PvcInNamespaceWorkLoad) SetNamespace(v string) *PvcInNamespaceWorkLoad {
  s.Namespace = &v
  return s
}

type PvcInNamespacePaths struct {
  // {"en":"namespace info", "zh_CN":"pvc namespace"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s PvcInNamespacePaths) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespacePaths) GoString() string {
  return s.String()
}

func (s *PvcInNamespacePaths) SetNamespace(v string) *PvcInNamespacePaths {
  s.Namespace = &v
  return s
}

type PvcInNamespaceParameters struct {
  // {"en":"pvc storageclass,multiple separated by commas", "zh_CN":"pvc storageclass，多个以逗号隔开"}
  Storageclass *string `json:"storageclass,omitempty" xml:"storageclass,omitempty" require:"true"`
  // {"en":"filter has been used", "zh_CN":"过滤已被使用的pvc"}
  Filterused *string `json:"filterused,omitempty" xml:"filterused,omitempty" require:"true"`
}

func (s PvcInNamespaceParameters) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceParameters) GoString() string {
  return s.String()
}

func (s *PvcInNamespaceParameters) SetStorageclass(v string) *PvcInNamespaceParameters {
  s.Storageclass = &v
  return s
}

func (s *PvcInNamespaceParameters) SetFilterused(v string) *PvcInNamespaceParameters {
  s.Filterused = &v
  return s
}

type PvcInNamespaceRequestHeader struct {
}

func (s PvcInNamespaceRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceRequestHeader) GoString() string {
  return s.String()
}

type PvcInNamespaceResponseHeader struct {
}

func (s PvcInNamespaceResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PvcInNamespaceResponseHeader) GoString() string {
  return s.String()
}




type PutPatchPvcsRequest struct {
}

func (s PutPatchPvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsRequest) GoString() string {
  return s.String()
}

type PutPatchPvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc", "zh_CN":"pvc"}
  Data *PutPatchPvcsPersistentVolumeClaim `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PutPatchPvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsResponse) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsResponse) SetCode(v int64) *PutPatchPvcsResponse {
  s.Code = &v
  return s
}

func (s *PutPatchPvcsResponse) SetMsg(v string) *PutPatchPvcsResponse {
  s.Msg = &v
  return s
}

func (s *PutPatchPvcsResponse) SetRequestId(v string) *PutPatchPvcsResponse {
  s.RequestId = &v
  return s
}

func (s *PutPatchPvcsResponse) SetData(v *PutPatchPvcsPersistentVolumeClaim) *PutPatchPvcsResponse {
  s.Data = v
  return s
}

type PutPatchPvcsPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"deployment name", "zh_CN":"pvc 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s PutPatchPvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsPaths) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsPaths) SetNamespace(v string) *PutPatchPvcsPaths {
  s.Namespace = &v
  return s
}

func (s *PutPatchPvcsPaths) SetName(v string) *PutPatchPvcsPaths {
  s.Name = &v
  return s
}

type PutPatchPvcsParameters struct {
}

func (s PutPatchPvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsParameters) GoString() string {
  return s.String()
}

type PutPatchPvcsRequestHeader struct {
}

func (s PutPatchPvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsRequestHeader) GoString() string {
  return s.String()
}

type PutPatchPvcsResponseHeader struct {
}

func (s PutPatchPvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsResponseHeader) GoString() string {
  return s.String()
}

type PutPatchPvcsObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
  OwnerReferences []*PutPatchPvcsOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
}

func (s PutPatchPvcsObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsObjectMeta) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsObjectMeta) SetName(v string) *PutPatchPvcsObjectMeta {
  s.Name = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetGenerateName(v string) *PutPatchPvcsObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetNamespace(v string) *PutPatchPvcsObjectMeta {
  s.Namespace = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetSelfLink(v string) *PutPatchPvcsObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetUid(v string) *PutPatchPvcsObjectMeta {
  s.Uid = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetResourceVersion(v string) *PutPatchPvcsObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetGeneration(v int64) *PutPatchPvcsObjectMeta {
  s.Generation = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetCreationTimestamp(v string) *PutPatchPvcsObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetDeletionTimestamp(v string) *PutPatchPvcsObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetDeletionGracePeriodSeconds(v int64) *PutPatchPvcsObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetLabels(v map[string]*string) *PutPatchPvcsObjectMeta {
  s.Labels = v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetAnnotations(v map[string]*string) *PutPatchPvcsObjectMeta {
  s.Annotations = v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetOwnerReferences(v []*PutPatchPvcsOwnerReference) *PutPatchPvcsObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetFinalizers(v []*string) *PutPatchPvcsObjectMeta {
  s.Finalizers = v
  return s
}

func (s *PutPatchPvcsObjectMeta) SetClusterName(v string) *PutPatchPvcsObjectMeta {
  s.ClusterName = &v
  return s
}

type PutPatchPvcsOwnerReference struct {
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

func (s PutPatchPvcsOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsOwnerReference) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsOwnerReference) SetApiVersion(v string) *PutPatchPvcsOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchPvcsOwnerReference) SetKind(v string) *PutPatchPvcsOwnerReference {
  s.Kind = &v
  return s
}

func (s *PutPatchPvcsOwnerReference) SetName(v string) *PutPatchPvcsOwnerReference {
  s.Name = &v
  return s
}

func (s *PutPatchPvcsOwnerReference) SetUid(v string) *PutPatchPvcsOwnerReference {
  s.Uid = &v
  return s
}

func (s *PutPatchPvcsOwnerReference) SetController(v bool) *PutPatchPvcsOwnerReference {
  s.Controller = &v
  return s
}

func (s *PutPatchPvcsOwnerReference) SetBlockOwnerDeletion(v bool) *PutPatchPvcsOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type PutPatchPvcsPersistentVolumeClaim struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *PutPatchPvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *PutPatchPvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"represents the current information/status of a persistent volume claim. Read-only", "zh_CN":"表示一个持久卷申领的当前信息/状态。只读"}
  Status *PutPatchPvcsPersistentVolumeClaimStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutPatchPvcsPersistentVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsPersistentVolumeClaim) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsPersistentVolumeClaim) SetApiVersion(v string) *PutPatchPvcsPersistentVolumeClaim {
  s.ApiVersion = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaim) SetKind(v string) *PutPatchPvcsPersistentVolumeClaim {
  s.Kind = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaim) SetMetadata(v *PutPatchPvcsObjectMeta) *PutPatchPvcsPersistentVolumeClaim {
  s.Metadata = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaim) SetSpec(v *PutPatchPvcsPersistentVolumeClaimSpec) *PutPatchPvcsPersistentVolumeClaim {
  s.Spec = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaim) SetStatus(v *PutPatchPvcsPersistentVolumeClaimStatus) *PutPatchPvcsPersistentVolumeClaim {
  s.Status = v
  return s
}

type PutPatchPvcsPersistentVolumeClaimSpec struct {
  // {"en":"contains the desired access modes the volume should have", "zh_CN":"包含卷应具备的预期访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"a label query over volumes to consider for binding", "zh_CN":"在绑定时对卷进行选择所执行的标签查询"}
  Selector *PutPatchPvcsMetaV1LabelSelector `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim", "zh_CN":"表示卷应拥有的最小资源。 如果启用了 RecoverVolumeExpansionFailure 功能特性，则允许用户指定这些资源要求， 此值必须低于之前的值，但必须高于申领的状态字段中记录的容量"}
  Resources *PutPatchPvcsResourceRequirements `json:"resources,omitempty" xml:"resources,omitempty"`
  // {"en":"the binding reference to the PersistentVolume backing this claim", "zh_CN":"对此申领所对应的 PersistentVolume 的绑定引用"}
  VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
  // {"en":"the name of the StorageClass required by the claim", "zh_CN":"此申领所要求的 StorageClass 名称"}
  StorageClassName *string `json:"storageClassName,omitempty" xml:"storageClassName,omitempty"`
  // {"en":"defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec", "zh_CN":"定义申领需要哪种类别的卷。当申领规约中未包含此字段时，意味着取值为 Filesystem"}
  VolumeMode *string `json:"volumeMode,omitempty" xml:"volumeMode,omitempty"`
  // {"en":"dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PutPatchPvcsPersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource", "zh_CN":"dataSource 字段可用于二选一：- 现有的 VolumeSnapshot 对象（snapshot.storage.k8s.io/VolumeSnapshot）- 现有的 PVC (PutPatchPvcsPersistentVolumeClaim)。如果制备器或外部控制器可以支持指定的数据源，则它将根据指定数据源的内容创建新的卷。 当 AnyVolumeDataSource 特性门控被启用时，dataSource 内容将被复制到 dataSourceRef， 当 dataSourceRef.namespace 未被指定时，dataSourceRef 内容将被复制到 dataSource。 如果名字空间被指定，则 dataSourceRef 不会被复制到 dataSource"}
  DataSource *PutPatchPvcsTypedLocalObjectReference `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
}

func (s PutPatchPvcsPersistentVolumeClaimSpec) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsPersistentVolumeClaimSpec) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetAccessModes(v []*string) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.AccessModes = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetSelector(v *PutPatchPvcsMetaV1LabelSelector) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.Selector = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetResources(v *PutPatchPvcsResourceRequirements) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.Resources = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetVolumeName(v string) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.VolumeName = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetStorageClassName(v string) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.StorageClassName = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetVolumeMode(v string) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.VolumeMode = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimSpec) SetDataSource(v *PutPatchPvcsTypedLocalObjectReference) *PutPatchPvcsPersistentVolumeClaimSpec {
  s.DataSource = v
  return s
}

type PutPatchPvcsMetaV1LabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*PutPatchPvcsLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s PutPatchPvcsMetaV1LabelSelector) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsMetaV1LabelSelector) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsMetaV1LabelSelector) SetMatchLabels(v map[string]*string) *PutPatchPvcsMetaV1LabelSelector {
  s.MatchLabels = v
  return s
}

func (s *PutPatchPvcsMetaV1LabelSelector) SetMatchExpressions(v []*PutPatchPvcsLabelSelectorRequirement) *PutPatchPvcsMetaV1LabelSelector {
  s.MatchExpressions = v
  return s
}

type PutPatchPvcsLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s PutPatchPvcsLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsLabelSelectorRequirement) SetKey(v string) *PutPatchPvcsLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *PutPatchPvcsLabelSelectorRequirement) SetOperator(v string) *PutPatchPvcsLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *PutPatchPvcsLabelSelectorRequirement) SetValues(v []*string) *PutPatchPvcsLabelSelectorRequirement {
  s.Values = v
  return s
}

type PutPatchPvcsResourceRequirements struct {
  // {"en":"describes the maximum amount of compute resources allowed", "zh_CN":"描述所允许的最大计算资源用量"}
  Limits map[string]*string `json:"limits,omitempty" xml:"limits,omitempty"`
  // {"en":"describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits", "zh_CN":"requests 描述所需的最小计算资源量。如果容器省略了 requests，但明确设定了 limits， 则 requests 默认值为 limits 值，否则为实现定义的值。请求不能超过限制"}
  Requests map[string]*string `json:"requests,omitempty" xml:"requests,omitempty"`
}

func (s PutPatchPvcsResourceRequirements) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsResourceRequirements) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsResourceRequirements) SetLimits(v map[string]*string) *PutPatchPvcsResourceRequirements {
  s.Limits = v
  return s
}

func (s *PutPatchPvcsResourceRequirements) SetRequests(v map[string]*string) *PutPatchPvcsResourceRequirements {
  s.Requests = v
  return s
}

type PutPatchPvcsTypedLocalObjectReference struct {
  // {"en":"the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required", "zh_CN":"被引用资源的组。如果不指定 APIGroup，则指定的 Kind 必须在核心 API 组中。对于任何其它第三方类型，都需要 APIGroup"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
  // {"en":" the type of resource being referenced", "zh_CN":"被引用的资源的类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"the name of resource being referenced", "zh_CN":"被引用的资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PutPatchPvcsTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsTypedLocalObjectReference) SetApiGroup(v string) *PutPatchPvcsTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

func (s *PutPatchPvcsTypedLocalObjectReference) SetKind(v string) *PutPatchPvcsTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *PutPatchPvcsTypedLocalObjectReference) SetName(v string) *PutPatchPvcsTypedLocalObjectReference {
  s.Name = &v
  return s
}

type PutPatchPvcsPersistentVolumeClaimStatus struct {
  // {"en":"represents the current phase of PutPatchPvcsPersistentVolumeClaim", "zh_CN":"表示 PutPatchPvcsPersistentVolumeClaim 的当前阶段"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
  // {"en":"contains the actual access modes the volume backing the PVC has", "zh_CN":"包含支持 PVC 的卷所具有的实际访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"represents the actual resources of the underlying volume", "zh_CN":"表示底层卷的实际资源"}
  Capacity map[string]*int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
  // {"en":"the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'", "zh_CN":"持久卷声明的当前的状况。 如果正在调整底层持久卷的大小，则状况将被设为 “ResizeStarted”"}
  Conditions []*PutPatchPvcsPersistentVolumeClaimCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s PutPatchPvcsPersistentVolumeClaimStatus) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsPersistentVolumeClaimStatus) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsPersistentVolumeClaimStatus) SetPhase(v string) *PutPatchPvcsPersistentVolumeClaimStatus {
  s.Phase = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimStatus) SetAccessModes(v []*string) *PutPatchPvcsPersistentVolumeClaimStatus {
  s.AccessModes = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimStatus) SetCapacity(v map[string]*int64) *PutPatchPvcsPersistentVolumeClaimStatus {
  s.Capacity = v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimStatus) SetConditions(v []*PutPatchPvcsPersistentVolumeClaimCondition) *PutPatchPvcsPersistentVolumeClaimStatus {
  s.Conditions = v
  return s
}

type PutPatchPvcsPersistentVolumeClaimCondition struct {
  // {"en":"type, required", "zh_CN":"类型，必需"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status, required", "zh_CN":"状态，必需"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized", "zh_CN":"reason 是唯一的，它应该是一个机器可理解的简短字符串，指明上次状况转换的原因。 如果它报告 “ResizeStarted”，则意味着正在调整底层持久卷的大小"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":" the human-readable message indicating details about last transition", "zh_CN":"人类可读的消息，指示有关上一次转换的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s PutPatchPvcsPersistentVolumeClaimCondition) String() string {
  return tea.Prettify(s)
}

func (s PutPatchPvcsPersistentVolumeClaimCondition) GoString() string {
  return s.String()
}

func (s *PutPatchPvcsPersistentVolumeClaimCondition) SetType(v string) *PutPatchPvcsPersistentVolumeClaimCondition {
  s.Type = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimCondition) SetStatus(v string) *PutPatchPvcsPersistentVolumeClaimCondition {
  s.Status = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimCondition) SetReason(v string) *PutPatchPvcsPersistentVolumeClaimCondition {
  s.Reason = &v
  return s
}

func (s *PutPatchPvcsPersistentVolumeClaimCondition) SetMessage(v string) *PutPatchPvcsPersistentVolumeClaimCondition {
  s.Message = &v
  return s
}




type GetPvcsRequest struct {
}

func (s GetPvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsRequest) GoString() string {
  return s.String()
}

type GetPvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc", "zh_CN":"pvc"}
  Data *GetPvcsPersistentVolumeClaim `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetPvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsResponse) GoString() string {
  return s.String()
}

func (s *GetPvcsResponse) SetCode(v int64) *GetPvcsResponse {
  s.Code = &v
  return s
}

func (s *GetPvcsResponse) SetMsg(v string) *GetPvcsResponse {
  s.Msg = &v
  return s
}

func (s *GetPvcsResponse) SetRequestId(v string) *GetPvcsResponse {
  s.RequestId = &v
  return s
}

func (s *GetPvcsResponse) SetData(v *GetPvcsPersistentVolumeClaim) *GetPvcsResponse {
  s.Data = v
  return s
}

type GetPvcsPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"pvc name", "zh_CN":"pvc 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s GetPvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsPaths) GoString() string {
  return s.String()
}

func (s *GetPvcsPaths) SetNamespace(v string) *GetPvcsPaths {
  s.Namespace = &v
  return s
}

func (s *GetPvcsPaths) SetName(v string) *GetPvcsPaths {
  s.Name = &v
  return s
}

type GetPvcsParameters struct {
}

func (s GetPvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsParameters) GoString() string {
  return s.String()
}

type GetPvcsRequestHeader struct {
}

func (s GetPvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsRequestHeader) GoString() string {
  return s.String()
}

type GetPvcsResponseHeader struct {
}

func (s GetPvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsResponseHeader) GoString() string {
  return s.String()
}

type GetPvcsObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
  OwnerReferences []*GetPvcsOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
}

func (s GetPvcsObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsObjectMeta) GoString() string {
  return s.String()
}

func (s *GetPvcsObjectMeta) SetName(v string) *GetPvcsObjectMeta {
  s.Name = &v
  return s
}

func (s *GetPvcsObjectMeta) SetGenerateName(v string) *GetPvcsObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *GetPvcsObjectMeta) SetNamespace(v string) *GetPvcsObjectMeta {
  s.Namespace = &v
  return s
}

func (s *GetPvcsObjectMeta) SetSelfLink(v string) *GetPvcsObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *GetPvcsObjectMeta) SetUid(v string) *GetPvcsObjectMeta {
  s.Uid = &v
  return s
}

func (s *GetPvcsObjectMeta) SetResourceVersion(v string) *GetPvcsObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *GetPvcsObjectMeta) SetGeneration(v int64) *GetPvcsObjectMeta {
  s.Generation = &v
  return s
}

func (s *GetPvcsObjectMeta) SetCreationTimestamp(v string) *GetPvcsObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *GetPvcsObjectMeta) SetDeletionTimestamp(v string) *GetPvcsObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *GetPvcsObjectMeta) SetDeletionGracePeriodSeconds(v int64) *GetPvcsObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *GetPvcsObjectMeta) SetLabels(v map[string]*string) *GetPvcsObjectMeta {
  s.Labels = v
  return s
}

func (s *GetPvcsObjectMeta) SetAnnotations(v map[string]*string) *GetPvcsObjectMeta {
  s.Annotations = v
  return s
}

func (s *GetPvcsObjectMeta) SetOwnerReferences(v []*GetPvcsOwnerReference) *GetPvcsObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *GetPvcsObjectMeta) SetFinalizers(v []*string) *GetPvcsObjectMeta {
  s.Finalizers = v
  return s
}

func (s *GetPvcsObjectMeta) SetClusterName(v string) *GetPvcsObjectMeta {
  s.ClusterName = &v
  return s
}

type GetPvcsOwnerReference struct {
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

func (s GetPvcsOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsOwnerReference) GoString() string {
  return s.String()
}

func (s *GetPvcsOwnerReference) SetApiVersion(v string) *GetPvcsOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *GetPvcsOwnerReference) SetKind(v string) *GetPvcsOwnerReference {
  s.Kind = &v
  return s
}

func (s *GetPvcsOwnerReference) SetName(v string) *GetPvcsOwnerReference {
  s.Name = &v
  return s
}

func (s *GetPvcsOwnerReference) SetUid(v string) *GetPvcsOwnerReference {
  s.Uid = &v
  return s
}

func (s *GetPvcsOwnerReference) SetController(v bool) *GetPvcsOwnerReference {
  s.Controller = &v
  return s
}

func (s *GetPvcsOwnerReference) SetBlockOwnerDeletion(v bool) *GetPvcsOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type GetPvcsPersistentVolumeClaim struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *GetPvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *GetPvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"represents the current information/status of a persistent volume claim. Read-only", "zh_CN":"表示一个持久卷申领的当前信息/状态。只读"}
  Status *GetPvcsPersistentVolumeClaimStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPvcsPersistentVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsPersistentVolumeClaim) GoString() string {
  return s.String()
}

func (s *GetPvcsPersistentVolumeClaim) SetApiVersion(v string) *GetPvcsPersistentVolumeClaim {
  s.ApiVersion = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaim) SetKind(v string) *GetPvcsPersistentVolumeClaim {
  s.Kind = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaim) SetMetadata(v *GetPvcsObjectMeta) *GetPvcsPersistentVolumeClaim {
  s.Metadata = v
  return s
}

func (s *GetPvcsPersistentVolumeClaim) SetSpec(v *GetPvcsPersistentVolumeClaimSpec) *GetPvcsPersistentVolumeClaim {
  s.Spec = v
  return s
}

func (s *GetPvcsPersistentVolumeClaim) SetStatus(v *GetPvcsPersistentVolumeClaimStatus) *GetPvcsPersistentVolumeClaim {
  s.Status = v
  return s
}

type GetPvcsPersistentVolumeClaimSpec struct {
  // {"en":"contains the desired access modes the volume should have", "zh_CN":"包含卷应具备的预期访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"a label query over volumes to consider for binding", "zh_CN":"在绑定时对卷进行选择所执行的标签查询"}
  Selector *GetPvcsMetaV1LabelSelector `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim", "zh_CN":"表示卷应拥有的最小资源。 如果启用了 RecoverVolumeExpansionFailure 功能特性，则允许用户指定这些资源要求， 此值必须低于之前的值，但必须高于申领的状态字段中记录的容量"}
  Resources *GetPvcsResourceRequirements `json:"resources,omitempty" xml:"resources,omitempty"`
  // {"en":"the binding reference to the PersistentVolume backing this claim", "zh_CN":"对此申领所对应的 PersistentVolume 的绑定引用"}
  VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
  // {"en":"the name of the StorageClass required by the claim", "zh_CN":"此申领所要求的 StorageClass 名称"}
  StorageClassName *string `json:"storageClassName,omitempty" xml:"storageClassName,omitempty"`
  // {"en":"defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec", "zh_CN":"定义申领需要哪种类别的卷。当申领规约中未包含此字段时，意味着取值为 Filesystem"}
  VolumeMode *string `json:"volumeMode,omitempty" xml:"volumeMode,omitempty"`
  // {"en":"dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (GetPvcsPersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource", "zh_CN":"dataSource 字段可用于二选一：- 现有的 VolumeSnapshot 对象（snapshot.storage.k8s.io/VolumeSnapshot）- 现有的 PVC (GetPvcsPersistentVolumeClaim)。如果制备器或外部控制器可以支持指定的数据源，则它将根据指定数据源的内容创建新的卷。 当 AnyVolumeDataSource 特性门控被启用时，dataSource 内容将被复制到 dataSourceRef， 当 dataSourceRef.namespace 未被指定时，dataSourceRef 内容将被复制到 dataSource。 如果名字空间被指定，则 dataSourceRef 不会被复制到 dataSource"}
  DataSource *GetPvcsTypedLocalObjectReference `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
}

func (s GetPvcsPersistentVolumeClaimSpec) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsPersistentVolumeClaimSpec) GoString() string {
  return s.String()
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetAccessModes(v []*string) *GetPvcsPersistentVolumeClaimSpec {
  s.AccessModes = v
  return s
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetSelector(v *GetPvcsMetaV1LabelSelector) *GetPvcsPersistentVolumeClaimSpec {
  s.Selector = v
  return s
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetResources(v *GetPvcsResourceRequirements) *GetPvcsPersistentVolumeClaimSpec {
  s.Resources = v
  return s
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetVolumeName(v string) *GetPvcsPersistentVolumeClaimSpec {
  s.VolumeName = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetStorageClassName(v string) *GetPvcsPersistentVolumeClaimSpec {
  s.StorageClassName = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetVolumeMode(v string) *GetPvcsPersistentVolumeClaimSpec {
  s.VolumeMode = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimSpec) SetDataSource(v *GetPvcsTypedLocalObjectReference) *GetPvcsPersistentVolumeClaimSpec {
  s.DataSource = v
  return s
}

type GetPvcsMetaV1LabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*GetPvcsLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s GetPvcsMetaV1LabelSelector) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsMetaV1LabelSelector) GoString() string {
  return s.String()
}

func (s *GetPvcsMetaV1LabelSelector) SetMatchLabels(v map[string]*string) *GetPvcsMetaV1LabelSelector {
  s.MatchLabels = v
  return s
}

func (s *GetPvcsMetaV1LabelSelector) SetMatchExpressions(v []*GetPvcsLabelSelectorRequirement) *GetPvcsMetaV1LabelSelector {
  s.MatchExpressions = v
  return s
}

type GetPvcsLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetPvcsLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *GetPvcsLabelSelectorRequirement) SetKey(v string) *GetPvcsLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *GetPvcsLabelSelectorRequirement) SetOperator(v string) *GetPvcsLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *GetPvcsLabelSelectorRequirement) SetValues(v []*string) *GetPvcsLabelSelectorRequirement {
  s.Values = v
  return s
}

type GetPvcsResourceRequirements struct {
  // {"en":"describes the maximum amount of compute resources allowed", "zh_CN":"描述所允许的最大计算资源用量"}
  Limits map[string]*string `json:"limits,omitempty" xml:"limits,omitempty"`
  // {"en":"describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits", "zh_CN":"requests 描述所需的最小计算资源量。如果容器省略了 requests，但明确设定了 limits， 则 requests 默认值为 limits 值，否则为实现定义的值。请求不能超过限制"}
  Requests map[string]*string `json:"requests,omitempty" xml:"requests,omitempty"`
}

func (s GetPvcsResourceRequirements) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsResourceRequirements) GoString() string {
  return s.String()
}

func (s *GetPvcsResourceRequirements) SetLimits(v map[string]*string) *GetPvcsResourceRequirements {
  s.Limits = v
  return s
}

func (s *GetPvcsResourceRequirements) SetRequests(v map[string]*string) *GetPvcsResourceRequirements {
  s.Requests = v
  return s
}

type GetPvcsTypedLocalObjectReference struct {
  // {"en":"the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required", "zh_CN":"被引用资源的组。如果不指定 APIGroup，则指定的 Kind 必须在核心 API 组中。对于任何其它第三方类型，都需要 APIGroup"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
  // {"en":" the type of resource being referenced", "zh_CN":"被引用的资源的类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"the name of resource being referenced", "zh_CN":"被引用的资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPvcsTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *GetPvcsTypedLocalObjectReference) SetApiGroup(v string) *GetPvcsTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

func (s *GetPvcsTypedLocalObjectReference) SetKind(v string) *GetPvcsTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *GetPvcsTypedLocalObjectReference) SetName(v string) *GetPvcsTypedLocalObjectReference {
  s.Name = &v
  return s
}

type GetPvcsPersistentVolumeClaimStatus struct {
  // {"en":"represents the current phase of GetPvcsPersistentVolumeClaim", "zh_CN":"表示 GetPvcsPersistentVolumeClaim 的当前阶段"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
  // {"en":"contains the actual access modes the volume backing the PVC has", "zh_CN":"包含支持 PVC 的卷所具有的实际访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"represents the actual resources of the underlying volume", "zh_CN":"表示底层卷的实际资源"}
  Capacity map[string]*int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
  // {"en":"the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'", "zh_CN":"持久卷声明的当前的状况。 如果正在调整底层持久卷的大小，则状况将被设为 “ResizeStarted”"}
  Conditions []*GetPvcsPersistentVolumeClaimCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s GetPvcsPersistentVolumeClaimStatus) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsPersistentVolumeClaimStatus) GoString() string {
  return s.String()
}

func (s *GetPvcsPersistentVolumeClaimStatus) SetPhase(v string) *GetPvcsPersistentVolumeClaimStatus {
  s.Phase = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimStatus) SetAccessModes(v []*string) *GetPvcsPersistentVolumeClaimStatus {
  s.AccessModes = v
  return s
}

func (s *GetPvcsPersistentVolumeClaimStatus) SetCapacity(v map[string]*int64) *GetPvcsPersistentVolumeClaimStatus {
  s.Capacity = v
  return s
}

func (s *GetPvcsPersistentVolumeClaimStatus) SetConditions(v []*GetPvcsPersistentVolumeClaimCondition) *GetPvcsPersistentVolumeClaimStatus {
  s.Conditions = v
  return s
}

type GetPvcsPersistentVolumeClaimCondition struct {
  // {"en":"type, required", "zh_CN":"类型，必需"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status, required", "zh_CN":"状态，必需"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized", "zh_CN":"reason 是唯一的，它应该是一个机器可理解的简短字符串，指明上次状况转换的原因。 如果它报告 “ResizeStarted”，则意味着正在调整底层持久卷的大小"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":" the human-readable message indicating details about last transition", "zh_CN":"人类可读的消息，指示有关上一次转换的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetPvcsPersistentVolumeClaimCondition) String() string {
  return tea.Prettify(s)
}

func (s GetPvcsPersistentVolumeClaimCondition) GoString() string {
  return s.String()
}

func (s *GetPvcsPersistentVolumeClaimCondition) SetType(v string) *GetPvcsPersistentVolumeClaimCondition {
  s.Type = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimCondition) SetStatus(v string) *GetPvcsPersistentVolumeClaimCondition {
  s.Status = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimCondition) SetReason(v string) *GetPvcsPersistentVolumeClaimCondition {
  s.Reason = &v
  return s
}

func (s *GetPvcsPersistentVolumeClaimCondition) SetMessage(v string) *GetPvcsPersistentVolumeClaimCondition {
  s.Message = &v
  return s
}




type ListPvcsRequest struct {
}

func (s ListPvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsRequest) GoString() string {
  return s.String()
}

type ListPvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc list", "zh_CN":"pvc列表"}
  Data *ListPvcsPvcList `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ListPvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsResponse) GoString() string {
  return s.String()
}

func (s *ListPvcsResponse) SetCode(v int64) *ListPvcsResponse {
  s.Code = &v
  return s
}

func (s *ListPvcsResponse) SetMsg(v string) *ListPvcsResponse {
  s.Msg = &v
  return s
}

func (s *ListPvcsResponse) SetRequestId(v string) *ListPvcsResponse {
  s.RequestId = &v
  return s
}

func (s *ListPvcsResponse) SetData(v *ListPvcsPvcList) *ListPvcsResponse {
  s.Data = v
  return s
}

type ListPvcsPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s ListPvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsPaths) GoString() string {
  return s.String()
}

func (s *ListPvcsPaths) SetNamespace(v string) *ListPvcsPaths {
  s.Namespace = &v
  return s
}

type ListPvcsParameters struct {
  // {"en":"The name of pvc", "zh_CN":"pvc 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"labelSelector", "zh_CN":"labelSelector"}
  LabelSelector *string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty"`
}

func (s ListPvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsParameters) GoString() string {
  return s.String()
}

func (s *ListPvcsParameters) SetName(v string) *ListPvcsParameters {
  s.Name = &v
  return s
}

func (s *ListPvcsParameters) SetLabelSelector(v string) *ListPvcsParameters {
  s.LabelSelector = &v
  return s
}

type ListPvcsRequestHeader struct {
}

func (s ListPvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsRequestHeader) GoString() string {
  return s.String()
}

type ListPvcsResponseHeader struct {
}

func (s ListPvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsResponseHeader) GoString() string {
  return s.String()
}

type ListPvcsPvcList struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"Standard list metadata", "zh_CN":"标准列表元数据"}
  Metadata *ListPvcsListMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"List of ListPvcsPersistentVolumeClaim", "zh_CN":"ListPvcsPersistentVolumeClaim 列表"}
  Items []*ListPvcsPersistentVolumeClaim `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListPvcsPvcList) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsPvcList) GoString() string {
  return s.String()
}

func (s *ListPvcsPvcList) SetApiVersion(v string) *ListPvcsPvcList {
  s.ApiVersion = &v
  return s
}

func (s *ListPvcsPvcList) SetKind(v string) *ListPvcsPvcList {
  s.Kind = &v
  return s
}

func (s *ListPvcsPvcList) SetMetadata(v *ListPvcsListMeta) *ListPvcsPvcList {
  s.Metadata = v
  return s
}

func (s *ListPvcsPvcList) SetItems(v []*ListPvcsPersistentVolumeClaim) *ListPvcsPvcList {
  s.Items = v
  return s
}

type ListPvcsListMeta struct {
  // {"en":"Deprecated: selfLink is a legacy read-only field that is no longer populated by the system", "zh_CN":"selfLink 表示此对象的 URL，由系统填充，只读。已弃用：selfLink 是一个遗留的只读字段，不再由系统填充。"}
  SelfLink *string `json:"selfLink,omitempty" xml:"selfLink,omitempty"`
  // {"en":"String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only", "zh_CN":"标识该对象的服务器内部版本的字符串，客户端可以用该字段来确定对象何时被更改。 该值对客户端是不透明的，并且应该原样传回给服务器。该值由系统填充，只读"}
  ResourceVersion *string `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
  // {"en":"continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message", "zh_CN":"如果用户对返回的条目数量设置了限制，则 continue 可能被设置，表示服务器有更多可用的数据。 该值是不透明的，可用于向提供此列表服务的端点发出另一个请求，以检索下一组可用的对象。 如果服务器配置已更改或时间已过去几分钟，则可能无法继续提供一致的列表。 除非你在错误消息中收到此令牌（token），否则使用此 continue 值时返回的 resourceVersion 字段应该和第一个响应中的值是相同的"}
  Continue *string `json:"continue,omitempty" xml:"continue,omitempty"`
  // {"en":"remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is estimating the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact", "zh_CN":"remainingItemCount 是列表中未包含在此列表响应中的后续项目的数量。 如果列表请求包含标签或字段选择器，则剩余项目的数量是未知的，并且在序列化期间该字段将保持未设置和省略。 如果列表是完整的（因为它没有分块或者这是最后一个块），那么就没有剩余的项目，并且在序列化过程中该字段将保持未设置和省略。 早于 v1.15 的服务器不设置此字段。remainingItemCount 的预期用途是估计集合的大小。 客户端不应依赖于设置准确的 remainingItemCount"}
  RemainingItemCount *int64 `json:"remainingItemCount,omitempty" xml:"remainingItemCount,omitempty"`
}

func (s ListPvcsListMeta) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsListMeta) GoString() string {
  return s.String()
}

func (s *ListPvcsListMeta) SetSelfLink(v string) *ListPvcsListMeta {
  s.SelfLink = &v
  return s
}

func (s *ListPvcsListMeta) SetResourceVersion(v string) *ListPvcsListMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListPvcsListMeta) SetContinue(v string) *ListPvcsListMeta {
  s.Continue = &v
  return s
}

func (s *ListPvcsListMeta) SetRemainingItemCount(v int64) *ListPvcsListMeta {
  s.RemainingItemCount = &v
  return s
}

type ListPvcsObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
  OwnerReferences []*ListPvcsOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
}

func (s ListPvcsObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsObjectMeta) GoString() string {
  return s.String()
}

func (s *ListPvcsObjectMeta) SetName(v string) *ListPvcsObjectMeta {
  s.Name = &v
  return s
}

func (s *ListPvcsObjectMeta) SetGenerateName(v string) *ListPvcsObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *ListPvcsObjectMeta) SetNamespace(v string) *ListPvcsObjectMeta {
  s.Namespace = &v
  return s
}

func (s *ListPvcsObjectMeta) SetSelfLink(v string) *ListPvcsObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *ListPvcsObjectMeta) SetUid(v string) *ListPvcsObjectMeta {
  s.Uid = &v
  return s
}

func (s *ListPvcsObjectMeta) SetResourceVersion(v string) *ListPvcsObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *ListPvcsObjectMeta) SetGeneration(v int64) *ListPvcsObjectMeta {
  s.Generation = &v
  return s
}

func (s *ListPvcsObjectMeta) SetCreationTimestamp(v string) *ListPvcsObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *ListPvcsObjectMeta) SetDeletionTimestamp(v string) *ListPvcsObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *ListPvcsObjectMeta) SetDeletionGracePeriodSeconds(v int64) *ListPvcsObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *ListPvcsObjectMeta) SetLabels(v map[string]*string) *ListPvcsObjectMeta {
  s.Labels = v
  return s
}

func (s *ListPvcsObjectMeta) SetAnnotations(v map[string]*string) *ListPvcsObjectMeta {
  s.Annotations = v
  return s
}

func (s *ListPvcsObjectMeta) SetOwnerReferences(v []*ListPvcsOwnerReference) *ListPvcsObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *ListPvcsObjectMeta) SetFinalizers(v []*string) *ListPvcsObjectMeta {
  s.Finalizers = v
  return s
}

func (s *ListPvcsObjectMeta) SetClusterName(v string) *ListPvcsObjectMeta {
  s.ClusterName = &v
  return s
}

type ListPvcsOwnerReference struct {
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

func (s ListPvcsOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsOwnerReference) GoString() string {
  return s.String()
}

func (s *ListPvcsOwnerReference) SetApiVersion(v string) *ListPvcsOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *ListPvcsOwnerReference) SetKind(v string) *ListPvcsOwnerReference {
  s.Kind = &v
  return s
}

func (s *ListPvcsOwnerReference) SetName(v string) *ListPvcsOwnerReference {
  s.Name = &v
  return s
}

func (s *ListPvcsOwnerReference) SetUid(v string) *ListPvcsOwnerReference {
  s.Uid = &v
  return s
}

func (s *ListPvcsOwnerReference) SetController(v bool) *ListPvcsOwnerReference {
  s.Controller = &v
  return s
}

func (s *ListPvcsOwnerReference) SetBlockOwnerDeletion(v bool) *ListPvcsOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type ListPvcsPersistentVolumeClaim struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *ListPvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *ListPvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"represents the current information/status of a persistent volume claim. Read-only", "zh_CN":"表示一个持久卷申领的当前信息/状态。只读"}
  Status *ListPvcsPersistentVolumeClaimStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPvcsPersistentVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsPersistentVolumeClaim) GoString() string {
  return s.String()
}

func (s *ListPvcsPersistentVolumeClaim) SetApiVersion(v string) *ListPvcsPersistentVolumeClaim {
  s.ApiVersion = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaim) SetKind(v string) *ListPvcsPersistentVolumeClaim {
  s.Kind = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaim) SetMetadata(v *ListPvcsObjectMeta) *ListPvcsPersistentVolumeClaim {
  s.Metadata = v
  return s
}

func (s *ListPvcsPersistentVolumeClaim) SetSpec(v *ListPvcsPersistentVolumeClaimSpec) *ListPvcsPersistentVolumeClaim {
  s.Spec = v
  return s
}

func (s *ListPvcsPersistentVolumeClaim) SetStatus(v *ListPvcsPersistentVolumeClaimStatus) *ListPvcsPersistentVolumeClaim {
  s.Status = v
  return s
}

type ListPvcsPersistentVolumeClaimSpec struct {
  // {"en":"contains the desired access modes the volume should have", "zh_CN":"包含卷应具备的预期访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"a label query over volumes to consider for binding", "zh_CN":"在绑定时对卷进行选择所执行的标签查询"}
  Selector *ListPvcsMetaV1LabelSelector `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim", "zh_CN":"表示卷应拥有的最小资源。 如果启用了 RecoverVolumeExpansionFailure 功能特性，则允许用户指定这些资源要求， 此值必须低于之前的值，但必须高于申领的状态字段中记录的容量"}
  Resources *ListPvcsResourceRequirements `json:"resources,omitempty" xml:"resources,omitempty"`
  // {"en":"the binding reference to the PersistentVolume backing this claim", "zh_CN":"对此申领所对应的 PersistentVolume 的绑定引用"}
  VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
  // {"en":"the name of the StorageClass required by the claim", "zh_CN":"此申领所要求的 StorageClass 名称"}
  StorageClassName *string `json:"storageClassName,omitempty" xml:"storageClassName,omitempty"`
  // {"en":"defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec", "zh_CN":"定义申领需要哪种类别的卷。当申领规约中未包含此字段时，意味着取值为 Filesystem"}
  VolumeMode *string `json:"volumeMode,omitempty" xml:"volumeMode,omitempty"`
  // {"en":"dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (ListPvcsPersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource", "zh_CN":"dataSource 字段可用于二选一：- 现有的 VolumeSnapshot 对象（snapshot.storage.k8s.io/VolumeSnapshot）- 现有的 PVC (ListPvcsPersistentVolumeClaim)。如果制备器或外部控制器可以支持指定的数据源，则它将根据指定数据源的内容创建新的卷。 当 AnyVolumeDataSource 特性门控被启用时，dataSource 内容将被复制到 dataSourceRef， 当 dataSourceRef.namespace 未被指定时，dataSourceRef 内容将被复制到 dataSource。 如果名字空间被指定，则 dataSourceRef 不会被复制到 dataSource"}
  DataSource *ListPvcsTypedLocalObjectReference `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
}

func (s ListPvcsPersistentVolumeClaimSpec) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsPersistentVolumeClaimSpec) GoString() string {
  return s.String()
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetAccessModes(v []*string) *ListPvcsPersistentVolumeClaimSpec {
  s.AccessModes = v
  return s
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetSelector(v *ListPvcsMetaV1LabelSelector) *ListPvcsPersistentVolumeClaimSpec {
  s.Selector = v
  return s
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetResources(v *ListPvcsResourceRequirements) *ListPvcsPersistentVolumeClaimSpec {
  s.Resources = v
  return s
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetVolumeName(v string) *ListPvcsPersistentVolumeClaimSpec {
  s.VolumeName = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetStorageClassName(v string) *ListPvcsPersistentVolumeClaimSpec {
  s.StorageClassName = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetVolumeMode(v string) *ListPvcsPersistentVolumeClaimSpec {
  s.VolumeMode = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimSpec) SetDataSource(v *ListPvcsTypedLocalObjectReference) *ListPvcsPersistentVolumeClaimSpec {
  s.DataSource = v
  return s
}

type ListPvcsMetaV1LabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*ListPvcsLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s ListPvcsMetaV1LabelSelector) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsMetaV1LabelSelector) GoString() string {
  return s.String()
}

func (s *ListPvcsMetaV1LabelSelector) SetMatchLabels(v map[string]*string) *ListPvcsMetaV1LabelSelector {
  s.MatchLabels = v
  return s
}

func (s *ListPvcsMetaV1LabelSelector) SetMatchExpressions(v []*ListPvcsLabelSelectorRequirement) *ListPvcsMetaV1LabelSelector {
  s.MatchExpressions = v
  return s
}

type ListPvcsLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListPvcsLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *ListPvcsLabelSelectorRequirement) SetKey(v string) *ListPvcsLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *ListPvcsLabelSelectorRequirement) SetOperator(v string) *ListPvcsLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *ListPvcsLabelSelectorRequirement) SetValues(v []*string) *ListPvcsLabelSelectorRequirement {
  s.Values = v
  return s
}

type ListPvcsResourceRequirements struct {
  // {"en":"describes the maximum amount of compute resources allowed", "zh_CN":"描述所允许的最大计算资源用量"}
  Limits map[string]*string `json:"limits,omitempty" xml:"limits,omitempty"`
  // {"en":"describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits", "zh_CN":"requests 描述所需的最小计算资源量。如果容器省略了 requests，但明确设定了 limits， 则 requests 默认值为 limits 值，否则为实现定义的值。请求不能超过限制"}
  Requests map[string]*string `json:"requests,omitempty" xml:"requests,omitempty"`
}

func (s ListPvcsResourceRequirements) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsResourceRequirements) GoString() string {
  return s.String()
}

func (s *ListPvcsResourceRequirements) SetLimits(v map[string]*string) *ListPvcsResourceRequirements {
  s.Limits = v
  return s
}

func (s *ListPvcsResourceRequirements) SetRequests(v map[string]*string) *ListPvcsResourceRequirements {
  s.Requests = v
  return s
}

type ListPvcsTypedLocalObjectReference struct {
  // {"en":"the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required", "zh_CN":"被引用资源的组。如果不指定 APIGroup，则指定的 Kind 必须在核心 API 组中。对于任何其它第三方类型，都需要 APIGroup"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
  // {"en":" the type of resource being referenced", "zh_CN":"被引用的资源的类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"the name of resource being referenced", "zh_CN":"被引用的资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPvcsTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *ListPvcsTypedLocalObjectReference) SetApiGroup(v string) *ListPvcsTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

func (s *ListPvcsTypedLocalObjectReference) SetKind(v string) *ListPvcsTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *ListPvcsTypedLocalObjectReference) SetName(v string) *ListPvcsTypedLocalObjectReference {
  s.Name = &v
  return s
}

type ListPvcsPersistentVolumeClaimStatus struct {
  // {"en":"represents the current phase of ListPvcsPersistentVolumeClaim", "zh_CN":"表示 ListPvcsPersistentVolumeClaim 的当前阶段"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
  // {"en":"contains the actual access modes the volume backing the PVC has", "zh_CN":"包含支持 PVC 的卷所具有的实际访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"represents the actual resources of the underlying volume", "zh_CN":"表示底层卷的实际资源"}
  Capacity map[string]*int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
  // {"en":"the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'", "zh_CN":"持久卷声明的当前的状况。 如果正在调整底层持久卷的大小，则状况将被设为 “ResizeStarted”"}
  Conditions []*ListPvcsPersistentVolumeClaimCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s ListPvcsPersistentVolumeClaimStatus) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsPersistentVolumeClaimStatus) GoString() string {
  return s.String()
}

func (s *ListPvcsPersistentVolumeClaimStatus) SetPhase(v string) *ListPvcsPersistentVolumeClaimStatus {
  s.Phase = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimStatus) SetAccessModes(v []*string) *ListPvcsPersistentVolumeClaimStatus {
  s.AccessModes = v
  return s
}

func (s *ListPvcsPersistentVolumeClaimStatus) SetCapacity(v map[string]*int64) *ListPvcsPersistentVolumeClaimStatus {
  s.Capacity = v
  return s
}

func (s *ListPvcsPersistentVolumeClaimStatus) SetConditions(v []*ListPvcsPersistentVolumeClaimCondition) *ListPvcsPersistentVolumeClaimStatus {
  s.Conditions = v
  return s
}

type ListPvcsPersistentVolumeClaimCondition struct {
  // {"en":"type, required", "zh_CN":"类型，必需"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status, required", "zh_CN":"状态，必需"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized", "zh_CN":"reason 是唯一的，它应该是一个机器可理解的简短字符串，指明上次状况转换的原因。 如果它报告 “ResizeStarted”，则意味着正在调整底层持久卷的大小"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":" the human-readable message indicating details about last transition", "zh_CN":"人类可读的消息，指示有关上一次转换的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListPvcsPersistentVolumeClaimCondition) String() string {
  return tea.Prettify(s)
}

func (s ListPvcsPersistentVolumeClaimCondition) GoString() string {
  return s.String()
}

func (s *ListPvcsPersistentVolumeClaimCondition) SetType(v string) *ListPvcsPersistentVolumeClaimCondition {
  s.Type = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimCondition) SetStatus(v string) *ListPvcsPersistentVolumeClaimCondition {
  s.Status = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimCondition) SetReason(v string) *ListPvcsPersistentVolumeClaimCondition {
  s.Reason = &v
  return s
}

func (s *ListPvcsPersistentVolumeClaimCondition) SetMessage(v string) *ListPvcsPersistentVolumeClaimCondition {
  s.Message = &v
  return s
}




type QueryChannelRecordFilesRequest struct {
  // {"en":"Start time:
  // 1. The format is yyyyy-MM-ddTHH: mm: SS + 08:00, for example, 2024-01-23T10:00 + 08:00 (10:00:00 Beijing time on January 23, 2024);
  // 2. Can not exceed the current time;", "zh_CN":"开始时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00，例如，2024-01-23T10:00:00+08:00（为北京时间2024年01月23日10点0分0秒）；
  // 2.不能大于当前时间"}
  DateFrom *string `json:"dateFrom,omitempty" xml:"dateFrom,omitempty"`
  // {"en":"End time:
  // 1. The format is yyyy-MM-ddTHH:mm:ss+08:00;
  // 2. The end time is greater than the start time.
  // 3. If the end time is greater than the current time, the current time is taken.
  // 4. DateFrom and dateTo are not uploaded, defaulting to query the past 24 hours; if only one is not uploaded, throw an exception;
  // 5. Maximum query interval allowed: 30 days, that is, the difference between dateFrom and dateTo can not exceed 30 days. ", "zh_CN":"结束时间：
  // 1.格式为yyyy-MM-ddTHH:mm:ss+08:00；
  // 2.结束时间需大于开始时间；
  // 3.结束时间如果大于当前时间，取当前时间；
  // 4.dateFrom，dateTo二者都未传，默认查询过去的24小时；如仅有一个未传，抛异常；
  // 5.允许查询最大间隔：30天，即dateFrom和dateTo相差不能超过30天。"}
  DateTo *string `json:"dateTo,omitempty" xml:"dateTo,omitempty"`
  // {"en":"Device GB28181 Id", "zh_CN":"设备国标ID"}
  DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty" require:"true"`
  // {"en":"ChannelGB28181 Id", "zh_CN":"通道国标ID"}
  ChannelId *string `json:"channelId,omitempty" xml:"channelId,omitempty" require:"true"`
  // {"en":"Paging Size.
  // 1. Default 10, maximum 50", "zh_CN":"分页大小。
  // 1、默认10，最大50"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
  // {"en":"Page index
  //  1. The default value is 1, indicating the initial page.", "zh_CN":"第几页
  // 1、默认为1，表示第一页"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
}

func (s QueryChannelRecordFilesRequest) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesRequest) GoString() string {
  return s.String()
}

func (s *QueryChannelRecordFilesRequest) SetDateFrom(v string) *QueryChannelRecordFilesRequest {
  s.DateFrom = &v
  return s
}

func (s *QueryChannelRecordFilesRequest) SetDateTo(v string) *QueryChannelRecordFilesRequest {
  s.DateTo = &v
  return s
}

func (s *QueryChannelRecordFilesRequest) SetDeviceId(v string) *QueryChannelRecordFilesRequest {
  s.DeviceId = &v
  return s
}

func (s *QueryChannelRecordFilesRequest) SetChannelId(v string) *QueryChannelRecordFilesRequest {
  s.ChannelId = &v
  return s
}

func (s *QueryChannelRecordFilesRequest) SetPageSize(v int32) *QueryChannelRecordFilesRequest {
  s.PageSize = &v
  return s
}

func (s *QueryChannelRecordFilesRequest) SetPageIndex(v int32) *QueryChannelRecordFilesRequest {
  s.PageIndex = &v
  return s
}

type QueryChannelRecordFilesResponse struct {
  // {"en":"Result status code, 0 indicates success", "zh_CN":"结果状态码，0为成功"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Return message", "zh_CN":"返回消息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
  // {"en":"Return data", "zh_CN":"返回数据"}
  QueryChannelRecordFilesData *QueryChannelRecordFilesData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryChannelRecordFilesResponse) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesResponse) GoString() string {
  return s.String()
}

func (s *QueryChannelRecordFilesResponse) SetCode(v int32) *QueryChannelRecordFilesResponse {
  s.Code = &v
  return s
}

func (s *QueryChannelRecordFilesResponse) SetMessage(v string) *QueryChannelRecordFilesResponse {
  s.Message = &v
  return s
}

func (s *QueryChannelRecordFilesResponse) SetData(v *QueryChannelRecordFilesData) *QueryChannelRecordFilesResponse {
  s.QueryChannelRecordFilesData = v
  return s
}

type QueryChannelRecordFilesData struct {
  // {"en":"QueryChannelRecordFilesFile List", "zh_CN":"文件列表"}
  Rows []*QueryChannelRecordFilesFile `json:"rows,omitempty" xml:"rows,omitempty" require:"true" type:"Repeated"`
  // {"en":"Page Index", "zh_CN":"第几页"}
  PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty" require:"true"`
  // {"en":"Paging Size", "zh_CN":"分页大小"}
  PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty" require:"true"`
  // {"en":"Total", "zh_CN":"总数"}
  Total *int32 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryChannelRecordFilesData) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesData) GoString() string {
  return s.String()
}

func (s *QueryChannelRecordFilesData) SetRows(v []*QueryChannelRecordFilesFile) *QueryChannelRecordFilesData {
  s.Rows = v
  return s
}

func (s *QueryChannelRecordFilesData) SetPageIndex(v int32) *QueryChannelRecordFilesData {
  s.PageIndex = &v
  return s
}

func (s *QueryChannelRecordFilesData) SetPageSize(v int32) *QueryChannelRecordFilesData {
  s.PageSize = &v
  return s
}

func (s *QueryChannelRecordFilesData) SetTotal(v int32) *QueryChannelRecordFilesData {
  s.Total = &v
  return s
}

type QueryChannelRecordFilesFile struct {
  // {"en":"QueryChannelRecordFilesFile name", "zh_CN":"文件名"}
  FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty" require:"true"`
  // {"en":"QueryChannelRecordFilesFile size
  // 1. Unit (M)", "zh_CN":"文件大小
  // 1、单位（M）"}
  FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty" require:"true"`
  // {"en":"QueryChannelRecordFilesFile duration
  // 1. Unit (second)", "zh_CN":"文件时长
  // 1、单位(秒)"}
  Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
  // {"en":"QueryChannelRecordFilesFile start time
  // 1. Timestamp format", "zh_CN":"文件开始时间
  // 1、时间戳格式"}
  StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty" require:"true"`
  // {"en":"QueryChannelRecordFilesFile expiration time
  // 1. Timestamp format", "zh_CN":"文件过期时间
  // 1、时间戳格式"}
  ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty" require:"true"`
  // {"en":"QueryChannelRecordFilesFile url", "zh_CN":"文件url"}
  FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty" require:"true"`
}

func (s QueryChannelRecordFilesFile) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesFile) GoString() string {
  return s.String()
}

func (s *QueryChannelRecordFilesFile) SetFileName(v string) *QueryChannelRecordFilesFile {
  s.FileName = &v
  return s
}

func (s *QueryChannelRecordFilesFile) SetFileSize(v string) *QueryChannelRecordFilesFile {
  s.FileSize = &v
  return s
}

func (s *QueryChannelRecordFilesFile) SetDuration(v int32) *QueryChannelRecordFilesFile {
  s.Duration = &v
  return s
}

func (s *QueryChannelRecordFilesFile) SetStartTime(v string) *QueryChannelRecordFilesFile {
  s.StartTime = &v
  return s
}

func (s *QueryChannelRecordFilesFile) SetExpireTime(v string) *QueryChannelRecordFilesFile {
  s.ExpireTime = &v
  return s
}

func (s *QueryChannelRecordFilesFile) SetFileUrl(v string) *QueryChannelRecordFilesFile {
  s.FileUrl = &v
  return s
}

type QueryChannelRecordFilesPaths struct {
}

func (s QueryChannelRecordFilesPaths) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesPaths) GoString() string {
  return s.String()
}

type QueryChannelRecordFilesParameters struct {
}

func (s QueryChannelRecordFilesParameters) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesParameters) GoString() string {
  return s.String()
}

type QueryChannelRecordFilesRequestHeader struct {
}

func (s QueryChannelRecordFilesRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesRequestHeader) GoString() string {
  return s.String()
}

type QueryChannelRecordFilesResponseHeader struct {
}

func (s QueryChannelRecordFilesResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s QueryChannelRecordFilesResponseHeader) GoString() string {
  return s.String()
}




type DeletePvcsRequest struct {
}

func (s DeletePvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsRequest) GoString() string {
  return s.String()
}

type DeletePvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"status", "zh_CN":"status"}
  Data *DeletePvcsStatus `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeletePvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsResponse) GoString() string {
  return s.String()
}

func (s *DeletePvcsResponse) SetCode(v int64) *DeletePvcsResponse {
  s.Code = &v
  return s
}

func (s *DeletePvcsResponse) SetMsg(v string) *DeletePvcsResponse {
  s.Msg = &v
  return s
}

func (s *DeletePvcsResponse) SetRequestId(v string) *DeletePvcsResponse {
  s.RequestId = &v
  return s
}

func (s *DeletePvcsResponse) SetData(v *DeletePvcsStatus) *DeletePvcsResponse {
  s.Data = v
  return s
}

type DeletePvcsPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"deployment name", "zh_CN":"pvc 名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeletePvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsPaths) GoString() string {
  return s.String()
}

func (s *DeletePvcsPaths) SetNamespace(v string) *DeletePvcsPaths {
  s.Namespace = &v
  return s
}

func (s *DeletePvcsPaths) SetName(v string) *DeletePvcsPaths {
  s.Name = &v
  return s
}

type DeletePvcsParameters struct {
}

func (s DeletePvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsParameters) GoString() string {
  return s.String()
}

type DeletePvcsRequestHeader struct {
}

func (s DeletePvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsRequestHeader) GoString() string {
  return s.String()
}

type DeletePvcsResponseHeader struct {
}

func (s DeletePvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsResponseHeader) GoString() string {
  return s.String()
}

type DeletePvcsStatus struct {
  // {"en":"APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values", "zh_CN":"APIVersion 定义对象表示的版本化模式。 服务器应将已识别的模式转换为最新的内部值，并可能拒绝无法识别的值"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase", "zh_CN":"Kind 是一个字符串值，表示此对象表示的 REST 资源。 服务器可以从客户端提交请求的端点推断出这一点。 无法更新。驼峰式规则"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"DeletePvcsStatus of the operation. One of: 'Success' or 'Failure'", "zh_CN":"操作状态。“Success”或“Failure” 之一"}
  DeletePvcsStatus *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
  // {"en":"Suggested HTTP return code for this status, 0 if not set", "zh_CN":"此状态的建议 HTTP 返回代码，如果未设置，则为 0"}
  Code *int32 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"Extended data associated with the reason. Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type", "zh_CN":"与原因（Reason）相关的扩展数据。每个原因都可以定义自己的扩展细节。 此字段是可选的，并且不保证返回的数据符合任何模式，除非由原因类型定义"}
  Details *DeletePvcsStatusDetails `json:"details,omitempty" xml:"details,omitempty" require:"true"`
}

func (s DeletePvcsStatus) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsStatus) GoString() string {
  return s.String()
}

func (s *DeletePvcsStatus) SetApiVersion(v string) *DeletePvcsStatus {
  s.ApiVersion = &v
  return s
}

func (s *DeletePvcsStatus) SetKind(v string) *DeletePvcsStatus {
  s.Kind = &v
  return s
}

func (s *DeletePvcsStatus) SetStatus(v string) *DeletePvcsStatus {
  s.DeletePvcsStatus = &v
  return s
}

func (s *DeletePvcsStatus) SetCode(v int32) *DeletePvcsStatus {
  s.Code = &v
  return s
}

func (s *DeletePvcsStatus) SetDetails(v *DeletePvcsStatusDetails) *DeletePvcsStatus {
  s.Details = v
  return s
}

type DeletePvcsStatusDetails struct {
  // {"en":"The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described)", "zh_CN":"与状态 StatusReason 关联的资源的名称属性（当有一个可以描述的名称时）"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  // {"en":"The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind", "zh_CN":"与状态 StatusReason 关联的资源的种类属性"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"The group attribute of the resource associated with the status StatusReason", "zh_CN":"与状态 StatusReason 关联的资源的组属性"}
  Group *string `json:"group,omitempty" xml:"group,omitempty" require:"true"`
  // {"en":"UID of the resource. (when there is a single resource which can be described)", "zh_CN":"资源的 UID（当有单个可以描述的资源时）"}
  Uid *string `json:"uid,omitempty" xml:"uid,omitempty" require:"true"`
}

func (s DeletePvcsStatusDetails) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcsStatusDetails) GoString() string {
  return s.String()
}

func (s *DeletePvcsStatusDetails) SetName(v string) *DeletePvcsStatusDetails {
  s.Name = &v
  return s
}

func (s *DeletePvcsStatusDetails) SetKind(v string) *DeletePvcsStatusDetails {
  s.Kind = &v
  return s
}

func (s *DeletePvcsStatusDetails) SetGroup(v string) *DeletePvcsStatusDetails {
  s.Group = &v
  return s
}

func (s *DeletePvcsStatusDetails) SetUid(v string) *DeletePvcsStatusDetails {
  s.Uid = &v
  return s
}




type CreatePvcsRequest struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreatePvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *CreatePvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreatePvcsRequest) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsRequest) GoString() string {
  return s.String()
}

func (s *CreatePvcsRequest) SetApiVersion(v string) *CreatePvcsRequest {
  s.ApiVersion = &v
  return s
}

func (s *CreatePvcsRequest) SetKind(v string) *CreatePvcsRequest {
  s.Kind = &v
  return s
}

func (s *CreatePvcsRequest) SetMetadata(v *CreatePvcsObjectMeta) *CreatePvcsRequest {
  s.Metadata = v
  return s
}

func (s *CreatePvcsRequest) SetSpec(v *CreatePvcsPersistentVolumeClaimSpec) *CreatePvcsRequest {
  s.Spec = v
  return s
}

type CreatePvcsResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
  // {"en":"pvc object", "zh_CN":"pvc对象"}
  Data *CreatePvcsPersistentVolumeClaim `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CreatePvcsResponse) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsResponse) GoString() string {
  return s.String()
}

func (s *CreatePvcsResponse) SetCode(v int64) *CreatePvcsResponse {
  s.Code = &v
  return s
}

func (s *CreatePvcsResponse) SetMsg(v string) *CreatePvcsResponse {
  s.Msg = &v
  return s
}

func (s *CreatePvcsResponse) SetRequestId(v string) *CreatePvcsResponse {
  s.RequestId = &v
  return s
}

func (s *CreatePvcsResponse) SetData(v *CreatePvcsPersistentVolumeClaim) *CreatePvcsResponse {
  s.Data = v
  return s
}

type CreatePvcsPaths struct {
  // {"en":"namespace", "zh_CN":"命名空间"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
}

func (s CreatePvcsPaths) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsPaths) GoString() string {
  return s.String()
}

func (s *CreatePvcsPaths) SetNamespace(v string) *CreatePvcsPaths {
  s.Namespace = &v
  return s
}

type CreatePvcsParameters struct {
}

func (s CreatePvcsParameters) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsParameters) GoString() string {
  return s.String()
}

type CreatePvcsRequestHeader struct {
}

func (s CreatePvcsRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsRequestHeader) GoString() string {
  return s.String()
}

type CreatePvcsResponseHeader struct {
}

func (s CreatePvcsResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsResponseHeader) GoString() string {
  return s.String()
}

type CreatePvcsObjectMeta struct {
  // {"en":"must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated", "zh_CN":"name 在命名空间内必须是唯一的。创建资源时需要，尽管某些资源可能允许客户端请求自动地生成适当的名称。 名称主要用于创建幂等性和配置定义。无法更新"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // {"en":"an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server", "zh_CN":"一个可选前缀，由服务器使用，仅在未提供 name 字段时生成唯一名称。 如果使用此字段，则返回给客户端的名称将与传递的名称不同。该值还将与唯一的后缀组合。 提供的值与 name 字段具有相同的验证规则，并且可能会根据所需的后缀长度被截断，以使该值在服务器上唯一"}
  GenerateName *string `json:"generateName,omitempty" xml:"generateName,omitempty"`
  // {"en":"Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.Must be a DNS_LABEL. Cannot be updated", "zh_CN":"namespace 定义了一个值空间，其中每个名称必须唯一。空命名空间相当于 “default” 命名空间，但 “default” 是规范表示。 并非所有对象都需要限定在命名空间中——这些对象的此字段的值将为空。必须是 DNS_LABEL。无法更新。"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
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
  OwnerReferences []*CreatePvcsOwnerReference `json:"ownerReferences,omitempty" xml:"ownerReferences,omitempty" type:"Repeated"`
  // {"en":"Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed. Finalizers may be processed and removed in any order. Order is NOT enforced because it introduces significant risk of stuck finalizers. finalizers is a shared field, any actor with permission can reorder it. If the finalizer list is processed in order, then this can lead to a situation in which the component responsible for the first finalizer in the list is waiting for a signal (field value, external system, or other) produced by a component responsible for a finalizer later in the list, resulting in a deadlock. Without enforced ordering finalizers are free to order amongst themselves and are not vulnerable to ordering changes in the list.", "zh_CN":"在从注册表中删除对象之前该字段必须为空。 每个条目都是负责的组件的标识符，各组件将从列表中删除自己对应的条目。 如果对象的 deletionTimestamp 非空，则只能删除此列表中的条目。 终结器可以按任何顺序处理和删除。没有按照顺序执行， 因为它引入了终结器卡住的重大风险。finalizers 是一个共享字段， 任何有权限的参与者都可以对其进行重新排序。如果按顺序处理终结器列表， 那么这可能导致列表中第一个负责终结器的组件正在等待列表中靠后负责终结器的组件产生的信号（字段值、外部系统或其他）， 从而导致死锁。在没有强制排序的情况下，终结者可以在它们之间自由排序， 并且不容易受到列表中排序更改的影响。"}
  Finalizers []*string `json:"finalizers,omitempty" xml:"finalizers,omitempty" type:"Repeated"`
  // {"en":"name of cluster", "zh_CN":"集群名称"}
  ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
}

func (s CreatePvcsObjectMeta) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsObjectMeta) GoString() string {
  return s.String()
}

func (s *CreatePvcsObjectMeta) SetName(v string) *CreatePvcsObjectMeta {
  s.Name = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetGenerateName(v string) *CreatePvcsObjectMeta {
  s.GenerateName = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetNamespace(v string) *CreatePvcsObjectMeta {
  s.Namespace = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetSelfLink(v string) *CreatePvcsObjectMeta {
  s.SelfLink = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetUid(v string) *CreatePvcsObjectMeta {
  s.Uid = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetResourceVersion(v string) *CreatePvcsObjectMeta {
  s.ResourceVersion = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetGeneration(v int64) *CreatePvcsObjectMeta {
  s.Generation = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetCreationTimestamp(v string) *CreatePvcsObjectMeta {
  s.CreationTimestamp = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetDeletionTimestamp(v string) *CreatePvcsObjectMeta {
  s.DeletionTimestamp = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetDeletionGracePeriodSeconds(v int64) *CreatePvcsObjectMeta {
  s.DeletionGracePeriodSeconds = &v
  return s
}

func (s *CreatePvcsObjectMeta) SetLabels(v map[string]*string) *CreatePvcsObjectMeta {
  s.Labels = v
  return s
}

func (s *CreatePvcsObjectMeta) SetAnnotations(v map[string]*string) *CreatePvcsObjectMeta {
  s.Annotations = v
  return s
}

func (s *CreatePvcsObjectMeta) SetOwnerReferences(v []*CreatePvcsOwnerReference) *CreatePvcsObjectMeta {
  s.OwnerReferences = v
  return s
}

func (s *CreatePvcsObjectMeta) SetFinalizers(v []*string) *CreatePvcsObjectMeta {
  s.Finalizers = v
  return s
}

func (s *CreatePvcsObjectMeta) SetClusterName(v string) *CreatePvcsObjectMeta {
  s.ClusterName = &v
  return s
}

type CreatePvcsOwnerReference struct {
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

func (s CreatePvcsOwnerReference) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsOwnerReference) GoString() string {
  return s.String()
}

func (s *CreatePvcsOwnerReference) SetApiVersion(v string) *CreatePvcsOwnerReference {
  s.ApiVersion = &v
  return s
}

func (s *CreatePvcsOwnerReference) SetKind(v string) *CreatePvcsOwnerReference {
  s.Kind = &v
  return s
}

func (s *CreatePvcsOwnerReference) SetName(v string) *CreatePvcsOwnerReference {
  s.Name = &v
  return s
}

func (s *CreatePvcsOwnerReference) SetUid(v string) *CreatePvcsOwnerReference {
  s.Uid = &v
  return s
}

func (s *CreatePvcsOwnerReference) SetController(v bool) *CreatePvcsOwnerReference {
  s.Controller = &v
  return s
}

func (s *CreatePvcsOwnerReference) SetBlockOwnerDeletion(v bool) *CreatePvcsOwnerReference {
  s.BlockOwnerDeletion = &v
  return s
}

type CreatePvcsPersistentVolumeClaim struct {
  // {"en":"apiVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.", "zh_CN":"APIVersion定义了表示对象的版本化模式。服务器应该将认可的模式转换为最新的内部值，并可以拒绝不被认可的值。"}
  ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty" require:"true"`
  // {"en":"kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.", "zh_CN":"kind是一个字符串值，表示此对象所代表的REST资源。服务器可以根据客户端提交请求的终点推断出这个值。不能更新。"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty" require:"true"`
  // {"en":"standard object metadata.", "zh_CN":"标准的对象元数据"}
  Metadata *CreatePvcsObjectMeta `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // {"en":"defines the desired characteristics of a volume requested by a pod author", "zh_CN":"定义 Pod 作者所请求的卷的预期特征"}
  Spec *CreatePvcsPersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec,omitempty"`
  // {"en":"represents the current information/status of a persistent volume claim. Read-only", "zh_CN":"表示一个持久卷申领的当前信息/状态。只读"}
  Status *CreatePvcsPersistentVolumeClaimStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePvcsPersistentVolumeClaim) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsPersistentVolumeClaim) GoString() string {
  return s.String()
}

func (s *CreatePvcsPersistentVolumeClaim) SetApiVersion(v string) *CreatePvcsPersistentVolumeClaim {
  s.ApiVersion = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaim) SetKind(v string) *CreatePvcsPersistentVolumeClaim {
  s.Kind = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaim) SetMetadata(v *CreatePvcsObjectMeta) *CreatePvcsPersistentVolumeClaim {
  s.Metadata = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaim) SetSpec(v *CreatePvcsPersistentVolumeClaimSpec) *CreatePvcsPersistentVolumeClaim {
  s.Spec = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaim) SetStatus(v *CreatePvcsPersistentVolumeClaimStatus) *CreatePvcsPersistentVolumeClaim {
  s.Status = v
  return s
}

type CreatePvcsPersistentVolumeClaimSpec struct {
  // {"en":"contains the desired access modes the volume should have", "zh_CN":"包含卷应具备的预期访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"a label query over volumes to consider for binding", "zh_CN":"在绑定时对卷进行选择所执行的标签查询"}
  Selector *CreatePvcsMetaV1LabelSelector `json:"selector,omitempty" xml:"selector,omitempty"`
  // {"en":"represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim", "zh_CN":"表示卷应拥有的最小资源。 如果启用了 RecoverVolumeExpansionFailure 功能特性，则允许用户指定这些资源要求， 此值必须低于之前的值，但必须高于申领的状态字段中记录的容量"}
  Resources *CreatePvcsResourceRequirements `json:"resources,omitempty" xml:"resources,omitempty"`
  // {"en":"the binding reference to the PersistentVolume backing this claim", "zh_CN":"对此申领所对应的 PersistentVolume 的绑定引用"}
  VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
  // {"en":"the name of the StorageClass required by the claim", "zh_CN":"此申领所要求的 StorageClass 名称"}
  StorageClassName *string `json:"storageClassName,omitempty" xml:"storageClassName,omitempty"`
  // {"en":"defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec", "zh_CN":"定义申领需要哪种类别的卷。当申领规约中未包含此字段时，意味着取值为 Filesystem"}
  VolumeMode *string `json:"volumeMode,omitempty" xml:"volumeMode,omitempty"`
  // {"en":"dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (CreatePvcsPersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource", "zh_CN":"dataSource 字段可用于二选一：- 现有的 VolumeSnapshot 对象（snapshot.storage.k8s.io/VolumeSnapshot）- 现有的 PVC (CreatePvcsPersistentVolumeClaim)。如果制备器或外部控制器可以支持指定的数据源，则它将根据指定数据源的内容创建新的卷。 当 AnyVolumeDataSource 特性门控被启用时，dataSource 内容将被复制到 dataSourceRef， 当 dataSourceRef.namespace 未被指定时，dataSourceRef 内容将被复制到 dataSource。 如果名字空间被指定，则 dataSourceRef 不会被复制到 dataSource"}
  DataSource *CreatePvcsTypedLocalObjectReference `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
}

func (s CreatePvcsPersistentVolumeClaimSpec) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsPersistentVolumeClaimSpec) GoString() string {
  return s.String()
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetAccessModes(v []*string) *CreatePvcsPersistentVolumeClaimSpec {
  s.AccessModes = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetSelector(v *CreatePvcsMetaV1LabelSelector) *CreatePvcsPersistentVolumeClaimSpec {
  s.Selector = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetResources(v *CreatePvcsResourceRequirements) *CreatePvcsPersistentVolumeClaimSpec {
  s.Resources = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetVolumeName(v string) *CreatePvcsPersistentVolumeClaimSpec {
  s.VolumeName = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetStorageClassName(v string) *CreatePvcsPersistentVolumeClaimSpec {
  s.StorageClassName = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetVolumeMode(v string) *CreatePvcsPersistentVolumeClaimSpec {
  s.VolumeMode = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimSpec) SetDataSource(v *CreatePvcsTypedLocalObjectReference) *CreatePvcsPersistentVolumeClaimSpec {
  s.DataSource = v
  return s
}

type CreatePvcsMetaV1LabelSelector struct {
  // {"en":"A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.", "zh_CN":"matchLabels 映射中的单个 {key,value} 键值对相当于 matchExpressions 的一个元素，其键字段为 key，运算符为 In，values 数组仅包含 value。"}
  MatchLabels map[string]*string `json:"matchLabels,omitempty" xml:"matchLabels,omitempty"`
  // {"en":"A list of label selector requirements. The requirements are ANDed.", "zh_CN":"标签选择器要求的列表，这些要求的结果按逻辑与的关系来计算。"}
  MatchExpressions []*CreatePvcsLabelSelectorRequirement `json:"matchExpressions,omitempty" xml:"matchExpressions,omitempty" type:"Repeated"`
}

func (s CreatePvcsMetaV1LabelSelector) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsMetaV1LabelSelector) GoString() string {
  return s.String()
}

func (s *CreatePvcsMetaV1LabelSelector) SetMatchLabels(v map[string]*string) *CreatePvcsMetaV1LabelSelector {
  s.MatchLabels = v
  return s
}

func (s *CreatePvcsMetaV1LabelSelector) SetMatchExpressions(v []*CreatePvcsLabelSelectorRequirement) *CreatePvcsMetaV1LabelSelector {
  s.MatchExpressions = v
  return s
}

type CreatePvcsLabelSelectorRequirement struct {
  // {"en":"key is the label key that the selector applies to.", "zh_CN":"选择器应用的标签键"}
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // {"en":"operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.", "zh_CN":"operator 表示键与一组值的关系。有效的运算符包括 In、NotIn、Exists 和 DoesNotExist。"}
  Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
  // {"en":"values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.", "zh_CN":"values 是一个字符串值数组。如果运算符为 In 或 NotIn，则 values 数组必须为非空。如果运算符是 Exists 或 DoesNotExist，则 values 数组必须为空。该数组在策略性合并补丁（Strategic Merge Patch）期间被替换。"}
  Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreatePvcsLabelSelectorRequirement) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsLabelSelectorRequirement) GoString() string {
  return s.String()
}

func (s *CreatePvcsLabelSelectorRequirement) SetKey(v string) *CreatePvcsLabelSelectorRequirement {
  s.Key = &v
  return s
}

func (s *CreatePvcsLabelSelectorRequirement) SetOperator(v string) *CreatePvcsLabelSelectorRequirement {
  s.Operator = &v
  return s
}

func (s *CreatePvcsLabelSelectorRequirement) SetValues(v []*string) *CreatePvcsLabelSelectorRequirement {
  s.Values = v
  return s
}

type CreatePvcsResourceRequirements struct {
  // {"en":"describes the maximum amount of compute resources allowed", "zh_CN":"描述所允许的最大计算资源用量"}
  Limits map[string]*string `json:"limits,omitempty" xml:"limits,omitempty"`
  // {"en":"describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits", "zh_CN":"requests 描述所需的最小计算资源量。如果容器省略了 requests，但明确设定了 limits， 则 requests 默认值为 limits 值，否则为实现定义的值。请求不能超过限制"}
  Requests map[string]*string `json:"requests,omitempty" xml:"requests,omitempty"`
}

func (s CreatePvcsResourceRequirements) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsResourceRequirements) GoString() string {
  return s.String()
}

func (s *CreatePvcsResourceRequirements) SetLimits(v map[string]*string) *CreatePvcsResourceRequirements {
  s.Limits = v
  return s
}

func (s *CreatePvcsResourceRequirements) SetRequests(v map[string]*string) *CreatePvcsResourceRequirements {
  s.Requests = v
  return s
}

type CreatePvcsTypedLocalObjectReference struct {
  // {"en":"the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required", "zh_CN":"被引用资源的组。如果不指定 APIGroup，则指定的 Kind 必须在核心 API 组中。对于任何其它第三方类型，都需要 APIGroup"}
  ApiGroup *string `json:"apiGroup,omitempty" xml:"apiGroup,omitempty"`
  // {"en":" the type of resource being referenced", "zh_CN":"被引用的资源的类型"}
  Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
  // {"en":"the name of resource being referenced", "zh_CN":"被引用的资源的名称"}
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreatePvcsTypedLocalObjectReference) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsTypedLocalObjectReference) GoString() string {
  return s.String()
}

func (s *CreatePvcsTypedLocalObjectReference) SetApiGroup(v string) *CreatePvcsTypedLocalObjectReference {
  s.ApiGroup = &v
  return s
}

func (s *CreatePvcsTypedLocalObjectReference) SetKind(v string) *CreatePvcsTypedLocalObjectReference {
  s.Kind = &v
  return s
}

func (s *CreatePvcsTypedLocalObjectReference) SetName(v string) *CreatePvcsTypedLocalObjectReference {
  s.Name = &v
  return s
}

type CreatePvcsPersistentVolumeClaimStatus struct {
  // {"en":"represents the current phase of CreatePvcsPersistentVolumeClaim", "zh_CN":"表示 CreatePvcsPersistentVolumeClaim 的当前阶段"}
  Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
  // {"en":"contains the actual access modes the volume backing the PVC has", "zh_CN":"包含支持 PVC 的卷所具有的实际访问模式"}
  AccessModes []*string `json:"accessModes,omitempty" xml:"accessModes,omitempty" type:"Repeated"`
  // {"en":"represents the actual resources of the underlying volume", "zh_CN":"表示底层卷的实际资源"}
  Capacity map[string]*int64 `json:"capacity,omitempty" xml:"capacity,omitempty"`
  // {"en":"the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'", "zh_CN":"持久卷声明的当前的状况。 如果正在调整底层持久卷的大小，则状况将被设为 “ResizeStarted”"}
  Conditions []*CreatePvcsPersistentVolumeClaimCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
}

func (s CreatePvcsPersistentVolumeClaimStatus) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsPersistentVolumeClaimStatus) GoString() string {
  return s.String()
}

func (s *CreatePvcsPersistentVolumeClaimStatus) SetPhase(v string) *CreatePvcsPersistentVolumeClaimStatus {
  s.Phase = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimStatus) SetAccessModes(v []*string) *CreatePvcsPersistentVolumeClaimStatus {
  s.AccessModes = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimStatus) SetCapacity(v map[string]*int64) *CreatePvcsPersistentVolumeClaimStatus {
  s.Capacity = v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimStatus) SetConditions(v []*CreatePvcsPersistentVolumeClaimCondition) *CreatePvcsPersistentVolumeClaimStatus {
  s.Conditions = v
  return s
}

type CreatePvcsPersistentVolumeClaimCondition struct {
  // {"en":"type, required", "zh_CN":"类型，必需"}
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // {"en":"status, required", "zh_CN":"状态，必需"}
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // {"en":"reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized", "zh_CN":"reason 是唯一的，它应该是一个机器可理解的简短字符串，指明上次状况转换的原因。 如果它报告 “ResizeStarted”，则意味着正在调整底层持久卷的大小"}
  Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
  // {"en":" the human-readable message indicating details about last transition", "zh_CN":"人类可读的消息，指示有关上一次转换的详细信息"}
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreatePvcsPersistentVolumeClaimCondition) String() string {
  return tea.Prettify(s)
}

func (s CreatePvcsPersistentVolumeClaimCondition) GoString() string {
  return s.String()
}

func (s *CreatePvcsPersistentVolumeClaimCondition) SetType(v string) *CreatePvcsPersistentVolumeClaimCondition {
  s.Type = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimCondition) SetStatus(v string) *CreatePvcsPersistentVolumeClaimCondition {
  s.Status = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimCondition) SetReason(v string) *CreatePvcsPersistentVolumeClaimCondition {
  s.Reason = &v
  return s
}

func (s *CreatePvcsPersistentVolumeClaimCondition) SetMessage(v string) *CreatePvcsPersistentVolumeClaimCondition {
  s.Message = &v
  return s
}




type DeletePvcFromEdgeRequest struct {
}

func (s DeletePvcFromEdgeRequest) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcFromEdgeRequest) GoString() string {
  return s.String()
}

type DeletePvcFromEdgeResponse struct {
  // {"en":"response code", "zh_CN":"请求返回码"}
  Code *int64 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  // {"en":"response message", "zh_CN":"请求返回信息"}
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  // {"en":"requestId", "zh_CN":"请求识别码"}
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty" require:"true"`
}

func (s DeletePvcFromEdgeResponse) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcFromEdgeResponse) GoString() string {
  return s.String()
}

func (s *DeletePvcFromEdgeResponse) SetCode(v int64) *DeletePvcFromEdgeResponse {
  s.Code = &v
  return s
}

func (s *DeletePvcFromEdgeResponse) SetMsg(v string) *DeletePvcFromEdgeResponse {
  s.Msg = &v
  return s
}

func (s *DeletePvcFromEdgeResponse) SetRequestId(v string) *DeletePvcFromEdgeResponse {
  s.RequestId = &v
  return s
}

type DeletePvcFromEdgePaths struct {
  // {"en":"namespace info", "zh_CN":"pvc namespace"}
  Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty" require:"true"`
  // {"en":"pvc name", "zh_CN":"pvc name"}
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DeletePvcFromEdgePaths) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcFromEdgePaths) GoString() string {
  return s.String()
}

func (s *DeletePvcFromEdgePaths) SetNamespace(v string) *DeletePvcFromEdgePaths {
  s.Namespace = &v
  return s
}

func (s *DeletePvcFromEdgePaths) SetName(v string) *DeletePvcFromEdgePaths {
  s.Name = &v
  return s
}

type DeletePvcFromEdgeParameters struct {
  // {"en":"pvc clusters,multiple separated by commas", "zh_CN":"pvc所在的集群，多个以逗号隔开"}
  Clusters *string `json:"clusters,omitempty" xml:"clusters,omitempty" require:"true"`
}

func (s DeletePvcFromEdgeParameters) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcFromEdgeParameters) GoString() string {
  return s.String()
}

func (s *DeletePvcFromEdgeParameters) SetClusters(v string) *DeletePvcFromEdgeParameters {
  s.Clusters = &v
  return s
}

type DeletePvcFromEdgeRequestHeader struct {
}

func (s DeletePvcFromEdgeRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcFromEdgeRequestHeader) GoString() string {
  return s.String()
}

type DeletePvcFromEdgeResponseHeader struct {
}

func (s DeletePvcFromEdgeResponseHeader) String() string {
  return tea.Prettify(s)
}

func (s DeletePvcFromEdgeResponseHeader) GoString() string {
  return s.String()
}





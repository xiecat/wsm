package dynamic

var CLASS_NAMES = []string{
	"org.apache.coyote.AbstractTypeResolver",
	"org.apache.coyote.AnnotationIntrospector",
	"org.apache.coyote.BeanDescription",
	"org.apache.coyote.BeanProperty",
	"org.apache.coyote.DatabindContext",
	"org.apache.coyote.DeserializationConfig",
	"org.apache.coyote.DeserializationContext",
	"org.apache.coyote.DeserializationFeature",
	"org.apache.coyote.InjectableValues",
	"org.apache.coyote.JavaType",
	"org.apache.coyote.JsonDeserializer",
	"org.apache.coyote.JsonMappingException",
	"org.apache.coyote.JsonNode",
	"org.apache.coyote.JsonSerializable",
	"org.apache.coyote.JsonSerializer",
	"org.apache.coyote.KeyDeserializer",
	"org.apache.coyote.MapperFeature",
	"org.apache.coyote.MappingIterator",
	"org.apache.coyote.MappingJsonFactory",
	"org.apache.coyote.Module",
	"org.apache.coyote.ObjectMapper",
	"org.apache.coyote.ObjectReader",
	"org.apache.coyote.ObjectWriter",
	"org.apache.coyote.PropertyMetadata",
	"org.apache.coyote.PropertyName",
	"org.apache.coyote.PropertyNamingStrategy",
	"org.apache.coyote.RuntimeJsonMappingException",
	"org.apache.coyote.SequenceWriter",
	"org.apache.coyote.SerializationConfig",
	"org.apache.coyote.SerializationFeature",
	"org.apache.coyote.SerializerProvider",
	"org.apache.coyote.annotation.JacksonStdImpl",
	"org.apache.coyote.annotation.JsonAppend",
	"org.apache.coyote.annotation.JsonDeserialize",
	"org.apache.coyote.annotation.JsonNaming",
	"org.apache.coyote.annotation.JsonPOJOBuilder",
	"org.apache.coyote.annotation.JsonSerialize",
	"org.apache.coyote.annotation.JsonTypeIdResolver",
	"org.apache.coyote.annotation.JsonTypeResolver",
	"org.apache.coyote.annotation.JsonValueInstantiator",
	"org.apache.coyote.annotation.NoClass",
	"org.apache.coyote.cfg.BaseSettings",
	"org.apache.coyote.cfg.ConfigFeature",
	"org.apache.coyote.cfg.ConfigOverride",
	"org.apache.coyote.cfg.ConfigOverrides",
	"org.apache.coyote.cfg.ContextAttributes",
	"org.apache.coyote.cfg.DeserializerFactoryConfig",
	"org.apache.coyote.cfg.HandlerInstantiator",
	"org.apache.coyote.cfg.MapperBuilder",
	"org.apache.coyote.cfg.MapperConfig",
	"org.apache.coyote.cfg.MapperConfigBase",
	"org.apache.coyote.cfg.MutableConfigOverride",
	"org.apache.coyote.cfg.PackageVersion",
	"org.apache.coyote.cfg.SerializerFactoryConfig",
	"org.apache.coyote.deser.AbstractDeserializer",
	"org.apache.coyote.deser.BasicDeserializerFactory",
	"org.apache.coyote.deser.BeanDeserializer",
	"org.apache.coyote.deser.BeanDeserializerBase",
	"org.apache.coyote.deser.BeanDeserializerBuilder",
	"org.apache.coyote.deser.BeanDeserializerFactory",
	"org.apache.coyote.deser.BeanDeserializerModifier",
	"org.apache.coyote.deser.BuilderBasedDeserializer",
	"org.apache.coyote.deser.ContextualDeserializer",
	"org.apache.coyote.deser.ContextualKeyDeserializer",
	"org.apache.coyote.deser.CreatorProperty",
	"org.apache.coyote.deser.DataFormatReaders",
	"org.apache.coyote.deser.DefaultDeserializationContext",
	"org.apache.coyote.deser.DeserializationProblemHandler",
	"org.apache.coyote.deser.DeserializerCache",
	"org.apache.coyote.deser.DeserializerFactory",
	"org.apache.coyote.deser.Deserializers",
	"org.apache.coyote.deser.KeyDeserializers",
	"org.apache.coyote.deser.NullValueProvider",
	"org.apache.coyote.deser.ResolvableDeserializer",
	"org.apache.coyote.deser.SettableAnyProperty",
	"org.apache.coyote.deser.SettableBeanProperty",
	"org.apache.coyote.deser.UnresolvedForwardReference",
	"org.apache.coyote.deser.UnresolvedId",
	"org.apache.coyote.deser.ValueInstantiator",
	"org.apache.coyote.deser.ValueInstantiators",
	"org.apache.coyote.deser.impl.BeanAsArrayBuilderDeserializer",
	"org.apache.coyote.deser.impl.BeanAsArrayDeserializer",
	"org.apache.coyote.deser.impl.BeanPropertyMap",
	"org.apache.coyote.deser.impl.CreatorCandidate",
	"org.apache.coyote.deser.impl.CreatorCollector",
	"org.apache.coyote.deser.impl.ErrorThrowingDeserializer",
	"org.apache.coyote.deser.impl.ExternalTypeHandler",
	"org.apache.coyote.deser.impl.FailingDeserializer",
	"org.apache.coyote.deser.impl.FieldProperty",
	"org.apache.coyote.deser.impl.InnerClassProperty",
	"org.apache.coyote.deser.impl.JDKValueInstantiators",
	"org.apache.coyote.deser.impl.JavaUtilCollectionsDeserializers",
	"org.apache.coyote.deser.impl.ManagedReferenceProperty",
	"org.apache.coyote.deser.impl.MergingSettableBeanProperty",
	"org.apache.coyote.deser.impl.MethodProperty",
	"org.apache.coyote.deser.impl.NullsAsEmptyProvider",
	"org.apache.coyote.deser.impl.NullsConstantProvider",
	"org.apache.coyote.deser.impl.NullsFailProvider",
	"org.apache.coyote.deser.impl.ObjectIdReader",
	"org.apache.coyote.deser.impl.ObjectIdReferenceProperty",
	"org.apache.coyote.deser.impl.ObjectIdValueProperty",
	"org.apache.coyote.deser.impl.PropertyBasedCreator",
	"org.apache.coyote.deser.impl.PropertyBasedObjectIdGenerator",
	"org.apache.coyote.deser.impl.PropertyValue",
	"org.apache.coyote.deser.impl.PropertyValueBuffer",
	"org.apache.coyote.deser.impl.ReadableObjectId",
	"org.apache.coyote.deser.impl.SetterlessProperty",
	"org.apache.coyote.deser.impl.TypeWrappedDeserializer",
	"org.apache.coyote.deser.impl.UnwrappedPropertyHandler",
	"org.apache.coyote.deser.impl.ValueInjector",
	"org.apache.coyote.deser.std.ArrayBlockingQueueDeserializer",
	"org.apache.coyote.deser.std.AtomicBooleanDeserializer",
	"org.apache.coyote.deser.std.AtomicReferenceDeserializer",
	"org.apache.coyote.deser.std.BaseNodeDeserializer",
	"org.apache.coyote.deser.std.ByteBufferDeserializer",
	"org.apache.coyote.deser.std.CollectionDeserializer",
	"org.apache.coyote.deser.std.ContainerDeserializerBase",
	"org.apache.coyote.deser.std.DateDeserializers",
	"org.apache.coyote.deser.std.DelegatingDeserializer",
	"org.apache.coyote.deser.std.EnumDeserializer",
	"org.apache.coyote.deser.std.EnumMapDeserializer",
	"org.apache.coyote.deser.std.EnumSetDeserializer",
	"org.apache.coyote.deser.std.FactoryBasedEnumDeserializer",
	"org.apache.coyote.deser.std.FromStringDeserializer",
	"org.apache.coyote.deser.std.JdkDeserializers",
	"org.apache.coyote.deser.std.JsonLocationInstantiator",
	"org.apache.coyote.deser.std.JsonNodeDeserializer",
	"org.apache.coyote.deser.std.MapDeserializer",
	"org.apache.coyote.deser.std.MapEntryDeserializer",
	"org.apache.coyote.deser.std.NullifyingDeserializer",
	"org.apache.coyote.deser.std.NumberDeserializers",
	"org.apache.coyote.deser.std.ObjectArrayDeserializer",
	"org.apache.coyote.deser.std.PrimitiveArrayDeserializers",
	"org.apache.coyote.deser.std.ReferenceTypeDeserializer",
	"org.apache.coyote.deser.std.StackTraceElementDeserializer",
	"org.apache.coyote.deser.std.StdDelegatingDeserializer",
	"org.apache.coyote.deser.std.StdDeserializer",
	"org.apache.coyote.deser.std.StdKeyDeserializer",
	"org.apache.coyote.deser.std.StdKeyDeserializers",
	"org.apache.coyote.deser.std.StdNodeBasedDeserializer",
	"org.apache.coyote.deser.std.StdScalarDeserializer",
	"org.apache.coyote.deser.std.StdValueInstantiator",
	"org.apache.coyote.deser.std.StringArrayDeserializer",
	"org.apache.coyote.deser.std.StringCollectionDeserializer",
	"org.apache.coyote.deser.std.StringDeserializer",
	"org.apache.coyote.deser.std.ThrowableDeserializer",
	"org.apache.coyote.deser.std.TokenBufferDeserializer",
	"org.apache.coyote.deser.std.UUIDDeserializer",
	"org.apache.coyote.deser.std.UntypedObjectDeserializer",
	"org.apache.coyote.exc.IgnoredPropertyException",
	"org.apache.coyote.exc.InvalidDefinitionException",
	"org.apache.coyote.exc.InvalidFormatException",
	"org.apache.coyote.exc.InvalidNullException",
	"org.apache.coyote.exc.InvalidTypeIdException",
	"org.apache.coyote.exc.MismatchedInputException",
	"org.apache.coyote.exc.PropertyBindingException",
	"org.apache.coyote.exc.UnrecognizedPropertyException",
	"org.apache.coyote.exc.ValueInstantiationException",
	"org.apache.coyote.ext.CoreXMLDeserializers",
	"org.apache.coyote.ext.CoreXMLSerializers",
	"org.apache.coyote.ext.DOMDeserializer",
	"org.apache.coyote.ext.DOMSerializer",
	"org.apache.coyote.ext.Java7Handlers",
	"org.apache.coyote.ext.Java7HandlersImpl",
	"org.apache.coyote.ext.Java7Support",
	"org.apache.coyote.ext.Java7SupportImpl",
	"org.apache.coyote.ext.NioPathDeserializer",
	"org.apache.coyote.ext.NioPathSerializer",
	"org.apache.coyote.ext.OptionalHandlerFactory",
	"org.apache.coyote.introspect.Annotated",
	"org.apache.coyote.introspect.AnnotatedClass",
	"org.apache.coyote.introspect.AnnotatedClassResolver",
	"org.apache.coyote.introspect.AnnotatedConstructor",
	"org.apache.coyote.introspect.AnnotatedCreatorCollector",
	"org.apache.coyote.introspect.AnnotatedField",
	"org.apache.coyote.introspect.AnnotatedFieldCollector",
	"org.apache.coyote.introspect.AnnotatedMember",
	"org.apache.coyote.introspect.AnnotatedMethod",
	"org.apache.coyote.introspect.AnnotatedMethodCollector",
	"org.apache.coyote.introspect.AnnotatedMethodMap",
	"org.apache.coyote.introspect.AnnotatedParameter",
	"org.apache.coyote.introspect.AnnotatedWithParams",
	"org.apache.coyote.introspect.AnnotationCollector",
	"org.apache.coyote.introspect.AnnotationIntrospectorPair",
	"org.apache.coyote.introspect.AnnotationMap",
	"org.apache.coyote.introspect.BasicBeanDescription",
	"org.apache.coyote.introspect.BasicClassIntrospector",
	"org.apache.coyote.introspect.BeanPropertyDefinition",
	"org.apache.coyote.introspect.ClassIntrospector",
	"org.apache.coyote.introspect.CollectorBase",
	"org.apache.coyote.introspect.ConcreteBeanPropertyBase",
	"org.apache.coyote.introspect.JacksonAnnotationIntrospector",
	"org.apache.coyote.introspect.MemberKey",
	"org.apache.coyote.introspect.NopAnnotationIntrospector",
	"org.apache.coyote.introspect.ObjectIdInfo",
	"org.apache.coyote.introspect.POJOPropertiesCollector",
	"org.apache.coyote.introspect.POJOPropertyBuilder",
	"org.apache.coyote.introspect.SimpleMixInResolver",
	"org.apache.coyote.introspect.TypeResolutionContext",
	"org.apache.coyote.introspect.VirtualAnnotatedMember",
	"org.apache.coyote.introspect.VisibilityChecker",
	"org.apache.coyote.introspect.WithMember",
	"org.apache.coyote.json.JsonMapper",
	"org.apache.coyote.jsonFormatVisitors.JsonAnyFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonArrayFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonBooleanFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonFormatTypes",
	"org.apache.coyote.jsonFormatVisitors.JsonFormatVisitable",
	"org.apache.coyote.jsonFormatVisitors.JsonFormatVisitorWithSerializerProvider",
	"org.apache.coyote.jsonFormatVisitors.JsonFormatVisitorWrapper",
	"org.apache.coyote.jsonFormatVisitors.JsonIntegerFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonMapFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonNullFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonNumberFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonObjectFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonStringFormatVisitor",
	"org.apache.coyote.jsonFormatVisitors.JsonValueFormat",
	"org.apache.coyote.jsonFormatVisitors.JsonValueFormatVisitor",
	"org.apache.coyote.jsonschema.JsonSchema",
	"org.apache.coyote.jsonschema.JsonSerializableSchema",
	"org.apache.coyote.jsonschema.SchemaAware",
	"org.apache.coyote.jsontype.BasicPolymorphicTypeValidator",
	"org.apache.coyote.jsontype.NamedType",
	"org.apache.coyote.jsontype.PolymorphicTypeValidator",
	"org.apache.coyote.jsontype.SubtypeResolver",
	"org.apache.coyote.jsontype.TypeDeserializer",
	"org.apache.coyote.jsontype.TypeIdResolver",
	"org.apache.coyote.jsontype.TypeResolverBuilder",
	"org.apache.coyote.jsontype.TypeSerializer",
	"org.apache.coyote.jsontype.impl.AsArrayTypeDeserializer",
	"org.apache.coyote.jsontype.impl.AsArrayTypeSerializer",
	"org.apache.coyote.jsontype.impl.AsExistingPropertyTypeSerializer",
	"org.apache.coyote.jsontype.impl.AsExternalTypeDeserializer",
	"org.apache.coyote.jsontype.impl.AsExternalTypeSerializer",
	"org.apache.coyote.jsontype.impl.AsPropertyTypeDeserializer",
	"org.apache.coyote.jsontype.impl.AsPropertyTypeSerializer",
	"org.apache.coyote.jsontype.impl.AsWrapperTypeDeserializer",
	"org.apache.coyote.jsontype.impl.AsWrapperTypeSerializer",
	"org.apache.coyote.jsontype.impl.ClassNameIdResolver",
	"org.apache.coyote.jsontype.impl.LaissezFaireSubTypeValidator",
	"org.apache.coyote.jsontype.impl.MinimalClassNameIdResolver",
	"org.apache.coyote.jsontype.impl.StdSubtypeResolver",
	"org.apache.coyote.jsontype.impl.StdTypeResolverBuilder",
	"org.apache.coyote.jsontype.impl.SubTypeValidator",
	"org.apache.coyote.jsontype.impl.TypeDeserializerBase",
	"org.apache.coyote.jsontype.impl.TypeIdResolverBase",
	"org.apache.coyote.jsontype.impl.TypeNameIdResolver",
	"org.apache.coyote.jsontype.impl.TypeSerializerBase",
	"org.apache.coyote.module.SimpleAbstractTypeResolver",
	"org.apache.coyote.module.SimpleDeserializers",
	"org.apache.coyote.module.SimpleKeyDeserializers",
	"org.apache.coyote.module.SimpleModule",
	"org.apache.coyote.module.SimpleSerializers",
	"org.apache.coyote.module.SimpleValueInstantiators",
	"org.apache.coyote.node.ArrayNode",
	"org.apache.coyote.node.BaseJsonNode",
	"org.apache.coyote.node.BigIntegerNode",
	"org.apache.coyote.node.BinaryNode",
	"org.apache.coyote.node.BooleanNode",
	"org.apache.coyote.node.ContainerNode",
	"org.apache.coyote.node.DecimalNode",
	"org.apache.coyote.node.DoubleNode",
	"org.apache.coyote.node.FloatNode",
	"org.apache.coyote.node.IntNode",
	"org.apache.coyote.node.InternalNodeMapper",
	"org.apache.coyote.node.JsonNodeCreator",
	"org.apache.coyote.node.JsonNodeFactory",
	"org.apache.coyote.node.JsonNodeType",
	"org.apache.coyote.node.LongNode",
	"org.apache.coyote.node.MissingNode",
	"org.apache.coyote.node.NodeCursor",
	"org.apache.coyote.node.NodeSerialization",
	"org.apache.coyote.node.NullNode",
	"org.apache.coyote.node.NumericNode",
	"org.apache.coyote.node.ObjectNode",
	"org.apache.coyote.node.POJONode",
	"org.apache.coyote.node.ShortNode",
	"org.apache.coyote.node.TextNode",
	"org.apache.coyote.node.TreeTraversingParser",
	"org.apache.coyote.node.ValueNode",
	"org.apache.coyote.ser.AnyGetterWriter",
	"org.apache.coyote.ser.BasicSerializerFactory",
	"org.apache.coyote.ser.BeanPropertyFilter",
	"org.apache.coyote.ser.BeanPropertyWriter",
	"org.apache.coyote.ser.BeanSerializer",
	"org.apache.coyote.ser.BeanSerializerBuilder",
	"org.apache.coyote.ser.BeanSerializerFactory",
	"org.apache.coyote.ser.BeanSerializerModifier",
	"org.apache.coyote.ser.ContainerSerializer",
	"org.apache.coyote.ser.ContextualSerializer",
	"org.apache.coyote.ser.DefaultSerializerProvider",
	"org.apache.coyote.ser.FilterProvider",
	"org.apache.coyote.ser.PropertyBuilder",
	"org.apache.coyote.ser.PropertyFilter",
	"org.apache.coyote.ser.PropertyWriter",
	"org.apache.coyote.ser.ResolvableSerializer",
	"org.apache.coyote.ser.SerializerCache",
	"org.apache.coyote.ser.SerializerFactory",
	"org.apache.coyote.ser.Serializers",
	"org.apache.coyote.ser.VirtualBeanPropertyWriter",
	"org.apache.coyote.ser.impl.AttributePropertyWriter",
	"org.apache.coyote.ser.impl.BeanAsArraySerializer",
	"org.apache.coyote.ser.impl.FailingSerializer",
	"org.apache.coyote.ser.impl.FilteredBeanPropertyWriter",
	"org.apache.coyote.ser.impl.IndexedListSerializer",
	"org.apache.coyote.ser.impl.IndexedStringListSerializer",
	"org.apache.coyote.ser.impl.IteratorSerializer",
	"org.apache.coyote.ser.impl.MapEntrySerializer",
	"org.apache.coyote.ser.impl.ObjectIdWriter",
	"org.apache.coyote.ser.impl.PropertyBasedObjectIdGenerator",
	"org.apache.coyote.ser.impl.PropertySerializerMap",
	"org.apache.coyote.ser.impl.ReadOnlyClassToSerializerMap",
	"org.apache.coyote.ser.impl.SimpleBeanPropertyFilter",
	"org.apache.coyote.ser.impl.SimpleFilterProvider",
	"org.apache.coyote.ser.impl.StringArraySerializer",
	"org.apache.coyote.ser.impl.StringCollectionSerializer",
	"org.apache.coyote.ser.impl.TypeWrappedSerializer",
	"org.apache.coyote.ser.impl.UnknownSerializer",
	"org.apache.coyote.ser.impl.UnwrappingBeanPropertyWriter",
	"org.apache.coyote.ser.impl.UnwrappingBeanSerializer",
	"org.apache.coyote.ser.impl.WritableObjectId",
	"org.apache.coyote.ser.std.ArraySerializerBase",
	"org.apache.coyote.ser.std.AsArraySerializerBase",
	"org.apache.coyote.ser.std.AtomicReferenceSerializer",
	"org.apache.coyote.ser.std.BeanSerializerBase",
	"org.apache.coyote.ser.std.BooleanSerializer",
	"org.apache.coyote.ser.std.ByteArraySerializer",
	"org.apache.coyote.ser.std.ByteBufferSerializer",
	"org.apache.coyote.ser.std.CalendarSerializer",
	"org.apache.coyote.ser.std.ClassSerializer",
	"org.apache.coyote.ser.std.CollectionSerializer",
	"org.apache.coyote.ser.std.DateSerializer",
	"org.apache.coyote.ser.std.DateTimeSerializerBase",
	"org.apache.coyote.ser.std.EnumSerializer",
	"org.apache.coyote.ser.std.EnumSetSerializer",
	"org.apache.coyote.ser.std.FileSerializer",
	"org.apache.coyote.ser.std.InetAddressSerializer",
	"org.apache.coyote.ser.std.InetSocketAddressSerializer",
	"org.apache.coyote.ser.std.IterableSerializer",
	"org.apache.coyote.ser.std.JsonValueSerializer",
	"org.apache.coyote.ser.std.MapProperty",
	"org.apache.coyote.ser.std.MapSerializer",
	"org.apache.coyote.ser.std.NonTypedScalarSerializerBase",
	"org.apache.coyote.ser.std.NullSerializer",
	"org.apache.coyote.ser.std.NumberSerializer",
	"org.apache.coyote.ser.std.NumberSerializers",
	"org.apache.coyote.ser.std.ObjectArraySerializer",
	"org.apache.coyote.ser.std.RawSerializer",
	"org.apache.coyote.ser.std.ReferenceTypeSerializer",
	"org.apache.coyote.ser.std.SerializableSerializer",
	"org.apache.coyote.ser.std.SqlDateSerializer",
	"org.apache.coyote.ser.std.SqlTimeSerializer",
	"org.apache.coyote.ser.std.StaticListSerializerBase",
	"org.apache.coyote.ser.std.StdArraySerializers",
	"org.apache.coyote.ser.std.StdDelegatingSerializer",
	"org.apache.coyote.ser.std.StdJdkSerializers",
	"org.apache.coyote.ser.std.StdKeySerializer",
	"org.apache.coyote.ser.std.StdKeySerializers",
	"org.apache.coyote.ser.std.StdScalarSerializer",
	"org.apache.coyote.ser.std.StdSerializer",
	"org.apache.coyote.ser.std.StringSerializer",
	"org.apache.coyote.ser.std.TimeZoneSerializer",
	"org.apache.coyote.ser.std.ToStringSerializer",
	"org.apache.coyote.ser.std.ToStringSerializerBase",
	"org.apache.coyote.ser.std.TokenBufferSerializer",
	"org.apache.coyote.ser.std.UUIDSerializer",
	"org.apache.coyote.type.ArrayType",
	"org.apache.coyote.type.ClassKey",
	"org.apache.coyote.type.ClassStack",
	"org.apache.coyote.type.CollectionLikeType",
	"org.apache.coyote.type.CollectionType",
	"org.apache.coyote.type.MapLikeType",
	"org.apache.coyote.type.MapType",
	"org.apache.coyote.type.PlaceholderForType",
	"org.apache.coyote.type.ReferenceType",
	"org.apache.coyote.type.ResolvedRecursiveType",
	"org.apache.coyote.type.SimpleType",
	"org.apache.coyote.type.TypeBase",
	"org.apache.coyote.type.TypeBindings",
	"org.apache.coyote.type.TypeFactory",
	"org.apache.coyote.type.TypeModifier",
	"org.apache.coyote.type.TypeParser",
	"org.apache.coyote.util.AccessPattern",
	"org.apache.coyote.util.Annotations",
	"org.apache.coyote.util.ArrayBuilders",
	"org.apache.coyote.util.ArrayIterator",
	"org.apache.coyote.util.BeanUtil",
	"org.apache.coyote.util.ByteBufferBackedInputStream",
	"org.apache.coyote.util.ByteBufferBackedOutputStream",
	"org.apache.coyote.util.ClassUtil",
	"org.apache.coyote.util.CompactStringObjectMap",
	"org.apache.coyote.util.Converter",
	"org.apache.coyote.util.EnumResolver",
	"org.apache.coyote.util.EnumValues",
	"org.apache.coyote.util.ISO8601DateFormat",
	"org.apache.coyote.util.ISO8601Utils",
	"org.apache.coyote.util.JSONPObject",
	"org.apache.coyote.util.JSONWrappedObject",
	"org.apache.coyote.util.LRUMap",
	"org.apache.coyote.util.LinkedNode",
	"org.apache.coyote.util.NameTransformer",
	"org.apache.coyote.util.Named",
	"org.apache.coyote.util.ObjectBuffer",
	"org.apache.coyote.util.PrimitiveArrayBuilder",
	"org.apache.coyote.util.RawValue",
	"org.apache.coyote.util.RootNameLookup",
	"org.apache.coyote.util.SimpleBeanPropertyDefinition",
	"org.apache.coyote.util.StdConverter",
	"org.apache.coyote.util.StdDateFormat",
	"org.apache.coyote.util.TokenBuffer",
	"org.apache.coyote.util.TokenBufferReadContext",
	"org.apache.coyote.util.TypeKey",
	"org.apache.coyote.util.ViewMatcher",
}
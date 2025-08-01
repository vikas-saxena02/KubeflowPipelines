# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: kubernetes_executor_config.proto
# Protobuf Python Version: 6.31.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    31,
    1,
    '',
    'kubernetes_executor_config.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from kfp.pipeline_spec import pipeline_spec_pb2 as pipeline__spec__pb2

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n kubernetes_executor_config.proto\x12\x0ekfp_kubernetes\x1a\x1cgoogle/protobuf/struct.proto\x1a\x13pipeline_spec.proto\"\xb4\x07\n\x18KubernetesExecutorConfig\x12\x38\n\x10secret_as_volume\x18\x01 \x03(\x0b\x32\x1e.kfp_kubernetes.SecretAsVolume\x12\x32\n\rsecret_as_env\x18\x02 \x03(\x0b\x32\x1b.kfp_kubernetes.SecretAsEnv\x12+\n\tpvc_mount\x18\x03 \x03(\x0b\x32\x18.kfp_kubernetes.PvcMount\x12\x33\n\rnode_selector\x18\x04 \x01(\x0b\x32\x1c.kfp_kubernetes.NodeSelector\x12\x31\n\x0cpod_metadata\x18\x05 \x01(\x0b\x32\x1b.kfp_kubernetes.PodMetadata\x12:\n\x11image_pull_secret\x18\x06 \x03(\x0b\x32\x1f.kfp_kubernetes.ImagePullSecret\x12\x19\n\x11image_pull_policy\x18\x07 \x01(\t\x12?\n\x14\x63onfig_map_as_volume\x18\x08 \x03(\x0b\x32!.kfp_kubernetes.ConfigMapAsVolume\x12\x39\n\x11\x63onfig_map_as_env\x18\t \x03(\x0b\x32\x1e.kfp_kubernetes.ConfigMapAsEnv\x12\x1f\n\x17\x61\x63tive_deadline_seconds\x18\n \x01(\x03\x12\x39\n\x11\x66ield_path_as_env\x18\x0b \x03(\x0b\x32\x1e.kfp_kubernetes.FieldPathAsEnv\x12/\n\x0btolerations\x18\x0c \x03(\x0b\x32\x1a.kfp_kubernetes.Toleration\x12H\n\x18generic_ephemeral_volume\x18\r \x03(\x0b\x32&.kfp_kubernetes.GenericEphemeralVolume\x12\x37\n\rnode_affinity\x18\x0e \x03(\x0b\x32 .kfp_kubernetes.NodeAffinityTerm\x12\x35\n\x0cpod_affinity\x18\x0f \x03(\x0b\x32\x1f.kfp_kubernetes.PodAffinityTerm\x12\x42\n\x15\x65nabled_shared_memory\x18\x10 \x01(\x0b\x32#.kfp_kubernetes.EnabledSharedMemory\x12\x37\n\x10\x65mpty_dir_mounts\x18\x11 \x03(\x0b\x32\x1d.kfp_kubernetes.EmptyDirMount\"8\n\x13\x45nabledSharedMemory\x12\x13\n\x0bvolume_name\x18\x01 \x01(\t\x12\x0c\n\x04size\x18\x02 \x01(\t\"\xb1\x01\n\x0eSecretAsVolume\x12\x17\n\x0bsecret_name\x18\x01 \x01(\tB\x02\x18\x01\x12\x12\n\nmount_path\x18\x02 \x01(\t\x12\x15\n\x08optional\x18\x03 \x01(\x08H\x00\x88\x01\x01\x12N\n\x15secret_name_parameter\x18\x04 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpecB\x0b\n\t_optional\"\xf3\x01\n\x0bSecretAsEnv\x12\x17\n\x0bsecret_name\x18\x01 \x01(\tB\x02\x18\x01\x12\x41\n\nkey_to_env\x18\x02 \x03(\x0b\x32-.kfp_kubernetes.SecretAsEnv.SecretKeyToEnvMap\x12N\n\x15secret_name_parameter\x18\x04 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpec\x1a\x38\n\x11SecretKeyToEnvMap\x12\x12\n\nsecret_key\x18\x01 \x01(\t\x12\x0f\n\x07\x65nv_var\x18\x02 \x01(\t\"\xab\x02\n\x08PvcMount\x12l\n\x15task_output_parameter\x18\x01 \x01(\x0b\x32G.ml_pipelines.TaskInputsSpec.InputParameterSpec.TaskOutputParameterSpecB\x02\x18\x01H\x00\x12\x16\n\x08\x63onstant\x18\x02 \x01(\tB\x02\x18\x01H\x00\x12\'\n\x19\x63omponent_input_parameter\x18\x03 \x01(\tB\x02\x18\x01H\x00\x12\x12\n\nmount_path\x18\x04 \x01(\t\x12K\n\x12pvc_name_parameter\x18\x05 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpecB\x0f\n\rpvc_reference\"\xe4\x01\n\tCreatePvc\x12\x12\n\x08pvc_name\x18\x01 \x01(\tH\x00\x12\x19\n\x0fpvc_name_suffix\x18\x02 \x01(\tH\x00\x12\x14\n\x0c\x61\x63\x63\x65ss_modes\x18\x03 \x03(\t\x12\x0c\n\x04size\x18\x04 \x01(\t\x12\x1d\n\x15\x64\x65\x66\x61ult_storage_class\x18\x05 \x01(\x08\x12\x1a\n\x12storage_class_name\x18\x06 \x01(\t\x12\x13\n\x0bvolume_name\x18\x07 \x01(\t\x12,\n\x0b\x61nnotations\x18\x08 \x01(\x0b\x32\x17.google.protobuf.StructB\x06\n\x04name\"\xbf\x01\n\tDeletePvc\x12h\n\x15task_output_parameter\x18\x01 \x01(\x0b\x32G.ml_pipelines.TaskInputsSpec.InputParameterSpec.TaskOutputParameterSpecH\x00\x12\x12\n\x08\x63onstant\x18\x02 \x01(\tH\x00\x12#\n\x19\x63omponent_input_parameter\x18\x03 \x01(\tH\x00\x42\x0f\n\rpvc_reference\"\xc4\x01\n\x0cNodeSelector\x12\x38\n\x06labels\x18\x01 \x03(\x0b\x32(.kfp_kubernetes.NodeSelector.LabelsEntry\x12K\n\x12node_selector_json\x18\x02 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpec\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xec\x01\n\x0bPodMetadata\x12\x37\n\x06labels\x18\x01 \x03(\x0b\x32\'.kfp_kubernetes.PodMetadata.LabelsEntry\x12\x41\n\x0b\x61nnotations\x18\x02 \x03(\x0b\x32,.kfp_kubernetes.PodMetadata.AnnotationsEntry\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x32\n\x10\x41nnotationsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xbc\x01\n\x11\x43onfigMapAsVolume\x12\x1b\n\x0f\x63onfig_map_name\x18\x01 \x01(\tB\x02\x18\x01\x12\x12\n\nmount_path\x18\x02 \x01(\t\x12\x15\n\x08optional\x18\x03 \x01(\x08H\x00\x88\x01\x01\x12R\n\x19\x63onfig_map_name_parameter\x18\x04 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpecB\x0b\n\t_optional\"\x8b\x02\n\x0e\x43onfigMapAsEnv\x12\x1b\n\x0f\x63onfig_map_name\x18\x01 \x01(\tB\x02\x18\x01\x12G\n\nkey_to_env\x18\x02 \x03(\x0b\x32\x33.kfp_kubernetes.ConfigMapAsEnv.ConfigMapKeyToEnvMap\x12R\n\x19\x63onfig_map_name_parameter\x18\x03 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpec\x1a?\n\x14\x43onfigMapKeyToEnvMap\x12\x16\n\x0e\x63onfig_map_key\x18\x01 \x01(\t\x12\x0f\n\x07\x65nv_var\x18\x02 \x01(\t\"\xcf\x01\n\x16GenericEphemeralVolume\x12\x13\n\x0bvolume_name\x18\x01 \x01(\t\x12\x12\n\nmount_path\x18\x02 \x01(\t\x12\x14\n\x0c\x61\x63\x63\x65ss_modes\x18\x03 \x03(\t\x12\x0c\n\x04size\x18\x04 \x01(\t\x12\x1d\n\x15\x64\x65\x66\x61ult_storage_class\x18\x05 \x01(\x08\x12\x1a\n\x12storage_class_name\x18\x06 \x01(\t\x12-\n\x08metadata\x18\x07 \x01(\x0b\x32\x1b.kfp_kubernetes.PodMetadata\"z\n\x0fImagePullSecret\x12\x17\n\x0bsecret_name\x18\x01 \x01(\tB\x02\x18\x01\x12N\n\x15secret_name_parameter\x18\x02 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpec\"2\n\x0e\x46ieldPathAsEnv\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x12\n\nfield_path\x18\x02 \x01(\t\"\xcc\x01\n\nToleration\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x10\n\x08operator\x18\x02 \x01(\t\x12\r\n\x05value\x18\x03 \x01(\t\x12\x0e\n\x06\x65\x66\x66\x65\x63t\x18\x04 \x01(\t\x12\x1f\n\x12toleration_seconds\x18\x05 \x01(\x03H\x00\x88\x01\x01\x12H\n\x0ftoleration_json\x18\x06 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpecB\x15\n\x13_toleration_seconds\"D\n\x13SelectorRequirement\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x10\n\x08operator\x18\x02 \x01(\t\x12\x0e\n\x06values\x18\x03 \x03(\t\"\xfa\x01\n\x10NodeAffinityTerm\x12>\n\x11match_expressions\x18\x01 \x03(\x0b\x32#.kfp_kubernetes.SelectorRequirement\x12\x39\n\x0cmatch_fields\x18\x02 \x03(\x0b\x32#.kfp_kubernetes.SelectorRequirement\x12\x13\n\x06weight\x18\x03 \x01(\x05H\x00\x88\x01\x01\x12K\n\x12node_affinity_json\x18\x04 \x01(\x0b\x32/.ml_pipelines.TaskInputsSpec.InputParameterSpecB\t\n\x07_weight\"\xa3\x04\n\x0fPodAffinityTerm\x12\x42\n\x15match_pod_expressions\x18\x01 \x03(\x0b\x32#.kfp_kubernetes.SelectorRequirement\x12M\n\x10match_pod_labels\x18\x02 \x03(\x0b\x32\x33.kfp_kubernetes.PodAffinityTerm.MatchPodLabelsEntry\x12\x14\n\x0ctopology_key\x18\x03 \x01(\t\x12\x12\n\nnamespaces\x18\x04 \x03(\t\x12H\n\x1bmatch_namespace_expressions\x18\x05 \x03(\x0b\x32#.kfp_kubernetes.SelectorRequirement\x12Y\n\x16match_namespace_labels\x18\x06 \x03(\x0b\x32\x39.kfp_kubernetes.PodAffinityTerm.MatchNamespaceLabelsEntry\x12\x13\n\x06weight\x18\x07 \x01(\x05H\x00\x88\x01\x01\x12\x11\n\x04\x61nti\x18\x08 \x01(\x08H\x01\x88\x01\x01\x1a\x35\n\x13MatchPodLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a;\n\x19MatchNamespaceLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\t\n\x07_weightB\x07\n\x05_anti\"\x80\x01\n\rEmptyDirMount\x12\x13\n\x0bvolume_name\x18\x01 \x01(\t\x12\x12\n\nmount_path\x18\x02 \x01(\t\x12\x13\n\x06medium\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x17\n\nsize_limit\x18\x04 \x01(\tH\x01\x88\x01\x01\x42\t\n\x07_mediumB\r\n\x0b_size_limitBIZGgithub.com/kubeflow/pipelines/kubernetes_platform/go/kubernetesplatformb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'kubernetes_executor_config_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZGgithub.com/kubeflow/pipelines/kubernetes_platform/go/kubernetesplatform'
  _globals['_SECRETASVOLUME'].fields_by_name['secret_name']._loaded_options = None
  _globals['_SECRETASVOLUME'].fields_by_name['secret_name']._serialized_options = b'\030\001'
  _globals['_SECRETASENV'].fields_by_name['secret_name']._loaded_options = None
  _globals['_SECRETASENV'].fields_by_name['secret_name']._serialized_options = b'\030\001'
  _globals['_PVCMOUNT'].fields_by_name['task_output_parameter']._loaded_options = None
  _globals['_PVCMOUNT'].fields_by_name['task_output_parameter']._serialized_options = b'\030\001'
  _globals['_PVCMOUNT'].fields_by_name['constant']._loaded_options = None
  _globals['_PVCMOUNT'].fields_by_name['constant']._serialized_options = b'\030\001'
  _globals['_PVCMOUNT'].fields_by_name['component_input_parameter']._loaded_options = None
  _globals['_PVCMOUNT'].fields_by_name['component_input_parameter']._serialized_options = b'\030\001'
  _globals['_NODESELECTOR_LABELSENTRY']._loaded_options = None
  _globals['_NODESELECTOR_LABELSENTRY']._serialized_options = b'8\001'
  _globals['_PODMETADATA_LABELSENTRY']._loaded_options = None
  _globals['_PODMETADATA_LABELSENTRY']._serialized_options = b'8\001'
  _globals['_PODMETADATA_ANNOTATIONSENTRY']._loaded_options = None
  _globals['_PODMETADATA_ANNOTATIONSENTRY']._serialized_options = b'8\001'
  _globals['_CONFIGMAPASVOLUME'].fields_by_name['config_map_name']._loaded_options = None
  _globals['_CONFIGMAPASVOLUME'].fields_by_name['config_map_name']._serialized_options = b'\030\001'
  _globals['_CONFIGMAPASENV'].fields_by_name['config_map_name']._loaded_options = None
  _globals['_CONFIGMAPASENV'].fields_by_name['config_map_name']._serialized_options = b'\030\001'
  _globals['_IMAGEPULLSECRET'].fields_by_name['secret_name']._loaded_options = None
  _globals['_IMAGEPULLSECRET'].fields_by_name['secret_name']._serialized_options = b'\030\001'
  _globals['_PODAFFINITYTERM_MATCHPODLABELSENTRY']._loaded_options = None
  _globals['_PODAFFINITYTERM_MATCHPODLABELSENTRY']._serialized_options = b'8\001'
  _globals['_PODAFFINITYTERM_MATCHNAMESPACELABELSENTRY']._loaded_options = None
  _globals['_PODAFFINITYTERM_MATCHNAMESPACELABELSENTRY']._serialized_options = b'8\001'
  _globals['_KUBERNETESEXECUTORCONFIG']._serialized_start=104
  _globals['_KUBERNETESEXECUTORCONFIG']._serialized_end=1052
  _globals['_ENABLEDSHAREDMEMORY']._serialized_start=1054
  _globals['_ENABLEDSHAREDMEMORY']._serialized_end=1110
  _globals['_SECRETASVOLUME']._serialized_start=1113
  _globals['_SECRETASVOLUME']._serialized_end=1290
  _globals['_SECRETASENV']._serialized_start=1293
  _globals['_SECRETASENV']._serialized_end=1536
  _globals['_SECRETASENV_SECRETKEYTOENVMAP']._serialized_start=1480
  _globals['_SECRETASENV_SECRETKEYTOENVMAP']._serialized_end=1536
  _globals['_PVCMOUNT']._serialized_start=1539
  _globals['_PVCMOUNT']._serialized_end=1838
  _globals['_CREATEPVC']._serialized_start=1841
  _globals['_CREATEPVC']._serialized_end=2069
  _globals['_DELETEPVC']._serialized_start=2072
  _globals['_DELETEPVC']._serialized_end=2263
  _globals['_NODESELECTOR']._serialized_start=2266
  _globals['_NODESELECTOR']._serialized_end=2462
  _globals['_NODESELECTOR_LABELSENTRY']._serialized_start=2417
  _globals['_NODESELECTOR_LABELSENTRY']._serialized_end=2462
  _globals['_PODMETADATA']._serialized_start=2465
  _globals['_PODMETADATA']._serialized_end=2701
  _globals['_PODMETADATA_LABELSENTRY']._serialized_start=2417
  _globals['_PODMETADATA_LABELSENTRY']._serialized_end=2462
  _globals['_PODMETADATA_ANNOTATIONSENTRY']._serialized_start=2651
  _globals['_PODMETADATA_ANNOTATIONSENTRY']._serialized_end=2701
  _globals['_CONFIGMAPASVOLUME']._serialized_start=2704
  _globals['_CONFIGMAPASVOLUME']._serialized_end=2892
  _globals['_CONFIGMAPASENV']._serialized_start=2895
  _globals['_CONFIGMAPASENV']._serialized_end=3162
  _globals['_CONFIGMAPASENV_CONFIGMAPKEYTOENVMAP']._serialized_start=3099
  _globals['_CONFIGMAPASENV_CONFIGMAPKEYTOENVMAP']._serialized_end=3162
  _globals['_GENERICEPHEMERALVOLUME']._serialized_start=3165
  _globals['_GENERICEPHEMERALVOLUME']._serialized_end=3372
  _globals['_IMAGEPULLSECRET']._serialized_start=3374
  _globals['_IMAGEPULLSECRET']._serialized_end=3496
  _globals['_FIELDPATHASENV']._serialized_start=3498
  _globals['_FIELDPATHASENV']._serialized_end=3548
  _globals['_TOLERATION']._serialized_start=3551
  _globals['_TOLERATION']._serialized_end=3755
  _globals['_SELECTORREQUIREMENT']._serialized_start=3757
  _globals['_SELECTORREQUIREMENT']._serialized_end=3825
  _globals['_NODEAFFINITYTERM']._serialized_start=3828
  _globals['_NODEAFFINITYTERM']._serialized_end=4078
  _globals['_PODAFFINITYTERM']._serialized_start=4081
  _globals['_PODAFFINITYTERM']._serialized_end=4628
  _globals['_PODAFFINITYTERM_MATCHPODLABELSENTRY']._serialized_start=4494
  _globals['_PODAFFINITYTERM_MATCHPODLABELSENTRY']._serialized_end=4547
  _globals['_PODAFFINITYTERM_MATCHNAMESPACELABELSENTRY']._serialized_start=4549
  _globals['_PODAFFINITYTERM_MATCHNAMESPACELABELSENTRY']._serialized_end=4608
  _globals['_EMPTYDIRMOUNT']._serialized_start=4631
  _globals['_EMPTYDIRMOUNT']._serialized_end=4759
# @@protoc_insertion_point(module_scope)

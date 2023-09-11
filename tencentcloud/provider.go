/*
The TencentCloud provider is used to interact with many resources supported by [TencentCloud](https://intl.cloud.tencent.com).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** From version 1.9.0 (June 18, 2019), the provider start to support Terraform 0.12.x.

# Example Usage

```hcl

	terraform {
	  required_providers {
	    tencentcloud = {
	      source = "tencentcloudstack/tencentcloud"
	    }
	  }
	}

# Configure the TencentCloud Provider

	provider "tencentcloud" {
	  secret_id  = var.secret_id
	  secret_key = var.secret_key
	  region     = var.region
	}

#Configure the TencentCloud Provider with STS

	provider "tencentcloud" {
	  secret_id  = var.secret_id
	  secret_key = var.secret_key
	  region     = var.region
	  assume_role {
	    role_arn         = var.assume_role_arn
	    session_name     = var.session_name
	    session_duration = var.session_duration
	    policy           = var.policy
	  }
	}

```

# Resources List

Provider Data Sources

	tencentcloud_availability_regions
	tencentcloud_availability_zones_by_product
	tencentcloud_availability_zones

Anti-DDoS(DayuV2)

	  Data Source
	    tencentcloud_dayu_eip
		tencentcloud_dayu_l4_rules_v2
		tencentcloud_dayu_l7_rules_v2

	  Resource
		tencentcloud_dayu_eip
		tencentcloud_dayu_l4_rule
		tencentcloud_dayu_l7_rule_v2
		tencentcloud_dayu_ddos_policy_v2
		tencentcloud_dayu_cc_policy_v2
		tencentcloud_dayu_ddos_ip_attachment_v2

Anti-DDoS(Dayu)

	Data Source
	  tencentcloud_dayu_cc_http_policies
	  tencentcloud_dayu_cc_https_policies
	  tencentcloud_dayu_ddos_policies
	  tencentcloud_dayu_ddos_policy_attachments
	  tencentcloud_dayu_ddos_policy_cases
	  tencentcloud_dayu_l4_rules
	  tencentcloud_dayu_l7_rules

	Resource
	  tencentcloud_dayu_cc_http_policy
	  tencentcloud_dayu_cc_https_policy
	  tencentcloud_dayu_ddos_policy
	  tencentcloud_dayu_ddos_policy_attachment
	  tencentcloud_dayu_ddos_policy_case
	  tencentcloud_dayu_l4_rule
	  tencentcloud_dayu_l7_rule

API GateWay

	  Data Source
		tencentcloud_api_gateway_apis
		tencentcloud_api_gateway_services
		tencentcloud_api_gateway_throttling_services
		tencentcloud_api_gateway_throttling_apis
		tencentcloud_api_gateway_usage_plans
		tencentcloud_api_gateway_ip_strategies
		tencentcloud_api_gateway_customer_domains
		tencentcloud_api_gateway_usage_plan_environments
		tencentcloud_api_gateway_api_keys
		tencentcloud_api_gateway_api_docs
		tencentcloud_api_gateway_api_apps

	  Resource
	  	tencentcloud_api_gateway_api
		tencentcloud_api_gateway_service
		tencentcloud_api_gateway_custom_domain
		tencentcloud_api_gateway_usage_plan
		tencentcloud_api_gateway_usage_plan_attachment
		tencentcloud_api_gateway_ip_strategy
		tencentcloud_api_gateway_strategy_attachment
		tencentcloud_api_gateway_api_key
		tencentcloud_api_gateway_api_key_attachment
	    tencentcloud_api_gateway_service_release
		tencentcloud_api_gateway_plugin
		tencentcloud_api_gateway_plugin_attachment
		tencentcloud_api_gateway_api_doc
		tencentcloud_api_gateway_api_app

Cloud Audit(Audit)

	  Data Source
		tencentcloud_audit_cos_regions
		tencentcloud_audit_key_alias
		tencentcloud_audits

	  Resource
		tencentcloud_audit
		tencentcloud_audit_track

Auto Scaling(AS)

	  Data Source
	    tencentcloud_as_scaling_configs
	    tencentcloud_as_scaling_groups
	    tencentcloud_as_scaling_policies
		tencentcloud_as_instances

	  Resource
	    tencentcloud_as_scaling_config
	    tencentcloud_as_scaling_group
	    tencentcloud_as_attachment
	    tencentcloud_as_scaling_policy
	    tencentcloud_as_schedule
	    tencentcloud_as_lifecycle_hook
	    tencentcloud_as_notification
		tencentcloud_as_remove_instances
	    tencentcloud_as_protect_instances
	    tencentcloud_as_start_instances
	    tencentcloud_as_stop_instances

Content Delivery Network(CDN)

	  Data Source
	    tencentcloud_cdn_domains
	    tencentcloud_cdn_domain_verifier

	  Resource
		tencentcloud_cdn_domain
		tencentcloud_cdn_url_push
		tencentcloud_cdn_url_purge

Cloud Kafka(ckafka)

	  Data Source
	    tencentcloud_ckafka_users
	    tencentcloud_ckafka_acls
	    tencentcloud_ckafka_topics
	    tencentcloud_ckafka_instances

	  Resource
		tencentcloud_ckafka_instance
	    tencentcloud_ckafka_user
	    tencentcloud_ckafka_acl
	    tencentcloud_ckafka_topic
		tencentcloud_ckafka_datahub_topic
		tencentcloud_ckafka_connect_resource

Cloud Access Management(CAM)

	  Data Source
	    tencentcloud_cam_group_memberships
	    tencentcloud_cam_group_policy_attachments
	    tencentcloud_cam_groups
	    tencentcloud_cam_policies
	    tencentcloud_cam_role_policy_attachments
	    tencentcloud_cam_roles
	    tencentcloud_cam_saml_providers
	    tencentcloud_cam_user_policy_attachments
	    tencentcloud_cam_users
	    tencentcloud_user_info

	  Resource
	    tencentcloud_cam_role
	    tencentcloud_cam_role_policy_attachment
	    tencentcloud_cam_policy
	    tencentcloud_cam_user
	    tencentcloud_cam_user_policy_attachment
	    tencentcloud_cam_group
	    tencentcloud_cam_group_policy_attachment
	    tencentcloud_cam_group_membership
	    tencentcloud_cam_saml_provider
		tencentcloud_cam_oidc_sso
		tencentcloud_cam_role_sso
		tencentcloud_cam_service_linked_role
		tencentcloud_cam_user_saml_config

Cloud Block Storage(CBS)

	  Data Source
	    tencentcloud_cbs_snapshots
	    tencentcloud_cbs_storages
		tencentcloud_cbs_storages_set
	    tencentcloud_cbs_snapshot_policies

	  Resource
	    tencentcloud_cbs_storage
		tencentcloud_cbs_storage_set
	    tencentcloud_cbs_storage_attachment
		tencentcloud_cbs_storage_set_attachment
	    tencentcloud_cbs_snapshot
	    tencentcloud_cbs_snapshot_policy
	    tencentcloud_cbs_snapshot_policy_attachment
		tencentcloud_cbs_snapshot_share_permission
		tencentcloud_cbs_disk_backup
		tencentcloud_cbs_disk_backup_rollback_operation

Cloud Connect Network(CCN)

	  Data Source
	    tencentcloud_ccn_bandwidth_limits
	    tencentcloud_ccn_instances
		tencentcloud_ccn_cross_border_compliance
		tencentcloud_ccn_tenant_instances
		tencentcloud_ccn_cross_border_flow_monitor
		tencentcloud_ccn_cross_border_region_bandwidth_limits

	  Resource
	    tencentcloud_ccn
	    tencentcloud_ccn_attachment
	    tencentcloud_ccn_bandwidth_limit
		tencentcloud_ccn_routes
		tencentcloud_ccn_instances_accept_attach
		tencentcloud_ccn_instances_reset_attach

CVM Dedicated Host(CDH)

	Data Source
	  tencentcloud_cdh_instances

	Resource
	  tencentcloud_cdh_instance

Cloud File Storage(CFS)

	  Data Source
	    tencentcloud_cfs_access_groups
	    tencentcloud_cfs_access_rules
	    tencentcloud_cfs_file_systems
		tencentcloud_cfs_mount_targets
		tencentcloud_cfs_file_system_clients
		tencentcloud_cfs_available_zone

	  Resource
	    tencentcloud_cfs_file_system
	    tencentcloud_cfs_access_group
	    tencentcloud_cfs_access_rule
		tencentcloud_cfs_auto_snapshot_policy
		tencentcloud_cfs_auto_snapshot_policy_attachment
		tencentcloud_cfs_snapshot
		tencentcloud_cfs_sign_up_cfs_service

Container Cluster

	Data Source
	  tencentcloud_container_cluster_instances
	  tencentcloud_container_clusters

	Resource
	  tencentcloud_container_cluster
	  tencentcloud_container_cluster_instance

Cloud Load Balancer(CLB)

	  Data Source
	    tencentcloud_clb_attachments
	    tencentcloud_clb_instances
	    tencentcloud_clb_listener_rules
	    tencentcloud_clb_listeners
	    tencentcloud_clb_redirections
	    tencentcloud_clb_target_groups

	  Resource
	    tencentcloud_clb_instance
	    tencentcloud_clb_listener
	    tencentcloud_clb_listener_rule
	    tencentcloud_clb_attachment
	    tencentcloud_clb_redirection
	    tencentcloud_lb
	    tencentcloud_alb_server_attachment
	    tencentcloud_clb_target_group
	    tencentcloud_clb_target_group_instance_attachment
	    tencentcloud_clb_target_group_attachment
	    tencentcloud_clb_log_set
		tencentcloud_clb_log_topic
		tencentcloud_clb_customized_config
	    tencentcloud_clb_snat_ip
		tencentcloud_clb_function_targets_attachment
		tencentcloud_clb_instance_sla_config
		tencentcloud_clb_instance_mix_ip_target_config
		tencentcloud_clb_replace_cert_for_lbs

Cloud Object Storage(COS)

	  Data Source
	    tencentcloud_cos_bucket_object
	    tencentcloud_cos_buckets

	  Resource
	    tencentcloud_cos_bucket
	    tencentcloud_cos_bucket_object
	    tencentcloud_cos_bucket_policy
		tencentcloud_cos_bucket_domain_certificate_attachment

Cloud Virtual Machine(CVM)

	  Data Source
	    tencentcloud_image
	    tencentcloud_images
	    tencentcloud_instance_types
	    tencentcloud_instances
		tencentcloud_instances_set
	    tencentcloud_key_pairs
	    tencentcloud_eip
	    tencentcloud_eips
		tencentcloud_eip_address_quota
		tencentcloud_eip_network_account_type
	    tencentcloud_placement_groups
	    tencentcloud_reserved_instance_configs
	    tencentcloud_reserved_instances
		tencentcloud_cvm_instances_modification
		tencentcloud_cvm_instance_vnc_url
		tencentcloud_cvm_disaster_recover_group_quota
		tencentcloud_cvm_chc_hosts
		tencentcloud_cvm_chc_denied_actions

	  Resource
	    tencentcloud_instance
		tencentcloud_instance_set
	    tencentcloud_eip
	    tencentcloud_eip_association
		tencentcloud_eip_address_transform
		tencentcloud_eip_public_address_adjust
		tencentcloud_eip_normal_address_return
	    tencentcloud_key_pair
	    tencentcloud_placement_group
	    tencentcloud_reserved_instance
	    tencentcloud_image
		tencentcloud_cvm_hpc_cluster
		tencentcloud_cvm_launch_template
		tencentcloud_cvm_launch_template_version
		tencentcloud_cvm_launch_template_default_version
		tencentcloud_cvm_security_group_attachment
		tencentcloud_cvm_reboot_instance
		tencentcloud_cvm_chc_config

TDSQL-C MySQL(CynosDB)

	  Data Source
		tencentcloud_cynosdb_clusters
		tencentcloud_cynosdb_instances
		tencentcloud_cynosdb_zone_config
		tencentcloud_cynosdb_accounts
		tencentcloud_cynosdb_cluster_instance_groups
		tencentcloud_cynosdb_cluster_params
		tencentcloud_cynosdb_param_templates

	  Resource
	    tencentcloud_cynosdb_cluster
	    tencentcloud_cynosdb_readonly_instance
		tencentcloud_cynosdb_security_group
		tencentcloud_cynosdb_audit_log_file

Direct Connect(DC)

	Data Source
	  tencentcloud_dc_instances
	  tencentcloud_dcx_instances

	Resource
	  tencentcloud_dcx

Direct Connect Gateway(DCG)

	Data Source
	  tencentcloud_dc_gateway_ccn_routes
	  tencentcloud_dc_gateway_instances

	Resource
	  tencentcloud_dc_gateway
	  tencentcloud_dc_gateway_ccn_route

Domain

	Data Source
	  tencentcloud_domains

Elasticsearch Service(ES)

	Data Source
	  tencentcloud_elasticsearch_instances

	Resource
	  tencentcloud_elasticsearch_instance

Global Application Acceleration(GAAP)

	Data Source
	  tencentcloud_gaap_certificates
	  tencentcloud_gaap_http_domains
	  tencentcloud_gaap_http_rules
	  tencentcloud_gaap_layer4_listeners
	  tencentcloud_gaap_layer7_listeners
	  tencentcloud_gaap_proxies
	  tencentcloud_gaap_realservers
	  tencentcloud_gaap_security_policies
	  tencentcloud_gaap_security_rules
	  tencentcloud_gaap_domain_error_pages

	Resource
	  tencentcloud_gaap_proxy
	  tencentcloud_gaap_realserver
	  tencentcloud_gaap_layer4_listener
	  tencentcloud_gaap_layer7_listener
	  tencentcloud_gaap_http_domain
	  tencentcloud_gaap_http_rule
	  tencentcloud_gaap_certificate
	  tencentcloud_gaap_security_policy
	  tencentcloud_gaap_security_rule
	  tencentcloud_gaap_domain_error_page

Key Management Service(KMS)

	Data Source
	  tencentcloud_kms_keys

	Resource
	  tencentcloud_kms_key
	  tencentcloud_kms_external_key

Tencent Kubernetes Engine(TKE)

	  Data Source
	    tencentcloud_kubernetes_clusters
	    tencentcloud_kubernetes_cluster_levels
	    tencentcloud_kubernetes_charts
	    tencentcloud_kubernetes_cluster_common_names
		tencentcloud_kubernetes_available_cluster_versions
		tencentcloud_kubernetes_cluster_authentication_options

	  Resource
	    tencentcloud_kubernetes_cluster
	    tencentcloud_kubernetes_scale_worker
	    tencentcloud_kubernetes_cluster_attachment
		tencentcloud_kubernetes_node_pool
		tencentcloud_kubernetes_serverless_node_pool
	    tencentcloud_kubernetes_backup_storage_location
	    tencentcloud_kubernetes_auth_attachment
	    tencentcloud_kubernetes_addon_attachment
		tencentcloud_kubernetes_cluster_endpoint

TDMQ for Pulsar(tpulsar)

	  Resource
	    tencentcloud_tdmq_instance
		tencentcloud_tdmq_namespace
		tencentcloud_tdmq_topic
		tencentcloud_tdmq_role
		tencentcloud_tdmq_namespace_role_attachment

TencentDB for MongoDB(mongodb)

	  Data Source
	    tencentcloud_mongodb_instances
	    tencentcloud_mongodb_zone_config
		tencentcloud_mongodb_instance_backups
		tencentcloud_mongodb_instance_connections
		tencentcloud_mongodb_instance_current_op
		tencentcloud_mongodb_instance_params
		tencentcloud_mongodb_instance_slow_log

	  Resource
	    tencentcloud_mongodb_instance
	    tencentcloud_mongodb_sharding_instance
	    tencentcloud_mongodb_standby_instance
		tencentcloud_mongodb_instance_account
		tencentcloud_mongodb_instance_backup

TencentDB for MySQL(cdb)

	  Data Source
	    tencentcloud_mysql_backup_list
	    tencentcloud_mysql_instance
	    tencentcloud_mysql_parameter_list
	    tencentcloud_mysql_default_params
	    tencentcloud_mysql_zone_config
		tencentcloud_mysql_backup_overview
		tencentcloud_mysql_backup_summaries
		tencentcloud_mysql_bin_log
		tencentcloud_mysql_binlog_backup_overview
		tencentcloud_mysql_clone_list
		tencentcloud_mysql_data_backup_overview
		tencentcloud_mysql_db_features
		tencentcloud_mysql_inst_tables
		tencentcloud_mysql_instance_charset
		tencentcloud_mysql_instance_info
		tencentcloud_mysql_instance_param_record
		tencentcloud_mysql_instance_reboot_time
		tencentcloud_mysql_rollback_range_time
		tencentcloud_mysql_slow_log
		tencentcloud_mysql_slow_log_data
		tencentcloud_mysql_supported_privileges
		tencentcloud_mysql_switch_record
		tencentcloud_mysql_user_task

	  Resource
	    tencentcloud_mysql_instance
	    tencentcloud_mysql_readonly_instance
	    tencentcloud_mysql_account
	    tencentcloud_mysql_privilege
	    tencentcloud_mysql_account_privilege
	    tencentcloud_mysql_backup_policy
		tencentcloud_mysql_time_window
		tencentcloud_mysql_param_template
		tencentcloud_mysql_deploy_group
		tencentcloud_mysql_security_groups_attachment
		tencentcloud_mysql_local_binlog_config
		tencentcloud_mysql_audit_log_file

Cloud Monitor(Monitor)

	  Data Source
		tencentcloud_monitor_policy_conditions
		tencentcloud_monitor_data
		tencentcloud_monitor_product_event
		tencentcloud_monitor_binding_objects
		tencentcloud_monitor_policy_groups
		tencentcloud_monitor_product_namespace
		tencentcloud_monitor_alarm_notices

	  Resource
	    tencentcloud_monitor_policy_group
	    tencentcloud_monitor_binding_object
		tencentcloud_monitor_policy_binding_object
	    tencentcloud_monitor_binding_receiver
		tencentcloud_monitor_alarm_policy
		tencentcloud_monitor_tmp_tke_template
		tencentcloud_monitor_tmp_tke_template_attachment
		tencentcloud_monitor_tmp_tke_alert_policy
		tencentcloud_monitor_tmp_tke_config
		tencentcloud_monitor_alarm_notice
		tencentcloud_monitor_tmp_tke_record_rule_yaml
		tencentcloud_monitor_tmp_tke_global_notification
		tencentcloud_monitor_tmp_tke_cluster_agent
		tencentcloud_monitor_tmp_tke_basic_config

Managed Service for Prometheus(TMP)

	  Resource
	  	tencentcloud_monitor_tmp_instance
		tencentcloud_monitor_tmp_alert_rule
		tencentcloud_monitor_tmp_exporter_integration
		tencentcloud_monitor_tmp_cvm_agent
		tencentcloud_monitor_tmp_scrape_job
		tencentcloud_monitor_tmp_recording_rule
		tencentcloud_monitor_tmp_manage_grafana_attachment

TencentCloud Managed Service for Grafana(TCMG)

	  Resource
		tencentcloud_monitor_grafana_instance
		tencentcloud_monitor_grafana_integration
		tencentcloud_monitor_grafana_notification_channel
		tencentcloud_monitor_grafana_plugin
		tencentcloud_monitor_grafana_sso_account
		tencentcloud_monitor_tmp_grafana_config

TencentDB for PostgreSQL(PostgreSQL)

	  Data Source
		tencentcloud_postgresql_instances
		tencentcloud_postgresql_specinfos
		tencentcloud_postgresql_xlogs
		tencentcloud_postgresql_parameter_templates

	  Resource
		tencentcloud_postgresql_instance
		tencentcloud_postgresql_readonly_instance
		tencentcloud_postgresql_readonly_group
		tencentcloud_postgresql_readonly_attachment
		tencentcloud_postgresql_parameter_template

TencentDB for Redis(crs)

	  Data Source
	    tencentcloud_redis_zone_config
	    tencentcloud_redis_instances
		tencentcloud_redis_backup
		tencentcloud_redis_backup_download_info
		tencentcloud_redis_param_records
		tencentcloud_redis_instance_shards
		tencentcloud_redis_instance_zone_info
		tencentcloud_redis_instance_task_list
		tencentcloud_redis_instance_node_info

	  Resource
	    tencentcloud_redis_instance
	    tencentcloud_redis_backup_config
	    tencentcloud_redis_param_template
		tencentcloud_redis_account
		tencentcloud_redis_read_only
		tencentcloud_redis_ssl
		tencentcloud_redis_backup_download_restriction
		tencentcloud_redis_clear_instance_operation
		tencentcloud_redis_renew_instance_operation
		tencentcloud_redis_startup_instance_operation
		tencentcloud_redis_upgrade_cache_version_operation
		tencentcloud_redis_upgrade_multi_zone_operation
		tencentcloud_redis_upgrade_proxy_version_operation
		tencentcloud_redis_maintenance_window
		tencentcloud_redis_replica_readonly
		tencentcloud_redis_switch_master

Serverless Cloud Function(SCF)

	  Data Source
	    tencentcloud_scf_functions
	    tencentcloud_scf_logs
	    tencentcloud_scf_namespaces

	  Resource
	    tencentcloud_scf_function
	    tencentcloud_scf_namespace
		tencentcloud_scf_layer
		tencentcloud_scf_function_alias

SQLServer

	  Data Source
	    tencentcloud_sqlserver_zone_config
		tencentcloud_sqlserver_instances
	    tencentcloud_sqlserver_dbs
		tencentcloud_sqlserver_accounts
		tencentcloud_sqlserver_account_db_attachments
		tencentcloud_sqlserver_backups
	  	tencentcloud_sqlserver_readonly_groups
		tencentcloud_sqlserver_publish_subscribes
		tencentcloud_sqlserver_basic_instances
		tencentcloud_sqlserver_backup_commands
		tencentcloud_sqlserver_backup_by_flow_id
		tencentcloud_sqlserver_backup_upload_size
		tencentcloud_sqlserver_cross_region_zone
		tencentcloud_sqlserver_db_charsets
		tencentcloud_sqlserver_instance_param_records
		tencentcloud_sqlserver_project_security_groups
		tencentcloud_sqlserver_regions
		tencentcloud_sqlserver_rollback_time
		tencentcloud_sqlserver_slowlogs
		tencentcloud_sqlserver_upload_backup_info
		tencentcloud_sqlserver_upload_incremental_info

	  Resource
		tencentcloud_sqlserver_instance
		tencentcloud_sqlserver_readonly_instance
	    tencentcloud_sqlserver_db
		tencentcloud_sqlserver_account
		tencentcloud_sqlserver_account_db_attachment
		tencentcloud_sqlserver_publish_subscribe
		tencentcloud_sqlserver_basic_instance
		tencentcloud_sqlserver_migration
		tencentcloud_sqlserver_config_backup_strategy
		tencentcloud_sqlserver_general_backup
		tencentcloud_sqlserver_general_clone
	    tencentcloud_sqlserver_full_backup_migration
		tencentcloud_sqlserver_incre_backup_migration
	    tencentcloud_sqlserver_business_intelligence_file
		tencentcloud_sqlserver_business_intelligence_instance
		tencentcloud_sqlserver_general_communication

SSL Certificates

	  Data Source
	    tencentcloud_ssl_certificates

	  Resource
	    tencentcloud_ssl_certificate
	    tencentcloud_ssl_pay_certificate
		tencentcloud_ssl_free_certificate

Secrets Manager(SSM)

	Data Source
	  tencentcloud_ssm_secrets
	  tencentcloud_ssm_secret_versions

	Resource
	  tencentcloud_ssm_secret
	  tencentcloud_ssm_secret_version

TcaplusDB

	Data Source
	  tencentcloud_tcaplus_clusters
	  tencentcloud_tcaplus_idls
	  tencentcloud_tcaplus_tables
	  tencentcloud_tcaplus_tablegroups

	Resource
	  tencentcloud_tcaplus_cluster
	  tencentcloud_tcaplus_tablegroup
	  tencentcloud_tcaplus_idl
	  tencentcloud_tcaplus_table

Tencent Container Registry(TCR)

	  Data Source
		tencentcloud_tcr_instances
		tencentcloud_tcr_namespaces
		tencentcloud_tcr_repositories
		tencentcloud_tcr_tokens
		tencentcloud_tcr_vpc_attachments
		tencentcloud_tcr_webhook_trigger_logs
		tencentcloud_tcr_images
		tencentcloud_tcr_image_manifests
		tencentcloud_tcr_tag_retention_execution_tasks

	  Resource
		tencentcloud_tcr_instance
		tencentcloud_tcr_namespace
		tencentcloud_tcr_repository
		tencentcloud_tcr_token
		tencentcloud_tcr_vpc_attachment
		tencentcloud_tcr_tag_retention_rule
		tencentcloud_tcr_webhook_trigger
		tencentcloud_tcr_manage_replication_operation
		tencentcloud_tcr_customized_domain
		tencentcloud_tcr_immutable_tag_rule
		tencentcloud_tcr_delete_image_operation
		tencentcloud_tcr_create_image_signature_operation

Video on Demand(VOD)

	  Data Source
		tencentcloud_vod_adaptive_dynamic_streaming_templates
		tencentcloud_vod_snapshot_by_time_offset_templates
		tencentcloud_vod_super_player_configs
		tencentcloud_vod_image_sprite_templates
		tencentcloud_vod_procedure_templates


	  Resource
	    tencentcloud_vod_adaptive_dynamic_streaming_template
	    tencentcloud_vod_procedure_template
	    tencentcloud_vod_snapshot_by_time_offset_template
	    tencentcloud_vod_image_sprite_template
	    tencentcloud_vod_super_player_config
		tencentcloud_vod_sub_application

Virtual Private Cloud(VPC)

	  Data Source
	    tencentcloud_route_table
	    tencentcloud_security_group
	    tencentcloud_security_groups
		tencentcloud_address_templates
		tencentcloud_address_template_groups
		tencentcloud_protocol_templates
		tencentcloud_protocol_template_groups
	    tencentcloud_subnet
	    tencentcloud_vpc
	    tencentcloud_vpc_acls
	    tencentcloud_vpc_instances
	    tencentcloud_vpc_route_tables
	    tencentcloud_vpc_subnets
	    tencentcloud_dnats
	    tencentcloud_enis
	    tencentcloud_ha_vip_eip_attachments
	    tencentcloud_ha_vips
	    tencentcloud_nat_gateways
	    tencentcloud_nat_gateway_snats
	    tencentcloud_nats
		tencentcloud_nat_dc_route
		tencentcloud_vpc_bandwidth_package_quota
		tencentcloud_vpc_bandwidth_package_bill_usage

	  Resource
	    tencentcloud_eni
	    tencentcloud_eni_attachment
	    tencentcloud_vpc
		tencentcloud_vpc_acl
		tencentcloud_vpc_acl_attachment
		tencentcloud_vpc_traffic_package
	    tencentcloud_subnet
	    tencentcloud_security_group
	    tencentcloud_security_group_rule
	    tencentcloud_security_group_rule_set
	    tencentcloud_security_group_lite_rule
		tencentcloud_address_template
		tencentcloud_address_template_group
		tencentcloud_protocol_template
		tencentcloud_protocol_template_group
	    tencentcloud_route_table
	    tencentcloud_route_entry
	    tencentcloud_route_table_entry
	    tencentcloud_dnat
	    tencentcloud_nat_gateway
	    tencentcloud_nat_gateway_snat
		tencentcloud_nat_refresh_nat_dc_route
	    tencentcloud_ha_vip
	    tencentcloud_ha_vip_eip_attachment
		tencentcloud_vpc_bandwidth_package
		tencentcloud_vpc_bandwidth_package_attachment
		tencentcloud_ipv6_address_bandwidth

Private Link(PLS)

	  Resource
		tencentcloud_vpc_end_point_service
		tencentcloud_vpc_end_point
		tencentcloud_vpc_end_point_service_white_list

Flow Logs(FL)

	 Resource
		tencentcloud_vpc_flow_log

VPN Connections(VPN)

	  Data Source
	    tencentcloud_vpn_connections
	    tencentcloud_vpn_customer_gateways
	    tencentcloud_vpn_gateways
	    tencentcloud_vpn_gateway_routes
		tencentcloud_vpn_customer_gateway_vendors

	  Resource
	    tencentcloud_vpn_customer_gateway
	    tencentcloud_vpn_gateway
	    tencentcloud_vpn_gateway_route
	    tencentcloud_vpn_connection
		tencentcloud_vpn_ssl_server
		tencentcloud_vpn_ssl_client
		tencentcloud_vpn_connection_reset
		tencentcloud_vpn_customer_gateway_configuration_download
		tencentcloud_vpn_gateway_ssl_client_cert
		tencentcloud_vpn_gateway_ccn_routes

MapReduce(EMR)

	Data Source
	  tencentcloud_emr
	  tencentcloud_emr_nodes

	Resource
	  tencentcloud_emr_cluster

DNSPOD

	Resource
	  tencentcloud_dnspod_domain_instance
	  tencentcloud_dnspod_record
	Data Source
	  tencentcloud_dnspod_records

PrivateDNS

	  Resource
	    tencentcloud_private_dns_zone
		tencentcloud_private_dns_record

Cloud Log Service(CLS)

	  Resource
		tencentcloud_cls_logset
		tencentcloud_cls_topic
		tencentcloud_cls_config
		tencentcloud_cls_config_extra
		tencentcloud_cls_config_attachment
		tencentcloud_cls_machine_group
		tencentcloud_cls_cos_shipper
		tencentcloud_cls_index

TencentCloud Lighthouse(Lighthouse)

	  Resource
		tencentcloud_lighthouse_instance
		tencentcloud_lighthouse_blueprint
		tencentcloud_lighthouse_firewall_rule
		tencentcloud_lighthouse_disk_backup
		tencentcloud_lighthouse_apply_disk_backup
		tencentcloud_lighthouse_disk_attachment
		tencentcloud_lighthouse_key_pair
		tencentcloud_lighthouse_snapshot
		tencentcloud_lighthouse_apply_instance_snapshot
		tencentcloud_lighthouse_start_instance
		tencentcloud_lighthouse_stop_instance
		tencentcloud_lighthouse_reboot_instance
		tencentcloud_lighthouse_key_pair_attachment
		tencentcloud_lighthouse_disk

	  Data Source
		tencentcloud_lighthouse_firewall_rules_template
		tencentcloud_lighthouse_bundle
		tencentcloud_lighthouse_zone
		tencentcloud_lighthouse_scene
		tencentcloud_lighthouse_reset_instance_blueprint
		tencentcloud_lighthouse_region
		tencentcloud_lighthouse_instance_vnc_url
		tencentcloud_lighthouse_instance_traffic_package
		tencentcloud_lighthouse_instance_disk_num
		tencentcloud_lighthouse_instance_blueprint
		tencentcloud_lighthouse_disk_config

TencentCloud Elastic Microservice(TEM)

	  Resource
		tencentcloud_tem_environment
		tencentcloud_tem_application
		tencentcloud_tem_workload
		tencentcloud_tem_app_config
		tencentcloud_tem_log_config
		tencentcloud_tem_scale_rule
		tencentcloud_tem_gateway
		tencentcloud_tem_application_service

TencentCloud EdgeOne(TEO)

	  Data Source
		tencentcloud_teo_zone_available_plans
		tencentcloud_teo_bot_managed_rules
		tencentcloud_teo_bot_portrait_rules
		tencentcloud_teo_rule_engine_settings
		tencentcloud_teo_security_policy_regions
		tencentcloud_teo_waf_rule_groups
		tencentcloud_teo_zone_ddos_policy

	  Resource
		tencentcloud_teo_zone
		tencentcloud_teo_zone_setting
		tencentcloud_teo_dns_record
		tencentcloud_teo_dns_sec
		tencentcloud_teo_load_balancing
		tencentcloud_teo_origin_group
		tencentcloud_teo_rule_engine
		tencentcloud_teo_rule_engine_priority
		tencentcloud_teo_application_proxy
		tencentcloud_teo_application_proxy_rule
		tencentcloud_teo_ddos_policy
		tencentcloud_teo_security_policy
		tencentcloud_teo_custom_error_page

TencentCloud ServiceMesh(TCM)

	  Data Source
		tencentcloud_tcm_mesh
	  Resource
		tencentcloud_tcm_mesh
		tencentcloud_tcm_cluster_attachment
		tencentcloud_tcm_prometheus_attachment
		tencentcloud_tcm_tracing_config
		tencentcloud_tcm_access_log_config

Simple Email Service(SES)

	  Resource
		tencentcloud_ses_domain
		tencentcloud_ses_template
		tencentcloud_ses_email_address

Security Token Service(STS)

	  Data Source
		tencentcloud_sts_caller_identity

TDSQL for MySQL(DCDB)

	  Data Source
		tencentcloud_dcdb_instances
		tencentcloud_dcdb_accounts
		tencentcloud_dcdb_databases
		tencentcloud_dcdb_parameters
		tencentcloud_dcdb_shards
		tencentcloud_dcdb_security_groups
		tencentcloud_dcdb_database_objects
		tencentcloud_dcdb_database_tables

	  Resource
		tencentcloud_dcdb_account
		tencentcloud_dcdb_hourdb_instance
		tencentcloud_dcdb_security_group_attachment
		tencentcloud_dcdb_account_privileges
		tencentcloud_dcdb_db_parameters

Short Message Service(SMS)

	  Resource
		tencentcloud_sms_sign
		tencentcloud_sms_template

Cloud Automated Testing(CAT)

	  Data Source
		tencentcloud_cat_probe_data
		tencentcloud_cat_node

	  Resource
	 	tencentcloud_cat_task_set

TencentDB for MariaDB(MariaDB)

	  Data Source
	    tencentcloud_mariadb_db_instances
		tencentcloud_mariadb_accounts
		tencentcloud_mariadb_security_groups
		tencentcloud_mariadb_database_objects
		tencentcloud_mariadb_databases
		tencentcloud_mariadb_database_table
	  Resource
		tencentcloud_mariadb_dedicatedcluster_db_instance
		tencentcloud_mariadb_instance
		tencentcloud_mariadb_hour_db_instance
		tencentcloud_mariadb_account
		tencentcloud_mariadb_parameters
		tencentcloud_mariadb_log_file_retention_period
		tencentcloud_mariadb_security_groups

Real User Monitoring(RUM)

	  Data Source
		tencentcloud_rum_project
		tencentcloud_rum_offline_log_config
		tencentcloud_rum_whitelist
		tencentcloud_rum_taw_instance
	  Resource
		tencentcloud_rum_project
		tencentcloud_rum_taw_instance
		tencentcloud_rum_whitelist
		tencentcloud_rum_offline_log_config_attachment

Cloud Streaming Services(CSS)

	  Resource
	    tencentcloud_css_watermark
		tencentcloud_css_pull_stream_task
		tencentcloud_css_live_transcode_template
		tencentcloud_css_live_transcode_rule_attachment
		tencentcloud_css_domain
		tencentcloud_css_authenticate_domain_owner_operation
		tencentcloud_css_play_domain_cert_attachment
		tencentcloud_css_play_auth_key_config
		tencentcloud_css_push_auth_key_config
	  Data Source
		tencentcloud_css_domains

Performance Testing Service(PTS)

	  Resource
		tencentcloud_pts_project
		tencentcloud_pts_alert_channel
		tencentcloud_pts_scenario
		tencentcloud_pts_file
		tencentcloud_pts_job
		tencentcloud_pts_cron_job

TencentCloud Automation Tools(TAT)

	  Data Source
		tencentcloud_tat_command
		tencentcloud_tat_invoker
		tencentcloud_tat_invoker_records
		tencentcloud_tat_agent
		tencentcloud_tat_invocation_task
	  Resource
		tencentcloud_tat_command
		tencentcloud_tat_invoker
		tencentcloud_tat_invoker_config
		tencentcloud_tat_invocation_invoke_attachment
		tencentcloud_tat_invocation_command_attachment

Tencent Cloud Organization (TCO)

	  Resource
		tencentcloud_organization_org_node
		tencentcloud_organization_org_member
		tencentcloud_organization_policy_sub_account_attachment

TDSQL-C for PostgreSQL(TDCPG)

	  Data Source
		tencentcloud_tdcpg_clusters
		tencentcloud_tdcpg_instances
	  Resource
		tencentcloud_tdcpg_cluster
		tencentcloud_tdcpg_instance

TencentDB for DBbrain(dbbrain)

	  Data Source
		tencentcloud_dbbrain_sql_filters
		tencentcloud_dbbrain_security_audit_log_export_tasks
		tencentcloud_dbbrain_diag_event
		tencentcloud_dbbrain_diag_events
		tencentcloud_dbbrain_diag_history
		tencentcloud_dbbrain_security_audit_log_download_urls
		tencentcloud_dbbrain_slow_log_time_series_stats
		tencentcloud_dbbrain_slow_log_top_sqls
		tencentcloud_dbbrain_slow_log_user_host_stats
		tencentcloud_dbbrain_slow_log_user_sql_advice
		tencentcloud_dbbrain_slow_logs
		tencentcloud_dbbrain_health_scores
		tencentcloud_dbbrain_sql_templates
		tencentcloud_dbbrain_db_space_status
		tencentcloud_dbbrain_top_space_schemas
		tencentcloud_dbbrain_top_space_tables
		tencentcloud_dbbrain_top_space_schema_time_series
		tencentcloud_dbbrain_top_space_table_time_series

	  Resource
		tencentcloud_dbbrain_sql_filter
		tencentcloud_dbbrain_security_audit_log_export_task
		tencentcloud_dbbrain_db_diag_report_task
		tencentcloud_dbbrain_modify_diag_db_instance_operation
		tencentcloud_dbbrain_tdsql_audit_log

Data Transmission Service(DTS)

	  Data Source
		tencentcloud_dts_sync_jobs
		tencentcloud_dts_migrate_jobs
		tencentcloud_dts_compare_tasks

	  Resource
		tencentcloud_dts_sync_job
		tencentcloud_dts_compare_task
		tencentcloud_dts_migrate_service
		tencentcloud_dts_migrate_job
		tencentcloud_dts_migrate_job_start_operation
		tencentcloud_dts_migrate_job_resume_operation
		tencentcloud_dts_compare_task_stop_operation
		tencentcloud_dts_migrate_job_config

TDMQ for RocketMQ(trocket)

	  Data Source
		tencentcloud_tdmq_rocketmq_cluster
		tencentcloud_tdmq_rocketmq_namespace
		tencentcloud_tdmq_rocketmq_topic
		tencentcloud_tdmq_rocketmq_role
		tencentcloud_tdmq_rocketmq_group

	  Resource
		tencentcloud_tdmq_rocketmq_cluster
		tencentcloud_tdmq_rocketmq_namespace
		tencentcloud_tdmq_rocketmq_role
		tencentcloud_tdmq_rocketmq_topic
		tencentcloud_tdmq_rocketmq_group
		tencentcloud_tdmq_rocketmq_environment_role

Cloud Infinite(CI)

	  Resource
		tencentcloud_ci_bucket_attachment
		tencentcloud_ci_bucket_pic_style
		tencentcloud_ci_hot_link
		tencentcloud_ci_media_snapshot_template
		tencentcloud_ci_media_transcode_template
		tencentcloud_ci_media_animation_template
		tencentcloud_ci_media_concat_template
		tencentcloud_ci_media_video_process_template
		tencentcloud_ci_media_video_montage_template
		tencentcloud_ci_media_voice_separate_template
		tencentcloud_ci_media_super_resolution_template
		tencentcloud_ci_media_pic_process_template
		tencentcloud_ci_media_watermark_template
		tencentcloud_ci_media_tts_template
		tencentcloud_ci_media_transcode_pro_template
		tencentcloud_ci_media_smart_cover_template
		tencentcloud_ci_media_speech_recognition_template
		tencentcloud_ci_guetzli
		tencentcloud_ci_original_image_protection

TDMQ for CMQ

	  Data Source
	    tencentcloud_tcmq_queue
		tencentcloud_tcmq_topic
		tencentcloud_tcmq_subscribe

	  Resource
	    tencentcloud_tcmq_queue
		tencentcloud_tcmq_topic
		tencentcloud_tcmq_subscribe

Tencent Service Framework(TSF)

	  Data Source
		tencentcloud_tsf_application
		tencentcloud_tsf_application_config
		tencentcloud_tsf_application_file_config
		tencentcloud_tsf_application_public_config
		tencentcloud_tsf_cluster
		tencentcloud_tsf_microservice
		tencentcloud_tsf_unit_rules
		tencentcloud_tsf_config_summary
		tencentcloud_tsf_delivery_config_by_group_id
		tencentcloud_tsf_delivery_configs
		tencentcloud_tsf_public_config_summary
		tencentcloud_tsf_api_group
		tencentcloud_tsf_application_attribute
		tencentcloud_tsf_business_log_configs
		tencentcloud_tsf_api_detail
		tencentcloud_tsf_microservice_api_version

	  Resource
	  	tencentcloud_tsf_cluster
		tencentcloud_tsf_microservice
		tencentcloud_tsf_application_config
		tencentcloud_tsf_api_group
		tencentcloud_tsf_namespace
		tencentcloud_tsf_path_rewrite
		tencentcloud_tsf_unit_rule
		tencentcloud_tsf_task
		tencentcloud_tsf_config_template
		tencentcloud_tsf_api_rate_limit_rule
		tencentcloud_tsf_application_release_config
		tencentcloud_tsf_lane
		tencentcloud_tsf_lane_rule
		tencentcloud_tsf_group
		tencentcloud_tsf_application
		tencentcloud_tsf_application_public_config_release
		tencentcloud_tsf_application_public_config
		tencentcloud_tsf_application_file_config_release
		tencentcloud_tsf_instances_attachment
		tencentcloud_tsf_bind_api_group

Media Processing Service(MPS)

	  Resource
		tencentcloud_mps_workflow
		tencentcloud_mps_transcode_template
		tencentcloud_mps_watermark_template
		tencentcloud_mps_image_sprite_template
		tencentcloud_mps_snapshot_by_timeoffset_template
		tencentcloud_mps_sample_snapshot_template
		tencentcloud_mps_animated_graphics_template
		tencentcloud_mps_ai_recognition_template
		tencentcloud_mps_ai_analysis_template
		tencentcloud_mps_adaptive_dynamic_streaming_template
		tencentcloud_mps_person_sample

Cloud HDFS(CHDFS)

	  Data Source
		tencentcloud_chdfs_access_groups
		tencentcloud_chdfs_mount_points

	  Resource
		tencentcloud_chdfs_access_group
		tencentcloud_chdfs_access_rule
		tencentcloud_chdfs_file_system
		tencentcloud_chdfs_life_cycle_rule
		tencentcloud_chdfs_mount_point
		tencentcloud_chdfs_mount_point_attachment

StreamLive(MDL)

	  Resource
		tencentcloud_mdl_stream_live_input

Application Performance Management(APM)

	  Resource
		tencentcloud_apm_instance

Tencent Cloud Service Engine(TSE)

	  Data Source
		tencentcloud_tse_access_address
		tencentcloud_tse_nacos_replicas
		tencentcloud_tse_zookeeper_replicas
		tencentcloud_tse_zookeeper_server_interfaces
		tencentcloud_tse_nacos_server_interfaces

	  Resource
		tencentcloud_tse_instance
*/
package tencentcloud

import (
	"net/url"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/connectivity"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

const (
	PROVIDER_SECRET_ID                    = "TENCENTCLOUD_SECRET_ID"
	PROVIDER_SECRET_KEY                   = "TENCENTCLOUD_SECRET_KEY"
	PROVIDER_SECURITY_TOKEN               = "TENCENTCLOUD_SECURITY_TOKEN"
	PROVIDER_REGION                       = "TENCENTCLOUD_REGION"
	PROVIDER_PROTOCOL                     = "TENCENTCLOUD_PROTOCOL"
	PROVIDER_DOMAIN                       = "TENCENTCLOUD_DOMAIN"
	PROVIDER_ASSUME_ROLE_ARN              = "TENCENTCLOUD_ASSUME_ROLE_ARN"
	PROVIDER_ASSUME_ROLE_SESSION_NAME     = "TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME"
	PROVIDER_ASSUME_ROLE_SESSION_DURATION = "TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION"
)

type TencentCloudClient struct {
	apiV3Conn *connectivity.TencentCloudClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "This is the TencentCloud access key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.",
				Sensitive:   true,
			},
			"security_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECURITY_TOKEN, nil),
				Description: "TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUD_SECURITY_TOKEN` environment variable. Notice: for supported products, please refer to: [temporary key supported products](https://intl.cloud.tencent.com/document/product/598/10588).",
				Sensitive:   true,
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_REGION, nil),
				Description:  "This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The default input value is ap-guangzhou.",
				InputDefault: "ap-guangzhou",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_PROTOCOL, "HTTPS"),
				ValidateFunc: validateAllowedStringValue([]string{"HTTP", "HTTPS"}),
				Description:  "The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTPS`.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_DOMAIN, nil),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"assume_role": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_ARN, nil),
							Description: "The ARN of the role to assume. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN`.",
						},
						"session_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_SESSION_NAME, nil),
							Description: "The session name to use when making the AssumeRole call. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`.",
						},
						"session_duration": {
							Type:     schema.TypeInt,
							Required: true,
							DefaultFunc: func() (interface{}, error) {
								if v := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); v != "" {
									return strconv.Atoi(v)
								}
								return 7200, nil
							},
							ValidateFunc: validateIntegerInRange(0, 43200),
							Description:  "The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION`.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A more restrictive policy when making the AssumeRole call. Its content must not contains `principal` elements. Notice: more syntax references, please refer to: [policies syntax logic](https://intl.cloud.tencent.com/document/product/598/10603).",
						},
					},
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tencentcloud_availability_regions":                      dataSourceTencentCloudAvailabilityRegions(),
			"tencentcloud_emr":                                       dataSourceTencentCloudEmr(),
			"tencentcloud_emr_nodes":                                 dataSourceTencentCloudEmrNodes(),
			"tencentcloud_availability_zones":                        dataSourceTencentCloudAvailabilityZones(),
			"tencentcloud_availability_zones_by_product":             dataSourceTencentCloudAvailabilityZonesByProduct(),
			"tencentcloud_instances":                                 dataSourceTencentCloudInstances(),
			"tencentcloud_instances_set":                             dataSourceTencentCloudInstancesSet(),
			"tencentcloud_reserved_instances":                        dataSourceTencentCloudReservedInstances(),
			"tencentcloud_placement_groups":                          dataSourceTencentCloudPlacementGroups(),
			"tencentcloud_key_pairs":                                 dataSourceTencentCloudKeyPairs(),
			"tencentcloud_image":                                     dataSourceTencentCloudImage(),
			"tencentcloud_images":                                    dataSourceTencentCloudImages(),
			"tencentcloud_instance_types":                            dataSourceInstanceTypes(),
			"tencentcloud_reserved_instance_configs":                 dataSourceTencentCloudReservedInstanceConfigs(),
			"tencentcloud_vpc_instances":                             dataSourceTencentCloudVpcInstances(),
			"tencentcloud_vpc_subnets":                               dataSourceTencentCloudVpcSubnets(),
			"tencentcloud_vpc_route_tables":                          dataSourceTencentCloudVpcRouteTables(),
			"tencentcloud_vpc":                                       dataSourceTencentCloudVpc(),
			"tencentcloud_vpc_acls":                                  dataSourceTencentCloudVpcAcls(),
			"tencentcloud_vpc_bandwidth_package_quota":               dataSourceTencentCloudVpcBandwidthPackageQuota(),
			"tencentcloud_vpc_bandwidth_package_bill_usage":          dataSourceTencentCloudVpcBandwidthPackageBillUsage(),
			"tencentcloud_subnet":                                    dataSourceTencentCloudSubnet(),
			"tencentcloud_route_table":                               dataSourceTencentCloudRouteTable(),
			"tencentcloud_domains":                                   dataSourceTencentCloudDomains(),
			"tencentcloud_eip":                                       dataSourceTencentCloudEip(),
			"tencentcloud_eips":                                      dataSourceTencentCloudEips(),
			"tencentcloud_eip_address_quota":                         dataSourceTencentCloudEipAddressQuota(),
			"tencentcloud_eip_network_account_type":                  dataSourceTencentCloudEipNetworkAccountType(),
			"tencentcloud_enis":                                      dataSourceTencentCloudEnis(),
			"tencentcloud_nats":                                      dataSourceTencentCloudNats(),
			"tencentcloud_dnats":                                     dataSourceTencentCloudDnats(),
			"tencentcloud_nat_gateways":                              dataSourceTencentCloudNatGateways(),
			"tencentcloud_nat_gateway_snats":                         dataSourceTencentCloudNatGatewaySnats(),
			"tencentcloud_nat_dc_route":                              dataSourceTencentCloudNatDcRoute(),
			"tencentcloud_vpn_customer_gateways":                     dataSourceTencentCloudVpnCustomerGateways(),
			"tencentcloud_vpn_gateways":                              dataSourceTencentCloudVpnGateways(),
			"tencentcloud_vpn_gateway_routes":                        dataSourceTencentCloudVpnGatewayRoutes(),
			"tencentcloud_vpn_connections":                           dataSourceTencentCloudVpnConnections(),
			"tencentcloud_vpn_customer_gateway_vendors":              dataSourceTencentCloudVpnCustomerGatewayVendors(),
			"tencentcloud_ha_vips":                                   dataSourceTencentCloudHaVips(),
			"tencentcloud_ha_vip_eip_attachments":                    dataSourceTencentCloudHaVipEipAttachments(),
			"tencentcloud_ccn_instances":                             dataSourceTencentCloudCcnInstances(),
			"tencentcloud_ccn_bandwidth_limits":                      dataSourceTencentCloudCcnBandwidthLimits(),
			"tencentcloud_ccn_cross_border_compliance":               dataSourceTencentCloudCcnCrossBorderCompliance(),
			"tencentcloud_ccn_tenant_instances":                      dataSourceTencentCloudCcnTenantInstance(),
			"tencentcloud_ccn_cross_border_flow_monitor":             dataSourceTencentCloudCcnCrossBorderFlowMonitor(),
			"tencentcloud_ccn_cross_border_region_bandwidth_limits":  dataSourceTencentCloudCcnCrossBorderRegionBandwidthLimits(),
			"tencentcloud_dc_instances":                              dataSourceTencentCloudDcInstances(),
			"tencentcloud_dcx_instances":                             dataSourceTencentCloudDcxInstances(),
			"tencentcloud_dc_gateway_instances":                      dataSourceTencentCloudDcGatewayInstances(),
			"tencentcloud_dc_gateway_ccn_routes":                     dataSourceTencentCloudDcGatewayCCNRoutes(),
			"tencentcloud_security_group":                            dataSourceTencentCloudSecurityGroup(),
			"tencentcloud_security_groups":                           dataSourceTencentCloudSecurityGroups(),
			"tencentcloud_kubernetes_clusters":                       dataSourceTencentCloudKubernetesClusters(),
			"tencentcloud_kubernetes_charts":                         dataSourceTencentCloudKubernetesCharts(),
			"tencentcloud_kubernetes_cluster_levels":                 datasourceTencentCloudKubernetesClusterLevels(),
			"tencentcloud_kubernetes_cluster_common_names":           datasourceTencentCloudKubernetesClusterCommonNames(),
			"tencentcloud_kubernetes_cluster_authentication_options": dataSourceTencentCloudKubernetesClusterAuthenticationOptions(),
			"tencentcloud_kubernetes_available_cluster_versions":     dataSourceTencentCloudKubernetesAvailableClusterVersions(),
			"tencentcloud_eks_clusters":                              dataSourceTencentCloudEKSClusters(),
			"tencentcloud_eks_cluster_credential":                    datasourceTencentCloudEksClusterCredential(),
			"tencentcloud_container_clusters":                        dataSourceTencentCloudContainerClusters(),
			"tencentcloud_container_cluster_instances":               dataSourceTencentCloudContainerClusterInstances(),
			"tencentcloud_mysql_backup_list":                         dataSourceTencentMysqlBackupList(),
			"tencentcloud_mysql_zone_config":                         dataSourceTencentMysqlZoneConfig(),
			"tencentcloud_mysql_parameter_list":                      dataSourceTencentCloudMysqlParameterList(),
			"tencentcloud_mysql_default_params":                      datasourceTencentCloudMysqlDefaultParams(),
			"tencentcloud_mysql_instance":                            dataSourceTencentCloudMysqlInstance(),
			"tencentcloud_mysql_backup_overview":                     dataSourceTencentCloudMysqlBackupOverview(),
			"tencentcloud_mysql_backup_summaries":                    dataSourceTencentCloudMysqlBackupSummaries(),
			"tencentcloud_mysql_bin_log":                             dataSourceTencentCloudMysqlBinLog(),
			"tencentcloud_mysql_binlog_backup_overview":              dataSourceTencentCloudMysqlBinlogBackupOverview(),
			"tencentcloud_mysql_clone_list":                          dataSourceTencentCloudMysqlCloneList(),
			"tencentcloud_mysql_data_backup_overview":                dataSourceTencentCloudMysqlDataBackupOverview(),
			"tencentcloud_mysql_db_features":                         dataSourceTencentCloudMysqlDbFeatures(),
			"tencentcloud_mysql_inst_tables":                         dataSourceTencentCloudMysqlInstTables(),
			"tencentcloud_mysql_instance_charset":                    dataSourceTencentCloudMysqlInstanceCharset(),
			"tencentcloud_mysql_instance_info":                       dataSourceTencentCloudMysqlInstanceInfo(),
			"tencentcloud_mysql_instance_param_record":               dataSourceTencentCloudMysqlInstanceParamRecord(),
			"tencentcloud_mysql_instance_reboot_time":                dataSourceTencentCloudMysqlInstanceRebootTime(),
			"tencentcloud_mysql_proxy_custom":                        dataSourceTencentCloudMysqlProxyCustom(),
			"tencentcloud_mysql_rollback_range_time":                 dataSourceTencentCloudMysqlRollbackRangeTime(),
			"tencentcloud_mysql_slow_log":                            dataSourceTencentCloudMysqlSlowLog(),
			"tencentcloud_mysql_slow_log_data":                       dataSourceTencentCloudMysqlSlowLogData(),
			"tencentcloud_mysql_supported_privileges":                dataSourceTencentCloudMysqlSupportedPrivileges(),
			"tencentcloud_mysql_switch_record":                       dataSourceTencentCloudMysqlSwitchRecord(),
			"tencentcloud_mysql_user_task":                           dataSourceTencentCloudMysqlUserTask(),
			"tencentcloud_cos_bucket_object":                         dataSourceTencentCloudCosBucketObject(),
			"tencentcloud_cos_buckets":                               dataSourceTencentCloudCosBuckets(),
			"tencentcloud_cfs_file_systems":                          dataSourceTencentCloudCfsFileSystems(),
			"tencentcloud_cfs_access_groups":                         dataSourceTencentCloudCfsAccessGroups(),
			"tencentcloud_cfs_access_rules":                          dataSourceTencentCloudCfsAccessRules(),
			"tencentcloud_cfs_mount_targets":                         dataSourceTencentCloudCfsMountTargets(),
			"tencentcloud_cfs_file_system_clients":                   dataSourceTencentCloudCfsFileSystemClients(),
			"tencentcloud_cfs_available_zone":                        dataSourceTencentCloudCfsAvailableZone(),
			"tencentcloud_redis_zone_config":                         dataSourceTencentRedisZoneConfig(),
			"tencentcloud_redis_instances":                           dataSourceTencentRedisInstances(),
			"tencentcloud_redis_backup":                              dataSourceTencentCloudRedisBackup(),
			"tencentcloud_redis_backup_download_info":                dataSourceTencentCloudRedisBackupDownloadInfo(),
			"tencentcloud_redis_param_records":                       dataSourceTencentCloudRedisRecordsParam(),
			"tencentcloud_redis_instance_shards":                     dataSourceTencentCloudRedisInstanceShards(),
			"tencentcloud_redis_instance_zone_info":                  dataSourceTencentCloudRedisInstanceZoneInfo(),
			"tencentcloud_redis_instance_task_list":                  dataSourceTencentCloudRedisInstanceTaskList(),
			"tencentcloud_redis_instance_node_info":                  dataSourceTencentCloudRedisInstanceNodeInfo(),
			"tencentcloud_as_scaling_configs":                        dataSourceTencentCloudAsScalingConfigs(),
			"tencentcloud_as_scaling_groups":                         dataSourceTencentCloudAsScalingGroups(),
			"tencentcloud_as_scaling_policies":                       dataSourceTencentCloudAsScalingPolicies(),
			"tencentcloud_cbs_storages":                              dataSourceTencentCloudCbsStorages(),
			"tencentcloud_cbs_storages_set":                          dataSourceTencentCloudCbsStoragesSet(),
			"tencentcloud_cbs_snapshots":                             dataSourceTencentCloudCbsSnapshots(),
			"tencentcloud_cbs_snapshot_policies":                     dataSourceTencentCloudCbsSnapshotPolicies(),
			"tencentcloud_clb_instances":                             dataSourceTencentCloudClbInstances(),
			"tencentcloud_clb_listeners":                             dataSourceTencentCloudClbListeners(),
			"tencentcloud_clb_listener_rules":                        dataSourceTencentCloudClbListenerRules(),
			"tencentcloud_clb_attachments":                           dataSourceTencentCloudClbServerAttachments(),
			"tencentcloud_clb_redirections":                          dataSourceTencentCloudClbRedirections(),
			"tencentcloud_clb_target_groups":                         dataSourceTencentCloudClbTargetGroups(),
			"tencentcloud_mongodb_zone_config":                       dataSourceTencentCloudMongodbZoneConfig(),
			"tencentcloud_mongodb_instances":                         dataSourceTencentCloudMongodbInstances(),
			"tencentcloud_mongodb_instance_backups":                  dataSourceTencentCloudMongodbInstanceBackups(),
			"tencentcloud_mongodb_instance_connections":              dataSourceTencentCloudMongodbInstanceConnections(),
			"tencentcloud_mongodb_instance_current_op":               dataSourceTencentCloudMongodbInstanceCurrentOp(),
			"tencentcloud_mongodb_instance_params":                   dataSourceTencentCloudMongodbInstanceParams(),
			"tencentcloud_mongodb_instance_slow_log":                 dataSourceTencentCloudMongodbInstanceSlowLog(),
			"tencentcloud_dayu_cc_https_policies":                    dataSourceTencentCloudDayuCCHttpsPolicies(),
			"tencentcloud_dayu_cc_http_policies":                     dataSourceTencentCloudDayuCCHttpPolicies(),
			"tencentcloud_dayu_ddos_policies":                        dataSourceTencentCloudDayuDdosPolicies(),
			"tencentcloud_dayu_ddos_policy_cases":                    dataSourceTencentCloudDayuDdosPolicyCases(),
			"tencentcloud_dayu_ddos_policy_attachments":              dataSourceTencentCloudDayuDdosPolicyAttachments(),
			"tencentcloud_dayu_l4_rules":                             dataSourceTencentCloudDayuL4Rules(),
			"tencentcloud_dayu_l4_rules_v2":                          dataSourceTencentCloudDayuL4RulesV2(),
			"tencentcloud_dayu_l7_rules":                             dataSourceTencentCloudDayuL7Rules(),
			"tencentcloud_dayu_l7_rules_v2":                          dataSourceTencentCloudDayuL7RulesV2(),
			"tencentcloud_gaap_proxies":                              dataSourceTencentCloudGaapProxies(),
			"tencentcloud_gaap_realservers":                          dataSourceTencentCloudGaapRealservers(),
			"tencentcloud_gaap_layer4_listeners":                     dataSourceTencentCloudGaapLayer4Listeners(),
			"tencentcloud_gaap_layer7_listeners":                     dataSourceTencentCloudGaapLayer7Listeners(),
			"tencentcloud_gaap_http_domains":                         dataSourceTencentCloudGaapHttpDomains(),
			"tencentcloud_gaap_http_rules":                           dataSourceTencentCloudGaapHttpRules(),
			"tencentcloud_gaap_security_policies":                    dataSourceTencentCloudGaapSecurityPolices(),
			"tencentcloud_gaap_security_rules":                       dataSourceTencentCloudGaapSecurityRules(),
			"tencentcloud_gaap_certificates":                         dataSourceTencentCloudGaapCertificates(),
			"tencentcloud_gaap_domain_error_pages":                   dataSourceTencentCloudGaapDomainErrorPageInfoList(),
			"tencentcloud_ssl_certificates":                          dataSourceTencentCloudSslCertificates(),
			"tencentcloud_cam_roles":                                 dataSourceTencentCloudCamRoles(),
			"tencentcloud_cam_users":                                 dataSourceTencentCloudCamUsers(),
			"tencentcloud_cam_groups":                                dataSourceTencentCloudCamGroups(),
			"tencentcloud_cam_group_memberships":                     dataSourceTencentCloudCamGroupMemberships(),
			"tencentcloud_cam_policies":                              dataSourceTencentCloudCamPolicies(),
			"tencentcloud_cam_role_policy_attachments":               dataSourceTencentCloudCamRolePolicyAttachments(),
			"tencentcloud_cam_user_policy_attachments":               dataSourceTencentCloudCamUserPolicyAttachments(),
			"tencentcloud_cam_group_policy_attachments":              dataSourceTencentCloudCamGroupPolicyAttachments(),
			"tencentcloud_cam_saml_providers":                        dataSourceTencentCloudCamSAMLProviders(),
			"tencentcloud_user_info":                                 datasourceTencentCloudUserInfo(),
			"tencentcloud_cdn_domains":                               dataSourceTencentCloudCdnDomains(),
			"tencentcloud_cdn_domain_verifier":                       dataSourceTencentCloudCdnDomainVerifyRecord(),
			"tencentcloud_scf_functions":                             dataSourceTencentCloudScfFunctions(),
			"tencentcloud_scf_namespaces":                            dataSourceTencentCloudScfNamespaces(),
			"tencentcloud_scf_logs":                                  dataSourceTencentCloudScfLogs(),
			"tencentcloud_tcaplus_clusters":                          dataSourceTencentCloudTcaplusClusters(),
			"tencentcloud_tcaplus_tablegroups":                       dataSourceTencentCloudTcaplusTableGroups(),
			"tencentcloud_tcaplus_tables":                            dataSourceTencentCloudTcaplusTables(),
			"tencentcloud_tcaplus_idls":                              dataSourceTencentCloudTcaplusIdls(),
			"tencentcloud_monitor_policy_conditions":                 dataSourceTencentMonitorPolicyConditions(),
			"tencentcloud_monitor_data":                              dataSourceTencentMonitorData(),
			"tencentcloud_monitor_product_event":                     dataSourceTencentMonitorProductEvent(),
			"tencentcloud_monitor_binding_objects":                   dataSourceTencentMonitorBindingObjects(),
			"tencentcloud_monitor_policy_groups":                     dataSourceTencentMonitorPolicyGroups(),
			"tencentcloud_monitor_product_namespace":                 dataSourceTencentMonitorProductNamespace(),
			"tencentcloud_monitor_alarm_notices":                     dataSourceTencentMonitorAlarmNotices(),
			"tencentcloud_elasticsearch_instances":                   dataSourceTencentCloudElasticsearchInstances(),
			"tencentcloud_postgresql_instances":                      dataSourceTencentCloudPostgresqlInstances(),
			"tencentcloud_postgresql_specinfos":                      dataSourceTencentCloudPostgresqlSpecinfos(),
			"tencentcloud_postgresql_xlogs":                          datasourceTencentCloudPostgresqlXlogs(),
			"tencentcloud_postgresql_parameter_templates":            dataSourceTencentCloudPostgresqlParameterTemplates(),
			"tencentcloud_sqlserver_zone_config":                     dataSourceTencentSqlserverZoneConfig(),
			"tencentcloud_sqlserver_instances":                       dataSourceTencentCloudSqlserverInstances(),
			"tencentcloud_sqlserver_backups":                         dataSourceTencentCloudSqlserverBackups(),
			"tencentcloud_sqlserver_dbs":                             dataSourceTencentSqlserverDBs(),
			"tencentcloud_sqlserver_accounts":                        dataSourceTencentCloudSqlserverAccounts(),
			"tencentcloud_sqlserver_account_db_attachments":          dataSourceTencentCloudSqlserverAccountDBAttachments(),
			"tencentcloud_sqlserver_readonly_groups":                 dataSourceTencentCloudSqlserverReadonlyGroups(),
			"tencentcloud_sqlserver_backup_commands":                 dataSourceTencentCloudSqlserverBackupCommands(),
			"tencentcloud_sqlserver_backup_by_flow_id":               dataSourceTencentCloudSqlserverBackupByFlowId(),
			"tencentcloud_sqlserver_backup_upload_size":              dataSourceTencentCloudSqlserverBackupUploadSize(),
			"tencentcloud_sqlserver_cross_region_zone":               dataSourceTencentCloudSqlserverCrossRegionZone(),
			"tencentcloud_sqlserver_db_charsets":                     dataSourceTencentCloudSqlserverDBCharsets(),
			"tencentcloud_ckafka_users":                              dataSourceTencentCloudCkafkaUsers(),
			"tencentcloud_ckafka_acls":                               dataSourceTencentCloudCkafkaAcls(),
			"tencentcloud_ckafka_topics":                             dataSourceTencentCloudCkafkaTopics(),
			"tencentcloud_ckafka_instances":                          dataSourceTencentCloudCkafkaInstances(),
			"tencentcloud_audit_cos_regions":                         dataSourceTencentCloudAuditCosRegions(),
			"tencentcloud_audit_key_alias":                           dataSourceTencentCloudAuditKeyAlias(),
			"tencentcloud_audits":                                    dataSourceTencentCloudAudits(),
			"tencentcloud_cynosdb_clusters":                          dataSourceTencentCloudCynosdbClusters(),
			"tencentcloud_cynosdb_instances":                         dataSourceTencentCloudCynosdbInstances(),
			"tencentcloud_cynosdb_zone_config":                       dataSourceTencentCynosdbZoneConfig(),
			"tencentcloud_vod_adaptive_dynamic_streaming_templates":  dataSourceTencentCloudVodAdaptiveDynamicStreamingTemplates(),
			"tencentcloud_vod_image_sprite_templates":                dataSourceTencentCloudVodImageSpriteTemplates(),
			"tencentcloud_vod_procedure_templates":                   dataSourceTencentCloudVodProcedureTemplates(),
			"tencentcloud_vod_snapshot_by_time_offset_templates":     dataSourceTencentCloudVodSnapshotByTimeOffsetTemplates(),
			"tencentcloud_vod_super_player_configs":                  dataSourceTencentCloudVodSuperPlayerConfigs(),
			"tencentcloud_sqlserver_publish_subscribes":              dataSourceTencentSqlserverPublishSubscribes(),
			"tencentcloud_sqlserver_instance_param_records":          dataSourceTencentCloudSqlserverInstanceParamRecords(),
			"tencentcloud_sqlserver_project_security_groups":         dataSourceTencentCloudSqlserverProjectSecurityGroups(),
			"tencentcloud_sqlserver_regions":                         dataSourceTencentCloudSqlserverRegions(),
			"tencentcloud_sqlserver_rollback_time":                   dataSourceTencentCloudSqlserverRollbackTime(),
			"tencentcloud_sqlserver_slowlogs":                        dataSourceTencentCloudSqlserverSlowlogs(),
			"tencentcloud_sqlserver_upload_backup_info":              dataSourceTencentCloudSqlserverUploadBackupInfo(),
			"tencentcloud_sqlserver_upload_incremental_info":         dataSourceTencentCloudSqlserverUploadIncrementalInfo(),
			"tencentcloud_api_gateway_usage_plans":                   dataSourceTencentCloudAPIGatewayUsagePlans(),
			"tencentcloud_api_gateway_ip_strategies":                 dataSourceTencentCloudAPIGatewayIpStrategy(),
			"tencentcloud_api_gateway_customer_domains":              dataSourceTencentCloudAPIGatewayCustomerDomains(),
			"tencentcloud_api_gateway_usage_plan_environments":       dataSourceTencentCloudAPIGatewayUsagePlanEnvironments(),
			"tencentcloud_api_gateway_throttling_services":           dataSourceTencentCloudAPIGatewayThrottlingServices(),
			"tencentcloud_api_gateway_throttling_apis":               dataSourceTencentCloudAPIGatewayThrottlingApis(),
			"tencentcloud_api_gateway_apis":                          dataSourceTencentCloudAPIGatewayAPIs(),
			"tencentcloud_api_gateway_services":                      dataSourceTencentCloudAPIGatewayServices(),
			"tencentcloud_api_gateway_api_keys":                      dataSourceTencentCloudAPIGatewayAPIKeys(),
			"tencentcloud_sqlserver_basic_instances":                 dataSourceTencentCloudSqlserverBasicInstances(),
			"tencentcloud_tcr_instances":                             dataSourceTencentCloudTCRInstances(),
			"tencentcloud_tcr_namespaces":                            dataSourceTencentCloudTCRNamespaces(),
			"tencentcloud_tcr_tokens":                                dataSourceTencentCloudTCRTokens(),
			"tencentcloud_tcr_vpc_attachments":                       dataSourceTencentCloudTCRVPCAttachments(),
			"tencentcloud_tcr_repositories":                          dataSourceTencentCloudTCRRepositories(),
			"tencentcloud_tcr_webhook_trigger_logs":                  dataSourceTencentCloudTcrWebhookTriggerLogs(),
			"tencentcloud_tcr_images":                                dataSourceTencentCloudTcrImages(),
			"tencentcloud_tcr_image_manifests":                       dataSourceTencentCloudTcrImageManifests(),
			"tencentcloud_tcr_tag_retention_execution_tasks":         dataSourceTencentCloudTcrTagRetentionExecutionTasks(),
			"tencentcloud_address_templates":                         dataSourceTencentCloudAddressTemplates(),
			"tencentcloud_address_template_groups":                   dataSourceTencentCloudAddressTemplateGroups(),
			"tencentcloud_protocol_templates":                        dataSourceTencentCloudProtocolTemplates(),
			"tencentcloud_protocol_template_groups":                  dataSourceTencentCloudProtocolTemplateGroups(),
			"tencentcloud_kms_keys":                                  dataSourceTencentCloudKmsKeys(),
			"tencentcloud_ssm_secrets":                               dataSourceTencentCloudSsmSecrets(),
			"tencentcloud_ssm_secret_versions":                       dataSourceTencentCloudSsmSecretVersions(),
			"tencentcloud_cdh_instances":                             dataSourceTencentCloudCdhInstances(),
			"tencentcloud_dayu_eip":                                  dataSourceTencentCloudDayuEip(),
			"tencentcloud_teo_zone_available_plans":                  dataSourceTencentCloudTeoZoneAvailablePlans(),
			"tencentcloud_teo_bot_managed_rules":                     dataSourceTencentCloudTeoBotManagedRules(),
			"tencentcloud_teo_bot_portrait_rules":                    dataSourceTencentCloudTeoBotPortraitRules(),
			"tencentcloud_teo_rule_engine_settings":                  dataSourceTencentCloudTeoRuleEngineSettings(),
			"tencentcloud_teo_security_policy_regions":               dataSourceTencentCloudTeoSecurityPolicyRegions(),
			"tencentcloud_teo_waf_rule_groups":                       dataSourceTencentCloudTeoWafRuleGroups(),
			"tencentcloud_teo_zone_ddos_policy":                      dataSourceTencentCloudTeoZoneDDoSPolicy(),
			"tencentcloud_sts_caller_identity":                       dataSourceTencentCloudStsCallerIdentity(),
			"tencentcloud_dcdb_instances":                            dataSourceTencentCloudDcdbInstances(),
			"tencentcloud_dcdb_accounts":                             dataSourceTencentCloudDcdbAccounts(),
			"tencentcloud_dcdb_databases":                            dataSourceTencentCloudDcdbDatabases(),
			"tencentcloud_dcdb_parameters":                           dataSourceTencentCloudDcdbParameters(),
			"tencentcloud_dcdb_shards":                               dataSourceTencentCloudDcdbShards(),
			"tencentcloud_dcdb_security_groups":                      dataSourceTencentCloudDcdbSecurityGroups(),
			"tencentcloud_dcdb_database_objects":                     dataSourceTencentCloudDcdbDatabaseObjects(),
			"tencentcloud_dcdb_database_tables":                      dataSourceTencentCloudDcdbDatabaseTables(),
			"tencentcloud_mariadb_db_instances":                      dataSourceTencentCloudMariadbDbInstances(),
			"tencentcloud_mariadb_accounts":                          dataSourceTencentCloudMariadbAccounts(),
			"tencentcloud_mariadb_security_groups":                   dataSourceTencentCloudMariadbSecurityGroups(),
			"tencentcloud_mariadb_database_objects":                  dataSourceTencentCloudMariadbDatabaseObjects(),
			"tencentcloud_mariadb_databases":                         dataSourceTencentCloudMariadbDatabases(),
			"tencentcloud_mariadb_database_table":                    dataSourceTencentCloudMariadbDatabaseTable(),
			"tencentcloud_tdcpg_clusters":                            dataSourceTencentCloudTdcpgClusters(),
			"tencentcloud_tdcpg_instances":                           dataSourceTencentCloudTdcpgInstances(),
			"tencentcloud_cat_probe_data":                            dataSourceTencentCloudCatProbeData(),
			"tencentcloud_cat_node":                                  dataSourceTencentCloudCatNode(),
			"tencentcloud_rum_project":                               dataSourceTencentCloudRumProject(),
			"tencentcloud_rum_offline_log_config":                    dataSourceTencentCloudRumOfflineLogConfig(),
			"tencentcloud_rum_whitelist":                             dataSourceTencentCloudRumWhitelist(),
			"tencentcloud_rum_taw_instance":                          dataSourceTencentCloudRumTawInstance(),
			"tencentcloud_dnspod_records":                            dataSourceTencentCloudDnspodRecords(),
			"tencentcloud_tat_command":                               dataSourceTencentCloudTatCommand(),
			"tencentcloud_tat_invoker":                               dataSourceTencentCloudTatInvoker(),
			"tencentcloud_tat_invoker_records":                       dataSourceTencentCloudTatInvokerRecords(),
			"tencentcloud_tat_agent":                                 dataSourceTencentCloudTatAgent(),
			"tencentcloud_tat_invocation_task":                       dataSourceTencentCloudTatInvocationTask(),
			"tencentcloud_dbbrain_sql_filters":                       dataSourceTencentCloudDbbrainSqlFilters(),
			"tencentcloud_dbbrain_security_audit_log_export_tasks":   dataSourceTencentCloudDbbrainSecurityAuditLogExportTasks(),
			"tencentcloud_dbbrain_diag_event":                        dataSourceTencentCloudDbbrainDiagEvent(),
			"tencentcloud_dbbrain_diag_events":                       dataSourceTencentCloudDbbrainDiagEvents(),
			"tencentcloud_dbbrain_diag_history":                      dataSourceTencentCloudDbbrainDiagHistory(),
			"tencentcloud_dbbrain_security_audit_log_download_urls":  dataSourceTencentCloudDbbrainSecurityAuditLogDownloadUrls(),
			"tencentcloud_dbbrain_slow_log_time_series_stats":        dataSourceTencentCloudDbbrainSlowLogTimeSeriesStats(),
			"tencentcloud_dbbrain_slow_log_top_sqls":                 dataSourceTencentCloudDbbrainSlowLogTopSqls(),
			"tencentcloud_dbbrain_slow_log_user_host_stats":          dataSourceTencentCloudDbbrainSlowLogUserHostStats(),
			"tencentcloud_dbbrain_slow_log_user_sql_advice":          dataSourceTencentCloudDbbrainSlowLogUserSqlAdvice(),
			"tencentcloud_dbbrain_slow_logs":                         dataSourceTencentCloudDbbrainSlowLogs(),
			"tencentcloud_dbbrain_health_scores":                     dataSourceTencentCloudDbbrainHealthScores(),
			"tencentcloud_dbbrain_sql_templates":                     dataSourceTencentCloudDbbrainSqlTemplates(),
			"tencentcloud_dbbrain_db_space_status":                   dataSourceTencentCloudDbbrainDbSpaceStatus(),
			"tencentcloud_dbbrain_top_space_schemas":                 dataSourceTencentCloudDbbrainTopSpaceSchemas(),
			"tencentcloud_dbbrain_top_space_tables":                  dataSourceTencentCloudDbbrainTopSpaceTables(),
			"tencentcloud_dbbrain_top_space_schema_time_series":      dataSourceTencentCloudDbbrainTopSpaceSchemaTimeSeries(),
			"tencentcloud_dbbrain_top_space_table_time_series":       dataSourceTencentCloudDbbrainTopSpaceTableTimeSeries(),
			"tencentcloud_dts_sync_jobs":                             dataSourceTencentCloudDtsSyncJobs(),
			"tencentcloud_dts_compare_tasks":                         dataSourceTencentCloudDtsCompareTasks(),
			"tencentcloud_dts_migrate_jobs":                          dataSourceTencentCloudDtsMigrateJobs(),
			"tencentcloud_tdmq_rocketmq_cluster":                     dataSourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloud_tdmq_rocketmq_namespace":                   dataSourceTencentCloudTdmqRocketmqNamespace(),
			"tencentcloud_tdmq_rocketmq_topic":                       dataSourceTencentCloudTdmqRocketmqTopic(),
			"tencentcloud_tdmq_rocketmq_role":                        dataSourceTencentCloudTdmqRocketmqRole(),
			"tencentcloud_tdmq_rocketmq_group":                       dataSourceTencentCloudTdmqRocketmqGroup(),
			"tencentcloud_tcmq_queue":                                dataSourceTencentCloudTcmqQueue(),
			"tencentcloud_tcmq_topic":                                dataSourceTencentCloudTcmqTopic(),
			"tencentcloud_tcmq_subscribe":                            dataSourceTencentCloudTcmqSubscribe(),
			"tencentcloud_as_instances":                              dataSourceTencentCloudAsInstances(),
			"tencentcloud_cynosdb_accounts":                          dataSourceTencentCloudCynosdbAccounts(),
			"tencentcloud_cynosdb_cluster_instance_groups":           dataSourceTencentCloudCynosdbClusterInstanceGroups(),
			"tencentcloud_cynosdb_cluster_params":                    dataSourceTencentCloudCynosdbClusterParams(),
			"tencentcloud_cynosdb_param_templates":                   dataSourceTencentCloudCynosdbParamTemplates(),
			"tencentcloud_cvm_instances_modification":                dataSourceTencentCloudCvmInstancesModification(),
			"tencentcloud_css_domains":                               dataSourceTencentCloudCssDomains(),
			"tencentcloud_chdfs_access_groups":                       dataSourceTencentCloudChdfsAccessGroups(),
			"tencentcloud_chdfs_mount_points":                        dataSourceTencentCloudChdfsMountPoints(),
			"tencentcloud_tcm_mesh":                                  dataSourceTencentCloudTcmMesh(),
			"tencentcloud_lighthouse_firewall_rules_template":        dataSourceTencentCloudLighthouseFirewallRulesTemplate(),
			"tencentcloud_cvm_instance_vnc_url":                      dataSourceTencentCloudCvmInstanceVncUrl(),
			"tencentcloud_cvm_disaster_recover_group_quota":          dataSourceTencentCloudCvmDisasterRecoverGroupQuota(),
			"tencentcloud_cvm_chc_hosts":                             dataSourceTencentCloudCvmChcHosts(),
			"tencentcloud_cvm_chc_denied_actions":                    dataSourceTencentCloudCvmChcDeniedActions(),
			"tencentcloud_tsf_application":                           dataSourceTencentCloudTsfApplication(),
			"tencentcloud_tsf_application_config":                    dataSourceTencentCloudTsfApplicationConfig(),
			"tencentcloud_tsf_application_file_config":               dataSourceTencentCloudTsfApplicationFileConfig(),
			"tencentcloud_tsf_application_public_config":             dataSourceTencentCloudTsfApplicationPublicConfig(),
			"tencentcloud_tsf_cluster":                               dataSourceTencentCloudTsfCluster(),
			"tencentcloud_tsf_microservice":                          dataSourceTencentCloudTsfMicroservice(),
			"tencentcloud_tsf_unit_rules":                            dataSourceTencentCloudTsfUnitRules(),
			"tencentcloud_tsf_config_summary":                        dataSourceTencentCloudTsfConfigSummary(),
			"tencentcloud_tsf_delivery_config_by_group_id":           dataSourceTencentCloudTsfDeliveryConfigByGroupId(),
			"tencentcloud_tsf_delivery_configs":                      dataSourceTencentCloudTsfDeliveryConfigs(),
			"tencentcloud_tsf_public_config_summary":                 dataSourceTencentCloudTsfPublicConfigSummary(),
			"tencentcloud_tsf_api_group":                             dataSourceTencentCloudTsfApiGroup(),
			"tencentcloud_tsf_application_attribute":                 dataSourceTencentCloudTsfApplicationAttribute(),
			"tencentcloud_tsf_business_log_configs":                  dataSourceTencentCloudTsfBusinessLogConfigs(),
			"tencentcloud_tsf_api_detail":                            dataSourceTencentCloudTsfApiDetail(),
			"tencentcloud_tsf_microservice_api_version":              dataSourceTencentCloudTsfMicroserviceApiVersion(),
			"tencentcloud_lighthouse_bundle":                         dataSourceTencentCloudLighthouseBundle(),
			"tencentcloud_api_gateway_api_docs":                      dataSourceTencentCloudAPIGatewayAPIDocs(),
			"tencentcloud_api_gateway_api_apps":                      dataSourceTencentCloudAPIGatewayAPIApps(),
			"tencentcloud_tse_access_address":                        dataSourceTencentCloudTseAccessAddress(),
			"tencentcloud_tse_nacos_replicas":                        dataSourceTencentCloudTseNacosReplicas(),
			"tencentcloud_tse_nacos_server_interfaces":               dataSourceTencentCloudTseNacosServerInterfaces(),
			"tencentcloud_tse_zookeeper_replicas":                    dataSourceTencentCloudTseZookeeperReplicas(),
			"tencentcloud_tse_zookeeper_server_interfaces":           dataSourceTencentCloudTseZookeeperServerInterfaces(),
			"tencentcloud_lighthouse_zone":                           dataSourceTencentCloudLighthouseZone(),
			"tencentcloud_lighthouse_scene":                          dataSourceTencentCloudLighthouseScene(),
			"tencentcloud_lighthouse_reset_instance_blueprint":       dataSourceTencentCloudLighthouseResetInstanceBlueprint(),
			"tencentcloud_lighthouse_region":                         dataSourceTencentCloudLighthouseRegion(),
			"tencentcloud_lighthouse_instance_vnc_url":               dataSourceTencentCloudLighthouseInstanceVncUrl(),
			"tencentcloud_lighthouse_instance_traffic_package":       dataSourceTencentCloudLighthouseInstanceTrafficPackage(),
			"tencentcloud_lighthouse_instance_disk_num":              dataSourceTencentCloudLighthouseInstanceDiskNum(),
			"tencentcloud_lighthouse_instance_blueprint":             dataSourceTencentCloudLighthouseInstanceBlueprint(),
			"tencentcloud_lighthouse_disk_config":                    dataSourceTencentCloudLighthouseDiskConfig(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"tencentcloud_emr_cluster":                                 resourceTencentCloudEmrCluster(),
			"tencentcloud_instance":                                    resourceTencentCloudInstance(),
			"tencentcloud_instance_set":                                resourceTencentCloudInstanceSet(),
			"tencentcloud_reserved_instance":                           resourceTencentCloudReservedInstance(),
			"tencentcloud_key_pair":                                    resourceTencentCloudKeyPair(),
			"tencentcloud_placement_group":                             resourceTencentCloudPlacementGroup(),
			"tencentcloud_cbs_snapshot":                                resourceTencentCloudCbsSnapshot(),
			"tencentcloud_cbs_snapshot_policy":                         resourceTencentCloudCbsSnapshotPolicy(),
			"tencentcloud_cbs_storage":                                 resourceTencentCloudCbsStorage(),
			"tencentcloud_cbs_storage_set":                             resourceTencentCloudCbsStorageSet(),
			"tencentcloud_cbs_storage_attachment":                      resourceTencentCloudCbsStorageAttachment(),
			"tencentcloud_cbs_storage_set_attachment":                  resourceTencentCloudCbsStorageSetAttachment(),
			"tencentcloud_cbs_snapshot_policy_attachment":              resourceTencentCloudCbsSnapshotPolicyAttachment(),
			"tencentcloud_vpc":                                         resourceTencentCloudVpcInstance(),
			"tencentcloud_vpc_acl":                                     resourceTencentCloudVpcACL(),
			"tencentcloud_vpc_acl_attachment":                          resourceTencentCloudVpcAclAttachment(),
			"tencentcloud_vpc_bandwidth_package":                       resourceTencentCloudVpcBandwidthPackage(),
			"tencentcloud_vpc_bandwidth_package_attachment":            resourceTencentCloudVpcBandwidthPackageAttachment(),
			"tencentcloud_vpc_traffic_package":                         resourceTencentCloudVpcTrafficPackage(),
			"tencentcloud_ipv6_address_bandwidth":                      resourceTencentCloudIpv6AddressBandwidth(),
			"tencentcloud_subnet":                                      resourceTencentCloudVpcSubnet(),
			"tencentcloud_route_entry":                                 resourceTencentCloudRouteEntry(),
			"tencentcloud_route_table_entry":                           resourceTencentCloudVpcRouteEntry(),
			"tencentcloud_route_table":                                 resourceTencentCloudVpcRouteTable(),
			"tencentcloud_dnat":                                        resourceTencentCloudDnat(),
			"tencentcloud_nat_gateway":                                 resourceTencentCloudNatGateway(),
			"tencentcloud_nat_gateway_snat":                            resourceTencentCloudNatGatewaySnat(),
			"tencentcloud_nat_refresh_nat_dc_route":                    resourceTencentCloudNatRefreshNatDcRoute(),
			"tencentcloud_eip":                                         resourceTencentCloudEip(),
			"tencentcloud_eip_association":                             resourceTencentCloudEipAssociation(),
			"tencentcloud_eip_address_transform":                       resourceTencentCloudEipAddressTransform(),
			"tencentcloud_eip_public_address_adjust":                   resourceTencentCloudEipPublicAddressAdjust(),
			"tencentcloud_eip_normal_address_return":                   resourceTencentCloudEipNormalAddressReturn(),
			"tencentcloud_eni":                                         resourceTencentCloudEni(),
			"tencentcloud_eni_attachment":                              resourceTencentCloudEniAttachment(),
			"tencentcloud_ccn":                                         resourceTencentCloudCcn(),
			"tencentcloud_ccn_attachment":                              resourceTencentCloudCcnAttachment(),
			"tencentcloud_ccn_bandwidth_limit":                         resourceTencentCloudCcnBandwidthLimit(),
			"tencentcloud_ccn_routes":                                  resourceTencentCloudCcnRoutes(),
			"tencentcloud_ccn_instances_accept_attach":                 resourceTencentCloudCcnInstancesAcceptAttach(),
			"tencentcloud_ccn_instances_reset_attach":                  resourceTencentCloudCcnInstancesResetAttach(),
			"tencentcloud_dcx":                                         resourceTencentCloudDcxInstance(),
			"tencentcloud_dc_gateway":                                  resourceTencentCloudDcGatewayInstance(),
			"tencentcloud_dc_gateway_ccn_route":                        resourceTencentCloudDcGatewayCcnRouteInstance(),
			"tencentcloud_vpn_customer_gateway":                        resourceTencentCloudVpnCustomerGateway(),
			"tencentcloud_vpn_gateway":                                 resourceTencentCloudVpnGateway(),
			"tencentcloud_vpn_gateway_route":                           resourceTencentCloudVpnGatewayRoute(),
			"tencentcloud_vpn_connection":                              resourceTencentCloudVpnConnection(),
			"tencentcloud_vpn_ssl_server":                              resourceTencentCloudVpnSslServer(),
			"tencentcloud_vpn_ssl_client":                              resourceTencentCloudVpnSslClient(),
			"tencentcloud_vpn_connection_reset":                        resourceTencentCloudVpnConnectionReset(),
			"tencentcloud_vpn_customer_gateway_configuration_download": resourceTencentCloudVpnCustomerGatewayConfigurationDownload(),
			"tencentcloud_vpn_gateway_ssl_client_cert":                 resourceTencentCloudVpnGatewaySslClientCert(),
			"tencentcloud_vpn_gateway_ccn_routes":                      resourceTencentCloudVpnGatewayCcnRoutes(),
			"tencentcloud_ha_vip":                                      resourceTencentCloudHaVip(),
			"tencentcloud_ha_vip_eip_attachment":                       resourceTencentCloudHaVipEipAttachment(),
			"tencentcloud_security_group":                              resourceTencentCloudSecurityGroup(),
			"tencentcloud_security_group_rule":                         resourceTencentCloudSecurityGroupRule(),
			"tencentcloud_security_group_rule_set":                     resourceTencentCloudSecurityGroupRuleSet(),
			"tencentcloud_security_group_lite_rule":                    resourceTencentCloudSecurityGroupLiteRule(),
			"tencentcloud_lb":                                          resourceTencentCloudLB(),
			"tencentcloud_alb_server_attachment":                       resourceTencentCloudAlbServerAttachment(),
			"tencentcloud_clb_instance":                                resourceTencentCloudClbInstance(),
			"tencentcloud_clb_listener":                                resourceTencentCloudClbListener(),
			"tencentcloud_clb_listener_rule":                           resourceTencentCloudClbListenerRule(),
			"tencentcloud_clb_attachment":                              resourceTencentCloudClbServerAttachment(),
			"tencentcloud_clb_redirection":                             resourceTencentCloudClbRedirection(),
			"tencentcloud_clb_target_group":                            resourceTencentCloudClbTargetGroup(),
			"tencentcloud_clb_target_group_instance_attachment":        resourceTencentCloudClbTGAttachmentInstance(),
			"tencentcloud_clb_target_group_attachment":                 resourceTencentCloudClbTargetGroupAttachment(),
			"tencentcloud_clb_log_set":                                 resourceTencentCloudClbLogSet(),
			"tencentcloud_clb_log_topic":                               resourceTencentCloudClbLogTopic(),
			"tencentcloud_clb_customized_config":                       resourceTencentCloudClbCustomizedConfig(),
			"tencentcloud_clb_snat_ip":                                 resourceTencentCloudClbSnatIp(),
			"tencentcloud_clb_function_targets_attachment":             resourceTencentCloudClbFunctionTargetsAttachment(),
			"tencentcloud_clb_instance_mix_ip_target_config":           resourceTencentCloudClbInstanceMixIpTargetConfig(),
			"tencentcloud_clb_instance_sla_config":                     resourceTencentCloudClbInstanceSlaConfig(),
			"tencentcloud_clb_replace_cert_for_lbs":                    resourceTencentCloudClbReplaceCertForLbs(),
			"tencentcloud_container_cluster":                           resourceTencentCloudContainerCluster(),
			"tencentcloud_container_cluster_instance":                  resourceTencentCloudContainerClusterInstance(),
			"tencentcloud_kubernetes_cluster":                          resourceTencentCloudTkeCluster(),
			"tencentcloud_kubernetes_cluster_endpoint":                 resourceTencentCloudTkeClusterEndpoint(),
			"tencentcloud_eks_cluster":                                 resourceTencentCloudEksCluster(),
			"tencentcloud_eks_container_instance":                      resourceTencentCloudEksContainerInstance(),
			"tencentcloud_kubernetes_addon_attachment":                 resourceTencentCloudTkeAddonAttachment(),
			"tencentcloud_kubernetes_auth_attachment":                  resourceTencentCloudTKEAuthAttachment(),
			"tencentcloud_kubernetes_as_scaling_group":                 resourceTencentCloudKubernetesAsScalingGroup(),
			"tencentcloud_kubernetes_scale_worker":                     resourceTencentCloudTkeScaleWorker(),
			"tencentcloud_kubernetes_cluster_attachment":               resourceTencentCloudTkeClusterAttachment(),
			"tencentcloud_kubernetes_node_pool":                        resourceTencentCloudKubernetesNodePool(),
			"tencentcloud_kubernetes_serverless_node_pool":             resourceTkeServerLessNodePool(),
			"tencentcloud_kubernetes_backup_storage_location":          resourceTencentCloudTkeBackupStorageLocation(),
			"tencentcloud_mysql_backup_policy":                         resourceTencentCloudMysqlBackupPolicy(),
			"tencentcloud_mysql_backup_policy_v2":                      resourceTencentCloudMysqlBackupPolicyV2(),
			"tencentcloud_mysql_account":                               resourceTencentCloudMysqlAccount(),
			"tencentcloud_mysql_account_privilege":                     resourceTencentCloudMysqlAccountPrivilege(),
			"tencentcloud_mysql_privilege":                             resourceTencentCloudMysqlPrivilege(),
			"tencentcloud_mysql_instance":                              resourceTencentCloudMysqlInstance(),
			"tencentcloud_mysql_database":                              resourceTencentCloudMysqlDatabase(),
			"tencentcloud_mysql_readonly_instance":                     resourceTencentCloudMysqlReadonlyInstance(),
			"tencentcloud_mysql_time_window":                           resourceTencentCloudMysqlTimeWindow(),
			"tencentcloud_mysql_param_template":                        resourceTencentCloudMysqlParamTemplate(),
			"tencentcloud_mysql_security_groups_attachment":            resourceTencentCloudMysqlSecurityGroupsAttachment(),
			"tencentcloud_mysql_deploy_group":                          resourceTencentCloudMysqlDeployGroup(),
			"tencentcloud_mysql_local_binlog_config":                   resourceTencentCloudMysqlLocalBinlogConfig(),
			"tencentcloud_mysql_audit_log_file":                        resourceTencentCloudMysqlAuditLogFile(),
			"tencentcloud_cos_bucket":                                  resourceTencentCloudCosBucket(),
			"tencentcloud_cos_bucket_object":                           resourceTencentCloudCosBucketObject(),
			"tencentcloud_cfs_file_system":                             resourceTencentCloudCfsFileSystem(),
			"tencentcloud_cfs_access_group":                            resourceTencentCloudCfsAccessGroup(),
			"tencentcloud_cfs_access_rule":                             resourceTencentCloudCfsAccessRule(),
			"tencentcloud_cfs_auto_snapshot_policy":                    resourceTencentCloudCfsAutoSnapshotPolicy(),
			"tencentcloud_cfs_auto_snapshot_policy_attachment":         resourceTencentCloudCfsAutoSnapshotPolicyAttachment(),
			"tencentcloud_cfs_snapshot":                                resourceTencentCloudCfsSnapshot(),
			"tencentcloud_cfs_user_quota":                              resourceTencentCloudCfsUserQuota(),
			"tencentcloud_cfs_sign_up_cfs_service":                     resourceTencentCloudCfsSignUpCfsService(),
			"tencentcloud_redis_instance":                              resourceTencentCloudRedisInstance(),
			"tencentcloud_redis_backup_config":                         resourceTencentCloudRedisBackupConfig(),
			"tencentcloud_redis_account":                               resourceTencentCloudRedisAccount(),
			"tencentcloud_redis_param_template":                        resourceTencentCloudRedisParamTemplate(),
			"tencentcloud_redis_connection_config":                     resourceTencentCloudRedisConnectionConfig(),
			"tencentcloud_redis_param":                                 resourceTencentCloudRedisParam(),
			"tencentcloud_redis_read_only":                             resourceTencentCloudRedisReadOnly(),
			"tencentcloud_redis_ssl":                                   resourceTencentCloudRedisSsl(),
			"tencentcloud_redis_backup_download_restriction":           resourceTencentCloudRedisBackupDownloadRestriction(),
			"tencentcloud_redis_clear_instance_operation":              resourceTencentCloudRedisClearInstanceOperation(),
			"tencentcloud_redis_renew_instance_operation":              resourceTencentCloudRedisRenewInstanceOperation(),
			"tencentcloud_redis_startup_instance_operation":            resourceTencentCloudRedisStartupInstanceOperation(),
			"tencentcloud_redis_upgrade_cache_version_operation":       resourceTencentCloudRedisUpgradeCacheVersionOperation(),
			"tencentcloud_redis_upgrade_multi_zone_operation":          resourceTencentCloudRedisUpgradeMultiZoneOperation(),
			"tencentcloud_redis_upgrade_proxy_version_operation":       resourceTencentCloudRedisUpgradeProxyVersionOperation(),
			"tencentcloud_redis_maintenance_window":                    resourceTencentCloudRedisMaintenanceWindow(),
			"tencentcloud_redis_replica_readonly":                      resourceTencentCloudRedisReplicaReadonly(),
			"tencentcloud_redis_switch_master":                         resourceTencentCloudRedisSwitchMaster(),
			"tencentcloud_as_scaling_config":                           resourceTencentCloudAsScalingConfig(),
			"tencentcloud_as_scaling_group":                            resourceTencentCloudAsScalingGroup(),
			"tencentcloud_as_attachment":                               resourceTencentCloudAsAttachment(),
			"tencentcloud_as_scaling_policy":                           resourceTencentCloudAsScalingPolicy(),
			"tencentcloud_as_schedule":                                 resourceTencentCloudAsSchedule(),
			"tencentcloud_as_lifecycle_hook":                           resourceTencentCloudAsLifecycleHook(),
			"tencentcloud_as_notification":                             resourceTencentCloudAsNotification(),
			"tencentcloud_as_remove_instances":                         resourceTencentCloudAsRemoveInstances(),
			"tencentcloud_as_protect_instances":                        resourceTencentCloudAsProtectInstances(),
			"tencentcloud_as_start_instances":                          resourceTencentCloudAsStartInstances(),
			"tencentcloud_as_stop_instances":                           resourceTencentCloudAsStopInstances(),
			"tencentcloud_mongodb_instance":                            resourceTencentCloudMongodbInstance(),
			"tencentcloud_mongodb_sharding_instance":                   resourceTencentCloudMongodbShardingInstance(),
			"tencentcloud_mongodb_instance_account":                    resourceTencentCloudMongodbInstanceAccount(),
			"tencentcloud_mongodb_instance_backup":                     resourceTencentCloudMongodbInstanceBackup(),
			"tencentcloud_mongodb_instance_backup_download_task":       resourceTencentCloudMongodbInstanceBackupDownloadTask(),
			"tencentcloud_dayu_cc_http_policy":                         resourceTencentCloudDayuCCHttpPolicy(),
			"tencentcloud_dayu_cc_https_policy":                        resourceTencentCloudDayuCCHttpsPolicy(),
			"tencentcloud_dayu_ddos_policy":                            resourceTencentCloudDayuDdosPolicy(),
			"tencentcloud_dayu_cc_policy_v2":                           resourceTencentCloudDayuCCPolicyV2(),
			"tencentcloud_dayu_ddos_policy_v2":                         resourceTencentCloudDayuDdosPolicyV2(),
			"tencentcloud_dayu_ddos_policy_case":                       resourceTencentCloudDayuDdosPolicyCase(),
			"tencentcloud_dayu_ddos_policy_attachment":                 resourceTencentCloudDayuDdosPolicyAttachment(),
			"tencentcloud_dayu_l4_rule":                                resourceTencentCloudDayuL4Rule(),
			"tencentcloud_dayu_l4_rule_v2":                             resourceTencentCloudDayuL4RuleV2(),
			"tencentcloud_dayu_l7_rule":                                resourceTencentCloudDayuL7Rule(),
			"tencentcloud_dayu_l7_rule_v2":                             resourceTencentCloudDayuL7RuleV2(),
			"tencentcloud_dayu_eip":                                    resourceTencentCloudDayuEip(),
			"tencentcloud_gaap_proxy":                                  resourceTencentCloudGaapProxy(),
			"tencentcloud_gaap_realserver":                             resourceTencentCloudGaapRealserver(),
			"tencentcloud_gaap_layer4_listener":                        resourceTencentCloudGaapLayer4Listener(),
			"tencentcloud_gaap_layer7_listener":                        resourceTencentCloudGaapLayer7Listener(),
			"tencentcloud_gaap_http_domain":                            resourceTencentCloudGaapHttpDomain(),
			"tencentcloud_gaap_http_rule":                              resourceTencentCloudGaapHttpRule(),
			"tencentcloud_gaap_certificate":                            resourceTencentCloudGaapCertificate(),
			"tencentcloud_gaap_security_policy":                        resourceTencentCloudGaapSecurityPolicy(),
			"tencentcloud_gaap_security_rule":                          resourceTencentCloudGaapSecurityRule(),
			"tencentcloud_gaap_domain_error_page":                      resourceTencentCloudGaapDomainErrorPageInfo(),
			"tencentcloud_ssl_certificate":                             resourceTencentCloudSslCertificate(),
			"tencentcloud_ssl_pay_certificate":                         resourceTencentCloudSSLInstance(),
			"tencentcloud_ssl_free_certificate":                        resourceTencentCloudSSLFreeCertificate(),
			"tencentcloud_cam_role":                                    resourceTencentCloudCamRole(),
			"tencentcloud_cam_role_by_name":                            resourceTencentCloudCamRoleByName(),
			"tencentcloud_cam_user":                                    resourceTencentCloudCamUser(),
			"tencentcloud_cam_policy":                                  resourceTencentCloudCamPolicy(),
			"tencentcloud_cam_policy_by_name":                          resourceTencentCloudCamPolicyByName(),
			"tencentcloud_cam_role_policy_attachment":                  resourceTencentCloudCamRolePolicyAttachment(),
			"tencentcloud_cam_role_policy_attachment_by_name":          resourceTencentCloudCamRolePolicyAttachmentByName(),
			"tencentcloud_cam_user_policy_attachment":                  resourceTencentCloudCamUserPolicyAttachment(),
			"tencentcloud_cam_group_policy_attachment":                 resourceTencentCloudCamGroupPolicyAttachment(),
			"tencentcloud_cam_group":                                   resourceTencentCloudCamGroup(),
			"tencentcloud_cam_oidc_sso":                                resourceTencentCloudCamOIDCSSO(),
			"tencentcloud_cam_role_sso":                                resourceTencentCloudCamRoleSSO(),
			"tencentcloud_cam_group_membership":                        resourceTencentCloudCamGroupMembership(),
			"tencentcloud_cam_saml_provider":                           resourceTencentCloudCamSAMLProvider(),
			"tencentcloud_cam_service_linked_role":                     resourceTencentCloudCamServiceLinkedRole(),
			"tencentcloud_cam_user_saml_config":                        resourceTencentCloudCamUserSamlConfig(),
			"tencentcloud_scf_function":                                resourceTencentCloudScfFunction(),
			"tencentcloud_scf_namespace":                               resourceTencentCloudScfNamespace(),
			"tencentcloud_scf_layer":                                   resourceTencentCloudScfLayer(),
			"tencentcloud_scf_function_alias":                          resourceTencentCloudScfFunctionAlias(),
			"tencentcloud_tcaplus_cluster":                             resourceTencentCloudTcaplusCluster(),
			"tencentcloud_tcaplus_tablegroup":                          resourceTencentCloudTcaplusTableGroup(),
			"tencentcloud_tcaplus_idl":                                 resourceTencentCloudTcaplusIdl(),
			"tencentcloud_tcaplus_table":                               resourceTencentCloudTcaplusTable(),
			"tencentcloud_cdn_domain":                                  resourceTencentCloudCdnDomain(),
			"tencentcloud_cdn_url_push":                                resourceTencentCloudUrlPush(),
			"tencentcloud_cdn_url_purge":                               resourceTencentCloudUrlPurge(),
			"tencentcloud_monitor_policy_group":                        resourceTencentCloudMonitorPolicyGroup(),
			"tencentcloud_monitor_binding_object":                      resourceTencentCloudMonitorBindingObject(),
			"tencentcloud_monitor_policy_binding_object":               resourceTencentCloudMonitorPolicyBindingObject(),
			"tencentcloud_monitor_binding_receiver":                    resourceTencentCloudMonitorBindingAlarmReceiver(),
			"tencentcloud_monitor_alarm_policy":                        resourceTencentCloudMonitorAlarmPolicy(),
			"tencentcloud_monitor_alarm_notice":                        resourceTencentCloudMonitorAlarmNotice(),
			"tencentcloud_monitor_tmp_instance":                        resourceTencentCloudMonitorTmpInstance(),
			"tencentcloud_monitor_tmp_cvm_agent":                       resourceTencentCloudMonitorTmpCvmAgent(),
			"tencentcloud_monitor_tmp_scrape_job":                      resourceTencentCloudMonitorTmpScrapeJob(),
			"tencentcloud_monitor_tmp_exporter_integration":            resourceTencentCloudMonitorTmpExporterIntegration(),
			"tencentcloud_monitor_tmp_alert_rule":                      resourceTencentCloudMonitorTmpAlertRule(),
			"tencentcloud_monitor_tmp_recording_rule":                  resourceTencentCloudMonitorTmpRecordingRule(),
			"tencentcloud_monitor_tmp_tke_template":                    resourceTencentCloudMonitorTmpTkeTemplate(),
			"tencentcloud_monitor_tmp_tke_template_attachment":         resourceTencentCloudMonitorTmpTkeTemplateAttachment(),
			"tencentcloud_monitor_tmp_tke_alert_policy":                resourceTencentCloudMonitorTmpTkeAlertPolicy(),
			"tencentcloud_monitor_tmp_tke_basic_config":                resourceTencentCloudMonitorTmpTkeBasicConfig(),
			"tencentcloud_monitor_tmp_tke_cluster_agent":               resourceTencentCloudMonitorTmpTkeClusterAgent(),
			"tencentcloud_monitor_tmp_tke_config":                      resourceTencentCloudMonitorTmpTkeConfig(),
			"tencentcloud_monitor_tmp_tke_record_rule_yaml":            resourceTencentCloudMonitorTmpTkeRecordRuleYaml(),
			"tencentcloud_monitor_tmp_tke_global_notification":         resourceTencentCloudMonitorTmpTkeGlobalNotification(),
			"tencentcloud_monitor_tmp_manage_grafana_attachment":       resourceTencentCloudMonitorTmpManageGrafanaAttachment(),
			"tencentcloud_monitor_grafana_instance":                    resourceTencentCloudMonitorGrafanaInstance(),
			"tencentcloud_monitor_grafana_integration":                 resourceTencentCloudMonitorGrafanaIntegration(),
			"tencentcloud_monitor_grafana_notification_channel":        resourceTencentCloudMonitorGrafanaNotificationChannel(),
			"tencentcloud_monitor_grafana_plugin":                      resourceTencentCloudMonitorGrafanaPlugin(),
			"tencentcloud_monitor_grafana_sso_account":                 resourceTencentCloudMonitorGrafanaSsoAccount(),
			"tencentcloud_monitor_tmp_grafana_config":                  resourceTencentCloudMonitorTmpGrafanaConfig(),
			"tencentcloud_mongodb_standby_instance":                    resourceTencentCloudMongodbStandbyInstance(),
			"tencentcloud_elasticsearch_instance":                      resourceTencentCloudElasticsearchInstance(),
			"tencentcloud_postgresql_instance":                         resourceTencentCloudPostgresqlInstance(),
			"tencentcloud_postgresql_readonly_instance":                resourceTencentCloudPostgresqlReadonlyInstance(),
			"tencentcloud_postgresql_readonly_group":                   resourceTencentCloudPostgresqlReadonlyGroup(),
			"tencentcloud_postgresql_readonly_attachment":              resourceTencentCloudPostgresqlReadonlyAttachment(),
			"tencentcloud_postgresql_parameter_template":               resourceTencentCloudPostgresqlParameterTemplate(),
			"tencentcloud_sqlserver_instance":                          resourceTencentCloudSqlserverInstance(),
			"tencentcloud_sqlserver_db":                                resourceTencentCloudSqlserverDB(),
			"tencentcloud_sqlserver_account":                           resourceTencentCloudSqlserverAccount(),
			"tencentcloud_sqlserver_account_db_attachment":             resourceTencentCloudSqlserverAccountDBAttachment(),
			"tencentcloud_sqlserver_readonly_instance":                 resourceTencentCloudSqlserverReadonlyInstance(),
			"tencentcloud_sqlserver_migration":                         resourceTencentCloudSqlserverMigration(),
			"tencentcloud_sqlserver_config_backup_strategy":            resourceTencentCloudSqlserverConfigBackupStrategy(),
			"tencentcloud_sqlserver_general_backup":                    resourceTencentCloudSqlserverGeneralBackup(),
			"tencentcloud_sqlserver_general_clone":                     resourceTencentCloudSqlserverGeneralClone(),
			"tencentcloud_sqlserver_full_backup_migration":             resourceTencentCloudSqlserverFullBackupMigration(),
			"tencentcloud_sqlserver_incre_backup_migration":            resourceTencentCloudSqlserverIncreBackupMigration(),
			"tencentcloud_sqlserver_business_intelligence_file":        resourceTencentCloudSqlserverBusinessIntelligenceFile(),
			"tencentcloud_sqlserver_business_intelligence_instance":    resourceTencentCloudSqlserverBusinessIntelligenceInstance(),
			"tencentcloud_sqlserver_general_communication":             resourceTencentCloudSqlserverGeneralCommunication(),
			"tencentcloud_ckafka_instance":                             resourceTencentCloudCkafkaInstance(),
			"tencentcloud_ckafka_user":                                 resourceTencentCloudCkafkaUser(),
			"tencentcloud_ckafka_acl":                                  resourceTencentCloudCkafkaAcl(),
			"tencentcloud_ckafka_topic":                                resourceTencentCloudCkafkaTopic(),
			"tencentcloud_ckafka_datahub_topic":                        resourceTencentCloudCkafkaDatahubTopic(),
			"tencentcloud_ckafka_connect_resource":                     resourceTencentCloudCkafkaConnectResource(),
			"tencentcloud_audit":                                       resourceTencentCloudAudit(),
			"tencentcloud_audit_track":                                 resourceTencentCloudAuditTrack(),
			"tencentcloud_image":                                       resourceTencentCloudImage(),
			"tencentcloud_cynosdb_cluster":                             resourceTencentCloudCynosdbCluster(),
			"tencentcloud_cynosdb_readonly_instance":                   resourceTencentCloudCynosdbReadonlyInstance(),
			"tencentcloud_vod_adaptive_dynamic_streaming_template":     resourceTencentCloudVodAdaptiveDynamicStreamingTemplate(),
			"tencentcloud_vod_image_sprite_template":                   resourceTencentCloudVodImageSpriteTemplate(),
			"tencentcloud_vod_procedure_template":                      resourceTencentCloudVodProcedureTemplate(),
			"tencentcloud_vod_snapshot_by_time_offset_template":        resourceTencentCloudVodSnapshotByTimeOffsetTemplate(),
			"tencentcloud_vod_super_player_config":                     resourceTencentCloudVodSuperPlayerConfig(),
			"tencentcloud_vod_sub_application":                         resourceTencentCloudVodSubApplication(),
			"tencentcloud_sqlserver_publish_subscribe":                 resourceTencentCloudSqlserverPublishSubscribe(),
			"tencentcloud_api_gateway_usage_plan":                      resourceTencentCloudAPIGatewayUsagePlan(),
			"tencentcloud_api_gateway_usage_plan_attachment":           resourceTencentCloudAPIGatewayUsagePlanAttachment(),
			"tencentcloud_api_gateway_api":                             resourceTencentCloudAPIGatewayAPI(),
			"tencentcloud_api_gateway_service":                         resourceTencentCloudAPIGatewayService(),
			"tencentcloud_api_gateway_custom_domain":                   resourceTencentCloudAPIGatewayCustomDomain(),
			"tencentcloud_api_gateway_ip_strategy":                     resourceTencentCloudAPIGatewayIPStrategy(),
			"tencentcloud_api_gateway_strategy_attachment":             resourceTencentCloudAPIGatewayStrategyAttachment(),
			"tencentcloud_api_gateway_api_key":                         resourceTencentCloudAPIGatewayAPIKey(),
			"tencentcloud_api_gateway_api_key_attachment":              resourceTencentCloudAPIGatewayAPIKeyAttachment(),
			"tencentcloud_api_gateway_service_release":                 resourceTencentCloudAPIGatewayServiceRelease(),
			"tencentcloud_api_gateway_plugin":                          resourceTencentCloudApiGatewayPlugin(),
			"tencentcloud_api_gateway_plugin_attachment":               resourceTencentCloudApiGatewayPluginAttachment(),
			"tencentcloud_sqlserver_basic_instance":                    resourceTencentCloudSqlserverBasicInstance(),
			"tencentcloud_tcr_instance":                                resourceTencentCloudTcrInstance(),
			"tencentcloud_tcr_namespace":                               resourceTencentCloudTcrNamespace(),
			"tencentcloud_tcr_repository":                              resourceTencentCloudTcrRepository(),
			"tencentcloud_tcr_token":                                   resourceTencentCloudTcrToken(),
			"tencentcloud_tcr_vpc_attachment":                          resourceTencentCloudTcrVpcAttachment(),
			"tencentcloud_tcr_tag_retention_rule":                      resourceTencentCloudTcrTagRetentionRule(),
			"tencentcloud_tcr_webhook_trigger":                         resourceTencentCloudTcrWebhookTrigger(),
			"tencentcloud_tcr_manage_replication_operation":            resourceTencentCloudTcrManageReplicationOperation(),
			"tencentcloud_tcr_customized_domain":                       resourceTencentCloudTcrCustomizedDomain(),
			"tencentcloud_tcr_immutable_tag_rule":                      resourceTencentCloudTcrImmutableTagRule(),
			"tencentcloud_tcr_delete_image_operation":                  resourceTencentCloudTcrDeleteImageOperation(),
			"tencentcloud_tcr_create_image_signature_operation":        resourceTencentCloudTcrCreateImageSignatureOperation(),
			"tencentcloud_tdmq_instance":                               resourceTencentCloudTdmqInstance(),
			"tencentcloud_tdmq_namespace":                              resourceTencentCloudTdmqNamespace(),
			"tencentcloud_tdmq_topic":                                  resourceTencentCloudTdmqTopic(),
			"tencentcloud_tdmq_role":                                   resourceTencentCloudTdmqRole(),
			"tencentcloud_tdmq_namespace_role_attachment":              resourceTencentCloudTdmqNamespaceRoleAttachment(),
			"tencentcloud_cos_bucket_policy":                           resourceTencentCloudCosBucketPolicy(),
			"tencentcloud_cos_bucket_domain_certificate_attachment":    resourceTencentCloudCosBucketDomainCertificateAttachment(),
			"tencentcloud_address_template":                            resourceTencentCloudAddressTemplate(),
			"tencentcloud_address_template_group":                      resourceTencentCloudAddressTemplateGroup(),
			"tencentcloud_protocol_template":                           resourceTencentCloudProtocolTemplate(),
			"tencentcloud_protocol_template_group":                     resourceTencentCloudProtocolTemplateGroup(),
			"tencentcloud_kms_key":                                     resourceTencentCloudKmsKey(),
			"tencentcloud_kms_external_key":                            resourceTencentCloudKmsExternalKey(),
			"tencentcloud_ssm_secret":                                  resourceTencentCloudSsmSecret(),
			"tencentcloud_ssm_secret_version":                          resourceTencentCloudSsmSecretVersion(),
			"tencentcloud_cdh_instance":                                resourceTencentCloudCdhInstance(),
			"tencentcloud_dnspod_domain_instance":                      resourceTencentCloudDnspodDomainInstance(),
			"tencentcloud_dnspod_record":                               resourceTencentCloudDnspodRecord(),
			"tencentcloud_private_dns_zone":                            resourceTencentCloudPrivateDnsZone(),
			"tencentcloud_private_dns_record":                          resourceTencentCloudPrivateDnsRecord(),
			"tencentcloud_cls_logset":                                  resourceTencentCloudClsLogset(),
			"tencentcloud_cls_topic":                                   resourceTencentCloudClsTopic(),
			"tencentcloud_cls_config":                                  resourceTencentCloudClsConfig(),
			"tencentcloud_cls_config_extra":                            resourceTencentCloudClsConfigExtra(),
			"tencentcloud_cls_config_attachment":                       resourceTencentCloudClsConfigAttachment(),
			"tencentcloud_cls_machine_group":                           resourceTencentCloudClsMachineGroup(),
			"tencentcloud_cls_cos_shipper":                             resourceTencentCloudClsCosShipper(),
			"tencentcloud_cls_index":                                   resourceTencentCloudClsIndex(),
			"tencentcloud_lighthouse_instance":                         resourceTencentCloudLighthouseInstance(),
			"tencentcloud_tem_environment":                             resourceTencentCloudTemEnvironment(),
			"tencentcloud_tem_application":                             resourceTencentCloudTemApplication(),
			"tencentcloud_tem_workload":                                resourceTencentCloudTemWorkload(),
			"tencentcloud_tem_app_config":                              resourceTencentCloudTemAppConfig(),
			"tencentcloud_tem_log_config":                              resourceTencentCloudTemLogConfig(),
			"tencentcloud_tem_scale_rule":                              resourceTencentCloudTemScaleRule(),
			"tencentcloud_tem_gateway":                                 resourceTencentCloudTemGateway(),
			"tencentcloud_tem_application_service":                     resourceTencentCloudTemApplicationService(),
			"tencentcloud_teo_zone":                                    resourceTencentCloudTeoZone(),
			"tencentcloud_teo_zone_setting":                            resourceTencentCloudTeoZoneSetting(),
			"tencentcloud_teo_dns_record":                              resourceTencentCloudTeoDnsRecord(),
			"tencentcloud_teo_dns_sec":                                 resourceTencentCloudTeoDnsSec(),
			"tencentcloud_teo_load_balancing":                          resourceTencentCloudTeoLoadBalancing(),
			"tencentcloud_teo_origin_group":                            resourceTencentCloudTeoOriginGroup(),
			"tencentcloud_teo_rule_engine":                             resourceTencentCloudTeoRuleEngine(),
			"tencentcloud_teo_rule_engine_priority":                    resourceTencentCloudTeoRuleEnginePriority(),
			"tencentcloud_teo_application_proxy":                       resourceTencentCloudTeoApplicationProxy(),
			"tencentcloud_teo_application_proxy_rule":                  resourceTencentCloudTeoApplicationProxyRule(),
			"tencentcloud_teo_ddos_policy":                             resourceTencentCloudTeoDdosPolicy(),
			"tencentcloud_teo_security_policy":                         resourceTencentCloudTeoSecurityPolicy(),
			"tencentcloud_teo_custom_error_page":                       resourceTencentCloudTeoCustomErrorPage(),
			// "tencentcloud_teo_host_certificate":                     resourceTencentCloudTeoHostCertificate(),
			// "tencentcloud_teo_default_certificate":                  resourceTencentCloudTeoDefaultCertificate(),
			"tencentcloud_tcm_mesh":                                   resourceTencentCloudTcmMesh(),
			"tencentcloud_tcm_cluster_attachment":                     resourceTencentCloudTcmClusterAttachment(),
			"tencentcloud_tcm_prometheus_attachment":                  resourceTencentCloudTcmPrometheusAttachment(),
			"tencentcloud_tcm_tracing_config":                         resourceTencentCloudTcmTracingConfig(),
			"tencentcloud_tcm_access_log_config":                      resourceTencentCloudTcmAccessLogConfig(),
			"tencentcloud_ses_domain":                                 resourceTencentCloudSesDomain(),
			"tencentcloud_ses_template":                               resourceTencentCloudSesTemplate(),
			"tencentcloud_ses_email_address":                          resourceTencentCloudSesEmailAddress(),
			"tencentcloud_sms_sign":                                   resourceTencentCloudSmsSign(),
			"tencentcloud_sms_template":                               resourceTencentCloudSmsTemplate(),
			"tencentcloud_dcdb_account":                               resourceTencentCloudDcdbAccount(),
			"tencentcloud_dcdb_hourdb_instance":                       resourceTencentCloudDcdbHourdbInstance(),
			"tencentcloud_dcdb_security_group_attachment":             resourceTencentCloudDcdbSecurityGroupAttachment(),
			"tencentcloud_dcdb_db_instance":                           resourceTencentCloudDcdbDbInstance(),
			"tencentcloud_dcdb_account_privileges":                    resourceTencentCloudDcdbAccountPrivileges(),
			"tencentcloud_dcdb_db_parameters":                         resourceTencentCloudDcdbDbParameters(),
			"tencentcloud_cat_task_set":                               resourceTencentCloudCatTaskSet(),
			"tencentcloud_mariadb_dedicatedcluster_db_instance":       resourceTencentCloudMariadbDedicatedclusterDbInstance(),
			"tencentcloud_mariadb_instance":                           resourceTencentCloudMariadbInstance(),
			"tencentcloud_mariadb_hour_db_instance":                   resourceTencentCloudMariadbHourDbInstance(),
			"tencentcloud_mariadb_account":                            resourceTencentCloudMariadbAccount(),
			"tencentcloud_mariadb_parameters":                         resourceTencentCloudMariadbParameters(),
			"tencentcloud_mariadb_log_file_retention_period":          resourceTencentCloudMariadbLogFileRetentionPeriod(),
			"tencentcloud_mariadb_security_groups":                    resourceTencentCloudMariadbSecurityGroups(),
			"tencentcloud_mariadb_encrypt_attributes":                 resourceTencentCloudMariadbEncryptAttributes(),
			"tencentcloud_tdcpg_cluster":                              resourceTencentCloudTdcpgCluster(),
			"tencentcloud_tdcpg_instance":                             resourceTencentCloudTdcpgInstance(),
			"tencentcloud_css_watermark":                              resourceTencentCloudCssWatermark(),
			"tencentcloud_css_pull_stream_task":                       resourceTencentCloudCssPullStreamTask(),
			"tencentcloud_css_live_transcode_template":                resourceTencentCloudCssLiveTranscodeTemplate(),
			"tencentcloud_css_live_transcode_rule_attachment":         resourceTencentCloudCssLiveTranscodeRuleAttachment(),
			"tencentcloud_css_domain":                                 resourceTencentCloudCssDomain(),
			"tencentcloud_css_authenticate_domain_owner_operation":    resourceTencentCloudCssAuthenticateDomainOwnerOperation(),
			"tencentcloud_css_play_domain_cert_attachment":            resourceTencentCloudCssPlayDomainCertAttachment(),
			"tencentcloud_css_play_auth_key_config":                   resourceTencentCloudCssPlayAuthKeyConfig(),
			"tencentcloud_css_push_auth_key_config":                   resourceTencentCloudCssPushAuthKeyConfig(),
			"tencentcloud_pts_project":                                resourceTencentCloudPtsProject(),
			"tencentcloud_pts_alert_channel":                          resourceTencentCloudPtsAlertChannel(),
			"tencentcloud_pts_scenario":                               resourceTencentCloudPtsScenario(),
			"tencentcloud_pts_file":                                   resourceTencentCloudPtsFile(),
			"tencentcloud_pts_job":                                    resourceTencentCloudPtsJob(),
			"tencentcloud_pts_cron_job":                               resourceTencentCloudPtsCronJob(),
			"tencentcloud_tat_command":                                resourceTencentCloudTatCommand(),
			"tencentcloud_tat_invoker":                                resourceTencentCloudTatInvoker(),
			"tencentcloud_tat_invoker_config":                         resourceTencentCloudTatInvokerConfig(),
			"tencentcloud_tat_invocation_invoke_attachment":           resourceTencentCloudTatInvocationInvokeAttachment(),
			"tencentcloud_tat_invocation_command_attachment":          resourceTencentCloudTatInvocationCommandAttachment(),
			"tencentcloud_organization_org_node":                      resourceTencentCloudOrganizationOrgNode(),
			"tencentcloud_organization_org_member":                    resourceTencentCloudOrganizationOrgMember(),
			"tencentcloud_organization_policy_sub_account_attachment": resourceTencentCloudOrganizationPolicySubAccountAttachment(),
			"tencentcloud_dbbrain_sql_filter":                         resourceTencentCloudDbbrainSqlFilter(),
			"tencentcloud_dbbrain_security_audit_log_export_task":     resourceTencentCloudDbbrainSecurityAuditLogExportTask(),
			"tencentcloud_dbbrain_db_diag_report_task":                resourceTencentCloudDbbrainDbDiagReportTask(),
			"tencentcloud_dbbrain_modify_diag_db_instance_operation":  resourceTencentCloudDbbrainModifyDiagDbInstanceOperation(),
			"tencentcloud_dbbrain_tdsql_audit_log":                    resourceTencentCloudDbbrainTdsqlAuditLog(),
			"tencentcloud_rum_project":                                resourceTencentCloudRumProject(),
			"tencentcloud_rum_taw_instance":                           resourceTencentCloudRumTawInstance(),
			"tencentcloud_rum_whitelist":                              resourceTencentCloudRumWhitelist(),
			"tencentcloud_rum_offline_log_config_attachment":          resourceTencentCloudRumOfflineLogConfigAttachment(),
			"tencentcloud_dts_sync_job":                               resourceTencentCloudDtsSyncJob(),
			"tencentcloud_tdmq_rocketmq_cluster":                      resourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloud_tdmq_rocketmq_namespace":                    resourceTencentCloudTdmqRocketmqNamespace(),
			"tencentcloud_tdmq_rocketmq_role":                         resourceTencentCloudTdmqRocketmqRole(),
			"tencentcloud_tdmq_rocketmq_topic":                        resourceTencentCloudTdmqRocketmqTopic(),
			"tencentcloud_tdmq_rocketmq_group":                        resourceTencentCloudTdmqRocketmqGroup(),
			"tencentcloud_tdmq_rocketmq_environment_role":             resourceTencentCloudTdmqRocketmqEnvironmentRole(),
			"tencentcloud_dts_migrate_service":                        resourceTencentCloudDtsMigrateService(),
			"tencentcloud_dts_migrate_job":                            resourceTencentCloudDtsMigrateJob(),
			"tencentcloud_dts_migrate_job_config":                     resourceTencentCloudDtsMigrateJobConfig(),
			"tencentcloud_dts_migrate_job_start_operation":            resourceTencentCloudDtsMigrateJobStartOperation(),
			"tencentcloud_dts_migrate_job_resume_operation":           resourceTencentCloudDtsMigrateJobResumeOperation(),
			"tencentcloud_dts_sync_check_job_operation":               resourceTencentCloudDtsSyncCheckJobOperation(),
			"tencentcloud_dts_sync_job_resume_operation":              resourceTencentCloudDtsSyncJobResumeOperation(),
			"tencentcloud_dts_compare_task_stop_operation":            resourceTencentCloudDtsCompareTaskStopOperation(),
			"tencentcloud_dts_compare_task":                           resourceTencentCloudDtsCompareTask(),
			"tencentcloud_cvm_hpc_cluster":                            resourceTencentCloudCvmHpcCluster(),
			"tencentcloud_vpc_flow_log":                               resourceTencentCloudVpcFlowLog(),
			"tencentcloud_vpc_end_point_service":                      resourceTencentCloudVpcEndPointService(),
			"tencentcloud_vpc_end_point":                              resourceTencentCloudVpcEndPoint(),
			"tencentcloud_vpc_end_point_service_white_list":           resourceTencentCloudVpcEndPointServiceWhiteList(),
			"tencentcloud_ci_bucket_attachment":                       resourceTencentCloudCiBucketAttachment(),
			"tencentcloud_tcmq_queue":                                 resourceTencentCloudTcmqQueue(),
			"tencentcloud_tcmq_topic":                                 resourceTencentCloudTcmqTopic(),
			"tencentcloud_tcmq_subscribe":                             resourceTencentCloudTcmqSubscribe(),
			"tencentcloud_ci_bucket_pic_style":                        resourceTencentCloudCiBucketPicStyle(),
			"tencentcloud_ci_hot_link":                                resourceTencentCloudCiHotLink(),
			"tencentcloud_ci_media_snapshot_template":                 resourceTencentCloudCiMediaSnapshotTemplate(),
			"tencentcloud_ci_media_transcode_template":                resourceTencentCloudCiMediaTranscodeTemplate(),
			"tencentcloud_ci_media_animation_template":                resourceTencentCloudCiMediaAnimationTemplate(),
			"tencentcloud_ci_media_concat_template":                   resourceTencentCloudCiMediaConcatTemplate(),
			"tencentcloud_ci_media_video_process_template":            resourceTencentCloudCiMediaVideoProcessTemplate(),
			"tencentcloud_ci_media_video_montage_template":            resourceTencentCloudCiMediaVideoMontageTemplate(),
			"tencentcloud_ci_media_voice_separate_template":           resourceTencentCloudCiMediaVoiceSeparateTemplate(),
			"tencentcloud_ci_media_super_resolution_template":         resourceTencentCloudCiMediaSuperResolutionTemplate(),
			"tencentcloud_ci_media_pic_process_template":              resourceTencentCloudCiMediaPicProcessTemplate(),
			"tencentcloud_ci_media_watermark_template":                resourceTencentCloudCiMediaWatermarkTemplate(),
			"tencentcloud_ci_media_tts_template":                      resourceTencentCloudCiMediaTtsTemplate(),
			"tencentcloud_ci_media_transcode_pro_template":            resourceTencentCloudCiMediaTranscodeProTemplate(),
			"tencentcloud_ci_media_smart_cover_template":              resourceTencentCloudCiMediaSmartCoverTemplate(),
			"tencentcloud_ci_media_speech_recognition_template":       resourceTencentCloudCiMediaSpeechRecognitionTemplate(),
			"tencentcloud_ci_guetzli":                                 resourceTencentCloudCIGuetzli(),
			"tencentcloud_ci_original_image_protection":               resourceTencentCloudCIOriginalImageProtection(),
			"tencentcloud_cynosdb_audit_log_file":                     resourceTencentCloudCynosdbAuditLogFile(),
			"tencentcloud_cynosdb_security_group":                     resourceTencentCloudCynosdbSecurityGroup(),
			"tencentcloud_dayu_ddos_ip_attachment_v2":                 resourceTencentCloudDayuDDosIpAttachmentV2(),
			"tencentcloud_tsf_microservice":                           resourceTencentCloudTsfMicroservice(),
			"tencentcloud_tsf_application_config":                     resourceTencentCloudTsfApplicationConfig(),
			"tencentcloud_cvm_launch_template":                        resourceTencentCloudCvmLaunchTemplate(),
			"tencentcloud_tsf_cluster":                                resourceTencentCloudTsfCluster(),
			"tencentcloud_tsf_api_group":                              resourceTencentCloudTsfApiGroup(),
			"tencentcloud_tsf_namespace":                              resourceTencentCloudTsfNamespace(),
			"tencentcloud_tsf_path_rewrite":                           resourceTencentCloudTsfPathRewrite(),
			"tencentcloud_tsf_unit_rule":                              resourceTencentCloudTsfUnitRule(),
			"tencentcloud_tsf_task":                                   resourceTencentCloudTsfTask(),
			"tencentcloud_tsf_config_template":                        resourceTencentCloudTsfConfigTemplate(),
			"tencentcloud_tsf_api_rate_limit_rule":                    resourceTencentCloudTsfApiRateLimitRule(),
			"tencentcloud_tsf_application_release_config":             resourceTencentCloudTsfApplicationReleaseConfig(),
			"tencentcloud_tsf_contain_group":                          resourceTencentCloudTsfContainGroup(),
			"tencentcloud_tsf_lane":                                   resourceTencentCloudTsfLane(),
			"tencentcloud_tsf_lane_rule":                              resourceTencentCloudTsfLaneRule(),
			"tencentcloud_tsf_group":                                  resourceTencentCloudTsfGroup(),
			"tencentcloud_tsf_repository":                             resourceTencentCloudTsfRepository(),
			"tencentcloud_tsf_application":                            resourceTencentCloudTsfApplication(),
			"tencentcloud_tsf_application_public_config_release":      resourceTencentCloudTsfApplicationPublicConfigRelease(),
			"tencentcloud_tsf_application_public_config":              resourceTencentCloudTsfApplicationPublicConfig(),
			"tencentcloud_tsf_application_file_config_release":        resourceTencentCloudTsfApplicationFileConfigRelease(),
			"tencentcloud_tsf_instances_attachment":                   resourceTencentCloudTsfInstancesAttachment(),
			"tencentcloud_tsf_bind_api_group":                         resourceTencentCloudTsfBindApiGroup(),
			"tencentcloud_mps_workflow":                               resourceTencentCloudMpsWorkflow(),
			"tencentcloud_mps_transcode_template":                     resourceTencentCloudMpsTranscodeTemplate(),
			"tencentcloud_mps_watermark_template":                     resourceTencentCloudMpsWatermarkTemplate(),
			"tencentcloud_mps_image_sprite_template":                  resourceTencentCloudMpsImageSpriteTemplate(),
			"tencentcloud_mps_snapshot_by_timeoffset_template":        resourceTencentCloudMpsSnapshotByTimeoffsetTemplate(),
			"tencentcloud_mps_sample_snapshot_template":               resourceTencentCloudMpsSampleSnapshotTemplate(),
			"tencentcloud_mps_animated_graphics_template":             resourceTencentCloudMpsAnimatedGraphicsTemplate(),
			"tencentcloud_mps_ai_recognition_template":                resourceTencentCloudMpsAiRecognitionTemplate(),
			"tencentcloud_mps_ai_analysis_template":                   resourceTencentCloudMpsAiAnalysisTemplate(),
			"tencentcloud_mps_adaptive_dynamic_streaming_template":    resourceTencentCloudMpsAdaptiveDynamicStreamingTemplate(),
			"tencentcloud_mps_person_sample":                          resourceTencentCloudMpsPersonSample(),
			"tencentcloud_cbs_disk_backup":                            resourceTencentCloudCbsDiskBackup(),
			"tencentcloud_cbs_snapshot_share_permission":              resourceTencentCloudCbsSnapshotSharePermission(),
			"tencentcloud_cbs_disk_backup_rollback_operation":         resourceTencentCloudCbsDiskBackupRollbackOperation(),
			"tencentcloud_chdfs_access_group":                         resourceTencentCloudChdfsAccessGroup(),
			"tencentcloud_chdfs_access_rule":                          resourceTencentCloudChdfsAccessRule(),
			"tencentcloud_chdfs_file_system":                          resourceTencentCloudChdfsFileSystem(),
			"tencentcloud_chdfs_life_cycle_rule":                      resourceTencentCloudChdfsLifeCycleRule(),
			"tencentcloud_chdfs_mount_point":                          resourceTencentCloudChdfsMountPoint(),
			"tencentcloud_chdfs_mount_point_attachment":               resourceTencentCloudChdfsMountPointAttachment(),
			"tencentcloud_mdl_stream_live_input":                      resourceTencentCloudMdlStreamLiveInput(),
			"tencentcloud_lighthouse_blueprint":                       resourceTencentCloudLighthouseBlueprint(),
			"tencentcloud_cvm_launch_template_version":                resourceTencentCloudCvmLaunchTemplateVersion(),
			"tencentcloud_apm_instance":                               resourceTencentCloudApmInstance(),
			"tencentcloud_cvm_launch_template_default_version":        resourceTencentCloudCvmLaunchTemplateDefaultVersion(),
			"tencentcloud_lighthouse_firewall_rule":                   resourceTencentCloudLighthouseFirewallRule(),
			"tencentcloud_cvm_security_group_attachment":              resourceTencentCloudCvmSecurityGroupAttachment(),
			"tencentcloud_cvm_reboot_instance":                        resourceTencentCloudCvmRebootInstance(),
			"tencentcloud_cvm_chc_config":                             resourceTencentCloudCvmChcConfig(),
			"tencentcloud_lighthouse_disk_backup":                     resourceTencentCloudLighthouseDiskBackup(),
			"tencentcloud_lighthouse_apply_disk_backup":               resourceTencentCloudLighthouseApplyDiskBackup(),
			"tencentcloud_lighthouse_disk_attachment":                 resourceTencentCloudLighthouseDiskAttachment(),
			"tencentcloud_lighthouse_key_pair":                        resourceTencentCloudLighthouseKeyPair(),
			"tencentcloud_lighthouse_snapshot":                        resourceTencentCloudLighthouseSnapshot(),
			"tencentcloud_lighthouse_apply_instance_snapshot":         resourceTencentCloudLighthouseApplyInstanceSnapshot(),
			"tencentcloud_lighthouse_start_instance":                  resourceTencentCloudLighthouseStartInstance(),
			"tencentcloud_lighthouse_stop_instance":                   resourceTencentCloudLighthouseStopInstance(),
			"tencentcloud_lighthouse_reboot_instance":                 resourceTencentCloudLighthouseRebootInstance(),
			"tencentcloud_lighthouse_key_pair_attachment":             resourceTencentCloudLighthouseKeyPairAttachment(),
			"tencentcloud_lighthouse_disk":                            resourceTencentCloudLighthouseDisk(),
			"tencentcloud_api_gateway_api_doc":                        resourceTencentCloudAPIGatewayAPIDoc(),
			"tencentcloud_api_gateway_api_app":                        resourceTencentCloudAPIGatewayAPIApp(),
			"tencentcloud_tse_instance":                               resourceTencentCloudTseInstance(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	secretId := d.Get("secret_id").(string)
	secretKey := d.Get("secret_key").(string)
	securityToken := d.Get("security_token").(string)
	region := d.Get("region").(string)
	protocol := d.Get("protocol").(string)
	domain := d.Get("domain").(string)

	// standard client
	var tcClient TencentCloudClient
	tcClient.apiV3Conn = &connectivity.TencentCloudClient{
		Credential: common.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		Region:   region,
		Protocol: protocol,
		Domain:   domain,
	}

	envRoleArn := os.Getenv(PROVIDER_ASSUME_ROLE_ARN)
	envSessionName := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_NAME)

	// get assume role from env
	if envRoleArn != "" && envSessionName != "" {
		var assumeRoleSessionDuration int
		if envSessionDuration := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); envSessionDuration != "" {
			var err error
			assumeRoleSessionDuration, err = strconv.Atoi(envSessionDuration)
			if err != nil {
				return nil, err
			}
		}
		if assumeRoleSessionDuration == 0 {
			assumeRoleSessionDuration = 7200
		}

		_ = genClientWithSTS(&tcClient, envRoleArn, envSessionName, assumeRoleSessionDuration, "")
	}

	// get assume role from tf config
	assumeRoleList := d.Get("assume_role").(*schema.Set).List()
	if len(assumeRoleList) == 1 {
		assumeRole := assumeRoleList[0].(map[string]interface{})
		assumeRoleArn := assumeRole["role_arn"].(string)
		assumeRoleSessionName := assumeRole["session_name"].(string)
		assumeRoleSessionDuration := assumeRole["session_duration"].(int)
		assumeRolePolicy := assumeRole["policy"].(string)

		_ = genClientWithSTS(&tcClient, assumeRoleArn, assumeRoleSessionName, assumeRoleSessionDuration, assumeRolePolicy)
	}
	return &tcClient, nil
}

func genClientWithSTS(tcClient *TencentCloudClient, assumeRoleArn, assumeRoleSessionName string, assumeRoleSessionDuration int, assumeRolePolicy string) error {
	// applying STS credentials
	request := sts.NewAssumeRoleRequest()
	request.RoleArn = helper.String(assumeRoleArn)
	request.RoleSessionName = helper.String(assumeRoleSessionName)
	request.DurationSeconds = helper.IntUint64(assumeRoleSessionDuration)
	if assumeRolePolicy != "" {
		request.Policy = helper.String(url.QueryEscape(assumeRolePolicy))
	}
	ratelimit.Check(request.GetAction())
	response, err := tcClient.apiV3Conn.UseStsClient().AssumeRole(request)
	if err != nil {
		return err
	}
	// using STS credentials
	tcClient.apiV3Conn.Credential = common.NewTokenCredential(
		*response.Response.Credentials.TmpSecretId,
		*response.Response.Credentials.TmpSecretKey,
		*response.Response.Credentials.Token,
	)
	return nil
}

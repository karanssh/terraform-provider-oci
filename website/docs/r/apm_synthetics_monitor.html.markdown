---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_monitor"
sidebar_current: "docs-oci-resource-apm_synthetics-monitor"
description: |-
  Provides the Monitor resource in Oracle Cloud Infrastructure Apm Synthetics service
---

# oci_apm_synthetics_monitor
This resource provides the Monitor resource in Oracle Cloud Infrastructure Apm Synthetics service.

Creates a new monitor.


## Example Usage

```hcl
resource "oci_apm_synthetics_monitor" "test_monitor" {
	#Required
	apm_domain_id = oci_apm_synthetics_apm_domain.test_apm_domain.id
	display_name = var.monitor_display_name
	monitor_type = var.monitor_monitor_type
	repeat_interval_in_seconds = var.monitor_repeat_interval_in_seconds
        vantage_points {
          #Required
          name  = var.monitor_vantage_points_name
          #Optional
          display_name = var.monitor_vantage_points_param_display_name
        }

	#Optional
	availability_configuration {

		#Optional
		max_allowed_failures_per_interval = var.monitor_availability_configuration_max_allowed_failures_per_interval
		min_allowed_runs_per_interval = var.monitor_availability_configuration_min_allowed_runs_per_interval
	}
	batch_interval_in_seconds = var.monitor_batch_interval_in_seconds
	configuration {

		#Optional
		client_certificate_details {

			#Optional
			client_certificate {

				#Optional
				content = var.monitor_configuration_client_certificate_details_client_certificate_content
				file_name = var.monitor_configuration_client_certificate_details_client_certificate_file_name
			}
			private_key {

				#Optional
				content = var.monitor_configuration_client_certificate_details_private_key_content
				file_name = var.monitor_configuration_client_certificate_details_private_key_file_name
			}
		}
		config_type = var.monitor_configuration_config_type
		dns_configuration {

			#Optional
			is_override_dns = var.monitor_configuration_dns_configuration_is_override_dns
			override_dns_ip = var.monitor_configuration_dns_configuration_override_dns_ip
		}
		is_certificate_validation_enabled = var.monitor_configuration_is_certificate_validation_enabled
		is_default_snapshot_enabled = var.monitor_configuration_is_default_snapshot_enabled
		is_failure_retried = var.monitor_configuration_is_failure_retried
		is_redirection_enabled = var.monitor_configuration_is_redirection_enabled
		network_configuration {

			#Optional
			number_of_hops = var.monitor_configuration_network_configuration_number_of_hops
			probe_mode = var.monitor_configuration_network_configuration_probe_mode
			probe_per_hop = var.monitor_configuration_network_configuration_probe_per_hop
			protocol = var.monitor_configuration_network_configuration_protocol
			transmission_rate = var.monitor_configuration_network_configuration_transmission_rate
		}
		req_authentication_details {

			#Optional
			auth_headers {

				#Optional
				header_name = var.monitor_configuration_req_authentication_details_auth_headers_header_name
				header_value = var.monitor_configuration_req_authentication_details_auth_headers_header_value
			}
			auth_request_method = var.monitor_configuration_req_authentication_details_auth_request_method
			auth_request_post_body = var.monitor_configuration_req_authentication_details_auth_request_post_body
			auth_token = var.monitor_configuration_req_authentication_details_auth_token
			auth_url = var.monitor_configuration_req_authentication_details_auth_url
			auth_user_name = oci_identity_user.test_user.name
			auth_user_password = var.monitor_configuration_req_authentication_details_auth_user_password
			oauth_scheme = var.monitor_configuration_req_authentication_details_oauth_scheme
		}
		req_authentication_scheme = var.monitor_configuration_req_authentication_scheme
		request_headers {

			#Optional
			header_name = var.monitor_configuration_request_headers_header_name
			header_value = var.monitor_configuration_request_headers_header_value
		}
		request_method = var.monitor_configuration_request_method
		request_post_body = var.monitor_configuration_request_post_body
		request_query_params {

			#Optional
			param_name = var.monitor_configuration_request_query_params_param_name
			param_value = var.monitor_configuration_request_query_params_param_value
		}
		verify_response_codes = var.monitor_configuration_verify_response_codes
		verify_response_content = var.monitor_configuration_verify_response_content
		verify_texts {

			#Optional
			text = var.monitor_configuration_verify_texts_text
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	is_run_now = var.monitor_is_run_now
	is_run_once = var.monitor_is_run_once
	maintenance_window_schedule {

		#Optional
		time_ended = var.monitor_maintenance_window_schedule_time_ended
		time_started = var.monitor_maintenance_window_schedule_time_started
	}
	scheduling_policy = var.monitor_scheduling_policy
	script_id = oci_apm_synthetics_script.test_script.id
	script_parameters {
		#Required
		param_name = var.monitor_script_parameters_param_name
		param_value = var.monitor_script_parameters_param_value
	}
	status = var.monitor_status
	target = var.monitor_target
	timeout_in_seconds = var.monitor_timeout_in_seconds
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) (Updatable) The APM domain ID the request is intended for.
* `availability_configuration` - (Optional) (Updatable) Monitor availability configuration details.
	* `max_allowed_failures_per_interval` - (Optional) (Updatable) Intervals with failed runs more than this value will be classified as UNAVAILABLE.
	* `min_allowed_runs_per_interval` - (Optional) (Updatable) Intervals with runs less than this value will be classified as UNKNOWN and excluded from the availability calculations.

* `batch_interval_in_seconds` - (Optional) (Updatable) Time interval between 2 runs in round robin batch mode (*SchedulingPolicy - BATCHED_ROUND_ROBIN).
* `display_name` - (Required) (Updatable) Unique name that can be edited. The name should not contain any confidential information.
* `monitor_type` - (Required) Type of monitor.
* `repeat_interval_in_seconds` - (Required) (Updatable) Interval in seconds after the start time when the job should be repeated. Minimum repeatIntervalInSeconds should be 300 seconds.
* `vantage_points` - (Required) (Updatable) A list of vantage points from which to execute the monitor. Use /publicVantagePoints to fetch public vantage points.
	* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
	* `name` - Name of the vantage point.
* `configuration` - (Optional) (Updatable) Details of monitor configuration.
	* `client_certificate_details` - (Applicable when config_type=REST_CONFIG) (Updatable) Details for client certificate.
		* `client_certificate` - (Applicable when config_type=REST_CONFIG) (Updatable) Client certificate in PEM format.
			* `content` - (Required when config_type=REST_CONFIG) (Updatable) Content of the client certificate file.
			* `file_name` - (Required when config_type=REST_CONFIG) (Updatable) Name of the certificate file. The name should not contain any confidential information.
		* `private_key` - (Applicable when config_type=REST_CONFIG) (Updatable) The private key associated with the client certificate in PEM format.
			* `content` - (Required when config_type=REST_CONFIG) (Updatable) Content of the private key file.
			* `file_name` - (Required when config_type=REST_CONFIG) (Updatable) Name of the private key file.
	* `config_type` - (Optional) (Updatable) Type of configuration.
	* `dns_configuration` - (Optional) (Updatable) Information about the DNS settings.
		* `is_override_dns` - (Optional) (Updatable) If isOverrideDns is true, then DNS settings will be overridden.
		* `override_dns_ip` - (Optional) (Updatable) Attribute to override the DNS IP value. This value will be honored only if isOverrideDns is set to true.
	* `is_certificate_validation_enabled` - (Applicable when config_type=BROWSER_CONFIG | REST_CONFIG | SCRIPTED_BROWSER_CONFIG) (Updatable) If certificate validation is enabled, then the call will fail in case of certification errors.
	* `is_default_snapshot_enabled` - (Applicable when config_type=BROWSER_CONFIG | SCRIPTED_BROWSER_CONFIG) (Updatable) If disabled, auto snapshots are not collected.
	* `is_failure_retried` - (Optional) (Updatable) If isFailureRetried is enabled, then a failed call will be retried.
	* `is_redirection_enabled` - (Applicable when config_type=REST_CONFIG) (Updatable) If redirection is enabled, then redirects will be allowed while accessing target URL.
	* `network_configuration` - (Optional) (Updatable) Details of the network configuration.
		* `number_of_hops` - (Applicable when config_type=BROWSER_CONFIG | NETWORK_CONFIG | REST_CONFIG | SCRIPTED_BROWSER_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Number of hops.
		* `probe_mode` - (Applicable when config_type=BROWSER_CONFIG | NETWORK_CONFIG | REST_CONFIG | SCRIPTED_BROWSER_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Type of probe mode when TCP protocol is selected.
		* `probe_per_hop` - (Applicable when config_type=BROWSER_CONFIG | NETWORK_CONFIG | REST_CONFIG | SCRIPTED_BROWSER_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Number of probes per hop.
		* `protocol` - (Applicable when config_type=BROWSER_CONFIG | NETWORK_CONFIG | REST_CONFIG | SCRIPTED_BROWSER_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Type of protocol.
		* `transmission_rate` - (Applicable when config_type=BROWSER_CONFIG | NETWORK_CONFIG | REST_CONFIG | SCRIPTED_BROWSER_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Number of probe packets sent out simultaneously.
	* `req_authentication_details` - (Applicable when config_type=REST_CONFIG) (Updatable) Details for request HTTP authentication.
		* `auth_headers` - (Applicable when config_type=REST_CONFIG) (Updatable) List of authentication headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
			* `header_name` - (Required when config_type=REST_CONFIG) (Updatable) Name of the header.
			* `header_value` - (Applicable when config_type=REST_CONFIG) (Updatable) Value of the header.
		* `auth_request_method` - (Applicable when config_type=REST_CONFIG) (Updatable) Request method.
		* `auth_request_post_body` - (Applicable when config_type=REST_CONFIG) (Updatable) Request post body.
		* `auth_token` - (Applicable when config_type=REST_CONFIG) (Updatable) Authentication token.
		* `auth_url` - (Applicable when config_type=REST_CONFIG) (Updatable) URL to get authentication token.
		* `auth_user_name` - (Applicable when config_type=REST_CONFIG) (Updatable) User name for authentication.
		* `auth_user_password` - (Applicable when config_type=REST_CONFIG) (Updatable) User password for authentication.
		* `oauth_scheme` - (Applicable when config_type=REST_CONFIG) (Updatable) Request HTTP OAuth scheme.
	* `req_authentication_scheme` - (Applicable when config_type=REST_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Request HTTP authentication scheme.
	* `request_headers` - (Applicable when config_type=REST_CONFIG) (Updatable) List of request headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
		* `header_name` - (Required when config_type=REST_CONFIG) (Updatable) Name of the header.
		* `header_value` - (Applicable when config_type=REST_CONFIG) (Updatable) Value of the header.
	* `request_method` - (Applicable when config_type=REST_CONFIG) (Updatable) Request HTTP method.
	* `request_post_body` - (Applicable when config_type=REST_CONFIG) (Updatable) Request post body content.
	* `request_query_params` - (Applicable when config_type=REST_CONFIG) (Updatable) List of request query params. Example: `[{"paramName": "sortOrder", "paramValue": "asc"}]` 
		* `param_name` - (Required when config_type=REST_CONFIG) (Updatable) Name of request query parameter.
		* `param_value` - (Applicable when config_type=REST_CONFIG) (Updatable) Value of request query parameter.
	* `verify_response_codes` - (Applicable when config_type=BROWSER_CONFIG | REST_CONFIG | SCRIPTED_REST_CONFIG) (Updatable) Expected HTTP response codes. For status code range, set values such as 2xx, 3xx. 
	* `verify_response_content` - (Applicable when config_type=REST_CONFIG) (Updatable) Verify response content against regular expression based string. If response content does not match the verifyResponseContent value, then it will be considered a failure. 
	* `verify_texts` - (Applicable when config_type=BROWSER_CONFIG) (Updatable) Verifies all the search strings present in the response. If any search string is not present in the response, then it will be considered as a failure. 
		* `text` - (Applicable when config_type=BROWSER_CONFIG) (Updatable) Verification text in the response.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_run_now` - (Optional) (Updatable) If isRunNow is enabled, then the monitor will run immediately.
* `is_run_once` - (Optional) (Updatable) If runOnce is enabled, then the monitor will run once.
* `maintenance_window_schedule` - (Optional) (Updatable) Details required to schedule maintenance window.
	* `time_ended` - (Optional) (Updatable) End time of the maintenance window, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
	* `time_started` - (Optional) (Updatable) Start time of the maintenance window, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `monitor_type` - (Required) Type of monitor.
* `repeat_interval_in_seconds` - (Required) (Updatable) Interval in seconds after the start time when the job should be repeated. Minimum repeatIntervalInSeconds should be 300 seconds for Scripted REST, Scripted Browser and Browser monitors, and 60 seconds for REST monitor. 
* `scheduling_policy` - (Optional) (Updatable) Scheduling policy to decide the distribution of monitor executions on vantage points.
* `script_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the script. scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null. 
* `script_parameters` - (Optional) (Updatable) List of script parameters in the monitor. This is valid only for SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null. Example: `[{"paramName": "userid", "paramValue":"testuser"}]` 
	* `param_name` - (Required) (Updatable) Name of the parameter.
	* `param_value` - (Required) (Updatable) Value of the parameter.
* `status` - (Optional) (Updatable) Enables or disables the monitor.
* `target` - (Optional) (Updatable) Specify the endpoint on which to run the monitor. For BROWSER and REST monitor types, target is mandatory. If target is specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script (specified by scriptId in monitor) against the specified target endpoint. If target is not specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script as it is. For NETWORK monitor with TCP protocol, a port needs to be provided along with target. Example: 192.168.0.1:80 
* `timeout_in_seconds` - (Optional) (Updatable) Timeout in seconds. If isFailureRetried is true, then timeout cannot be more than 30% of repeatIntervalInSeconds time for monitors. If isFailureRetried is false, then timeout cannot be more than 50% of repeatIntervalInSeconds time for monitors. Also, timeoutInSeconds should be a multiple of 60 for Scripted REST, Scripted Browser and Browser monitors. Monitor will be allowed to run only for timeoutInSeconds time. It would be terminated after that. 
* `vantage_points` - (Required) (Updatable) A list of public and dedicated vantage points from which to execute the monitor. Use /publicVantagePoints to fetch public vantage points, and /dedicatedVantagePoints to fetch dedicated vantage points. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_configuration` - Monitor availability configuration details.
	* `max_allowed_failures_per_interval` - Maximum number of failed runs allowed in an interval. If an interval has more failed runs than the specified value, then the interval will be classified as UNAVAILABLE.
	* `min_allowed_runs_per_interval` - Minimum number of runs allowed in an interval. If an interval has fewer runs than the specified value, then the interval will be classified as UNKNOWN and will be excluded from the availability calculations.
* `batch_interval_in_seconds` - Time interval between two runs in round robin batch mode (SchedulingPolicy - BATCHED_ROUND_ROBIN).
* `configuration` - Details of monitor configuration.
	* `client_certificate_details` - Details for client certificate.
		* `client_certificate` - Client certificate in PEM format.
			* `content` - Content of the client certificate file.
			* `file_name` - Name of the certificate file. The name should not contain any confidential information.
		* `private_key` - The private key associated with the client certificate in PEM format.
			* `content` - Content of the private key file.
			* `file_name` - Name of the private key file.
	* `config_type` - Type of configuration.
	* `dns_configuration` - Information about the DNS settings.
		* `is_override_dns` - If isOverrideDns is true, then DNS settings will be overridden.
		* `override_dns_ip` - Attribute to override the DNS IP value. This value will be honored only if isOverrideDns is set to true.
	* `is_certificate_validation_enabled` - If certificate validation is enabled, then the call will fail in case of certification errors.
	* `is_default_snapshot_enabled` - If disabled, auto snapshots are not collected.
	* `is_failure_retried` - If isFailureRetried is enabled, then a failed call will be retried.
	* `is_redirection_enabled` - If redirection is enabled, then redirects will be allowed while accessing target URL.
	* `network_configuration` - Details of the network configuration.
		* `number_of_hops` - Number of hops.
		* `probe_mode` - Type of probe mode when TCP protocol is selected.
		* `probe_per_hop` - Number of probes per hop.
		* `protocol` - Type of protocol.
		* `transmission_rate` - Number of probe packets sent out simultaneously.
	* `req_authentication_details` - Details for request HTTP authentication.
		* `auth_headers` - List of authentication headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
			* `header_name` - Name of the header.
			* `header_value` - Value of the header.
		* `auth_request_method` - Request method.
		* `auth_request_post_body` - Request post body.
		* `auth_token` - Authentication token.
		* `auth_url` - URL to get authentication token.
		* `auth_user_name` - User name for authentication.
		* `auth_user_password` - User password for authentication.
		* `oauth_scheme` - Request HTTP OAuth scheme.
	* `req_authentication_scheme` - Request HTTP authentication scheme.
	* `request_headers` - List of request headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]` 
		* `header_name` - Name of the header.
		* `header_value` - Value of the header.
	* `request_method` - Request HTTP method.
	* `request_post_body` - Request post body content.
	* `request_query_params` - List of request query params. Example: `[{"paramName": "sortOrder", "paramValue": "asc"}]` 
		* `param_name` - Name of request query parameter.
		* `param_value` - Value of request query parameter.
	* `verify_response_codes` - Expected HTTP response codes. For status code range, set values such as 2xx, 3xx. 
	* `verify_response_content` - Verify response content against regular expression based string. If response content does not match the verifyResponseContent value, then it will be considered a failure. 
	* `verify_texts` - Verifies all the search strings present in the response. If any search string is not present in the response, then it will be considered as a failure. 
		* `text` - Verification text in the response.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitor.
* `is_run_now` - If isRunNow is enabled, then the monitor will run immediately.
* `is_run_once` - If runOnce is enabled, then the monitor will run once.
* `maintenance_window_schedule` - Details required to schedule maintenance window.
	* `time_ended` - End time of the maintenance window, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
	* `time_started` - Start time of the maintenance window, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `monitor_type` - Type of monitor.
* `repeat_interval_in_seconds` - Interval in seconds after the start time when the job should be repeated. Minimum repeatIntervalInSeconds should be 300 seconds for Scripted REST, Scripted Browser and Browser monitors, and 60 seconds for REST monitor. 
* `scheduling_policy` - Scheduling policy to decide the distribution of monitor executions on vantage points.
* `script_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the script. scriptId is mandatory for creation of SCRIPTED_BROWSER and SCRIPTED_REST monitor types. For other monitor types, it should be set to null. 
* `script_name` - Name of the script.
* `script_parameters` - List of script parameters. Example: `[{"monitorScriptParameter": {"paramName": "userid", "paramValue":"testuser"}, "isSecret": false, "isOverwritten": false}]` 
	* `is_overwritten` - If parameter value is default or overwritten. 
	* `is_secret` - Describes if  the parameter value is secret and should be kept confidential. isSecret is specified in either CreateScript or UpdateScript API. 
	* `monitor_script_parameter` - Details of the script parameter that can be used to overwrite the parameter present in the script. 
		* `param_name` - Name of the parameter.
		* `param_value` - Value of the parameter.
* `status` - Enables or disables the monitor.
* `target` - Specify the endpoint on which to run the monitor. For BROWSER and REST monitor types, target is mandatory. If target is specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script (specified by scriptId in monitor) against the specified target endpoint. If target is not specified in the SCRIPTED_BROWSER monitor type, then the monitor will run the selected script as it is. For NETWORK monitor with TCP protocol, a port needs to be provided along with target. Example: 192.168.0.1:80 
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `timeout_in_seconds` - Timeout in seconds. If isFailureRetried is true, then timeout cannot be more than 30% of repeatIntervalInSeconds time for monitors. If isFailureRetried is false, then timeout cannot be more than 50% of repeatIntervalInSeconds time for monitors. Also, timeoutInSeconds should be a multiple of 60 for Scripted REST, Scripted Browser and Browser monitors. Monitor will be allowed to run only for timeoutInSeconds time. It would be terminated after that. 
* `vantage_point_count` - Number of vantage points where monitor is running.
* `vantage_points` - List of public and dedicated vantage points where the monitor is running.
	* `display_name` - Unique name that can be edited. The name should not contain any confidential information.
	* `name` - Name of the vantage point.
	
## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitor
	* `update` - (Defaults to 20 minutes), when updating the Monitor
	* `delete` - (Defaults to 20 minutes), when destroying the Monitor


## Import

Monitors can be imported using the `id`, e.g.

```
$ terraform import oci_apm_synthetics_monitor.test_monitor "monitors/{monitorId}/apmDomainId/{apmDomainId}" 
```


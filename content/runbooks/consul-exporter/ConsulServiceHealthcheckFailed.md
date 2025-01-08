---
title: ConsulServiceHealthcheckFailed
description: Troubleshooting for alert ConsulServiceHealthcheckFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ConsulServiceHealthcheckFailed

Service: `{{ $labels.service_name }}` Healthcheck: `{{ $labels.service_id }}`

<details>
  <summary>Alert Rule</summary>

{{% rule "consul/consul-exporter.yml" "ConsulServiceHealthcheckFailed" %}}

{{% comment %}}

```yaml
alert: ConsulServiceHealthcheckFailed
expr: consul_catalog_service_node_healthy == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: Consul service healthcheck failed (instance {{ $labels.instance }})
    description: |-
        Service: `{{ $labels.service_name }}` Healthcheck: `{{ $labels.service_id }}`
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/consul-exporter/ConsulServiceHealthcheckFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the ConsulServiceHealthcheckFailed alert:

## Meaning

The ConsulServiceHealthcheckFailed alert indicates that a Consul service health check has failed. This means that the service is not responding to health checks, which can indicate a problem with the service or the underlying infrastructure.

## Impact

The impact of this alert can be severe, as it may indicate that a critical service is not functioning properly. This can lead to issues with application availability, data loss, or other downstream effects. The failed health check may also indicate a larger issue with the Consul cluster or the underlying infrastructure.

## Diagnosis

To diagnose the root cause of the ConsulServiceHealthcheckFailed alert, follow these steps:

1. Check the Consul UI to verify the status of the service and the health check.
2. Check the service logs to determine if there are any error messages or exceptions that may indicate the cause of the failure.
3. Verify that the service is properly configured and that the health check is correctly set up.
4. Check the underlying infrastructure, such as the network, CPU, and memory usage, to ensure that it is functioning properly.
5. Review the Consul cluster status to ensure that it is healthy and functioning as expected.

## Mitigation

To mitigate the ConsulServiceHealthcheckFailed alert, follow these steps:

1. Restart the service to see if it self-heals.
2. Check the service configuration and verify that it is correct.
3. Update the service or Consul configuration to fix any issues identified during diagnosis.
4. Check the underlying infrastructure and perform any necessary maintenance or repairs.
5. Consider increasing the monitoring and logging for the service to provide more visibility into its operation.
6. If the issue persists, consider escalating to a higher-level support team or seeking guidance from a subject matter expert.

By following these steps, you should be able to diagnose and mitigate the root cause of the ConsulServiceHealthcheckFailed alert, and restore the service to a healthy state.
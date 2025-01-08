---
title: RabbitmqInstancesDifferentVersions
description: Troubleshooting for alert RabbitmqInstancesDifferentVersions
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# RabbitmqInstancesDifferentVersions

Running different version of RabbitMQ in the same cluster, can lead to failure.

<details>
  <summary>Alert Rule</summary>

{{% rule "rabbitmq/rabbitmq-exporter.yml" "RabbitmqInstancesDifferentVersions" %}}

{{% comment %}}

```yaml
alert: RabbitmqInstancesDifferentVersions
expr: count(count(rabbitmq_build_info) by (rabbitmq_version)) > 1
for: 1h
labels:
    severity: warning
annotations:
    summary: RabbitMQ instances different versions (instance {{ $labels.instance }})
    description: |-
        Running different version of RabbitMQ in the same cluster, can lead to failure.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/rabbitmq-exporter/RabbitmqInstancesDifferentVersions.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule "RabbitmqInstancesDifferentVersions":

## Meaning

This alert is triggered when multiple RabbitMQ instances in the same cluster are running different versions of RabbitMQ. This is detected by the presence of multiple distinct values for the `rabbitmq_version` label in the `rabbitmq_build_info` metric.

## Impact

Running different versions of RabbitMQ in the same cluster can lead to failures, inconsistencies, and unpredictable behavior. This may cause issues with message processing, queuing, and overall system reliability.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the `rabbitmq_build_info` metric to identify the different versions of RabbitMQ running in the cluster.
2. Verify the versions of RabbitMQ installed on each instance.
3. Check the RabbitMQ cluster configuration to ensure that all instances are properly connected and configured.
4. Review the system logs for any errors or warnings related to version inconsistencies.

## Mitigation

To mitigate the issue, follow these steps:

1. Upgrade all RabbitMQ instances to the same version.
2. Ensure that the RabbitMQ cluster configuration is correct and consistent across all instances.
3. Validate that all instances are properly connected and communicating with each other.
4. Monitor the cluster for any further issues or inconsistencies.

Additional resources:

* Refer to the RabbitMQ documentation for version upgrade and cluster configuration guidelines.
* Consult with the RabbitMQ cluster administrator or operation team for assistance with the mitigation process.
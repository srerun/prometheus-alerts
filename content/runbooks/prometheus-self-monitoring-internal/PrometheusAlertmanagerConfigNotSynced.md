---
title: PrometheusAlertmanagerConfigNotSynced
description: Troubleshooting for alert PrometheusAlertmanagerConfigNotSynced
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# PrometheusAlertmanagerConfigNotSynced

Configurations of AlertManager cluster instances are out of sync

<details>
  <summary>Alert Rule</summary>

{{% rule "prometheus-self-monitoring/prometheus-self-monitoring-internal.yml" "PrometheusAlertmanagerConfigNotSynced" %}}

{{% comment %}}

```yaml
alert: PrometheusAlertmanagerConfigNotSynced
expr: count(count_values("config_hash", alertmanager_config_hash)) > 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Prometheus AlertManager config not synced (instance {{ $labels.instance }})
    description: |-
        Configurations of AlertManager cluster instances are out of sync
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/prometheus-self-monitoring-internal/PrometheusAlertmanagerConfigNotSynced.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus AlertmanagerConfigNotSynced alert:

## Meaning

The Prometheus AlertmanagerConfigNotSynced alert is triggered when the configuration of AlertManager cluster instances are not synchronized. This means that the configuration hash of the AlertManager instances does not match, indicating that there is a discrepancy in the configuration across the cluster.

## Impact

The impact of this alert is that AlertManager may not function correctly, leading to potential issues with alert routing, notification, and escalation. This can result in delayed or missed alerts, which can have significant consequences in terms of service availability, performance, and security.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the AlertManager configuration files on each instance to identify any differences.
2. Verify that the configuration has been updated correctly on each instance.
3. Check the AlertManager logs for any errors or warnings related to configuration synchronization.
4. Review the Prometheus metrics and labels to identify which instances are affected.

## Mitigation

To mitigate the issue, follow these steps:

1. Identify the instance(s) with the mismatched configuration and update the configuration to match the others.
2. Verify that the configuration has been updated correctly on each instance.
3. Restart the AlertManager service on the affected instances to apply the updated configuration.
4. Monitor the AlertManager logs and Prometheus metrics to ensure that the configuration is synchronized across all instances.
5. Consider implementing a configuration management system to ensure consistency across the cluster.

Remember to also review and update the runbook to reflect any changes to your AlertManager configuration or deployment.
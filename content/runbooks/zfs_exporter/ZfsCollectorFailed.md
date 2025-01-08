---
title: ZfsCollectorFailed
description: Troubleshooting for alert ZfsCollectorFailed
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ZfsCollectorFailed

ZFS collector for {{ $labels.instance }} has failed to collect information

<details>
  <summary>Alert Rule</summary>

{{% rule "zfs/zfs_exporter.yml" "ZfsCollectorFailed" %}}

{{% comment %}}

```yaml
alert: ZfsCollectorFailed
expr: zfs_scrape_collector_success != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: ZFS collector failed (instance {{ $labels.instance }})
    description: |-
        ZFS collector for {{ $labels.instance }} has failed to collect information
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/zfs_exporter/ZfsCollectorFailed.md

```

{{% /comment %}}

</details>


Here is a runbook for the ZfsCollectorFailed alert rule:

## Meaning

The ZfsCollectorFailed alert is triggered when the ZFS collector fails to collect information from a ZFS instance. This means that Prometheus is unable to scrape metrics from the ZFS exporter, which can lead to incomplete or missing data in monitoring dashboards and potential issues with alerting and automation.

## Impact

The impact of this alert is moderate. Without ZFS metrics, monitoring and alerting capabilities may be impaired, leading to potential issues with data integrity, capacity planning, and performance troubleshooting. Additionally, automated workflows that rely on ZFS metrics may be affected, leading to delays or failures in critical system operations.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the ZFS exporter logs for errors or issues that may be preventing the collector from scraping metrics.
2. Verify that the ZFS exporter is running and configured correctly.
3. Check the network connectivity between the Prometheus server and the ZFS exporter to ensure that there are no connectivity issues.
4. Review the Prometheus configuration to ensure that the ZFS collector is correctly configured and enabled.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the ZFS exporter service to attempt to re-establish connectivity with Prometheus.
2. Check and update the ZFS exporter configuration to ensure that it is correct and up-to-date.
3. Verify that the Prometheus configuration is correct and up-to-date.
4. If the issue persists, consider increasing the logging verbosity of the ZFS exporter to diagnose any underlying issues.
5. If all else fails, consider escalating the issue to a senior administrator or ZFS expert for further assistance.

Remember to always follow proper troubleshooting procedures and to verify changes before making them to ensure that the issue is fully resolved and that no additional issues are introduced.
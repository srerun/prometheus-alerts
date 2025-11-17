---
title: OPNsenseExporterDown
description: Troubleshooting for alert OPNsenseExporterDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseExporterDown

OPNsense exporter is not responding

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseExporterDown" %}}

{{% comment %}}

```yaml
alert: OPNsenseExporterDown
expr: opnsense_up == 0
for: 1m
labels:
    severity: critical
annotations:
    summary: OPNsense exporter down (instance {{ $labels.opnsense_instance }})
    description: |-
        OPNsense exporter is not responding
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsenseexporterdown/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseExporterDown alert is triggered when the Prometheus server is unable to scrape metrics from the OPNsense exporter for a period of 1 minute. This indicates that the exporter is not responding or is down, which can prevent Prometheus from collecting vital metrics and alerts from the OPNsense device.

## Impact
The impact of this alert is that Prometheus will not be able to collect metrics from the OPNsense device, which can lead to a lack of visibility into the device's performance and health. This can result in delayed detection of issues, making it more difficult to troubleshoot and resolve problems. Additionally, this can also affect the accuracy of alerts and notifications, as Prometheus relies on the exporter to provide up-to-date metrics.

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. Check the OPNsense exporter logs for any errors or issues that may indicate why it is not responding.
2. Verify that the OPNsense exporter is running and configured correctly.
3. Check the network connectivity between the Prometheus server and the OPNsense exporter to ensure that there are no issues preventing communication.
4. Verify that the OPNsense instance is configured correctly and is reachable.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. Restart the OPNsense exporter to see if it resolves the issue.
2. Check and update the OPNsense exporter configuration to ensure it is correct and aligned with the Prometheus server.
3. Verify that the OPNsense instance is properly configured and reachable.
4. If the issue persists, check the Prometheus server logs for any errors or issues related to the OPNsense exporter.
5. Consider increasing the timeout or retry settings for the OPNsense exporter in the Prometheus configuration to allow for temporary connectivity issues.
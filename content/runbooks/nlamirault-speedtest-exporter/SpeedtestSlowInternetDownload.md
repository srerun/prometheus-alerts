---
title: SpeedtestSlowInternetDownload
description: Troubleshooting for alert SpeedtestSlowInternetDownload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SpeedtestSlowInternetDownload

Internet download speed is currently {{humanize $value}} Mbps.

<details>
  <summary>Alert Rule</summary>

{{% rule "speedtest/nlamirault-speedtest-exporter.yml" "SpeedtestSlowInternetDownload" %}}

{{% comment %}}

```yaml
alert: SpeedtestSlowInternetDownload
expr: avg_over_time(speedtest_download[10m]) < 100
for: 0m
labels:
    severity: warning
annotations:
    summary: SpeedTest Slow Internet Download (instance {{ $labels.instance }})
    description: |-
        Internet download speed is currently {{humanize $value}} Mbps.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nlamirault-speedtest-exporter/SpeedtestSlowInternetDownload.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The SpeedtestSlowInternetDownload alert is triggered when the average download speed measured by the Speedtest exporter falls below 100 Mbps over a 10-minute period. This indicates that the internet connection is experiencing slow download speeds, which can impact the performance of applications and services that rely on internet connectivity.

## Impact

A slow internet download speed can have a significant impact on the usability and responsiveness of applications and services. This can lead to frustrated users, decreased productivity, and potential revenue loss. Additionally, slow internet speeds can also impact the performance of critical services such as voice and video conferencing, online backups, and file transfers.

## Diagnosis

To diagnose the root cause of the slow internet download speed, follow these steps:

1. Check the Speedtest exporter logs to verify that the measurements are accurate and not affected by any issues with the exporter itself.
2. Investigate the network infrastructure to identify any bottlenecks or congestion that may be contributing to the slow speeds.
3. Check for any recent changes to the network configuration or topology that may be causing the issue.
4. Verify that the internet service provider (ISP) is not experiencing any outages or maintenance that may be affecting the connection.
5. Check for any malware or viruses that may be consuming bandwidth and slowing down the connection.

## Mitigation

To mitigate the impact of the slow internet download speed, follow these steps:

1. Investigate with the ISP to determine if there are any issues on their end that need to be addressed.
2. Check for any software or firmware updates for the network equipment and apply them as necessary.
3. Consider optimizing network traffic by implementing Quality of Service (QoS) policies to prioritize critical services.
4. Identify and address any bottlenecks in the network infrastructure, such as outdated or malfunctioning equipment.
5. Consider implementing a redundant internet connection to ensure continued availability of critical services in the event of an outage.
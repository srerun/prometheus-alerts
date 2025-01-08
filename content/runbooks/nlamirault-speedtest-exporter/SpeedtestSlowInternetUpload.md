---
title: SpeedtestSlowInternetUpload
description: Troubleshooting for alert SpeedtestSlowInternetUpload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# SpeedtestSlowInternetUpload

Internet upload speed is currently {{humanize $value}} Mbps.

<details>
  <summary>Alert Rule</summary>

{{% rule "speedtest/nlamirault-speedtest-exporter.yml" "SpeedtestSlowInternetUpload" %}}

{{% comment %}}

```yaml
alert: SpeedtestSlowInternetUpload
expr: avg_over_time(speedtest_upload[10m]) < 20
for: 0m
labels:
    severity: warning
annotations:
    summary: SpeedTest Slow Internet Upload (instance {{ $labels.instance }})
    description: |-
        Internet upload speed is currently {{humanize $value}} Mbps.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/nlamirault-speedtest-exporter/SpeedtestSlowInternetUpload.md

```

{{% /comment %}}

</details>


Here is a runbook for the SpeedtestSlowInternetUpload alert:

## Meaning

The SpeedtestSlowInternetUpload alert is triggered when the average internet upload speed, as measured by the speedtest exporter, falls below 20 Mbps over a 10-minute period. This indicates that the internet upload speed is slower than expected, which may impact the performance of applications and services that rely on uploading data to the internet.

## Impact

A slow internet upload speed can have several impacts on the system and its users, including:

* Poor performance of applications that rely on uploading data to the internet, such as cloud-based services or online backups
* Increased latency and timeouts for users uploading data to the internet
* Potential impact on business-critical operations that rely on fast and reliable internet connectivity

## Diagnosis

To diagnose the root cause of the slow internet upload speed, follow these steps:

1. Check the speedtest exporter logs for any errors or issues that may be causing the slow upload speeds.
2. Verify that the internet connection is stable and not experiencing any outages or packet loss.
3. Check the network infrastructure, including routers, switches, and firewalls, to ensure they are functioning correctly and not bottlenecking the upload speed.
4. Investigate any recent changes to the network configuration or infrastructure that may be contributing to the slow upload speeds.

## Mitigation

To mitigate the impact of a slow internet upload speed, follow these steps:

1. Contact the internet service provider (ISP) to report the slow upload speeds and request assistance in troubleshooting and resolving the issue.
2. Consider upgrading the internet plan or infrastructure to improve the upload speeds.
3. Implement Quality of Service (QoS) policies to prioritize critical applications and services that rely on uploading data to the internet.
4. Consider implementing a backup internet connection or redundant infrastructure to ensure business continuity in the event of an internet outage or slow speeds.
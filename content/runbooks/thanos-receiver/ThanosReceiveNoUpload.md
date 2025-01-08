---
title: ThanosReceiveNoUpload
description: Troubleshooting for alert ThanosReceiveNoUpload
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosReceiveNoUpload

Thanos Receive {{$labels.instance}} has not uploaded latest data to object storage.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-receiver.yml" "ThanosReceiveNoUpload" %}}

{{% comment %}}

```yaml
alert: ThanosReceiveNoUpload
expr: (up{job=~".*thanos-receive.*"} - 1) + on (job, instance) (sum by (job, instance) (increase(thanos_shipper_uploads_total{job=~".*thanos-receive.*"}[3h])) == 0)
for: 3h
labels:
    severity: critical
annotations:
    summary: Thanos Receive No Upload (instance {{ $labels.instance }})
    description: |-
        Thanos Receive {{$labels.instance}} has not uploaded latest data to object storage.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-receiver/ThanosReceiveNoUpload.md

```

{{% /comment %}}

</details>


## Meaning

The ThanosReceiveNoUpload alert is triggered when a Thanos Receive instance has not uploaded any data to object storage within the last 3 hours. This indicates a potential issue with the data ingestion pipeline, which may lead to incomplete or missing data in the monitoring system.

## Impact

The impact of this alert is that the monitoring system may not have access to the latest data, which can affect the accuracy of alerts, dashboards, and Grafana visualizations. This can lead to:

* Delayed or missed alerting on critical issues
* Incomplete or inaccurate monitoring data
* Inability to troubleshoot issues due to missing data

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Receive instance logs for any errors or issues related to data upload.
2. Verify that the instance is properly configured and authenticated to access object storage.
3. Check the object storage bucket for any issues or errors related to data upload.
4. Verify that the Thanos Shipper is properly configured and running.

## Mitigation

To mitigate the issue, follow these steps:

1. Investigate and resolve any errors or issues found in the Thanos Receive instance logs.
2. Verify that the instance is properly configured and authenticated to access object storage.
3. Attempt to manually trigger a data upload to object storage.
4. If the issue persists, consider restarting the Thanos Receive instance or consulting with the storage team to resolve any object storage issues.

Additional resources:

* Refer to the Thanos Receive documentation for configuration and troubleshooting guidelines.
* Consult with the storage team for assistance with object storage issues.
* Review the Thanos Shipper configuration and logs for any related issues.
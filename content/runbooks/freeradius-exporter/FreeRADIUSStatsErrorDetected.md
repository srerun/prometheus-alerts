---
title: FreeRADIUSStatsErrorDetected
description: Troubleshooting for alert FreeRADIUSStatsErrorDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# FreeRADIUSStatsErrorDetected

A stats error has been reported in the FreeRADIUS server. Investigate the source of the error.

<details>
  <summary>Alert Rule</summary>

{{% rule "freeradius/freeradius-exporter.yml" "FreeRADIUSStatsErrorDetected" %}}

{{% comment %}}

```yaml
alert: FreeRADIUSStatsErrorDetected
expr: freeradius_stats_error{error!~"^$"} == 1
for: 5m
labels:
    severity: warning
annotations:
    summary: Stats Error Detected
    description: A stats error has been reported in the FreeRADIUS server. Investigate the source of the error.
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/freeradius-exporter/freeradiusstatserrordetected/

```

{{% /comment %}}

</details>


## Meaning

The FreeRADIUSStatsErrorDetected alert is triggered when the FreeRADIUS server reports an error in its statistics. This error can indicate a problem with the source of the error, which may affect the accuracy of the statistics reported by the FreeRADIUS server. This can lead to incorrect decisions based on faulty data.

## Impact

The impact of this alert is moderate, as it may affect the ability to make informed decisions based on accurate statistics. However, it does not necessarily indicate a critical issue with the FreeRADIUS server or its functionality.

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the FreeRADIUS server logs for error messages related to statistics.
2. Investigate the source of the error, which may be a specific module or component of the FreeRADIUS server.
4. Verify that the statistics are being reported correctly and accurately reflect the current state of the system.

## Mitigation

To mitigate the issue, follow these steps:

1. Correct the source of the statistics error, which may require updating or configuring the FreeRADIUS server module.
2. Restart the FreeRADIUS server to ensure that the changes take effect.
3. Verify that the statistics are being reported correctly and accurately reflect the current state of the system.
4. Update the monitoring system to reflect the changes made to the FreeRADIUS server.
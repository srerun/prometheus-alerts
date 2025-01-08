---
title: HostClockSkew
description: Troubleshooting for alert HostClockSkew
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostClockSkew

Clock skew detected. Clock is out of sync. Ensure NTP is configured correctly on this host.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostClockSkew" %}}

{{% comment %}}

```yaml
alert: HostClockSkew
expr: ((node_timex_offset_seconds > 0.05 and deriv(node_timex_offset_seconds[5m]) >= 0) or (node_timex_offset_seconds < -0.05 and deriv(node_timex_offset_seconds[5m]) <= 0)) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 10m
labels:
    severity: warning
annotations:
    summary: Host clock skew (instance {{ $labels.instance }})
    description: |-
        Clock skew detected. Clock is out of sync. Ensure NTP is configured correctly on this host.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostClockSkew.md

```

{{% /comment %}}

</details>


Here is the runbook for the HostClockSkew alert:

## Meaning

The HostClockSkew alert is triggered when the clock on a host is out of sync with the expected time. This can cause issues with timing-sensitive applications and log analysis. The alert is raised when the clock skew exceeds 50 milliseconds and is increasing over a 5-minute period.

## Impact

A clock skew on a host can have several consequences, including:

* Inaccurate timestamps in log files
* Issues with timing-sensitive applications, such as scheduling and caching
* Difficulty correlating logs from multiple hosts
* Potential security issues, such as invalid certificate validation

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the NTP configuration on the host to ensure it is correctly configured.
2. Verify that the NTP service is running and functioning correctly.
3. Check the system logs for any errors related to NTP or clock synchronization.
4. Use the `ntpq` command to check the NTP server status and clock offset.

## Mitigation

To mitigate the issue, follow these steps:

1. Ensure NTP is correctly configured on the host, including setting the correct NTP server and enabling the NTP service.
2. Restart the NTP service to apply any changes.
3. Monitor the clock skew to ensure it returns to a normal value.
4. Consider adjusting the NTP server configuration to use a more reliable or redundant time source.

Remember to investigate and resolve the underlying cause of the clock skew to prevent the issue from recurring.
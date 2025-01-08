---
title: BlackboxProbeSlowPing
description: Troubleshooting for alert BlackboxProbeSlowPing
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxProbeSlowPing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Blackbox ping took more than 1s

<details>
  <summary>Alert Rule</summary>

{{% rule "blackbox/blackbox-exporter.yml" "BlackboxProbeSlowPing" %}}

{{% comment %}}

```yaml
alert: BlackboxProbeSlowPing
expr: avg_over_time(probe_icmp_duration_seconds[1m]) > 1
for: 1m
labels:
    severity: warning
annotations:
    summary: Blackbox probe slow ping (instance {{ $labels.instance }})
    description: |-
        Blackbox ping took more than 1s
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/blackbox-exporter/BlackboxProbeSlowPing.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"

The impact depends on the device being monitored.  If services depend on the device it could start causing downstream service failures.


## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"

- [ ] Manually ping the device to determine response time
- [ ] Login to the device to see if it is responsive.  Check the load
  ```bash
  uptime
  ```
  If the load is excessively high (above 10) that could indicate an issue.
- [ ] If the device appears ok the problem may be network devices in between.  Run a traceroute
	```bash
	traceroute
	```
	Look for any high response times between you and the device
  

## Mitigation
[//]: # "The steps necessary to resolve the alert"

If the load on the box appears ok the issue may resove on its own, otherwise, the issue may be resolved by simply rebooting the device.

If the device appears ok, there may be issues with a device in the network path.

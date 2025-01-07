---
title: BlackboxProbeSlowPing
description: Troubleshooting for alert BlackboxProbeSlowPing
#published: true
date: 2023-12-15T12:40:32.296Z
tags: 
editor: markdown
dateCreated: 2023-12-13T16:44:41.576Z
---

# BlackboxProbeSlowPing

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Ping took more than 1s.  Latency to the monitored device is high.  Could indicate high traffic to the device or issues, such as high load, on the device itself.


<details>
  <summary>Alert Rule</summary>

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
    runbook: http://wiki.ringsq.io/runbook/BlackboxProbeSlowPing

```
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

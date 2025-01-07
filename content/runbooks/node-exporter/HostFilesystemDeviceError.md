---
title: HostFilesystemDeviceError
description: Troubleshooting for alert HostFilesystemDeviceError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostFilesystemDeviceError

## Meaning
[//]: # "Short paragraph that explains what the alert means"
{{ $labels.instance }}: Device error with the {{ $labels.mountpoint }} filesystem

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HostFilesystemDeviceError
expr: node_filesystem_device_error == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Host filesystem device error (instance {{ $labels.instance }})
    description: |-
        {{ $labels.instance }}: Device error with the {{ $labels.mountpoint }} filesystem
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HostFilesystemDeviceError

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

---
title: HostRequiresReboot
description: Troubleshooting for alert HostRequiresReboot
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostRequiresReboot

## Meaning
[//]: # "Short paragraph that explains what the alert means"
{{ $labels.instance }} requires a reboot.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostRequiresReboot" %}}

{{% comment %}}

```yaml
alert: HostRequiresReboot
expr: (node_reboot_required > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 4h
labels:
    severity: info
annotations:
    summary: Host requires reboot (instance {{ $labels.instance }})
    description: |-
        {{ $labels.instance }} requires a reboot.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostRequiresReboot.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

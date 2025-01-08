---
title: HostEdacCorrectableErrorsDetected
description: Troubleshooting for alert HostEdacCorrectableErrorsDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostEdacCorrectableErrorsDetected

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} correctable memory errors reported by EDAC in the last 5 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostEdacCorrectableErrorsDetected" %}}

{{% comment %}}

```yaml
alert: HostEdacCorrectableErrorsDetected
expr: (increase(node_edac_correctable_errors_total[1m]) > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: info
annotations:
    summary: Host EDAC Correctable Errors detected (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} correctable memory errors reported by EDAC in the last 5 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostEdacCorrectableErrorsDetected.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

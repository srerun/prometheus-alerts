---
title: HostEdacUncorrectableErrorsDetected
description: Troubleshooting for alert HostEdacUncorrectableErrorsDetected
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostEdacUncorrectableErrorsDetected

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} uncorrectable memory errors reported by EDAC in the last 5 minutes.

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostEdacUncorrectableErrorsDetected" %}}

<!-- Rule when generated

```yaml
alert: HostEdacUncorrectableErrorsDetected
expr: (node_edac_uncorrectable_errors_total > 0) * on(instance) group_left (nodename) node_uname_info{nodename=~".+"}
for: 0m
labels:
    severity: warning
annotations:
    summary: Host EDAC Uncorrectable Errors detected (instance {{ $labels.instance }})
    description: |-
        Host {{ $labels.instance }} has had {{ printf "%.0f" $value }} uncorrectable memory errors reported by EDAC in the last 5 minutes.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostEdacUncorrectableErrorsDetected.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

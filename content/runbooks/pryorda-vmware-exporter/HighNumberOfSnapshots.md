---
title: HighNumberOfSnapshots
description: Troubleshooting for alert HighNumberOfSnapshots
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HighNumberOfSnapshots

## Meaning
[//]: # "Short paragraph that explains what the alert means"
High snapshots number on {{ $labels.instance }}: {{ $value }}

<details>
  <summary>Alert Rule</summary>

{{% rule "vmware/pryorda-vmware-exporter.yml" "HighNumberOfSnapshots" %}}

<!-- Rule when generated

```yaml
alert: HighNumberOfSnapshots
expr: vmware_vm_snapshots > 3
for: 30m
labels:
    severity: warning
annotations:
    summary: High Number of Snapshots (instance {{ $labels.instance }})
    description: |-
        High snapshots number on {{ $labels.instance }}: {{ $value }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pryorda-vmware-exporter/HighNumberOfSnapshots.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

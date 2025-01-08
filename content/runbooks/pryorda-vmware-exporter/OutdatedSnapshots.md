---
title: OutdatedSnapshots
description: Troubleshooting for alert OutdatedSnapshots
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OutdatedSnapshots

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Outdated snapshots on {{ $labels.instance }}: {{ $value | printf "%.0f"}} days

<details>
  <summary>Alert Rule</summary>

{{% rule "vmware/pryorda-vmware-exporter.yml" "OutdatedSnapshots" %}}

<!-- Rule when generated

```yaml
alert: OutdatedSnapshots
expr: (time() - vmware_vm_snapshot_timestamp_seconds) / (60 * 60 * 24) >= 3
for: 5m
labels:
    severity: warning
annotations:
    summary: Outdated Snapshots (instance {{ $labels.instance }})
    description: |-
        Outdated snapshots on {{ $labels.instance }}: {{ $value | printf "%.0f"}} days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/pryorda-vmware-exporter/OutdatedSnapshots.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

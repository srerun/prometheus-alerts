---
title: NetdataMdMismatchCntUnsynchronizedBlocks
description: Troubleshooting for alert NetdataMdMismatchCntUnsynchronizedBlocks
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# NetdataMdMismatchCntUnsynchronizedBlocks

## Meaning
[//]: # "Short paragraph that explains what the alert means"
RAID Array have unsynchronized blocks

<details>
  <summary>Alert Rule</summary>

{{% rule "netdata/netdata-internal.yml" "NetdataMdMismatchCntUnsynchronizedBlocks" %}}

<!-- Rule when generated

```yaml
alert: NetdataMdMismatchCntUnsynchronizedBlocks
expr: netdata_md_mismatch_cnt_unsynchronized_blocks_average > 1024
for: 2m
labels:
    severity: warning
annotations:
    summary: Netdata MD mismatch cnt unsynchronized blocks (instance {{ $labels.instance }})
    description: |-
        RAID Array have unsynchronized blocks
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/netdata-internal/NetdataMdMismatchCntUnsynchronizedBlocks.md

```

-->

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

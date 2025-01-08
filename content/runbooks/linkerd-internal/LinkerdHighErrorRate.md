---
title: LinkerdHighErrorRate
description: Troubleshooting for alert LinkerdHighErrorRate
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# LinkerdHighErrorRate

Linkerd error rate for {{ $labels.deployment | $labels.statefulset | $labels.daemonset }} is over 10%

<details>
  <summary>Alert Rule</summary>

{{% rule "linkerd/linkerd-internal.yml" "LinkerdHighErrorRate" %}}

{{% comment %}}

```yaml
alert: LinkerdHighErrorRate
expr: sum(rate(request_errors_total[1m])) by (deployment, statefulset, daemonset) / sum(rate(request_total[1m])) by (deployment, statefulset, daemonset) * 100 > 10
for: 1m
labels:
    severity: warning
annotations:
    summary: Linkerd high error rate (instance {{ $labels.instance }})
    description: |-
        Linkerd error rate for {{ $labels.deployment | $labels.statefulset | $labels.daemonset }} is over 10%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/linkerd-internal/LinkerdHighErrorRate.md

```

{{% /comment %}}

</details>


## Meaning
[//]: # "Short paragraph that explains what the alert means"


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

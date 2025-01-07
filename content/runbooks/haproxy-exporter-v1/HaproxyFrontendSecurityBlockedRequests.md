---
title: HaproxyFrontendSecurityBlockedRequests
description: Troubleshooting for alert HaproxyFrontendSecurityBlockedRequests
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HaproxyFrontendSecurityBlockedRequests

## Meaning
[//]: # "Short paragraph that explains what the alert means"
HAProxy is blocking requests for security reason

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: HaproxyFrontendSecurityBlockedRequests
expr: sum by (frontend) (rate(haproxy_frontend_requests_denied_total[2m])) > 10
for: 2m
labels:
    severity: warning
annotations:
    summary: HAProxy frontend security blocked requests (instance {{ $labels.instance }})
    description: |-
        HAProxy is blocking requests for security reason
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/HaproxyFrontendSecurityBlockedRequests

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

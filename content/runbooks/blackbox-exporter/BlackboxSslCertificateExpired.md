---
title: BlackboxSslCertificateExpired
description: Troubleshooting for alert BlackboxSslCertificateExpired
#published: true
date: 2023-12-12T21:12:32.022Z
tags: LGTM
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# BlackboxSslCertificateExpired

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SSL certificate has expired already

<details>
  <summary>Alert Rule</summary>

  ```yaml
alert: BlackboxSslCertificateExpired
expr: round((last_over_time(probe_ssl_earliest_cert_expiry[10m]) - time()) / 86400, 0.1) < 0
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox SSL certificate expired (instance {{ $labels.instance }})
    description: |-
        SSL certificate has expired already
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/content/runbooks/BlackboxSslCertificateExpired

  ```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

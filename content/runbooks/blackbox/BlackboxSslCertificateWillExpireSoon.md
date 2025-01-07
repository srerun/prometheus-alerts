---
title: BlackboxSslCertificateWillExpireSoon
description: Troubleshooting for alert BlackboxSslCertificateWillExpireSoon
#published: true
date: 2023-12-15T17:15:01.389Z
tags: 
editor: markdown
dateCreated: 2023-12-13T16:44:46.826Z
---

# BlackboxSslCertificateWillExpireSoon

## Meaning
[//]: # "Short paragraph that explains what the alert means"
SSL certificate expires in less than X days


<details>
  <summary>Alert Rule</summary>

```yaml
alert: BlackboxSslCertificateWillExpireSoon
expr: 0 <= round((last_over_time(probe_ssl_earliest_cert_expiry[10m]) - time()) / 86400, 0.1) < 3
for: 0m
labels:
    severity: critical
annotations:
    summary: Blackbox SSL certificate will expire soon (instance {{ $labels.instance }})
    description: |-
        SSL certificate expires in less than 3 days
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/BlackboxSslCertificateWillExpireSoon

```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

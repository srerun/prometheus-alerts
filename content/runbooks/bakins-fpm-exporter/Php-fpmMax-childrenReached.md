---
title: Php-fpmMax-childrenReached
description: Troubleshooting for alert Php-fpmMax-childrenReached
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# Php-fpmMax-childrenReached

## Meaning
[//]: # "Short paragraph that explains what the alert means"
PHP-FPM reached max children - {{ $labels.instance }}

<details>
  <summary>Alert Rule</summary>

{{% rule "php-fpm/bakins-fpm-exporter.yml" "Php-fpmMax-childrenReached" %}}

{{% comment %}}

```yaml
alert: Php-fpmMax-childrenReached
expr: sum(phpfpm_max_children_reached_total) by (instance) > 0
for: 0m
labels:
    severity: warning
annotations:
    summary: PHP-FPM max-children reached (instance {{ $labels.instance }})
    description: |-
        PHP-FPM reached max children - {{ $labels.instance }}
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/bakins-fpm-exporter/Php-fpmMax-childrenReached.md

```

{{% /comment %}}

</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"



## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"



## Mitigation
[//]: # "The steps necessary to resolve the alert"

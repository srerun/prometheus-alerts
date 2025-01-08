---
title: ThanosRuleConfigReloadFailure
description: Troubleshooting for alert ThanosRuleConfigReloadFailure
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ThanosRuleConfigReloadFailure

Thanos Rule {{$labels.job}} has not been able to reload its configuration.

<details>
  <summary>Alert Rule</summary>

{{% rule "thanos/thanos-ruler.yml" "ThanosRuleConfigReloadFailure" %}}

{{% comment %}}

```yaml
alert: ThanosRuleConfigReloadFailure
expr: avg by (job, instance) (thanos_rule_config_last_reload_successful{job=~".*thanos-rule.*"}) != 1
for: 5m
labels:
    severity: info
annotations:
    summary: Thanos Rule Config Reload Failure (instance {{ $labels.instance }})
    description: |-
        Thanos Rule {{$labels.job}} has not been able to reload its configuration.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/thanos-ruler/ThanosRuleConfigReloadFailure.md

```

{{% /comment %}}

</details>


## Meaning

The `ThanosRuleConfigReloadFailure` alert is triggered when Thanos Rule cannot reload its configuration. This alert indicates that there is an issue with the configuration reload process, which may prevent Thanos Rule from functioning correctly.

## Impact

If this alert is not addressed, Thanos Rule may not be able to function correctly, leading to:

* Inaccurate or incomplete metric data
* Failure to detect anomalies or alert on critical events
* Increased risk of errors or mistakes in the monitoring and alerting system

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Thanos Rule logs for errors related to configuration reload.
2. Verify that the configuration file is correctly formatted and valid.
3. Check the Thanos Rule service status and ensure it is running correctly.
4. Verify that the configuration reload is not being blocked by any network or firewall issues.
5. Check the Prometheus metric `thanos_rule_config_last_reload_successful` to see if it is consistently reporting 0 or other non-successful values.

## Mitigation

To mitigate the issue, follow these steps:

1. Check the Thanos Rule configuration file for errors and correct any syntax issues.
2. Restart the Thanos Rule service to attempt to reload the configuration.
3. Verify that the configuration reload is successful by checking the `thanos_rule_config_last_reload_successful` metric.
4. If the issue persists, check the Thanos Rule logs for any additional error messages and investigate further.
5. If none of the above steps resolve the issue, consider seeking assistance from a Thanos Rule expert or the Prometheus community.
---
title: BlackboxConfigurationReloadFailure
description: Troubleshooting for alert BlackboxConfigurationReloadFailure
#published: true
date: 2023-12-15T13:48:54.534Z
tags: 
editor: markdown
dateCreated: 2023-12-13T16:44:33.808Z
---

# BlackboxConfigurationReloadFailure

## Meaning
[//]: # "Short paragraph that explains what the alert means"
Blackbox configuration reload failure.  This indicates that the configuration for the blackbox_exporter has changed, but the the new configuration is invalid.

<details>
  <summary>Alert Rule</summary>

```yaml
alert: BlackboxConfigurationReloadFailure
expr: blackbox_exporter_config_last_reload_successful != 1
for: 0m
labels:
    severity: warning
annotations:
    summary: Blackbox configuration reload failure (instance {{ $labels.instance }})
    description: |-
        Blackbox configuration reload failure
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: http://wiki.ringsq.io/runbook/BlackboxConfigurationReloadFailure

```
</details>


## Impact
[//]: # "What could / will happen if the alert is not addressed"

Whatever change was being made to the configuration has not been applied.  This will cause checks to not have the options desired or fail entirely if it was a new configuration section.

## Diagnosis
[//]: # "Steps to take to identify the cause of the problem"

- [ ] Login to the device reporting the error
- [ ] Check the logs to see what the reload indicated
  ```bash
  systemctl status blackbox_exporter
  ```
  If the error occurred some time ago the error may not appear.  In that case, check the journal:
  ```bash
  journalctl -xeu blackbox_exporter -S today
  ```
  You can also list just log entries for yesterday.  eg.
  ```bash
  journalctl -xeu blackbox_exporter -S yesterday -U today
  ```
  or for a date/time range (if time is omitted 00:00:00 is used)
  ```bash
  journalctl -xeu blackbox_exporter -S 2023-12-01 -U "2023-12-01 12:00:00"

## Mitigation
[//]: # "The steps necessary to resolve the alert"

Once the underlying cause of the issue is identified, the configuration must be corrected.  While this can be done on the server issuing the error, it also needs to be done in the ansible playbook so that subsequent deployments do not reintroduce the same error.

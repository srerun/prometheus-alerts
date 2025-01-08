---
title: ClickhouseAccessDeniedErrors
description: Troubleshooting for alert ClickhouseAccessDeniedErrors
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ClickhouseAccessDeniedErrors

Access denied errors have been logged, which could indicate permission issues or unauthorized access attempts.

<details>
  <summary>Alert Rule</summary>

{{% rule "clickhouse/clickhouse-internal.yml" "ClickhouseAccessDeniedErrors" %}}

{{% comment %}}

```yaml
alert: ClickhouseAccessDeniedErrors
expr: increase(ClickHouseErrorMetric_RESOURCE_ACCESS_DENIED[5m]) > 0
for: 0m
labels:
    severity: info
annotations:
    summary: ClickHouse Access Denied Errors (instance {{ $labels.instance }})
    description: |-
        Access denied errors have been logged, which could indicate permission issues or unauthorized access attempts.
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/clickhouse-internal/ClickhouseAccessDeniedErrors.md

```

{{% /comment %}}

</details>


## Meaning

The ClickhouseAccessDeniedErrors alert indicates that ClickHouse has logged access denied errors, which may be caused by permission issues or unauthorized access attempts. This alert is triggered when the number of access denied errors increases over a 5-minute period.

## Impact

The impact of this alert can be significant, as access denied errors can indicate:

* Unauthorized access attempts to sensitive data
* Permission issues that prevent legitimate users from accessing data
* Potential security breaches or data leaks
* Performance degradation due to repeated access attempts

If left unaddressed, these errors can lead to data compromise, system instability, and reputational damage.

## Diagnosis

To diagnose the root cause of the ClickhouseAccessDeniedErrors alert, follow these steps:

1. **Check ClickHouse logs**: Review the ClickHouse logs to identify the specific errors and the users or applications that are generating the access denied errors.
2. **Verify permission settings**: Check the permission settings for the affected users or applications to ensure they have the necessary access to the data.
3. **Investigate access patterns**: Analyze the access patterns to identify any unusual or suspicious activity.
4. **Check system configuration**: Verify that the ClickHouse system configuration is correct and up-to-date.

## Mitigation

To mitigate the ClickhouseAccessDeniedErrors alert, follow these steps:

1. **Update permission settings**: Update the permission settings to ensure that users or applications have the necessary access to the data.
2. **Implement access controls**: Implement additional access controls, such as IP whitelisting or two-factor authentication, to prevent unauthorized access attempts.
3. **Monitor system activity**: Continuously monitor system activity to detect and respond to potential security breaches.
4. **Review system configuration**: Review and update the ClickHouse system configuration to ensure it is correct and up-to-date.

Remember to update the alert annotations with the results of the diagnosis and mitigation steps, and to notify relevant stakeholders of the issue and its resolution.
---
title: OPNsenseServiceDown
description: Troubleshooting for alert OPNsenseServiceDown
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# OPNsenseServiceDown

Service {{ $labels.name }} ({{ $labels.description }}) is down

<details>
  <summary>Alert Rule</summary>

{{% rule "opnsense/opnsense.yml" "OPNsenseServiceDown" %}}

{{% comment %}}

```yaml
alert: OPNsenseServiceDown
expr: opnsense_services_status == 0
for: 2m
labels:
    severity: critical
annotations:
    summary: OPNsense service down (instance {{ $labels.opnsense_instance }})
    description: |-
        Service {{ $labels.name }} ({{ $labels.description }}) is down
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://srerun.github.io/prometheus-alerts/runbooks/opnsense/opnsenseservicedown/

```

{{% /comment %}}

</details>


## Meaning
The OPNsenseServiceDown alert is triggered when the status of an OPNsense service is reported as down (indicated by a value of 0). This alert is critical and indicates that a service, which is a crucial component of the OPNsense firewall, is not functioning properly. The alert provides information about the specific service that is down, including its name and description.

## Impact
The impact of this alert can be significant, as a down service can compromise the security and functionality of the network. Depending on the specific service that is down, the impact may include:
* Loss of network connectivity or access to certain resources
* Compromised security posture, potentially allowing unauthorized access to the network
* Disruption of critical business operations or services
* Potential data loss or corruption

## Diagnosis
To diagnose the issue, the following steps can be taken:
1. **Check the OPNsense logs**: Review the OPNsense logs to determine the cause of the service failure.
2. **Verify service configuration**: Check the service configuration to ensure that it is correctly set up and that there are no errors.
3. **Check for software updates**: Verify that the OPNsense software is up to date, as outdated software may be the cause of the issue.
4. **Check system resources**: Verify that the system has sufficient resources (e.g., CPU, memory, disk space) to run the service.
5. **Check for conflicts with other services**: Verify that there are no conflicts with other services that may be causing the issue.

## Mitigation
To mitigate the issue, the following steps can be taken:
1. **Restart the service**: Attempt to restart the service to see if it will come back online.
2. **Restore from backup**: If the service is not recoverable, restore the service from a backup.
3. **Reconfigure the service**: If the issue is due to a configuration error, reconfigure the service to correct the issue.
4. **Apply software updates**: If the issue is due to outdated software, apply the necessary software updates.
5. **Contact support**: If none of the above steps resolve the issue, contact OPNsense support for further assistance.
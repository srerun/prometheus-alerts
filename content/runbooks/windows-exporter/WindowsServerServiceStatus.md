---
title: WindowsServerServiceStatus
description: Troubleshooting for alert WindowsServerServiceStatus
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# WindowsServerServiceStatus

Windows Service state is not OK

<details>
  <summary>Alert Rule</summary>

{{% rule "windows-server/windows-exporter.yml" "WindowsServerServiceStatus" %}}

{{% comment %}}

```yaml
alert: WindowsServerServiceStatus
expr: windows_service_status{status="ok"} != 1
for: 1m
labels:
    severity: critical
annotations:
    summary: Windows Server service Status (instance {{ $labels.instance }})
    description: |-
        Windows Service state is not OK
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/windows-exporter/WindowsServerServiceStatus.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule:

## Meaning

The WindowsServerServiceStatus alert is triggered when the status of a Windows Server service is not "ok". This alert is critical and indicates that there is an issue with the service that needs to be addressed promptly.

## Impact

The impact of this alert can be significant, as it may indicate that a critical service is not functioning properly. This can lead to a range of issues, including:

* Disruption to business-critical applications and services
* Data loss or corruption
* Security vulnerabilities
* System instability or crashes

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the service status: Verify the status of the affected service using the Windows Service Manager or PowerShell.
2. Review system logs: Check the system logs for errors or warnings related to the service.
3. Verify service configuration: Check the service configuration to ensure that it is set up correctly.
4. Check for dependencies: Verify that all dependencies required by the service are available and functioning properly.

## Mitigation

To mitigate the issue, follow these steps:

1. Restart the service: Attempt to restart the affected service to see if it resolves the issue.
2. Investigate and resolve underlying issues: Identify and resolve any underlying issues causing the service to fail, such as configuration problems, dependencies, or system resource issues.
3. Contact support: If the issue persists, contact Windows support or a system administrator for further assistance.
4. Monitor service status: Continuously monitor the service status to ensure that it remains stable and functional.

Remember to update the runbook with specific steps and procedures relevant to your organization's Windows Server environment.
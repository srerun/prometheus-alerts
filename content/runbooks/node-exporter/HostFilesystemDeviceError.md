---
title: HostFilesystemDeviceError
description: Troubleshooting for alert HostFilesystemDeviceError
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# HostFilesystemDeviceError

{{ $labels.instance }}: Device error with the {{ $labels.mountpoint }} filesystem

<details>
  <summary>Alert Rule</summary>

{{% rule "host-and-hardware/node-exporter.yml" "HostFilesystemDeviceError" %}}

{{% comment %}}

```yaml
alert: HostFilesystemDeviceError
expr: node_filesystem_device_error == 1
for: 2m
labels:
    severity: critical
annotations:
    summary: Host filesystem device error (instance {{ $labels.instance }})
    description: |-
        {{ $labels.instance }}: Device error with the {{ $labels.mountpoint }} filesystem
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/node-exporter/HostFilesystemDeviceError.md

```

{{% /comment %}}

</details>


Here is a runbook for the Prometheus alert rule `HostFilesystemDeviceError`:

## Meaning

The `HostFilesystemDeviceError` alert is triggered when a device error is detected on a filesystem of a host machine. This error is reported by the Node Exporter, which is a Prometheus exporter that collects metrics about a host's kernel, network, and disk performance. The alert indicates that there is a problem with the underlying device that stores the filesystem, which can lead to data loss, corruption, or unavailability.

## Impact

The impact of this alert can be severe, as it may prevent the host from writing data to the affected filesystem. This can lead to:

* Data loss or corruption
* Service unavailability
* System instability
* Increased risk of further errors or failures

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the Node Exporter metrics for the affected host to identify the specific device and filesystem that are experiencing errors.
2. Review system logs for errors related to the device or filesystem.
3. Check the disk's SMART status to see if it's reporting any errors.
4. Run a disk check (e.g. `fsck`) to identify and potentially fix any file system errors.
5. Verify that the system is configured to send alerts and notifications for device errors.

## Mitigation

To mitigate the issue, follow these steps:

1. Immediately stop writing data to the affected filesystem to prevent further data loss or corruption.
2. Unmount the affected filesystem to prevent further errors.
3. Run a disk check (e.g. `fsck`) to identify and fix any file system errors.
4. If the error is related to a hardware failure, replace the faulty device as soon as possible.
5. If the error is related to a software configuration issue, correct the configuration and restart the affected services.
6. Verify that the system is configured to send alerts and notifications for device errors.
7. Once the issue is resolved, remount the filesystem and resume normal operations.

Remember to investigate the root cause of the issue and take corrective action to prevent similar errors from occurring in the future.
---
title: ContainerVolumeUsage
description: Troubleshooting for alert ContainerVolumeUsage
#published: true
date: 2023-12-12T21:12:32.022Z
tags: 
  - LGTM
  - generated
editor: markdown
dateCreated: 2020-04-10T18:32:27.079Z
---

# ContainerVolumeUsage

Container Volume usage is above 80%

<details>
  <summary>Alert Rule</summary>

{{% rule "docker-containers/google-cadvisor.yml" "ContainerVolumeUsage" %}}

{{% comment %}}

```yaml
alert: ContainerVolumeUsage
expr: (1 - (sum(container_fs_inodes_free{name!=""}) BY (instance) / sum(container_fs_inodes_total) BY (instance))) * 100 > 80
for: 2m
labels:
    severity: warning
annotations:
    summary: Container Volume usage (instance {{ $labels.instance }})
    description: |-
        Container Volume usage is above 80%
          VALUE = {{ $value }}
          LABELS = {{ $labels }}
    runbook: https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerVolumeUsage.md

```

{{% /comment %}}

</details>


## Meaning

The ContainerVolumeUsage alert is triggered when the available inodes on a container's filesystem falls below 20%, indicating that the container's volume usage is above 80%. This alert is critical because it can lead to container crashes, data loss, and performance degradation.

## Impact

If left unaddressed, high container volume usage can cause:

* Container crashes and failures
* Data loss and corruption
* Performance degradation and slow response times
* Increased risk of security breaches due to overflowing logs and temporary files

## Diagnosis

To diagnose the issue, follow these steps:

1. Check the container's filesystem usage using commands like `docker exec -it <container_name> df -h` or `kubectl exec -it <pod_name> -c <container_name> df -h`
2. Verify that the container is not writing excessive amounts of data to its filesystem
3. Check for any unnecessary files or directories that can be safely removed
4. Investigate if the container's volume is properly configured and resized
5. Review the container's logging configuration to ensure that logs are not overflowing and filling up the filesystem

## Mitigation

To mitigate the issue, follow these steps:

1. Identify and remove any unnecessary files or directories taking up space on the container's filesystem
2. Resize the container's volume to provide more storage capacity
3. Implement log rotation and compression to prevent log files from growing indefinitely
4. Configure the container to write data to an external storage volume or a cloud-based object store
5. Consider implementing quotas or limits on the container's filesystem usage to prevent future instances of high volume usage

Note: For more detailed steps and recommendations, refer to the [runbook](https://github.com/srerun/prometheus-alerts/blob/main/content/runbooks/google-cadvisor/ContainerVolumeUsage.md) provided in the alert annotation.
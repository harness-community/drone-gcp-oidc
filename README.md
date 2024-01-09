# drone-gcp-oidc

- [Synopsis](#Synopsis)
- [Parameters](#Parameters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin generates an access token through the OIDC token and outputs it as an environment variable. This variable can be utilized in subsequent pipeline steps to control Google Cloud Services through the gcloud CLI or API using curl.

To learn how to utilize Drone plugins in Harness CI, please consult the provided [documentation](https://developer.harness.io/docs/continuous-integration/use-ci/use-drone-plugins/run-a-drone-plugin-in-ci).

## Parameters

| Parameter                                                                                                                           | Choices/<span style="color:blue;">Defaults</span> | Comments                                             |
| :---------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | ---------------------------------------------------- |
| oidc_id_token <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>         |                                                   | This token is used for authenticating with GCP.      |
| project_number <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>        |                                                   | The project number associated with your GCP project. |
| pool_id <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>               |                                                   | The pool ID for OIDC authentication.                 |
| provider_id <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>           |                                                   | The provider ID for OIDC authentication.             |
| service_account_email <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | The email address of the service account.            |

## Notes

The plugin outputs the access token in the form of an environment variable that can be accessed in the subsequent pipeline steps like this: `<+steps.STEP_ID.output.outputVariables.GCLOUD_ACCESS_TOKEN>`

## Plugin Image

The plugin is available for the following architectures:

| OS            | Tag             |
| ------------- | --------------- |
| linux/amd64   | `linux-amd64`   |
| linux/arm64   | `linux-arm64`   |
| windows/amd64 | `windows-amd64` |

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: drone-gcp-oidc-plugin
    identifier: drone_gcp_oidc_plugin
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunity/drone-gcp-oidc:linux-amd64
        settings:
                oidc_id_token: <+secrets.getValue("oidc-token")>
                project_number: 22819301
                pool_id: d8291ka22
                pool_id: kda91fa
                service_account_email: test-gcp@harness.io


# Run step to use the access token to list the compute zones
- step:
    type: Run
    name: List Compute Engine Zone
    identifier: list_zones
    spec:
        shell: Sh
        command: |-
            curl -H "Authorization: Bearer <+steps.STEP_ID.output.outputVariables.GCLOUD_ACCESS_TOKEN>" \
            "https://compute.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances"
```

> <span style="font-size: 14px; margin-left:5px; background-color: #d3d3d3; padding: 4px; border-radius: 4px;">ℹ️ If you notice any issues in this documentation, you can [edit this document](https://github.com/harness-community/drone-gcp-oidc/blob/main/README.md) to improve it.</span>

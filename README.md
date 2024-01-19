# drone-docker-image-migration

- [Synopsis](#Synopsis)
- [Parameters](#Paramaters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin is designed to migrate a Docker image from one registry to another.

To learn how to utilize Drone plugins in Harness CI, please consult the provided [documentation](https://developer.harness.io/docs/continuous-integration/use-ci/use-drone-plugins/run-a-drone-plugin-in-ci).

## Parameters

| Parameter                                                                                                                           | Choices/<span style="color:blue;">Defaults</span> | Comments                                                                 |
| :---------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | :----------------------------------------------------------------------- |
| source_registry <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>       |                                                   | The source docker registry                                               |
| source_username <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>       |                                                   | Username to login to the source registry                                 |
| source_password <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>       |                                                   | PAT / access token to authenticate with the source registry              |
| destination_registry <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>  |                                                   | The destination docker registry                                          |
| destination_username <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>  |                                                   | Username to login to the destination registry                            |
| destination_password <span style="font-size: 10px"><br/>`string`</span>                                                             |                                                   | PAT / access token to authenticate with the destination registry         |
| source_namespace <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>      |                                                   | Source namespace to pull the image from                                  |
| destination_namespace <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | Destination namespace to push the image to                               |
| image_name <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>            |                                                   | The docker image name to be migrated from source to destination registry |
| image_tag <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>             |                                                   | The docker image tag to be migrated from source to destination registry  |
| aws_access_key_id <span style="font-size: 10px"><br/>`string`</span>                                                                |                                                   | AWS access key ID for generating access token                            |
| aws_secret_access_key <span style="font-size: 10px"><br/>`string`</span>                                                            |                                                   | AWS secret access key for generating access token                        |
| aws_region <span style="font-size: 10px"><br/>`string`</span>                                                                       |                                                   | AWS region containing the ECR registry                                   |

## Notes

While using AWS ECR as destination registy, set `destination_username` as `AWS`, and either provide the AWS access token as `destination_password`, or provide the `aws_access_key_id`, `aws_secret_access_key` and `aws_region`.

## Plugin Image

The plugin `harnesscommunity/drone-docker-image-migration` is available for the following architectures:

| OS          | Tag           |
| ----------- | ------------- |
| linux/amd64 | `linux-amd64` |
| linux/arm64 | `linux-arm64` |

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: Migration Plugin
    identifier: Migration_Plugin
    spec:
        connectorRef: harness-connector
        image: harnesscommunity/drone-docker-image-migration:linux-amd64
        settings:
                source_registry: registry.hub.docker.com
                source_username: <+variable.source_username>
                source_password: <+secrets.getValue("source_pat")>
                image_name: image
                image_tag: latest
                destination_registry: account-id.dkr.ecr.region.amazonaws.com
                destination_username: AWS
                destination_password: <+secrets.getValue("aws_acess_token")>
                source_namespace: <+variable.docker_username>
                destination_namespace: aws-repo-name

# Without destination_password
- step:
    type: Plugin
    name: Migration Plugin
    identifier: Migration_Plugin
    spec:
        connectorRef: harness-connector
        image: harnesscommunity/drone-docker-image-migration:linux-amd64
        settings:
                source_registry: registry.hub.docker.com
                source_username: <+variable.source_username>
                source_password: <+secrets.getValue("source_pat")>
                image_name: image
                image_tag: latest
                destination_registry: account-id.dkr.ecr.region.amazonaws.com
                destination_username: AWS
                source_namespace: <+variable.docker_username>
                destination_namespace: aws-repo-name
                aws_access_key_id: <+secrets.getValue("access_key_id")>
                aws_secret_access_key: <+secrets.getValue("secret_access_key")>
                aws_region: us-east-1

```

> <span style="font-size: 14px; margin-left:5px; background-color: #d3d3d3; padding: 4px; border-radius: 4px;">ℹ️ If you notice any issues in this documentation, you can [edit this document](https://github.com/harness-community/drone-docker-image-migration/blob/main/README.md) to improve it.</span>

# parameter.aws-ssm
Provides parameters from AWS SSM (Parameter Store)

## Usage

```
parameters:
  - use: https://github.com/velocity-ci/parameter.aws-ssm/releases/download/0.1.0/aws-ssm
    arguments:
      name: <parameter name in AWS SSM Parameter Store>
    exports:
      value: <parameter name>
```

See `tasks/publish.yml` for example.
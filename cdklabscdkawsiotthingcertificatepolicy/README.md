# AWS IoT Thing, Certificate, and Policy Construct Library

[![NPM](https://img.shields.io/npm/v/@cdklabs/cdk-aws-iot-thing-certificate-policy?label=npm+cdk+v2)](https://www.npmjs.com/package/@cdklabs/cdk-aws-iot-thing-certificate-policy)
[![PyPI](https://img.shields.io/pypi/v/cdklabs.cdk-aws-iot-thing-certificate-policy?label=pypi+cdk+v2)](https://pypi.org/project/cdklabs.cdk-aws-iot-thing-certificate-policy/)
[![Maven version](https://img.shields.io/maven-central/v/io.github.cdklabs/cdk-aws-iot-thing-certificate-policy?label=maven+cdk+v2)](https://central.sonatype.com/artifact/io.github.cdklabs/cdk-aws-iot-thing-certificate-policy)
[![NuGet version](https://img.shields.io/nuget/v/Cdklabs.CdkAwsIotThingCertificatePolicy?label=nuget+cdk+v2)](https://www.nuget.org/packages/Cdklabs.CdkNag)
[![Go version](https://img.shields.io/github/go-mod/go-version/cdklabs/cdk-aws-iot-thing-certificate-policy?label=go+cdk+v2)](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy-go)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](https://github.com/DataDog/datadog-cdk-constructs/blob/main/LICENSE)

![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

[![View on Construct Hub](https://constructs.dev/badge?package=@cdklabs/cdk-aws-iot-thing-certificate-policy)](https://constructs.dev/packages/@cdklabs/cdk-aws-iot-thing-certificate-policy)

An [L3 CDK construct](https://docs.aws.amazon.com/cdk/v2/guide/constructs.html#constructs_lib) to create and associate a singular AWS IoT Thing, Certificate, and IoT Policy. The construct also retrieves and returns AWS IoT account specific details such as the AWS IoT data endpoint and the AWS IoT Credential provider endpoint.

The certificate and its private key are stored as AWS Systems Manager Parameter Store parameters that can be retrieved via the AWS Console or programmatically via construct members.

## Installation and use

<details>
  <summary><b>TypeScript</b></summary>

**Installation:**

```shell
npm install cdk-aws-iot-thing-certificate-policy
```

[**API Reference**](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy/blob/main/doc/api-typescript.md)

**Example:**

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/cdklabs/cdk-aws-iot-thing-certificate-policy-go/cdklabscdkawsiotthingcertificatepolicy"

/**
 * A minimum IoT Policy template using substitution variables for actual
 * policy to be deployed for "region", "account", and "thingname". Allows
 * the thing to publish and subscribe on any topics under "thing/*" topic
 * namespace. Normal IoT Policy conventions such as "*", apply.
 */
minimalIotPolicy := `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["iot:Connect"],
      "Resource": "arn:aws:iot:{{region}}:{{account}}:client/{{thingname}}"
    },
    {
      "Effect": "Allow",
      "Action": ["iot:Publish"],
      "Resource": [
        "arn:aws:iot:{{region}}:{{account}}:topic/{{thingname}}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": ["iot:Subscribe"],
      "Resource": [
        "arn:aws:iot:{{region}}:{{account}}:topicfilter/{{thingname}}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": ["iot:Receive"],
      "Resource": [
        "arn:aws:iot:{{region}}:{{account}}:topic/{{thingname}}/*"
      ]
    }
  ]
}`

app := cdk.NewApp()

/**
 * Create the thing, certificate, and policy, then associate the
 * certificate to both the thing and the policy and fully activate.
 */
fooThing := cdklabscdkawsiotthingcertificatepolicy.NewIotThingCertificatePolicy(app, jsii.String("MyFooThing"), &IotThingCertificatePolicyProps{
	ThingName: jsii.String("foo-thing"),
	 // Name to assign to AWS IoT thing, and value for {{thingname}} in policy template
	IotPolicyName: jsii.String("foo-iot-policy"),
	 // Name to assign to AWS IoT policy
	IotPolicy: minimalIotPolicy,
	 // Policy with or without substitution parameters from above
	EncryptionAlgorithm: jsii.String("ECC"),
	 // Algorithm to use to private key (RSA or ECC)
	PolicyParameterMapping: []policyMapping{
		&policyMapping{
			Name: jsii.String("region"),
			Value: cdk.Fn_Ref(jsii.String("AWS::Region")),
		},
		&policyMapping{
			Name: jsii.String("account"),
			Value: cdk.Fn_*Ref(jsii.String("AWS::AccountId")),
		},
	},
})

// The AWS IoT Thing Arn as a stack output
// The AWS IoT Thing Arn as a stack output
cdk.NewCfnOutput(app, jsii.String("ThingArn"), &CfnOutputProps{
	Value: fooThing.ThingArn,
})
// The AWS account unique endpoint for the MQTT data connection
// See API for other available public values that can be referenced
// The AWS account unique endpoint for the MQTT data connection
// See API for other available public values that can be referenced
cdk.NewCfnOutput(app, jsii.String("IotEndpoint"), &CfnOutputProps{
	Value: fooThing.DataAtsEndpointAddress,
})
```

</details><details>
  <summary><b>Python</b></summary>

**Installation:**

```shell
pip install cdklabs.cdk-aws-iot-thing-certificate-policy
```

[**API Reference**](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy/blob/main/doc/api-python.md)

**Example:**

```python
import aws_cdk as cdk
from cdklabs.cdk_aws_iot_thing_certificate_policy import (
    IotThingCertificatePolicy,
)

minimal_iot_policy = """{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": ["iot:Connect"],
      "Resource": "arn:aws:iot:{{region}}:{{account}}:client/{{thingname}}"
    },
    {
      "Effect": "Allow",
      "Action": ["iot:Publish"],
      "Resource": [
        "arn:aws:iot:{{region}}:{{account}}:topic/{{thingname}}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": ["iot:Subscribe"],
      "Resource": [
        "arn:aws:iot:{{region}}:{{account}}:topicfilter/{{thingname}}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": ["iot:Receive"],
      "Resource": [
        "arn:aws:iot:{{region}}:{{account}}:topic/{{thingname}}/*"
      ]
    }
  ]
}"""

app = cdk.App()

foo_thing = IotThingCertificatePolicy(
    app,
    "MyFooThing",
    thing_name="foo-thing",
    iot_policy_name="foo-iot-policy",
    iot_policy=minimal_iot_policy,
    encryption_algorithm="ECC",
    policy_parameter_mapping=[
        {
            "name": "region",
            "value": cdk.Fn.ref("AWS::Region")
        },
        {
            "name": "account",
            "value": cdk.Fn.ref("AWS::AccountId")
        }
    ],
)
cdk.CfnOutput(app, "ThingArn", value=foo_thing.thing_arn)
cdk.CfnOutput(app, "IotEndpoint", value=foo_thing.data_ats_endpoint_address)
```

</details><details>
  <summary><b>Java</b></summary>

**Installation:**

```shell
Coming Soon
```

[**API Reference**](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy/blob/main/doc/api-java.md)

**Example:** *Coming soon*

</details><details>
  <summary><b>C#</b></summary>

**Installation:**

```shell
dotnet add package Cdklabs.CdkAwsIotThingCertificatePolicy
```

[**API Reference**](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy/blob/main/doc/api-csharp.md)

**Example:** *coming soon*

</details><details>
  <summary><b>Go</b></summary>

**Installation:**

```shell
Coming soon
```

[**API Reference**](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy/blob/main/doc/api-go.md)

**Example:** *coming soon*

</details>
## AWS IoT Thing, Certificate, and Policy Construct Library

[![NPM](https://img.shields.io/npm/v/@cdklabs/cdk-aws-iot-thing-certificate-policy?label=npm+cdk+v2)](https://www.npmjs.com/package/@cdklabs/cdk-aws-iot-thing-certificate-policy)
[![PyPI](https://img.shields.io/pypi/v/cdklabs.cdk-aws-iot-thing-certificate-policy?label=pypi+cdk+v2)](https://pypi.org/project/cdklabs.cdk-aws-iot-thing-certificate-policy/)
[![Maven version](https://img.shields.io/maven-central/v/io.github.cdklabs/cdk-aws-iot-thing-certificate-policy?label=maven+cdk+v2)](https://central.sonatype.com/artifact/io.github.cdklabs/cdk-aws-iot-thing-certificate-policy)
[![NuGet version](https://img.shields.io/nuget/v/Cdklabs.CdkAwsIotThingCertificatePolicy?label=nuget+cdk+v2)](https://www.nuget.org/packages/Cdklabs.CdkAwsIotThingCertificatePolicy)
[![Go version](https://img.shields.io/github/go-mod/go-version/cdklabs/cdk-aws-iot-thing-certificate-policy-go?label=go+cdk+v2&&filename=cdklabscdkawsiotthingcertificatepolicy%2Fgo.mod)](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy-go)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue)](https://github.com/cdklabs/cdk-aws-iot-thing-certificate-policy/blob/main/LICENSE)

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

[![View on Construct Hub](https://constructs.dev/badge?package=%40cdklabs%2Fcdk-aws-iot-thing-certificate-policy)](https://constructs.dev/packages/@cdklabs/cdk-aws-iot-thing-certificate-policy)

An [L3 CDK construct](https://docs.aws.amazon.com/cdk/v2/guide/constructs.html#constructs_lib) to create and associate a singular AWS IoT Thing, Certificate, and IoT Policy. The construct also retrieves and returns AWS IoT account specific details such as the AWS IoT data endpoint and the AWS IoT Credential provider endpoint.

The certificate and its private key are stored as AWS Systems Manager Parameter Store parameters that can be retrieved via the AWS Console or programmatically via construct members.

## Installing

### TypeScript/JavaScript

```shell
npm install @cdklabs/cdk-aws-iot-thing-certificate-policy
```

### Python

```shell
pip install cdklabs.cdk-aws-iot-thing-certificate-policy
```

### Java

```xml
// add this to your pom.xml
<dependency>
    <groupId>io.github.cdklabs</groupId>
    <artifactId>cdk-aws-iot-thing-certificate-policy</artifactId>
    <version>0.0.0</version> // replace with version
</dependency>
```

### .NET

```plaintext
dotnet add package Cdklabs.CdkAwsIotThingCertificatePolicy --version X.X.X
```

## Go

```go
// Add this
import "github.com/cdklabs/cdk-aws-iot-thing-certificate-policy-go/cdklabscdkawsiotthingcertificatepolicy"
```

## Usage

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

package cdklabscdkawsiotthingcertificatepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-aws-iot-thing-certificate-policy-go/cdklabscdkawsiotthingcertificatepolicy/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-aws-iot-thing-certificate-policy-go/cdklabscdkawsiotthingcertificatepolicy/internal"
)

// Creates and associates an AWS IoT thing, AWS IoT certificate, and AWS IoT policy.
//
// It attaches the certificate to the thing and policy, and then stores the certificate
// and private key in AWS Systems Manager Parameter Store parameters for reference
// outside of the CloudFormation stack or by other constructs.
//
// Use this construct to create and delete a thing, certificate (principal), and IoT policy for
// testing or other singular uses. **Note:** Destroying this stack will fully detach and delete
// all created IoT resources including the AWS IoT thing, certificate, and policy.
// Experimental.
type IotThingCertificatePolicy interface {
	constructs.Construct
	// Arn of created AWS IoT Certificate.
	// Experimental.
	CertificateArn() *string
	// Fully qualified name in AWS Systems Manager Parameter Store of the certificate in `PEM` format.
	// Experimental.
	CertificatePemParameter() *string
	// Fully qualified domain name of the AWS IoT Credential provider endpoint specific to this AWS account and AWS region.
	// Experimental.
	CredentialProviderEndpointAddress() *string
	// Fully qualified domain name of the AWS IoT Core data plane endpoint specific to this AWS account and AWS region.
	// Experimental.
	DataAtsEndpointAddress() *string
	// Arn of created AWS IoT Policy.
	// Experimental.
	IotPolicyArn() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Fully qualified name in AWS Systems Manager Parameter Store of the certificate's private key in `PEM` format.
	// Experimental.
	PrivateKeySecretParameter() *string
	// Arn of created AWS IoT Thing.
	// Experimental.
	ThingArn() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotThingCertificatePolicy
type jsiiProxy_IotThingCertificatePolicy struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IotThingCertificatePolicy) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) CertificatePemParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePemParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) CredentialProviderEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) DataAtsEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAtsEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) IotPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iotPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) PrivateKeySecretParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeySecretParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotThingCertificatePolicy) ThingArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thingArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewIotThingCertificatePolicy(scope constructs.Construct, id *string, props *IotThingCertificatePolicyProps) IotThingCertificatePolicy {
	_init_.Initialize()

	if err := validateNewIotThingCertificatePolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotThingCertificatePolicy{}

	_jsii_.Create(
		"@cdklabs/cdk-aws-iot-thing-certificate-policy.IotThingCertificatePolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIotThingCertificatePolicy_Override(i IotThingCertificatePolicy, scope constructs.Construct, id *string, props *IotThingCertificatePolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/cdk-aws-iot-thing-certificate-policy.IotThingCertificatePolicy",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IotThingCertificatePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotThingCertificatePolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/cdk-aws-iot-thing-certificate-policy.IotThingCertificatePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotThingCertificatePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


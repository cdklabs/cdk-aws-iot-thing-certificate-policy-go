// Creates an AWS IoT thing, certificate, policy, and associates the three together
package cdklabscdkawsiotthingcertificatepolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/cdk-aws-iot-thing-certificate-policy.IotThingCertificatePolicy",
		reflect.TypeOf((*IotThingCertificatePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "certificatePemParameter", GoGetter: "CertificatePemParameter"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderEndpointAddress", GoGetter: "CredentialProviderEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "dataAtsEndpointAddress", GoGetter: "DataAtsEndpointAddress"},
			_jsii_.MemberProperty{JsiiProperty: "iotPolicyArn", GoGetter: "IotPolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeySecretParameter", GoGetter: "PrivateKeySecretParameter"},
			_jsii_.MemberProperty{JsiiProperty: "thingArn", GoGetter: "ThingArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IotThingCertificatePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-iot-thing-certificate-policy.IotThingCertificatePolicyProps",
		reflect.TypeOf((*IotThingCertificatePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/cdk-aws-iot-thing-certificate-policy.PolicyMapping",
		reflect.TypeOf((*PolicyMapping)(nil)).Elem(),
	)
}

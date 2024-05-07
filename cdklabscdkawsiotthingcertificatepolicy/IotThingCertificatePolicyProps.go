package cdklabscdkawsiotthingcertificatepolicy


// Properties for defining an AWS IoT thing, AWS IoT certificate, and AWS IoT policy.
// Experimental.
type IotThingCertificatePolicyProps struct {
	// The AWS IoT policy in JSON format to be created and attached to the certificate.
	//
	// This is a JSON string that uses [mustache-compatible](https://handlebarsjs.com/guide/)
	// template substitution to create the AWS IoT policy.
	// Default: - None.
	//
	// Experimental.
	IotPolicy *string `field:"required" json:"iotPolicy" yaml:"iotPolicy"`
	// Name of the AWS IoT Core policy to create.
	// Default: - None.
	//
	// Experimental.
	IotPolicyName *string `field:"required" json:"iotPolicyName" yaml:"iotPolicyName"`
	// Name of AWS IoT thing to create.
	// Default: - None.
	//
	// Experimental.
	ThingName *string `field:"required" json:"thingName" yaml:"thingName"`
	// Selects RSA or ECC private key and certificate generation.
	//
	// If not provided, `RSA` will be used.
	// Default: - RSA.
	//
	// Experimental.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Optional: A `PolicyMapping` object of parameters and values to be replaced if a [mustache-compatible](https://handlebarsjs.com/guide/) template is provided as the `iotPolicy` (see example). For each matching parameter in the policy template, the value will be used. If not provided, only the `{{thingname}}` mapping will be available for the `iotPolicy` template.
	// Default: - None.
	//
	// Experimental.
	PolicyParameterMapping *[]*PolicyMapping `field:"optional" json:"policyParameterMapping" yaml:"policyParameterMapping"`
}


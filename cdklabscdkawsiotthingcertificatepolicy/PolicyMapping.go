package cdklabscdkawsiotthingcertificatepolicy


// Policy substitutions provided as key-value pairs.
//
// Done this way to be JSII compatible.
// Experimental.
type PolicyMapping struct {
	// Name of substitution variable, e.g., `region` or `account`.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of substitution variable, e.g., `us-east-1` or `12345689012`.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}


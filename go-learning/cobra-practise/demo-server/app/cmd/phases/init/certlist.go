package phases

type Certificates []*Cert

type Cert struct {
	Name     string
	LongName string
}

func GetDefaultCertList() Certificates {
	return Certificates{
		CertRootCA(),
		CertAPIServer(),
		KubeletClient(),
	}
}

func CertRootCA() *Cert {
	return &Cert{
		Name:     "ca",
		LongName: "self-signed Kubernetes CA to provision identities for other Kubernetes components",
	}
}

func CertAPIServer() *Cert {
	return &Cert{
		Name:     "apiserver",
		LongName: "certificate for serving the Kubernetes API",
	}
}

func KubeletClient() *Cert {
	return &Cert{
		Name:     "apiserver-kubelet-client",
		LongName: "certificate for the API server to connect to kubelet",
	}
}

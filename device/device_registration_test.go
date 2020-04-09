package device_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("DeviceRegistration", func() {


	Context("Using Valid Device Info", func() {

		var identityServiceHost string = os.Getenv("FAIRX_IDENTITY_SERVICE_HOST")
		var identityServicePort string = os.Getenv("FAIRX_IDENTITY_SERVICE_PORT")


		It("Device registers itself with the Identity Service using valid data and authenticates successfully", func(){

			addr := fmt.Sprintf("%s:%s", identityServiceHost, identityServicePort)
			persona := testutils.EstablishValidStandardHumanPersona(addr, "alpha.fairx.io")
			Expect(len(persona.AuthTokens)).To(Equal(1))


		})


	})


})
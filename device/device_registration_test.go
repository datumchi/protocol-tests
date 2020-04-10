package device_test

import (
	"fmt"
	"github.com/datumchi/protocol-tests/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Device Registration", func() {


	Context("Using Valid DeviceInfo", func() {

		identityServiceAddr := fmt.Sprintf("%s:%s", os.Getenv("DATUMCHI_IDENTITY_SERVICE_HOST"), os.Getenv("DATUMCHI_IDENTITY_SERVICE_PORT"))
		persona := testutils.EstablishValidStandardHumanPersona(identityServiceAddr, "alpha.datumchi.com")


		It("Device registers itself with the Identity Service using valid data and authenticates successfully", func(){

			Expect(len(persona.AuthTokens)).To(Equal(1))


		})


	})


})
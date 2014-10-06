package metrix

import (
	"os"
	"testing"

	"github.com/dynport/dgtk/expect"
)

func TestEc2Metadata(t *testing.T) {
	expect := expect.New(t)
	f, err := os.Open("fixtures/ec2metadata.txt")
	expect(err).ToBeNil()
	defer f.Close()

	em := &Ec2Metadata{}
	err = em.Load(f)
	expect(err).ToBeNil()
	expect(em.AvailabilityZone).ToEqual("eu-west-1a")
	expect(em.InstanceId).ToEqual("i-72d0fa30")
	expect(em.AmiId).ToEqual("ami-905e81e7")
}

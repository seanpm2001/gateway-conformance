package ipns

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var YEAR_100 time.Time

func init() {
	YEAR_100 = time.Time(time.Date(2123, time.March, 01, 01, 0, 0, 801257000, time.UTC))
}

func TestExtractPath(t *testing.T) {
	path := "a/b/c/key_suffix.ipns-record"
	k, err := extractPubkeyFromPath(path)

	assert.Nil(t, err)
	assert.Equal(t, "key", k)

	path = "a/b/c/something.ipns-record"
	k, err = extractPubkeyFromPath(path)

	assert.Nil(t, err)
	assert.Equal(t, "something", k)

	path = "a/b/c/brokenrecord.ipns"
	k, err = extractPubkeyFromPath(path)

	assert.NotNil(t, err)
	assert.Equal(t, "", k)
}

func TestLoadIPNSRecord(t *testing.T) {
	path := "../../fixtures/t0124/k51qzi5uqu5dhjjqhpcuvdnskq4mz84a2xg1rpqzi6s5460q2612egkfjsk42x.ipns-record"
	ipns, err := OpenIPNSRecord(path)

	assert.Nil(t, err)
	assert.Equal(t, "k51qzi5uqu5dhjjqhpcuvdnskq4mz84a2xg1rpqzi6s5460q2612egkfjsk42x", ipns.Key())
	assert.Equal(t, ipns.Value(), "/ipfs/bafkreidfdrlkeq4m4xnxuyx6iae76fdm4wgl5d4xzsb77ixhyqwumhz244")
	assert.True(t, ipns.entry.Validity.After(YEAR_100))

	verify, err := ipns.Verify()
	assert.Nil(t, err)
	assert.True(t, verify)
}

func TestLoadTestRecord(t *testing.T) {
	// Test record created with (100 years):
	// ipfs name publish --allow-offline -t 876000h --key=self "/ipfs/$( echo "helloworld" | ipfs add --inline -q )"
	// ipfs routing get /ipns/${K} > ${K}.ipns-record

	path := "../../k51qzi5uqu5dgh7y9l90nqs6tvnzcm9erbt8fhzg3fu79p5qt9zb2izvfu51ki.ipns-record"
	ipns, err := OpenIPNSRecord(path)

	assert.Nil(t, err)
	assert.Equal(t, "k51qzi5uqu5dgh7y9l90nqs6tvnzcm9erbt8fhzg3fu79p5qt9zb2izvfu51ki", ipns.Key())
	assert.Equal(t, ipns.Value(), "/ipfs/bafyaaeykceeaeeqlnbswy3dpo5xxe3debimaw")
	assert.True(t, ipns.entry.Validity.After(YEAR_100))

	verify, err := ipns.Verify()
	assert.Nil(t, err)
	assert.True(t, verify)
}

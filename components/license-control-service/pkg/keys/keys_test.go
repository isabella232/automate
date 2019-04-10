package keys_test

import (
	"io/ioutil"
	"path"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chef/automate/components/license-control-service/pkg/keys"
)

var keysFixture = path.Join("testdata", "keys.json")

// In production code `Asset()` is defined in code generated by go-bindata,
// but in the test code we don't have that generated code available, so calls
// in the code under test for `Asset()` break because the function isn't
// defined. So, we define it here.
func Asset(string) ([]byte, error) {
	data, err := ioutil.ReadFile(keysFixture)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestLoadPublicKeys(t *testing.T) {
	keysData, err := Asset(keysFixture)
	if err != nil {
		log.Error(err)
		log.Fatalf("Couldn't load %s", keysFixture)
	}

	keyMap, err := keys.LoadPublicKeys(keysData)
	if err != nil {
		t.Fatal(err)
	}

	expectedKey := "e0df28c8bc68150edbfef98cd6b7dc439c1f80c7e7ef747893a6893a2f7b60f7"
	_, ok := keyMap[expectedKey]

	assert.True(t, ok, "We didn't find the checksum we expected in the keys map")
}

// This checks that the bindata compiled version of keys.bindata.go matches
// the files that are currently in keys/data so that we don't accidentally
// forget to generate them.
// If this test fails `make generate` should get you fixed.
func TestBindataKeys(t *testing.T) {
	assert.Equal(t, 1, len(keys.AssetNames()), "expecting only keys.json and got more assets")

	for _, asset := range keys.AssetNames() {
		compiled := keys.MustAsset(asset)

		onDisk, err := ioutil.ReadFile(asset)
		require.Nil(t, err, "read asset %v from disk", "foo")

		assert.Equal(t, onDisk, compiled, "compare asset %v on disk vs compiled-in", asset)
	}
}

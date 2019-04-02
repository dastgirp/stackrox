// +build !release

package singleton

import (
	"time"

	"github.com/stackrox/rox/pkg/license/validator"
	"github.com/stackrox/rox/pkg/timeutil"
	"github.com/stackrox/rox/pkg/utils"
)

func init() {
	utils.Must(
		validatorInstance.RegisterSigningKey(
			"stackrox-dev/licensing-dev/dev-license/2",
			validator.EC256,
			[]byte{
				0x30, 0x59, 0x30, 0x13, 0x06, 0x07, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x02,
				0x01, 0x06, 0x08, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x03, 0x01, 0x07, 0x03,
				0x42, 0x00, 0x04, 0x2e, 0x93, 0xd3, 0xdc, 0xa9, 0x93, 0x8c, 0xe2, 0xf4,
				0xf9, 0x46, 0xdb, 0x30, 0x8a, 0xc5, 0x0e, 0xcf, 0x52, 0x7d, 0xfc, 0x48,
				0x49, 0x00, 0xd9, 0xf5, 0x20, 0xd0, 0x02, 0x5b, 0xb3, 0xb7, 0x4c, 0x10,
				0x82, 0xbb, 0xbf, 0x25, 0x0d, 0xc8, 0xe6, 0x90, 0xf4, 0xea, 0x8f, 0x3d,
				0x0b, 0xb6, 0xcf, 0xac, 0xa7, 0xcb, 0x3f, 0xe0, 0xd5, 0x08, 0xd1, 0x44,
				0x9f, 0x9f, 0x2a, 0x4b, 0xf8, 0xd8, 0x70,
			},
			validator.SigningKeyRestrictions{
				EarliestNotValidBefore:                  timeutil.MustParse(time.RFC3339, "2019-04-01T00:00:00Z"),
				LatestNotValidAfter:                     timeutil.MustParse(time.RFC3339, "2019-04-12T00:00:00Z"),
				MaxDuration:                             14 * 24 * time.Hour,
				AllowOffline:                            true,
				MaxNodeLimit:                            50,
				AllowNoBuildFlavorRestriction:           true,
				AllowNoDeploymentEnvironmentRestriction: true,
			}),
	)
}

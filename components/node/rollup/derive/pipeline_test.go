package derive

import "github.com/wemixkanvas/kanvas/components/node/testutils"

var _ Engine = (*testutils.MockEngine)(nil)

var _ L1Fetcher = (*testutils.MockL1Source)(nil)

var _ Metrics = (*testutils.TestDerivationMetrics)(nil)
